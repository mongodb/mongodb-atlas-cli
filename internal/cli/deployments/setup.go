// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package deployments

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/briandowns/spinner"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/setup"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/workflows"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/compass"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/container"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mongodbclient"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mongosh"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/search"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/templatewriter"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/vscode"
	"github.com/spf13/cobra"
)

const (
	internalMongodPort         = 27017
	mdb7                       = "7.0"
	mdb8                       = "8.0"
	defaultSettings            = "default"
	customSettings             = "custom"
	cancelSettings             = "cancel"
	skipConnect                = "skip"
	autoassignPort             = "autoassign"
	spinnerSpeed               = 100 * time.Millisecond
	steps                      = 3
	atlasM2                    = "M2"
	atlasM5                    = "M5"
	deprecateMessageSharedTier = "The '%s' tier is deprecated. For the migration guide and timeline, visit: https://dochub.mongodb.org/core/flex-migration.\n"
)

var (
	errInvalidInit              = errors.New("invalid --initdb flag")
	errInitMustBeDir            = fmt.Errorf("%w: must be a directory", errInvalidInit)
	errCancel                   = errors.New("the setup was cancelled")
	ErrDeploymentExists         = errors.New("deployment already exists")
	errMustBeInt                = errors.New("you must specify an integer")
	errPortOutOfRange           = errors.New("you must specify a port within the range 1..65535")
	errPortNotAvailable         = errors.New("the port is unavailable")
	errFlagTypeRequired         = fmt.Errorf("the --%s flag is required when the --%s flag is set", flag.TypeFlag, flag.Force)
	errFlagsTypeAndAuthRequired = fmt.Errorf("the --%s, --%s and --%s flags are required when the --%s and --%s flags are set",
		flag.TypeFlag, flag.Username, flag.Password, flag.Force, flag.BindIPAll)
	errInvalidDeploymentType      = errors.New("the deployment type is invalid")
	errIncompatibleDeploymentType = fmt.Errorf("the --%s flag applies only to LOCAL deployments", flag.BindIPAll)
	errInvalidMongoDBVersion      = errors.New("the mongodb version is invalid")
	errUnsupportedConnectWith     = fmt.Errorf("the --%s flag is unsupported", flag.ConnectWith)
	errFlagUsernameRequired       = fmt.Errorf("the --%s is required to enable authentication when --%s flag is set",
		flag.Username, flag.BindIPAll)
	errFailedToDownloadImage   = errors.New("failed to download the MongoDB image")
	errDeploymentUnhealthy     = errors.New("the deployment is unhealthy")
	errDeploymentNoHealthCheck = errors.New("the deployment does not have a healthcheck")
	errHealthCheckTimeout      = errors.New("timed out waiting for the deployment to be healthy")
	errDownloadImage           = errors.New("image download failed")
	errInspectHealthCheck      = errors.New("inspect healthcheck failed")
	errQueryHealthCheckStatus  = errors.New("query healthcheck status failed")
	errConfigContainer         = errors.New("container configuration failed")
	errRunContainer            = errors.New("container run failed")
	errListContainer           = errors.New("listing containers failed")
	settingOptions             = []string{defaultSettings, customSettings, cancelSettings}
	settingsDescription        = map[string]string{
		defaultSettings: "With default settings",
		customSettings:  "With custom settings",
		cancelSettings:  "Cancel setup",
	}
	connectWithOptions     = []string{options.MongoshConnect, options.CompassConnect, options.VsCodeConnect, skipConnect}
	connectWithDescription = map[string]string{
		options.MongoshConnect: "MongoDB Shell",
		options.CompassConnect: "MongoDB Compass",
		options.VsCodeConnect:  "MongoDB for VsCode",
		skipConnect:            "Skip Connection",
	}
	mdbVersions = []string{mdb7, mdb8}
)

type SetupOpts struct {
	options.DeploymentOpts
	cli.OutputOpts
	cli.ProjectOpts
	cli.InputOpts
	mongodbClient mongodbclient.MongoDBClient
	settings      string
	connectWith   string
	force         bool
	bindIPAll     bool
	mongodIP      string
	initdb        string
	s             *spinner.Spinner
	atlasSetup    *setup.Opts
}

func (opts *SetupOpts) initMongoDBClient() error {
	opts.mongodbClient = mongodbclient.NewClient()
	return nil
}

func (opts *SetupOpts) logStepStarted(msg string, currentStep int, totalSteps int) {
	fullMessage := fmt.Sprintf("%d/%d: %s", currentStep, totalSteps, msg)
	_, _ = log.Warningln(fullMessage)
	opts.start()
}

func (opts *SetupOpts) downloadImage(ctx context.Context, currentStep int, steps int) error {
	opts.logStepStarted("Downloading the latest MongoDB image to your local environment...", currentStep, steps)
	defer opts.stop()

	err := opts.ContainerEngine.ImagePull(ctx, opts.MongodDockerImageName())
	if err == nil {
		return nil
	}

	// In case we already have an image present and the download fails, we can continue with the existing image
	images, _ := opts.ContainerEngine.ImageList(ctx, opts.MongodDockerImageName())
	if len(images) != 0 {
		return nil
	}

	return errFailedToDownloadImage
}

func (opts *SetupOpts) startEnvironment(ctx context.Context, currentStep int, steps int) error {
	opts.logStepStarted("Starting your local environment...", currentStep, steps)
	defer opts.stop()

	containers, err := opts.GetLocalContainers(ctx)
	if err != nil {
		return fmt.Errorf("%w: %w", errListContainer, err)
	}

	return opts.validateLocalDeploymentsSettings(containers)
}

func (opts *SetupOpts) createLocalDeployment(ctx context.Context) error {
	currentStep := 1

	_, _ = log.Warningf("Creating your cluster %s\n", opts.DeploymentName)

	// verify that the host meets the minimum requirements, if not, print a warning
	if err := opts.ValidateMinimumRequirements(); err != nil {
		return err
	}

	// containers check
	if err := opts.startEnvironment(ctx, currentStep, steps); err != nil {
		return err
	}
	currentStep++

	// always download the latest image
	if err := opts.downloadImage(ctx, currentStep, steps); err != nil {
		return fmt.Errorf("%w: %w", errDownloadImage, err)
	}
	currentStep++

	// create local deployment
	opts.logStepStarted(fmt.Sprintf("Creating your deployment %s...", opts.DeploymentName), currentStep, steps)
	defer opts.stop()

	if err := opts.configureContainer(ctx); err != nil {
		return fmt.Errorf("%w: %w", errConfigContainer, err)
	}

	return nil
}

func (opts *SetupOpts) configureContainer(ctx context.Context) error {
	envVars := map[string]string{
		"TOOL": "ATLASCLI",
	}

	if log.IsDebugLevel() {
		envVars["RUNNER_LOG_FILE"] = "/dev/stdout"
	}

	if !config.TelemetryEnabled() {
		envVars["DO_NOT_TRACK"] = "1"
	}

	if opts.IsAuthEnabled() {
		envVars["MONGODB_INITDB_ROOT_USERNAME"] = opts.DBUsername
		envVars["MONGODB_INITDB_ROOT_PASSWORD"] = opts.DBUserPassword
	}

	flags := container.RunFlags{}
	flags.Detach = pointer.Get(true)
	flags.Name = pointer.Get(opts.LocalMongodHostname())
	flags.Hostname = pointer.Get(opts.LocalMongodHostname())
	flags.Env = envVars
	flags.BindIPAll = &opts.bindIPAll
	flags.IP = &opts.mongodIP
	flags.Ports = []container.PortMapping{{HostPort: opts.Port, ContainerPort: internalMongodPort}}
	if opts.initdb != "" {
		flags.Volumes = []container.VolumeMapping{
			{
				HostPath:      opts.initdb,
				ContainerPath: "/docker-entrypoint-initdb.d",
			},
		}
	}
	healthCheck, err := opts.ContainerEngine.ImageHealthCheck(ctx, opts.MongodDockerImageName())
	if err != nil {
		return fmt.Errorf("%w: %w", errInspectHealthCheck, err)
	}

	// Temporary fix until https://github.com/containers/podman/issues/18904 is closed
	if healthCheck == nil {
		const HealthcheckInterval = 30
		const HealthStartPeriod = 1
		const HealthTimeout = 30
		const HealthRetries = 3

		flags.HealthCmd = &[]string{"/usr/local/bin/runner", "healthcheck"}
		flags.HealthInterval = pointer.Get(HealthcheckInterval * time.Second)
		flags.HealthStartPeriod = pointer.Get(HealthStartPeriod * time.Second)
		flags.HealthTimeout = pointer.Get(HealthTimeout * time.Second)
		flags.HealthRetries = pointer.Get(HealthRetries)
	}

	_, err = opts.ContainerEngine.ContainerRun(ctx, opts.MongodDockerImageName(), &flags)
	if err != nil {
		return fmt.Errorf("%w: %w", errRunContainer, err)
	}

	// This can be a high number because the container will become unhealthy before the 10 minutes is reached
	const healthyDeploymentTimeout = 10 * time.Minute

	err = opts.WaitForHealthyDeployment(ctx, healthyDeploymentTimeout)
	if err != nil && log.IsDebugLevel() {
		logs, err := opts.ContainerEngine.ContainerLogs(ctx, opts.LocalMongodHostname())
		if err != nil {
			_, _ = log.Debugf("Container unhealthy (hostname=%s), failed to get container logs : %s\n", opts.LocalMongodHostname(), err)
		} else {
			_, _ = log.Debugf("Container unhealthy (hostname=%s), container logs:\n", opts.LocalMongodHostname())
			for _, logLine := range logs {
				_, _ = log.Debugln(logLine)
			}
		}
	}

	return err
}

func (opts *SetupOpts) WaitForHealthyDeployment(ctx context.Context, duration time.Duration) error {
	start := time.Now()

	for {
		if time.Since(start) > duration {
			return errHealthCheckTimeout
		}

		status, err := opts.ContainerEngine.ContainerHealthStatus(ctx, opts.LocalMongodHostname())
		if err != nil {
			return fmt.Errorf("%w: %w", errQueryHealthCheckStatus, err)
		}

		switch status {
		case container.DockerHealthcheckStatusHealthy:
			return nil
		case container.DockerHealthcheckStatusUnhealthy:
			return errDeploymentUnhealthy
		case container.DockerHealthcheckStatusNone:
			return errDeploymentNoHealthCheck
		case container.DockerHealthcheckStatusStarting:
			time.Sleep(1 * time.Second)
		}
	}
}

func (opts *SetupOpts) validateLocalDeploymentsSettings(containers []container.Container) error {
	mongodContainerName := opts.LocalMongodHostname()
	for _, c := range containers {
		for _, n := range c.Names {
			if n == mongodContainerName {
				return fmt.Errorf("%w: \"%s\", state:\"%s\"", ErrDeploymentExists, opts.DeploymentName, c.State)
			}
		}
	}

	return nil
}

func (opts *SetupOpts) promptSettings() error {
	p := &survey.Select{
		Message: "How do you want to set up your local Atlas deployment?",
		Options: settingOptions,
		Default: opts.settings,
		Description: func(value string, _ int) string {
			return settingsDescription[value]
		},
	}

	return telemetry.TrackAskOne(p, &opts.settings, nil)
}

func (opts *SetupOpts) generateDeploymentName() {
	opts.DeploymentName = fmt.Sprintf("local%v", rand.Intn(10000)) //nolint // no need for crypto here
}

func (opts *SetupOpts) promptDeploymentName() error {
	p := &survey.Input{
		Message: "Deployment Name [You can't change this value later]",
		Help:    usage.ClusterName,
		Default: opts.DeploymentName,
	}

	return telemetry.TrackAskOne(p, &opts.DeploymentName, survey.WithValidator(func(ans any) error {
		name, _ := ans.(string)
		return options.ValidateDeploymentName(name)
	}))
}

func (opts *SetupOpts) promptMdbVersion() error {
	p := &survey.Select{
		Message: "Major MongoDB Version",
		Options: mdbVersions,
		Default: opts.MdbVersion,
		Help:    "Major MongoDB Version for the deployment. Atlas CLI applies the latest minor version available.",
	}

	return telemetry.TrackAskOne(p, &opts.MdbVersion, nil)
}

func checkPort(p int) error {
	server, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", p))
	if err != nil {
		return fmt.Errorf("%w: %d", errPortNotAvailable, p)
	}
	_ = server.Close()
	return nil
}

func validatePort(p int) error {
	if p <= 0 || p > 65535 {
		return errPortOutOfRange
	}
	return checkPort(p)
}

func (opts *SetupOpts) promptPort() error {
	exportPort := autoassignPort
	if opts.Port != 0 {
		exportPort = strconv.Itoa(opts.Port)
	}

	p := &survey.Input{
		Message: "Specify a port",
		Default: exportPort,
	}

	err := telemetry.TrackAskOne(p, &exportPort, survey.WithValidator(func(ans any) error {
		input, _ := ans.(string)
		if input == autoassignPort {
			return nil
		}
		value, err := strconv.Atoi(input)
		if err != nil {
			return errMustBeInt
		}

		return validatePort(value)
	}))

	if err != nil {
		return err
	}

	if exportPort != autoassignPort {
		if opts.Port, err = strconv.Atoi(exportPort); err != nil {
			return err
		}
	}

	return nil
}

func (opts *SetupOpts) validateDeploymentTypeFlag() error {
	if opts.DeploymentType == "" && opts.force {
		return errFlagTypeRequired
	}

	if !opts.IsLocalDeploymentType() && opts.bindIPAll {
		return errIncompatibleDeploymentType
	}

	if opts.DeploymentType != "" && !strings.EqualFold(opts.DeploymentType, options.AtlasCluster) && !strings.EqualFold(opts.DeploymentType, options.LocalCluster) {
		return fmt.Errorf("%w: %s", errInvalidDeploymentType, opts.DeploymentType)
	}

	return nil
}

func (opts *SetupOpts) validateBindIPAllFlag() error {
	if !opts.bindIPAll {
		return nil
	}

	if opts.force && (opts.DeploymentType == "" || opts.DBUsername == "" || opts.DBUserPassword == "") {
		return errFlagsTypeAndAuthRequired
	}

	if opts.DBUsername == "" {
		return errFlagUsernameRequired
	}

	return nil
}

func (opts *SetupOpts) validateFlags() error {
	if err := opts.validateDeploymentTypeFlag(); err != nil {
		return err
	}

	if opts.DeploymentName != "" {
		if err := options.ValidateDeploymentName(opts.DeploymentName); err != nil {
			return err
		}
	}

	if opts.MdbVersion != "" && !slices.Contains(mdbVersions, opts.MdbVersion) {
		return fmt.Errorf("%w: %s", errInvalidMongoDBVersion, opts.MdbVersion)
	}

	if opts.Port != 0 {
		if err := validatePort(opts.Port); err != nil {
			return err
		}
	}

	if opts.connectWith != "" && !search.StringInSliceFold(connectWithOptions, opts.connectWith) {
		return fmt.Errorf("%w: %s", errUnsupportedConnectWith, opts.connectWith)
	}

	if opts.initdb != "" {
		info, err := os.Stat(opts.initdb)
		if err != nil {
			return fmt.Errorf("%w: %w", errInvalidInit, err)
		}
		if !info.IsDir() {
			return errInitMustBeDir
		}
		opts.initdb, err = filepath.Abs(opts.initdb)
		if err != nil {
			return fmt.Errorf("%w: %w", errInvalidInit, err)
		}
	}

	return opts.validateBindIPAllFlag()
}

func (opts *SetupOpts) promptLocalAdminPassword() error {
	if !opts.IsTerminalInput() {
		_, err := fmt.Fscanln(opts.InReader, &opts.DBUserPassword)
		return err
	}

	p := &survey.Password{
		Message: "Password for authenticating to local deployment",
	}
	return telemetry.TrackAskOne(p, &opts.DBUserPassword)
}

func (opts *SetupOpts) setDefaultSettings() error {
	opts.settings = defaultSettings
	defaultValuesSet := false

	if opts.DeploymentName == "" {
		opts.generateDeploymentName()
		defaultValuesSet = true
	}

	if opts.MdbVersion == "" {
		opts.MdbVersion = mdb8
		defaultValuesSet = true
	}

	if opts.Port == 0 {
		defaultValuesSet = true
	}

	if defaultValuesSet {
		if err := templatewriter.Print(os.Stderr, `
[Default Settings]
Deployment Name	{{.DeploymentName}}
MongoDB Version	{{.MdbVersion}}

`, opts); err != nil {
			return err
		}
		if !opts.force {
			if err := opts.promptSettings(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (opts *SetupOpts) promptConnect() error {
	p := &survey.Select{
		Message: fmt.Sprintf("How would you like to connect to %s?", opts.DeploymentName),
		Options: connectWithOptions,
		Description: func(value string, _ int) string {
			return connectWithDescription[value]
		},
	}

	return telemetry.TrackAskOne(p, &opts.connectWith, nil)
}

func (opts *SetupOpts) runConnectWith(cs string) error {
	if opts.connectWith == "" {
		if opts.force {
			opts.connectWith = skipConnect
		} else {
			if err := opts.promptConnect(); err != nil {
				return err
			}
		}
	}

	switch opts.connectWith {
	case skipConnect:
		_, _ = fmt.Fprintln(os.Stderr, "connection skipped")
	case options.CompassConnect:
		if !compass.Detect() {
			return compass.ErrCompassNotInstalled
		}
		if _, err := log.Warningln("Launching MongoDB Compass..."); err != nil {
			return err
		}
		return compass.Run("", "", cs)
	case options.MongoshConnect:
		if !mongosh.Detect() {
			return mongosh.ErrMongoshNotInstalled
		}
		return mongosh.Run("", "", cs)
	case options.VsCodeConnect:
		if !vscode.Detect() {
			return vscode.ErrVsCodeCliNotInstalled
		}
		if _, err := log.Warningln("Launching VsCode..."); err != nil {
			return err
		}
		return vscode.SaveConnection(cs, opts.DeploymentName, opts.DeploymentType)
	}

	return nil
}

func (opts *SetupOpts) validateAndPrompt() error {
	if err := opts.validateFlags(); err != nil {
		return err
	}

	if err := opts.ValidateAndPromptDeploymentType(); err != nil {
		return err
	}

	// Defer prompts to Atlas command
	if opts.IsAtlasDeploymentType() {
		return nil
	}

	if opts.DBUsername != "" && opts.DBUserPassword == "" {
		if err := opts.promptLocalAdminPassword(); err != nil {
			return err
		}
	}

	if err := opts.setDefaultSettings(); err != nil {
		return err
	}

	switch opts.settings {
	case cancelSettings:
		return errCancel
	case customSettings:
		if err := opts.promptDeploymentName(); err != nil {
			return err
		}

		if err := opts.promptMdbVersion(); err != nil {
			return err
		}

		if err := opts.promptPort(); err != nil {
			return err
		}
	}

	return nil
}

func (opts *SetupOpts) runLocal(ctx context.Context) error {
	if err := opts.LocalDeploymentPreRun(ctx); err != nil {
		return err
	}

	if err := opts.createLocalDeployment(ctx); err != nil {
		// in case the deployment already exists we shouldn't delete it
		if !errors.Is(err, ErrDeploymentExists) {
			_ = opts.RemoveLocal(ctx)
		}
		return err
	}

	cs, err := opts.ConnectionString(ctx)
	if err != nil {
		return err
	}

	_, _ = log.Warningln("Deployment created!")
	_, _ = fmt.Fprintf(opts.OutWriter, `Connection string: "%s"
`, cs)
	_, _ = log.Warningln("")

	return opts.runConnectWith(cs)
}

func (opts *SetupOpts) runAtlas(ctx context.Context) error {
	s := setup.Builder()

	// remove global flags and unknown flags
	var newArgs []string
	_, _ = log.Debugf("Removing flags and args from original args %s\n", os.Args)

	flagstoRemove := map[string]string{
		flag.TypeFlag: "1",
	}

	newArgs, err := workflows.RemoveFlagsAndArgs(flagstoRemove, map[string]bool{opts.DeploymentName: true}, os.Args)
	if err != nil {
		return err
	}

	// replace deployment name with cluster name
	if opts.DeploymentName != "" {
		newArgs = append(newArgs, "--"+flag.ClusterName, opts.DeploymentName)
	}

	// update args
	s.SetArgs(newArgs)

	// run atlas setup
	_, _ = log.Debugf("Starting to run atlas setup with args %s\n", newArgs)
	_, err = s.ExecuteContextC(ctx)
	return err
}

func (opts *SetupOpts) Run(ctx context.Context) error {
	if err := opts.validateAndPrompt(); err != nil {
		if errors.Is(err, errCancel) {
			_, _ = log.Warningln(err)
			return nil
		}

		return err
	}

	if strings.EqualFold(options.LocalCluster, opts.DeploymentType) {
		return opts.runLocal(ctx)
	}

	return opts.runAtlas(ctx)
}

func (opts *SetupOpts) PostRun() {
	opts.DeploymentTelemetry.AppendDeploymentType()
	opts.DeploymentTelemetry.AppendDeploymentUUID()
}

func (opts *SetupOpts) validateTier() error {
	opts.atlasSetup.Tier = strings.ToUpper(opts.atlasSetup.Tier)
	if opts.atlasSetup.Tier == atlasM2 || opts.atlasSetup.Tier == atlasM5 {
		_, _ = fmt.Fprintf(os.Stderr, deprecateMessageSharedTier, opts.atlasSetup.Tier)
	}
	return nil
}

// SetupBuilder builds a cobra.Command that can run as:
// atlas deployments setup.
func SetupBuilder() *cobra.Command {
	opts := &SetupOpts{
		settings:   defaultSettings,
		atlasSetup: &setup.Opts{},
	}
	cmd := &cobra.Command{
		Use:     "setup [deploymentName]",
		Short:   "Create a local deployment.",
		Long:    "To learn more about local atlas deployments, see https://www.mongodb.com/docs/atlas/cli/current/atlas-cli-deploy-local/",
		Args:    require.MaximumNArgs(1),
		GroupID: "all",
		Annotations: map[string]string{
			"deploymentNameDesc": "Name of the deployment that you want to set up.",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				opts.DeploymentName = args[0]
			}

			opts.force = opts.atlasSetup.Confirm
			opts.DBUsername = opts.atlasSetup.DBUsername
			opts.DBUserPassword = opts.atlasSetup.DBUserPassword

			return opts.PreRunE(
				opts.validateTier,
				opts.InitOutput(cmd.OutOrStdout(), ""),
				opts.InitInput(cmd.InOrStdin()),
				opts.InitStore(cmd.Context(), cmd.OutOrStdout()),
				opts.initMongoDBClient,
			)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return opts.Run(cmd.Context())
		},
		PostRun: func(_ *cobra.Command, _ []string) {
			opts.PostRun()
		},
	}

	// Local and Atlas
	cmd.Flags().StringVar(&opts.DeploymentType, flag.TypeFlag, "", usage.DeploymentTypeSetup)
	cmd.Flags().StringVar(&opts.MdbVersion, flag.MDBVersion, "", usage.DeploymentMDBVersion)
	cmd.Flags().StringVar(&opts.connectWith, flag.ConnectWith, "", usage.ConnectWithSetup)

	// Local only
	cmd.Flags().IntVar(&opts.Port, flag.Port, 0, usage.MongodPort)
	cmd.Flags().BoolVar(&opts.bindIPAll, flag.BindIPAll, false, usage.BindIPAll)
	cmd.Flags().StringVar(&opts.initdb, flag.InitDB, "", usage.InitDB)

	// Atlas only
	opts.atlasSetup.SetupAtlasFlags(cmd)
	opts.atlasSetup.SetupFlowFlags(cmd)
	cmd.Flags().Lookup(flag.Region).Usage = usage.DeploymentRegion
	cmd.Flags().Lookup(flag.Tag).Usage = usage.DeploymentTag
	cmd.Flags().Lookup(flag.Tier).Usage = usage.DeploymentTier
	cmd.Flags().Lookup(flag.EnableTerminationProtection).Usage = usage.EnableTerminationProtectionForDeployment
	cmd.Flags().Lookup(flag.SkipSampleData).Usage = usage.SkipSampleDataDeployment
	cmd.Flags().Lookup(flag.Force).Usage = usage.ForceDeploymentsSetup

	_ = cmd.RegisterFlagCompletionFunc(flag.MDBVersion, func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return mdbVersions, cobra.ShellCompDirectiveDefault
	})
	_ = cmd.RegisterFlagCompletionFunc(flag.TypeFlag, func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return options.DeploymentTypeOptions, cobra.ShellCompDirectiveDefault
	})
	_ = cmd.RegisterFlagCompletionFunc(flag.ConnectWith, func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return connectWithOptions, cobra.ShellCompDirectiveDefault
	})
	return cmd
}

func (opts *SetupOpts) start() {
	if opts.IsTerminal() {
		opts.s = spinner.New(spinner.CharSets[9], spinnerSpeed)
		_ = opts.s.Color("cyan", "bold")
		opts.s.Start()
	}
}

func (opts *SetupOpts) stop() {
	if opts.IsTerminal() {
		opts.s.Stop()
	}
}

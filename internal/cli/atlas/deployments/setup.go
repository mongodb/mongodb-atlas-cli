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
	_ "embed"
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/briandowns/spinner"
	"github.com/google/uuid"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/compass"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/internal/mongodbclient"
	"github.com/mongodb/mongodb-atlas-cli/internal/mongosh"
	"github.com/mongodb/mongodb-atlas-cli/internal/podman"
	"github.com/mongodb/mongodb-atlas-cli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/internal/templatewriter"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	internalMongodPort = 27017
	internalMongotPort = 27027
	localCluster       = "local"
	atlasCluster       = "atlas"
	mdb6               = "6.0"
	mdb7               = "7.0"
	replicaSetName     = "rs-localdev"
	defaultSettings    = "default"
	customSettings     = "custom"
	skipSettings       = "skip"
	compassConnect     = "compass"
	mongoshConnect     = "mongosh"
	skipConnect        = "skip"
	spinnerSpeed       = 100 * time.Millisecond
	shortStepCount     = 2
)

var (
	errSkip                         = errors.New("setup skipped")
	errMustBeInt                    = errors.New("input must be an integer")
	errPortOutOfRange               = errors.New("port must within the range 1..65535")
	errPortNotAvailable             = errors.New("port not available")
	errFlagTypeRequired             = errors.New("flag --type is required when --force is set")
	errInvalidDeploymentType        = errors.New("invalid deployment type")
	errInvalidMongoDBVersion        = errors.New("invalid mongodb version")
	errUnsupportedConnectWith       = errors.New("flag --connectWith unsupported")
	errDeploymentTypeNotImplemented = errors.New("deployment type not implemented")
	deploymentTypeOptions           = []string{localCluster, atlasCluster}
	deploymentTypeDescription       = map[string]string{
		localCluster: "Local Database",
		atlasCluster: "Atlas Database",
	}
	settingOptions      = []string{defaultSettings, customSettings, skipSettings}
	settingsDescription = map[string]string{
		defaultSettings: "With default settings",
		customSettings:  "With custom settings",
		skipSettings:    "Skip set up",
	}
	connectWithOptions     = []string{mongoshConnect, compassConnect, skipConnect}
	connectWithDescription = map[string]string{
		mongoshConnect: "MongoDB Shell",
		compassConnect: "MongoDB Compass",
		skipConnect:    "Skip Connection",
	}
	mdbVersions = []string{mdb7, mdb6}
	//go:embed scripts/start_mongod.sh
	mongodStartScript []byte
	//go:embed scripts/start_mongot.sh
	mongotStartScript []byte
)

type SetupOpts struct {
	options.DeploymentOpts
	cli.OutputOpts
	cli.GlobalOpts
	podmanClient  podman.Client
	mongodbClient mongodbclient.MongoDBClient
	settings      string
	connectWith   string
	force         bool
	s             *spinner.Spinner
}

func (opts *SetupOpts) initPodmanClient() error {
	opts.podmanClient = podman.NewClient(log.IsDebugLevel(), log.Writer())
	return nil
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

func (opts *SetupOpts) downloadImagesIfNotAvailable(ctx context.Context, currentStep int, steps int) error {
	opts.logStepStarted("Downloading MongoDB binaries to your local environment...", currentStep, steps)
	defer opts.stop()

	var mongodImages []*podman.Image
	var mongotImages []*podman.Image
	var err error

	if mongodImages, err = opts.podmanClient.ListImages(ctx, opts.MongodDockerImageName()); err != nil {
		return err
	}

	if len(mongodImages) == 0 {
		if _, err = opts.podmanClient.PullImage(ctx, opts.MongodDockerImageName()); err != nil {
			return err
		}
	}

	if mongotImages, err = opts.podmanClient.ListImages(ctx, opts.MongodDockerImageName()); err != nil {
		return err
	}

	if len(mongotImages) == 0 {
		if _, err := opts.podmanClient.PullImage(ctx, options.MongotDockerImageName); err != nil {
			return err
		}
	}

	return nil
}

func (opts *SetupOpts) setupPodman(ctx context.Context, currentStep int, steps int) error {
	opts.logStepStarted("Downloading and completing configuration...", currentStep, steps)
	defer opts.stop()

	return opts.podmanClient.Ready(ctx)
}

func (opts *SetupOpts) startEnvironment(ctx context.Context, currentStep int, steps int) error {
	opts.logStepStarted("Starting your local environment...", currentStep, steps)
	defer opts.stop()

	containers, errList := opts.podmanClient.ListContainers(ctx, options.MongodHostnamePrefix)
	if errList != nil {
		return errList
	}

	if err := opts.validateLocalDeploymentsSettings(containers); err != nil {
		return err
	}

	if _, err := opts.podmanClient.CreateNetwork(ctx, opts.LocalNetworkName()); err != nil {
		return err
	}
	return nil
}

func (opts *SetupOpts) planSteps(ctx context.Context) (steps int, needPodmanSetup bool, needToPullImages bool) {
	steps = 2
	needPodmanSetup = false
	needToPullImages = false

	setupState := opts.podmanClient.Diagnostics(ctx)

	if !setupState.MachineFound || setupState.MachineState != "running" {
		steps++
		needPodmanSetup = true
	}

	foundMongod := false
	foundMongot := false
	for _, image := range setupState.Images {
		foundMongod = foundMongod || image == opts.MongodDockerImageName()
		foundMongot = foundMongot || image == options.MongotDockerImageName
	}

	if !foundMongod || !foundMongot {
		steps++
		needToPullImages = true
	}
	return steps, needPodmanSetup, needToPullImages
}

func (opts *SetupOpts) createLocalDeployment(ctx context.Context) error {
	if err := podman.Installed(); err != nil {
		return err
	}

	steps, needPodmanSetup, needToPullImages := opts.planSteps(ctx)
	currentStep := 1
	longWaitWarning := ""
	if steps > shortStepCount {
		longWaitWarning = " [this might take several minutes]"
	}

	_, _ = log.Warningf("Creating your cluster %s%s\n", opts.DeploymentName, longWaitWarning)

	// podman config
	if needPodmanSetup {
		if err := opts.setupPodman(ctx, currentStep, steps); err != nil {
			return err
		}
		currentStep++
	}

	// containers check and network init
	if err := opts.startEnvironment(ctx, currentStep, steps); err != nil {
		return err
	}
	currentStep++

	// pull images if not available
	if needToPullImages {
		if err := opts.downloadImagesIfNotAvailable(ctx, currentStep, steps); err != nil {
			return err
		}
		currentStep++
	}

	// create local deployment
	opts.logStepStarted(fmt.Sprintf("Creating your cluster %s...", opts.DeploymentName), currentStep, steps)
	defer opts.stop()
	keyFileContents := base64.URLEncoding.EncodeToString([]byte(uuid.NewString()))

	if err := opts.configureMongod(ctx, keyFileContents); err != nil {
		return err
	}

	return opts.configureMongot(ctx, keyFileContents)
}

func (opts *SetupOpts) configureMongod(ctx context.Context, keyFileContents string) error {
	mongodDataVolume := opts.LocalMongodDataVolume()
	if _, err := opts.podmanClient.CreateVolume(ctx, mongodDataVolume); err != nil {
		return err
	}

	if _, err := opts.podmanClient.RunContainer(ctx,
		podman.RunContainerOpts{
			Detach:   true,
			Image:    opts.MongodDockerImageName(),
			Name:     opts.LocalMongodHostname(),
			Hostname: opts.LocalMongodHostname(),
			EnvVars: map[string]string{
				"KEYFILECONTENTS": keyFileContents,
				"DBPATH":          "/data/db",
				"KEYFILE":         "/data/configdb/keyfile",
				"REPLSETNAME":     replicaSetName,
				"MONGOTHOST":      fmt.Sprintf("%s:%d", opts.LocalMongotHostname(), internalMongotPort),
			},
			Volumes: map[string]string{
				mongodDataVolume: "/data/db",
			},
			Ports: map[int]int{
				opts.Port: internalMongodPort,
			},
			Network: opts.LocalNetworkName(),
			// wrap the entrypoint with a chain of commands that
			// creates the keyfile in the container and sets the 400 permissions for it,
			// then starts the entrypoint with the local dev config
			Cmd:  "/bin/sh",
			Args: []string{"-c", string(mongodStartScript)},
		}); err != nil {
		return err
	}

	return opts.initReplicaSet(ctx)
}

func (opts *SetupOpts) initReplicaSet(ctx context.Context) error {
	// connect to local deployment
	connectionString, e := opts.ConnectionString(ctx)
	if e != nil {
		return e
	}

	const waitSeconds = 60
	if err := opts.mongodbClient.Connect(ctx, connectionString, waitSeconds); err != nil {
		return err
	}
	defer opts.mongodbClient.Disconnect(ctx)
	db := opts.mongodbClient.Database("admin")

	// initiate ReplicaSet
	if err := db.InitiateReplicaSet(ctx, replicaSetName, opts.LocalMongodHostname(), internalMongodPort, opts.Port); err != nil {
		return err
	}

	const waitForPrimarySeconds = 60
	for i := 0; i < waitForPrimarySeconds; i++ {
		r, err := db.RunCommand(ctx, bson.D{{Key: "hello", Value: 1}})
		if err != nil {
			continue
		}
		result := r.(bson.M)

		if state, ok := result["isWritablePrimary"].(bool); ok && state {
			break
		}
		time.Sleep(1 * time.Second)
	}

	// insert local dev cluster marker
	_, err := db.InsertOne(ctx, "atlascli", bson.M{"managedClusterType": "atlasCliLocalDevCluster"})
	return err
}

func (opts *SetupOpts) configureMongot(ctx context.Context, keyFileContents string) error {
	mongotDataVolume := opts.LocalMongotDataVolume()
	if _, err := opts.podmanClient.CreateVolume(ctx, mongotDataVolume); err != nil {
		return err
	}

	mongotMetricsVolume := opts.LocalMongoMetricsVolume()
	if _, err := opts.podmanClient.CreateVolume(ctx, mongotMetricsVolume); err != nil {
		return err
	}

	if _, err := opts.podmanClient.RunContainer(ctx, podman.RunContainerOpts{
		Detach:     true,
		Image:      options.MongotDockerImageName,
		Name:       opts.LocalMongotHostname(),
		Hostname:   opts.LocalMongotHostname(),
		Entrypoint: "/bin/sh",
		EnvVars: map[string]string{
			"MONGODHOST":      fmt.Sprintf("%s:%d", opts.LocalMongodHostname(), internalMongodPort),
			"DATADIR":         "/var/lib/mongot",
			"KEYFILEPATH":     "/var/lib/mongot/keyfile",
			"KEYFILECONTENTS": keyFileContents,
		},
		Args: []string{"-c", string(mongotStartScript)},
		Volumes: map[string]string{
			mongotDataVolume:    "/var/lib/mongot",
			mongotMetricsVolume: "/var/lib/mongot/metrics",
		},
		Network: opts.LocalNetworkName(),
	}); err != nil {
		return err
	}

	for { // wait for mongot
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			_, err := opts.podmanClient.RunContainer(ctx, podman.RunContainerOpts{
				Remove:   true,
				Detach:   false,
				Image:    opts.MongodDockerImageName(),
				Name:     opts.LocalCheckHostname(),
				Hostname: opts.LocalCheckHostname(),
				EnvVars: map[string]string{
					"MONGOTHOST": fmt.Sprintf("%s:%d", opts.LocalMongotHostname(), internalMongotPort),
				},
				Entrypoint: "/bin/sh",
				Args:       []string{"-c", `mongosh $MONGOTHOST --eval "db.adminCommand('ping')"`},
				Network:    opts.LocalNetworkName(),
			})

			if err == nil { // ping was successful
				return nil
			}
		}
	}
}

func (opts *SetupOpts) validateLocalDeploymentsSettings(containers []*podman.Container) error {
	mongodContainerName := opts.LocalMongodHostname()
	for _, c := range containers {
		for _, n := range c.Names {
			if n == mongodContainerName {
				return fmt.Errorf("\"%s\" deployment was already created and is currently in \"%s\" state", opts.DeploymentName, c.State)
			}
		}
	}

	return nil
}

func (opts *SetupOpts) promptSettings() error {
	p := &survey.Select{
		Message: "How do you want to setup your local MongoDB database?",
		Options: settingOptions,
		Default: opts.settings,
		Description: func(value string, index int) string {
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
		Message: "Cluster Name [This can't be changed later]",
		Default: opts.DeploymentName,
	}

	return telemetry.TrackAskOne(p, &opts.DeploymentName, survey.WithValidator(func(ans interface{}) error {
		name, _ := ans.(string)
		return options.ValidateDeploymentName(name)
	}))
}

func (opts *SetupOpts) promptMdbVersion() error {
	p := &survey.Select{
		Message: "MongoDB Version",
		Options: mdbVersions,
		Default: opts.MdbVersion,
	}

	return telemetry.TrackAskOne(p, &opts.MdbVersion, nil)
}

func availablePort() (int, error) {
	// prefer mongodb default's port
	if err := checkPort(internalMongodPort); err == nil {
		return internalMongodPort, nil
	}

	server, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	defer server.Close()

	_, port, err := net.SplitHostPort(server.Addr().String())
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(port)
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
	exportPort := strconv.Itoa(opts.Port)

	p := &survey.Input{
		Message: "Specify a port",
		Default: exportPort,
	}

	err := telemetry.TrackAskOne(p, &exportPort, survey.WithValidator(func(ans interface{}) error {
		input, _ := ans.(string)
		value, err := strconv.Atoi(input)
		if err != nil {
			return errMustBeInt
		}

		return validatePort(value)
	}))

	if err != nil {
		return err
	}

	opts.Port, err = strconv.Atoi(exportPort)
	return err
}

func (opts *SetupOpts) validateDeploymentTypeFlag() error {
	if opts.DeploymentType == "" && opts.force {
		return errFlagTypeRequired
	}

	if opts.DeploymentType != "" && !strings.EqualFold(opts.DeploymentType, atlasCluster) && !strings.EqualFold(opts.DeploymentType, localCluster) {
		return fmt.Errorf("%w: %s", errInvalidDeploymentType, opts.DeploymentType)
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

	if opts.MdbVersion != "" && opts.MdbVersion != mdb6 && opts.MdbVersion != mdb7 {
		return fmt.Errorf("%w: %s", errInvalidMongoDBVersion, opts.MdbVersion)
	}

	if opts.Port != 0 {
		if err := validatePort(opts.Port); err != nil {
			return err
		}
	}

	if opts.connectWith != "" &&
		!strings.EqualFold(opts.connectWith, compassConnect) &&
		!strings.EqualFold(opts.connectWith, mongoshConnect) &&
		!strings.EqualFold(opts.connectWith, skipConnect) {
		return fmt.Errorf("%w: %s", errUnsupportedConnectWith, opts.connectWith)
	}

	return nil
}

func (opts *SetupOpts) promptDeploymentType() error {
	p := &survey.Select{
		Message: "What would you like to deploy?",
		Options: deploymentTypeOptions,
		Help:    usage.DeploymentType,
		Description: func(value string, index int) string {
			return deploymentTypeDescription[value]
		},
	}

	return telemetry.TrackAskOne(p, &opts.DeploymentType, nil)
}

func (opts *SetupOpts) setDefaultSettings() (ok bool, err error) {
	opts.settings = defaultSettings

	if opts.DeploymentName == "" {
		opts.generateDeploymentName()
		ok = true
	}

	if opts.MdbVersion == "" {
		opts.MdbVersion = mdb7
		ok = true
	}

	if opts.Port == 0 {
		opts.Port, err = availablePort()
		ok = true
	}

	return
}

func (opts *SetupOpts) promptConnect() error {
	p := &survey.Select{
		Message: fmt.Sprintf("How would you like to connect to %s?", opts.DeploymentName),
		Options: connectWithOptions,
		Description: func(value string, index int) string {
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
		fmt.Fprintln(os.Stderr, "connection skipped")
	case compassConnect:
		if !compass.Detect() {
			return errCompassNotInstalled
		}
		if _, err := fmt.Fprintln(opts.OutWriter, "Launching MongoDB Compass..."); err != nil {
			return err
		}
		return compass.Run("", "", cs)
	case mongoshConnect:
		if !mongosh.Detect() {
			return errMongoshNotInstalled
		}
		return mongosh.Run("", "", cs)
	}

	return nil
}

func (opts *SetupOpts) validateAndPrompt() error {
	if err := opts.validateFlags(); err != nil {
		return err
	}

	if opts.DeploymentType == "" {
		if err := opts.promptDeploymentType(); err != nil {
			return err
		}
	}

	if strings.EqualFold(opts.DeploymentType, atlasCluster) {
		return fmt.Errorf("%w: %s", errDeploymentTypeNotImplemented, deploymentTypeDescription[opts.DeploymentType])
	}

	ok, err := opts.setDefaultSettings()
	if err != nil {
		return err
	}
	if ok {
		templatewriter.Print(os.Stderr, `
[Default Settings]
Cluster Name	{{.DeploymentName}}
MongoDB Version	{{.MdbVersion}}
Port	{{.Port}}

`, opts)

		if !opts.force {
			if err := opts.promptSettings(); err != nil {
				return err
			}
		}
	}

	switch opts.settings {
	case skipSettings:
		return errSkip
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

func (opts *SetupOpts) Run(ctx context.Context) error {
	if err := opts.validateAndPrompt(); err != nil {
		if errors.Is(err, errSkip) {
			_, _ = log.Warningln(err)
			return nil
		}

		return err
	}

	if err := opts.createLocalDeployment(ctx); err != nil {
		return err
	}

	cs := fmt.Sprintf("mongodb://localhost:%d/?directConnection=true", opts.Port)

	_, _ = log.Warningln("Cluster created!")
	_, _ = fmt.Fprintf(opts.OutWriter, "Connection string: %s\n", cs)
	_, _ = log.Warningln("")

	return opts.runConnectWith(cs)
}

// atlas deployments setup.
func SetupBuilder() *cobra.Command {
	opts := &SetupOpts{
		settings: defaultSettings,
	}
	cmd := &cobra.Command{
		Use:   "setup [deploymentName]",
		Short: "Create a local deployment.",
		Args:  require.MaximumNArgs(1),
		Annotations: map[string]string{
			"deploymentNameDesc": "Name of the deployment you want to setup.",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				opts.DeploymentName = args[0]
			}
			return opts.PreRunE(
				opts.InitOutput(cmd.OutOrStdout(), ""),
				opts.initPodmanClient,
				opts.initMongoDBClient,
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}

	cmd.Flags().StringVar(&opts.DeploymentType, flag.TypeFlag, "", usage.DeploymentType)
	cmd.Flags().IntVar(&opts.Port, flag.Port, 0, usage.MongodPort)
	cmd.Flags().StringVar(&opts.MdbVersion, flag.MDBVersion, "", usage.MDBVersion)
	cmd.Flags().StringVar(&opts.connectWith, flag.ConnectWith, "", usage.ConnectWith)

	cmd.Flags().BoolVar(&opts.force, flag.Force, false, usage.Force)

	_ = cmd.RegisterFlagCompletionFunc(flag.MDBVersion, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return mdbVersions, cobra.ShellCompDirectiveDefault
	})
	_ = cmd.RegisterFlagCompletionFunc(flag.TypeFlag, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return deploymentTypeOptions, cobra.ShellCompDirectiveDefault
	})
	_ = cmd.RegisterFlagCompletionFunc(flag.ConnectWith, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
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

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
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/briandowns/spinner"
	"github.com/google/uuid"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mongosh"
	"github.com/mongodb/mongodb-atlas-cli/internal/podman"
	"github.com/mongodb/mongodb-atlas-cli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/internal/templatewriter"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

const (
	startHostPort      = 37017
	internalMongodPort = 27017
	internalMongotPort = 27027
	localCluster       = "local"
	atlasCluster       = "atlas"
	mdb6               = "6.0"
	mdb7               = "7.0"
	replicaSetName     = "rs-localdev"
	// based on https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Clusters/operation/createCluster
	clusterNamePattern = "^[a-zA-Z0-9][a-zA-Z0-9-]*$"
	defaultSettings    = "default"
	customSettings     = "custom"
	skipSettings       = "skip"
	mongoshConnect     = "mongosh"
	skipConnect        = "skip"
	spinnerSpeed       = 100 * time.Millisecond
)

var (
	errSkip                         = errors.New("setup skipped")
	errMustBeInt                    = errors.New("input must be an integer")
	errWaitFailed                   = errors.New("waitConnection failed")
	errPortOutOfRange               = errors.New("port must within the range 1..65535")
	errPortNotAvailable             = errors.New("port not available")
	errInvalidClusterName           = errors.New("invalid cluster name")
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
	connectWithOptions     = []string{mongoshConnect, skipConnect}
	connectWithDescription = map[string]string{
		mongoshConnect: "MongoDB Shell",
		skipConnect:    "Skip Connection",
	}
	mdbVersions = []string{mdb7, mdb6}
)

type SetupOpts struct {
	options.DeploymentOpts
	cli.OutputOpts
	cli.GlobalOpts
	podmanClient podman.Client
	debug        bool
	settings     string
	connectWith  string
	force        bool
	s            *spinner.Spinner
}

const startTemplate = `local environment started at {{.ConnectionString}}
`

func (opts *SetupOpts) initPodmanClient() error {
	opts.podmanClient = podman.NewClient(opts.debug, opts.OutWriter)
	return nil
}

func (opts *SetupOpts) createLocalDeployment() error {
	fmt.Fprintf(os.Stderr, `
Creating your cluster %s [this might take several minutes]
`, opts.DeploymentName)
	opts.start()

	defer opts.stop()

	if err := opts.podmanClient.Ready(); err != nil {
		return err
	}

	if err := opts.podmanClient.Setup(); err != nil {
		return err
	}

	containers, errList := opts.podmanClient.ListContainers(options.MongodHostnamePrefix)
	if errList != nil {
		return errList
	}

	if err := opts.validateLocalDeploymentsSettings(containers); err != nil {
		return err
	}

	if _, err := opts.podmanClient.CreateNetwork(opts.LocalNetworkName()); err != nil {
		return err
	}

	if err := opts.configureMongod(); err != nil {
		return err
	}

	return opts.configureMongot()
}

func (opts *SetupOpts) configureMongod() error {
	mongodDataVolume := opts.LocalMongodDataVolume()
	if _, err := opts.podmanClient.CreateVolume(mongodDataVolume); err != nil {
		return err
	}

	keyfile := "/data/configdb/keyfile"
	keyfilePerm := 400
	mongodArgs := []string{
		"--transitionToAuth",
		"--keyFile", keyfile,
		"--dbpath", "/data/db",
		"--replSet", replicaSetName,
		"--setParameter", fmt.Sprintf("mongotHost=%s:%d", opts.LocalMongotHostname(), internalMongotPort),
		"--setParameter", fmt.Sprintf("searchIndexManagementHostAndPort=%s:%d", opts.LocalMongotHostname(), internalMongotPort),
	}

	// wrap the entrypoint with a chain of commands that
	// creates the keyfile in the container and sets the 400 permissions for it,
	// then starts the entrypoint with the local dev config
	cmdTemplate := "echo '%[1]s' > %[2]s && chmod %[3]d %[2]s && python3 /usr/local/bin/docker-entrypoint.py %[4]s"
	cmd := fmt.Sprintf(cmdTemplate,
		base64.URLEncoding.EncodeToString([]byte(opts.DeploymentID)),
		keyfile,
		keyfilePerm,
		strings.Join(mongodArgs, " "))

	if _, err := opts.podmanClient.RunContainer(
		podman.RunContainerOpts{
			Detach:   true,
			Image:    fmt.Sprintf("mongodb/mongodb-enterprise-server:%s-ubi8", opts.MdbVersion),
			Name:     opts.LocalMongodHostname(),
			Hostname: opts.LocalMongodHostname(),
			Volumes: map[string]string{
				mongodDataVolume: "/data/db",
			},
			Ports: map[int]int{
				opts.Port: internalMongodPort,
			},
			Network: opts.LocalNetworkName(),
			Args:    []string{"sh", "-c", cmd},
		}); err != nil {
		return err
	}

	// init ReplicaSet
	if err := opts.waitConnection(opts.Port); err != nil {
		return err
	}

	seedRs := fmt.Sprintf(`try {
		rs.status();
	  } catch {
		rs.initiate({
		  _id: "%s",
		  version: 1,
		  configsvr: false,
		  members: [{ _id: 0, host: "%s:%d", horizons: { external: "localhost:%d" } }],
		});
	  }`,
		replicaSetName,
		opts.LocalMongodHostname(),
		internalMongodPort,
		opts.Port)
	if err := opts.seed(opts.Port, seedRs); err != nil {
		return err
	}

	return opts.seed(opts.Port, "db.getSiblingDB('admin').atlascli.insertOne({ managedClusterType: 'atlasCliLocalDevCluster' })")
}

func (opts *SetupOpts) configureMongot() error {
	mongotDataVolume := opts.LocalMongotDataVolume()
	if _, err := opts.podmanClient.CreateVolume(mongotDataVolume); err != nil {
		return err
	}

	mongotMetricsVolume := opts.LocalMongoMetricsVolume()
	if _, err := opts.podmanClient.CreateVolume(mongotMetricsVolume); err != nil {
		return err
	}

	_, err := opts.podmanClient.RunContainer(podman.RunContainerOpts{
		Detach:   true,
		Image:    "mongodb/apix_test:mongot",
		Name:     opts.LocalMongotHostname(),
		Hostname: opts.LocalMongotHostname(),
		Args: []string{
			"--mongodHostAndPort", fmt.Sprintf("%s:%d", opts.LocalMongodHostname(), internalMongodPort),
			"--keyFileContent", base64.URLEncoding.EncodeToString([]byte(opts.DeploymentID)),
		},
		Volumes: map[string]string{
			mongotDataVolume:    "/var/lib/mongot",
			mongotMetricsVolume: "/var/lib/mongot/metrics",
		},
		Network: opts.LocalNetworkName(),
	})

	return err
}

func connString(port int) string {
	return fmt.Sprintf("mongodb://localhost:%d/admin", port)
}

func (opts *SetupOpts) seed(port int, script string) error {
	return mongosh.Exec(opts.debug, connString(port), "--eval", script)
}

func (opts *SetupOpts) validateLocalDeploymentsSettings(containers []podman.Container) error {
	mongodContainerName := opts.LocalMongodHostname()
	for _, c := range containers {
		for _, n := range c.Names {
			if n == mongodContainerName {
				return fmt.Errorf("\"%s\" deployment was already created and is currently in \"%s\" state", opts.DeploymentName, c.State)
			}
		}

		for _, p := range c.Ports {
			if p.HostPort == opts.Port {
				return fmt.Errorf("port %d is already used by \"%s\" local deployment", opts.Port, c.Names[0])
			}
		}
	}

	return nil
}

func (opts *SetupOpts) waitConnection(port int) error {
	const waitSeconds = 60
	for i := 0; i < waitSeconds; i++ {
		if err := mongosh.Exec(opts.debug, connString(port), "--eval", "db.runCommand('ping').ok"); err == nil {
			return nil
		}
		time.Sleep(1 * time.Second)
	}
	return errWaitFailed
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
		return validateDeploymentName(name)
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

func checkPort(p int) error {
	server, err := net.Listen("tcp", fmt.Sprintf(":%d", p))
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

func validateDeploymentName(n string) error {
	if matched, _ := regexp.MatchString(clusterNamePattern, n); !matched {
		return fmt.Errorf("%w: %s", errInvalidClusterName, n)
	}
	return nil
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
		if err := validateDeploymentName(opts.DeploymentName); err != nil {
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

	if opts.connectWith != "" && !strings.EqualFold(opts.connectWith, mongoshConnect) && !strings.EqualFold(opts.connectWith, skipConnect) {
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

func (opts *SetupOpts) setDefaultSettings() bool {
	set := false
	opts.settings = defaultSettings

	if opts.DeploymentName == "" {
		opts.generateDeploymentName()
		set = true
	}

	if opts.MdbVersion == "" {
		opts.MdbVersion = mdb7
		set = true
	}

	if opts.Port == 0 {
		opts.Port = 27017
		set = true
	}

	return set
}

func (opts *SetupOpts) promptConnect() error {
	if opts.connectWith != "" {
		return nil
	}

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
	if err := opts.promptConnect(); err != nil {
		return err
	}

	switch opts.connectWith {
	case skipConnect:
		fmt.Fprintln(os.Stderr, "connection skipped")
	case mongoshConnect:
		if !mongosh.Detect() {
			return errors.New("mongosh not found in your system")
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

	if opts.setDefaultSettings() {
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

func (opts *SetupOpts) Run(_ context.Context) error {
	if err := opts.validateAndPrompt(); err != nil {
		if errors.Is(err, errSkip) {
			_, _ = fmt.Fprintf(opts.OutWriter, "%s\n", err)
			return nil
		}

		return err
	}

	if err := opts.createLocalDeployment(); err != nil {
		return err
	}

	cs := fmt.Sprintf("mongodb://localhost:%d", opts.Port)

	fmt.Fprintf(opts.OutWriter, `Cluster created!
Connection string: %s

`, cs)

	return opts.runConnectWith(cs)
}

// atlas deployments setup.
func SetupBuilder() *cobra.Command {
	opts := &SetupOpts{
		settings: defaultSettings,
	}
	cmd := &cobra.Command{
		Use:   "setup <clusterName>",
		Short: "Create a local deployment.",
		Args:  require.MaximumNArgs(1),
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the cluster you want to setup.",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				opts.DeploymentName = args[0]
			}
			opts.DeploymentID = uuid.NewString()

			return opts.PreRunE(opts.InitOutput(cmd.OutOrStdout(), startTemplate), opts.initPodmanClient)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}

	cmd.Flags().StringVar(&opts.DeploymentType, flag.TypeFlag, "", usage.DeploymentType)
	cmd.Flags().IntVar(&opts.Port, flag.Port, 0, usage.MongodPort)
	cmd.Flags().StringVar(&opts.MdbVersion, flag.MDBVersion, "", usage.MDBVersion)
	cmd.Flags().StringVar(&opts.connectWith, flag.ConnectWith, "", usage.ConnectWith)

	cmd.Flags().BoolVarP(&opts.debug, flag.Debug, flag.DebugShort, false, usage.Debug)
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
		opts.s.Start()
	}
}

func (opts *SetupOpts) stop() {
	if opts.IsTerminal() {
		opts.s.Stop()
	}
}

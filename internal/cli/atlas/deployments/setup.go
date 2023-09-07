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
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/podman"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/compass"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mongosh"
	"github.com/mongodb/mongodb-atlas-cli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/internal/templatewriter"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

const (
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
	compassConnect     = "compass"
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
	connectWithOptions     = []string{mongoshConnect, compassConnect, skipConnect}
	connectWithDescription = map[string]string{
		mongoshConnect: "MongoDB Shell",
		compassConnect: "MongoDB Compass",
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

func (opts *SetupOpts) initPodmanClient() error {
	opts.podmanClient = podman.NewClient(opts.debug, opts.OutWriter)
	return nil
}

func (opts *SetupOpts) createLocalDeployment(ctx context.Context) error {
	fmt.Fprintf(os.Stderr, `
Creating your cluster %s [this might take several minutes]
`, opts.DeploymentName)
	opts.start()

	defer opts.stop()

	if err := opts.podmanClient.Ready(ctx); err != nil {
		return err
	}

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
			Image:    fmt.Sprintf("docker.io/mongodb/mongodb-enterprise-server:%s-ubi8", opts.MdbVersion),
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
			Cmd: "/bin/sh",
			Args: []string{
				"-c",
				`echo $KEYFILECONTENTS > $KEYFILE && \
                   chmod 400 $KEYFILE && \
                   python3 /usr/local/bin/docker-entrypoint.py \
                                          --transitionToAuth \
                                          --dbpath $DBPATH \
                                          --keyFile $KEYFILE \
                                          --replSet $REPLSETNAME \
                                          --setParameter "mongotHost=$MONGOTHOST" \
                                          --setParameter "searchIndexManagementHostAndPort=$MONGOTHOST"`,
			},
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

func (opts *SetupOpts) configureMongot(ctx context.Context, keyFileContents string) error {
	mongotDataVolume := opts.LocalMongotDataVolume()
	if _, err := opts.podmanClient.CreateVolume(ctx, mongotDataVolume); err != nil {
		return err
	}

	mongotMetricsVolume := opts.LocalMongoMetricsVolume()
	if _, err := opts.podmanClient.CreateVolume(ctx, mongotMetricsVolume); err != nil {
		return err
	}

	_, err := opts.podmanClient.RunContainer(ctx, podman.RunContainerOpts{
		Detach:     true,
		Image:      "docker.io/mongodb/apix_test:mongot-preview",
		Name:       opts.LocalMongotHostname(),
		Hostname:   opts.LocalMongotHostname(),
		Entrypoint: "/bin/sh",
		EnvVars: map[string]string{
			"MONGODHOST":      fmt.Sprintf("%s:%d", opts.LocalMongodHostname(), internalMongodPort),
			"DATADIR":         "/var/lib/mongot",
			"KEYFILEPATH":     "/var/lib/mongot/keyfile",
			"KEYFILECONTENTS": keyFileContents,
		},
		Args: []string{
			"-c",
			`echo $KEYFILECONTENTS > $KEYFILEPATH && chmod 400 $KEYFILEPATH && /etc/mongot-localdev/mongot \
                             --data-dir $DATADIR \
                             --mongodHostAndPort $MONGODHOST \
                             --keyFile $KEYFILEPATH`,
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
			return errors.New("MongoDB Compass not found in your system")
		}
		if err := compass.Run("", "", cs); err != nil {
			return err
		}
		_, err := fmt.Fprintln(opts.OutWriter, "Launching MongoDB Compass...")
		return err
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
			_, _ = fmt.Fprintf(opts.OutWriter, "%s\n", err)
			return nil
		}

		return err
	}

	if err := opts.createLocalDeployment(ctx); err != nil {
		return err
	}

	cs := fmt.Sprintf("mongodb://localhost:%d/?directConnection=true", opts.Port)

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
			return opts.PreRunE(opts.InitOutput(cmd.OutOrStdout(), ""), opts.initPodmanClient)
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

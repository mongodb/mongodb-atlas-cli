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
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/andreangiolillo/mongocli-test/internal/cli"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/deployments/options"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/setup"
	"github.com/andreangiolillo/mongocli-test/internal/cli/require"
	"github.com/andreangiolillo/mongocli-test/internal/cli/workflows"
	"github.com/andreangiolillo/mongocli-test/internal/compass"
	"github.com/andreangiolillo/mongocli-test/internal/flag"
	"github.com/andreangiolillo/mongocli-test/internal/log"
	"github.com/andreangiolillo/mongocli-test/internal/mongodbclient"
	"github.com/andreangiolillo/mongocli-test/internal/mongosh"
	"github.com/andreangiolillo/mongocli-test/internal/podman"
	"github.com/andreangiolillo/mongocli-test/internal/search"
	"github.com/andreangiolillo/mongocli-test/internal/telemetry"
	"github.com/andreangiolillo/mongocli-test/internal/templatewriter"
	"github.com/andreangiolillo/mongocli-test/internal/usage"
	"github.com/briandowns/spinner"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	internalMongodPort = 27017
	mdb6               = "6.0"
	mdb7               = "7.0"
	replicaSetName     = "rs-localdev"
	maxConns           = "32200" // --maxConns https://jira.mongodb.org/browse/SERVER-51233: Given the default max_map_count is 65530, we can support ~32200 connections
	defaultSettings    = "default"
	customSettings     = "custom"
	cancelSettings     = "cancel"
	skipConnect        = "skip"
	spinnerSpeed       = 100 * time.Millisecond
	shortStepCount     = 2
	tempRootUser       = "tempRootUser"
)

var (
	errCancel                   = errors.New("the setup was cancelled")
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
	settingOptions      = []string{defaultSettings, customSettings, cancelSettings}
	settingsDescription = map[string]string{
		defaultSettings: "With default settings",
		customSettings:  "With custom settings",
		cancelSettings:  "Cancel setup",
	}
	connectWithOptions     = []string{options.MongoshConnect, options.CompassConnect, skipConnect}
	connectWithDescription = map[string]string{
		options.MongoshConnect: "MongoDB Shell",
		options.CompassConnect: "MongoDB Compass",
		skipConnect:            "Skip Connection",
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
	cli.InputOpts
	mongodbClient mongodbclient.MongoDBClient
	settings      string
	connectWith   string
	force         bool
	bindIPAll     bool
	mongodIP      string
	mongotIP      string
	s             *spinner.Spinner
	atlasSetup    *setup.Opts
}

func (opts *SetupOpts) initMongoDBClient(ctx context.Context) func() error {
	return func() error {
		opts.mongodbClient = mongodbclient.NewClientWithContext(ctx)
		return nil
	}
}

func (opts *SetupOpts) logStepStarted(msg string, currentStep int, totalSteps int) {
	fullMessage := fmt.Sprintf("%d/%d: %s", currentStep, totalSteps, msg)
	_, _ = log.Warningln(fullMessage)
	opts.start()
}

func (opts *SetupOpts) downloadImagesIfNotAvailable(ctx context.Context, currentStep int, steps int) error {
	opts.logStepStarted("Downloading the MongoDB binaries to your local environment...", currentStep, steps)
	defer opts.stop()

	var mongodImages []*podman.Image
	var mongotImages []*podman.Image
	var err error

	if mongodImages, err = opts.PodmanClient.ListImages(ctx, opts.MongodDockerImageName()); err != nil {
		return err
	}

	if len(mongodImages) == 0 {
		if _, err = opts.PodmanClient.PullImage(ctx, opts.MongodDockerImageName()); err != nil {
			return err
		}
	}

	if mongotImages, err = opts.PodmanClient.ListImages(ctx, options.MongotDockerImageName); err != nil {
		return err
	}

	if len(mongotImages) == 0 {
		if _, err := opts.PodmanClient.PullImage(ctx, options.MongotDockerImageName); err != nil {
			return err
		}
	}

	return nil
}

func (opts *SetupOpts) setupPodman(ctx context.Context, currentStep int, steps int) error {
	opts.logStepStarted("Downloading and completing the configuration...", currentStep, steps)
	defer opts.stop()

	return opts.PodmanClient.Ready(ctx)
}

func (opts *SetupOpts) startEnvironment(ctx context.Context, currentStep int, steps int) error {
	opts.logStepStarted("Starting your local environment...", currentStep, steps)
	defer opts.stop()

	containers, errList := opts.PodmanClient.ListContainers(ctx, options.MongodHostnamePrefix)
	if errList != nil {
		return errList
	}

	return opts.validateLocalDeploymentsSettings(containers)
}

func (opts *SetupOpts) internalIPs(ctx context.Context) error {
	n, err := opts.PodmanClient.Network(ctx, opts.LocalNetworkName())
	if err != nil {
		return err
	}
	if len(n) < 1 {
		return podman.ErrNetworkNotFound
	}
	if len(n[0].Subnets) < 1 {
		return podman.ErrNetworkNotFound
	}
	ipNet := n[0].Subnets[0].Subnet
	if len(ipNet.IP) < 4 { //nolint:gomnd // ip ranges are meant to have 4 segments (e.g. 127.0.0.1)
		return podman.ErrNetworkNotFound
	}
	ipNet.IP[3] = 10
	opts.mongodIP = ipNet.IP.String()
	ipNet.IP[3] = 11
	opts.mongotIP = ipNet.IP.String()

	return nil
}

func (opts *SetupOpts) planSteps(ctx context.Context) (steps int, needPodmanSetup bool, needToPullImages bool) {
	steps = 2
	needPodmanSetup = false
	needToPullImages = false

	setupState := opts.PodmanClient.Diagnostics(ctx)

	if setupState.MachineRequired &&
		(!setupState.MachineFound || setupState.MachineState != "running") {
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
	opts.logStepStarted(fmt.Sprintf("Creating your deployment %s...", opts.DeploymentName), currentStep, steps)
	defer opts.stop()

	if _, err := opts.PodmanClient.CreateNetwork(ctx, opts.LocalNetworkName()); err != nil {
		return err
	}

	if err := opts.internalIPs(ctx); err != nil {
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
	if _, err := opts.PodmanClient.CreateVolume(ctx, mongodDataVolume); err != nil {
		return err
	}

	tempRootUserPassword := base64.URLEncoding.EncodeToString([]byte(uuid.NewString()))

	envVars := map[string]string{
		"KEYFILECONTENTS": keyFileContents,
		"DBPATH":          "/data/db",
		"KEYFILE":         "/data/configdb/keyfile",
		"MAXCONNS":        maxConns,
		"REPLSETNAME":     replicaSetName,
		"MONGOTHOST":      opts.InternalMongotAddress(opts.mongotIP),
	}

	if opts.IsAuthEnabled() {
		envVars["MONGODB_INITDB_ROOT_USERNAME"] = tempRootUser
		envVars["MONGODB_INITDB_ROOT_PASSWORD"] = tempRootUserPassword
	}

	if _, err := opts.PodmanClient.RunContainer(ctx,
		podman.RunContainerOpts{
			Detach:   true,
			Image:    opts.MongodDockerImageName(),
			Name:     opts.LocalMongodHostname(),
			Hostname: opts.LocalMongodHostname(),
			EnvVars:  envVars,
			Volumes: map[string]string{
				mongodDataVolume: "/data/db",
			},
			Ports: map[int]int{
				opts.Port: internalMongodPort,
			},
			BindIPAll: opts.bindIPAll,
			Network:   opts.LocalNetworkName(),
			IP:        opts.mongodIP,
			// wrap the entrypoint with a chain of commands that
			// creates the keyfile in the container and sets the 400 permissions for it,
			// then starts the entrypoint with the local dev config
			Cmd:  "/bin/sh",
			Args: []string{"-c", string(mongodStartScript)},
		}); err != nil {
		return err
	}

	return opts.initLocalDeployment(ctx, tempRootUserPassword)
}

func (opts *SetupOpts) initLocalDeployment(ctx context.Context, tempRootUserPassword string) error {
	// connect to local deployment
	connectionString := fmt.Sprintf("mongodb://localhost:%d/?directConnection=true", opts.Port)
	if opts.IsAuthEnabled() {
		connectionString = fmt.Sprintf("mongodb://%s:%s@localhost:%d/?directConnection=true",
			url.QueryEscape(tempRootUser),
			url.QueryEscape(tempRootUserPassword),
			opts.Port)
	}

	const waitSeconds = 60
	if err := opts.mongodbClient.Connect(connectionString, waitSeconds); err != nil {
		return err
	}
	defer opts.mongodbClient.Disconnect()
	adminDB := opts.mongodbClient.Database("admin")

	if err := opts.initReplicaSet(ctx, adminDB); err != nil {
		return err
	}

	if !opts.IsAuthEnabled() {
		return nil
	}

	if err := opts.createAdminUser(ctx, adminDB); err != nil {
		return err
	}

	return adminDB.DropUser(ctx, tempRootUser)
}

func (opts *SetupOpts) initReplicaSet(ctx context.Context, adminDB mongodbclient.Database) error {
	if err := adminDB.InitiateReplicaSet(ctx, replicaSetName, opts.LocalMongodHostname(), internalMongodPort, opts.Port); err != nil {
		return err
	}

	const waitForPrimarySeconds = 60
	for i := 0; i < waitForPrimarySeconds; i++ {
		r, err := adminDB.RunCommand(ctx, bson.D{{Key: "hello", Value: 1}})
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
	_, err := adminDB.InsertOne(ctx, "atlascli", bson.M{"managedClusterType": "atlasCliLocalDevCluster"})
	return err
}

func (opts *SetupOpts) createAdminUser(ctx context.Context, db mongodbclient.Database) error {
	return db.CreateUser(ctx,
		opts.DBUsername,
		opts.DBUserPassword,
		[]string{"userAdminAnyDatabase", "readWriteAnyDatabase"},
	)
}

func (opts *SetupOpts) internalMongodAddress() string {
	return fmt.Sprintf("%s:%d", opts.mongodIP, internalMongodPort)
}

func (opts *SetupOpts) configureMongot(ctx context.Context, keyFileContents string) error {
	mongotDataVolume := opts.LocalMongotDataVolume()
	if _, err := opts.PodmanClient.CreateVolume(ctx, mongotDataVolume); err != nil {
		return err
	}

	mongotMetricsVolume := opts.LocalMongoMetricsVolume()
	if _, err := opts.PodmanClient.CreateVolume(ctx, mongotMetricsVolume); err != nil {
		return err
	}

	if _, err := opts.PodmanClient.RunContainer(ctx, podman.RunContainerOpts{
		Detach:     true,
		Image:      options.MongotDockerImageName,
		Name:       opts.LocalMongotHostname(),
		Hostname:   opts.LocalMongotHostname(),
		Entrypoint: "/bin/sh",
		EnvVars: map[string]string{
			"MONGOD_HOST_AND_PORT": opts.internalMongodAddress(),
			"KEY_FILE_CONTENTS":    keyFileContents,
		},
		Args: []string{"-c", string(mongotStartScript)},
		Volumes: map[string]string{
			mongotDataVolume:    "/var/lib/mongot",
			mongotMetricsVolume: "/var/lib/mongot/metrics",
		},
		Network: opts.LocalNetworkName(),
		IP:      opts.mongotIP,
	}); err != nil {
		return err
	}

	return opts.WaitForMongot(ctx, opts.mongotIP)
}

func (opts *SetupOpts) validateLocalDeploymentsSettings(containers []*podman.Container) error {
	mongodContainerName := opts.LocalMongodHostname()
	for _, c := range containers {
		for _, n := range c.Names {
			if n == mongodContainerName {
				return fmt.Errorf("\"%s\" deployment already exists and is currently in \"%s\" state", opts.DeploymentName, c.State)
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
		Message: "Deployment Name [You can't change this value later]",
		Default: opts.DeploymentName,
	}

	return telemetry.TrackAskOne(p, &opts.DeploymentName, survey.WithValidator(func(ans interface{}) error {
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

	if opts.MdbVersion != "" && opts.MdbVersion != mdb6 && opts.MdbVersion != mdb7 {
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
		opts.MdbVersion = mdb7
		defaultValuesSet = true
	}

	if opts.Port == 0 {
		port, err := availablePort()
		if err != nil {
			return err
		}
		opts.Port = port
		defaultValuesSet = true
	}

	if defaultValuesSet {
		templatewriter.Print(os.Stderr, `
[Default Settings]
Deployment Name	{{.DeploymentName}}
MongoDB Version	{{.MdbVersion}}
Port	{{.Port}}

`, opts)
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
	case options.CompassConnect:
		if !compass.Detect() {
			return options.ErrCompassNotInstalled
		}
		if _, err := log.Warningln("Launching MongoDB Compass..."); err != nil {
			return err
		}
		return compass.Run("", "", cs)
	case options.MongoshConnect:
		if !mongosh.Detect() {
			return options.ErrMongoshNotInstalled
		}
		return mongosh.Run("", "", cs)
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
		_ = opts.RemoveLocal(ctx)
		return err
	}

	cs, err := opts.ConnectionString(ctx)
	if err != nil {
		return err
	}

	_, _ = log.Warningln("Deployment created!")
	_, _ = fmt.Fprintf(opts.OutWriter, "Connection string: %s\n", cs)
	_, _ = log.Warningln("")

	return opts.runConnectWith(cs)
}

func (opts *SetupOpts) runAtlas(ctx context.Context) error {
	s := setup.Builder()

	// remove global flags and unknown flags
	var newArgs []string
	_, _ = log.Debugf("Removing flags and args from original args %s\n", os.Args)

	flagstoRemove := map[string]string{
		flag.TypeFlag:    "1",
		flag.MDBVersion:  "1", // TODO: CLOUDP-200331
		flag.ConnectWith: "1", // TODO: CLOUDP-199422
	}

	newArgs, err := workflows.RemoveFlagsAndArgs(flagstoRemove, map[string]bool{opts.DeploymentName: true}, os.Args)
	if err != nil {
		return err
	}

	// replace deployment name with cluster name
	if opts.DeploymentName != "" {
		newArgs = append(newArgs, fmt.Sprintf("--%s", flag.ClusterName), opts.DeploymentName)
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

	telemetry.AppendOption(telemetry.WithDeploymentType(opts.DeploymentType))
	if strings.EqualFold(options.LocalCluster, opts.DeploymentType) {
		return opts.runLocal(ctx)
	}

	return opts.runAtlas(ctx)
}

// atlas deployments setup.
func SetupBuilder() *cobra.Command {
	opts := &SetupOpts{
		settings:   defaultSettings,
		atlasSetup: &setup.Opts{},
	}
	cmd := &cobra.Command{
		Use:     "setup [deploymentName]",
		Short:   "Create a local deployment.",
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
				opts.InitOutput(cmd.OutOrStdout(), ""),
				opts.InitInput(cmd.InOrStdin()),
				opts.InitStore(cmd.Context(), cmd.OutOrStdout()),
				opts.initMongoDBClient(cmd.Context()),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}

	// Local and Atlas
	cmd.Flags().StringVar(&opts.DeploymentType, flag.TypeFlag, "", usage.DeploymentType)
	cmd.Flags().StringVar(&opts.MdbVersion, flag.MDBVersion, "", usage.DeploymentMDBVersion)
	cmd.Flags().StringVar(&opts.connectWith, flag.ConnectWith, "", usage.ConnectWith)

	// Local only
	cmd.Flags().IntVar(&opts.Port, flag.Port, 0, usage.MongodPort)
	cmd.Flags().BoolVar(&opts.bindIPAll, flag.BindIPAll, false, usage.BindIPAll)

	// Atlas only
	opts.atlasSetup.SetupAtlasFlags(cmd)
	opts.atlasSetup.SetupFlowFlags(cmd)
	cmd.Flags().Lookup(flag.Region).Usage = usage.DeploymentRegion
	cmd.Flags().Lookup(flag.Tag).Usage = usage.DeploymentTag
	cmd.Flags().Lookup(flag.Tier).Usage = usage.DeploymentTier
	cmd.Flags().Lookup(flag.EnableTerminationProtection).Usage = usage.EnableTerminationProtectionForDeployment
	cmd.Flags().Lookup(flag.SkipSampleData).Usage = usage.SkipSampleDataDeployment
	cmd.Flags().Lookup(flag.Force).Usage = usage.ForceDeploymentsSetup

	_ = cmd.RegisterFlagCompletionFunc(flag.MDBVersion, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return mdbVersions, cobra.ShellCompDirectiveDefault
	})
	_ = cmd.RegisterFlagCompletionFunc(flag.TypeFlag, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return options.DeploymentTypeOptions, cobra.ShellCompDirectiveDefault
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

// Copyright 2022 MongoDB Inc
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

package setup

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/auth"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/compass"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mongosh"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/prerun"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/sighandle"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/validate"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/vscode"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

const (
	DefaultAtlasTier           = "M0"
	defaultAtlasGovTier        = "M30"
	atlasAdmin                 = "atlasAdmin"
	replicaSet                 = "REPLICASET"
	defaultProvider            = "AWS"
	defaultRegion              = "US_EAST_1"
	defaultRegionGCP           = "US_EAST_4"
	defaultRegionAzure         = "US_EAST_2"
	defaultRegionGov           = "US_GOV_EAST_1"
	defaultSettings            = "default"
	customSettings             = "custom"
	cancelSettings             = "cancel"
	skipConnect                = "skip"
	compassConnect             = "compass"
	mongoshConnect             = "mongosh"
	vsCodeConnect              = "vscode"
	clusterWideScaling         = "clusterWideScaling"
	independentShardScaling    = "independentShardScaling"
	deprecateMessageSharedTier = "The '%s' tier is deprecated. For the migration guide and timeline, visit: https://dochub.mongodb.org/core/flex-migration.\n"
)

var (
	settingOptions      = []string{defaultSettings, customSettings, cancelSettings}
	settingsDescription = map[string]string{
		defaultSettings: "With default settings",
		customSettings:  "With custom settings",
		cancelSettings:  "Cancel setup",
	}
	connectWithOptions     = []string{mongoshConnect, compassConnect, vsCodeConnect, skipConnect}
	connectWithDescription = map[string]string{
		mongoshConnect: "MongoDB Shell",
		compassConnect: "MongoDB Compass",
		vsCodeConnect:  "MongoDB for VsCode",
		skipConnect:    "Skip Connection",
	}
)

var errNeedsProject = errors.New("ensure you select or add a project to the profile")

const setupTemplateCloseHandler = `
Enter 'atlas cluster watch %s' to learn when your cluster is available.
`

const setupTemplateStoreWarning = `
Store your database authentication access details in a secure location:
Database User Username: %s
Database User Password: %s
`

const setupTemplateIntro = `Press [Enter] to use the default values.

Enter [?] on any option to get help.
`

const setupTemplateCluster = `
Creating your cluster... [It's safe to 'Ctrl + C']
`
const setupTemplateIPNotFound = `
We could not find your public IP address. To add your IP address run:
  atlas accesslist create

`

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=setup_mock_test.go -package=setup . AtlasClusterQuickStarter

type profileReader interface {
	ProjectID() string
	OrgID() string
}

type AtlasClusterQuickStarter interface {
	AddSampleData(string, string) (*atlasv2.SampleDatasetStatus, error)
	SampleDataStatus(string, string) (*atlasv2.SampleDatasetStatus, error)
	CloudProviderRegions(string, string, []string) (*atlasv2.PaginatedApiAtlasProviderRegions, error)
	CreateCluster(v15 *atlasClustersPinned.AdvancedClusterDescription) (*atlasClustersPinned.AdvancedClusterDescription, error)
	LatestAtlasCluster(string, string) (*atlasv2.ClusterDescription20240805, error)
	CreateClusterLatest(v15 *atlasv2.ClusterDescription20240805) (*atlasv2.ClusterDescription20240805, error)
	MDBVersions(projectID string, opt *store.MDBVersionListOptions) (*atlasv2.PaginatedAvailableVersion, error)
	CreateDatabaseUser(*atlasv2.CloudDatabaseUser) (*atlasv2.CloudDatabaseUser, error)
	DatabaseUser(string, string, string) (*atlasv2.CloudDatabaseUser, error)
	CreateProjectIPAccessList([]*atlasv2.NetworkPermissionEntry) (*atlasv2.PaginatedNetworkAccess, error)
}

type Opts struct {
	cli.ProjectOpts
	cli.WatchOpts
	register                    auth.RegisterOpts
	config                      profileReader
	store                       AtlasClusterQuickStarter
	defaultName                 string
	ClusterName                 string
	Provider                    string
	Region                      string
	IPAddresses                 []string
	IPAddressesResponse         string
	DBUsername                  string
	DBUserPassword              string
	SampleDataJobID             string
	MDBVersion                  string
	Tier                        string
	Tag                         map[string]string
	SkipSampleData              bool
	SkipMongosh                 bool
	connectWith                 string
	DefaultValue                bool
	Confirm                     bool
	CurrentIP                   bool
	EnableTerminationProtection bool
	AutoScalingMode             string
	flags                       *pflag.FlagSet
	flagSet                     map[string]struct{}
	settings                    string
	connectionString            string

	// control
	skipRegister bool
	skipLogin    bool
}

type clusterSettings struct {
	ClusterName                 string
	Provider                    string
	Region                      string
	Tier                        string
	DBUsername                  string
	DBUserPassword              string
	IPAddresses                 []string
	EnableTerminationProtection bool
	SkipSampleData              bool
	Tag                         map[string]string
	MdbVersion                  string
}

func (opts *Opts) providerAndRegionToConstant() {
	opts.Provider = strings.ToUpper(opts.Provider)
	opts.Region = strings.ReplaceAll(strings.ToUpper(opts.Region), "-", "_")
}

func (opts *Opts) trackFlags() {
	if opts.flags == nil {
		opts.flagSet = make(map[string]struct{})
		return
	}

	opts.flagSet = make(map[string]struct{}, opts.flags.NFlag())
	opts.flags.Visit(func(f *pflag.Flag) {
		opts.flagSet[f.Name] = struct{}{}
	})
}

func (opts *Opts) newDefaultValues() (*clusterSettings, error) {
	values := &clusterSettings{}
	values.SkipSampleData = opts.SkipSampleData

	values.ClusterName = opts.ClusterName
	if opts.ClusterName == "" {
		values.ClusterName = opts.defaultName
	}

	values.Provider = opts.Provider
	if opts.Provider == "" {
		values.Provider = defaultProvider
	}

	values.Region = opts.Region
	if opts.Region == "" {
		if config.CloudGovService == config.Service() {
			values.Region = defaultRegionGov
		} else {
			switch strings.ToUpper(opts.Provider) {
			case "AZURE":
				values.Region = defaultRegionAzure
			case "GCP":
				values.Region = defaultRegionGCP
			default:
				values.Region = defaultRegion
			}
		}
	}

	values.MdbVersion = opts.MDBVersion
	if opts.MDBVersion == "" {
		opts.MDBVersion, _ = cli.DefaultMongoDBMajorVersion()
	}

	values.DBUsername = opts.DBUsername
	if opts.DBUsername == "" {
		values.DBUsername = opts.defaultName
	}

	values.DBUserPassword = opts.DBUserPassword
	if opts.DBUserPassword == "" {
		pwd, err := generatePassword()
		if err != nil {
			return nil, err
		}
		values.DBUserPassword = pwd
	}

	values.IPAddresses = opts.IPAddresses
	if len(opts.IPAddresses) == 0 {
		if publicIP := store.IPAddress(); publicIP != "" {
			values.IPAddresses = []string{publicIP}
		} else {
			_, _ = log.Warning(setupTemplateIPNotFound)
		}
	}

	values.Tier = opts.Tier
	values.EnableTerminationProtection = opts.EnableTerminationProtection
	values.Tag = opts.Tag

	return values, nil
}

func (opts *Opts) clusterCreationWatcher() (any, bool, error) {
	result, err := opts.store.LatestAtlasCluster(opts.ConfigProjectID(), opts.ClusterName)
	if err != nil {
		return nil, false, err
	}
	return nil, result.GetStateName() == "IDLE", nil
}

func (opts *Opts) sampleDataWatcher() (any, bool, error) {
	result, err := opts.store.SampleDataStatus(opts.ConfigProjectID(), opts.SampleDataJobID)
	if err != nil {
		return nil, false, err
	}
	if result.GetState() == "FAILED" {
		return nil, false, fmt.Errorf("failed to load data: %s", result.GetErrorMessage())
	}
	return nil, result.GetState() == "COMPLETED", nil
}

func (opts *Opts) loadSampleData() error {
	if opts.SkipSampleData {
		return nil
	}

	fmt.Print(`
Loading sample data into your cluster... [It's safe to 'Ctrl + C']
`)
	sampleDataJob, err := opts.store.AddSampleData(opts.ConfigProjectID(), opts.ClusterName)

	if err != nil {
		return nil
	}

	opts.SampleDataJobID = sampleDataJob.GetId()

	_, err = opts.Watch(opts.sampleDataWatcher)
	return err
}

func (opts *Opts) createResources() error {
	if err := opts.createDatabaseUser(); err != nil {
		return err
	}

	if err := opts.createAccessList(); err != nil {
		return err
	}

	if err := opts.createCluster(); err != nil {
		target, _ := atlasClustersPinned.AsError(err)
		if target.GetErrorCode() == "CANNOT_CREATE_FREE_CLUSTER_VIA_PUBLIC_API" && strings.Contains(strings.ToLower(target.GetDetail()), cli.ErrFreeClusterAlreadyExists.Error()) {
			return cli.ErrFreeClusterAlreadyExists
		} else if target.GetErrorCode() == "INVALID_ATTRIBUTE" && strings.Contains(target.GetDetail(), "regionName") {
			return cli.ErrNoRegionExistsTryCommand
		}
		return err
	}
	return nil
}

func (opts *Opts) askSampleDataQuestion() error {
	if opts.SkipSampleData {
		return nil
	}

	q := newSampleDataQuestion()
	var addSampleData bool
	if err := telemetry.TrackAskOne(q, &addSampleData); err != nil {
		return err
	}
	opts.SkipSampleData = !addSampleData

	return nil
}

func (opts *Opts) interactiveSetup() error {
	if err := opts.askClusterOptions(); err != nil {
		return err
	}

	if err := opts.askSampleDataQuestion(); err != nil {
		return err
	}

	if err := opts.askDBUserOptions(); err != nil {
		return err
	}

	return opts.askAccessListOptions()
}

func (opts *Opts) shouldAskForValue(f string) bool {
	_, isFlagSet := opts.flagSet[f]
	return !isFlagSet
}

func (opts *Opts) replaceWithDefaultSettings(values *clusterSettings) {
	if values.ClusterName != "" {
		opts.ClusterName = values.ClusterName
	}

	if values.Provider != "" {
		opts.Provider = values.Provider
	}

	if values.Region != "" {
		opts.Region = values.Region
	}

	if values.MdbVersion != "" {
		opts.MDBVersion = values.MdbVersion
	}

	if values.DBUsername != "" {
		opts.DBUsername = values.DBUsername
	}

	if values.DBUserPassword != "" {
		opts.DBUserPassword = values.DBUserPassword
	}

	if values.IPAddresses != nil {
		opts.IPAddresses = values.IPAddresses
	}

	opts.EnableTerminationProtection = values.EnableTerminationProtection
	opts.SkipSampleData = values.SkipSampleData
	opts.Tag = values.Tag
}

// setupCloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by printing
// the dbUsername and dbPassword.
func (opts *Opts) setupCloseHandler() {
	sighandle.Notify(func(sig os.Signal) {
		fmt.Printf(setupTemplateCloseHandler, opts.ClusterName)
		telemetry.FinishTrackingCommand(telemetry.TrackOptions{
			Signal: sig.String(),
		})
		os.Exit(0)
	}, os.Interrupt, syscall.SIGTERM)
}

func (opts *Opts) Run(ctx context.Context) error {
	if !opts.skipRegister {
		_, _ = fmt.Fprintf(opts.OutWriter, `
This command will help you:
1. Create and verify your MongoDB Atlas account in your browser.
2. Return to the terminal to create your first free MongoDB database in Atlas.
`)
		if err := opts.register.RegisterRun(ctx); err != nil {
			return err
		}
	} else if !opts.skipLogin {
		_, _ = fmt.Fprintf(opts.OutWriter, `Next steps:
1. Log in and verify your MongoDB Atlas account in your browser.
2. Return to the terminal to create your first free MongoDB database in Atlas.
`)

		if err := opts.register.LoginRun(ctx); err != nil {
			return err
		}
	}

	if err := opts.clusterPreRun(ctx, opts.OutWriter); err != nil {
		return err
	}

	if opts.config.ProjectID() == "" {
		return fmt.Errorf("%w: %s", errNeedsProject, config.Default().Name())
	}

	return opts.setupCluster()
}

func (opts *Opts) clusterPreRun(ctx context.Context, outWriter io.Writer) error {
	opts.setTier()
	defaultProfile := config.Default()

	return opts.PreRunE(
		opts.initStore(ctx),
		opts.register.SyncWithOAuthAccessProfile(defaultProfile),
		opts.register.InitFlow(defaultProfile),
		opts.InitOutput(outWriter, ""),
	)
}

func (opts *Opts) setupCluster() error {
	const base10 = 10
	opts.defaultName = "Cluster" + strconv.FormatInt(time.Now().Unix(), base10)[5:]
	opts.providerAndRegionToConstant()
	opts.trackFlags()

	if opts.CurrentIP {
		if publicIP := store.IPAddress(); publicIP != "" {
			opts.IPAddresses = []string{publicIP}
		} else {
			_, _ = log.Warning(setupTemplateIPNotFound)
		}
	}

	values, dErr := opts.newDefaultValues()
	if dErr != nil {
		return dErr
	}

	if opts.Confirm {
		opts.settings = defaultSettings
	} else {
		if err := opts.askConfirmDefaultQuestion(values); err != nil {
			return err
		}
	}

	switch opts.settings {
	case customSettings:
		fmt.Print(setupTemplateIntro)

		if err := opts.interactiveSetup(); err != nil {
			return err
		}
	case defaultSettings:
		opts.replaceWithDefaultSettings(values)
	case cancelSettings:
		_, _ = fmt.Println("user-aborted. Not creating cluster")
		return nil
	}

	// Create db user, access list and cluster
	if err := opts.createResources(); err != nil {
		return err
	}

	fmt.Printf(`We are deploying %s...
`, opts.ClusterName)

	fmt.Printf(setupTemplateStoreWarning, opts.DBUsername, opts.DBUserPassword)
	opts.setupCloseHandler()

	fmt.Print(setupTemplateCluster)

	// Watch cluster creation
	if _, er := opts.Watch(opts.clusterCreationWatcher); er != nil {
		return er
	}

	fmt.Println("Cluster created.")

	// Get cluster's connection string
	cluster, err := opts.store.LatestAtlasCluster(opts.ConfigProjectID(), opts.ClusterName)
	if err != nil {
		return err
	}

	opts.connectionString = cluster.ConnectionStrings.GetStandardSrv()

	fmt.Printf("Your connection string: %v\n", opts.connectionString)

	if err := opts.loadSampleData(); err != nil {
		return err
	}

	return opts.runConnectWith()
}

func (opts *Opts) runConnectWith() error {
	if opts.connectWith == "" {
		if opts.SkipMongosh { // deprecated flag --skipMongosh
			return nil
		}

		if opts.Confirm { // --force
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
	case compassConnect:
		if !compass.Detect() {
			return compass.ErrCompassNotInstalled
		}
		if _, err := log.Warningln("Launching MongoDB Compass..."); err != nil {
			return err
		}
		return compass.Run(opts.DBUsername, opts.DBUserPassword, opts.connectionString)
	case mongoshConnect:
		if !mongosh.Detect() {
			return mongosh.ErrMongoshNotInstalled
		}
		return mongosh.Run(opts.DBUsername, opts.DBUserPassword, opts.connectionString)
	case vsCodeConnect:
		if !vscode.Detect() {
			return vscode.ErrVsCodeCliNotInstalled
		}
		if _, err := log.Warningln("Launching VsCode..."); err != nil {
			return err
		}
		return vscode.SaveConnection(opts.connectionString, opts.ClusterName, "atlas")
	}

	return nil
}

func (opts *Opts) promptConnect() error {
	p := &survey.Select{
		Message: "How would you like to connect to your cluster?",
		Options: connectWithOptions,
		Description: func(value string, _ int) string {
			return connectWithDescription[value]
		},
	}

	return telemetry.TrackAskOne(p, &opts.connectWith, nil)
}

func (opts *Opts) PreRun(ctx context.Context) error {
	opts.skipRegister = true
	opts.skipLogin = true

	if err := validate.NoAPIKeys(); err != nil {
		// Why are we ignoring the error?
		// Because if the user has API keys, we just want to proceed with the flow
		// Then why not remove the error?
		// The error is useful in other components that call `validate.NoAPIKeys()`
		return nil
	}
	if err := opts.register.RefreshAccessToken(ctx); err != nil && errors.Is(err, cli.ErrInvalidRefreshToken) {
		opts.skipLogin = false
		return nil
	}
	if _, err := auth.AccountWithAccessToken(); err == nil {
		return nil
	}
	opts.skipRegister = false
	return nil
}

func (opts *Opts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *Opts) setTier() {
	if config.CloudGovService == config.Service() && opts.Tier == DefaultAtlasTier {
		opts.Tier = defaultAtlasGovTier
	}
}

func (opts *Opts) SetupAtlasFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&opts.Tier, flag.Tier, DefaultAtlasTier, usage.Tier)
	cmd.Flags().StringVar(&opts.Provider, flag.Provider, "", usage.Provider)
	cmd.Flags().StringVarP(&opts.Region, flag.Region, flag.RegionShort, "", usage.Region)
	cmd.Flags().StringSliceVar(&opts.IPAddresses, flag.AccessListIP, []string{}, usage.NetworkAccessListIPEntry)
	cmd.Flags().StringVar(&opts.DBUsername, flag.Username, "", usage.DBUsername)
	cmd.Flags().StringVar(&opts.DBUserPassword, flag.Password, "", usage.Password)
	cmd.Flags().BoolVar(&opts.EnableTerminationProtection, flag.EnableTerminationProtection, false, usage.EnableTerminationProtection)
	cmd.Flags().BoolVar(&opts.CurrentIP, flag.CurrentIP, false, usage.CurrentIPSimplified)
	cmd.Flags().StringToStringVar(&opts.Tag, flag.Tag, nil, usage.Tag)
	cmd.Flags().StringVar(&opts.AutoScalingMode, flag.AutoScalingMode, clusterWideScaling, usage.AutoScalingMode)

	opts.AddProjectOptsFlags(cmd)

	cmd.MarkFlagsMutuallyExclusive(flag.CurrentIP, flag.AccessListIP)
}

func (opts *Opts) SetupFlowFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVar(&opts.SkipSampleData, flag.SkipSampleData, false, usage.SkipSampleData)
	cmd.Flags().BoolVar(&opts.SkipMongosh, flag.SkipMongosh, false, usage.SkipMongosh)
	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.ForceQuickstart)
	_ = cmd.Flags().MarkDeprecated(flag.SkipMongosh, "Use --connectWith instead")
	cmd.MarkFlagsMutuallyExclusive(flag.SkipMongosh, flag.ConnectWith)
}

func (opts *Opts) validateTier() error {
	opts.Tier = strings.ToUpper(opts.Tier)
	if opts.Tier == atlasM2 || opts.Tier == atlasM5 {
		_, _ = fmt.Fprintf(os.Stderr, deprecateMessageSharedTier, opts.Tier)
	}
	return nil
}

// Builder
// atlas setup
//
//	[--clusterName clusterName]
//	[--provider provider]
//	[--region regionName]
//	[--username username]
//	[--password password]
//	[--skipMongosh skipMongosh]
func Builder() *cobra.Command {
	opts := &Opts{}

	cmd := &cobra.Command{
		Use:     "setup",
		Aliases: []string{"quickstart"},
		Short:   "Register, authenticate, create, and access an Atlas cluster.",
		Long:    `This command takes you through registration, login, default profile creation, creating your first free tier cluster and connecting to it using MongoDB Shell.`,
		Example: `  # Override default cluster settings like name, provider, or database username by using the command options
  atlas setup --clusterName Test --provider GCP --username dbuserTest`,
		Hidden: false,
		Args:   require.NoArgs,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			defaultProfile := config.Default()
			opts.config = defaultProfile
			opts.OutWriter = cmd.OutOrStdout()
			opts.register.OutWriter = opts.OutWriter

			if err := opts.register.SyncWithOAuthAccessProfile(defaultProfile)(); err != nil {
				return err
			}
			if err := opts.register.InitFlow(defaultProfile)(); err != nil {
				return err
			}
			if err := opts.PreRun(cmd.Context()); err != nil {
				return nil
			}
			var preRun []prerun.CmdOpt
			// registration pre run if applicable
			if !opts.skipRegister {
				preRun = append(preRun,
					opts.register.LoginPreRun(cmd.Context()),
					validate.NoAPIKeys,
					validate.NoAccessToken,
				)
			}

			if !opts.skipLogin && opts.skipRegister {
				preRun = append(preRun, opts.register.LoginPreRun(cmd.Context()))
			}
			preRun = append(preRun, opts.validateTier)
			preRun = append(preRun, validate.AutoScalingMode(opts.AutoScalingMode))

			return opts.PreRunE(preRun...)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			opts.flags = cmd.Flags()
			return opts.Run(cmd.Context())
		},
	}

	// Register and login related
	cmd.Flags().BoolVar(&opts.register.IsGov, "gov", false, "Register with Atlas for Government.")
	cmd.Flags().BoolVar(&opts.register.NoBrowser, "noBrowser", false, "Don't try to open a browser session.")
	// Setup related
	cmd.Flags().StringVar(&opts.MDBVersion, flag.MDBVersion, "", usage.DeploymentMDBVersion)
	cmd.Flags().StringVar(&opts.connectWith, flag.ConnectWith, "", usage.ConnectWithAtlasSetup)
	opts.SetupAtlasFlags(cmd)
	opts.SetupFlowFlags(cmd)

	cmd.Flags().StringVar(&opts.ClusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().BoolVarP(&opts.DefaultValue, flag.Default, "Y", false, usage.QuickstartDefault)
	_ = cmd.Flags().MarkDeprecated(flag.Default, "please use --force instead")

	return cmd
}

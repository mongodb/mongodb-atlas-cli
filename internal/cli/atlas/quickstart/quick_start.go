// Copyright 2020 MongoDB Inc
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
package quickstart

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/convert"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/mongosh"
	"github.com/mongodb/mongocli/internal/search"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/mongodb/mongocli/internal/validate"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const quickstartTemplate = `
Now you can connect to your Atlas cluster with: mongosh -u %s -p %s %s

`
const quickstartTemplateCloseHandler = `
You can connect to your Atlas cluster with the following user: 
username: %s 
password: %s
`

const clusterDetails = `
[Set up your Atlas cluster]
`

const databaseUserDetails = `
[Set up your database access details]
`

const accessListDetails = `
[Set up your network access list details]
`

const mongoShellDetails = `
[Connect to your new cluster]
`

const creatingClusterDetails = `
Creating your cluster... [It's safe to 'Ctrl + C']
`

const loadingSampleData = `
Loading sample data into your cluster... [It's safe to 'Ctrl + C']
`

const (
	replicaSet        = "REPLICASET"
	mdbVersion        = "4.4"
	shards            = 1
	tier              = "M2"
	tenant            = "TENANT"
	members           = 3
	zoneName          = "Zone 1"
	accessListComment = "IP added with mongocli atlas quickstart"
	atlasAdmin        = "atlasAdmin"
	none              = "NONE"
	mongoshURL        = "https://www.mongodb.com/try/download/shell"
	atlasAccountURL   = "https://docs.atlas.mongodb.com/tutorial/create-atlas-account/?utm_campaign=atlas_quickstart&utm_source=mongocli&utm_medium=product/"
	profileDocURL     = "https://docs.mongodb.com/mongocli/stable/configure/?utm_campaign=atlas_quickstart&utm_source=mongocli&utm_medium=product#std-label-mcli-configure"
)

type Opts struct {
	cli.GlobalOpts
	cli.WatchOpts
	ClusterName     string
	Provider        string
	Region          string
	IPAddresses     []string
	IPAddress       string
	DBUsername      string
	DBUserPassword  string
	SampleDataJobID string
	SkipSampleData  bool
	SkipMongosh     bool
	store           store.AtlasClusterQuickStarter
}

func (opts *Opts) initStore() error {
	var err error
	opts.store, err = store.New(store.PublicAuthenticatedPreset(config.Default()))
	return err
}

func (opts *Opts) Run() error {
	if err := opts.createCluster(); err != nil {
		return err
	}

	fmt.Println("We are deploying your cluster...")

	if err := opts.createDatabaseUser(); err != nil {
		return err
	}

	if err := opts.createAccessList(); err != nil {
		return err
	}

	opts.setupCloseHandler()

	runMongoShell, er := opts.askMongoShellQuestion()
	if er != nil {
		return er
	}

	fmt.Print(creatingClusterDetails)
	// Watch cluster creation
	if er := opts.Watch(opts.clusterCreationWatcher); er != nil {
		return er
	}

	if err := opts.loadSampleData(); err != nil {
		return nil
	}

	// Get cluster's connection string
	cluster, err := opts.store.AtlasCluster(opts.ConfigProjectID(), opts.ClusterName)
	if err != nil {
		return err
	}

	fmt.Printf(quickstartTemplate, opts.DBUsername, opts.DBUserPassword, cluster.SrvAddress)

	if runMongoShell {
		if err := mongosh.Run(config.MongoShellPath(), opts.DBUsername, opts.DBUserPassword, cluster.SrvAddress); err != nil {
			return err
		}
	}

	return nil
}

func (opts *Opts) createAccessList() error {
	if err := opts.askAccessListOptions(); err != nil {
		return err
	}
	// Add IP to project’s IP access list
	entries := opts.newProjectIPAccessList()
	if _, err := opts.store.CreateProjectIPAccessList(entries); err != nil {
		return err
	}

	return nil
}

func (opts *Opts) createDatabaseUser() error {
	if err := opts.askDBUserOptions(); err != nil {
		return err
	}

	if _, err := opts.store.CreateDatabaseUser(opts.newDatabaseUser()); err != nil {
		return err
	}

	return nil
}

func (opts *Opts) createCluster() error {
	if err := opts.askClusterOptions(); err != nil {
		return err
	}

	if _, err := opts.store.CreateCluster(opts.newCluster()); err != nil {
		return err
	}

	return nil
}

func (opts *Opts) loadSampleData() error {
	if opts.SkipSampleData {
		return nil
	}

	fmt.Print(loadingSampleData)
	sampleDataJob, err := opts.store.AddSampleData(opts.ConfigProjectID(), opts.ClusterName)

	if err != nil {
		return nil
	}

	opts.SampleDataJobID = sampleDataJob.ID

	if er := opts.Watch(opts.sampleDataWatcher); er != nil {
		return er
	}

	return nil
}

func (opts *Opts) sampleDataWatcher() (bool, error) {
	result, err := opts.store.SampleDataStatus(opts.ConfigProjectID(), opts.SampleDataJobID)
	if err != nil {
		return false, err
	}
	return result.State == "COMPLETED", nil
}

func (opts *Opts) clusterCreationWatcher() (bool, error) {
	result, err := opts.store.AtlasCluster(opts.ConfigProjectID(), opts.ClusterName)
	if err != nil {
		return false, err
	}
	return result.StateName == "IDLE", nil
}

func (opts *Opts) newDatabaseUser() *atlas.DatabaseUser {
	return &atlas.DatabaseUser{
		Roles:        convert.BuildAtlasRoles([]string{atlasAdmin}),
		GroupID:      opts.ConfigProjectID(),
		Password:     opts.DBUserPassword,
		X509Type:     none,
		AWSIAMType:   none,
		LDAPAuthType: none,
		DatabaseName: convert.AdminDB,
		Username:     opts.DBUsername,
	}
}

func (opts *Opts) newProjectIPAccessList() []*atlas.ProjectIPAccessList {
	accessListArray := make([]*atlas.ProjectIPAccessList, len(opts.IPAddresses))
	for i, addr := range opts.IPAddresses {
		accessList := &atlas.ProjectIPAccessList{
			GroupID:   opts.ConfigProjectID(),
			Comment:   accessListComment,
			IPAddress: addr,
		}

		accessListArray[i] = accessList
	}
	return accessListArray
}

func (opts *Opts) newCluster() *atlas.Cluster {
	diskSizeGB := atlas.DefaultDiskSizeGB[strings.ToUpper(tenant)][tier]
	cluster := &atlas.Cluster{
		GroupID:             opts.ConfigProjectID(),
		ClusterType:         replicaSet,
		ReplicationSpecs:    []atlas.ReplicationSpec{opts.newReplicationSpec()},
		ProviderSettings:    opts.newProviderSettings(),
		MongoDBMajorVersion: mdbVersion,
		DiskSizeGB:          &diskSizeGB,
		Name:                opts.ClusterName,
		Labels: []atlas.Label{
			{
				Key:   "Infrastructure Tool",
				Value: "MongoDB CLI Quickstart",
			},
		},
	}

	return cluster
}

func (opts *Opts) newReplicationSpec() atlas.ReplicationSpec {
	var (
		readOnlyNodes int64 = 0
		priority      int64 = 7
		shards        int64 = shards
		members       int64 = members
	)
	replicationSpec := atlas.ReplicationSpec{
		NumShards: &shards,
		ZoneName:  zoneName,
		RegionsConfig: map[string]atlas.RegionsConfig{
			opts.Region: {
				ReadOnlyNodes:  &readOnlyNodes,
				ElectableNodes: &members,
				Priority:       &priority,
			},
		},
	}
	return replicationSpec
}

func (opts *Opts) newProviderSettings() *atlas.ProviderSettings {
	return &atlas.ProviderSettings{
		InstanceSizeName:    tier,
		ProviderName:        tenant,
		RegionName:          opts.Region,
		BackingProviderName: opts.Provider,
	}
}

func (opts *Opts) askClusterOptions() error {
	var qs []*survey.Question

	if q := clusterNameQuestion(opts.ClusterName); q != nil {
		qs = append(qs, q)
	}

	if q := providerQuestion(opts.Provider); q != nil {
		qs = append(qs, q)
	}

	if opts.Provider == "" || opts.ClusterName == "" || opts.Region == "" {
		fmt.Print(clusterDetails)
	}

	if err := survey.Ask(qs, opts); err != nil {
		return err
	}

	// We need the provider to ask for the region
	if err := opts.askClusterRegion(); err != nil {
		return err
	}

	// We need the cluster name to ask for adding sample data
	return opts.askSampleDataQuestion()
}

func (opts *Opts) askSampleDataQuestion() error {
	if opts.SkipSampleData {
		return nil
	}

	q := newSampleDataQuestion(opts.ClusterName)
	addSampleData := false
	if err := survey.AskOne(q, &addSampleData); err != nil {
		return err
	}

	opts.SkipSampleData = !addSampleData

	return nil
}

func (opts *Opts) askClusterRegion() error {
	if opts.Region == "" {
		regions, err := opts.defaultRegions()
		if err != nil {
			return err
		}
		if regionQ := newRegionQuestions(regions); regionQ != nil {
			return survey.Ask([]*survey.Question{regionQ}, opts)
		}
	}

	return nil
}

func (opts *Opts) askDBUserOptions() error {
	var qs []*survey.Question

	if q := dbUsernameQuestion(opts.DBUsername, opts.validateUniqueUsername); q != nil {
		qs = append(qs, q)
	}

	if pwd, q := dbUserPasswordQuestion(opts.DBUserPassword); q != nil {
		opts.DBUsername = pwd
		qs = append(qs, q)
	}

	if len(qs) > 0 {
		fmt.Print(databaseUserDetails)
		if err := survey.Ask(qs, opts); err != nil {
			return err
		}
	}

	return nil
}

func (opts *Opts) askAccessListOptions() error {
	q := accessListQuestion(opts.IPAddresses)
	if q == nil {
		return nil
	}

	fmt.Print(accessListDetails)
	if err := survey.Ask([]*survey.Question{q}, opts); err != nil {
		return err
	}

	if len(opts.IPAddresses) == 0 && opts.IPAddress != "" {
		opts.IPAddresses = []string{opts.IPAddress}
	}

	return nil
}

func (opts *Opts) askMongoShellQuestion() (bool, error) {
	if response, err := askAccessDeploymentQuestion(opts.SkipMongosh, opts.ClusterName); !response || err != nil {
		return response, err
	}

	if config.MongoShellPath() != "" {
		return true, nil
	}

	fmt.Println("No MongoDB shell version detected.")

	if isInstalled, err := askIsMongoShellInstalledQuestion(); !isInstalled || err != nil {
		if err != nil {
			return isInstalled, err
		}

		runShell, err := askOpenMongoShellDownloadPage()
		if !runShell || err != nil {
			return runShell, err
		}
	}

	return askMongoShellPathQuestion()
}

func askMongoShellPathQuestion() (bool, error) {
	wantToProvidePath := false
	q := newMongoShellPathQuestion()

	if err := survey.AskOne(q, &wantToProvidePath); !wantToProvidePath || err != nil {
		return wantToProvidePath, err
	}

	if err := askMongoShellAndSetConfig(); err != nil {
		return false, err
	}

	return true, nil
}

func askIsMongoShellInstalledQuestion() (bool, error) {
	isInstalled := false
	q := newIsMongoShellInstalledQuestion()
	if err := survey.AskOne(q, &isInstalled); !isInstalled || err != nil {
		return isInstalled, err
	}

	return true, nil
}

func askAccessDeploymentQuestion(skip bool, clusterName string) (bool, error) {
	if q := accessDeploymentQuestion(skip, clusterName); q != nil {
		fmt.Print(mongoShellDetails)

		runMongoShell := false
		if err := survey.AskOne(q, &runMongoShell); !runMongoShell || err != nil {
			return runMongoShell, err
		}

		return true, nil
	}

	return false, nil
}

func (opts *Opts) validateUniqueUsername(val interface{}) error {
	username, ok := val.(string)
	if !ok {
		return fmt.Errorf("the username %s is not valid", username)
	}

	_, err := opts.store.DatabaseUser(convert.AdminDB, opts.ConfigProjectID(), username)
	var target *atlas.ErrorResponse

	if err != nil && errors.As(err, &target) {
		if target.ErrorCode == "USERNAME_NOT_FOUND" {
			return nil
		}
		return err
	}

	return fmt.Errorf("a user with this username %s already exists", username)
}

func askOpenMongoShellDownloadPage() (bool, error) {
	if openURL, err := askOpenBrowserQuestion(); !openURL || err != nil {
		return openURL, err
	}

	if err := browser.OpenURL(mongoshURL); err != nil {
		return false, err
	}

	return true, nil
}

func askOpenBrowserQuestion() (bool, error) {
	openURL := false
	prompt := newMongoShellQuestionOpenBrowser()
	if err := survey.AskOne(prompt, &openURL); !openURL || err != nil {
		return openURL, err
	}

	return true, nil
}

func askMongoShellAndSetConfig() error {
	var mongoShellPath string
	q := newMongoShellPathInput(mongosh.Path())
	if err := survey.Ask([]*survey.Question{q}, &mongoShellPath); err != nil {
		return err
	}

	config.SetMongoShellPath(mongoShellPath)
	if err := config.Save(); err != nil {
		return err
	}
	return nil
}

func askAtlasAccountAndProfile() error {
	_, _ = fmt.Fprintln(os.Stderr, "No API credentials set.")

	if err := openBrowserAtlasAccount(); err != nil {
		return err
	}

	if err := openBrowserProfile(); err != nil {
		return err
	}

	return validate.Credentials()
}

func openBrowserProfile() error {
	openBrowserProfileDoc := false
	q := newProfileDocQuestionOpenBrowser()
	if err := survey.AskOne(q, &openBrowserProfileDoc); !openBrowserProfileDoc || err != nil {
		return err
	}

	if err := browser.OpenURL(profileDocURL); err != nil {
		return err
	}

	return nil
}

func openBrowserAtlasAccount() error {
	openBrowserAtlasAccount := false
	q := newAtlasAccountQuestionOpenBrowser()
	if err := survey.AskOne(q, &openBrowserAtlasAccount); !openBrowserAtlasAccount || err != nil {
		return err
	}

	if err := browser.OpenURL(atlasAccountURL); err != nil {
		return err
	}

	return nil
}

func (opts *Opts) defaultRegions() ([]string, error) {
	cloudProviders, err := opts.store.CloudProviderRegions(opts.ConfigProjectID(), tier, []*string{&opts.Provider})

	if err != nil {
		return nil, err
	}

	if len(cloudProviders.Results) == 0 || len(cloudProviders.Results[0].InstanceSizes) == 0 {
		return nil, errors.New("no regions available")
	}

	availableRegions := cloudProviders.Results[0].InstanceSizes[0].AvailableRegions

	defaultRegions := make([]string, 0, len(availableRegions))
	popularRegionIndex := search.DefaultRegion(availableRegions)

	if popularRegionIndex != -1 {
		// the most popular region must be the first in the list
		popularRegion := availableRegions[popularRegionIndex]
		defaultRegions = append(defaultRegions, popularRegion.Name)

		// remove popular region from availableRegions
		availableRegions = append(availableRegions[:popularRegionIndex], availableRegions[popularRegionIndex+1:]...)
	}

	for _, v := range availableRegions {
		defaultRegions = append(defaultRegions, v.Name)
	}

	return defaultRegions, nil
}

// setupCloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by printing
// the dbUsername and dbPassword.
func (opts *Opts) setupCloseHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Printf(quickstartTemplateCloseHandler, opts.DBUsername, opts.DBUserPassword)
		os.Exit(0)
	}()
}

func (opts *Opts) providerAndRegionToConstant() {
	opts.Provider = strings.ToUpper(opts.Provider)
	opts.Region = strings.ReplaceAll(strings.ToUpper(opts.Region), "-", "_")
}

// mongocli atlas dbuser(s) quickstart [--clusterName clusterName] [--provider provider] [--region regionName] [--projectId projectId] [--username username] [--password password] [--skipMongosh skipMongosh].
func Builder() *cobra.Command {
	opts := &Opts{}
	cmd := &cobra.Command{
		Use: "quickstart",
		Example: `Skip setting cluster name, provider or database username by using the command options
  $ mongocli atlas quickstart --clusterName Test --provider GPC --username dbuserTest
`,
		Short: "Create and access an Atlas Cluster.",
		Long:  "This command creates a cluster, adds your public IP to the atlas access list and creates a db user to access your MongoDB instance.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if config.PublicAPIKey() == "" || config.PrivateAPIKey() == "" {
				// no profile set
				return askAtlasAccountAndProfile()
			}

			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), ""),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.providerAndRegionToConstant()
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ClusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().StringVar(&opts.Provider, flag.Provider, "", usage.Provider)
	cmd.Flags().StringVarP(&opts.Region, flag.Region, flag.RegionShort, "", usage.Region)
	cmd.Flags().StringSliceVar(&opts.IPAddresses, flag.AccessListIP, []string{}, usage.NetworkAccessListIPEntry)
	cmd.Flags().StringVar(&opts.DBUsername, flag.Username, "", usage.DBUsername)
	cmd.Flags().StringVar(&opts.DBUserPassword, flag.Password, "", usage.Password)
	cmd.Flags().BoolVar(&opts.SkipSampleData, flag.SkipSampleData, false, usage.SkipSampleData)
	cmd.Flags().BoolVar(&opts.SkipMongosh, flag.SkipMongosh, false, usage.SkipMongosh)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}

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
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/convert"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/mongosh"
	"github.com/mongodb/mongocli/internal/randgen"
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
	passwordLength    = 12
	mongoshURL        = "https://www.mongodb.com/try/download/shell"
	atlasAccountURL   = "https://docs.atlas.mongodb.com/tutorial/create-atlas-account/"
	profileDocURL     = "https://docs.mongodb.com/mongocli/stable/configure/#std-label-mcli-configure"
)

type Opts struct {
	cli.GlobalOpts
	cli.WatchOpts
	ClusterName    string
	Provider       string
	Region         string
	IPAddresses    []string
	IPAddress      string
	DBUsername     string
	DBUserPassword string
	SkipMongosh    bool
	store          store.AtlasClusterQuickStarter
}

func (opts *Opts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *Opts) Run() error {
	if err := opts.askClusterOptions(); err != nil {
		return err
	}

	if _, err := opts.store.CreateCluster(opts.newCluster()); err != nil {
		return err
	}

	fmt.Println("We are deploying your cluster...")

	if err := opts.askDBUserOptions(); err != nil {
		return err
	}

	if err := opts.askAccessListOptions(); err != nil {
		return err
	}

	if _, err := opts.store.CreateDatabaseUser(opts.newDatabaseUser()); err != nil {
		return err
	}

	opts.setupCloseHandler()

	// Add IP to project’s IP access list
	entries := opts.newProjectIPAccessList()
	if _, err := opts.store.CreateProjectIPAccessList(entries); err != nil {
		return err
	}

	runMongoShell, err := opts.askMongoShellQuestion()
	if err != nil {
		return err
	}

	fmt.Print(creatingClusterDetails)
	if er := opts.Watch(opts.watcher); er != nil {
		return er
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

func (opts *Opts) watcher() (bool, error) {
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
	return &atlas.Cluster{
		GroupID:             opts.ConfigProjectID(),
		ClusterType:         replicaSet,
		ReplicationSpecs:    []atlas.ReplicationSpec{opts.newReplicationSpec()},
		ProviderSettings:    opts.newProviderSettings(),
		MongoDBMajorVersion: mdbVersion,
		DiskSizeGB:          &diskSizeGB,
		Name:                opts.ClusterName,
	}
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

	clusterName := opts.ClusterName
	if clusterName == "" {
		message := ""
		clusterName = autogeneratedName()
		if clusterName != "" {
			message = fmt.Sprintf(" [Press Enter to use the auto-generated name '%s']", clusterName)
		}
		qs = append(qs, newClusterNameQuestion(clusterName, message))
	}

	if opts.Provider == "" {
		qs = append(qs, newClusterProviderQuestion())
	}

	if opts.Provider == "" || opts.ClusterName == "" || opts.Region == "" {
		fmt.Print(clusterDetails)
	}

	if err := survey.Ask(qs, opts); err != nil {
		return err
	}

	if opts.Region == "" {
		regions, err := opts.defaultRegions()
		if err != nil {
			return err
		}
		if regionQ := newRegionQuestions(regions); regionQ != nil {
			// we call survey.Ask two times because the region question needs opts.Provider to be populated
			return survey.Ask([]*survey.Question{regionQ}, opts)
		}
	}

	return nil
}

func (opts *Opts) defaultRegions() ([]string, error) {
	cloudProviders, err := opts.store.CloudProviderRegions(opts.ConfigProjectID(), opts.Provider, tier, false)

	if err != nil {
		return nil, err
	}

	if len(cloudProviders.Results) == 0 || len(cloudProviders.Results[0].InstanceSizes) == 0 {
		return nil, errors.New("no regions available")
	}

	availableRegions := cloudProviders.Results[0].InstanceSizes[0].AvailableRegions

	defaultRegions := make([]string, 0, len(availableRegions))
	popularRegionIndex := findPopularRegionIndex(availableRegions)

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

func findPopularRegionIndex(regions []*atlas.AvailableRegion) int {
	for i, v := range regions {
		if v.Default {
			return i
		}
	}
	return -1
}

// setupCloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by printing
// the dbUsername and dbPassword
func (opts *Opts) setupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Printf(quickstartTemplateCloseHandler, opts.DBUsername, opts.DBUserPassword)
		os.Exit(0)
	}()
}

// askDBUserOptions allows the user to set required flags by using interactive prompts
func (opts *Opts) askDBUserOptions() error {
	var qs []*survey.Question

	dbUser := opts.DBUsername
	if dbUser == "" {
		message := ""
		dbUser = autogeneratedName()
		if dbUser != "" {
			message = fmt.Sprintf(" [Press Enter to use '%s']", dbUser)
		}

		qs = append(qs, newDBUsernameQuestion(dbUser, message, opts.validateUniqueUsername))
	}

	if opts.DBUserPassword == "" {
		pwd, err := randgen.GenerateRandomBase64String(passwordLength)

		message := ""
		if err == nil {
			message = fmt.Sprintf(" [Press Enter to use an auto-generated password '%s']", pwd)
			opts.DBUserPassword = pwd
		}

		qs = append(qs, newDBUserPasswordQuestion(opts.DBUserPassword, message))
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
	if len(opts.IPAddresses) > 0 {
		return nil
	}

	message := ""
	publicIP := store.IPAddress()
	if publicIP != "" {
		message = fmt.Sprintf(" [Press Enter to use your public IP address '%s']", publicIP)
	}
	q := newAccessListQuestion(publicIP, message)

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
	if opts.SkipMongosh {
		return false, nil
	}

	fmt.Print(mongoShellDetails)

	runMongoShell := false
	prompt := newMongoShellQuestionAccessDeployment(opts.ClusterName)
	err := survey.AskOne(prompt, &runMongoShell)

	if !runMongoShell || err != nil {
		return false, err
	}

	if config.MongoShellPath() != "" {
		return true, nil
	}

	fmt.Println("No MongoDB shell version detected.")

	isInstalled := false
	prompt = newMongoShellQuestion()
	if err := survey.AskOne(prompt, &isInstalled); err != nil {
		return false, err
	}

	if isInstalled {
		wantToProvidePath := false
		prompt = newMongoShellQuestionProvidePath()
		if err := survey.AskOne(prompt, &wantToProvidePath); err != nil {
			return false, err
		}

		if wantToProvidePath {
			if err := askMongoShellAndSetConfig(); err != nil {
				return false, err
			}
		}
	} else {
		runShell, err := openMogoshDownloadPageAndSetPath()
		if !runShell || err != nil {
			return runShell, err
		}
	}

	return runMongoShell, nil
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

func openMogoshDownloadPageAndSetPath() (bool, error) {
	openURL := false
	prompt := newMongoShellQuestionOpenBrowser()
	if err := survey.AskOne(prompt, &openURL); err != nil {
		return false, err
	}

	if openURL {
		if err := browser.OpenURL(mongoshURL); err != nil {
			return false, err
		}

		if err := askMongoShellAndSetConfig(); err != nil {
			return false, err
		}
	} else {
		return false, nil
	}

	return true, nil
}

func askMongoShellAndSetConfig() error {
	var mongoShellPath string
	q := newMongoShellPathInput(mongosh.FindBinaryInPath(), mongosh.ValidateUniqueUsername)
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
	openBrowserAtlasAccount := false
	prompt := newAtlasAccountQuestionOpenBrowser()
	if err := survey.AskOne(prompt, &openBrowserAtlasAccount); err != nil {
		return err
	}

	if openBrowserAtlasAccount {
		if err := browser.OpenURL(atlasAccountURL); err != nil {
			return err
		}
	}

	openBrowserProfileDoc := false
	prompt = newProfileDocQuestionOpenBrowser()
	if err := survey.AskOne(prompt, &openBrowserProfileDoc); err != nil {
		return err
	}

	if openBrowserProfileDoc {
		if err := browser.OpenURL(profileDocURL); err != nil {
			return err
		}
	}

	return validate.Credentials()
}

func autogeneratedName() string {
	return "Quickstart-" + strconv.FormatInt(time.Now().Unix(), 10)
}

func (opts *Opts) convertFlags() {
	opts.Provider = strings.ToUpper(opts.Provider)
	opts.Region = strings.ReplaceAll(strings.ToUpper(opts.Region), "-", "_")
}

// mongocli atlas dbuser(s) quickstart [--clusterName clusterName] [--provider provider] [--region regionName] [--projectId projectId] [--username username] [--password password] [--skipMongosh skipMongosh]
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
			opts.convertFlags()
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ClusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().StringVar(&opts.Provider, flag.Provider, "", usage.Provider)
	cmd.Flags().StringVarP(&opts.Region, flag.Region, flag.RegionShort, "", usage.Region)
	cmd.Flags().StringSliceVar(&opts.IPAddresses, flag.AccessListIP, []string{}, usage.NetworkAccessListIPEntry)
	cmd.Flags().StringVar(&opts.DBUsername, flag.Username, "", usage.DBUsername)
	cmd.Flags().StringVar(&opts.DBUserPassword, flag.Password, "", usage.Password)
	cmd.Flags().BoolVar(&opts.SkipMongosh, flag.SkipMongosh, false, usage.SkipMongosh)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}

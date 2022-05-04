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
	"context"
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
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/mongosh"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/mongodb/mongocli/internal/validate"
	"github.com/spf13/cobra"
)

const quickstartTemplate = `
Now you can connect to your Atlas cluster with: mongosh -u %s -p %s %s

`
const quickstartTemplateCloseHandler = `
Enter 'atlas cluster watch %s' to learn when your cluster is available.
`

const quickstartTemplateStoreWarning = `
Please store your database authentication access details in a secure location: 
username: %s 
password: %s
`

const quickstartTemplateIntro = `Press [Enter] to use the default values.

Enter [?] on any option to get help.
`

const quickstartTemplateCluster = `
Creating your cluster... [It's safe to 'Ctrl + C']
`
const quickstartTemplateIPNotFound = `
We could not find your public IP address. To add your IP address run:
  mongocli atlas accesslist create`

const (
	replicaSet          = "REPLICASET"
	DefaultAtlasTier    = "M0"
	defaultAtlasGovTier = "M30"
	atlasAdmin          = "atlasAdmin"
	mongoshURL          = "https://www.mongodb.com/try/download/shell"
	defaultProvider     = "AWS"
	defaultRegion       = "US_EAST_1"
	defaultRegionGov    = "US_GOV_EAST_1"
)

type Opts struct {
	cli.GlobalOpts
	cli.WatchOpts
	defaultName         string
	ClusterName         string
	Tier                string
	Provider            string
	Region              string
	IPAddresses         []string
	IPAddressesResponse string
	DBUsername          string
	DBUserPassword      string
	SampleDataJobID     string
	SkipSampleData      bool
	SkipMongosh         bool
	runMongoShell       bool
	mongoShellInstalled bool
	defaultValue        bool
	Confirm             bool
	store               store.AtlasClusterQuickStarter
}

type quickstart struct {
	ClusterName    string
	Provider       string
	Region         string
	DBUsername     string
	DBUserPassword string
	IPAddresses    []string
	SkipSampleData bool
	SkipMongosh    bool
}

func (opts *Opts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *Opts) Run() error {
	values, dErr := opts.newDefaultValues()
	if dErr != nil {
		return dErr
	}

	if err := opts.askConfirmDefaultQuestion(values); err != nil || !opts.Confirm {
		fmt.Print(quickstartTemplateIntro)

		err = opts.interactiveSetup()
		if err != nil {
			return err
		}
	} else {
		opts.replaceWithDefaultSettings(values)
	}

	if err := opts.createDatabaseUser(); err != nil {
		return err
	}

	if err := opts.createAccessList(); err != nil {
		return err
	}

	fmt.Printf(`We are deploying %s...`, opts.ClusterName)
	if err := opts.createCluster(); err != nil {
		return err
	}

	fmt.Printf(quickstartTemplateStoreWarning, opts.DBUsername, opts.DBUserPassword)
	opts.setupCloseHandler()

	fmt.Print(quickstartTemplateCluster)

	// Watch cluster creation
	if er := opts.Watch(opts.clusterCreationWatcher); er != nil {
		return er
	}

	fmt.Print(quickstartTemplateCluster)

	fmt.Print("Cluster created.")

	if err := opts.loadSampleData(); err != nil {
		return err
	}

	if err := opts.askMongoShellQuestion(); err != nil {
		return err
	}
	// Get cluster's connection string
	cluster, err := opts.store.AtlasCluster(opts.ConfigProjectID(), opts.ClusterName)
	if err != nil {
		return err
	}

	fmt.Printf(quickstartTemplate, opts.DBUsername, opts.DBUserPassword, cluster.ConnectionStrings.StandardSrv)

	if opts.runMongoShell && config.MongoShellPath() != "" {
		return mongosh.Run(config.MongoShellPath(), opts.DBUsername, opts.DBUserPassword, cluster.ConnectionStrings.StandardSrv)
	}

	return nil
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

	opts.SampleDataJobID = sampleDataJob.ID

	return opts.Watch(opts.sampleDataWatcher)
}

func (opts *Opts) sampleDataWatcher() (bool, error) {
	result, err := opts.store.SampleDataStatus(opts.ConfigProjectID(), opts.SampleDataJobID)
	if err != nil {
		return false, err
	}
	if result.State == "FAILED" {
		return false, fmt.Errorf("failed to load data: %s", result.ErrorMessage)
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

func (opts *Opts) askSampleDataQuestion() error {
	if opts.SkipSampleData {
		return nil
	}

	q := newSampleDataQuestion(opts.ClusterName)
	var addSampleData bool
	if err := survey.AskOne(q, &addSampleData); err != nil {
		return err
	}
	opts.SkipSampleData = !addSampleData

	return nil
}

func askMongoShellAndSetConfig() error {
	var mongoShellPath string
	q := newMongoShellPathInput()
	if err := survey.AskOne(q, &mongoShellPath, survey.WithValidator(validate.Path)); err != nil {
		return err
	}

	config.SetMongoShellPath(mongoShellPath)
	return config.Save()
}

// setupCloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by printing
// the dbUsername and dbPassword.
func (opts *Opts) setupCloseHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Printf(quickstartTemplateCloseHandler, opts.ClusterName)
		os.Exit(0)
	}()
}

func (opts *Opts) providerAndRegionToConstant() {
	opts.Provider = strings.ToUpper(opts.Provider)
	opts.Region = strings.ReplaceAll(strings.ToUpper(opts.Region), "-", "_")
}

func (opts *Opts) setTier() {
	if config.CloudGovService == config.Service() && opts.Tier == DefaultAtlasTier {
		opts.Tier = defaultAtlasGovTier
	}
}

func (opts *Opts) newDefaultValues() (*quickstart, error) {
	values := &quickstart{}
	values.SkipMongosh = opts.SkipMongosh
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
		values.Region = defaultRegion
		if config.CloudGovService == config.Service() {
			values.Region = defaultRegionGov
		}
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
			_, _ = fmt.Fprintln(os.Stderr, quickstartTemplateIPNotFound)
		}
	}

	return values, nil
}

func (opts *Opts) replaceWithDefaultSettings(values *quickstart) {
	if values.ClusterName != "" {
		opts.ClusterName = values.ClusterName
	}

	if values.Provider != "" {
		opts.Provider = values.Provider
	}

	if values.Region != "" {
		opts.Region = values.Region
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

	opts.SkipSampleData = values.SkipSampleData
	opts.SkipMongosh = values.SkipMongosh
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

	if err := opts.askAccessListOptions(); err != nil {
		return err
	}

	return opts.askConfirmConfigQuestion()
}

// Builder
// mongocli atlas dbuser(s) quickstart
//	[--clusterName clusterName]
//	[--provider provider]
//	[--region regionName]
//	[--projectId projectId]
//	[--username username]
//	[--password password]
//	[--skipMongosh skipMongosh]
//	[--default]
func Builder() *cobra.Command {
	opts := &Opts{}
	cmd := &cobra.Command{
		Use: "quickstart",
		Example: `Skip setting cluster name, provider or database username by using the command options
  $ mongocli atlas quickstart --clusterName Test --provider GCP --username dbuserTest
`,
		Short: "Create and access an Atlas Cluster.",
		Long:  "This command creates a new cluster, adds your public IP to the atlas access list and creates a db user to access your new MongoDB instance.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.setTier()
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), ""),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			const base10 = 10
			opts.defaultName = "Cluster" + strconv.FormatInt(time.Now().Unix(), base10)[5:]
			opts.providerAndRegionToConstant()

			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ClusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().StringVar(&opts.Tier, flag.Tier, DefaultAtlasTier, usage.Tier)
	cmd.Flags().StringVar(&opts.Provider, flag.Provider, "", usage.Provider)
	cmd.Flags().StringVarP(&opts.Region, flag.Region, flag.RegionShort, "", usage.Region)
	cmd.Flags().StringSliceVar(&opts.IPAddresses, flag.AccessListIP, []string{}, usage.NetworkAccessListIPEntry)
	cmd.Flags().StringVar(&opts.DBUsername, flag.Username, "", usage.DBUsername)
	cmd.Flags().StringVar(&opts.DBUserPassword, flag.Password, "", usage.Password)
	cmd.Flags().BoolVar(&opts.SkipSampleData, flag.SkipSampleData, false, usage.SkipSampleData)
	cmd.Flags().BoolVar(&opts.SkipMongosh, flag.SkipMongosh, false, usage.SkipMongosh)
	cmd.Flags().BoolVarP(&opts.defaultValue, flag.Default, "Y", false, usage.QuickstartDefault)
	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}

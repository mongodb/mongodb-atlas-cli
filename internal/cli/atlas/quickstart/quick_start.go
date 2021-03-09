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
	"os/user"
	"regexp"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/convert"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/net"
	"github.com/mongodb/mongocli/internal/randgen"
	"github.com/mongodb/mongocli/internal/search"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const quickstartTemplate = "Now you can connect to your Atlas cluster with: mongo -u %s -p %s %s\n"

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
)

// DefaultRegions represents the regions available for each cloud service provider
var DefaultRegions = map[string][]string{
	"AWS":   {"US_EAST_1", "US_WEST_2", "AP_SOUTH_1", "AP_EAST_2", "EU_WEST_1", "EU_CENTRAL_1", "ME_SOUTH_1", "AF_SOUTH_1"},
	"GCP":   {"CENTRAL_US", "CANADA_CENTRAL", "WESTERN_EUROPE", "ASIA_SOUTH_EAST", "SOUTH_AFRICA_NORTH", "UAE_NORTH"},
	"AZURE": {"US_EAST_2", "US_WEST", "EUROPE_NORTH"},
}

type Opts struct {
	cli.GlobalOpts
	cli.WatchOpts
	ClusterName    string
	Provider       string
	Region         string
	IPAddresses    []string
	DBUsername     string
	DBUserPassword string
	store          store.AtlasClusterQuickStarter
}

func (opts *Opts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *Opts) Run() error {
	if err := opts.askClusterFlags(); err != nil {
		return err
	}

	if _, err := opts.store.CreateCluster(opts.newCluster()); err != nil {
		return err
	}

	if err := opts.askDBUserAccessListFlags(); err != nil {
		return err
	}

	if _, err := opts.store.CreateDatabaseUser(opts.newDatabaseUser()); err != nil {
		return err
	}

	// Add IP to project’s IP access list
	entries := opts.newProjectIPAccessList()
	if _, err := opts.store.CreateProjectIPAccessList(entries); err != nil {
		return err
	}

	fmt.Println("Creating your cluster...")
	if er := opts.Watch(opts.watcher); er != nil {
		return er
	}

	// Get cluster's connection string
	cluster, err := opts.store.AtlasCluster(opts.ConfigProjectID(), opts.ClusterName)
	if err != nil {
		return err
	}

	fmt.Printf(quickstartTemplate, opts.DBUsername, opts.DBUserPassword, cluster.SrvAddress)
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
	diskSizeGB := atlas.DefaultDiskSizeGB[strings.ToUpper(opts.Provider)]["M10"]
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

func (opts *Opts) askClusterFlags() error {
	var qs []*survey.Question

	message := "Insert the cluster name"
	clusterName := opts.ClusterName
	if clusterName == "" {
		clusterName = opts.newClusterName()
		if clusterName != "" {
			message = fmt.Sprintf("Insert the cluster name [Press Enter to use the auto-generated name '%s']", clusterName)
		}
		qs = append(qs, newClusterNameQuestion(clusterName, message))
	}

	if opts.Provider == "" {
		qs = append(qs, newClusterProviderQuestion())
	}

	if err := survey.Ask(qs, opts); err != nil {
		return err
	}

	if regionQ := newRegionQuestions(opts.Region, opts.Provider); regionQ != nil {
		// we call survey.Ask two times because the region question needs opts.Provider to be populated
		if err := survey.Ask([]*survey.Question{regionQ}, opts); err != nil {
			return err
		}
	}

	return nil
}

// askDBUserAccessListFlags allows the user to set required flags by using interactive prompts
func (opts *Opts) askDBUserAccessListFlags() error {
	var qs []*survey.Question

	message := "Insert the Username for authenticating to MongoDB"
	dbUser := opts.DBUsername
	if dbUser == "" {
		dbUser = dbUsername()
		if dbUser != "" {
			message = fmt.Sprintf("Insert the Username for authenticating to MongoDB [Press Enter to use '%s']", dbUser)
		}

		qs = append(qs, newDBUsernameQuestion(dbUser, message, opts.validateUniqueUsername))
	}

	if opts.DBUserPassword == "" {
		qs = append(qs, newDBUserPassword())
	}

	if opts.DBUserPassword == "" {
		// The user wants to auto-generate the password
		pwd, err := randgen.GenerateRandomBase64String(passwordLength)
		if err != nil {
			return err
		}
		opts.DBUserPassword = pwd
	}

	if len(opts.IPAddresses) == 0 {
		message = "Insert the IP entry to add to the Access List"
		publicIP := net.IPAddress()
		if publicIP != "" {
			message = fmt.Sprintf("Insert the IP entry to add to the Access List [Press Enter to use your public IP address '%s']", publicIP)
		}
		q := newAccessListQuestion(publicIP, message)
		qs = append(qs, q)
	}

	if len(qs) > 0 {
		if err := survey.Ask(qs, opts); err != nil {
			return err
		}
	}

	return nil
}

func (opts *Opts) validateUniqueUsername(val interface{}) error {
	username, _ := val.(string)
	dbUser, err := opts.store.DatabaseUser(convert.AdminDB, opts.ConfigProjectID(), username)
	if err != nil {
		if !strings.Contains(err.Error(), fmt.Sprintf("No user with username %s exists.", username)) {
			return err
		}
	}

	if dbUser != nil {
		return errors.New("a user with this username already exists")
	}

	return nil
}

// dbUsername returns the username of the user by running the command 'whoami'
func dbUsername() string {
	userStruct, err := user.Current()
	if err != nil {
		return ""
	}

	// dbUsername can only contain ASCII letters, numbers, hyphens and underscores
	out := strings.TrimSpace(userStruct.Username)
	var re = regexp.MustCompile("([^A-Za-z0-9_-])")
	return re.ReplaceAllString(out, "_")
}

// newClusterName returns an auto-generate Cluster name
func (opts *Opts) newClusterName() string {
	cs, _ := opts.store.ProjectClusters(opts.ConfigProjectID(), nil)
	i := 0
	if clusters, ok := cs.([]atlas.Cluster); ok {
		for {
			clusterName := "QuickstartCluster" + strconv.Itoa(i)
			if !search.IsClusterFound(clusters, clusterName) {
				return clusterName
			}
			i++
		}
	}

	return ""
}

// mongocli atlas dbuser(s) quickstart [--clusterName clusterName] [--provider provider] [--region regionName] [--projectId projectId] [--username username] [--password password]
func Builder() *cobra.Command {
	opts := &Opts{}
	cmd := &cobra.Command{
		Use: "quickstart",
		Example: `mongocli atlas quickstart
mongocli atlas quickstart --clusterName Test --provider GPC --username dbuserTest --password Test!
`,
		Short: QuickStart,
		Long:  LongQuickStart,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), ""),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ClusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().StringVar(&opts.Provider, flag.Provider, "", usage.Provider)
	cmd.Flags().StringVarP(&opts.Region, flag.Region, flag.RegionShort, "", usage.Region)
	cmd.Flags().StringSliceVar(&opts.IPAddresses, flag.IP, []string{}, usage.AccessListIPEntry)
	cmd.Flags().StringVar(&opts.DBUsername, flag.Username, "", usage.DBUsername)
	cmd.Flags().StringVar(&opts.DBUserPassword, flag.Password, "", usage.Password)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}

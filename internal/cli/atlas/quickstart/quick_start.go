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
	"os/exec"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/convert"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const quickstartTemplate = `Now you can connect to your Atlas cluster with:
mongo -u %s -p %s %s`
const (
	replicaSet        = "REPLICASET"
	diskSizeGB        = 10
	mdbVersion        = "4.2"
	shards            = 1
	tier              = "M10"
	members           = 3
	zoneName          = "Zone 1"
	accessListComment = "IP added with mongocli atlas quickstart"
	atlasAdmin        = "atlasAdmin"
	dbUsername        = "quickstartUser"
	dbUserPassword    = "Password1!"
	none              = "NONE"
)

type Opts struct {
	cli.GlobalOpts
	cli.WatchOpts
	clusterName      string
	provider         string
	region           string
	ipAddress        string
	dbUsername       string
	connectionString string
	store            store.AtlasClusterQuickStarter
}

func (opts *Opts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *Opts) Run() error {
	if _, err := opts.store.CreateCluster(opts.newCluster()); err != nil {
		return err
	}

	// Add IP to project’s IP access list
	entry, err := opts.newWhitelist()
	if err != nil {
		return err
	}

	if _, err = opts.store.CreateProjectIPAccessList(entry); err != nil {
		return err
	}

	// Create dbUser
	if _, err = opts.store.CreateDatabaseUser(opts.newDatabaseUser()); err != nil {
		return err
	}

	fmt.Println("Creating your cluster...")
	if er := opts.Watch(opts.watcher); er != nil {
		return er
	}

	// Get cluster's connection string
	cluster, err := opts.store.AtlasCluster(opts.ConfigProjectID(), opts.clusterName)
	if err != nil {
		return err
	}
	opts.connectionString = cluster.SrvAddress

	fmt.Printf(quickstartTemplate, opts.dbUsername, dbUserPassword, opts.connectionString)
	return opts.Print(nil)
}

func (opts *Opts) watcher() (bool, error) {
	result, err := opts.store.AtlasCluster(opts.ConfigProjectID(), opts.clusterName)
	if err != nil {
		return false, err
	}
	return result.StateName == "IDLE", nil
}

func (opts *Opts) newDatabaseUser() *atlas.DatabaseUser {
	if opts.dbUsername == "" {
		opts.dbUsername = dbUsername
	}

	return &atlas.DatabaseUser{
		Roles:        convert.BuildAtlasRoles([]string{atlasAdmin}),
		GroupID:      opts.ConfigProjectID(),
		Password:     dbUserPassword,
		X509Type:     none,
		AWSIAMType:   none,
		LDAPAuthType: none,
		DatabaseName: convert.AdminDB,
		Username:     opts.dbUsername,
	}
}

// newIPAddress returns client's public ip
func (opts *Opts) newIPAddress() (string, error) {
	command := "curl ifconfig.me"
	cmd := exec.Command("bash", "-c", command)
	stdout, err := cmd.Output()

	if err != nil {
		return "", errors.New("error in finding your public IP, please use --ip to provide your public ip")
	}

	return string(stdout), nil
}

func (opts *Opts) newWhitelist() (*atlas.ProjectIPWhitelist, error) {
	whitelist := &atlas.ProjectIPWhitelist{
		GroupID: opts.ConfigProjectID(),
		Comment: accessListComment,
	}

	if opts.ipAddress == "" {
		ipAddress, err := opts.newIPAddress()
		if err != nil {
			return nil, err
		}
		whitelist.IPAddress = ipAddress
	}

	return whitelist, nil
}

func (opts *Opts) newCluster() *atlas.Cluster {
	diskSizeGB := float64(diskSizeGB)
	return &atlas.Cluster{
		GroupID:             opts.ConfigProjectID(),
		ClusterType:         replicaSet,
		ReplicationSpecs:    []atlas.ReplicationSpec{opts.newReplicationSpec()},
		ProviderSettings:    opts.newProviderSettings(),
		MongoDBMajorVersion: mdbVersion,
		DiskSizeGB:          &diskSizeGB,
		Name:                opts.clusterName,
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
			opts.region: {
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
		InstanceSizeName: tier,
		ProviderName:     opts.provider,
		RegionName:       opts.region,
	}
}

// mongocli atlas dbuser(s) quickstart [--clusterName clusterName] [--provider provider] [--region regionName] [--projectId projectId]
func Builder() *cobra.Command {
	opts := &Opts{}
	cmd := &cobra.Command{
		Use:   "quickstart",
		Short: QuickStart,
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

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "GetStarted", usage.ClusterName)
	cmd.Flags().StringVar(&opts.provider, flag.Provider, "AWS", usage.Provider)
	cmd.Flags().StringVarP(&opts.region, flag.Region, flag.RegionShort, "US_EAST_1", usage.Region)
	cmd.Flags().StringVar(&opts.ipAddress, flag.IP, "", usage.AccessListIPEntry)
	cmd.Flags().StringVar(&opts.dbUsername, flag.Username, "", usage.DBUsername)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}

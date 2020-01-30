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

package cli

import (
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mcli/internal/flags"
	"github.com/mongodb/mcli/internal/json"
	"github.com/mongodb/mcli/internal/store"
	"github.com/mongodb/mcli/internal/usage"
	"github.com/spf13/cobra"
)

const (
	replicaSet        = "REPLICASET"
	tenant            = "TENANT"
	atlasM2           = "M2"
	atlasM5           = "M5"
	zoneName          = "Zone 1"
	currentMDBVersion = "4.2"
)

type atlasClustersCreateOpts struct {
	*globalOpts
	name         string
	provider     string
	region       string
	instanceSize string
	members      int64
	diskSizeGB   float64
	backup       bool
	mdbVersion   string
	store        store.ClusterCreator
}

func (opts *atlasClustersCreateOpts) init() error {
	if opts.ProjectID() == "" {
		return errMissingProjectID
	}

	s, err := store.New()
	if err != nil {
		return err
	}
	opts.store = s
	return nil
}

func (opts *atlasClustersCreateOpts) Run() error {
	cluster := opts.newCluster()
	result, err := opts.store.CreateCluster(cluster)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func (opts *atlasClustersCreateOpts) newCluster() *atlas.Cluster {
	replicationSpec := opts.newReplicationSpec()
	providerSettings := opts.newProviderSettings()

	cluster := &atlas.Cluster{
		BackupEnabled:       &opts.backup,
		ClusterType:         replicaSet,
		DiskSizeGB:          &opts.diskSizeGB,
		GroupID:             opts.ProjectID(),
		MongoDBMajorVersion: opts.mdbVersion,
		Name:                opts.name,
		ProviderSettings:    providerSettings,
		ReplicationSpecs:    []atlas.ReplicationSpec{replicationSpec},
	}
	return cluster
}

func (opts *atlasClustersCreateOpts) newProviderSettings() *atlas.ProviderSettings {
	providerName := opts.providerName()

	var backingProviderName string
	if providerName == tenant {
		backingProviderName = opts.provider
	}

	return &atlas.ProviderSettings{
		InstanceSizeName:    opts.instanceSize,
		ProviderName:        providerName,
		RegionName:          opts.region,
		BackingProviderName: backingProviderName,
	}
}

func (opts *atlasClustersCreateOpts) providerName() string {
	if opts.instanceSize == atlasM2 || opts.instanceSize == atlasM5 {
		return tenant
	}
	return opts.provider
}

func (opts *atlasClustersCreateOpts) newReplicationSpec() atlas.ReplicationSpec {
	var (
		readOnlyNodes int64 = 0
		NumShards     int64 = 1
		Priority      int64 = 7
	)
	replicationSpec := atlas.ReplicationSpec{
		NumShards: &NumShards,
		ZoneName:  zoneName,
		RegionsConfig: map[string]atlas.RegionsConfig{
			opts.region: {
				ReadOnlyNodes:  &readOnlyNodes,
				ElectableNodes: &opts.members,
				Priority:       &Priority,
			},
		},
	}
	return replicationSpec
}

// mcli atlas cluster(s) create name --projectId projectId --provider AWS|GCP|AZURE --region regionName [--members N] [--instanceSize M#] [--diskSizeGB N] [--backup] [--mdbVersion]
func AtlasClustersCreateBuilder() *cobra.Command {
	opts := &atlasClustersCreateOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:     "create [name]",
		Short:   "Create a MongoDB cluster in Atlas.",
		Example: `  mcli atlas cluster create myCluster --projectId=1 --region US_EAST_1 --members 3 --instanceSize M2 --provider AWS --mdbVersion 4.2 --diskSizeGB 2`,
		Args:    cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.provider, flags.Provider, "", usage.Provider)
	cmd.Flags().StringVarP(&opts.region, flags.Region, flags.RegionShort, "", usage.Region)
	cmd.Flags().Int64VarP(&opts.members, flags.Members, flags.MembersShort, 3, usage.Members)
	cmd.Flags().StringVar(&opts.instanceSize, flags.InstanceSize, atlasM2, usage.InstanceSize)
	cmd.Flags().Float64Var(&opts.diskSizeGB, flags.DiskSizeGB, 2, usage.DiskSizeGB)
	cmd.Flags().StringVar(&opts.mdbVersion, flags.MDBVersion, currentMDBVersion, usage.MDBVersion)
	cmd.Flags().BoolVar(&opts.backup, flags.Backup, false, usage.Backup)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flags.Provider)
	_ = cmd.MarkFlagRequired(flags.Region)

	return cmd
}

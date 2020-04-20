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
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/file"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/afero"
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
	globalOpts
	name         string
	provider     string
	region       string
	instanceSize string
	members      int64
	diskSizeGB   float64
	backup       bool
	mdbVersion   string
	filename     string
	fs           afero.Fs
	store        store.ClusterCreator
}

func (opts *atlasClustersCreateOpts) initStore() error {
	var err error
	opts.store, err = store.New()
	return err
}

func (opts *atlasClustersCreateOpts) Run() error {
	cluster, err := opts.newCluster()
	if err != nil {
		return err
	}
	result, err := opts.store.CreateCluster(cluster)
	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func (opts *atlasClustersCreateOpts) newCluster() (*atlas.Cluster, error) {
	cluster := new(atlas.Cluster)
	if opts.filename != "" {
		if err := file.Load(opts.fs, opts.filename, cluster); err != nil {
			return nil, err
		}
	} else {
		cluster.ClusterType = replicaSet
		opts.applyOpts(cluster)
	}

	if opts.name != "" {
		cluster.Name = opts.name
	}

	cluster.GroupID = opts.ProjectID()
	return cluster, nil
}

func (opts *atlasClustersCreateOpts) applyOpts(out *atlas.Cluster) {
	replicationSpec := opts.newReplicationSpec()
	out.BackupEnabled = &opts.backup
	out.DiskSizeGB = &opts.diskSizeGB
	out.MongoDBMajorVersion = opts.mdbVersion
	out.ProviderSettings = opts.newProviderSettings()
	out.ReplicationSpecs = []atlas.ReplicationSpec{replicationSpec}
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

// AtlasClustersCreateBuilder builds a cobra.Command that can run as:
// create <name> --projectId projectId --provider AWS|GCP|AZURE --region regionName [--members N] [--instanceSize M#] [--diskSizeGB N] [--backup] [--mdbVersion]
func AtlasClustersCreateBuilder() *cobra.Command {
	opts := &atlasClustersCreateOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:     "create [name]",
		Short:   description.CreateCluster,
		Example: `  mongocli atlas cluster create myCluster --projectId=<projectId> --region US_EAST_1 --members 3 --instanceSize M2 --provider AWS --mdbVersion 4.2 --diskSizeGB 2`,
		Args:    cobra.MaximumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if opts.filename == "" {
				_ = cmd.MarkFlagRequired(flags.Provider)
				_ = cmd.MarkFlagRequired(flags.Region)

				if len(args) == 0 {
					return errMissingClusterName
				}
			}
			if len(args) != 0 {
				opts.name = args[0]
			}
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
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
	cmd.Flags().StringVarP(&opts.filename, flags.File, flags.FileShort, "", usage.Filename)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}

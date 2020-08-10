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

package clusters

import (
	"errors"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/file"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/output"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const (
	replicaSet        = "REPLICASET"
	tenant            = "TENANT"
	atlasM2           = "M2"
	atlasM5           = "M5"
	zoneName          = "Zone 1"
	currentMDBVersion = "4.2"
	labelKey          = "Infrastructure Tool"
	labelValue        = "mongoCLI"
)

type CreateOpts struct {
	cli.GlobalOpts
	name        string
	provider    string
	region      string
	tier        string
	members     int64
	shards      int64
	clusterType string
	diskSizeGB  float64
	backup      bool
	mdbVersion  string
	filename    string
	fs          afero.Fs
	store       store.ClusterCreator
}

func (opts *CreateOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

var createTmpl = "Deploying cluster {{.Name}}.\n"

func (opts *CreateOpts) Run() error {
	cluster, err := opts.newCluster()
	if err != nil {
		return err
	}
	r, err := opts.store.CreateCluster(cluster)
	if err != nil {
		return err
	}

	return output.Print(config.Default(), createTmpl, r)
}

func (opts *CreateOpts) newCluster() (*atlas.Cluster, error) {
	cluster := new(atlas.Cluster)
	if opts.filename != "" {
		if err := file.Load(opts.fs, opts.filename, cluster); err != nil {
			return nil, err
		}
	} else {
		opts.applyOpts(cluster)
	}

	if opts.name != "" {
		cluster.Name = opts.name
	}

	updateLabels(cluster)

	cluster.GroupID = opts.ConfigProjectID()
	return cluster, nil
}

func updateLabels(out *atlas.Cluster) {
	found := false
	for _, v := range out.Labels {
		if v.Key == labelKey && v.Value == labelValue {
			found = true
			break
		}
	}

	if !found {
		out.Labels = append(out.Labels, atlas.Label{
			Key:   labelKey,
			Value: labelValue,
		})
	}
}

func (opts *CreateOpts) applyOpts(out *atlas.Cluster) {
	replicationSpec := opts.newReplicationSpec()
	if opts.backup {
		out.ProviderBackupEnabled = &opts.backup
		out.PitEnabled = &opts.backup
	}
	out.ClusterType = opts.clusterType
	out.DiskSizeGB = &opts.diskSizeGB
	out.MongoDBMajorVersion = opts.mdbVersion
	out.ProviderSettings = opts.newProviderSettings()
	out.ReplicationSpecs = []atlas.ReplicationSpec{replicationSpec}
}

func (opts *CreateOpts) newProviderSettings() *atlas.ProviderSettings {
	providerName := opts.providerName()

	var backingProviderName string
	if providerName == tenant {
		backingProviderName = opts.provider
	}

	return &atlas.ProviderSettings{
		InstanceSizeName:    opts.tier,
		ProviderName:        providerName,
		RegionName:          opts.region,
		BackingProviderName: backingProviderName,
	}
}

func (opts *CreateOpts) providerName() string {
	if opts.tier == atlasM2 || opts.tier == atlasM5 {
		return tenant
	}
	return opts.provider
}

func (opts *CreateOpts) newReplicationSpec() atlas.ReplicationSpec {
	var (
		readOnlyNodes int64 = 0
		Priority      int64 = 7
	)
	replicationSpec := atlas.ReplicationSpec{
		NumShards: &opts.shards,
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

// CreateBuilder builds a cobra.Command that can run as:
// create <name> --projectId projectId --provider AWS|GCP|AZURE --region regionName [--members N] [--tier M#] [--diskSizeGB N] [--backup] [--mdbVersion]
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "create [name]",
		Short: CreateCluster,
		Long:  CreateClusterLong,
		Example: `  
  Deploy a 3 members replica set in AWS
  $ mongocli atlas cluster create <clusterName> --projectId <projectId> --provider AWS --region US_EAST_1 --members 3 --tier M10 --mdbVersion 4.2 --diskSizeGB 10

  Deploy a 3 members replica set in AZURE
  $ mongocli atlas cluster create <clusterName> --projectId <projectId> --provider AZURE --region US_EAST_2 --members 3 --tier M10  --mdbVersion 4.2 --diskSizeGB 10
  
  Deploy a 3 members replica set in GCP
  $ mongocli atlas cluster create <clusterName> --projectId <projectId> --provider GCP --region EASTERN_US --members 3 --tier M10  --mdbVersion 4.2 --diskSizeGB 10

  Deploy a cluster from a config file
  $ mongocli atlas cluster create --projectId <projectId> --file <path/to/cluster.json>
`,
		Args: cobra.MaximumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if opts.filename == "" {
				_ = cmd.MarkFlagRequired(flag.Provider)
				_ = cmd.MarkFlagRequired(flag.Region)

				if len(args) == 0 {
					return errors.New("cluster name missing")
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

	cmd.Flags().StringVar(&opts.provider, flag.Provider, "", usage.Provider)
	cmd.Flags().StringVarP(&opts.region, flag.Region, flag.RegionShort, "", usage.Region)
	cmd.Flags().Int64VarP(&opts.members, flag.Members, flag.MembersShort, 3, usage.Members)
	cmd.Flags().StringVar(&opts.tier, flag.Tier, atlasM2, usage.Tier)
	cmd.Flags().Float64Var(&opts.diskSizeGB, flag.DiskSizeGB, 2, usage.DiskSizeGB)
	cmd.Flags().StringVar(&opts.mdbVersion, flag.MDBVersion, currentMDBVersion, usage.MDBVersion)
	cmd.Flags().BoolVar(&opts.backup, flag.Backup, false, usage.Backup)
	cmd.Flags().StringVarP(&opts.filename, flag.File, flag.FileShort, "", usage.Filename)
	cmd.Flags().StringVar(&opts.clusterType, flag.Type, replicaSet, usage.ClusterTypes)
	cmd.Flags().Int64VarP(&opts.shards, flag.Shards, flag.ShardsShort, 1, usage.Shards)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagFilename(flag.File)

	return cmd
}

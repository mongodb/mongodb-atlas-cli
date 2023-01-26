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
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/commonerrors"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const (
	updateTmpl = "Updating cluster '{{.Name}}'.\n"
)

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	name       string
	tier       string
	diskSizeGB float64
	mdbVersion string
	filename   string
	fs         afero.Fs
	store      store.AtlasClusterGetterUpdater
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *UpdateOpts) Run() error {
	cluster, err := opts.cluster()
	if err != nil {
		return err
	}
	if opts.filename == "" {
		opts.patchOpts(cluster)
	}

	r, err := opts.store.UpdateCluster(opts.ConfigProjectID(), opts.name, cluster)
	if err != nil {
		return commonerrors.Check(err)
	}

	return opts.Print(r)
}

func (opts *UpdateOpts) cluster() (*atlas.AdvancedCluster, error) {
	var cluster *atlas.AdvancedCluster
	if opts.filename != "" {
		err := file.Load(opts.fs, opts.filename, &cluster)
		if err != nil {
			return nil, err
		}
		if opts.name == "" {
			opts.name = cluster.Name
		}
		return cluster, nil
	}
	return opts.store.AtlasCluster(opts.ProjectID, opts.name)
}

func (opts *UpdateOpts) patchOpts(out *atlas.AdvancedCluster) {
	RemoveReadOnlyAttributes(out)
	if opts.mdbVersion != "" {
		out.MongoDBMajorVersion = opts.mdbVersion
	}
	if opts.diskSizeGB > 0 {
		out.DiskSizeGB = &opts.diskSizeGB
	}
	if opts.tier != "" {
		opts.addTierToAdvancedCluster(out)
	}

	AddLabel(out, NewCLILabel())
}

func (opts *UpdateOpts) addTierToAdvancedCluster(out *atlas.AdvancedCluster) {
	for _, replicationSpec := range out.ReplicationSpecs {
		for _, regionConf := range replicationSpec.RegionConfigs {
			if regionConf.ReadOnlySpecs != nil {
				regionConf.ReadOnlySpecs.InstanceSize = opts.tier
			}
			if regionConf.AnalyticsSpecs != nil {
				regionConf.AnalyticsSpecs.InstanceSize = opts.tier
			}
			if regionConf.ElectableSpecs != nil {
				regionConf.ElectableSpecs.InstanceSize = opts.tier
			}
		}
	}
}

// mongocli atlas cluster(s) update [clusterName] --projectId projectId [--tier M#] [--diskSizeGB N] [--mdbVersion].
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "update [clusterName]",
		Short: "Modify the settings of the specified cluster.",
		Long: `You can specify modifications in a JSON configuration file with the --file flag.
		
You can modify only M10 or larger clusters that are single-region replica sets.
		
You can't change the name of the cluster or downgrade the MongoDB version of your cluster.`,
		Example: fmt.Sprintf(`  # Update the tier for a cluster named myCluster for the project with ID 5e2211c17a3e5a48f5497de3:
  %[1]s cluster update myCluster --projectId 5e2211c17a3e5a48f5497de3 --tier M50

  # Update the disk size for a cluster named myCluster for the project with ID 5e2211c17a3e5a48f5497de3:
  %[1]s cluster update myCluster --projectId 5e2211c17a3e5a48f5497de3 --diskSizeGB 20

  # Update the MongoDB version for a cluster named myCluster for the project with ID 5e2211c17a3e5a48f5497de3:
  %[1]s cluster update myCluster --projectId 5e2211c17a3e5a48f5497de3 --mdbVersion 5.0
  
  # Use a configuration file named cluster-config.json to update a cluster named myCluster for the project with ID 5e2211c17a3e5a48f5497de3:
  %[1]s cluster update myCluster --projectId 5e2211c17a3e5a48f5497de3 --file cluster-config.json --output json`,
			cli.ExampleAtlasEntryPoint()),
		Args: require.MaximumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				opts.name = args[0]
			}
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTmpl),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the cluster to update.",
		},
	}

	cmd.Flags().StringVar(&opts.tier, flag.Tier, "", usage.Tier)
	cmd.Flags().Float64Var(&opts.diskSizeGB, flag.DiskSizeGB, 0, usage.DiskSizeGB)
	cmd.Flags().StringVar(&opts.mdbVersion, flag.MDBVersion, "", usage.MDBVersion)
	cmd.Flags().StringVarP(&opts.filename, flag.File, flag.FileShort, "", usage.ClusterFilename)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagFilename(flag.File)

	return cmd
}

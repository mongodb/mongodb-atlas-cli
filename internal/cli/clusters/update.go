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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/commonerrors"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const (
	updateTmpl = "Updating cluster '{{.Name}}'.\n"
)

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	name                         string
	tier                         string
	diskSizeGB                   float64
	mdbVersion                   string
	enableTerminationProtection  bool
	disableTerminationProtection bool
	filename                     string
	tag                          map[string]string
	fs                           afero.Fs
	store                        store.AtlasClusterGetterUpdater
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

	removeReadOnlyAttributes(cluster)

	if opts.filename == "" {
		opts.patchOpts(cluster)
	}

	r, err := opts.store.UpdateCluster(opts.ConfigProjectID(), opts.name, cluster)
	if err != nil {
		return commonerrors.Check(err)
	}

	return opts.Print(r)
}

func (opts *UpdateOpts) cluster() (*atlasv2.AdvancedClusterDescription, error) {
	var cluster *atlasv2.AdvancedClusterDescription
	if opts.filename != "" {
		err := file.Load(opts.fs, opts.filename, &cluster)
		if err != nil {
			return nil, err
		}
		if opts.name == "" {
			opts.name = cluster.GetName()
		}

		// The cluster name cannot be updated by the Update operation
		if cluster.GetName() != "" && opts.name != cluster.GetName() {
			cluster.Name = nil
		}

		return cluster, nil
	}
	return opts.store.AtlasCluster(opts.ProjectID, opts.name)
}

func (opts *UpdateOpts) patchOpts(out *atlasv2.AdvancedClusterDescription) {
	if opts.mdbVersion != "" {
		out.MongoDBMajorVersion = &opts.mdbVersion
	}
	if opts.diskSizeGB > 0 {
		out.DiskSizeGB = &opts.diskSizeGB
	}
	if opts.tier != "" {
		opts.addTierToAdvancedCluster(out)
	}
	out.TerminationProtectionEnabled = cli.ReturnValueForSetting(opts.enableTerminationProtection, opts.disableTerminationProtection)

	if len(opts.tag) > 0 {
		out.Tags = &[]atlasv2.ResourceTag{}
	}
	addTags(out, opts.tag)
}

func (opts *UpdateOpts) addTierToAdvancedCluster(out *atlasv2.AdvancedClusterDescription) {
	for _, replicationSpec := range out.GetReplicationSpecs() {
		for regionIdx := range replicationSpec.GetRegionConfigs() {
			regionConf := (*replicationSpec.RegionConfigs)[regionIdx]
			if regionConf.ReadOnlySpecs != nil {
				regionConf.ReadOnlySpecs.InstanceSize = &opts.tier
			}
			if regionConf.AnalyticsSpecs != nil {
				regionConf.AnalyticsSpecs.InstanceSize = &opts.tier
			}
			if regionConf.ElectableSpecs != nil {
				regionConf.ElectableSpecs.InstanceSize = &opts.tier
			}
		}
	}
}

// UpdateBuilder atlas cluster(s) update [clusterName] --projectId projectId [--tier M#] [--diskSizeGB N] [--mdbVersion] [--tag key=value].
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "update [clusterName]",
		Short: "Modify the settings of the specified cluster.",
		Long: `You can specify modifications in a JSON configuration file with the --file flag.
		
You can't change the name of the cluster or downgrade the MongoDB version of your cluster.

You can only update a replica set to a single-shard cluster; you cannot update a replica set to a multi-sharded cluster. To learn more, see https://www.mongodb.com/docs/atlas/scale-cluster/#convert-a-replica-set-to-a-sharded-cluster and https://www.mongodb.com/docs/upcoming/tutorial/convert-replica-set-to-replicated-shard-cluster.

` + fmt.Sprintf("%s\n%s", fmt.Sprintf(usage.RequiredRole, "Project Cluster Manager"), "Atlas supports this command only for M10+ clusters"),
		Example: `  # Update the tier for a cluster named myCluster for the project with ID 5e2211c17a3e5a48f5497de3:
  atlas cluster update myCluster --projectId 5e2211c17a3e5a48f5497de3 --tier M50

  # Replace tags cluster named myCluster for the project with ID 5e2211c17a3e5a48f5497de3:
  atlas cluster update myCluster --projectId 5e2211c17a3e5a48f5497de3 --tag key1=value1

  # Remove all tags from cluster named myCluster for the project with ID 5e2211c17a3e5a48f5497de3:
  atlas cluster update myCluster --projectId 5e2211c17a3e5a48f5497de3 --tag =

  # Update the disk size for a cluster named myCluster for the project with ID 5e2211c17a3e5a48f5497de3:
  atlas cluster update myCluster --projectId 5e2211c17a3e5a48f5497de3 --diskSizeGB 20

  # Update the MongoDB version for a cluster named myCluster for the project with ID 5e2211c17a3e5a48f5497de3:
  atlas cluster update myCluster --projectId 5e2211c17a3e5a48f5497de3 --mdbVersion 5.0
  
  # Use a configuration file named cluster-config.json to update a cluster named myCluster for the project with ID 5e2211c17a3e5a48f5497de3:
  atlas cluster update myCluster --projectId 5e2211c17a3e5a48f5497de3 --file cluster-config.json --output json`,
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
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the cluster to update.",
			"output":          updateTmpl,
		},
	}

	cmd.Flags().StringVar(&opts.tier, flag.Tier, "", usage.Tier)
	cmd.Flags().Float64Var(&opts.diskSizeGB, flag.DiskSizeGB, 0, usage.DiskSizeGB)
	cmd.Flags().StringVar(&opts.mdbVersion, flag.MDBVersion, "", usage.MDBVersion)
	cmd.Flags().StringVarP(&opts.filename, flag.File, flag.FileShort, "", usage.ClusterFilename)

	cmd.Flags().BoolVar(&opts.enableTerminationProtection, flag.EnableTerminationProtection, false, usage.EnableTerminationProtection)
	cmd.Flags().BoolVar(&opts.disableTerminationProtection, flag.DisableTerminationProtection, false, usage.DisableTerminationProtection)
	cmd.MarkFlagsMutuallyExclusive(flag.EnableTerminationProtection, flag.DisableTerminationProtection)
	cmd.Flags().StringToStringVar(&opts.tag, flag.Tag, nil, usage.Tag+usage.UpdateWarning)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagFilename(flag.File)

	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.Tier)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.DiskSizeGB)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.EnableTerminationProtection)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.DisableTerminationProtection)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.Tag)

	autocomplete := &autoCompleteOpts{}
	_ = cmd.RegisterFlagCompletionFunc(flag.Tier, autocomplete.autocompleteTier())

	return cmd
}

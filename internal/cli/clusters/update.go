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
	"os"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/commonerrors"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/validate"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312003/admin"
)

const (
	updateTmpl = "Updating cluster '{{.Name}}'.\n"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=update_mock_test.go -package=clusters . AtlasClusterGetterUpdater

type AtlasClusterGetterUpdater interface {
	AtlasCluster(string, string) (*atlasClustersPinned.AdvancedClusterDescription, error)
	LatestAtlasCluster(string, string) (*atlasv2.ClusterDescription20240805, error)
	FlexCluster(string, string) (*atlasv2.FlexClusterDescription20241113, error)
	UpdateCluster(string, string, *atlasClustersPinned.AdvancedClusterDescription) (*atlasClustersPinned.AdvancedClusterDescription, error)
	UpdateFlexCluster(string, string, *atlasv2.FlexClusterDescriptionUpdate20241113) (*atlasv2.FlexClusterDescription20241113, error)
	UpdateClusterLatest(string, string, *atlasv2.ClusterDescription20240805) (*atlasv2.ClusterDescription20240805, error)
	GetClusterAutoScalingConfig(string, string) (*atlasv2.ClusterDescriptionAutoScalingModeConfiguration, error)
}

type UpdateOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	name                         string
	tier                         string
	diskSizeGB                   float64
	mdbVersion                   string
	autoScalingMode              string
	enableTerminationProtection  bool
	disableTerminationProtection bool
	isFlexCluster                bool
	filename                     string
	tag                          map[string]string
	fs                           afero.Fs
	store                        AtlasClusterGetterUpdater
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *UpdateOpts) Run() error {
	if opts.isFlexCluster {
		return opts.RunFlexCluster()
	}

	// Update will override the autoscaling mode to ISS if the flag is set to ISS
	if isIndependentShardScaling(opts.autoScalingMode) {
		return opts.RunDedicatedIndependentShardScaling()
	}

	// Get the cluster auto scaling config
	targetClusterAutoScalingConfig, err := opts.store.GetClusterAutoScalingConfig(opts.ConfigProjectID(), opts.name)
	if err != nil {
		targetClusterAutoScalingConfig = &atlasv2.ClusterDescriptionAutoScalingModeConfiguration{
			AutoScalingMode: &opts.autoScalingMode,
		}
	}
	appendAutoScalingModeTelemetry(opts.autoScalingMode)

	// If the flag is set to cluster wide scaling, warn the user that they are using the wrong flag
	if isIndependentShardScaling(targetClusterAutoScalingConfig.GetAutoScalingMode()) {
		if isClusterWideScaling(opts.autoScalingMode) {
			fmt.Fprintf(os.Stderr, "'independentShardScaling' autoscaling cluster detected, updating it to clusterWideScaling is not possible, use  --autoScalingMode 'independentShardScaling' instead")
		}
	}

	return opts.RunDedicatedClusterWideScaling()
}

func (opts *UpdateOpts) RunDedicatedIndependentShardScaling() error {
	cluster, err := opts.clusterLatest()
	if err != nil {
		return err
	}

	removeReadOnlyAttributesLatest(cluster)

	r, err := opts.store.UpdateClusterLatest(opts.ConfigProjectID(), opts.name, cluster)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *UpdateOpts) RunFlexCluster() error {
	flexClusterReq, err := opts.newFlexCluster()
	if err != nil {
		return err
	}

	flexCluster, err = opts.store.UpdateFlexCluster(opts.ConfigProjectID(), opts.name, flexClusterReq)

	apiError, ok := atlasv2.AsError(err)
	code := apiError.GetErrorCode()
	if ok {
		if code == invalidAttributeErrorCode && strings.Contains(apiError.GetDetail(), regionName) {
			return cli.ErrNoRegionExistsTryCommand
		}
		if code == duplicateClusterNameErrorCode {
			return cli.ErrNameExists
		}

		return err
	}

	return opts.Print(flexCluster)
}

func (opts *UpdateOpts) newFlexCluster() (*atlasv2.FlexClusterDescriptionUpdate20241113, error) {
	cluster := new(atlasv2.FlexClusterDescriptionUpdate20241113)
	if opts.filename != "" {
		if err := file.Load(opts.fs, opts.filename, cluster); err != nil {
			return nil, err
		}
	} else {
		flexClusterFromGet, err := opts.store.FlexCluster(opts.ConfigProjectID(), opts.name)
		if err != nil {
			return nil, err
		}
		cluster = opts.newFlexClusterDescriptionUpdate20241113(flexClusterFromGet)
	}

	return cluster, nil
}

func (opts *UpdateOpts) newFlexClusterDescriptionUpdate20241113(cluster *atlasv2.FlexClusterDescription20241113) *atlasv2.FlexClusterDescriptionUpdate20241113 {
	out := &atlasv2.FlexClusterDescriptionUpdate20241113{}

	if opts.disableTerminationProtection {
		out.TerminationProtectionEnabled = pointer.Get(false)
	}

	if opts.enableTerminationProtection {
		out.TerminationProtectionEnabled = pointer.Get(true)
	}

	// add existing tags
	if cluster.Tags != nil && len(*cluster.Tags) > 0 {
		out.SetTags(cluster.GetTags())
	}

	if len(opts.tag) > 0 {
		tags := newResourceTags(opts.tag)
		if out.HasTags() {
			newTags := append(out.GetTags(), *tags...)
			out.SetTags(newTags)
		} else {
			out.SetTags(*tags)
		}
	}

	return out
}

func (opts *UpdateOpts) RunDedicatedClusterWideScaling() error {
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

func (opts *UpdateOpts) clusterLatest() (*atlasv2.ClusterDescription20240805, error) {
	var cluster *atlasv2.ClusterDescription20240805
	var err error
	if opts.filename != "" {
		err = file.Load(opts.fs, opts.filename, &cluster)
		if err != nil {
			return nil, err
		}
	} else {
		cluster, err = opts.store.LatestAtlasCluster(opts.ConfigProjectID(), opts.name)
		if err != nil {
			return nil, err
		}
	}
	return cluster, nil
}

func (opts *UpdateOpts) cluster() (*atlasClustersPinned.AdvancedClusterDescription, error) {
	var cluster *atlasClustersPinned.AdvancedClusterDescription
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
	return opts.store.AtlasCluster(opts.ConfigProjectID(), opts.name)
}

func (opts *UpdateOpts) patchOpts(out *atlasClustersPinned.AdvancedClusterDescription) {
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
		out.Tags = &[]atlasClustersPinned.ResourceTag{}
	}
	addTags(out, opts.tag)
}

func (opts *UpdateOpts) addTierToAdvancedCluster(out *atlasClustersPinned.AdvancedClusterDescription) {
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

// newIsFlexCluster sets the opts.isFlexCluster that indicates if the cluster to create is
// a FlexCluster. The function uses the AtlasCluster to get the cluster description and in the event of the error
// cannotUseFlexWithClusterApisErrorCode sets the opts.isFlexCluster to true.
func (opts *UpdateOpts) newIsFlexCluster() error {
	_, err := opts.store.AtlasCluster(opts.ConfigProjectID(), opts.name)
	if err == nil {
		opts.isFlexCluster = false
		return nil
	}

	apiError, ok := atlasClustersPinned.AsError(err)
	if !ok {
		return err
	}
	if *apiError.ErrorCode != cannotUseFlexWithClusterApisErrorCode {
		return err
	}

	opts.isFlexCluster = true
	return nil
}

func (opts *UpdateOpts) validateTier() error {
	opts.tier = strings.ToUpper(opts.tier)
	if opts.tier == atlasM2 || opts.tier == atlasM5 {
		_, _ = fmt.Fprintf(os.Stderr, deprecateMessageSharedTier, opts.tier)
	}
	return nil
}

func (opts *UpdateOpts) validateAutoScalingMode() error {
	if opts.filename != "" {
		opts.autoScalingMode = detectIsFileISS(opts.fs, opts.filename)
	}

	return validate.AutoScalingMode(opts.autoScalingMode)()
}

// UpdateBuilder builds a cobra.Command that can run as:
// atlas cluster(s) update [clusterName] --projectId projectId [--tier M#] [--diskSizeGB N] [--mdbVersion] [--tag key=value].
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

Deprecation note: the M2 and M5 tiers are now deprecated; when selecting M2 or M5, a FLEX tier will be created instead. For the migration guide, visit: https://dochub.mongodb.org/core/flex-migration.\n

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
				opts.validateTier,
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTmpl),
				opts.newIsFlexCluster,
				opts.validateAutoScalingMode,
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
	cmd.Flags().StringVar(&opts.autoScalingMode, flag.AutoScalingMode, "", usage.AutoScalingMode)
	cmd.MarkFlagsMutuallyExclusive(flag.EnableTerminationProtection, flag.DisableTerminationProtection)
	cmd.Flags().StringToStringVar(&opts.tag, flag.Tag, nil, usage.Tag+usage.UpdateWarning)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagFilename(flag.File)

	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.Tier)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.DiskSizeGB)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.EnableTerminationProtection)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.DisableTerminationProtection)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.Tag)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.AutoScalingMode)

	autocomplete := &autoCompleteOpts{}
	_ = cmd.RegisterFlagCompletionFunc(flag.Tier, autocomplete.autocompleteTier())

	return cmd
}

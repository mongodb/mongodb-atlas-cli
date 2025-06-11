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
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/watchers"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312003/admin"
)

const (
	replicaSet                    = "REPLICASET"
	tenant                        = "TENANT"
	atlasM0                       = "M0"
	atlasM2                       = "M2"
	atlasFlex                     = "FLEX"
	atlasM5                       = "M5"
	zoneName                      = "Zone 1"
	invalidAttributeErrorCode     = "INVALID_ATTRIBUTE"
	duplicateClusterNameErrorCode = "DUPLICATE_CLUSTER_NAME"
	regionName                    = "regionName"
	priority                      = 7
	readOnlyNode                  = 0
	independentShardScalingFlag   = "independentShardScaling"
	clusterWideScalingFlag        = "clusterWideScaling"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=create_mock_test.go -package=clusters . ClusterCreator

type ClusterCreator interface {
	CreateCluster(v15 *atlasClustersPinned.AdvancedClusterDescription) (*atlasClustersPinned.AdvancedClusterDescription, error)
	CreateFlexCluster(string, *atlasv2.FlexClusterDescriptionCreate20241113) (*atlasv2.FlexClusterDescription20241113, error)
	CreateClusterLatest(*atlasv2.ClusterDescription20240805) (*atlasv2.ClusterDescription20240805, error)
}

type CreateOpts struct {
	cli.ProjectOpts
	cli.WatchOpts
	name                        string
	provider                    string
	region                      string
	tier                        string
	members                     int
	shards                      int
	clusterType                 string
	diskSizeGB                  float64
	backup                      bool
	biConnector                 bool
	enableTerminationProtection bool
	isFlexCluster               bool
	mdbVersion                  string
	filename                    string
	tag                         map[string]string
	fs                          afero.Fs
	store                       ClusterCreator
	autoScalingMode             string
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

const (
	createTemplate      = "Cluster '{{.Name}}' is being created.\n"
	createWatchTemplate = "Cluster '{{.Name}}' created successfully.\n"
)

var clusterObj *atlasClustersPinned.AdvancedClusterDescription
var flexCluster *atlasv2.FlexClusterDescription20241113
var clusterObjLatest *atlasv2.ClusterDescription20240805

func (opts *CreateOpts) Run() error {
	if opts.autoScalingMode == independentShardScalingFlag {
		return opts.RunDedicatedClusterLatest()
	}

	if opts.isFlexCluster {
		return opts.RunFlexCluster()
	}

	return opts.RunDedicatedCluster()
}

func (opts *CreateOpts) RunFlexCluster() error {
	flexClusterReq, err := opts.newFlexCluster()
	if err != nil {
		return err
	}

	flexCluster, err = opts.store.CreateFlexCluster(opts.ConfigProjectID(), flexClusterReq)

	apiError, ok := atlasv2.AsError(err)
	code := apiError.GetErrorCode()
	if ok {
		if apiError.GetErrorCode() == invalidAttributeErrorCode && strings.Contains(apiError.GetDetail(), regionName) {
			return cli.ErrNoRegionExistsTryCommand
		}
		if ok && code == duplicateClusterNameErrorCode {
			return cli.ErrNameExists
		}
	}

	return err
}

func (opts *CreateOpts) newFlexCluster() (*atlasv2.FlexClusterDescriptionCreate20241113, error) {
	cluster := new(atlasv2.FlexClusterDescriptionCreate20241113)
	if opts.filename != "" {
		if err := file.Load(opts.fs, opts.filename, cluster); err != nil {
			return nil, err
		}
	} else {
		cluster = opts.newFlexClusterDescriptionCreate20241113()
	}

	if opts.name != "" {
		cluster.Name = opts.name
	}

	return cluster, nil
}

func (opts *CreateOpts) RunDedicatedClusterLatest() error {
	cluster, err := opts.newClusterLatest()
	if err != nil {
		return err
	}

	clusterObjLatest, err = opts.store.CreateClusterLatest(cluster)
	apiError, ok := atlasv2.AsError(err)
	code := apiError.GetErrorCode()
	if ok {
		if apiError.GetErrorCode() == invalidAttributeErrorCode && strings.Contains(apiError.GetDetail(), regionName) {
			return cli.ErrNoRegionExistsTryCommand
		}
		if ok && code == duplicateClusterNameErrorCode {
			return cli.ErrNameExists
		}
	}

	return err
}

func (opts *CreateOpts) RunDedicatedCluster() error {
	cluster, err := opts.newCluster()
	if err != nil {
		return err
	}

	clusterObj, err = opts.store.CreateCluster(cluster)
	apiError, ok := atlasClustersPinned.AsError(err)
	code := apiError.GetErrorCode()
	if ok {
		if apiError.GetErrorCode() == invalidAttributeErrorCode && strings.Contains(apiError.GetDetail(), regionName) {
			return cli.ErrNoRegionExistsTryCommand
		}
		if ok && code == duplicateClusterNameErrorCode {
			return cli.ErrNameExists
		}
	}

	return err
}

func (opts *CreateOpts) PostRun() error {
	if opts.autoScalingMode == independentShardScalingFlag {
		return opts.PostRunDedicatedClusterLatest()
	}

	if opts.isFlexCluster {
		return opts.PostRunFlexCluster()
	}

	return opts.PostRunDedicatedCluster()
}

func (opts *CreateOpts) PostRunFlexCluster() error {
	if !opts.EnableWatch {
		return opts.Print(flexCluster)
	}
	opts.Template = createWatchTemplate

	watcher := watchers.NewWatcherWithDefaultWait(
		*watchers.ClusterCreated,
		watchers.NewAtlasFlexClusterStateDescriber(
			opts.store.(store.ClusterDescriber),
			opts.ConfigProjectID(),
			opts.name,
		),
		opts.GetDefaultWait(),
	)

	watcher.Timeout = time.Duration(opts.Timeout)
	if err := opts.WatchWatcher(watcher); err != nil {
		return err
	}

	return opts.Print(flexCluster)
}

func (opts *CreateOpts) PostRunDedicatedClusterLatest() error {
	if !opts.EnableWatch {
		return opts.Print(clusterObjLatest)
	}

	watcher := watchers.NewWatcherWithDefaultWait(
		*watchers.ClusterCreated,
		watchers.NewAtlasClusterStateDescriber(
			opts.store.(store.ClusterDescriber),
			opts.ConfigProjectID(),
			opts.name,
		),
		opts.GetDefaultWait(),
	)

	watcher.Timeout = time.Duration(opts.Timeout)
	if err := opts.WatchWatcher(watcher); err != nil {
		return err
	}

	return opts.Print(clusterObjLatest)
}

func (opts *CreateOpts) PostRunDedicatedCluster() error {
	if !opts.EnableWatch {
		return opts.Print(clusterObj)
	}

	opts.Template = createWatchTemplate

	watcher := watchers.NewWatcherWithDefaultWait(
		*watchers.ClusterCreated,
		watchers.NewAtlasClusterStateDescriber(
			opts.store.(store.ClusterDescriber),
			opts.ConfigProjectID(),
			opts.name,
		),
		opts.GetDefaultWait(),
	)

	watcher.Timeout = time.Duration(opts.Timeout)
	if err := opts.WatchWatcher(watcher); err != nil {
		return err
	}

	return opts.Print(clusterObj)
}

func (opts *CreateOpts) newCluster() (*atlasClustersPinned.AdvancedClusterDescription, error) {
	cluster := new(atlasClustersPinned.AdvancedClusterDescription)
	if opts.filename != "" {
		if err := file.Load(opts.fs, opts.filename, cluster); err != nil {
			return nil, err
		}
		removeReadOnlyAttributes(cluster)
	} else {
		opts.applyOptsAdvancedCluster(cluster)
	}

	if opts.name != "" {
		cluster.Name = &opts.name
	}

	cluster.GroupId = pointer.Get(opts.ConfigProjectID())
	return cluster, nil
}

func (opts *CreateOpts) newClusterLatest() (*atlasv2.ClusterDescription20240805, error) {
	cluster := new(atlasv2.ClusterDescription20240805)
	if opts.filename != "" {
		if err := file.Load(opts.fs, opts.filename, cluster); err != nil {
			return nil, err
		}

		removeReadOnlyAttributesLatest(cluster)
		cluster.GroupId = pointer.Get(opts.ConfigProjectID())
		if opts.name != "" {
			cluster.Name = &opts.name
		}

		return cluster, nil
	}

	opts.applyOptsClusterLatest(cluster)

	return cluster, nil
}

func (opts *CreateOpts) applyOptsAdvancedCluster(out *atlasClustersPinned.AdvancedClusterDescription) {
	replicationSpec := opts.newAdvanceReplicationSpec()
	if opts.backup {
		out.BackupEnabled = &opts.backup
		out.PitEnabled = &opts.backup
	}
	if opts.biConnector {
		out.BiConnector = &atlasClustersPinned.BiConnector{Enabled: &opts.biConnector}
	}
	out.TerminationProtectionEnabled = &opts.enableTerminationProtection
	out.ClusterType = &opts.clusterType

	if !opts.isTenant() {
		out.DiskSizeGB = &opts.diskSizeGB
		out.MongoDBMajorVersion = &opts.mdbVersion
	}

	out.ReplicationSpecs = &[]atlasClustersPinned.ReplicationSpec{replicationSpec}

	addTags(out, opts.tag)
}

func (opts *CreateOpts) applyOptsClusterLatest(out *atlasv2.ClusterDescription20240805) {
	out.GroupId = pointer.Get(opts.ConfigProjectID())
	out.ClusterType = &opts.clusterType
	out.TerminationProtectionEnabled = &opts.enableTerminationProtection
	out.ReplicationSpecs = opts.newAdvanceReplicationSpecsLatest()

	if opts.name != "" {
		out.Name = &opts.name
	}

	if opts.backup {
		out.BackupEnabled = &opts.backup
		out.PitEnabled = &opts.backup
	}

	if opts.biConnector {
		out.BiConnector = &atlasv2.BiConnector{Enabled: &opts.biConnector}
	}

	if len(opts.tag) > 0 {
		out.Tags = newResourceTags(opts.tag)
	}
}

func (opts *CreateOpts) newFlexClusterDescriptionCreate20241113() *atlasv2.FlexClusterDescriptionCreate20241113 {
	return &atlasv2.FlexClusterDescriptionCreate20241113{
		Name: opts.name,
		ProviderSettings: atlasv2.FlexProviderSettingsCreate20241113{
			BackingProviderName: opts.provider,
			ProviderName:        pointer.Get(opts.tier),
			RegionName:          opts.region,
		},
		TerminationProtectionEnabled: &opts.enableTerminationProtection,
		Tags:                         newResourceTags(opts.tag),
	}
}

func (opts *CreateOpts) isTenant() bool {
	return opts.tier == atlasM0 || opts.tier == atlasM2 || opts.tier == atlasM5
}

func (opts *CreateOpts) providerName() string {
	if opts.isTenant() {
		return tenant
	}
	return opts.provider
}

func (opts *CreateOpts) newAdvanceReplicationSpec() atlasClustersPinned.ReplicationSpec {
	return atlasClustersPinned.ReplicationSpec{
		NumShards:     &opts.shards,
		ZoneName:      pointer.Get(zoneName),
		RegionConfigs: &[]atlasClustersPinned.CloudRegionConfig{opts.newAdvancedRegionConfig()},
	}
}

func (opts *CreateOpts) newAdvanceReplicationSpecsLatest() *[]atlasv2.ReplicationSpec20240805 {
	replicationSpecs := make([]atlasv2.ReplicationSpec20240805, opts.shards)
	for i := range opts.shards {
		replicationSpecs[i] = atlasv2.ReplicationSpec20240805{
			ZoneName: pointer.Get(zoneName),
			RegionConfigs: &[]atlasv2.CloudRegionConfig20240805{
				opts.newAdvanceRegionConfigLatest(),
			},
		}
	}
	return &replicationSpecs
}

func (opts *CreateOpts) newAdvancedRegionConfig() atlasClustersPinned.CloudRegionConfig {
	providerName := opts.providerName()

	regionConfig := atlasClustersPinned.CloudRegionConfig{
		Priority:     pointer.Get(priority),
		RegionName:   &opts.region,
		ProviderName: &providerName,
		ElectableSpecs: &atlasClustersPinned.HardwareSpec{
			InstanceSize: &opts.tier,
		},
	}

	if providerName == tenant {
		regionConfig.BackingProviderName = &opts.provider
	} else {
		regionConfig.ElectableSpecs.NodeCount = &opts.members
	}

	readOnlySpec := &atlasClustersPinned.DedicatedHardwareSpec{
		InstanceSize: &opts.tier,
		NodeCount:    pointer.Get(readOnlyNode),
	}
	regionConfig.ReadOnlySpecs = readOnlySpec

	return regionConfig
}

func (opts *CreateOpts) newAdvanceRegionConfigLatest() atlasv2.CloudRegionConfig20240805 {
	providerName := opts.providerName()
	regionConfig := atlasv2.CloudRegionConfig20240805{
		ProviderName: pointer.Get(providerName),
		Priority:     pointer.Get(priority),
		RegionName:   pointer.Get(opts.region),
		ElectableSpecs: &atlasv2.HardwareSpec20240805{
			InstanceSize: pointer.Get(opts.tier),
		},
		ReadOnlySpecs: &atlasv2.DedicatedHardwareSpec20240805{
			InstanceSize: pointer.Get(opts.tier),
			NodeCount:    pointer.Get(readOnlyNode),
		},
	}

	if providerName == tenant {
		regionConfig.BackingProviderName = &opts.provider
	} else {
		regionConfig.ElectableSpecs.NodeCount = pointer.Get(opts.members)
	}

	return regionConfig
}

// newIsFlexCluster sets the opts.isFlexCluster that indicates if the cluster to create is
// a FlexCluster. When opts.filename is not provided, a FlexCluster has the opts.tier==FLEX.
// When opts.filename is provided, the function loads the file and check that the field replicationSpecs
// (available only for Dedicated Cluster) is present.
func (opts *CreateOpts) newIsFlexCluster() error {
	if opts.filename == "" {
		opts.isFlexCluster = opts.tier == atlasFlex
		return nil
	}

	var m map[string]any
	if err := file.Load(opts.fs, opts.filename, &m); err != nil {
		opts.isFlexCluster = false
		return fmt.Errorf("%w: %w", errFailedToLoadClusterFileMessage, err)
	}

	_, ok := m["replicationSpecs"]
	opts.isFlexCluster = !ok
	return nil
}

func (opts *CreateOpts) validateTier() error {
	opts.tier = strings.ToUpper(opts.tier)
	if opts.tier == atlasM2 || opts.tier == atlasM5 {
		_, _ = fmt.Fprintf(os.Stderr, deprecateMessageSharedTier, opts.tier)
	}
	return nil
}

func (opts *CreateOpts) validateAutoScalingMode() error {
	if opts.isFlexCluster && opts.autoScalingMode != clusterWideScalingFlag {
		return fmt.Errorf("flex is incompatible with %s auto scaling mode", opts.autoScalingMode)
	}

	if opts.autoScalingMode != clusterWideScalingFlag && opts.autoScalingMode != independentShardScalingFlag {
		return fmt.Errorf("invalid auto scaling mode: %s", opts.autoScalingMode)
	}

	if opts.isFlexCluster {
		return nil
	}

	// If the file is provided and it is not flex,
	// we need to check the format of the file
	if opts.filename != "" {
		// First try to load as a default dedicated cluster
		oldCluster := new(atlasClustersPinned.AdvancedClusterDescription)
		oldLoadErr := file.Load(opts.fs, opts.filename, oldCluster)
		if oldLoadErr == nil {
			opts.autoScalingMode = clusterWideScalingFlag
			return nil
		}

		// Then try to load as an ISS cluster
		cluster := new(atlasv2.ClusterDescription20240805)
		latestLoadErr := file.Load(opts.fs, opts.filename, cluster)
		if latestLoadErr == nil {
			opts.autoScalingMode = independentShardScalingFlag
			return nil
		}

		return fmt.Errorf("failed to load cluster file: %w", latestLoadErr)
	}

	return nil
}

// CreateBuilder builds a cobra.Command that can run as:
// create <name> --projectId projectId --provider AWS|GCP|AZURE --region regionName [--members N] [--tier M#] [--diskSizeGB N] [--backup] [--mdbVersion] [--tag key=value].
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{
		fs: afero.NewOsFs(),
	}

	cmd := &cobra.Command{
		Use:   "create [name]",
		Short: "Create a cluster for your project.",
		Long: `To get started quickly, specify a name for your cluster, a cloud provider, and a region to deploy a three-member replica set with the latest MongoDB server version.
For full control of your deployment, or to create multi-cloud clusters, provide a JSON configuration file with the --file flag.

Deprecation note: the M2 and M5 tiers are now deprecated; when selecting M2 or M5, a FLEX tier will be created instead. For the migration guide, visit: https://dochub.mongodb.org/core/flex-migration.\n

` + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Example: `  # Deploy a free cluster named myCluster for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas cluster create myCluster --projectId 5e2211c17a3e5a48f5497de3 --provider AWS --region US_EAST_1 --tier M0

  # Deploy a Flex cluster named myFlexCluster for the project with the ID 5e2211c17a3e5a48f5497de3 and tag "env=dev":
  atlas cluster create myFlexCluster --projectId 5e2211c17a3e5a48f5497de3 --provider AWS --region US_EAST_1 --tier FLEX --tag env=dev

  # Deploy a free cluster named myCluster for the project with the ID 5e2211c17a3e5a48f5497de3 and tag "env=dev":
  atlas cluster create myCluster --projectId 5e2211c17a3e5a48f5497de3 --provider AWS --region US_EAST_1 --tier M0 --tag env=dev

  # Deploy a three-member replica set named myRS in AWS for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas cluster create myRS --projectId 5e2211c17a3e5a48f5497de3 --provider AWS --region US_EAST_1 --members 3 --tier M10 --mdbVersion 5.0 --diskSizeGB 10

  # Deploy a three-member replica set named myRS in AZURE for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas cluster create myRS --projectId 5e2211c17a3e5a48f5497de3 --provider AZURE --region US_EAST_2 --members 3 --tier M10  --mdbVersion 5.0 --diskSizeGB 10
  
  # Deploy a three-member replica set named myRS in GCP for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas cluster create myRS --projectId 5e2211c17a3e5a48f5497de3 --provider GCP --region EASTERN_US --members 3 --tier M10  --mdbVersion 5.0 --diskSizeGB 10

  # Deploy a cluster or a multi-cloud cluster from a JSON configuration file named myfile.json for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas cluster create --projectId <projectId> --file myfile.json`,
		Args: require.MaximumNArgs(1),
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

			opts.tier = strings.ToUpper(opts.tier)
			opts.region = strings.ToUpper(opts.region)
			return opts.PreRunE(
				opts.validateTier,
				opts.newIsFlexCluster,
				opts.validateAutoScalingMode,
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
		PostRunE: func(_ *cobra.Command, _ []string) error {
			return opts.PostRun()
		},
		Annotations: map[string]string{
			"nameDesc": "Name of the cluster. The cluster name cannot be changed after the cluster is created. Cluster name can contain ASCII letters, numbers, and hyphens. You must specify the cluster name argument if you don't use the --file option.",
			"output":   createTemplate,
		},
	}

	currentMDBVersion, _ := cli.DefaultMongoDBMajorVersion()

	const (
		defaultMembersSize = 3
		defaultDiskSize    = 2
		defaultShardSize   = 1
	)
	cmd.Flags().StringVar(&opts.provider, flag.Provider, "", usage.CreateProvider)
	cmd.Flags().StringVarP(&opts.region, flag.Region, flag.RegionShort, "", usage.CreateRegion)
	cmd.Flags().IntVarP(&opts.members, flag.Members, flag.MembersShort, defaultMembersSize, usage.Members)
	cmd.Flags().StringVar(&opts.tier, flag.Tier, atlasFlex, usage.Tier)
	cmd.Flags().Float64Var(&opts.diskSizeGB, flag.DiskSizeGB, defaultDiskSize, usage.DiskSizeGB)
	cmd.Flags().StringVar(&opts.mdbVersion, flag.MDBVersion, currentMDBVersion, usage.MDBVersion)
	cmd.Flags().BoolVar(&opts.backup, flag.Backup, false, usage.Backup)
	cmd.Flags().BoolVar(&opts.biConnector, flag.BIConnector, false, usage.BIConnector)
	cmd.Flags().StringVarP(&opts.filename, flag.File, flag.FileShort, "", usage.ClusterFilename)
	cmd.Flags().StringVar(&opts.clusterType, flag.TypeFlag, replicaSet, usage.ClusterTypes)
	cmd.Flags().IntVarP(&opts.shards, flag.Shards, flag.ShardsShort, defaultShardSize, usage.Shards)
	cmd.Flags().BoolVar(&opts.enableTerminationProtection, flag.EnableTerminationProtection, false, usage.EnableTerminationProtection)
	cmd.Flags().StringToStringVar(&opts.tag, flag.Tag, nil, usage.Tag)
	cmd.Flags().StringVar(&opts.autoScalingMode, flag.AutoScalingMode, clusterWideScalingFlag, usage.AutoScalingMode)

	cmd.Flags().BoolVarP(&opts.EnableWatch, flag.EnableWatch, flag.EnableWatchShort, false, usage.EnableWatch)
	cmd.Flags().Int64Var(&opts.Timeout, flag.WatchTimeout, 0, usage.WatchTimeout)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagFilename(flag.File)

	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.Tier)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.Provider)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.Members)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.Region)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.DiskSizeGB)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.MDBVersion)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.BIConnector)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.TypeFlag)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.Shards)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.Tag)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.AutoScalingMode)

	_ = cmd.RegisterFlagCompletionFunc(flag.TypeFlag, func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return []string{"REPLICASET", "SHARDED", "GEOSHARDED"}, cobra.ShellCompDirectiveDefault
	})

	_ = cmd.RegisterFlagCompletionFunc(flag.Provider, func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return []string{"AWS", "AZURE", "GCP"}, cobra.ShellCompDirectiveDefault
	})

	_ = cmd.RegisterFlagCompletionFunc(flag.AutoScalingMode, func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return []string{clusterWideScalingFlag, independentShardScalingFlag}, cobra.ShellCompDirectiveDefault
	})

	autocomplete := &autoCompleteOpts{}
	_ = cmd.RegisterFlagCompletionFunc(flag.Tier, autocomplete.autocompleteTier())
	_ = cmd.RegisterFlagCompletionFunc(flag.Region, autocomplete.autocompleteRegion())

	return cmd
}

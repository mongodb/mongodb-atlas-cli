// Copyright 2022 MongoDB Inc
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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113004/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const (
	upgradeTemplate     = "Upgrading cluster '{{.Name}}'.\n"
	replicaSetNodeCount = 3
	replicaSetPriority  = 7
)

type UpgradeOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	name                         string
	mdbVersion                   string
	diskSizeGB                   float64
	tier                         string
	filename                     string
	enableTerminationProtection  bool
	disableTerminationProtection bool
	isFlexCluster                bool
	tag                          map[string]string
	fs                           afero.Fs
	store                        store.AtlasSharedClusterGetterUpgrader
}

func (opts *UpgradeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *UpgradeOpts) Run() error {
	if opts.isFlexCluster {
		return opts.RunFlexCluster()
	}

	return opts.RunSharedCluster()
}

func (opts *UpgradeOpts) RunFlexCluster() error {
	cluster, err := opts.atlasTenantClusterUpgradeRequest20240805()
	if err != nil {
		return err
	}

	r, err := opts.store.UpgradeFlexCluster(opts.ConfigProjectID(), cluster)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *UpgradeOpts) RunSharedCluster() error {
	cluster, err := opts.cluster()
	if err != nil {
		return err
	}
	if opts.filename == "" {
		opts.patchOpts(cluster)
	}

	r, err := opts.store.UpgradeCluster(opts.ConfigProjectID(), cluster)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *UpgradeOpts) atlasTenantClusterUpgradeRequest20240805() (*atlasv2.AtlasTenantClusterUpgradeRequest20240805, error) {
	var cluster *atlasv2.AtlasTenantClusterUpgradeRequest20240805
	if opts.filename != "" {
		err := file.Load(opts.fs, opts.filename, &cluster)
		if err != nil {
			return nil, err
		}

		cluster.Name = opts.name

		return cluster, nil
	}

	flexClusterDescription, err := opts.store.FlexCluster(opts.ConfigProjectID(), opts.name)
	if err != nil {
		return nil, err
	}

	return opts.newAtlasTenantClusterUpgradeRequestFromFlexClusterDescription(flexClusterDescription), nil
}

func (opts *UpgradeOpts) newAtlasTenantClusterUpgradeRequestFromFlexClusterDescription(flexCluster *atlasv2.FlexClusterDescription20241113) *atlasv2.AtlasTenantClusterUpgradeRequest20240805 {
	mdbVersion := flexCluster.MongoDBVersion
	if opts.mdbVersion != "" {
		mdbVersion = &opts.mdbVersion
	}

	terminationProtectionEnabled := flexCluster.GetTerminationProtectionEnabled()
	if opts.disableTerminationProtection {
		terminationProtectionEnabled = false
	}

	if opts.enableTerminationProtection {
		terminationProtectionEnabled = true
	}

	var tags []atlasv2.ResourceTag
	if flexCluster.Tags != nil {
		tags = flexCluster.GetTags()
	}

	if len(opts.tag) > 0 {
		newTags := newResourceTags(opts.tag)
		tags = append(tags, *newTags...)
	}

	backupEnabled := false
	if settings, ok := flexCluster.GetBackupSettingsOk(); ok {
		backupEnabled = settings.GetEnabled()
	}

	flexGroupID, _ := flexCluster.GetGroupIdOk()
	versionRelease, _ := flexCluster.GetVersionReleaseSystemOk()
	flexClusterType, _ := flexCluster.GetClusterTypeOk()

	return &atlasv2.AtlasTenantClusterUpgradeRequest20240805{
		BackupEnabled:                &backupEnabled,
		ClusterType:                  flexClusterType,
		GroupId:                      flexGroupID,
		MongoDBVersion:               mdbVersion,
		Name:                         flexCluster.GetName(),
		Tags:                         &tags,
		TerminationProtectionEnabled: &terminationProtectionEnabled,
		VersionReleaseSystem:         versionRelease,
		ReplicationSpecs:             opts.newReplicationSpecFromOpts(flexCluster),
	}
}

func (opts *UpgradeOpts) newReplicationSpecFromOpts(flexCluster *atlasv2.FlexClusterDescription20241113) *[]atlasv2.ReplicationSpec20240805 {
	if opts.tier == "" {
		return nil
	}

	diskSizeGb := 0.0
	backingProviderName := ""
	regionN := ""
	if settings, ok := flexCluster.GetProviderSettingsOk(); ok {
		diskSizeGb = settings.GetDiskSizeGB()
		backingProviderName = settings.GetBackingProviderName()
		regionN = settings.GetRegionName()
	}

	if opts.diskSizeGB != 0 {
		diskSizeGb = opts.diskSizeGB
	}

	replicaSetCount := replicaSetNodeCount
	priority := replicaSetPriority

	replicaSpec := atlasv2.ReplicationSpec20240805{
		RegionConfigs: &[]atlasv2.CloudRegionConfig20240805{
			{
				ElectableSpecs: &atlasv2.HardwareSpec20240805{
					InstanceSize: &opts.tier,
					DiskSizeGB:   &diskSizeGb,
					NodeCount:    &replicaSetCount,
				},
				ProviderName: &backingProviderName,
				RegionName:   &regionN,
				Priority:     &priority,
			},
		},
	}

	return &[]atlasv2.ReplicationSpec20240805{replicaSpec}
}

func (opts *UpgradeOpts) cluster() (*atlas.Cluster, error) {
	var cluster *atlas.Cluster
	if opts.filename != "" {
		err := file.Load(opts.fs, opts.filename, &cluster)
		if err != nil {
			return nil, err
		}
		cluster.Name = opts.name

		return cluster, nil
	}

	return opts.store.AtlasSharedCluster(opts.ConfigProjectID(), opts.name)
}

func (opts *UpgradeOpts) patchOpts(out *atlas.Cluster) {
	removeReadOnlyAttributesSharedCluster(out)
	if opts.mdbVersion != "" {
		out.MongoDBMajorVersion = opts.mdbVersion
	}
	if opts.diskSizeGB > 0 {
		out.DiskSizeGB = &opts.diskSizeGB
	}
	if out.ProviderSettings != nil {
		if opts.tier != "" {
			out.ProviderSettings.InstanceSizeName = opts.tier
		}
		if isTenant(out.ProviderSettings.InstanceSizeName) {
			out.BiConnector = nil
		} else {
			out.ProviderSettings.ProviderName = out.ProviderSettings.BackingProviderName
			out.ProviderSettings.BackingProviderName = ""
		}
	}
	out.TerminationProtectionEnabled = cli.ReturnValueForSetting(
		opts.enableTerminationProtection,
		opts.disableTerminationProtection,
	)

	var tags []*atlas.Tag
	if len(opts.tag) > 0 {
		tags = make([]*atlas.Tag, 0, len(opts.tag))
	}
	for k, v := range opts.tag {
		if k != "" && v != "" {
			tags = append(tags, &atlas.Tag{Key: k, Value: v})
		}
	}
	out.Tags = &tags
}

func isTenant(instanceSizeName string) bool {
	return instanceSizeName == atlasM0 ||
		instanceSizeName == atlasM2 ||
		instanceSizeName == atlasM5
}

// newIsFlexCluster sets the opts.isFlexCluster that indicates if the cluster to create is
// a FlexCluster. The function calls the AtlasSharedAPI to get the cluster, and it sets the opts.isFlexCluster = true
// in the event of a cannotUseFlexWithClusterApisErrorCode.
func (opts *UpgradeOpts) newIsFlexCluster() error {
	_, err := opts.store.AtlasSharedCluster(opts.ConfigProjectID(), opts.name)
	var errorResponse *atlas.ErrorResponse
	ok := errors.As(err, &errorResponse)
	if !ok {
		opts.isFlexCluster = false
		return err
	}

	if errorResponse.ErrorCode != cannotUseFlexWithClusterApisErrorCode {
		return err
	}

	opts.isFlexCluster = true
	return nil
}

// UpgradeBuilder builds a cobra.Command that can run as:
// atlas cluster(s) upgrade [clusterName] --projectId projectId [--tier M#] [--diskSizeGB N] [--mdbVersion] [--tag key=value].
func UpgradeBuilder() *cobra.Command {
	opts := UpgradeOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "upgrade [clusterName]",
		Short: "Upgrade a shared cluster's tier, disk size, and/or MongoDB version.",
		Long: `This command is unavailable for dedicated clusters.

` + fmt.Sprintf(usage.RequiredRole, "Project Cluster Manager"),
		Example: `  # Upgrade the tier, disk size, and MongoDB version for the shared cluster named myCluster in the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas cluster upgrade myCluster --projectId 5e2211c17a3e5a48f5497de3 --tier M50 --diskSizeGB 20 --mdbVersion 7.0 --tag env=dev`,
		Args: require.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				opts.name = args[0]
			}
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), upgradeTemplate),
				opts.newIsFlexCluster,
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the cluster to upgrade.",
			"output":          upgradeTemplate,
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

	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.Tier)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.DiskSizeGB)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.MDBVersion)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.EnableTerminationProtection)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.DisableTerminationProtection)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.Tag)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagFilename(flag.File)

	cmd.MarkFlagsOneRequired(flag.File, flag.Tier, flag.DiskSizeGB, flag.MDBVersion, flag.EnableTerminationProtection, flag.DisableTerminationProtection, flag.Tag)

	return cmd
}

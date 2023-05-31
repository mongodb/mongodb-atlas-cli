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
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const (
	replicaSet        = "REPLICASET"
	tenant            = "TENANT"
	atlasM0           = "M0"
	atlasM2           = "M2"
	atlasM5           = "M5"
	zoneName          = "Zone 1"
	awsProviderName   = "AWS"
	gcpProviderName   = "GCP"
	azureProviderName = "Azure"
)

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
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
	mdbVersion                  string
	filename                    string
	fs                          afero.Fs
	store                       store.ClusterCreator
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var createTmpl = "Deploying cluster '{{.Name}}'.\n"

func (opts *CreateOpts) Run() error {
	cluster, err := opts.newCluster()
	if err != nil {
		return err
	}

	r, err := opts.store.CreateCluster(cluster)
	var target *atlas.ErrorResponse
	if errors.As(err, &target) && target.ErrorCode == "INVALID_ATTRIBUTE" && strings.Contains(target.Detail, "regionName") {
		return cli.ErrNoRegionExistsTryCommand
	} else if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newCluster() (*atlasv2.ClusterDescriptionV15, error) {
	cluster := new(atlasv2.ClusterDescriptionV15)
	if opts.filename != "" {
		if err := file.Load(opts.fs, opts.filename, cluster); err != nil {
			return nil, err
		}
		RemoveReadOnlyAttributes(cluster)
	} else {
		opts.applyOpts(cluster)
	}

	if opts.name != "" {
		cluster.Name = &opts.name
	}

	AddLabel(cluster, NewCLILabel())

	cluster.GroupId = pointer.Get(opts.ConfigProjectID())
	return cluster, nil
}

func (opts *CreateOpts) applyOpts(out *atlasv2.ClusterDescriptionV15) {
	replicationSpec := opts.newAdvanceReplicationSpec()
	if opts.backup {
		out.BackupEnabled = &opts.backup
		out.PitEnabled = &opts.backup
	}
	if opts.biConnector {
		out.BiConnector = &atlasv2.BiConnector{Enabled: &opts.biConnector}
	}
	out.TerminationProtectionEnabled = &opts.enableTerminationProtection
	out.ClusterType = &opts.clusterType

	if !opts.isTenant() {
		out.DiskSizeGB = &opts.diskSizeGB
		out.MongoDBMajorVersion = &opts.mdbVersion
	}

	out.ReplicationSpecs = []atlasv2.ReplicationSpec{replicationSpec}
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

func (opts *CreateOpts) newAdvanceReplicationSpec() atlasv2.ReplicationSpec {
	return atlasv2.ReplicationSpec{
		NumShards:     &opts.shards,
		ZoneName:      pointer.Get(zoneName),
		RegionConfigs: []atlasv2.RegionConfig{opts.newAdvancedRegionConfig()},
	}
}

func (opts *CreateOpts) newAdvancedRegionConfig() atlasv2.RegionConfig {
	priority := 7
	readOnlyNode := 0
	providerName := opts.providerName()
	readOnlySpec := &atlasv2.DedicatedHardwareSpec{
		InstanceSize: &opts.tier,
		NodeCount:    &readOnlyNode,
	}
	regionConfig := atlasv2.RegionConfig{}

	switch providerName {
	case tenant:
		regionConfig.TenantRegionConfig = &atlasv2.TenantRegionConfig{
			ProviderName: &providerName,
			Priority:     &priority,
			RegionName:   &opts.region,
			ElectableSpecs: &atlasv2.HardwareSpec{
				InstanceSize: &opts.tier,
			},
			BackingProviderName: &opts.provider,
		}
	case awsProviderName:
		regionConfig.AWSRegionConfig = &atlasv2.AWSRegionConfig{
			ProviderName: &providerName,
			Priority:     &priority,
			RegionName:   &opts.region,
			ElectableSpecs: &atlasv2.HardwareSpec{
				InstanceSize: &opts.tier,
				NodeCount:    &opts.members,
			},
			ReadOnlySpecs: readOnlySpec,
		}
	case azureProviderName:
		regionConfig.AzureRegionConfig = &atlasv2.AzureRegionConfig{
			ProviderName: &providerName,
			Priority:     &priority,
			RegionName:   &opts.region,
			ElectableSpecs: &atlasv2.HardwareSpec{
				InstanceSize: &opts.tier,
				NodeCount:    &opts.members,
			},
			ReadOnlySpecs: readOnlySpec,
		}
	case gcpProviderName:
		regionConfig.GCPRegionConfig = &atlasv2.GCPRegionConfig{
			ProviderName: &providerName,
			Priority:     &priority,
			RegionName:   &opts.region,
			ElectableSpecs: &atlasv2.HardwareSpec{
				InstanceSize: &opts.tier,
				NodeCount:    &opts.members,
			},
			ReadOnlySpecs: readOnlySpec,
		}
	}

	return regionConfig
}

// CreateBuilder builds a cobra.Command that can run as:
// create <name> --projectId projectId --provider AWS|GCP|AZURE --region regionName [--members N] [--tier M#] [--diskSizeGB N] [--backup] [--mdbVersion].
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{
		fs: afero.NewOsFs(),
	}

	cmd := &cobra.Command{
		Use:   "create [name]",
		Short: "Create a cluster for your project.",
		Long: `To get started quickly, specify a name for your cluster, a cloud provider, and a region to deploy a three-member replica set with the latest MongoDB server version.
For full control of your deployment, or to create multi-cloud clusters, provide a JSON configuration file with the --file flag.

` + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Example: fmt.Sprintf(`  # Deploy a free cluster named myCluster for the project with the ID 5e2211c17a3e5a48f5497de3:
  %[1]s cluster create myCluster --projectId 5e2211c17a3e5a48f5497de3 --provider AWS --region US_EAST_1 --tier M0

  # Deploy a three-member replica set named myRS in AWS for the project with the ID 5e2211c17a3e5a48f5497de3:
  %[1]s cluster create myRS --projectId 5e2211c17a3e5a48f5497de3 --provider AWS --region US_EAST_1 --members 3 --tier M10 --mdbVersion 5.0 --diskSizeGB 10

  # Deploy a three-member replica set named myRS in AZURE for the project with the ID 5e2211c17a3e5a48f5497de3:
  %[1]s cluster create myRS --projectId 5e2211c17a3e5a48f5497de3 --provider AZURE --region US_EAST_2 --members 3 --tier M10  --mdbVersion 5.0 --diskSizeGB 10
  
  # Deploy a three-member replica set named myRS in GCP for the project with the ID 5e2211c17a3e5a48f5497de3:
  %[1]s cluster create myRS --projectId 5e2211c17a3e5a48f5497de3 --provider GCP --region EASTERN_US --members 3 --tier M10  --mdbVersion 5.0 --diskSizeGB 10

  # Deploy a cluster or a multi-cloud cluster from a JSON configuration file named myfile.json for the project with the ID 5e2211c17a3e5a48f5497de3:
  %[1]s cluster create --projectId <projectId> --file myfile.json`, cli.ExampleAtlasEntryPoint()),
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
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTmpl),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
		Annotations: map[string]string{
			"nameDesc": "Name of the cluster. The cluster name cannot be changed after the cluster is created. Cluster name can contain ASCII letters, numbers, and hyphens.",
			"output":   createTmpl,
		},
	}

	currentMDBVersion, _ := cli.DefaultMongoDBMajorVersion()

	const (
		defaultMembersSize = 3
		defaultDiskSize    = 2
		defaultShardSize   = 1
	)
	cmd.Flags().StringVar(&opts.provider, flag.Provider, "", usage.Provider)
	cmd.Flags().StringVarP(&opts.region, flag.Region, flag.RegionShort, "", usage.Region)
	cmd.Flags().IntVarP(&opts.members, flag.Members, flag.MembersShort, defaultMembersSize, usage.Members)
	cmd.Flags().StringVar(&opts.tier, flag.Tier, atlasM2, usage.Tier)
	cmd.Flags().Float64Var(&opts.diskSizeGB, flag.DiskSizeGB, defaultDiskSize, usage.DiskSizeGB)
	cmd.Flags().StringVar(&opts.mdbVersion, flag.MDBVersion, currentMDBVersion, usage.MDBVersion)
	cmd.Flags().BoolVar(&opts.backup, flag.Backup, false, usage.Backup)
	cmd.Flags().BoolVar(&opts.biConnector, flag.BIConnector, false, usage.BIConnector)
	cmd.Flags().StringVarP(&opts.filename, flag.File, flag.FileShort, "", usage.ClusterFilename)
	cmd.Flags().StringVar(&opts.clusterType, flag.TypeFlag, replicaSet, usage.ClusterTypes)
	cmd.Flags().IntVarP(&opts.shards, flag.Shards, flag.ShardsShort, defaultShardSize, usage.Shards)
	cmd.Flags().BoolVar(&opts.enableTerminationProtection, flag.EnableTerminationProtection, false, usage.EnableTerminationProtection)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

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

	_ = cmd.RegisterFlagCompletionFunc(flag.TypeFlag, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"REPLICASET", "SHARDED", "GEOSHARDED"}, cobra.ShellCompDirectiveDefault
	})

	_ = cmd.RegisterFlagCompletionFunc(flag.Provider, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"AWS", "AZURE", "GCP"}, cobra.ShellCompDirectiveDefault
	})

	autocomplete := &autoCompleteOpts{}
	_ = cmd.RegisterFlagCompletionFunc(flag.Tier, autocomplete.autocompleteTier())
	_ = cmd.RegisterFlagCompletionFunc(flag.Region, autocomplete.autocompleteRegion())

	return cmd
}

package cli

import (
	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/flags"
	"github.com/10gen/mcli/internal/store"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
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
	nodes        int64
	diskSize     float64
	backup       bool
	mdbVersion   string
	store        store.ClusterCreator
}

func (opts *atlasClustersCreateOpts) init() error {
	if err := opts.loadConfig(); err != nil {
		return err
	}

	if opts.ProjectID() == "" {
		return errMissingProjectID
	}

	s, err := store.New(opts.Config)

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

	return prettyJSON(result)
}

func (opts *atlasClustersCreateOpts) newCluster() *atlas.Cluster {
	replicationSpec := opts.newReplicationSpec()
	providerSettings := opts.newProviderSettings()

	cluster := &atlas.Cluster{
		BackupEnabled:       &opts.backup,
		ClusterType:         replicaSet,
		DiskSizeGB:          &opts.diskSize,
		GroupID:             opts.ProjectID(),
		MongoDBMajorVersion: opts.mdbVersion,
		Name:                opts.name,
		ProviderSettings:    providerSettings,
		ReplicationSpecs:    []atlas.ReplicationSpec{*replicationSpec},
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

func (opts *atlasClustersCreateOpts) newReplicationSpec() *atlas.ReplicationSpec {
	var (
		readOnlyNodes int64 = 0
		NumShards     int64 = 1
		Priority      int64 = 7
	)
	replicationSpec := &atlas.ReplicationSpec{
		NumShards: &NumShards,
		ZoneName:  zoneName,
		RegionsConfig: map[string]atlas.RegionsConfig{
			opts.region: {
				ReadOnlyNodes:  &readOnlyNodes,
				ElectableNodes: &opts.nodes,
				Priority:       &Priority,
			},
		},
	}
	return replicationSpec
}

// mcli atlas cluster(s) create name --projectId projectId --provider AWS|GCP|AZURE --region regionName [--nodes N] [--instanceSize M#] [--diskSize N] [--backup] [--mdbVersion]
func AtlasClustersCreateBuilder() *cobra.Command {
	opts := &atlasClustersCreateOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:   "create [name]",
		Short: "Command to create a cluster with Atlas",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", "Project ID")
	cmd.Flags().StringVar(&opts.provider, flags.Provider, "", "Provider name")
	cmd.Flags().StringVar(&opts.region, flags.Region, "", "Provider region name")
	cmd.Flags().Int64Var(&opts.nodes, flags.Nodes, 3, "Number of nodes")
	cmd.Flags().StringVar(&opts.instanceSize, flags.InstanceSize, atlasM2, "Instance size")
	cmd.Flags().Float64Var(&opts.diskSize, flags.DiskSize, 2, "Storage size")
	cmd.Flags().StringVar(&opts.mdbVersion, flags.MDBVersion, currentMDBVersion, "mongoDB major version")
	cmd.Flags().BoolVar(&opts.backup, flags.Backup, false, "Backup")

	cmd.Flags().StringVar(&opts.profile, flags.Profile, config.DefaultProfile, "Profile")

	_ = cmd.MarkFlagRequired(flags.Provider)
	_ = cmd.MarkFlagRequired(flags.Region)

	return cmd
}

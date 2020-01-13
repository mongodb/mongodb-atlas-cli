// Copyright (C) 2020 - present MongoDB, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the Server Side Public License, version 1,
// as published by MongoDB, Inc.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// Server Side Public License for more details.
//
// You should have received a copy of the Server Side Public License
// along with this program. If not, see
// http://www.mongodb.com/licensing/server-side-public-license
//
// As a special exception, the copyright holders give permission to link the
// code of portions of this program with the OpenSSL library under certain
// conditions as described in each individual source file and distribute
// linked combinations including the program with the OpenSSL library. You
// must comply with the Server Side Public License in all respects for
// all of the code used other than as permitted herein. If you modify file(s)
// with this exception, you may extend this exception to your version of the
// file(s), but you are not obligated to do so. If you do not wish to do so,
// delete this exception statement from your version. If you delete this
// exception statement from all source files in the program, then also delete
// it in the license file.

package cli

import (
	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/flags"
	"github.com/10gen/mcli/internal/store"
	"github.com/10gen/mcli/internal/usage"
	"github.com/10gen/mcli/internal/utils"
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

	return utils.PrettyJSON(result)
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

// mcli atlas cluster(s) create name --projectId projectId --provider AWS|GCP|AZURE --region regionName [--members N] [--instanceSize M#] [--diskSize N] [--backup] [--mdbVersion]
func AtlasClustersCreateBuilder() *cobra.Command {
	opts := &atlasClustersCreateOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:     "create [name]",
		Short:   "Create a MongoDB cluster in Atlas.",
		Example: `  mcli atlas cluster create myCluster --projectId=1 --region US_EAST_1 --nodes 3 --instanceSize M2 --provider AWS --profile qa --mdbVersion 4.2 --diskSize 2`,
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
	cmd.Flags().StringVarP(&opts.region, flags.Region, "r", "", usage.Region)
	cmd.Flags().Int64VarP(&opts.nodes, flags.Members, "m", 3, usage.Members)
	cmd.Flags().StringVar(&opts.instanceSize, flags.InstanceSize, atlasM2, usage.InstanceSize)
	cmd.Flags().Float64Var(&opts.diskSize, flags.DiskSize, 2, usage.DiskSize)
	cmd.Flags().StringVar(&opts.mdbVersion, flags.MDBVersion, currentMDBVersion, usage.MDBVersion)
	cmd.Flags().BoolVar(&opts.backup, flags.Backup, false, usage.Backup)

	cmd.Flags().StringVarP(&opts.profile, flags.Profile, "p", config.DefaultProfile, usage.Profile)

	_ = cmd.MarkFlagRequired(flags.Provider)
	_ = cmd.MarkFlagRequired(flags.Region)

	return cmd
}

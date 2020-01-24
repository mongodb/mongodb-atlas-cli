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
	"github.com/10gen/mcli/internal/json"
	"github.com/10gen/mcli/internal/store"
	"github.com/10gen/mcli/internal/usage"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/spf13/cobra"
)

type atlasClustersUpdateOpts struct {
	*globalOpts
	name         string
	instanceSize string
	nodes        int64
	diskSizeGB   float64
	mdbVersion   string
	store        store.ClusterStore
}

func (opts *atlasClustersUpdateOpts) init() error {
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

func (opts *atlasClustersUpdateOpts) Run() error {
	current, err := opts.store.Cluster(opts.projectID, opts.name)

	if err != nil {
		return err
	}

	opts.update(current)

	result, err := opts.store.UpdateCluster(current)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func (opts *atlasClustersUpdateOpts) update(out *atlas.Cluster) {
	// There can only be one
	if out.ReplicationSpecs != nil {
		out.ReplicationSpec = nil
	}
	// This can't be sent
	out.MongoURI = ""
	out.MongoURIWithOptions = ""
	out.MongoURIUpdated = ""
	out.StateName = ""
	out.MongoDBVersion = ""

	if opts.mdbVersion != "" {
		out.MongoDBMajorVersion = opts.mdbVersion
	}

	if opts.diskSizeGB > 0 {
		out.DiskSizeGB = &opts.diskSizeGB
	}

	if opts.instanceSize != "" {
		out.ProviderSettings.InstanceSizeName = opts.instanceSize
	}
}

// mcli atlas cluster(s) update name --projectId projectId [--instanceSize M#] [--diskSizeGB N] [--mdbVersion]
func AtlasClustersUpdateBuilder() *cobra.Command {
	opts := &atlasClustersUpdateOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:     "update [name]",
		Short:   "Update a MongoDB cluster in Atlas.",
		Example: `  mcli atlas cluster update myCluster --projectId=1 --instanceSize M2 --mdbVersion 4.2 --diskSizeGB 2`,
		Args:    cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().Int64VarP(&opts.nodes, flags.Members, flags.MembersShort, 0, usage.Members)
	cmd.Flags().StringVar(&opts.instanceSize, flags.InstanceSize, "", usage.InstanceSize)
	cmd.Flags().Float64Var(&opts.diskSizeGB, flags.DiskSizeGB, 0, usage.DiskSizeGB)
	cmd.Flags().StringVar(&opts.mdbVersion, flags.MDBVersion, "", usage.MDBVersion)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.profile, flags.Profile, flags.ProfileShort, config.DefaultProfile, usage.Profile)

	return cmd
}

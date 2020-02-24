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

package cli

import (
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mcli/internal/file"
	"github.com/mongodb/mcli/internal/flags"
	"github.com/mongodb/mcli/internal/json"
	"github.com/mongodb/mcli/internal/store"
	"github.com/mongodb/mcli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

type atlasClustersUpdateOpts struct {
	*globalOpts
	name         string
	instanceSize string
	diskSizeGB   float64
	mdbVersion   string
	filename     string
	fs           afero.Fs
	store        store.ClusterStore
}

func (opts *atlasClustersUpdateOpts) init() error {
	if opts.ProjectID() == "" {
		return errMissingProjectID
	}
	var err error
	opts.store, err = store.New()
	return err
}

func (opts *atlasClustersUpdateOpts) Run() error {
	cluster, err := opts.cluster()
	if err != nil {
		return err
	}
	if opts.filename == "" {
		opts.patchOpts(cluster)
	} else {
		cluster.GroupID = opts.ProjectID()
	}

	result, err := opts.store.UpdateCluster(cluster)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func (opts *atlasClustersUpdateOpts) cluster() (*atlas.Cluster, error) {
	var cluster *atlas.Cluster
	var err error
	if opts.filename != "" {
		cluster = new(atlas.Cluster)
		err = file.Load(opts.fs, opts.filename, cluster)
	} else {
		cluster, err = opts.store.Cluster(opts.projectID, opts.name)
	}
	return cluster, err
}

func (opts *atlasClustersUpdateOpts) patchOpts(out *atlas.Cluster) {
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
		fs:         afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:     "update [name]",
		Short:   "Update a MongoDB cluster in Atlas.",
		Example: `  mcli atlas cluster update myCluster --projectId=1 --instanceSize M2 --mdbVersion 4.2 --diskSizeGB 2`,
		Args:    cobra.MaximumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if opts.filename == "" {
				if len(args) == 0 {
					return errMissingClusterName
				}
				opts.name = args[0]
			}
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.instanceSize, flags.InstanceSize, "", usage.InstanceSize)
	cmd.Flags().Float64Var(&opts.diskSizeGB, flags.DiskSizeGB, 0, usage.DiskSizeGB)
	cmd.Flags().StringVar(&opts.mdbVersion, flags.MDBVersion, "", usage.MDBVersion)
	cmd.Flags().StringVarP(&opts.filename, flags.File, flags.FileShort, "", usage.Filename)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}

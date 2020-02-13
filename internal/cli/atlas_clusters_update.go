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
	"github.com/mongodb/mcli/internal/flags"
	"github.com/mongodb/mcli/internal/json"
	"github.com/mongodb/mcli/internal/store"
	"github.com/mongodb/mcli/internal/usage"
	"github.com/spf13/cobra"
)

type atlasClustersUpdateOpts struct {
	*globalOpts
	name         string
	instanceSize string
	diskSizeGB   float64
	mdbVersion   string
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

	cmd.Flags().StringVar(&opts.instanceSize, flags.InstanceSize, "", usage.InstanceSize)
	cmd.Flags().Float64Var(&opts.diskSizeGB, flags.DiskSizeGB, 0, usage.DiskSizeGB)
	cmd.Flags().StringVar(&opts.mdbVersion, flags.MDBVersion, "", usage.MDBVersion)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}

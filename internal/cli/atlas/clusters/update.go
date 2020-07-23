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
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/file"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/output"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type UpdateOpts struct {
	cli.GlobalOpts
	name       string
	tier       string
	diskSizeGB float64
	mdbVersion string
	filename   string
	fs         afero.Fs
	store      store.ClusterStore
}

func (opts *UpdateOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *UpdateOpts) Run() error {
	cluster, err := opts.cluster()
	if err != nil {
		return err
	}
	if opts.filename == "" {
		opts.patchOpts(cluster)
	}

	r, err := opts.store.UpdateCluster(opts.ConfigProjectID(), opts.name, cluster)
	if err != nil {
		return err
	}

	return output.Print(config.Default(), "", r)
}

func (opts *UpdateOpts) cluster() (*atlas.Cluster, error) {
	var cluster *atlas.Cluster
	var err error
	if opts.filename != "" {
		cluster = new(atlas.Cluster)
		err = file.Load(opts.fs, opts.filename, cluster)
		if opts.name == "" {
			opts.name = cluster.Name
		}
	} else {
		cluster, err = opts.store.Cluster(opts.ProjectID, opts.name)
	}
	return cluster, err
}

func (opts *UpdateOpts) patchOpts(out *atlas.Cluster) {
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
	if opts.tier != "" {
		out.ProviderSettings.InstanceSizeName = opts.tier
	}

	updateLabels(out)
}

// mongocli atlas cluster(s) update [name] --projectId projectId [--tier M#] [--diskSizeGB N] [--mdbVersion]
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "update [name]",
		Short: description.UpdateCluster,
		Example: `
  Update tier for a cluster
  $ mongocli atlas cluster update <clusterName> --projectId <projectId> --tier M50

  Update disk size for a cluster
  $ mongocli atlas cluster update <clusterName> --projectId <projectId> --diskSizeGB 20

  Update MongoDB version for a cluster
  $ mongocli atlas cluster update <clusterName> --projectId <projectId> --mdbVersion 4.2`,
		Args: cobra.MaximumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				opts.name = args[0]
			}
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.tier, flag.Tier, "", usage.Tier)
	cmd.Flags().Float64Var(&opts.diskSizeGB, flag.DiskSizeGB, 0, usage.DiskSizeGB)
	cmd.Flags().StringVar(&opts.mdbVersion, flag.MDBVersion, "", usage.MDBVersion)
	cmd.Flags().StringVarP(&opts.filename, flag.File, flag.FileShort, "", usage.Filename)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagFilename(flag.File)

	return cmd
}

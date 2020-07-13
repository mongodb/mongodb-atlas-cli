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

package onlinearchive

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type UpdateOpts struct {
	cli.GlobalOpts
	id           string
	clusterName  string
	archiveAfter float64
	store        store.OnlineArchiveUpdater
}

func (opts *UpdateOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *UpdateOpts) Run() error {
	archive := opts.newOnlineArchive()
	result, err := opts.store.UpdateOnlineArchive(opts.ConfigProjectID(), opts.clusterName, archive)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func (opts *UpdateOpts) newOnlineArchive() *atlas.OnlineArchive {
	archive := &atlas.OnlineArchive{
		ID: opts.id,
		Criteria: &atlas.OnlineArchiveCriteria{
			ExpireAfterDays: opts.archiveAfter,
		},
	}
	return archive
}

// mongocli atlas cluster(s) start <name> [--projectId projectId]
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	cmd := &cobra.Command{
		Use:   "update <ID>",
		Short: description.UpdateOnlineArchive,
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().Float64Var(&opts.archiveAfter, flag.ArchiveAfter, 0, usage.ArchiveAfter)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flag.ClusterName)
	_ = cmd.MarkFlagRequired(flag.ArchiveAfter)

	return cmd
}

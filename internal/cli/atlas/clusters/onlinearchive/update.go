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
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	id           string
	clusterName  string
	archiveAfter float64
	store        store.OnlineArchiveUpdater
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var updateTemplate = "Online archive '{{.ID}}' updated.\n"

func (opts *UpdateOpts) Run() error {
	archive := opts.newOnlineArchive()
	r, err := opts.store.UpdateOnlineArchive(opts.ConfigProjectID(), opts.clusterName, archive)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *UpdateOpts) newOnlineArchive() *atlas.OnlineArchive {
	archive := &atlas.OnlineArchive{
		ID: opts.id,
		Criteria: &atlas.OnlineArchiveCriteria{
			ExpireAfterDays: &opts.archiveAfter,
		},
	}
	return archive
}

// mongocli atlas cluster(s) onlineArchive(s) start <archiveId> [--clusterName name][--archiveAfter N] [--projectId projectId].
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	cmd := &cobra.Command{
		Use:   "update <archiveId>",
		Short: "Modify the archiving rule for the specified online archive for a cluster.",
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"archiveIdDesc": "Unique identifier of the online archive to update.",
		},
		Example: fmt.Sprintf(`  # Update the archiving rule to archive after 5 days for the online archive with the ID 5f189832e26ec075e10c32d3 for the cluster named myCluster:
  %s clusters onlineArchives update 5f189832e26ec075e10c32d3 --clusterName --archiveAfter 5 myCluster --output json`, cli.ExampleAtlasEntryPoint()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().Float64Var(&opts.archiveAfter, flag.ArchiveAfter, 0, usage.ArchiveAfter)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.ClusterName)
	_ = cmd.MarkFlagRequired(flag.ArchiveAfter)

	return cmd
}

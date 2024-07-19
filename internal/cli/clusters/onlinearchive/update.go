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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	id              string
	clusterName     string
	archiveAfter    int
	expireAfterDays int
	store           store.OnlineArchiveUpdater
	filename        string
	fs              afero.Fs
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var updateTemplate = "Online archive '{{.Id}}' updated.\n"

func (opts *UpdateOpts) Run() error {
	archive, err := opts.newOnlineArchive()
	if err != nil {
		return err
	}

	r, err := opts.store.UpdateOnlineArchive(opts.ConfigProjectID(), opts.clusterName, archive)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *UpdateOpts) newOnlineArchive() (*atlasv2.BackupOnlineArchive, error) {
	if opts.filename != "" {
		var archive *atlasv2.BackupOnlineArchive
		if err := file.Load(opts.fs, opts.filename, &archive); err != nil {
			return nil, err
		}
		return archive, nil
	}

	archive := &atlasv2.BackupOnlineArchive{
		Id: &opts.id,
		Criteria: &atlasv2.Criteria{
			ExpireAfterDays: pointer.Get(opts.archiveAfter),
		},
	}

	if opts.expireAfterDays > 0 {
		archive.DataExpirationRule = &atlasv2.DataExpirationRule{
			ExpireAfterDays: &opts.expireAfterDays,
		}
	}

	return archive, nil
}

// atlas cluster(s) onlineArchive(s) start <archiveId> [--clusterName name][--archiveAfter N] [--projectId projectId].
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "update <archiveId>",
		Short: "Modify the archiving rule for the specified online archive for a cluster.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Data Access Admin"),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"archiveIdDesc": "Unique identifier of the online archive to update.",
			"output":        updateTemplate,
		},
		Example: `  # Update the archiving rule to archive after 5 days for the online archive with the ID 5f189832e26ec075e10c32d3 for the cluster named myCluster:
  atlas clusters onlineArchives update 5f189832e26ec075e10c32d3 --clusterName --archiveAfter 5 myCluster --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().IntVar(&opts.archiveAfter, flag.ArchiveAfter, 0, usage.ArchiveAfter)
	cmd.Flags().IntVar(&opts.expireAfterDays, flag.ExpireAfterDays, 0, usage.ExpireAfterDays)
	cmd.Flags().StringVar(&opts.filename, flag.File, "", usage.OnlineArchiveFilename)
	_ = cmd.MarkFlagFilename(flag.File)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.ClusterName)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.ArchiveAfter)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.ExpireAfterDays)

	cmd.MarkFlagsOneRequired(flag.File, flag.ArchiveAfter)

	return cmd
}

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

package snapshots

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/commonerrors"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	cli.GlobalOpts
	*cli.DeleteOpts
	clusterName string
	store       store.SnapshotsDeleter
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return commonerrors.Check(opts.Delete(opts.store.DeleteSnapshot, opts.ConfigProjectID(), opts.clusterName))
}

// atlas snapshot(s) delete <snapshotId> --force --clusterName [--projectId projectId].
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Snapshot '%s' deleted\n", "Snapshot not deleted"),
	}
	cmd := &cobra.Command{
		Use:     "delete <snapshotId>",
		Aliases: []string{"rm"},
		Short:   "Remove the specified backup snapshot.",
		Long:    fmt.Sprintf("%s\n%s", fmt.Sprintf(usage.RequiredRole, "Project Owner"), "Atlas supports this command only for M10+ clusters."),
		Args:    require.ExactArgs(1),
		Example: `  # Remove the backup snapshot with the ID 5f4007f327a3bd7b6f4103c5 for the cluster named myDemo:
  atlas backups snapshots delete 5f4007f327a3bd7b6f4103c5 --clusterName myDemo`,
		Annotations: map[string]string{
			"snapshotIdDesc": "Unique identifier of the snapshot you want to delete.",
			"output":         opts.SuccessMessage(),
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(opts.ValidateProjectID, opts.initStore(cmd.Context()))
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.Entry = args[0]
			if err := opts.Prompt(); err != nil {
				return err
			}
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flag.ClusterName)

	return cmd
}

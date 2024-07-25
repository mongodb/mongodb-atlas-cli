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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

type WatchOpts struct {
	cli.GlobalOpts
	cli.WatchOpts
	id          string
	clusterName string
	store       store.SnapshotsDescriber
}

var watchTemplate = "\nSnapshot changes completed.\n"

func (opts *WatchOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *WatchOpts) watcher() (any, bool, error) {
	result, err := opts.store.Snapshot(opts.ConfigProjectID(), opts.clusterName, opts.id)
	if err != nil {
		return nil, false, err
	}
	return nil, result.GetStatus() == "completed" || result.GetStatus() == "failed", nil
}

func (opts *WatchOpts) Run() error {
	if _, err := opts.Watch(opts.watcher); err != nil {
		return err
	}

	return opts.Print(nil)
}

// atlas snapshot(s) watch <snapshotId> --clusterName clusterName [--projectId projectId].
func WatchBuilder() *cobra.Command {
	opts := &WatchOpts{}
	cmd := &cobra.Command{
		Use:   "watch <snapshotId>",
		Short: "Watch the specified snapshot in your project until it becomes available.",
		Long: `This command checks the snapshot's status periodically until it reaches a completed or failed status. 
Once the snapshot reaches the expected status, the command prints "Snapshot changes completed."
If you run the command in the terminal, it blocks the terminal session until the resource status completes or fails.
You can interrupt the command's polling at any time with CTRL-C.

` + fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Example: `  # Watch the backup snapshot with the ID 5f4007f327a3bd7b6f4103c5 in the cluster named myDemo until it becomes available:
  atlas backups snapshots watch 5f4007f327a3bd7b6f4103c5 --clusterName myDemo`,
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"snapshotIdDesc": "Unique identifier of the snapshot you want to watch.",
			"output":         watchTemplate,
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), watchTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flag.ClusterName)

	return cmd
}

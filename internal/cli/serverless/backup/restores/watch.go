// Copyright 2023 MongoDB Inc
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

package restores

import (
	"context"
	"fmt"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312006/admin"
)

type WatchOpts struct {
	cli.ProjectOpts
	cli.WatchOpts
	id          string
	clusterName string
	store       ServerlessRestoreJobsDescriber
}

var watchTemplate = "\nRestore completed.\n"
var result *atlasv2.ServerlessBackupRestoreJob

func (opts *WatchOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *WatchOpts) watcher() (any, bool, error) {
	var err error
	result, err = opts.store.ServerlessRestoreJob(opts.ConfigProjectID(), opts.clusterName, opts.id)
	if err != nil {
		return nil, false, err
	}
	return nil, result.GetExpired() || result.GetCancelled() || result.GetFailed() || result.HasFinishedAt(), nil
}

func (opts *WatchOpts) Run() error {
	if _, err := opts.Watch(opts.watcher); err != nil {
		return err
	}

	return opts.Print(result)
}

// WatchBuilder atlas serverless backup(s) restore(s) watch.
func WatchBuilder() *cobra.Command {
	opts := new(WatchOpts)
	cmd := &cobra.Command{
		Use:   "watch",
		Short: "Watch the specified backup restore job until it completes.",
		Long: `This command checks the restore job's status periodically until it reaches a completed, failed or canceled status. 
Command finishes once one of the expected statuses are reached.
If you run the command in the terminal, it blocks the terminal session until the resource status completes or fails.
You can interrupt the command's polling at any time with CTRL-C.
` + fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Args: require.NoArgs,
		Annotations: map[string]string{
			"output": watchTemplate,
		},
		Example: `  # Watch the continuous backup restore job with the ID 507f1f77bcf86cd799439011 for the cluster named Cluster0 until it becomes available:
  atlas serverless backup restore watch --restoreJobId 507f1f77bcf86cd799439011 --clusterName Cluster0`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), watchTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
		Deprecated: "please use the 'atlas backup restores watch' command instead. For the migration guide and timeline, visit: https://dochub.mongodb.org/core/flex-migration",
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	opts.AddProjectOptsFlags(cmd)
	cmd.Flags().StringVar(&opts.id, flag.RestoreJobID, "", usage.RestoreJobID)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagRequired(flag.ClusterName)
	_ = cmd.MarkFlagRequired(flag.RestoreJobID)

	return cmd
}

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

package jobs

import (
	"context"
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

type WatchOpts struct {
	cli.ProjectOpts
	cli.WatchOpts
	id          string
	clusterName string
	store       ExportJobsDescriber
}

var watchTemplate = "\nExport completed.\n"
var errExportFailed = errors.New("export failed")
var errExportCancelled = errors.New("export cancelled")

const cancelledState = "Cancelled"
const failedState = "Failed"
const successfulState = "Successful"

func (opts *WatchOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *WatchOpts) watcher() (any, bool, error) {
	result, err := opts.store.ExportJob(opts.ConfigProjectID(), opts.clusterName, opts.id)
	if err != nil {
		return nil, false, err
	}
	return result, result.GetState() == successfulState || result.GetState() == failedState || result.GetState() == cancelledState, nil
}

func (opts *WatchOpts) Run() error {
	result, err := opts.Watch(opts.watcher)
	if err != nil {
		return err
	}

	res, ok := result.(*atlasv2.DiskBackupExportJob)
	if !ok {
		return errExportFailed
	}

	switch res.GetState() {
	case failedState:
		return errExportFailed
	case cancelledState:
		return errExportCancelled
	}

	return opts.Print(nil)
}

// WatchBuilder atlas backup(s) export(s) job(s) watch <exportJobId>.
func WatchBuilder() *cobra.Command {
	opts := new(WatchOpts)
	cmd := &cobra.Command{
		Use:   "watch <exportJobId>",
		Short: "Watch for the specified export job to complete.",
		Long: `This command checks the export job's status periodically until it reaches a completed, cancelled or failed status. 
Once the export reaches the expected status, the command prints "Export completed."
If you run the command in the terminal, it blocks the terminal session until the resource status completes or fails.
You can interrupt the command's polling at any time with CTRL-C.
` + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"exportJobIdDesc": "Unique string that identifies the export job.",
			"output":          watchTemplate,
		},
		Example: `  # Watch the continuous backup restore job with the ID 507f1f77bcf86cd799439011 for the cluster named Cluster0 until it becomes available:
  atlas backup restore watch 507f1f77bcf86cd799439011 --clusterName Cluster0`,
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
	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagRequired(flag.ClusterName)

	return cmd
}

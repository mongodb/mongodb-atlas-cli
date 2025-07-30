// Copyright 2022 MongoDB Inc
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
	"errors"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

const failedStatus = "FAILED"
const completedStatus = "COMPLETED"

var watchTemplate = "\nRestore completed.\n"
var result *atlasv2.DiskBackupSnapshotRestoreJob
var errRestoreFailed = errors.New("restore failed")
var errRestoreExpired = errors.New("restore expired")
var errRestoreCancelled = errors.New("restore cancelled")

type WatchOpts struct {
	cli.ProjectOpts
	cli.WatchOpts
	id            string
	clusterName   string
	isFlexCluster bool
	store         RestoreJobsDescriber
}

func (opts *WatchOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *WatchOpts) watcher() (any, bool, error) {
	var err error
	result, err = opts.store.RestoreJob(opts.ConfigProjectID(), opts.clusterName, opts.id)
	if err != nil {
		return nil, false, err
	}

	return result, stopWatcher(result), nil
}

func stopWatcher(result *atlasv2.DiskBackupSnapshotRestoreJob) bool {
	if result.GetExpired() || result.GetCancelled() || result.GetFailed() || result.HasFinishedAt() {
		return true
	}

	if strings.EqualFold(result.GetDeliveryType(), DeliveryTypeDownload) && len(result.GetDeliveryUrl()) > 0 {
		return true
	}

	return false
}

func (opts *WatchOpts) watcherFlexCluster() (any, bool, error) {
	result, err := opts.store.RestoreFlexClusterJob(opts.ConfigProjectID(), opts.clusterName, opts.id)
	if err != nil {
		return nil, false, err
	}
	return result, strings.EqualFold(result.GetStatus(), completedStatus) || strings.EqualFold(result.GetStatus(), failedStatus), nil
}

func (opts *WatchOpts) Run() error {
	if opts.isFlexCluster {
		return opts.RunFlexCluster()
	}

	return opts.RunDedicatedCluster()
}

func (opts *WatchOpts) RunFlexCluster() error {
	result, err := opts.Watch(opts.watcherFlexCluster)
	if err != nil {
		return err
	}

	res, ok := result.(*atlasv2.FlexBackupRestoreJob20241113)
	if !ok {
		return errRestoreFailed
	}

	if strings.EqualFold(res.GetStatus(), failedStatus) {
		return errRestoreFailed
	}

	return opts.Print(result)
}

func (opts *WatchOpts) RunDedicatedCluster() error {
	result, err := opts.Watch(opts.watcher)
	if err != nil {
		return err
	}

	res, ok := result.(*atlasv2.DiskBackupSnapshotRestoreJob)
	if !ok {
		return errRestoreFailed
	}

	if res.GetFailed() {
		return errRestoreFailed
	}

	if res.GetExpired() {
		return errRestoreExpired
	}

	if res.GetCancelled() {
		return errRestoreCancelled
	}

	return opts.Print(result)
}

// newIsFlexCluster sets the opts.isFlexCluster that indicates if the cluster to create is
// a FlexCluster. The function calls the RestoreFlexClusterJob to get a flex cluster snapshot,
// and it sets the opts.isFlexCluster = true in the event of a cannotUseNotFlexWithFlexApisErrorCode.
func (opts *WatchOpts) newIsFlexCluster() error {
	_, err := opts.store.RestoreFlexClusterJob(opts.ConfigProjectID(), opts.clusterName, opts.id)
	if err == nil {
		opts.isFlexCluster = true
		return nil
	}

	apiError, ok := atlasv2.AsError(err)
	if !ok {
		return err
	}

	if apiError.ErrorCode != CannotUseNotFlexWithFlexApisErrorCode && apiError.ErrorCode != FeatureUnsupported && apiError.ErrorCode != ClusterNotFoundErrorCode {
		return err
	}

	opts.isFlexCluster = false
	return nil
}

// WatchBuilder atlas backup(s) restore(s) job(s) watch <restoreJobId>.
func WatchBuilder() *cobra.Command {
	opts := new(WatchOpts)
	cmd := &cobra.Command{
		Use:   "watch <restoreJobId>",
		Short: "Watch for a restore job to complete.",
		Long: `This command checks the restore job's status periodically until it reaches a completed, failed or canceled status. 
Once the restore reaches the expected status, the command prints "Restore completed."
If you run the command in the terminal, it blocks the terminal session until the resource status completes or fails.
You can interrupt the command's polling at any time with CTRL-C.`,
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"restoreJobIdDesc": "ID of the restore job.",
			"output":           watchTemplate,
		},
		Example: `  # Watch the continuous backup restore job with the ID 507f1f77bcf86cd799439011 for the restore source cluster named Cluster0 until it becomes available:
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
			if err := opts.newIsFlexCluster(); err != nil {
				return nil
			}
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagRequired(flag.ClusterName)

	return cmd
}

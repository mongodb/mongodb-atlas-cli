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

package restores

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
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const (
	automatedRestore   = "automated"
	downloadRestore    = "download"
	pointInTimeRestore = "pointInTime"
)

type StartOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	method                string
	clusterName           string
	targetProjectID       string
	targetClusterName     string
	oplogTS               int
	oplogInc              int
	snapshotID            string
	pointInTimeUTCSeconds int
	store                 store.RestoreJobsCreator
}

func (opts *StartOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var startTemplate = "Restore job '{{.Id}}' successfully started\n"

func (opts *StartOpts) Run() error {
	request := opts.newCloudProviderSnapshotRestoreJob()
	r, err := opts.store.CreateRestoreJobs(opts.ConfigProjectID(), opts.clusterName, request)

	if err != nil {
		return commonerrors.Check(err)
	}

	return opts.Print(r)
}

func (opts *StartOpts) newCloudProviderSnapshotRestoreJob() *admin.DiskBackupSnapshotRestoreJob {
	request := new(admin.DiskBackupSnapshotRestoreJob)
	request.DeliveryType = opts.method

	if opts.targetProjectID != "" {
		request.TargetGroupId = &opts.targetProjectID
	}

	if opts.targetClusterName != "" {
		request.TargetClusterName = &opts.targetClusterName
	}

	if opts.snapshotID != "" {
		request.SnapshotId = &opts.snapshotID
	}

	// Set only in pointInTimeRestore
	if opts.oplogTS != 0 && opts.oplogInc != 0 {
		request.OplogTs = &opts.oplogTS
		request.OplogInc = &opts.oplogInc
	} else if opts.pointInTimeUTCSeconds != 0 {
		// Set only when oplogTS and oplogInc are not set
		request.PointInTimeUTCSeconds = &opts.pointInTimeUTCSeconds
	}

	return request
}

func (opts *StartOpts) isAutomatedRestore() bool {
	return opts.method == automatedRestore
}

func (opts *StartOpts) isPointInTimeRestore() bool {
	return opts.method == pointInTimeRestore
}

func (opts *StartOpts) isDownloadRestore() bool {
	return opts.method == downloadRestore
}

func markRequiredAutomatedRestoreFlags(cmd *cobra.Command) error {
	if err := cmd.MarkFlagRequired(flag.TargetProjectID); err != nil {
		return err
	}

	if err := cmd.MarkFlagRequired(flag.SnapshotID); err != nil {
		return err
	}

	if err := cmd.MarkFlagRequired(flag.TargetClusterName); err != nil {
		return err
	}

	return cmd.MarkFlagRequired(flag.ClusterName)
}

func markRequiredPointInTimeRestoreFlags(cmd *cobra.Command) error {
	if err := cmd.MarkFlagRequired(flag.TargetProjectID); err != nil {
		return err
	}

	return cmd.MarkFlagRequired(flag.TargetClusterName)
}

// atlas backup(s) restore(s) job(s) start <automated|download|pointInTime>.
func StartBuilder() *cobra.Command {
	opts := new(StartOpts)
	cmd := &cobra.Command{
		Use:   fmt.Sprintf("start <%s|%s|%s>", automatedRestore, downloadRestore, pointInTimeRestore),
		Short: "Start a restore job for your project and cluster.",
		Long: `If you create an automated or pointInTime restore job, Atlas removes all existing data on the target cluster prior to the restore.

` + fmt.Sprintf("%s\n%s", fmt.Sprintf(usage.RequiredRole, "Project Owner"), "Atlas supports this command only for M10+ clusters."),
		Args:      require.ExactValidArgs(1),
		ValidArgs: []string{automatedRestore, downloadRestore, pointInTimeRestore},
		Annotations: map[string]string{
			"automated|download|pointInTimeDesc": "Type of restore job to create. Valid values include: automated, download, pointInTime. To learn more about types of restore jobs, see https://www.mongodb.com/docs/atlas/backup-restore-cluster/.",
			"output":                             startTemplate,
		},
		Example: `  # Create an automated restore:
  atlas backup restore start automated \
         --clusterName myDemo \
         --snapshotId 5e7e00128f8ce03996a47179 \
         --targetClusterName myDemo2 \
         --targetProjectId 1a2345b67c8e9a12f3456de7

  # Create a point-in-time restore:
  atlas backup restore start pointInTime \
         --clusterName myDemo \
         --pointInTimeUTCSeconds 1588523147 \
         --targetClusterName myDemo2 \
         --targetProjectId 1a2345b67c8e9a12f3456de7
  
  # Create a download restore:
  atlas backup restore start download \
         --clusterName myDemo \
         --snapshotId 5e7e00128f8ce03996a47179`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.method = args[0]

			if opts.isAutomatedRestore() {
				if err := markRequiredAutomatedRestoreFlags(cmd); err != nil {
					return err
				}
			}

			if opts.isPointInTimeRestore() {
				if err := markRequiredPointInTimeRestoreFlags(cmd); err != nil {
					return err
				}
			}

			if opts.isDownloadRestore() {
				if err := cmd.MarkFlagRequired(flag.SnapshotID); err != nil {
					return err
				}
			}

			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), startTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.snapshotID, flag.SnapshotID, "", usage.RestoreSnapshotID)
	// Atlas uses cluster name
	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)

	cmd.Flags().StringVar(&opts.targetProjectID, flag.TargetProjectID, "", usage.TargetProjectID)
	// Atlas uses cluster name
	cmd.Flags().StringVar(&opts.targetClusterName, flag.TargetClusterName, "", usage.TargetClusterName)
	cmd.Flags().IntVar(&opts.oplogTS, flag.OplogTS, 0, usage.OplogTS)
	cmd.Flags().IntVar(&opts.oplogInc, flag.OplogInc, 0, usage.OplogInc)
	cmd.Flags().IntVar(&opts.pointInTimeUTCSeconds, flag.PointInTimeUTCMillis, 0, usage.PointInTimeUTCMillis)
	_ = cmd.Flags().MarkDeprecated(flag.PointInTimeUTCMillis, fmt.Sprintf("please use --%s instead", flag.PointInTimeUTCSeconds))
	cmd.Flags().IntVar(&opts.pointInTimeUTCSeconds, flag.PointInTimeUTCSeconds, 0, usage.PointInTimeUTCSeconds)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.ClusterName)

	return cmd
}

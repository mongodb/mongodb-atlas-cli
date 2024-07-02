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
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/commonerrors"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const (
	automatedRestore   = "automated"
	downloadRestore    = "download"
	pointInTimeRestore = "pointInTime"
)

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	deliveryType          string
	clusterName           string
	targetProjectID       string
	targetClusterName     string
	oplogTS               int
	oplogInc              int
	snapshotID            string
	pointInTimeUTCSeconds int
	store                 store.ServerlessRestoreJobsCreator
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var createTemplate = "Restore job '{{.Id}}' successfully started\n"

var ErrInvalidDeliveryType = errors.New("delivery type invalid, choose 'automated', 'download' or 'pointInTime'")

func (opts *CreateOpts) Run() error {
	request := opts.newCloudProviderSnapshotRestoreJob()
	r, err := opts.store.ServerlessCreateRestoreJobs(opts.ConfigProjectID(), opts.clusterName, request)

	if err != nil {
		return commonerrors.Check(err)
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newCloudProviderSnapshotRestoreJob() *atlasv2.ServerlessBackupRestoreJob {
	request := new(atlasv2.ServerlessBackupRestoreJob)
	request.DeliveryType = opts.deliveryType

	if opts.targetProjectID != "" {
		request.TargetGroupId = opts.targetProjectID
	}

	if opts.targetClusterName != "" {
		request.TargetClusterName = opts.targetClusterName
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

func (opts *CreateOpts) isAutomatedRestore() bool {
	return opts.deliveryType == automatedRestore
}

func (opts *CreateOpts) isPointInTimeRestore() bool {
	return opts.deliveryType == pointInTimeRestore
}

func (opts *CreateOpts) isDownloadRestore() bool {
	return opts.deliveryType == downloadRestore
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

// atlas serverless backup(s) restore(s) job(s) create <automated|download|pointInTime>.
func CreateBuilder() *cobra.Command {
	opts := new(CreateOpts)
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Start a restore job for your serverless instance.",
		Long: `If you create an automated or pointInTime restore job, Atlas removes all existing data on the target cluster prior to the restore.

` + fmt.Sprintf("%s\n%s", fmt.Sprintf(usage.RequiredRole, "Project Owner"), "Atlas supports this command only for M10+ clusters."),
		Args: require.NoArgs,
		Example: `  # Create an automated restore:
  atlas serverless backup restore create \
         --deliveryType automated \
         --clusterName myDemo \
         --snapshotId 5e7e00128f8ce03996a47179 \
         --targetClusterName myDemo2 \
         --targetProjectId 1a2345b67c8e9a12f3456de7

  # Create a point-in-time restore:
  atlas serverless backup restore create \
         --deliveryType pointInTime \
         --clusterName myDemo \
         --pointInTimeUTCSeconds 1588523147 \
         --targetClusterName myDemo2 \
         --targetProjectId 1a2345b67c8e9a12f3456de7
  
  # Create a download restore:
  atlas serverless backup restore create \
         --deliveryType download \
         --clusterName myDemo \
         --snapshotId 5e7e00128f8ce03996a47179`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			set := false

			if opts.isAutomatedRestore() {
				set = true
				if err := markRequiredAutomatedRestoreFlags(cmd); err != nil {
					return err
				}
			}

			if opts.isPointInTimeRestore() {
				set = true
				if err := markRequiredPointInTimeRestoreFlags(cmd); err != nil {
					return err
				}
			}

			if opts.isDownloadRestore() {
				set = true
				if err := cmd.MarkFlagRequired(flag.SnapshotID); err != nil {
					return err
				}
			}

			if !set {
				return ErrInvalidDeliveryType
			}

			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.snapshotID, flag.SnapshotID, "", usage.SnapshotID)
	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().StringVar(&opts.deliveryType, flag.DeliveryType, "", usage.DeliveryType)
	cmd.Flags().StringVar(&opts.targetProjectID, flag.TargetProjectID, "", usage.TargetProjectID)
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
	_ = cmd.MarkFlagRequired(flag.DeliveryType)

	return cmd
}

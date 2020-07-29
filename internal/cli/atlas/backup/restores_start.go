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

package backup

import (
	"errors"
	"fmt"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/output"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const (
	automatedRestore   = "automated"
	downloadRestore    = "download"
	pointInTimeRestore = "pointInTime"
	onlyFor            = "'%s' can only be used with %s"
)

type RestoresStartOpts struct {
	cli.GlobalOpts
	method               string
	clusterName          string
	targetProjectID      string
	targetClusterName    string
	oplogTS              int64
	oplogInc             int64
	snapshotID           string
	pointInTimeUTCMillis int64
	outputTemplate       string
	store                store.RestoreJobsCreator
}

func (opts *RestoresStartOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

var automatedTemplate = "Restoring cluster {{.TargetClusterName}} using snapshot {{.SnapshotId}}"
var pointInTimeTemplate = "Restoring cluster {{.TargetClusterName}} using point in time provided"
var downloadTemplate = "Links to download the snapshot:"

func (opts *RestoresStartOpts) Run() error {
	request := opts.newCloudProviderSnapshotRestoreJob()
	r, err := opts.store.CreateRestoreJobs(opts.ConfigProjectID(), opts.clusterName, request)

	if err != nil {
		return err
	}

	return output.Print(config.Default(), opts.outputTemplate, r)
}

func (opts *RestoresStartOpts) newCloudProviderSnapshotRestoreJob() *atlas.CloudProviderSnapshotRestoreJob {
	request := new(atlas.CloudProviderSnapshotRestoreJob)
	request.DeliveryType = opts.method

	if opts.targetProjectID != "" {
		request.TargetGroupID = opts.targetProjectID
	}

	if opts.targetClusterName != "" {
		request.TargetClusterName = opts.targetClusterName
	}

	if opts.snapshotID != "" {
		request.SnapshotID = opts.snapshotID
	}

	// Set only in pointInTimeRestore
	if opts.oplogTS != 0 && opts.oplogInc != 0 {
		request.OplogTs = opts.oplogTS
		request.OplogInc = opts.oplogInc
	} else if opts.pointInTimeUTCMillis != 0 {
		//set only when oplogTS and oplogInc are not set
		request.PointInTimeUTCSeconds = opts.pointInTimeUTCMillis
	}

	return request
}

func (opts *RestoresStartOpts) isAutomatedRestore() bool {
	if opts.method == automatedRestore {
		opts.outputTemplate = automatedTemplate
		return true
	}
	return false
}

func (opts *RestoresStartOpts) isPointInTimeRestore() bool {
	if opts.method == pointInTimeRestore {
		opts.outputTemplate = pointInTimeTemplate
		return true
	}
	return false
}

func (opts *RestoresStartOpts) isDownloadRestore() bool {
	if opts.method == downloadRestore {
		opts.outputTemplate = downloadTemplate
		return true
	}
	return false
}

func (opts *RestoresStartOpts) validateParams() error {
	if opts.clusterName == "" {
		return errors.New("needs clusterName")
	}

	return nil
}

func markRequiredAutomatedRestoreFlags(cmd *cobra.Command) error {
	if err := cmd.MarkFlagRequired(flag.TargetProjectID); err != nil {
		return err
	}

	if err := cmd.MarkFlagRequired(flag.SnapshotID); err != nil {
		return err
	}

	if err := cmd.MarkFlagRequired(flag.TargetClusterID); err != nil {
		return err
	}

	return cmd.MarkFlagRequired(flag.ClusterName)
}

func markRequiredPointInTimeRestoreFlags(cmd *cobra.Command) error {
	if err := cmd.MarkFlagRequired(flag.TargetProjectID); err != nil {
		return err
	}

	if err := cmd.MarkFlagRequired(flag.TargetClusterID); err != nil {
		return err
	}

	return cmd.MarkFlagRequired(flag.ClusterName)
}

func markRequiredDownloadRestoreFlags(cmd *cobra.Command) error {
	if err := cmd.MarkFlagRequired(flag.SnapshotID); err != nil {
		return err
	}

	return nil
}

// mongocli atlas backup(s) restore(s) job(s) start <automated|download|pointInTime>
func RestoresStartBuilder() *cobra.Command {
	opts := new(RestoresStartOpts)
	cmd := &cobra.Command{
		Use:       fmt.Sprintf("start <%s|%s|%s>", automatedRestore, downloadRestore, pointInTimeRestore),
		Short:     description.StartRestore,
		Args:      cobra.ExactValidArgs(1),
		ValidArgs: []string{automatedRestore, downloadRestore, pointInTimeRestore},
		PreRunE: func(cmd *cobra.Command, args []string) error {
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
				if err := markRequiredDownloadRestoreFlags(cmd); err != nil {
					return err
				}
			}

			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.method = args[0]

			if e := opts.validateParams(); e != nil {
				return e
			}

			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.snapshotID, flag.SnapshotID, "", usage.SnapshotID)
	// Atlas uses cluster name
	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)

	cmd.Flags().StringVar(&opts.targetProjectID, flag.TargetProjectID, "", usage.TargetProjectID)
	// Atlas uses cluster name
	cmd.Flags().StringVar(&opts.targetClusterName, flag.TargetClusterName, "", usage.TargetClusterName)
	cmd.Flags().Int64Var(&opts.oplogTS, flag.OplogTS, 0, usage.OplogTS)
	cmd.Flags().Int64Var(&opts.oplogInc, flag.OplogInc, 0, usage.OplogInc)
	cmd.Flags().Int64Var(&opts.pointInTimeUTCMillis, flag.PointInTimeUTCMillis, 0, usage.PointInTimeUTCMillis)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}

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

package cli

import (
	"errors"

	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

const (
	automatedRestore = "AUTOMATED_RESTORE"
	httpRestore      = "HTTP"
)

type atlasBackupsRestoresStartOpts struct {
	*globalOpts
	method               string
	clusterName          string
	clusterID            string
	targetProjectID      string
	targetClusterID      string
	targetClusterName    string
	checkpointID         string
	oplogTs              string
	oplogInc             int64
	snapshotID           string
	expirationHours      int64
	expires              string
	maxDownloads         int64
	pointInTimeUTCMillis float64
	store                store.ContinuousJobCreator
}

func (opts *atlasBackupsRestoresStartOpts) init() error {
	if opts.ProjectID() == "" {
		return errMissingProjectID
	}

	var err error
	opts.store, err = store.New()
	return err
}

func (opts *atlasBackupsRestoresStartOpts) Run() error {
	request := opts.newContinuousJobRequest()

	result, err := opts.store.CreateContinuousRestoreJob(opts.ProjectID(), opts.fromCluster(), request)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func (opts *atlasBackupsRestoresStartOpts) newContinuousJobRequest() *atlas.ContinuousJobRequest {
	request := new(atlas.ContinuousJobRequest)
	request.Delivery.MethodName = opts.method
	request.SnapshotID = opts.snapshotID

	if opts.isAutomatedRestore() {
		request.Delivery.TargetGroupID = opts.targetProjectID
		opts.setTargetCluster(request)

		if opts.oplogTs != "" && opts.oplogInc != 0 {
			request.OplogTs = opts.oplogTs
			request.OplogInc = opts.oplogInc
		}
		if opts.pointInTimeUTCMillis != 0 {
			request.PointInTimeUTCMillis = opts.pointInTimeUTCMillis
		}
	}

	if opts.isHTTP() {
		if opts.expires != "" {
			request.Delivery.Expires = opts.expires
		}
		if opts.maxDownloads > 0 {
			request.Delivery.MaxDownloads = opts.maxDownloads
		}
		if opts.expirationHours > 0 {
			request.Delivery.ExpirationHours = opts.expirationHours
		}
	}
	return request
}

func (opts *atlasBackupsRestoresStartOpts) fromCluster() string {
	if opts.clusterName != "" {
		return opts.clusterName
	}
	return opts.clusterID
}

func (opts *atlasBackupsRestoresStartOpts) setTargetCluster(out *atlas.ContinuousJobRequest) {
	if opts.targetClusterID != "" {
		out.Delivery.TargetClusterID = opts.targetClusterID
	} else if opts.targetClusterName != "" {
		out.Delivery.TargetClusterName = opts.targetClusterName
	}
}

func (opts *atlasBackupsRestoresStartOpts) isAutomatedRestore() bool {
	return opts.method == automatedRestore
}

func (opts *atlasBackupsRestoresStartOpts) isHTTP() bool {
	return opts.method == httpRestore
}

func (opts *atlasBackupsRestoresStartOpts) validateParams() error {
	if opts.clusterName == "" && opts.clusterID == "" {
		return errors.New("missing clusterName or clusterId")
	}
	if opts.clusterName != "" && opts.clusterID != "" {
		return errors.New("clusterName and clusterId specified")
	}
	if !opts.isAutomatedRestore() && opts.checkpointID != "" {
		return errors.New("checkpointId can only be used with AUTOMATED_RESTORE")
	}
	if !opts.isAutomatedRestore() && opts.oplogTs != "" {
		return errors.New("oplogTs can only be used with AUTOMATED_RESTORE")
	}
	if !opts.isAutomatedRestore() && opts.oplogInc > 0 {
		return errors.New("oplogInc can only be used with AUTOMATED_RESTORE")
	}
	if !opts.isAutomatedRestore() && opts.pointInTimeUTCMillis > 0 {
		return errors.New("pointInTimeUTCMillis can only be used with AUTOMATED_RESTORE")
	}
	if opts.isAutomatedRestore() && opts.targetProjectID == "" {
		return errors.New("missing targetProject")
	}
	if opts.isAutomatedRestore() && opts.targetClusterID == "" && opts.targetClusterName == "" {
		return errors.New("missing targetClusterID or targetClusterName")
	}
	if !opts.isHTTP() && opts.expires != "" {
		return errors.New("expires can only be used with HTTP")
	}
	if !opts.isHTTP() && opts.maxDownloads > 0 {
		return errors.New("maxDownloads can only be used with HTTP")
	}
	if !opts.isHTTP() && opts.expirationHours > 0 {
		return errors.New("expirationHours can only be used with HTTP")
	}
	return nil
}

// mongocli atlas backup(s) restore(s) job(s) start
func AtlasBackupsRestoresStartBuilder() *cobra.Command {
	opts := &atlasBackupsRestoresStartOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:       "start",
		Short:     "Start a restore job.",
		Args:      cobra.ExactValidArgs(1),
		ValidArgs: []string{automatedRestore, httpRestore},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.method = args[0]

			if e := opts.validateParams(); e != nil {
				return e
			}

			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.snapshotID, flags.SnapshotID, "", usage.SnapshotID)
	// Atlas uses cluster name
	cmd.Flags().StringVar(&opts.clusterName, flags.ClusterName, "", usage.ClusterName)
	// C/OM uses cluster ID
	cmd.Flags().StringVar(&opts.clusterID, flags.ClusterID, "", usage.ClusterID)

	// For Automatic restore
	cmd.Flags().StringVar(&opts.targetProjectID, flags.TargetProjectID, "", usage.TargetProjectID)
	cmd.Flags().StringVar(&opts.targetClusterID, flags.TargetClusterID, "", usage.TargetClusterID)
	cmd.Flags().StringVar(&opts.targetClusterName, flags.TargetClusterName, "", usage.TargetClusterName)
	cmd.Flags().StringVar(&opts.checkpointID, flags.CheckpointID, "", usage.CheckpointID)
	cmd.Flags().StringVar(&opts.oplogTs, flags.OplogTs, "", usage.OplogTs)
	cmd.Flags().Int64Var(&opts.oplogInc, flags.OplogInc, 0, usage.OplogInc)
	cmd.Flags().Float64Var(&opts.pointInTimeUTCMillis, flags.PointInTimeUTCMillis, 0, usage.PointInTimeUTCMillis)

	// For http restore
	cmd.Flags().StringVar(&opts.expires, flags.Expires, "", usage.Expires)
	cmd.Flags().Int64Var(&opts.maxDownloads, flags.MaxDownloads, 0, usage.MaxDownloads)
	cmd.Flags().Int64Var(&opts.expirationHours, flags.ExpirationHours, 0, usage.ExpirationHours)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}

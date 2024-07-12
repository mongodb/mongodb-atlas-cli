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

package jobs

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
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	exportBucketID string
	snapshotID     string
	clusterName    string
	customData     map[string]string
	store          store.ExportJobsCreator
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var createTemplate = "Export job '{{.Id}}' created in a bucket with ID '{{.ExportBucketId}}'.\n"

func (opts *CreateOpts) Run() error {
	createRequest := opts.newExportJob()

	r, err := opts.store.CreateExportJob(opts.ConfigProjectID(), opts.clusterName, createRequest)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newExportJob() *atlasv2.DiskBackupExportJobRequest {
	customData := make([]atlasv2.BackupLabel, 0, len(opts.customData))
	for k, v := range opts.customData {
		customData = append(customData, newBackupLabel(k, v))
	}
	createRequest := &atlasv2.DiskBackupExportJobRequest{
		SnapshotId:     opts.snapshotID,
		ExportBucketId: opts.exportBucketID,
		CustomData:     &customData,
	}
	return createRequest
}

func newBackupLabel(k, v string) atlasv2.BackupLabel {
	return atlasv2.BackupLabel{
		Key:   &k,
		Value: &v,
	}
}

// atlas backup(s) export(s) job(s) –-clusterName clusterName [--bucketId bucketId] [--projectId projectId].
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Export one backup snapshot for an M10 or higher Atlas cluster to an existing AWS S3 bucket.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Example: `  # The following command exports one backup snapshot of the ExampleCluster cluster to an existing AWS S3 bucket:
  atlas backup export jobs create --clusterName ExampleCluster --bucketId 62c569f85b7a381c093cc539 --snapshotId 62c808ceeb4e021d850dfe1b --customData name=test,info=test`,
		Annotations: map[string]string{
			"output": createTemplate,
		},
		Args: require.NoArgs,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
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
	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().StringVar(&opts.exportBucketID, flag.BucketID, "", usage.BucketID)
	cmd.Flags().StringVar(&opts.snapshotID, flag.SnapshotID, "", usage.SnapshotID)
	cmd.Flags().StringToStringVar(&opts.customData, flag.CustomData, nil, usage.CustomData)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.ClusterName)
	_ = cmd.MarkFlagRequired(flag.BucketID)
	_ = cmd.MarkFlagRequired(flag.SnapshotID)

	return cmd
}

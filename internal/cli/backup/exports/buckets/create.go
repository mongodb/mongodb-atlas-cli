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

package buckets

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/validate"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	iamRoleID     string
	bucketName    string
	cloudProvider string
	store         store.ExportBucketsCreator
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var createTemplate = "Export destination created using '{{.BucketName}}'.\n"

func (opts *CreateOpts) Run() error {
	createRequest := opts.newExportBucket()

	r, err := opts.store.CreateExportBucket(opts.ConfigProjectID(), createRequest)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newExportBucket() *atlasv2.DiskBackupSnapshotExportBucket {
	createRequest := &atlasv2.DiskBackupSnapshotExportBucket{
		BucketName:    &opts.bucketName,
		CloudProvider: &opts.cloudProvider,
		IamRoleId:     &opts.iamRoleID,
	}
	return createRequest
}

// atlas backup(s) export(s) bucket(s) create <bucketName> --cloudProvider AWS.
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create <bucketName>",
		Short: "Create an export destination for Atlas backups using an existing AWS S3 bucket.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Example: `  # The following command creates an export destination for Atlas backups using the existing AWS S3 bucket named test-bucket:
  atlas backup export buckets create test-bucket --cloudProvider AWS --iamRoleId 12345678f901a234dbdb00ca`,
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"bucketNameDesc": "Name of the existing S3 bucket that the provided role ID is authorized to access.",
			"output":         createTemplate,
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				func() error {
					return validate.OptionalObjectID(opts.iamRoleID)
				},
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.bucketName = args[0]
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.cloudProvider, flag.CloudProvider, "", usage.CloudProvider)
	cmd.Flags().StringVar(&opts.iamRoleID, flag.IAMRoleID, "", usage.ExportBucketIAMRoleID)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_, _ = cmd.MarkFlagRequired(flag.CloudProvider), cmd.MarkFlagRequired(flag.IAMRoleID)

	return cmd
}

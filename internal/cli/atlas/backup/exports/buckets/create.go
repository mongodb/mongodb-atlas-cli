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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas/mongodbatlas"
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

var createTemplate = "Bucket '{{.BucketName}}' created.\n"

func (opts *CreateOpts) Run() error {
	createRequest := opts.newExportBucket()

	r, err := opts.store.CreateExportBucket(opts.ConfigProjectID(), createRequest)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newExportBucket() *mongodbatlas.CloudProviderSnapshotExportBucket {
	createRequest := &mongodbatlas.CloudProviderSnapshotExportBucket{
		BucketName:    opts.bucketName,
		CloudProvider: opts.cloudProvider,
		IAMRoleID:     opts.iamRoleID,
	}
	return createRequest
}

// mongocli atlas backup(s) export(s) bucket(s) create <bucketName> --cloudProvider AWS.
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create <bucketName>",
		Short: "Create new export bucket.",
		Example: fmt.Sprintf(`  Create new export bucket:
  $ %s config rename myProfile testProfile`, config.BinName()),
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"args":           "bucketName",
			"bucketNameDesc": "Name of the bucket that the provided role ID is authorized to access.",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.bucketName = args[0]
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.cloudProvider, flag.CloudProvider, "", usage.CloudProvider)
	cmd.Flags().StringVar(&opts.iamRoleID, flag.IAMRoleID, "", usage.ExportBucketIAMRoleID)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_, _ = cmd.MarkFlagRequired(flag.CloudProvider), cmd.MarkFlagRequired(flag.IAMRoleID)

	return cmd
}

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
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312002/admin"
)

var describeTemplate = `ID	BUCKET NAME	CLOUD PROVIDER	IAM ROLE ID
{{.Id}}	{{.BucketName}}	{{.CloudProvider}}	{{.IamRoleId}}
`

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=describe_mock_test.go -package=buckets . ExportBucketsDescriber

type ExportBucketsDescriber interface {
	DescribeExportBucket(string, string) (*atlasv2.DiskBackupSnapshotExportBucketResponse, error)
}

type DescribeOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	bucketID string
	store    ExportBucketsDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.DescribeExportBucket(opts.ConfigProjectID(), opts.bucketID)
	if err != nil {
		return err
	}
	return opts.Print(r)
}

// atlas backup(s) export(s) bucket(s) describe --bucketId bucketId [--projectId projectId].
func DescribeBuilder() *cobra.Command {
	opts := &DescribeOpts{}
	cmd := &cobra.Command{
		Use:     "describe",
		Aliases: []string{"get"},
		Short:   "Return one snapshot export bucket.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output": describeTemplate,
		},
		Example: `  # Return the details for the continuous backup export bucket with the ID dbdb00ca12345678f901a234:
  atlas backup exports buckets describe dbdb00ca12345678f901a234`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.bucketID, flag.BucketID, "", usage.BucketID)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagRequired(flag.BucketID)

	return cmd
}

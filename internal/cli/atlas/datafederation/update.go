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

// This code was autogenerated at 2023-06-21T13:32:21+01:00. Note: Manual updates are allowed, but may be overwritten.

package datafederation

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	store "github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/admin"
)

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store         store.DataFederationUpdater
	fs            afero.Fs
	name          string
	filename      string
	region        string
	awsRoleId     string
	awsTestBucket string
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var updateTemplate = `Pipeline {{.Name}} updated.`

func (opts *UpdateOpts) Run() error {
	updateRequest, err := opts.newUpdateRequest()
	if err != nil {
		return err
	}

	r, err := opts.store.UpdateDataFederation(opts.ConfigProjectID(), opts.name, updateRequest)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *UpdateOpts) newUpdateRequest() (*admin.DataLakeTenant, error) {
	if opts.filename != "" {
		tenant := admin.DataLakeTenant{}
		if err := file.Load(opts.fs, opts.filename, &tenant); err != nil {
			return nil, err
		}
		tenant.Name = &opts.name
		return &tenant, nil
	}

	return &admin.DataLakeTenant{
		Name: &opts.name,
		DataProcessRegion: &admin.DataLakeDataProcessRegion{
			CloudProvider: "AWS",
			Region:        opts.region,
		},
		CloudProviderConfig: &admin.DataLakeCloudProviderConfig{
			Aws: admin.DataLakeAWSCloudProviderConfig{
				RoleId:       opts.awsRoleId,
				TestS3Bucket: opts.awsTestBucket,
			},
		},
	}, nil
}

// atlas dataFederation update <name> [--projectId projectId].
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "update <name>",
		Short: "Modify the details of the specified data federation database for your project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"nameDesc": "Name of the data federation database",
			"output":   updateTemplate,
		},
		Example: `# update data lake pipeline:
  atlas dataFederation update DataFederation1
`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.region, flag.Region, "", usage.DataLakeRegion)
	cmd.Flags().StringVar(&opts.awsRoleId, flag.AWSRoleID, "", usage.DataLakeRole)
	cmd.Flags().StringVar(&opts.awsTestBucket, flag.AWSTestS3Bucket, "", usage.DataLakeTestBucket)
	cmd.Flags().StringVarP(&opts.filename, flag.File, flag.FileShort, "", usage.DataFederationFile)

	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.Region)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.AWSRoleID)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.AWSTestS3Bucket)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}

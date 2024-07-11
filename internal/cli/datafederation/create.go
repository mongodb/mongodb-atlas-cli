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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20231115014/admin"
)

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store         store.DataFederationCreator
	fs            afero.Fs
	name          string
	filename      string
	region        string
	awsRoleID     string
	awsTestBucket string
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var createTemplate = `Data federation {{.Name}} created.`

func (opts *CreateOpts) Run() error {
	createRequest, err := opts.newCreateRequest()
	if err != nil {
		return err
	}

	r, err := opts.store.CreateDataFederation(opts.ConfigProjectID(), createRequest)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newCreateRequest() (*admin.DataLakeTenant, error) {
	if opts.filename != "" {
		tenant := admin.DataLakeTenant{}
		if err := file.Load(opts.fs, opts.filename, &tenant); err != nil {
			return nil, err
		}
		tenant.Name = &opts.name
		return &tenant, nil
	}

	ret := admin.NewDataLakeTenant()
	ret.Name = &opts.name

	if opts.region != "" {
		ret.DataProcessRegion = &admin.DataLakeDataProcessRegion{
			CloudProvider: "AWS",
			Region:        opts.region,
		}
	}

	if opts.awsRoleID != "" || opts.awsTestBucket != "" {
		ret.CloudProviderConfig = &admin.DataLakeCloudProviderConfig{
			Aws: admin.DataLakeAWSCloudProviderConfig{
				RoleId:       opts.awsRoleID,
				TestS3Bucket: opts.awsTestBucket,
			},
		}
	}

	return ret, nil
}

// atlas dataFederation create <name> [--projectId projectId].
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Creates a new Data Federation database.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"nameDesc": "Name of the data federation database.",
			"output":   createTemplate,
		},
		Example: `# create data federation database:
  atlas dataFederation create DataFederation1 --region us_east_1 --awsRoleId role --awsTestS3Bucket bucket
`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
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

	cmd.Flags().StringVar(&opts.region, flag.Region, "", usage.DataLakeRegion)
	cmd.Flags().StringVar(&opts.awsRoleID, flag.AWSRoleID, "", usage.DataLakeRole)
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

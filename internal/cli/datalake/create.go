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

package datalake

import (
	"context"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas/mongodbatlas"
)

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store      store.DataLakeCreator
	name       string
	awsRoleID  string
	testBucket string
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var createTemplate = "Data lake '{{.Name}}' created.\n"

func (opts *CreateOpts) Run() error {
	createRequest := opts.newDataLakeRequest()

	r, err := opts.store.CreateDataLake(opts.ConfigProjectID(), createRequest)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newDataLakeRequest() *mongodbatlas.DataLakeCreateRequest {
	return &mongodbatlas.DataLakeCreateRequest{
		Name: opts.name,
		CloudProviderConfig: &mongodbatlas.CloudProviderConfig{
			AWSConfig: mongodbatlas.AwsCloudProviderConfig{
				RoleID:       opts.awsRoleID,
				TestS3Bucket: opts.testBucket,
			},
		},
	}
}

// atlas datalake(s) create <name> --projectId projectId.
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Create a new federated database instance for your project.",
		Long:  `To learn more about Atlas Data Federation (previously named Atlas Data Lake), see https://www.mongodb.com/docs/atlas/data-federation/overview/.`,
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"nameDesc": "Name of the federated database instance to create.",
			"output":   createTemplate,
		},
		Example: `  # Create a federated database instance named myFDI in the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas dataLakes create myFDI --projectId 5e2211c17a3e5a48f5497de3 --output json`,
		Deprecated: "Please use 'atlas datafederation create'",
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

	cmd.Flags().StringVar(&opts.awsRoleID, flag.Role, "", usage.DataLakeRole)
	cmd.Flags().StringVar(&opts.testBucket, flag.TestBucket, "", usage.DataLakeTestBucket)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.Role)
	_ = cmd.MarkFlagRequired(flag.TestBucket)

	return cmd
}

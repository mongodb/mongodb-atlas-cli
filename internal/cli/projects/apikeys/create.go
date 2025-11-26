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

package apikeys

import (
	"context"
	"fmt"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312010/admin"
)

var createTemplate = `API Key '{{.Id}}' created.
Public API Key {{.PublicKey}}
Private API Key {{.PrivateKey}}
`

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=create_mock_test.go -package=apikeys . ProjectAPIKeyCreator

type ProjectAPIKeyCreator interface {
	CreateProjectAPIKey(string, *atlasv2.CreateAtlasProjectApiKey) (*atlasv2.ApiKeyUserDetails, error)
}

type CreateOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	store       ProjectAPIKeyCreator
	description string
	roles       []string
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *CreateOpts) Run() error {
	apiKeyInput := &atlasv2.CreateAtlasProjectApiKey{
		Desc:  opts.description,
		Roles: opts.roles,
	}

	r, err := opts.store.CreateProjectAPIKey(opts.ProjectID, apiKeyInput)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas project apiKey create --roles roles --description description.
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create an organization API key and assign it to your project.",
		Long: `MongoDB returns the private API key only once. After you run this command, immediately copy, save, and secure both the public and private API keys.
If you don't provide an organization level role, the API Key defaults to organization member of the project parent organization.

` + fmt.Sprintf(usage.RequiredOneOfRoles, "Project User Admin or Organization User Admin to manage organization level roles"),
		Annotations: map[string]string{
			"output": createTemplate,
		},
		Args: require.NoArgs,
		Example: `  # Create an organization API key with the GROUP_OWNER role and assign it to the project with ID 5e2211c17a3e5a48f5497de3:
  atlas projects apiKeys create --desc "My API key" --projectId 5e1234c17a3e5a48f5497de3 --role GROUP_OWNER --output json
  
  # Create an organization API key with the GROUP_SEARCH_INDEX_EDITOR and GROUP_DATABASE_ACCESS_ADMIN roles and assign it to the project with ID 5e2211c17a3e5a48f5497de3:
  atlas projects apiKeys create --desc "My API key" --projectId 5e1234c17a3e5a48f5497de3 --role GROUP_SEARCH_INDEX_EDITOR,GROUP_DATABASE_ACCESS_ADMIN --output json`,
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

	cmd.Flags().StringSliceVar(&opts.roles, flag.Role, []string{}, usage.CreateProjectAPIKeyRoles)
	cmd.Flags().StringVar(&opts.description, flag.Description, "", usage.APIKeyDescription)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagRequired(flag.Description)
	_ = cmd.MarkFlagRequired(flag.Role)

	return cmd
}

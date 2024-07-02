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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type AssignOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	id    string
	roles []string
	store store.ProjectAPIKeyAssigner
}

func (opts *AssignOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *AssignOpts) newAssignAPIKey() *atlasv2.UpdateAtlasProjectApiKey {
	return &atlasv2.UpdateAtlasProjectApiKey{
		Roles: &opts.roles,
	}
}

var updateTemplate = "API Key successfully assigned.\n"

func (opts *AssignOpts) Run() error {
	err := opts.store.AssignProjectAPIKey(opts.ConfigProjectID(), opts.id, opts.newAssignAPIKey())
	if err != nil {
		return err
	}
	return opts.Print(nil)
}

// atlas project(s) apiKey(s)|apikey(s) assign <ID> [--role role][--projectId projectId].
func AssignBuilder() *cobra.Command {
	opts := new(AssignOpts)
	cmd := &cobra.Command{
		Use:     "assign <ID>",
		Aliases: []string{"update"},
		Args:    require.ExactArgs(1),
		Short:   "Assign the specified organization API key to your project and modify the API key's roles for the project.",
		Long: `When you modify the roles for an organization API key with this command, the values you specify overwrite the existing roles assigned to the API key.
		
To view possible values for the ID argument, run atlas organizations apiKeys list.`,
		Annotations: map[string]string{
			"IDDesc": "Unique 24-digit string that identifies your API key.",
			"output": updateTemplate,
		},
		Example: `  # Assign an organization API key with the ID 5f46ae53d58b421fe3edc115 and grant the GROUP_DATA_ACCESS_READ_WRITE role for the project with ID 5e2211c17a3e5a48f5497de3:
  atlas projects apiKeys assign 5f46ae53d58b421fe3edc115 --projectId 5e1234c17a3e5a48f5497de3 --role GROUP_DATA_ACCESS_READ_WRITE --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringSliceVar(&opts.roles, flag.Role, []string{}, usage.ProjectAPIKeyRoles)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flag.Role)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}

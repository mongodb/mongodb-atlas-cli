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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	*cli.DeleteOpts
	cli.GlobalOpts
	store store.ProjectAPIKeyDeleter
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteProjectAPIKey, opts.ConfigProjectID())
}

// atlas project(s) apiKey(s)|apikey(s) delete <ID> [--projectId projectId].
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("API Key '%s' deleted\n", "APIKey not deleted"),
	}

	cmd := &cobra.Command{
		Use:     "delete <ID>",
		Aliases: []string{"rm"},
		Short:   "Remove the specified organization API key from your project.",
		Long: `The API key still exists at the organization level. To reassign the organization API key to a project, run the  atlas projects apiKeys assign command.
		
To view possible values for the ID argument, run atlas organizations apiKeys list.

` + fmt.Sprintf(usage.RequiredRole, "Project User Admin"),
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"IDDesc": "Unique 24-digit string that identifies your API key.",
			"output": opts.SuccessMessage(),
		},
		Example: `  # Remove an organization API key with the ID 5f46ae53d58b421fe3edc115 from the project with ID 5e2211c17a3e5a48f5497de3:
  atlas projects apiKeys delete 5f46ae53d58b421fe3edc115 --projectId 5e1234c17a3e5a48f5497de3`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(opts.ValidateProjectID, opts.initStore(cmd.Context())); err != nil {
				return err
			}
			opts.Entry = args[0]
			return opts.Prompt()
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}

// Copyright 2021 MongoDB Inc
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

package azure

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

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=delete_mock_test.go -package=azure . PrivateEndpointDeleter

type PrivateEndpointDeleter interface {
	DeletePrivateEndpoint(string, string, string) error
}

type DeleteOpts struct {
	cli.ProjectOpts
	*cli.DeleteOpts
	store PrivateEndpointDeleter
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeletePrivateEndpoint, opts.ConfigProjectID(), provider)
}

// atlas privateEndpoint(s) delete <privateEndpointId> --projectId projectId.
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Private endpoint '%s' deleted\n", "Private endpoint not deleted"),
	}
	cmd := &cobra.Command{
		Use:     "delete <privateEndpointId>",
		Aliases: []string{"rm"},
		Short:   "Remove the specified Azure private endpoint from your project.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"privateEndpointIdDesc": "Unique 24-character alphanumeric string that identifies the private endpoint in Atlas.",
			"output":                opts.SuccessMessage(),
		},
		Example: `  # Remove the Azure private endpoint with the ID 5f4fc14da2b47835a58c63a2 from the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas privateEndpoints azure delete 5f4fc14da2b47835a58c63a2 --projectId 5e2211c17a3e5a48f5497de3`,
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

	opts.AddProjectOptsFlags(cmd)

	return cmd
}

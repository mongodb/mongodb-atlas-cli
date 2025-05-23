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

package organizations

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

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=delete_mock_test.go -package=organizations . OrganizationDeleter

type OrganizationDeleter interface {
	DeleteOrganization(string) error
}

type DeleteOpts struct {
	*cli.DeleteOpts
	store OrganizationDeleter
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteOrganization)
}

// DeleteBuilder atlas organization(s) delete <ID> [--force].
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Organization '%s' deleted\n", "Organization not deleted"),
	}
	cmd := &cobra.Command{
		Use:     "delete <ID>",
		Aliases: []string{"rm"},
		Short:   "Remove the specified organization.",
		Long: `Organizations with active projects can't be removed.

` + fmt.Sprintf(usage.RequiredRole, "Organization Owner"),
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"IDDesc": "Unique 24-digit string that identifies the organization.",
			"output": opts.SuccessMessage(),
		},
		Example: `  # Remove the organization with the ID 5e2211c17a3e5a48f5497de3:
  atlas organizations delete 5e2211c17a3e5a48f5497de3`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.initStore(cmd.Context())(); err != nil {
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

	return cmd
}

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

package settings

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
	cli.GlobalOpts
	*cli.DeleteOpts
	store store.AlertConfigurationDeleter
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteAlertConfiguration, opts.ConfigProjectID())
}

// atlas alerts config(s) delete <ID> --projectId projectId [--confirm].
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Alert config '%s' deleted\n", "Alert config not deleted"),
	}

	cmd := &cobra.Command{
		Use:     "delete <alertConfigId>",
		Aliases: []string{"rm"},
		Short:   "Remove the specified alert configuration for your project.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"alertConfigIdDesc": "Unique identifier of the alert configuration to delete.",
			"output":            opts.SuccessMessage(),
		},
		Example: `  # Remove the alert configuration with the ID 5d1113b25a115342acc2d1aa in the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas alerts settings delete 5d1113b25a115342acc2d1aa --projectId 5e2211c17a3e5a48f5497de3`,
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

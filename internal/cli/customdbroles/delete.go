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

package customdbroles

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

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=delete_mock_test.go -package=customdbroles . DatabaseRoleDeleter

type DatabaseRoleDeleter interface {
	DeleteDatabaseRole(string, string) error
}

type DeleteOpts struct {
	cli.ProjectOpts
	*cli.DeleteOpts
	store DatabaseRoleDeleter
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteDatabaseRole, opts.ConfigProjectID(), opts.Entry)
}

// atlas dbrole(s) delete <roleName> --force.
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Custom database role '%s' deleted\n", "Custom database role not deleted"),
	}
	cmd := &cobra.Command{
		Use:     "delete <roleName>",
		Aliases: []string{"rm"},
		Short:   "Remove the specified custom database role from your project.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"roleNameDesc": "Name of the custom role to delete.",
			"output":       opts.SuccessMessage(),
		},
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

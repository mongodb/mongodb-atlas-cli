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

package users

import (
	"context"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	*cli.DeleteOpts
	store store.UserDeleter
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteUser)
}

// mongocli iam users(s) delete <ID> [--force].
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("User '%s' deleted\n", "User not deleted"),
	}
	cmd := &cobra.Command{
		Use:     "delete <userId>",
		Aliases: []string{"rm"},
		Short:   "Remove the specified Ops Manager user.",
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"userIdDesc": "Unique 24-digit string that identifies the user in Ops Manager.",
			"output":     opts.SuccessMessage(),
		},
		Example: `  # Remove the Ops Manager user with the ID 5e44445ef10fab20b49c0f31:
  mongocli iam users delete 5e44445ef10fab20b49c0f31`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.initStore(cmd.Context())(); err != nil {
				return err
			}
			opts.Entry = args[0]
			return opts.Prompt()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	return cmd
}

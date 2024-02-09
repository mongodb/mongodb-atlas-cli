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

package globalaccesslists

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
	store store.GlobalAPIKeyWhitelistDeleter
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteGlobalAPIKeyWhitelist)
}

// mongocli iam globalAccessList(s) delete <ID> [--force].
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Global access list entry '%s' deleted\n", "Global access list entry not deleted"),
	}

	cmd := &cobra.Command{
		Use:     "delete <ID>",
		Aliases: []string{"rm"},
		Short:   "Remove the specified IP access list entry for your global API key.",
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"IDDesc": "Unique 24-digit string that identifies the access list entry.",
		},
		Example: `  # Remove the IP access list entry with the ID 5f5bad7a57aef32b04ed0210 from the access list for the global API key:
  mongocli iam globalAccessLists delete 5f5bad7a57aef32b04ed0210`,
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

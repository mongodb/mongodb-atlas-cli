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

package globalwhitelist

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	*cli.DeleteOpts
	store store.GlobalAPIKeyWhitelistDeleter
}

func (opts *DeleteOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteGlobalAPIKeyWhitelist)
}

// mongocli iam globalWhitelist(s) delete <ID> [--force]
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Whitelist entry '%s' deleted\n", "Whitelist entry not deleted"),
	}

	cmd := &cobra.Command{
		Use:     "delete <ID>",
		Aliases: []string{"rm"},
		Short:   deleteWhitelist,
		Args:    cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.init(); err != nil {
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

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

package cli

import (
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/convert"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type atlasDBUsersDeleteOpts struct {
	globalOpts
	deleteOpts
	authDB string
	store  store.DatabaseUserDeleter
}

func (opts *atlasDBUsersDeleteOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *atlasDBUsersDeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteDatabaseUser, opts.authDB, opts.ProjectID())
}

// mongocli atlas dbuser(s) delete <username> --force
func AtlasDBUsersDeleteBuilder() *cobra.Command {
	opts := &atlasDBUsersDeleteOpts{}
	opts.successMessage = "DB user '%s' deleted\n"
	opts.failMessage = "DB user not deleted"
	cmd := &cobra.Command{
		Use:     "delete [username]",
		Short:   description.DeleteDBUser,
		Aliases: []string{"rm"},
		Args:    cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(opts.initStore); err != nil {
				return err
			}
			opts.entry = args[0]
			return opts.Confirm()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.confirm, flag.Force, false, usage.Force)
	cmd.Flags().StringVar(&opts.authDB, flag.AuthDB, convert.AdminDB, usage.AuthDB)

	cmd.Flags().StringVar(&opts.projectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}

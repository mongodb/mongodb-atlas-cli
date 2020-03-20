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
	"fmt"

	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/convert"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/messages"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type opsManagerDBUsersDeleteOpts struct {
	*globalOpts
	*deleteOpts
	authDB string
	store  store.AutomationPatcher
}

func (opts *opsManagerDBUsersDeleteOpts) init() error {
	if opts.ProjectID() == "" {
		return errMissingProjectID
	}

	var err error
	opts.store, err = store.New()
	return err
}

func (opts *opsManagerDBUsersDeleteOpts) Run() error {
	current, err := opts.store.GetAutomationConfig(opts.ProjectID())

	if err != nil {
		return err
	}

	convert.RemoveUser(current, opts.entry, opts.authDB)

	if err = opts.store.UpdateAutomationConfig(opts.ProjectID(), current); err != nil {
		return err
	}

	fmt.Print(messages.DeploymentStatus(config.OpsManagerURL(), opts.ProjectID()))

	return nil
}

// mongocli atlas dbuser(s) delete <username> [--projectId projectId]
func OpsManagerDBUsersDeleteBuilder() *cobra.Command {
	opts := &opsManagerDBUsersDeleteOpts{
		globalOpts: newGlobalOpts(),
		deleteOpts: &deleteOpts{
			successMessage: "DB user '%s' deleted\n",
			failMessage:    "DB user not deleted",
		},
	}
	cmd := &cobra.Command{
		Use:     "delete [username]",
		Short:   "Delete a database user for a project.",
		Aliases: []string{"rm"},
		Args:    cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.init(); err != nil {
				return err
			}
			opts.entry = args[0]
			return opts.Confirm()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.authDB, flags.AuthDB, convert.AdminDB, usage.AuthDB)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}

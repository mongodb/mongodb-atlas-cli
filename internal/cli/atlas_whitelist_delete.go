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

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mcli/internal/flags"
	"github.com/mongodb/mcli/internal/prompts"
	"github.com/mongodb/mcli/internal/store"
	"github.com/mongodb/mcli/internal/usage"
	"github.com/spf13/cobra"
)

type atlasWhitelistDeleteOpts struct {
	*globalOpts
	entry   string
	confirm bool
	store   store.ProjectIPWhitelistDeleter
}

func (opts *atlasWhitelistDeleteOpts) init() error {
	if opts.ProjectID() == "" {
		return errMissingProjectID
	}

	s, err := store.New()

	if err != nil {
		return err
	}

	opts.store = s
	return nil
}

func (opts *atlasWhitelistDeleteOpts) Run() error {
	err := opts.store.DeleteProjectIPWhitelist(opts.ProjectID(), opts.entry)

	if err != nil {
		return err
	}

	fmt.Printf("Project whitelist entry '%s' deleted\n", opts.entry)

	return nil
}

func (opts *atlasWhitelistDeleteOpts) Confirm() error {
	if opts.confirm {
		return nil
	}
	prompt := prompts.NewDeleteConfirmation(opts.entry)
	return survey.AskOne(prompt, &opts.confirm)
}

// mcli atlas whitelist delete <entry> --force
func AtlasWhitelistDeleteBuilder() *cobra.Command {
	opts := &atlasWhitelistDeleteOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:     "delete [entry]",
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

	cmd.Flags().BoolVar(&opts.confirm, flags.Force, false, usage.Force)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}

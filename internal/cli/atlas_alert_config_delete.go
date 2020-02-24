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
	"github.com/mongodb/mcli/internal/flags"
	"github.com/mongodb/mcli/internal/store"
	"github.com/mongodb/mcli/internal/usage"
	"github.com/spf13/cobra"
)

type atlasAlertConfigDeleteOpts struct {
	*globalOpts
	*deleteOpts
	store          store.AlertConfigurationDeleter
}

func (opts *atlasAlertConfigDeleteOpts) init() error {
	if opts.ProjectID() == "" {
		return errMissingProjectID
	}

	var err error
	opts.store, err = store.New()
	return err
}

func (opts *atlasAlertConfigDeleteOpts) Run() error {
	return opts.DeleteFromProject(opts.store. DeleteAlertConfiguration, opts.ProjectID())
}

// mcli atlas alerts config(s) delete id --projectId projectId [--confirm]
func AtlasAlertConfigDeleteBuilder() *cobra.Command {
	opts := &atlasAlertConfigDeleteOpts{
		globalOpts:     newGlobalOpts(),
		deleteOpts: &deleteOpts{
			successMessage: "Alert Config '%s' deleted\n",
			failMessage:    "Alert Config not deleted",
		},
	}
	cmd := &cobra.Command{
		Use:     "delete [id]",
		Short:   "Delete an Atlas Alert Config.",
		Aliases: []string{"rm", "Delete", "Remove"},
		Args:    cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.entry = args[0]
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.confirm, flags.Force, false, usage.Force)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}

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
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type atlasAlertsConfigDeleteOpts struct {
	globalOpts
	deleteOpts
	store store.AlertConfigurationDeleter
}

func (opts *atlasAlertsConfigDeleteOpts) initStore() error {
	var err error
	opts.store, err = store.New()
	return err
}

func (opts *atlasAlertsConfigDeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteAlertConfiguration, opts.ProjectID())
}

// mongocli atlas alerts config(s) delete id --projectId projectId [--confirm]
func AtlasAlertsConfigDeleteBuilder() *cobra.Command {
	opts := new(atlasAlertsConfigDeleteOpts)
	opts.successMessage = "Alert config '%s' deleted\n"
	opts.failMessage = "Alert config not deleted"
	cmd := &cobra.Command{
		Use:     "delete [id]",
		Short:   description.DeleteAlertsConfig,
		Aliases: []string{"rm", "Delete", "Remove"},
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

	cmd.Flags().BoolVar(&opts.confirm, flags.Force, false, usage.Force)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}

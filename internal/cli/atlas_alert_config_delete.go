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

	"github.com/mongodb/mcli/internal/flags"
	"github.com/mongodb/mcli/internal/store"
	"github.com/mongodb/mcli/internal/usage"
	"github.com/spf13/cobra"
)

type atlasAlertConfigDeleteOpts struct {
	*globalOpts
	entry          string
	successMessage string
	failMessage    string
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
	error := opts.store.DeleteAlertConfiguration(opts.ProjectID(), opts.entry)

	if error != nil {
		fmt.Println(opts.failMessage, opts.entry)
	} else {
		fmt.Println(opts.successMessage, opts.entry)
	}

	return error
}

// mcli atlas alert_config(s) delete id --projectId projectId [--confirm]
func AtlasAlertConfigDeleteBuilder() *cobra.Command {
	opts := &atlasAlertConfigDeleteOpts{
		globalOpts:     newGlobalOpts(),
		successMessage: "Alert Config '%s' deleted\n",
		failMessage:    "Alert Config '%s' not deleted",
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

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}

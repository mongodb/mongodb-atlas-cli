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
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type opsManagerAutomationStatusOpts struct {
	*globalOpts
	store store.AutomationStatusGetter
}

func (opts *opsManagerAutomationStatusOpts) init() error {
	if opts.ProjectID() == "" {
		return errMissingProjectID
	}
	var err error
	opts.store, err = store.New()
	return err
}

func (opts *opsManagerAutomationStatusOpts) Run() error {
	result, err := opts.store.GetAutomationStatus(opts.projectID)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

// mongocli ops-manager automation status [--projectId projectId]
func OpsManagerAutomationStatusBuilder() *cobra.Command {
	opts := &opsManagerAutomationStatusOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Show the current status of the automation config.",
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}

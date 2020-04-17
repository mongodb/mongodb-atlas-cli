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

type opsManagerAutomationShowOpts struct {
	globalOpts
	store store.AutomationGetter
}

func (opts *opsManagerAutomationShowOpts) initStore() error {
	var err error
	opts.store, err = store.New()
	return err
}

func (opts *opsManagerAutomationShowOpts) Run() error {
	r, err := opts.store.GetAutomationConfig(opts.projectID)

	if err != nil {
		return err
	}

	return json.PrettyPrint(r)
}

// mongocli ops-manager automation show [--projectId projectId]
func OpsManagerAutomationShowBuilder() *cobra.Command {
	opts := &opsManagerAutomationShowOpts{}
	cmd := &cobra.Command{
		Use:    "show",
		Args:   cobra.NoArgs,
		Hidden: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}

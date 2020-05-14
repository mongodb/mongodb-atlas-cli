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
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type opsManagerAutomationDescribeOpts struct {
	globalOpts
	store store.AutomationGetter
}

func (opts *opsManagerAutomationDescribeOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *opsManagerAutomationDescribeOpts) Run() error {
	r, err := opts.store.GetAutomationConfig(opts.ProjectID())

	if err != nil {
		return err
	}

	return json.PrettyPrint(r)
}

// mongocli ops-manager automation describe [--projectId projectId]
func OpsManagerAutomationDescribeBuilder() *cobra.Command {
	opts := &opsManagerAutomationDescribeOpts{}
	cmd := &cobra.Command{
		Use:     "describe",
		Aliases: []string{"show", "get"},
		Args:    cobra.NoArgs,
		Hidden:  true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.projectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}

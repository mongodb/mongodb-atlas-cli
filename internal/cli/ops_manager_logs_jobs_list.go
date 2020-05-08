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
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/ops-manager/opsmngr"
)

type opsManagerLogsJobsListOpts struct {
	globalOpts
	verbose bool
	store   store.LogJobLister
}

func (opts *opsManagerLogsJobsListOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *opsManagerLogsJobsListOpts) Run() error {
	result, err := opts.store.LogCollectionJobs(opts.ProjectID(), opts.newLogListOptions())
	if err != nil {
		return err
	}
	return json.PrettyPrint(result)
}

func (opts *opsManagerLogsJobsListOpts) newLogListOptions() *opsmngr.LogListOptions {
	return &opsmngr.LogListOptions{Verbose: opts.verbose}
}

// mongocli om logs jobs list --verbose verbose [--projectId projectId]
func OpsManagerLogsJobsListOptsBuilder() *cobra.Command {
	opts := &opsManagerLogsJobsListOpts{}
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   description.ListLogCollectionJobs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.verbose, flags.Verbose, false, usage.Verbose)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}

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
	om "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type opsManagerLogsDownloadOpts struct {
	globalOpts
	id string
	filePath string
	store store.LogsLister
}

func (opts *opsManagerLogsDownloadOpts) initStore() error {
	var err error
	opts.store, err = store.New()
	return err
}

func (opts *opsManagerLogsDownloadOpts) Run() error {
	result, err := opts.store.ListLogJobs(opts.ProjectID(), opts.newLogListOptions())
	if err != nil {
		return err
	}
	return json.PrettyPrint(result)
}

func (opts *opsManagerLogsDownloadOpts) newLogListOptions() *om.LogListOptions {
	return &om.LogListOptions{Verbose: opts.Verbose}
}

// mongocli om logs download id [--filePath filePath] [--projectId projectId]
func OpsManagerLogsDownloadOptsBuilder() *cobra.Command {
	opts := &opsManagerLogsDownloadOpts{}
	cmd := &cobra.Command{
		Use:     "download",
		Short:   description.DownloadLogs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.id, flags.ID, "", usage.Verbose)
	cmd.Flags().StringVar(&opts.filePath, flags.FilePath, "", usage.Verbose)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}

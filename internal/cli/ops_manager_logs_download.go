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
	"io"
	"os"

	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

type opsManagerLogsDownloadOpts struct {
	globalOpts
	id    string
	out   string
	fs    afero.Fs
	store store.LogJobsDownloader
}

func (opts *opsManagerLogsDownloadOpts) initStore() error {
	var err error
	opts.store, err = store.New()
	return err
}

func (opts *opsManagerLogsDownloadOpts) Run() error {
	out, err := opts.newWriteCloser()
	if err != nil {
		return err
	}

	if err := opts.store.DownloadLogJob(opts.ProjectID(), opts.id, out); err != nil {
		return err
	}

	return nil
}

func (opts *opsManagerLogsDownloadOpts) newWriteCloser() (io.WriteCloser, error) {
	// Create file only if is not there already (don't overwrite)
	ff := os.O_CREATE | os.O_TRUNC | os.O_WRONLY | os.O_EXCL
	f, err := opts.fs.OpenFile(opts.out, ff, 0777)
	return f, err
}

// mongocli om logs download id [--out out] [--projectId projectId]
func OpsManagerLogsDownloadOptsBuilder() *cobra.Command {
	opts := &opsManagerLogsDownloadOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "download [id]",
		Short: description.DownloadLogs,
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.out, flags.Out, flags.OutShort, "", usage.LogOut)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flags.Out)

	return cmd
}

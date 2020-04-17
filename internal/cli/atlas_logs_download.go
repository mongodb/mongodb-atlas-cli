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

	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

type atlasLogsDownloadOpts struct {
	globalOpts
	host  string
	name  string
	out   string
	start string
	end   string
	fs    afero.Fs
	store store.LogsDownloader
}

func (opts *atlasLogsDownloadOpts) initStore() error {
	var err error
	opts.store, err = store.New()
	return err
}

func (opts *atlasLogsDownloadOpts) Run() error {
	f, err := opts.newWriteCloser()
	if err != nil {
		return err
	}
	defer f.Close()

	r := opts.newDateRangeOpts()
	if err := opts.store.DownloadLog(opts.ProjectID(), opts.host, opts.name, f, r); err != nil {
		return err
	}
	return nil
}

func (opts *atlasLogsDownloadOpts) output() string {
	if opts.out == "" {
		opts.out = opts.name
	}
	return opts.out
}

func (opts *atlasLogsDownloadOpts) newWriteCloser() (io.WriteCloser, error) {
	// Create file only if is not there already (don't overwrite)
	ff := os.O_CREATE | os.O_TRUNC | os.O_WRONLY | os.O_EXCL
	f, err := opts.fs.OpenFile(opts.output(), ff, 0777)
	return f, err
}

func (opts *atlasLogsDownloadOpts) newDateRangeOpts() *atlas.DateRangetOptions {
	return &atlas.DateRangetOptions{
		StartDate: opts.start,
		EndDate:   opts.end,
	}
}

// mongocli atlas logs download [hostname] [logname] [--type type] [--output destination] [--projectId projectId]
func AtlasLogsDownloadBuilder() *cobra.Command {
	opts := &atlasLogsDownloadOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:     "download [hostname] [logname]",
		Short:   description.ListDisks,
		Aliases: []string{"ls"},
		Args:    cobra.ExactArgs(2),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.host = args[0]
			opts.name = args[1]

			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.out, flags.Out, flags.OutShort, "", usage.End)

	cmd.Flags().StringVar(&opts.start, flags.Start, "", usage.Start)
	cmd.Flags().StringVar(&opts.end, flags.End, "", usage.End)
	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}

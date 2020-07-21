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

package atlas

import (
	"fmt"
	"io"
	"os"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/search"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type LogsDownloadOpts struct {
	cli.GlobalOpts
	host  string
	name  string
	out   string
	start string
	end   string
	fs    afero.Fs
	store store.LogsDownloader
}

func (opts *LogsDownloadOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *LogsDownloadOpts) Run() error {
	f, err := opts.newWriteCloser()
	if err != nil {
		return err
	}

	r := opts.newDateRangeOpts()
	if err := opts.store.DownloadLog(opts.ConfigProjectID(), opts.host, opts.name, f, r); err != nil {
		_ = opts.handleError(f)
		return err
	}
	return f.Close()
}

func (opts *LogsDownloadOpts) output() string {
	if opts.out == "" {
		opts.out = opts.name
	}
	return opts.out
}

func (opts *LogsDownloadOpts) handleError(f io.Closer) error {
	_ = f.Close()
	return opts.fs.Remove(opts.output())
}

func (opts *LogsDownloadOpts) newWriteCloser() (io.WriteCloser, error) {
	// Create file only if is not there already (don't overwrite)
	ff := os.O_CREATE | os.O_TRUNC | os.O_WRONLY | os.O_EXCL
	f, err := opts.fs.OpenFile(opts.output(), ff, 0777)
	return f, err
}

func (opts *LogsDownloadOpts) newDateRangeOpts() *atlas.DateRangetOptions {
	return &atlas.DateRangetOptions{
		StartDate: opts.start,
		EndDate:   opts.end,
	}
}

var logNames = []string{"mongodb.gz", "mongos.gz", "mongodb-audit-log.gz", "mongos-audit-log.gz"}

// mongocli atlas logs download <hostname> <logname> [--type type] [--output destination] [--projectId projectId]
func LogsDownloadBuilder() *cobra.Command {
	const argsN = 2
	opts := &LogsDownloadOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "download <hostname> <logname>",
		Short: description.LogsDownload,
		Long:  description.LogsDownloadLong,
		Args:  cobra.ExactArgs(argsN),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.host = args[0]
			opts.name = args[1]
			if !search.StringInSlice(logNames, opts.name) {
				return fmt.Errorf("<logname> must be one of %s", logNames)
			}
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.out, flag.Out, flag.OutShort, "", usage.LogOut)

	cmd.Flags().StringVar(&opts.start, flag.Start, "", usage.LogStart)
	cmd.Flags().StringVar(&opts.end, flag.End, "", usage.LogEnd)
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagFilename(flag.Out)

	return cmd
}

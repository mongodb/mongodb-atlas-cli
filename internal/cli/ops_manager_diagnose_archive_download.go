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

	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/ops-manager/opsmngr"
)

type opsManagerDiagnoseArchiveDownloadOpts struct {
	globalOpts
	out     string
	limit   int64
	minutes int64
	fs      afero.Fs
	store   store.ArchivesDownloader
}

func (opts *opsManagerDiagnoseArchiveDownloadOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *opsManagerDiagnoseArchiveDownloadOpts) Run() error {
	out, err := opts.newWriteCloser()
	if err != nil {
		return err
	}
	defer out.Close()

	if err := opts.store.DownloadArchive(opts.ProjectID(), opts.newDiagnosticsListOpts(), out); err != nil {
		return err
	}

	return nil
}

func (opts *opsManagerDiagnoseArchiveDownloadOpts) newWriteCloser() (io.WriteCloser, error) {
	// Create file only if is not there already (don't overwrite)
	ff := os.O_CREATE | os.O_TRUNC | os.O_WRONLY | os.O_EXCL
	f, err := opts.fs.OpenFile(opts.out, ff, 0777)
	return f, err
}

func (opts *opsManagerDiagnoseArchiveDownloadOpts) newDiagnosticsListOpts() *opsmngr.DiagnosticsListOpts {
	return &opsmngr.DiagnosticsListOpts{
		Limit:   opts.limit,
		Minutes: opts.minutes,
	}
}

// mongocli om diagnose-archive download [--out out] [--projectId projectId]
func OpsManagerDiagnoseArchiveDownloadBuilder() *cobra.Command {
	opts := &opsManagerDiagnoseArchiveDownloadOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:     "download",
		Aliases: []string{"get"},
		Short:   description.DownloadDiagnoseArchive,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.out, flags.Out, flags.OutShort, "diagnose-archive.tar.gz", usage.DiagnoseOut)
	cmd.Flags().Int64Var(&opts.limit, flags.Limit, 0, usage.ArchiveLimit)
	cmd.Flags().Int64Var(&opts.minutes, flags.Minutes, 0, usage.ArchiveMinutes)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}

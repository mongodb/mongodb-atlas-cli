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

package opsmanager

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/ops-manager/opsmngr"
)

type DiagnoseArchiveDownloadOpts struct {
	cli.GlobalOpts
	cli.DownloaderOpts
	limit   int64
	minutes int64
	store   store.ArchivesDownloader
}

func (opts *DiagnoseArchiveDownloadOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *DiagnoseArchiveDownloadOpts) Run() error {
	out, err := opts.NewWriteCloser()
	if err != nil {
		return err
	}

	if err := opts.store.DownloadArchive(opts.ConfigProjectID(), opts.newDiagnosticsListOpts(), out); err != nil {
		_ = opts.OnError(out)
		return err
	}

	return out.Close()
}

func (opts *DiagnoseArchiveDownloadOpts) newDiagnosticsListOpts() *opsmngr.DiagnosticsListOpts {
	return &opsmngr.DiagnosticsListOpts{
		Limit:   opts.limit,
		Minutes: opts.minutes,
	}
}

// mongocli om diagnose-archive download [--out out] [--projectId projectId]
func DiagnoseArchiveDownloadBuilder() *cobra.Command {
	opts := &DiagnoseArchiveDownloadOpts{}
	opts.Fs = afero.NewOsFs()
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

	cmd.Flags().StringVarP(&opts.Out, flag.Out, flag.OutShort, "diagnose-archive.tar.gz", usage.DiagnoseOut)
	cmd.Flags().Int64Var(&opts.limit, flag.Limit, 0, usage.ArchiveLimit)
	cmd.Flags().Int64Var(&opts.minutes, flag.Minutes, 0, usage.ArchiveMinutes)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagFilename(flag.Out)

	return cmd
}

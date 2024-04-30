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

package serverusage

import (
	"context"
	"fmt"

	"github.com/andreaangiolillo/mongocli-test/internal/cli"
	"github.com/andreaangiolillo/mongocli-test/internal/config"
	"github.com/andreaangiolillo/mongocli-test/internal/flag"
	"github.com/andreaangiolillo/mongocli-test/internal/store"
	"github.com/andreaangiolillo/mongocli-test/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/ops-manager/opsmngr"
)

type DownloadOpts struct {
	cli.DownloaderOpts
	startDate     string
	endDate       string
	format        string
	skipRedaction bool
	store         store.ServerUsageReportDownloader
}

var downloadMessage = "Download of %s completed.\n"

func (opts *DownloadOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DownloadOpts) Run() error {
	f, err := opts.NewWriteCloser()
	if err != nil {
		return err
	}

	if err := opts.store.DownloadServerUsageReport(opts.newServerTypeOptions(), f); err != nil {
		_ = opts.OnError(f)
		return err
	}
	if !opts.ShouldDownloadToStdout() {
		fmt.Printf(downloadMessage, opts.Out)
	}
	return f.Close()
}

func (opts *DownloadOpts) newServerTypeOptions() *opsmngr.ServerTypeOptions {
	redact := !opts.skipRedaction
	return &opsmngr.ServerTypeOptions{
		StartDate:  opts.startDate,
		EndDate:    opts.endDate,
		FileFormat: opts.format,
		Redact:     &redact,
	}
}

func (opts *DownloadOpts) initDefaultOut() {
	if opts.Out == "" {
		opts.Out = fmt.Sprintf("serverUsage-%s-%s.%s", opts.startDate, opts.endDate, opts.format)
	}
}

// mongocli om serverUsage download [--startDate startDate] [--endDate endDate] [--format format] [--force] [--output destination] [--projectId projectId].
func DownloadBuilder() *cobra.Command {
	opts := &DownloadOpts{}
	opts.Fs = afero.NewOsFs()
	cmd := &cobra.Command{
		Use:     "download",
		Short:   "Download the server usage report.",
		Example: `  mongocli ops-manager serverUsage download --endDate 2022-12-12 --startDate 2022-01-01`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.initDefaultOut()
			return opts.initStore(cmd.Context())()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.startDate, flag.StartDate, "", usage.ServerUsageStartDate)
	cmd.Flags().StringVar(&opts.endDate, flag.EndDate, "", usage.ServerUsageEndDate)
	cmd.Flags().StringVar(&opts.format, flag.Format, "tar.gz", usage.ServerUsageFormat)
	cmd.Flags().StringVar(&opts.Out, flag.Out, "", usage.LogOut)
	cmd.Flags().BoolVar(&opts.skipRedaction, flag.SkipRedaction, false, usage.ServerUsageSkipRedacted)
	cmd.Flags().BoolVar(&opts.Force, flag.Force, false, usage.ForceFile)

	_ = cmd.MarkFlagRequired(flag.StartDate)
	_ = cmd.MarkFlagRequired(flag.EndDate)
	cmd.MarkFlagsRequiredTogether(flag.StartDate, flag.EndDate)

	return cmd
}

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

package logs

import (
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"slices"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
)

var errEmptyLog = errors.New("log is empty")

type DownloadOpts struct {
	cli.GlobalOpts
	cli.DownloaderOpts
	host       string
	name       string
	start      int64
	end        int64
	decompress bool
	store      store.LogsDownloader
}

var downloadMessage = "Download of %s completed.\n"

func (opts *DownloadOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

// maxBytes  1k each write to avoid compression bomb.
const maxBytes = 1024

func (opts *DownloadOpts) write(w io.Writer, r io.Reader) error {
	if !opts.decompress {
		_, err := io.Copy(w, r)
		return err
	}

	gr, errGz := gzip.NewReader(r)
	if errGz != nil {
		return errGz
	}

	written := false
	for {
		n, err := io.CopyN(w, gr, maxBytes)
		if n > 0 {
			written = true
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}

	if !written {
		return errEmptyLog
	}

	return nil
}

func (opts *DownloadOpts) Run() error {
	w, err := opts.NewWriteCloser()
	if err != nil {
		return err
	}
	defer w.Close()

	r, err := opts.store.DownloadLog(opts.newHostLogsParams())
	if err != nil {
		_ = opts.OnError(w)
		return err
	}
	defer r.Close()

	if err := opts.write(w, r); err != nil {
		_ = opts.OnError(w)
		return err
	}

	if !opts.ShouldDownloadToStdout() {
		fmt.Printf(downloadMessage, opts.Out)
	}

	return nil
}

func (opts *DownloadOpts) initDefaultOut() error {
	if opts.Out == "" {
		opts.Out = strings.ReplaceAll(opts.name, ".gz", ".log.gz")
	}
	return nil
}

func (opts *DownloadOpts) newHostLogsParams() *admin.GetHostLogsApiParams {
	fileBaseName := strings.TrimSuffix(opts.name, filepath.Ext(opts.name))
	params := &admin.GetHostLogsApiParams{
		GroupId:  opts.ConfigProjectID(),
		HostName: opts.host,
		LogName:  fileBaseName,
	}
	if opts.start != 0 {
		params.StartDate = &opts.start
	}
	if opts.end != 0 {
		params.EndDate = &opts.end
	}
	return params
}

// atlas logs download <hostname> <mongodb.gz|mongos.gz|mongosqld.gz|mongodb-audit-log.gz|mongos-audit-log.gz> [--force] [--output destination] [--projectId projectId].
func DownloadBuilder() *cobra.Command {
	const argsN = 2
	opts := &DownloadOpts{}
	opts.Fs = afero.NewOsFs()
	cmd := &cobra.Command{
		Use:   "download <hostname> <mongodb.gz|mongos.gz|mongosqld.gz|mongodb-audit-log.gz|mongos-audit-log.gz>",
		Short: "Download a compressed file that contains the MongoDB logs for the specified host.",
		Long: `This command downloads a file with a .gz extension.

To find the hostnames for an Atlas project, use the process list command.

` + fmt.Sprintf(usage.RequiredRole, "Project Data Access Read/Write"),
		Args: cobra.MatchAll(
			require.ExactArgs(argsN),
			func(cmd *cobra.Command, args []string) error {
				if !slices.Contains(cmd.ValidArgs, args[1]) {
					return fmt.Errorf("<logname> must be one of %s", cmd.ValidArgs)
				}
				return nil
			},
		),
		Example: `  # Download the mongodb log file from the host atlas-123abc-shard-00-00.111xx.mongodb.net for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas logs download  atlas-123abc-shard-00-00.111xx.mongodb.net mongodb.gz --projectId 5e2211c17a3e5a48f5497de3`,
		Annotations: map[string]string{
			"hostnameDesc": "Label that identifies the host that stores the log files that you want to download.",
			"mongodb.gz|mongos.gz|mongosqld.gz|mongodb-audit-log.gz|mongos-audit-log.gzDesc": "Log file that you want to return.",
			"output": downloadMessage,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.host = args[0]
			opts.name = args[1]
			return opts.PreRunE(opts.ValidateProjectID, opts.initStore(cmd.Context()), opts.initDefaultOut)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},

		ValidArgs: []string{"mongodb.gz", "mongos.gz", "mongosqld.gz", "mongodb-audit-log.gz", "mongos-audit-log.gz"},
	}

	cmd.Flags().StringVar(&opts.Out, flag.Out, "", usage.LogOut)
	cmd.Flags().Int64Var(&opts.start, flag.Start, 0, usage.LogStart)
	cmd.Flags().Int64Var(&opts.end, flag.End, 0, usage.LogEnd)
	cmd.Flags().BoolVar(&opts.Force, flag.Force, false, usage.ForceFile)
	cmd.Flags().BoolVarP(&opts.decompress, flag.Decompress, flag.DecompressShort, false, usage.Decompress)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagFilename(flag.Out)

	return cmd
}

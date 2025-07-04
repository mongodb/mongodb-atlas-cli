// Copyright 2024 MongoDB Inc
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

package instance

import (
	"context"
	"fmt"
	"io"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

var downloadMessage = "Download of %s completed.\n"

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=download_mock_test.go -package=instance . StreamsDownloader

type StreamsDownloader interface {
	DownloadAuditLog(*atlasv2.DownloadStreamTenantAuditLogsApiParams) (io.ReadCloser, error)
}

type DownloadOpts struct {
	cli.ProjectOpts
	cli.DownloaderOpts
	tenantName string
	start      int64
	end        int64
	store      StreamsDownloader
}

func (opts *DownloadOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DownloadOpts) Run() error {
	params := atlasv2.DownloadStreamTenantAuditLogsApiParams{
		GroupId:    opts.ConfigProjectID(),
		TenantName: opts.tenantName,
	}

	if opts.start != 0 {
		params.StartDate = &opts.start
	}

	if opts.end != 0 {
		params.EndDate = &opts.end
	}

	f, err := opts.store.DownloadAuditLog(&params)
	if err != nil {
		return err
	}

	defer f.Close()

	out, err := opts.NewWriteCloser()
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, f)
	return err
}

// DownloadBuilder
// atlas streams download [tenantName] --projectId [projectID].
func DownloadBuilder() *cobra.Command {
	const argsN = 1
	opts := &DownloadOpts{}
	opts.Fs = afero.NewOsFs()
	cmd := &cobra.Command{
		Use:   "download <tenantName>",
		Short: "Download a compressed file that contains the logs for the specified Atlas Stream Processing instance.",
		Long:  `This command downloads a file with a .gz extension. ` + fmt.Sprintf(usage.RequiredRole, "Project Data Access Read/Write"),
		Args: cobra.MatchAll(
			require.ExactArgs(argsN),
		),
		Example: `  # Download the audit log file from the instance myProcessor for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas streams instance download myProcessor --projectId 5e2211c17a3e5a48f5497de3`,
		Annotations: map[string]string{
			"tenantNameDesc": "Label that identifies the tenant that stores the log files that you want to download.",
			"output":         downloadMessage,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.tenantName = args[0]
			return opts.PreRunE(opts.ValidateProjectID, opts.initStore(cmd.Context()))
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.Out, flag.Out, "", usage.LogOut)
	cmd.Flags().Int64Var(&opts.start, flag.Start, 0, usage.LogStart)
	cmd.Flags().Int64Var(&opts.end, flag.End, 0, usage.LogEnd)
	cmd.Flags().BoolVar(&opts.Force, flag.Force, false, usage.ForceFile)

	opts.AddProjectOptsFlags(cmd)

	_ = cmd.MarkFlagRequired(flag.Out)
	_ = cmd.MarkFlagFilename(flag.Out)

	return cmd
}

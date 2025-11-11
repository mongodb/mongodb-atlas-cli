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

package snapshots

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312009/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=download_mock_test.go -package=snapshots . Downloader

type Downloader interface {
	DownloadFlexClusterSnapshot(string, string, *atlasv2.FlexBackupSnapshotDownloadCreate20241113) (*atlasv2.FlexBackupRestoreJob20241113, error)
}

type DownloadOpts struct {
	cli.ProjectOpts
	cli.DownloaderOpts
	cli.OutputOpts
	store       Downloader
	clusterName string
	id          string
}

func (opts *DownloadOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var downloadTemplate = "Snapshot '%s' downloaded.\n"
var errEmptyURL = errors.New("'snapshotUrl' is empty")
var errExtNotSupported = errors.New("only the '.tgz' extension is supported")

func (opts *DownloadOpts) Run() error {
	r, err := opts.store.DownloadFlexClusterSnapshot(opts.ConfigProjectID(), opts.clusterName, opts.newFlexBackupSnapshotDownloadCreate())
	if err != nil {
		return err
	}

	return opts.Download(r.SnapshotUrl)
}

func (opts *DownloadOpts) newFlexBackupSnapshotDownloadCreate() *atlasv2.FlexBackupSnapshotDownloadCreate20241113 {
	return &atlasv2.FlexBackupSnapshotDownloadCreate20241113{
		SnapshotId: opts.id,
	}
}

func (opts *DownloadOpts) Download(url *string) error {
	if url == nil {
		return errEmptyURL
	}

	w, err := opts.NewWriteCloser()
	if err != nil {
		return err
	}
	defer w.Close()

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, *url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if err != nil {
		_ = opts.OnError(w)
		return err
	}

	if _, err := io.Copy(w, resp.Body); err != nil {
		_ = opts.OnError(w)
		return err
	}

	fmt.Printf(downloadTemplate, opts.Out)
	return nil
}

func (opts *DownloadOpts) initDefaultOut() error {
	if opts.Out == "" {
		opts.Out = opts.id + ".tgz"
	} else if !strings.Contains(opts.Out, ".tgz") {
		return errExtNotSupported
	}

	return nil
}

// DownloadBuilder builds a cobra.Command that can run as:
// atlas backup snapshots download snapshotId --clusterName string [--projectId string] [--out string].
func DownloadBuilder() *cobra.Command {
	opts := &DownloadOpts{}
	opts.Fs = afero.NewOsFs()
	cmd := &cobra.Command{
		Use:   "download <snapshotId>",
		Short: "Download one snapshot for the specified flex cluster.",
		Long: `You can download a snapshot for an Atlas Flex cluster.
` + fmt.Sprintf("%s\n%s", fmt.Sprintf(usage.RequiredRole, "Project Owner"), "Atlas supports this command only for Flex clusters."),
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"snapshotIdDesc": "Unique 24-hexadecimal digit string that identifies the snapshot to download.",
			"output":         downloadTemplate,
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.id = args[0]
			if err := opts.initDefaultOut(); err != nil {
				return err
			}
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().StringVar(&opts.Out, flag.Out, "", usage.SnapshotOut)

	opts.AddProjectOptsFlags(cmd)

	_ = cmd.MarkFlagRequired(flag.ClusterName)

	return cmd
}

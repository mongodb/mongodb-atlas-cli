// Copyright 2023 MongoDB Inc
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

package deployments

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/internal/search"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

type DownloadOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	cli.DownloaderOpts
	options.DeploymentOpts
	downloadStore store.LogsDownloader
	host          string
	name          string
	start         int64
	end           int64
}

var downloadMessage = "Download of %s completed.\n"

var ErrAtlasNotSupported = errors.New("atlas deployments are not supported")

func (opts *DownloadOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.downloadStore, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DownloadOpts) Run(ctx context.Context) error {
	if _, err := opts.SelectDeployments(ctx, opts.ConfigProjectID()); err != nil {
		return err
	}

	if opts.IsLocalDeploymentType() {
		return opts.RunLocal(ctx)
	}

	if opts.IsAtlasDeploymentType() {
		return opts.RunAtlas()
	}

	return errors.New("atlas deployments are not supported")
}

func (opts *DownloadOpts) RunAtlas() error {
	r := opts.newHostLogsParams()
	if err := opts.downloadStore.DownloadLog(opts.OutWriter, r); err != nil {
		return err
	}
	if !opts.ShouldDownloadToStdout() {
		fmt.Printf(downloadMessage, opts.Out)
	}
	return opts.Print(r)
}

func (opts *DownloadOpts) newHostLogsParams() *admin.GetHostLogsApiParams {
	fileBaseName := strings.TrimSuffix(opts.name, filepath.Ext(opts.name))
	params := &admin.GetHostLogsApiParams{
		GroupId:  opts.ConfigProjectID(),
		HostName: opts.host,
		LogName:  fileBaseName,
	}
	if opts.start > 0 {
		params.StartDate = &opts.start
	}
	if opts.end > 0 {
		params.StartDate = &opts.end
	}
	return params
}

func (opts *DownloadOpts) RunLocal(ctx context.Context) error {
	logs, err := opts.PodmanClient.ContainerLogs(ctx, opts.LocalMongodHostname())
	if err != nil {
		return err
	}
	// format log entries into lines
	if opts.IsJSONOutput() {
		return opts.Print(logs)
	}
	return opts.Print(strings.Join(logs, "\n"))
}

func (opts *DownloadOpts) validateAtlasFlags() error {
	if opts.host == "" {
		return errors.New("missing --hostname flag")
	}
	if opts.name == "" {
		return errors.New("missing --name flag")
	}

	validNameFlags := []string{"mongodb.gz", "mongos.gz", "mongosqld.gz", "mongodb-audit-log.gz", "mongos-audit-log.gz"}
	if !search.StringInSliceFold(validNameFlags, opts.name) {
		return fmt.Errorf("invalid --name flag: %s", opts.name)
	}
	return nil
}

// atlas deployments logs.
func LogsBuilder() *cobra.Command {
	opts := &DownloadOpts{}
	cmd := &cobra.Command{
		Use:     "logs",
		Short:   "Get deployments logs.",
		Aliases: []string{"log"},
		Args:    require.NoArgs,
		GroupID: "local",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			w := cmd.OutOrStdout()
			log.SetWriter(w)

			if opts.IsAtlasDeploymentType() {
				if err := opts.validateAtlasFlags(); err != nil {
					return err
				}
			}

			return opts.PreRunE(
				opts.InitStore(cmd.Context(), cmd.OutOrStdout()),
				opts.initStore(cmd.Context()),
				opts.InitOutput(w, ""))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}

	cmd.Flags().StringVar(&opts.DeploymentType, flag.TypeFlag, "", usage.DeploymentType)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	cmd.Flags().StringVar(&opts.DeploymentName, flag.DeploymentName, "", usage.DeploymentName)

	// Atlas flags
	cmd.Flags().Int64Var(&opts.start, flag.Start, 0, usage.LogStart)
	cmd.Flags().Int64Var(&opts.end, flag.End, 0, usage.LogEnd)
	cmd.Flags().BoolVar(&opts.Force, flag.Force, false, usage.ForceFile)
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVar(&opts.host, flag.Hostname, "", usage.LogHostName)
	cmd.Flags().StringVar(&opts.name, flag.Name, "", usage.LogName)

	return cmd
}

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
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/search"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20241113003/admin"
)

type DownloadOpts struct {
	cli.OutputOpts
	cli.ProjectOpts
	cli.DownloaderOpts
	options.DeploymentOpts
	downloadStore store.LogsDownloader
	Host          string
	Name          string
	start         int64
	end           int64
}

var (
	errEmptyLog = errors.New("log is empty")
)

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
		if err := opts.promptMissingAtlasFlags(); err != nil {
			return err
		}
		if err := opts.validateAtlasFlags(); err != nil {
			return err
		}
		return opts.RunAtlas()
	}

	return errors.New("atlas deployments are not supported")
}

func (opts *DownloadOpts) RunAtlas() error {
	if err := opts.downloadLogFile(); err != nil {
		return err
	}
	defer func() {
		_ = opts.Fs.Remove(opts.Out)
	}()

	return nil
}

// maxBytes  1k each write to avoid compression bomb.
const maxBytes = 1024

func (*DownloadOpts) write(w io.Writer, r io.Reader) error {
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

func (opts *DownloadOpts) downloadLogFile() error {
	w, err := opts.NewWriteCloser()
	if err != nil {
		_ = opts.OnError(w)
		return err
	}
	defer w.Close()

	r, err := opts.downloadStore.DownloadLog(opts.newHostLogsParams())
	if err != nil {
		_ = opts.OnError(w)
		return err
	}
	defer r.Close()

	if err := opts.write(w, r); err != nil {
		_ = opts.OnError(w)
		return err
	}

	return nil
}

func (opts *DownloadOpts) newHostLogsParams() *admin.GetHostLogsApiParams {
	fileBaseName := strings.TrimSuffix(opts.Name, filepath.Ext(opts.Name))
	params := &admin.GetHostLogsApiParams{
		GroupId:  opts.ConfigProjectID(),
		HostName: opts.Host,
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
	logs, err := opts.ContainerEngine.ContainerLogs(ctx, opts.LocalMongodHostname())
	if err != nil {
		return err
	}
	// format log entries into lines
	if opts.IsJSONOutput() {
		return opts.Print(logs)
	}
	return opts.Print(strings.Join(logs, "\n"))
}

func (opts *DownloadOpts) promptMissingAtlasFlags() error {
	questions := make([]*survey.Question, 0)

	if opts.Host == "" {
		questions = append(questions, &survey.Question{
			Name: "host",
			Prompt: &survey.Input{
				Message: "Hostname:",
			},
			Validate: survey.Required,
		})
	}

	if opts.Name == "" {
		questions = append(questions, &survey.Question{
			Name: "name",
			Prompt: &survey.Select{
				Message: "Choose a log:",
				Options: []string{
					"mongodb.gz",
					"mongos.gz",
					"mongosqld.gz",
					"mongodb-audit-log.gz",
					"mongos-audit-log.gz",
				},
			},
			Validate: survey.Required,
		})
	}

	if len(questions) > 0 {
		return survey.Ask(questions, opts)
	}

	return nil
}

func (opts *DownloadOpts) validateAtlasFlags() error {
	if opts.Host == "" {
		return errors.New("missing --hostname flag")
	}
	if opts.Name == "" {
		return errors.New("missing --name flag")
	}

	validNameFlags := []string{"mongodb.gz", "mongos.gz", "mongosqld.gz", "mongodb-audit-log.gz", "mongos-audit-log.gz"}
	if !search.StringInSliceFold(validNameFlags, opts.Name) {
		return fmt.Errorf("invalid --name flag: %s", opts.Name)
	}
	return nil
}

func (opts *DownloadOpts) PostRun() {
	opts.DeploymentTelemetry.AppendDeploymentType()
}

// atlas deployments logs.
func LogsBuilder() *cobra.Command {
	opts := &DownloadOpts{
		DownloaderOpts: cli.DownloaderOpts{
			Fs:  afero.NewOsFs(),
			Out: "-", // stdout
		},
	}
	cmd := &cobra.Command{
		Use:     "logs",
		Short:   "Get deployment logs.",
		Aliases: []string{"log"},
		Args:    require.NoArgs,
		GroupID: "all",
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.InitStore(cmd.Context(), cmd.OutOrStdout()),
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), ""))
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return opts.Run(cmd.Context())
		},
		PostRun: func(_ *cobra.Command, _ []string) {
			opts.PostRun()
		},
	}

	cmd.Flags().StringVar(&opts.DeploymentType, flag.TypeFlag, "", usage.DeploymentType)
	opts.AddOutputOptFlags(cmd)
	cmd.Flags().StringVar(&opts.DeploymentName, flag.DeploymentName, "", usage.DeploymentName)

	// Atlas flags
	cmd.Flags().Int64Var(&opts.start, flag.Start, 0, usage.LogStart)
	cmd.Flags().Int64Var(&opts.end, flag.End, 0, usage.LogEnd)
	cmd.Flags().BoolVar(&opts.Force, flag.Force, false, usage.ForceFile)
	opts.AddProjectOptsFlags(cmd)
	cmd.Flags().StringVar(&opts.Host, flag.Hostname, "", usage.LogHostName)
	cmd.Flags().StringVar(&opts.Name, flag.Name, "", usage.LogName)

	return cmd
}

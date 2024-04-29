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

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20231115012/admin"
)

type downloadSharedClusterBackupOpts struct {
	client      *admin.APIClient
	clusterName string
	groupId     string

	filename string
	fs       afero.Fs
	format   string
	tmpl     *template.Template
	resp     *admin.TenantRestore
}

func (opts *downloadSharedClusterBackupOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(config.UserAgent, config.Default()); err != nil {
		return err
	}

	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}
	if opts.groupId == "" {
		return errors.New(`required flag(s) "projectId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.groupId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.groupId)
	}

	if opts.format != "" {
		if opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func (opts *downloadSharedClusterBackupOpts) readData(r io.Reader) (*admin.TenantRestore, error) {
	var out *admin.TenantRestore

	var buf []byte
	var err error
	if opts.filename == "" {
		buf, err = io.ReadAll(r)
	} else {
		if exists, errExists := afero.Exists(opts.fs, opts.filename); !exists || errExists != nil {
			return nil, fmt.Errorf("file not found: %s", opts.filename)
		}
		buf, err = afero.ReadFile(opts.fs, opts.filename)
	}
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(buf, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (opts *downloadSharedClusterBackupOpts) run(ctx context.Context, r io.Reader) error {
	data, errData := opts.readData(r)
	if errData != nil {
		return errData
	}

	params := &admin.DownloadSharedClusterBackupApiParams{
		ClusterName: opts.clusterName,
		GroupId:     opts.groupId,

		TenantRestore: data,
	}

	var err error
	opts.resp, _, err = opts.client.SharedTierSnapshotsApi.DownloadSharedClusterBackupWithParams(ctx, params).Execute()
	return err
}

func (opts *downloadSharedClusterBackupOpts) postRun(_ context.Context, w io.Writer) error {

	prettyJSON, errJson := json.MarshalIndent(opts.resp, "", " ")
	if errJson != nil {
		return errJson
	}

	if opts.format == "" {
		_, err := fmt.Fprintln(w, string(prettyJSON))
		return err
	}

	var parsedJSON interface{}
	if err := json.Unmarshal([]byte(prettyJSON), &parsedJSON); err != nil {
		return err
	}

	return opts.tmpl.Execute(w, parsedJSON)
}

func downloadSharedClusterBackupBuilder() *cobra.Command {
	opts := downloadSharedClusterBackupOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "downloadSharedClusterBackup",
		Short: "Download One M2 or M5 Cluster Snapshot",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin())
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return opts.postRun(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster.`)
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("clusterName")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type getSharedClusterBackupOpts struct {
	client      *admin.APIClient
	groupId     string
	clusterName string
	snapshotId  string
	format      string
	tmpl        *template.Template
	resp        *admin.BackupTenantSnapshot
}

func (opts *getSharedClusterBackupOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(config.UserAgent, config.Default()); err != nil {
		return err
	}

	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}
	if opts.groupId == "" {
		return errors.New(`required flag(s) "projectId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.groupId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.groupId)
	}

	if opts.format != "" {
		if opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func (opts *getSharedClusterBackupOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.GetSharedClusterBackupApiParams{
		GroupId:     opts.groupId,
		ClusterName: opts.clusterName,
		SnapshotId:  opts.snapshotId,
	}

	var err error
	opts.resp, _, err = opts.client.SharedTierSnapshotsApi.GetSharedClusterBackupWithParams(ctx, params).Execute()
	return err
}

func (opts *getSharedClusterBackupOpts) postRun(_ context.Context, w io.Writer) error {

	prettyJSON, errJson := json.MarshalIndent(opts.resp, "", " ")
	if errJson != nil {
		return errJson
	}

	if opts.format == "" {
		_, err := fmt.Fprintln(w, string(prettyJSON))
		return err
	}

	var parsedJSON interface{}
	if err := json.Unmarshal([]byte(prettyJSON), &parsedJSON); err != nil {
		return err
	}

	return opts.tmpl.Execute(w, parsedJSON)
}

func getSharedClusterBackupBuilder() *cobra.Command {
	opts := getSharedClusterBackupOpts{}
	cmd := &cobra.Command{
		Use:   "getSharedClusterBackup",
		Short: "Return One Snapshot for One M2 or M5 Cluster",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin())
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return opts.postRun(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster.`)
	cmd.Flags().StringVar(&opts.snapshotId, "snapshotId", "", `Unique 24-hexadecimal digit string that identifies the desired snapshot.`)

	_ = cmd.MarkFlagRequired("clusterName")
	_ = cmd.MarkFlagRequired("snapshotId")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type listSharedClusterBackupsOpts struct {
	client      *admin.APIClient
	groupId     string
	clusterName string
	format      string
	tmpl        *template.Template
	resp        *admin.PaginatedTenantSnapshot
}

func (opts *listSharedClusterBackupsOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(config.UserAgent, config.Default()); err != nil {
		return err
	}

	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}
	if opts.groupId == "" {
		return errors.New(`required flag(s) "projectId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.groupId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.groupId)
	}

	if opts.format != "" {
		if opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func (opts *listSharedClusterBackupsOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.ListSharedClusterBackupsApiParams{
		GroupId:     opts.groupId,
		ClusterName: opts.clusterName,
	}

	var err error
	opts.resp, _, err = opts.client.SharedTierSnapshotsApi.ListSharedClusterBackupsWithParams(ctx, params).Execute()
	return err
}

func (opts *listSharedClusterBackupsOpts) postRun(_ context.Context, w io.Writer) error {

	prettyJSON, errJson := json.MarshalIndent(opts.resp, "", " ")
	if errJson != nil {
		return errJson
	}

	if opts.format == "" {
		_, err := fmt.Fprintln(w, string(prettyJSON))
		return err
	}

	var parsedJSON interface{}
	if err := json.Unmarshal([]byte(prettyJSON), &parsedJSON); err != nil {
		return err
	}

	return opts.tmpl.Execute(w, parsedJSON)
}

func listSharedClusterBackupsBuilder() *cobra.Command {
	opts := listSharedClusterBackupsOpts{}
	cmd := &cobra.Command{
		Use:   "listSharedClusterBackups",
		Short: "Return All Snapshots for One M2 or M5 Cluster",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin())
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return opts.postRun(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster.`)

	_ = cmd.MarkFlagRequired("clusterName")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

func sharedTierSnapshotsBuilder() *cobra.Command {
	const use = "sharedTierSnapshots"
	cmd := &cobra.Command{
		Use:     use,
		Short:   `Returns and requests to download shared-tier database deployment snapshots.`,
		Aliases: cli.GenerateAliases(use),
	}
	cmd.AddCommand(
		downloadSharedClusterBackupBuilder(),
		getSharedClusterBackupBuilder(),
		listSharedClusterBackupsBuilder(),
	)
	return cmd
}

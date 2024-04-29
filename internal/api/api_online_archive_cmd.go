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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20231115012/admin"
)

type createOnlineArchiveOpts struct {
	client      *admin.APIClient
	groupId     string
	clusterName string

	filename string
	fs       afero.Fs
	format   string
	tmpl     *template.Template
	resp     *admin.BackupOnlineArchive
}

func (opts *createOnlineArchiveOpts) preRun() (err error) {
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

func (opts *createOnlineArchiveOpts) readData(r io.Reader) (*admin.BackupOnlineArchiveCreate, error) {
	var out *admin.BackupOnlineArchiveCreate

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

func (opts *createOnlineArchiveOpts) run(ctx context.Context, r io.Reader) error {
	data, errData := opts.readData(r)
	if errData != nil {
		return errData
	}

	params := &admin.CreateOnlineArchiveApiParams{
		GroupId:     opts.groupId,
		ClusterName: opts.clusterName,

		BackupOnlineArchiveCreate: data,
	}

	var err error
	opts.resp, _, err = opts.client.OnlineArchiveApi.CreateOnlineArchiveWithParams(ctx, params).Execute()
	return err
}

func (opts *createOnlineArchiveOpts) postRun(_ context.Context, w io.Writer) error {

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

func createOnlineArchiveBuilder() *cobra.Command {
	opts := createOnlineArchiveOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createOnlineArchive",
		Short: "Create One Online Archive",
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
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster that contains the collection for which you want to create one online archive.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("clusterName")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type deleteOnlineArchiveOpts struct {
	client      *admin.APIClient
	groupId     string
	archiveId   string
	clusterName string
	format      string
	tmpl        *template.Template
	resp        map[string]interface{}
}

func (opts *deleteOnlineArchiveOpts) preRun() (err error) {
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

func (opts *deleteOnlineArchiveOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.DeleteOnlineArchiveApiParams{
		GroupId:     opts.groupId,
		ArchiveId:   opts.archiveId,
		ClusterName: opts.clusterName,
	}

	var err error
	opts.resp, _, err = opts.client.OnlineArchiveApi.DeleteOnlineArchiveWithParams(ctx, params).Execute()
	return err
}

func (opts *deleteOnlineArchiveOpts) postRun(_ context.Context, w io.Writer) error {

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

func deleteOnlineArchiveBuilder() *cobra.Command {
	opts := deleteOnlineArchiveOpts{}
	cmd := &cobra.Command{
		Use:   "deleteOnlineArchive",
		Short: "Remove One Online Archive",
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
	cmd.Flags().StringVar(&opts.archiveId, "archiveId", "", `Unique 24-hexadecimal digit string that identifies the online archive to delete.`)
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster that contains the collection from which you want to remove an online archive.`)

	_ = cmd.MarkFlagRequired("archiveId")
	_ = cmd.MarkFlagRequired("clusterName")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type downloadOnlineArchiveQueryLogsOpts struct {
	client      *admin.APIClient
	groupId     string
	clusterName string
	startDate   int64
	endDate     int64
	archiveOnly bool
	format      string
	tmpl        *template.Template
	resp        io.ReadCloser
}

func (opts *downloadOnlineArchiveQueryLogsOpts) preRun() (err error) {
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

func (opts *downloadOnlineArchiveQueryLogsOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.DownloadOnlineArchiveQueryLogsApiParams{
		GroupId:     opts.groupId,
		ClusterName: opts.clusterName,
		StartDate:   &opts.startDate,
		EndDate:     &opts.endDate,
		ArchiveOnly: &opts.archiveOnly,
	}

	var err error
	opts.resp, _, err = opts.client.OnlineArchiveApi.DownloadOnlineArchiveQueryLogsWithParams(ctx, params).Execute()
	return err
}

func (opts *downloadOnlineArchiveQueryLogsOpts) postRun(_ context.Context, w io.Writer) error {

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

func downloadOnlineArchiveQueryLogsBuilder() *cobra.Command {
	opts := downloadOnlineArchiveQueryLogsOpts{}
	cmd := &cobra.Command{
		Use:   "downloadOnlineArchiveQueryLogs",
		Short: "Download Online Archive Query Logs",
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
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster that contains the collection for which you want to return the query logs from one online archive.`)
	cmd.Flags().Int64Var(&opts.startDate, "startDate", 0, `Date and time that specifies the starting point for the range of log messages to return. This resource expresses this value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).`)
	cmd.Flags().Int64Var(&opts.endDate, "endDate", 0, `Date and time that specifies the end point for the range of log messages to return. This resource expresses this value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).`)
	cmd.Flags().BoolVar(&opts.archiveOnly, "archiveOnly", false, `Flag that indicates whether to download logs for queries against your online archive only or both your online archive and cluster.`)

	_ = cmd.MarkFlagRequired("clusterName")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type getOnlineArchiveOpts struct {
	client      *admin.APIClient
	groupId     string
	archiveId   string
	clusterName string
	format      string
	tmpl        *template.Template
	resp        *admin.BackupOnlineArchive
}

func (opts *getOnlineArchiveOpts) preRun() (err error) {
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

func (opts *getOnlineArchiveOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.GetOnlineArchiveApiParams{
		GroupId:     opts.groupId,
		ArchiveId:   opts.archiveId,
		ClusterName: opts.clusterName,
	}

	var err error
	opts.resp, _, err = opts.client.OnlineArchiveApi.GetOnlineArchiveWithParams(ctx, params).Execute()
	return err
}

func (opts *getOnlineArchiveOpts) postRun(_ context.Context, w io.Writer) error {

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

func getOnlineArchiveBuilder() *cobra.Command {
	opts := getOnlineArchiveOpts{}
	cmd := &cobra.Command{
		Use:   "getOnlineArchive",
		Short: "Return One Online Archive",
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
	cmd.Flags().StringVar(&opts.archiveId, "archiveId", "", `Unique 24-hexadecimal digit string that identifies the online archive to return.`)
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster that contains the specified collection from which Application created the online archive.`)

	_ = cmd.MarkFlagRequired("archiveId")
	_ = cmd.MarkFlagRequired("clusterName")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type listOnlineArchivesOpts struct {
	client       *admin.APIClient
	groupId      string
	clusterName  string
	includeCount bool
	itemsPerPage int
	pageNum      int
	format       string
	tmpl         *template.Template
	resp         *admin.PaginatedOnlineArchive
}

func (opts *listOnlineArchivesOpts) preRun() (err error) {
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

func (opts *listOnlineArchivesOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.ListOnlineArchivesApiParams{
		GroupId:      opts.groupId,
		ClusterName:  opts.clusterName,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
	}

	var err error
	opts.resp, _, err = opts.client.OnlineArchiveApi.ListOnlineArchivesWithParams(ctx, params).Execute()
	return err
}

func (opts *listOnlineArchivesOpts) postRun(_ context.Context, w io.Writer) error {

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

func listOnlineArchivesBuilder() *cobra.Command {
	opts := listOnlineArchivesOpts{}
	cmd := &cobra.Command{
		Use:   "listOnlineArchives",
		Short: "Return All Online Archives for One Cluster",
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
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster that contains the collection for which you want to return the online archives.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)

	_ = cmd.MarkFlagRequired("clusterName")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type updateOnlineArchiveOpts struct {
	client      *admin.APIClient
	groupId     string
	archiveId   string
	clusterName string

	filename string
	fs       afero.Fs
	format   string
	tmpl     *template.Template
	resp     *admin.BackupOnlineArchive
}

func (opts *updateOnlineArchiveOpts) preRun() (err error) {
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

func (opts *updateOnlineArchiveOpts) readData(r io.Reader) (*admin.BackupOnlineArchive, error) {
	var out *admin.BackupOnlineArchive

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

func (opts *updateOnlineArchiveOpts) run(ctx context.Context, r io.Reader) error {
	data, errData := opts.readData(r)
	if errData != nil {
		return errData
	}

	params := &admin.UpdateOnlineArchiveApiParams{
		GroupId:     opts.groupId,
		ArchiveId:   opts.archiveId,
		ClusterName: opts.clusterName,

		BackupOnlineArchive: data,
	}

	var err error
	opts.resp, _, err = opts.client.OnlineArchiveApi.UpdateOnlineArchiveWithParams(ctx, params).Execute()
	return err
}

func (opts *updateOnlineArchiveOpts) postRun(_ context.Context, w io.Writer) error {

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

func updateOnlineArchiveBuilder() *cobra.Command {
	opts := updateOnlineArchiveOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "updateOnlineArchive",
		Short: "Update One Online Archive",
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
	cmd.Flags().StringVar(&opts.archiveId, "archiveId", "", `Unique 24-hexadecimal digit string that identifies the online archive to update.`)
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster that contains the specified collection from which Application created the online archive.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("archiveId")
	_ = cmd.MarkFlagRequired("clusterName")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

func onlineArchiveBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "onlineArchive",
		Short: `Returns, adds, edits, or removes an online archive.`,
	}
	cmd.AddCommand(
		createOnlineArchiveBuilder(),
		deleteOnlineArchiveBuilder(),
		downloadOnlineArchiveQueryLogsBuilder(),
		getOnlineArchiveBuilder(),
		listOnlineArchivesBuilder(),
		updateOnlineArchiveBuilder(),
	)
	return cmd
}

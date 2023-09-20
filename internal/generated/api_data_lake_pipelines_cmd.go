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

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package generated

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/jsonwriter"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

type createPipelineOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client  *admin.APIClient
	groupId string

	filename string
	fs       afero.Fs
}

func (opts *createPipelineOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *createPipelineOpts) readData() (*admin.IngestionPipeline, error) {
	var out *admin.IngestionPipeline

	var buf []byte
	var err error
	if opts.filename == "" {
		buf, err = io.ReadAll(os.Stdin)
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

func (opts *createPipelineOpts) Run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	params := &admin.CreatePipelineApiParams{
		GroupId: opts.groupId,

		IngestionPipeline: data,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.CreatePipelineWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func createPipelineBuilder() *cobra.Command {
	opts := createPipelineOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createPipeline",
		Short: "Create One Data Lake Pipeline",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type deletePipelineOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	groupId      string
	pipelineName string
}

func (opts *deletePipelineOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *deletePipelineOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.DeletePipelineApiParams{
		GroupId:      opts.groupId,
		PipelineName: opts.pipelineName,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.DeletePipelineWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func deletePipelineBuilder() *cobra.Command {
	opts := deletePipelineOpts{}
	cmd := &cobra.Command{
		Use:   "deletePipeline",
		Short: "Remove One Data Lake Pipeline",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", "", `Human-readable label that identifies the Data Lake Pipeline.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("pipelineName")
	return cmd
}

type deletePipelineRunDatasetOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client        *admin.APIClient
	groupId       string
	pipelineName  string
	pipelineRunId string
}

func (opts *deletePipelineRunDatasetOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *deletePipelineRunDatasetOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.DeletePipelineRunDatasetApiParams{
		GroupId:       opts.groupId,
		PipelineName:  opts.pipelineName,
		PipelineRunId: opts.pipelineRunId,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.DeletePipelineRunDatasetWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func deletePipelineRunDatasetBuilder() *cobra.Command {
	opts := deletePipelineRunDatasetOpts{}
	cmd := &cobra.Command{
		Use:   "deletePipelineRunDataset",
		Short: "Delete Pipeline Run Dataset",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", "", `Human-readable label that identifies the Data Lake Pipeline.`)
	cmd.Flags().StringVar(&opts.pipelineRunId, "pipelineRunId", "", `Unique 24-hexadecimal character string that identifies a Data Lake Pipeline run.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("pipelineName")
	_ = cmd.MarkFlagRequired("pipelineRunId")
	return cmd
}

type getPipelineOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	groupId      string
	pipelineName string
}

func (opts *getPipelineOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getPipelineOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.GetPipelineApiParams{
		GroupId:      opts.groupId,
		PipelineName: opts.pipelineName,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.GetPipelineWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func getPipelineBuilder() *cobra.Command {
	opts := getPipelineOpts{}
	cmd := &cobra.Command{
		Use:   "getPipeline",
		Short: "Return One Data Lake Pipeline",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", "", `Human-readable label that identifies the Data Lake Pipeline.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("pipelineName")
	return cmd
}

type getPipelineRunOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client        *admin.APIClient
	groupId       string
	pipelineName  string
	pipelineRunId string
}

func (opts *getPipelineRunOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getPipelineRunOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.GetPipelineRunApiParams{
		GroupId:       opts.groupId,
		PipelineName:  opts.pipelineName,
		PipelineRunId: opts.pipelineRunId,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.GetPipelineRunWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func getPipelineRunBuilder() *cobra.Command {
	opts := getPipelineRunOpts{}
	cmd := &cobra.Command{
		Use:   "getPipelineRun",
		Short: "Return One Data Lake Pipeline Run",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", "", `Human-readable label that identifies the Data Lake Pipeline.`)
	cmd.Flags().StringVar(&opts.pipelineRunId, "pipelineRunId", "", `Unique 24-hexadecimal character string that identifies a Data Lake Pipeline run.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("pipelineName")
	_ = cmd.MarkFlagRequired("pipelineRunId")
	return cmd
}

type listPipelineRunsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client        *admin.APIClient
	groupId       string
	pipelineName  string
	includeCount  bool
	itemsPerPage  int
	pageNum       int
	createdBefore string
}

func (opts *listPipelineRunsOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *listPipelineRunsOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.ListPipelineRunsApiParams{
		GroupId:       opts.groupId,
		PipelineName:  opts.pipelineName,
		IncludeCount:  &opts.includeCount,
		ItemsPerPage:  &opts.itemsPerPage,
		PageNum:       &opts.pageNum,
		CreatedBefore: convertTime(&opts.createdBefore),
	}
	resp, _, err := opts.client.DataLakePipelinesApi.ListPipelineRunsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func listPipelineRunsBuilder() *cobra.Command {
	opts := listPipelineRunsOpts{}
	cmd := &cobra.Command{
		Use:   "listPipelineRuns",
		Short: "Return All Data Lake Pipeline Runs from One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", "", `Human-readable label that identifies the Data Lake Pipeline.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)
	cmd.Flags().StringVar(&opts.createdBefore, "createdBefore", "", `If specified, Atlas returns only Data Lake Pipeline runs initiated before this time and date.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("pipelineName")
	return cmd
}

type listPipelineSchedulesOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	groupId      string
	pipelineName string
}

func (opts *listPipelineSchedulesOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *listPipelineSchedulesOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.ListPipelineSchedulesApiParams{
		GroupId:      opts.groupId,
		PipelineName: opts.pipelineName,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.ListPipelineSchedulesWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func listPipelineSchedulesBuilder() *cobra.Command {
	opts := listPipelineSchedulesOpts{}
	cmd := &cobra.Command{
		Use:   "listPipelineSchedules",
		Short: "Return Available Ingestion Schedules for One Data Lake Pipeline",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", "", `Human-readable label that identifies the Data Lake Pipeline.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("pipelineName")
	return cmd
}

type listPipelineSnapshotsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client         *admin.APIClient
	groupId        string
	pipelineName   string
	includeCount   bool
	itemsPerPage   int
	pageNum        int
	completedAfter string
}

func (opts *listPipelineSnapshotsOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *listPipelineSnapshotsOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.ListPipelineSnapshotsApiParams{
		GroupId:        opts.groupId,
		PipelineName:   opts.pipelineName,
		IncludeCount:   &opts.includeCount,
		ItemsPerPage:   &opts.itemsPerPage,
		PageNum:        &opts.pageNum,
		CompletedAfter: convertTime(&opts.completedAfter),
	}
	resp, _, err := opts.client.DataLakePipelinesApi.ListPipelineSnapshotsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func listPipelineSnapshotsBuilder() *cobra.Command {
	opts := listPipelineSnapshotsOpts{}
	cmd := &cobra.Command{
		Use:   "listPipelineSnapshots",
		Short: "Return Available Backup Snapshots for One Data Lake Pipeline",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", "", `Human-readable label that identifies the Data Lake Pipeline.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)
	cmd.Flags().StringVar(&opts.completedAfter, "completedAfter", "", `Date and time after which MongoDB Cloud created the snapshot. If specified, MongoDB Cloud returns available backup snapshots created after this time and date only. This parameter expresses its value in the ISO 8601 timestamp format in UTC.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("pipelineName")
	return cmd
}

type listPipelinesOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client  *admin.APIClient
	groupId string
}

func (opts *listPipelinesOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *listPipelinesOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.ListPipelinesApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.ListPipelinesWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func listPipelinesBuilder() *cobra.Command {
	opts := listPipelinesOpts{}
	cmd := &cobra.Command{
		Use:   "listPipelines",
		Short: "Return All Data Lake Pipelines from One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type pausePipelineOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	groupId      string
	pipelineName string
}

func (opts *pausePipelineOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *pausePipelineOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.PausePipelineApiParams{
		GroupId:      opts.groupId,
		PipelineName: opts.pipelineName,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.PausePipelineWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func pausePipelineBuilder() *cobra.Command {
	opts := pausePipelineOpts{}
	cmd := &cobra.Command{
		Use:   "pausePipeline",
		Short: "Pause One Data Lake Pipeline",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", "", `Human-readable label that identifies the Data Lake Pipeline.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("pipelineName")
	return cmd
}

type resumePipelineOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	groupId      string
	pipelineName string
}

func (opts *resumePipelineOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *resumePipelineOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.ResumePipelineApiParams{
		GroupId:      opts.groupId,
		PipelineName: opts.pipelineName,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.ResumePipelineWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func resumePipelineBuilder() *cobra.Command {
	opts := resumePipelineOpts{}
	cmd := &cobra.Command{
		Use:   "resumePipeline",
		Short: "Resume One Data Lake Pipeline",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", "", `Human-readable label that identifies the Data Lake Pipeline.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("pipelineName")
	return cmd
}

type triggerSnapshotIngestionOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	groupId      string
	pipelineName string

	filename string
	fs       afero.Fs
}

func (opts *triggerSnapshotIngestionOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *triggerSnapshotIngestionOpts) readData() (*admin.TriggerIngestionRequest, error) {
	var out *admin.TriggerIngestionRequest

	var buf []byte
	var err error
	if opts.filename == "" {
		buf, err = io.ReadAll(os.Stdin)
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

func (opts *triggerSnapshotIngestionOpts) Run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	params := &admin.TriggerSnapshotIngestionApiParams{
		GroupId:      opts.groupId,
		PipelineName: opts.pipelineName,

		TriggerIngestionRequest: data,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.TriggerSnapshotIngestionWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func triggerSnapshotIngestionBuilder() *cobra.Command {
	opts := triggerSnapshotIngestionOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "triggerSnapshotIngestion",
		Short: "Trigger on demand snapshot ingestion",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", "", `Human-readable label that identifies the Data Lake Pipeline.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("pipelineName")
	return cmd
}

type updatePipelineOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	groupId      string
	pipelineName string

	filename string
	fs       afero.Fs
}

func (opts *updatePipelineOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *updatePipelineOpts) readData() (*admin.IngestionPipeline, error) {
	var out *admin.IngestionPipeline

	var buf []byte
	var err error
	if opts.filename == "" {
		buf, err = io.ReadAll(os.Stdin)
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

func (opts *updatePipelineOpts) Run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	params := &admin.UpdatePipelineApiParams{
		GroupId:      opts.groupId,
		PipelineName: opts.pipelineName,

		IngestionPipeline: data,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.UpdatePipelineWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func updatePipelineBuilder() *cobra.Command {
	opts := updatePipelineOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "updatePipeline",
		Short: "Update One Data Lake Pipeline",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", "", `Human-readable label that identifies the Data Lake Pipeline.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("pipelineName")
	return cmd
}

func dataLakePipelinesBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dataLakePipelines",
		Short: `Returns, adds, edits, and removes Atlas Data Lake Pipelines and associated runs.`,
	}
	cmd.AddCommand(
		createPipelineBuilder(),
		deletePipelineBuilder(),
		deletePipelineRunDatasetBuilder(),
		getPipelineBuilder(),
		getPipelineRunBuilder(),
		listPipelineRunsBuilder(),
		listPipelineSchedulesBuilder(),
		listPipelineSnapshotsBuilder(),
		listPipelinesBuilder(),
		pausePipelineBuilder(),
		resumePipelineBuilder(),
		triggerSnapshotIngestionBuilder(),
		updatePipelineBuilder(),
	)
	return cmd
}

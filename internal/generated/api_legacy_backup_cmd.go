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
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/admin"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
)

type DeleteLegacySnapshotOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
	snapshotId string
}

func (opts *DeleteLegacySnapshotOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DeleteLegacySnapshotOpts) Run(ctx context.Context) error {
	params := &admin.DeleteLegacySnapshotApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
		SnapshotId: opts.snapshotId,
	}
	resp, _, err := opts.client.LegacyBackupApi.DeleteLegacySnapshotWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func DeleteLegacySnapshotBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := DeleteLegacySnapshotOpts{}
	cmd := &cobra.Command{
		Use:     "deleteLegacySnapshot",
		// Aliases: []string{"?"},
		Short:   "Remove One Legacy Backup Snapshot",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", , "Human-readable label that identifies the cluster.")
	cmd.Flags().StringVar(&opts.snapshotId, "snapshotId", , "Unique 24-hexadecimal digit string that identifies the desired snapshot.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	_ = cmd.MarkFlagRequired("snapshotId")

	return cmd
}
type GetLegacyBackupCheckpointOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	checkpointId string
	clusterName string
}

func (opts *GetLegacyBackupCheckpointOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetLegacyBackupCheckpointOpts) Run(ctx context.Context) error {
	params := &admin.GetLegacyBackupCheckpointApiParams{
		GroupId: opts.groupId,
		CheckpointId: opts.checkpointId,
		ClusterName: opts.clusterName,
	}
	resp, _, err := opts.client.LegacyBackupApi.GetLegacyBackupCheckpointWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetLegacyBackupCheckpointBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetLegacyBackupCheckpointOpts{}
	cmd := &cobra.Command{
		Use:     "getLegacyBackupCheckpoint",
		// Aliases: []string{"?"},
		Short:   "Return One Legacy Backup Checkpoint",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	cmd.Flags().StringVar(&opts.checkpointId, "checkpointId", , "Unique 24-hexadecimal digit string that identifies the checkpoint.")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", , "Human-readable label that identifies the cluster that contains the checkpoints that you want to return.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("checkpointId")
	_ = cmd.MarkFlagRequired("clusterName")

	return cmd
}
type GetLegacyBackupRestoreJobOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
	jobId string
}

func (opts *GetLegacyBackupRestoreJobOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetLegacyBackupRestoreJobOpts) Run(ctx context.Context) error {
	params := &admin.GetLegacyBackupRestoreJobApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
		JobId: opts.jobId,
	}
	resp, _, err := opts.client.LegacyBackupApi.GetLegacyBackupRestoreJobWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetLegacyBackupRestoreJobBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetLegacyBackupRestoreJobOpts{}
	cmd := &cobra.Command{
		Use:     "getLegacyBackupRestoreJob",
		// Aliases: []string{"?"},
		Short:   "Return One Legacy Backup Restore Job",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", , "Human-readable label that identifies the cluster with the snapshot you want to return.")
	cmd.Flags().StringVar(&opts.jobId, "jobId", , "Unique 24-hexadecimal digit string that identifies the restore job.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	_ = cmd.MarkFlagRequired("jobId")

	return cmd
}
type GetLegacySnapshotOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
	snapshotId string
}

func (opts *GetLegacySnapshotOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetLegacySnapshotOpts) Run(ctx context.Context) error {
	params := &admin.GetLegacySnapshotApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
		SnapshotId: opts.snapshotId,
	}
	resp, _, err := opts.client.LegacyBackupApi.GetLegacySnapshotWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetLegacySnapshotBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetLegacySnapshotOpts{}
	cmd := &cobra.Command{
		Use:     "getLegacySnapshot",
		// Aliases: []string{"?"},
		Short:   "Return One Legacy Backup Snapshot",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", , "Human-readable label that identifies the cluster.")
	cmd.Flags().StringVar(&opts.snapshotId, "snapshotId", , "Unique 24-hexadecimal digit string that identifies the desired snapshot.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	_ = cmd.MarkFlagRequired("snapshotId")

	return cmd
}
type GetLegacySnapshotScheduleOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
}

func (opts *GetLegacySnapshotScheduleOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetLegacySnapshotScheduleOpts) Run(ctx context.Context) error {
	params := &admin.GetLegacySnapshotScheduleApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
	}
	resp, _, err := opts.client.LegacyBackupApi.GetLegacySnapshotScheduleWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetLegacySnapshotScheduleBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetLegacySnapshotScheduleOpts{}
	cmd := &cobra.Command{
		Use:     "getLegacySnapshotSchedule",
		// Aliases: []string{"?"},
		Short:   "Return One Snapshot Schedule",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", , "Human-readable label that identifies the cluster with the snapshot you want to return.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")

	return cmd
}
type ListLegacyBackupCheckpointsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
	includeCount bool
	itemsPerPage int
	pageNum int
}

func (opts *ListLegacyBackupCheckpointsOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListLegacyBackupCheckpointsOpts) Run(ctx context.Context) error {
	params := &admin.ListLegacyBackupCheckpointsApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
	}
	resp, _, err := opts.client.LegacyBackupApi.ListLegacyBackupCheckpointsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListLegacyBackupCheckpointsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListLegacyBackupCheckpointsOpts{}
	cmd := &cobra.Command{
		Use:     "listLegacyBackupCheckpoints",
		// Aliases: []string{"?"},
		Short:   "Return All Legacy Backup Checkpoints",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", , "Human-readable label that identifies the cluster that contains the checkpoints that you want to return.")
	cmd.Flags().StringVar(&opts.includeCount, "includeCount", true, "Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.")
	cmd.Flags().StringVar(&opts.itemsPerPage, "itemsPerPage", 100, "Number of items that the response returns per page.")
	cmd.Flags().StringVar(&opts.pageNum, "pageNum", 1, "Number of the page that displays the current set of the total objects that the response returns.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")

	return cmd
}
type ListLegacyBackupRestoreJobsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
	includeCount bool
	itemsPerPage int
	pageNum int
	batchId string
}

func (opts *ListLegacyBackupRestoreJobsOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListLegacyBackupRestoreJobsOpts) Run(ctx context.Context) error {
	params := &admin.ListLegacyBackupRestoreJobsApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
		BatchId: opts.batchId,
	}
	resp, _, err := opts.client.LegacyBackupApi.ListLegacyBackupRestoreJobsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListLegacyBackupRestoreJobsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListLegacyBackupRestoreJobsOpts{}
	cmd := &cobra.Command{
		Use:     "listLegacyBackupRestoreJobs",
		// Aliases: []string{"?"},
		Short:   "Return All Legacy Backup Restore Jobs",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", , "Human-readable label that identifies the cluster with the snapshot you want to return.")
	cmd.Flags().StringVar(&opts.includeCount, "includeCount", true, "Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.")
	cmd.Flags().StringVar(&opts.itemsPerPage, "itemsPerPage", 100, "Number of items that the response returns per page.")
	cmd.Flags().StringVar(&opts.pageNum, "pageNum", 1, "Number of the page that displays the current set of the total objects that the response returns.")
	cmd.Flags().StringVar(&opts.batchId, "batchId", , "Unique 24-hexadecimal digit string that identifies the batch of restore jobs to return. Timestamp in ISO 8601 date and time format in UTC when creating a restore job for a sharded cluster, Application creates a separate job for each shard, plus another for the config host. Each of these jobs comprise one batch. A restore job for a replica set can&#39;t be part of a batch.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")

	return cmd
}
type ListLegacySnapshotsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
	includeCount bool
	itemsPerPage int
	pageNum int
	completed string
}

func (opts *ListLegacySnapshotsOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListLegacySnapshotsOpts) Run(ctx context.Context) error {
	params := &admin.ListLegacySnapshotsApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
		Completed: opts.completed,
	}
	resp, _, err := opts.client.LegacyBackupApi.ListLegacySnapshotsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListLegacySnapshotsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListLegacySnapshotsOpts{}
	cmd := &cobra.Command{
		Use:     "listLegacySnapshots",
		// Aliases: []string{"?"},
		Short:   "Return All Legacy Backup Snapshots",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", , "Human-readable label that identifies the cluster.")
	cmd.Flags().StringVar(&opts.includeCount, "includeCount", true, "Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.")
	cmd.Flags().StringVar(&opts.itemsPerPage, "itemsPerPage", 100, "Number of items that the response returns per page.")
	cmd.Flags().StringVar(&opts.pageNum, "pageNum", 1, "Number of the page that displays the current set of the total objects that the response returns.")
	cmd.Flags().StringVar(&opts.completed, "completed", &quot;true&quot;, "Human-readable label that specifies whether to return only completed, incomplete, or all snapshots. By default, MongoDB Cloud only returns completed snapshots.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")

	return cmd
}
type UpdateLegacySnapshotRetentionOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
	snapshotId string
}

func (opts *UpdateLegacySnapshotRetentionOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *UpdateLegacySnapshotRetentionOpts) Run(ctx context.Context) error {
	params := &admin.UpdateLegacySnapshotRetentionApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
		SnapshotId: opts.snapshotId,
	}
	resp, _, err := opts.client.LegacyBackupApi.UpdateLegacySnapshotRetentionWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func UpdateLegacySnapshotRetentionBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := UpdateLegacySnapshotRetentionOpts{}
	cmd := &cobra.Command{
		Use:     "updateLegacySnapshotRetention",
		// Aliases: []string{"?"},
		Short:   "Change One Legacy Backup Snapshot Expiration",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", , "Human-readable label that identifies the cluster.")
	cmd.Flags().StringVar(&opts.snapshotId, "snapshotId", , "Unique 24-hexadecimal digit string that identifies the desired snapshot.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	_ = cmd.MarkFlagRequired("snapshotId")

	return cmd
}
type UpdateLegacySnapshotScheduleOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
}

func (opts *UpdateLegacySnapshotScheduleOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *UpdateLegacySnapshotScheduleOpts) Run(ctx context.Context) error {
	params := &admin.UpdateLegacySnapshotScheduleApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
	}
	resp, _, err := opts.client.LegacyBackupApi.UpdateLegacySnapshotScheduleWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func UpdateLegacySnapshotScheduleBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := UpdateLegacySnapshotScheduleOpts{}
	cmd := &cobra.Command{
		Use:     "updateLegacySnapshotSchedule",
		// Aliases: []string{"?"},
		Short:   "Update Snapshot Schedule for One Cluster",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", , "Human-readable label that identifies the cluster with the snapshot you want to return.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")

	return cmd
}

func LegacyBackupBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "legacyBackup",
		Short:   "Manages Legacy Backup snapshots, restore jobs, schedules and checkpoints.",
	}
	cmd.AddCommand(
		DeleteLegacySnapshotBuilder(),
		GetLegacyBackupCheckpointBuilder(),
		GetLegacyBackupRestoreJobBuilder(),
		GetLegacySnapshotBuilder(),
		GetLegacySnapshotScheduleBuilder(),
		ListLegacyBackupCheckpointsBuilder(),
		ListLegacyBackupRestoreJobsBuilder(),
		ListLegacySnapshotsBuilder(),
		UpdateLegacySnapshotRetentionBuilder(),
		UpdateLegacySnapshotScheduleBuilder(),
	)
	return cmd
}

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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

type createSharedClusterBackupRestoreJobOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client      *admin.APIClient
	clusterName string
	groupId     string
}

func (opts *createSharedClusterBackupRestoreJobOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *createSharedClusterBackupRestoreJobOpts) Run(ctx context.Context) error {
	params := &admin.CreateSharedClusterBackupRestoreJobApiParams{
		ClusterName: opts.clusterName,
		GroupId:     opts.groupId,
	}
	resp, _, err := opts.client.SharedTierRestoreJobsApi.CreateSharedClusterBackupRestoreJobWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func createSharedClusterBackupRestoreJobBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := createSharedClusterBackupRestoreJobOpts{}
	cmd := &cobra.Command{
		Use:   "createSharedClusterBackupRestoreJob",
		Short: "Create One Restore Job from One M2 or M5 Cluster",
		Annotations: map[string]string{
			"output": template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster.`)
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)

	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the source cluster.`)

	cmd.Flags().StringVar(&opts.deliveryType, "deliveryType", "", `Means by which this resource returns the snapshot to the requesting MongoDB Cloud user.`)

	cmd.Flags().StringVar(&opts.expirationDate, "expirationDate", "", `Date and time when the download link no longer works. This parameter expresses its value in the ISO 8601 timestamp format in UTC.`)

	cmd.Flags().StringVar(&opts.id, "id", "", `Unique 24-hexadecimal digit string that identifies the restore job.`)

	cmd.Flags().ArraySliceVar(&opts.links, "links", nil, `List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.`)

	cmd.Flags().StringVar(&opts.projectId, "projectId", "", `Unique 24-hexadecimal digit string that identifies the project from which the restore job originated.`)

	cmd.Flags().StringVar(&opts.restoreFinishedDate, "restoreFinishedDate", "", `Date and time when MongoDB Cloud completed writing this snapshot. MongoDB Cloud changes the status of the restore job to &#x60;CLOSED&#x60;. This parameter expresses its value in the ISO 8601 timestamp format in UTC.`)

	cmd.Flags().StringVar(&opts.restoreScheduledDate, "restoreScheduledDate", "", `Date and time when MongoDB Cloud will restore this snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC.`)

	cmd.Flags().StringVar(&opts.snapshotFinishedDate, "snapshotFinishedDate", "", `Date and time when MongoDB Cloud completed writing this snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC.`)

	cmd.Flags().StringVar(&opts.snapshotId, "snapshotId", "", `Unique 24-hexadecimal digit string that identifies the snapshot to restore.`)

	cmd.Flags().StringVar(&opts.snapshotUrl, "snapshotUrl", "", `Internet address from which you can download the compressed snapshot files. The resource returns this parameter when  &#x60;&quot;deliveryType&quot; : &quot;DOWNLOAD&quot;&#x60;.`)

	cmd.Flags().StringVar(&opts.status, "status", "", `Phase of the restore workflow for this job at the time this resource made this request.`)

	cmd.Flags().StringVar(&opts.targetDeploymentItemName, "targetDeploymentItemName", "", `Human-readable label that identifies the cluster on the target project to which you want to restore the snapshot. You can restore the snapshot to a cluster tier *M2* or greater.`)

	cmd.Flags().StringVar(&opts.targetProjectId, "targetProjectId", "", `Unique 24-hexadecimal digit string that identifies the project that contains the cluster to which you want to restore the snapshot.`)

	_ = cmd.MarkFlagRequired("clusterName")
	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type getSharedClusterBackupRestoreJobOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client      *admin.APIClient
	clusterName string
	groupId     string
	restoreId   string
}

func (opts *getSharedClusterBackupRestoreJobOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getSharedClusterBackupRestoreJobOpts) Run(ctx context.Context) error {
	params := &admin.GetSharedClusterBackupRestoreJobApiParams{
		ClusterName: opts.clusterName,
		GroupId:     opts.groupId,
		RestoreId:   opts.restoreId,
	}
	resp, _, err := opts.client.SharedTierRestoreJobsApi.GetSharedClusterBackupRestoreJobWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func getSharedClusterBackupRestoreJobBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := getSharedClusterBackupRestoreJobOpts{}
	cmd := &cobra.Command{
		Use:   "getSharedClusterBackupRestoreJob",
		Short: "Return One Restore Job for One M2 or M5 Cluster",
		Annotations: map[string]string{
			"output": template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster.`)
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.restoreId, "restoreId", "", `Unique 24-hexadecimal digit string that identifies the restore job to return.`)

	_ = cmd.MarkFlagRequired("clusterName")
	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("restoreId")
	return cmd
}

type listSharedClusterBackupRestoreJobsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client      *admin.APIClient
	clusterName string
	groupId     string
}

func (opts *listSharedClusterBackupRestoreJobsOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *listSharedClusterBackupRestoreJobsOpts) Run(ctx context.Context) error {
	params := &admin.ListSharedClusterBackupRestoreJobsApiParams{
		ClusterName: opts.clusterName,
		GroupId:     opts.groupId,
	}
	resp, _, err := opts.client.SharedTierRestoreJobsApi.ListSharedClusterBackupRestoreJobsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func listSharedClusterBackupRestoreJobsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := listSharedClusterBackupRestoreJobsOpts{}
	cmd := &cobra.Command{
		Use:   "listSharedClusterBackupRestoreJobs",
		Short: "Return All Restore Jobs for One M2 or M5 Cluster",
		Annotations: map[string]string{
			"output": template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster.`)
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)

	_ = cmd.MarkFlagRequired("clusterName")
	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

func sharedTierRestoreJobsBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sharedTierRestoreJobs",
		Short: `Returns and adds restore jobs for shared-tier database deployments.`,
	}
	cmd.AddCommand(
		createSharedClusterBackupRestoreJobBuilder(),
		getSharedClusterBackupRestoreJobBuilder(),
		listSharedClusterBackupRestoreJobsBuilder(),
	)
	return cmd
}

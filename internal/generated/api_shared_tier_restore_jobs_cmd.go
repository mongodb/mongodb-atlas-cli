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

type createSharedClusterBackupRestoreJobOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client      *admin.APIClient
	clusterName string
	groupId     string

	filename string
	fs       afero.Fs
}

func (opts *createSharedClusterBackupRestoreJobOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *createSharedClusterBackupRestoreJobOpts) readData() (*admin.TenantRestore, error) {
	var out *admin.TenantRestore

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

func (opts *createSharedClusterBackupRestoreJobOpts) Run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	params := &admin.CreateSharedClusterBackupRestoreJobApiParams{
		ClusterName: opts.clusterName,
		GroupId:     opts.groupId,

		TenantRestore: data,
	}
	resp, _, err := opts.client.SharedTierRestoreJobsApi.CreateSharedClusterBackupRestoreJobWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func createSharedClusterBackupRestoreJobBuilder() *cobra.Command {
	opts := createSharedClusterBackupRestoreJobOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createSharedClusterBackupRestoreJob",
		Short: "Create One Restore Job from One M2 or M5 Cluster",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster.`)
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

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

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

func (opts *getSharedClusterBackupRestoreJobOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.GetSharedClusterBackupRestoreJobApiParams{
		ClusterName: opts.clusterName,
		GroupId:     opts.groupId,
		RestoreId:   opts.restoreId,
	}
	resp, _, err := opts.client.SharedTierRestoreJobsApi.GetSharedClusterBackupRestoreJobWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func getSharedClusterBackupRestoreJobBuilder() *cobra.Command {
	opts := getSharedClusterBackupRestoreJobOpts{}
	cmd := &cobra.Command{
		Use:   "getSharedClusterBackupRestoreJob",
		Short: "Return One Restore Job for One M2 or M5 Cluster",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster.`)
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.restoreId, "restoreId", "", `Unique 24-hexadecimal digit string that identifies the restore job to return.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

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

func (opts *listSharedClusterBackupRestoreJobsOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.ListSharedClusterBackupRestoreJobsApiParams{
		ClusterName: opts.clusterName,
		GroupId:     opts.groupId,
	}
	resp, _, err := opts.client.SharedTierRestoreJobsApi.ListSharedClusterBackupRestoreJobsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func listSharedClusterBackupRestoreJobsBuilder() *cobra.Command {
	opts := listSharedClusterBackupRestoreJobsOpts{}
	cmd := &cobra.Command{
		Use:   "listSharedClusterBackupRestoreJobs",
		Short: "Return All Restore Jobs for One M2 or M5 Cluster",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster.`)
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

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

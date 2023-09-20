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
	"io"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/jsonwriter"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

type listAccessLogsByClusterNameOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client      *admin.APIClient
	groupId     string
	clusterName string
	authResult  bool
	end         int64
	ipAddress   string
	nLogs       int
	start       int64
}

func (opts *listAccessLogsByClusterNameOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *listAccessLogsByClusterNameOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.ListAccessLogsByClusterNameApiParams{
		GroupId:     opts.groupId,
		ClusterName: opts.clusterName,
		AuthResult:  &opts.authResult,
		End:         &opts.end,
		IpAddress:   &opts.ipAddress,
		NLogs:       &opts.nLogs,
		Start:       &opts.start,
	}
	resp, _, err := opts.client.AccessTrackingApi.ListAccessLogsByClusterNameWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func listAccessLogsByClusterNameBuilder() *cobra.Command {
	opts := listAccessLogsByClusterNameOpts{}
	cmd := &cobra.Command{
		Use:   "listAccessLogsByClusterName",
		Short: "Return Database Access History for One Cluster using Its Cluster Name",
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
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster.`)
	cmd.Flags().BoolVar(&opts.authResult, "authResult", false, `Flag that indicates whether the response returns the successful authentication attempts only.`)
	cmd.Flags().Int64Var(&opts.end, "end", 0, `Date and time when to stop retrieving database history. If you specify **end**, you must also specify **start**. This parameter uses UNIX epoch time in milliseconds.`)
	cmd.Flags().StringVar(&opts.ipAddress, "ipAddress", "", `One Internet Protocol address that attempted to authenticate with the database.`)
	cmd.Flags().IntVar(&opts.nLogs, "nLogs", 20000, `Maximum number of lines from the log to return.`)
	cmd.Flags().Int64Var(&opts.start, "start", 0, `Date and time when MongoDB Cloud begins retrieving database history. If you specify **start**, you must also specify **end**. This parameter uses UNIX epoch time in milliseconds.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}

type listAccessLogsByHostnameOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client     *admin.APIClient
	groupId    string
	hostname   string
	authResult bool
	end        int64
	ipAddress  string
	nLogs      int
	start      int64
}

func (opts *listAccessLogsByHostnameOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *listAccessLogsByHostnameOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.ListAccessLogsByHostnameApiParams{
		GroupId:    opts.groupId,
		Hostname:   opts.hostname,
		AuthResult: &opts.authResult,
		End:        &opts.end,
		IpAddress:  &opts.ipAddress,
		NLogs:      &opts.nLogs,
		Start:      &opts.start,
	}
	resp, _, err := opts.client.AccessTrackingApi.ListAccessLogsByHostnameWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func listAccessLogsByHostnameBuilder() *cobra.Command {
	opts := listAccessLogsByHostnameOpts{}
	cmd := &cobra.Command{
		Use:   "listAccessLogsByHostname",
		Short: "Return Database Access History for One Cluster using Its Hostname",
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
	cmd.Flags().StringVar(&opts.hostname, "hostname", "", `Fully qualified domain name or IP address of the MongoDB host that stores the log files that you want to download.`)
	cmd.Flags().BoolVar(&opts.authResult, "authResult", false, `Flag that indicates whether the response returns the successful authentication attempts only.`)
	cmd.Flags().Int64Var(&opts.end, "end", 0, `Date and time when to stop retrieving database history. If you specify **end**, you must also specify **start**. This parameter uses UNIX epoch time in milliseconds.`)
	cmd.Flags().StringVar(&opts.ipAddress, "ipAddress", "", `One Internet Protocol address that attempted to authenticate with the database.`)
	cmd.Flags().IntVar(&opts.nLogs, "nLogs", 20000, `Maximum number of lines from the log to return.`)
	cmd.Flags().Int64Var(&opts.start, "start", 0, `Date and time when MongoDB Cloud begins retrieving database history. If you specify **start**, you must also specify **end**. This parameter uses UNIX epoch time in milliseconds.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("hostname")
	return cmd
}

func accessTrackingBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accessTracking",
		Short: `Returns access logs for authentication attempts made to Atlas database deployments. To view database access history, you must have either the Project Owner or Organization Owner role.`,
	}
	cmd.AddCommand(
		listAccessLogsByClusterNameBuilder(),
		listAccessLogsByHostnameBuilder(),
	)
	return cmd
}

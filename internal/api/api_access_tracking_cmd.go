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
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20231115012/admin"
)

type listAccessLogsByClusterNameOpts struct {
	client      *admin.APIClient
	groupId     string
	clusterName string
	authResult  bool
	end         int64
	ipAddress   string
	nLogs       int
	start       int64
	format      string
	tmpl        *template.Template
	resp        *admin.MongoDBAccessLogsList
}

func (opts *listAccessLogsByClusterNameOpts) preRun() (err error) {
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

func (opts *listAccessLogsByClusterNameOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.ListAccessLogsByClusterNameApiParams{
		GroupId:     opts.groupId,
		ClusterName: opts.clusterName,
		AuthResult:  &opts.authResult,
		End:         &opts.end,
		IpAddress:   &opts.ipAddress,
		NLogs:       &opts.nLogs,
		Start:       &opts.start,
	}

	var err error
	opts.resp, _, err = opts.client.AccessTrackingApi.ListAccessLogsByClusterNameWithParams(ctx, params).Execute()
	return err
}

func (opts *listAccessLogsByClusterNameOpts) postRun(_ context.Context, w io.Writer) error {

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

func listAccessLogsByClusterNameBuilder() *cobra.Command {
	opts := listAccessLogsByClusterNameOpts{}
	cmd := &cobra.Command{
		Use:   "listAccessLogsByClusterName",
		Short: "Return Database Access History for One Cluster using Its Cluster Name",
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
	cmd.Flags().BoolVar(&opts.authResult, "authResult", false, `Flag that indicates whether the response returns the successful authentication attempts only.`)
	cmd.Flags().Int64Var(&opts.end, "end", 0, `Date and time when to stop retrieving database history. If you specify **end**, you must also specify **start**. This parameter uses UNIX epoch time in milliseconds.`)
	cmd.Flags().StringVar(&opts.ipAddress, "ipAddress", "", `One Internet Protocol address that attempted to authenticate with the database.`)
	cmd.Flags().IntVar(&opts.nLogs, "nLogs", 20000, `Maximum number of lines from the log to return.`)
	cmd.Flags().Int64Var(&opts.start, "start", 0, `Date and time when MongoDB Cloud begins retrieving database history. If you specify **start**, you must also specify **end**. This parameter uses UNIX epoch time in milliseconds.`)

	_ = cmd.MarkFlagRequired("clusterName")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type listAccessLogsByHostnameOpts struct {
	client     *admin.APIClient
	groupId    string
	hostname   string
	authResult bool
	end        int64
	ipAddress  string
	nLogs      int
	start      int64
	format     string
	tmpl       *template.Template
	resp       *admin.MongoDBAccessLogsList
}

func (opts *listAccessLogsByHostnameOpts) preRun() (err error) {
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

func (opts *listAccessLogsByHostnameOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.ListAccessLogsByHostnameApiParams{
		GroupId:    opts.groupId,
		Hostname:   opts.hostname,
		AuthResult: &opts.authResult,
		End:        &opts.end,
		IpAddress:  &opts.ipAddress,
		NLogs:      &opts.nLogs,
		Start:      &opts.start,
	}

	var err error
	opts.resp, _, err = opts.client.AccessTrackingApi.ListAccessLogsByHostnameWithParams(ctx, params).Execute()
	return err
}

func (opts *listAccessLogsByHostnameOpts) postRun(_ context.Context, w io.Writer) error {

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

func listAccessLogsByHostnameBuilder() *cobra.Command {
	opts := listAccessLogsByHostnameOpts{}
	cmd := &cobra.Command{
		Use:   "listAccessLogsByHostname",
		Short: "Return Database Access History for One Cluster using Its Hostname",
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
	cmd.Flags().StringVar(&opts.hostname, "hostname", "", `Fully qualified domain name or IP address of the MongoDB host that stores the log files that you want to download.`)
	cmd.Flags().BoolVar(&opts.authResult, "authResult", false, `Flag that indicates whether the response returns the successful authentication attempts only.`)
	cmd.Flags().Int64Var(&opts.end, "end", 0, `Date and time when to stop retrieving database history. If you specify **end**, you must also specify **start**. This parameter uses UNIX epoch time in milliseconds.`)
	cmd.Flags().StringVar(&opts.ipAddress, "ipAddress", "", `One Internet Protocol address that attempted to authenticate with the database.`)
	cmd.Flags().IntVar(&opts.nLogs, "nLogs", 20000, `Maximum number of lines from the log to return.`)
	cmd.Flags().Int64Var(&opts.start, "start", 0, `Date and time when MongoDB Cloud begins retrieving database history. If you specify **start**, you must also specify **end**. This parameter uses UNIX epoch time in milliseconds.`)

	_ = cmd.MarkFlagRequired("hostname")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

func accessTrackingBuilder() *cobra.Command {
	const use = "accessTracking"
	cmd := &cobra.Command{
		Use:     use,
		Short:   `Returns access logs for authentication attempts made to Atlas database deployments. To view database access history, you must have either the Project Owner or Organization Owner role.`,
		Aliases: cli.GenerateAliases(use),
	}
	cmd.AddCommand(
		listAccessLogsByClusterNameBuilder(),
		listAccessLogsByHostnameBuilder(),
	)
	return cmd
}

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

package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

type disableSlowOperationThresholdingOpts struct {
	client  *admin.APIClient
	groupId string
}

func (opts *disableSlowOperationThresholdingOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *disableSlowOperationThresholdingOpts) run(ctx context.Context, _ io.Writer) error {

	params := &admin.DisableSlowOperationThresholdingApiParams{
		GroupId: opts.groupId,
	}

	_, err := opts.client.PerformanceAdvisorApi.DisableSlowOperationThresholdingWithParams(ctx, params).Execute()
	return err
}

func disableSlowOperationThresholdingBuilder() *cobra.Command {
	opts := disableSlowOperationThresholdingOpts{}
	cmd := &cobra.Command{
		Use:   "disableSlowOperationThresholding",
		Short: "Disable Managed Slow Operation Threshold",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type enableSlowOperationThresholdingOpts struct {
	client  *admin.APIClient
	groupId string
}

func (opts *enableSlowOperationThresholdingOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *enableSlowOperationThresholdingOpts) run(ctx context.Context, _ io.Writer) error {

	params := &admin.EnableSlowOperationThresholdingApiParams{
		GroupId: opts.groupId,
	}

	_, err := opts.client.PerformanceAdvisorApi.EnableSlowOperationThresholdingWithParams(ctx, params).Execute()
	return err
}

func enableSlowOperationThresholdingBuilder() *cobra.Command {
	opts := enableSlowOperationThresholdingOpts{}
	cmd := &cobra.Command{
		Use:   "enableSlowOperationThresholding",
		Short: "Enable Managed Slow Operation Threshold",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type listSlowQueriesOpts struct {
	client     *admin.APIClient
	groupId    string
	processId  string
	duration   int64
	namespaces []string
	nLogs      int64
	since      int64
}

func (opts *listSlowQueriesOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *listSlowQueriesOpts) run(ctx context.Context, w io.Writer) error {

	params := &admin.ListSlowQueriesApiParams{
		GroupId:    opts.groupId,
		ProcessId:  opts.processId,
		Duration:   &opts.duration,
		Namespaces: &opts.namespaces,
		NLogs:      &opts.nLogs,
		Since:      &opts.since,
	}

	resp, _, err := opts.client.PerformanceAdvisorApi.ListSlowQueriesWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	_, err = fmt.Fprintln(w, string(prettyJSON))
	return err
}

func listSlowQueriesBuilder() *cobra.Command {
	opts := listSlowQueriesOpts{}
	cmd := &cobra.Command{
		Use:   "listSlowQueries",
		Short: "Return Slow Queries",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.processId, "processId", "", `Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests.`)
	cmd.Flags().Int64Var(&opts.duration, "duration", 0, `Length of time expressed during which the query finds slow queries among the managed namespaces in the cluster. This parameter expresses its value in milliseconds.

- If you don&#39;t specify the **since** parameter, the endpoint returns data covering the duration before the current time.
- If you specify neither the **duration** nor **since** parameters, the endpoint returns data from the previous 24 hours.`)
	cmd.Flags().StringSliceVar(&opts.namespaces, "namespaces", nil, `Namespaces from which to retrieve slow queries. A namespace consists of one database and one collection resource written as &#x60;.&#x60;: &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each namespace. Omit this parameter to return results for all namespaces.`)
	cmd.Flags().Int64Var(&opts.nLogs, "nLogs", 20000, `Maximum number of lines from the log to return.`)
	cmd.Flags().Int64Var(&opts.since, "since", 0, `Date and time from which the query retrieves the slow queries. This parameter expresses its value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).

- If you don&#39;t specify the **duration** parameter, the endpoint returns data covering from the **since** value and the current time.
- If you specify neither the **duration** nor the **since** parameters, the endpoint returns data from the previous 24 hours.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("processId")
	return cmd
}

type listSlowQueryNamespacesOpts struct {
	client    *admin.APIClient
	groupId   string
	processId string
	duration  int64
	since     int64
}

func (opts *listSlowQueryNamespacesOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *listSlowQueryNamespacesOpts) run(ctx context.Context, w io.Writer) error {

	params := &admin.ListSlowQueryNamespacesApiParams{
		GroupId:   opts.groupId,
		ProcessId: opts.processId,
		Duration:  &opts.duration,
		Since:     &opts.since,
	}

	resp, _, err := opts.client.PerformanceAdvisorApi.ListSlowQueryNamespacesWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	_, err = fmt.Fprintln(w, string(prettyJSON))
	return err
}

func listSlowQueryNamespacesBuilder() *cobra.Command {
	opts := listSlowQueryNamespacesOpts{}
	cmd := &cobra.Command{
		Use:   "listSlowQueryNamespaces",
		Short: "Return All Namespaces for One Host",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.processId, "processId", "", `Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests.`)
	cmd.Flags().Int64Var(&opts.duration, "duration", 0, `Length of time expressed during which the query finds suggested indexes among the managed namespaces in the cluster. This parameter expresses its value in milliseconds.

- If you don&#39;t specify the **since** parameter, the endpoint returns data covering the duration before the current time.
- If you specify neither the **duration** nor **since** parameters, the endpoint returns data from the previous 24 hours.`)
	cmd.Flags().Int64Var(&opts.since, "since", 0, `Date and time from which the query retrieves the suggested indexes. This parameter expresses its value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).

- If you don&#39;t specify the **duration** parameter, the endpoint returns data covering from the **since** value and the current time.
- If you specify neither the **duration** nor the **since** parameters, the endpoint returns data from the previous 24 hours.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("processId")
	return cmd
}

type listSuggestedIndexesOpts struct {
	client       *admin.APIClient
	groupId      string
	processId    string
	includeCount bool
	itemsPerPage int
	pageNum      int
	duration     int64
	namespaces   []string
	nExamples    int64
	nIndexes     int64
	since        int64
}

func (opts *listSuggestedIndexesOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *listSuggestedIndexesOpts) run(ctx context.Context, w io.Writer) error {

	params := &admin.ListSuggestedIndexesApiParams{
		GroupId:      opts.groupId,
		ProcessId:    opts.processId,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
		Duration:     &opts.duration,
		Namespaces:   &opts.namespaces,
		NExamples:    &opts.nExamples,
		NIndexes:     &opts.nIndexes,
		Since:        &opts.since,
	}

	resp, _, err := opts.client.PerformanceAdvisorApi.ListSuggestedIndexesWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	_, err = fmt.Fprintln(w, string(prettyJSON))
	return err
}

func listSuggestedIndexesBuilder() *cobra.Command {
	opts := listSuggestedIndexesOpts{}
	cmd := &cobra.Command{
		Use:   "listSuggestedIndexes",
		Short: "Return Suggested Indexes",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.processId, "processId", "", `Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)
	cmd.Flags().Int64Var(&opts.duration, "duration", 0, `Length of time expressed during which the query finds suggested indexes among the managed namespaces in the cluster. This parameter expresses its value in milliseconds.

- If you don&#39;t specify the **since** parameter, the endpoint returns data covering the duration before the current time.
- If you specify neither the **duration** nor **since** parameters, the endpoint returns data from the previous 24 hours.`)
	cmd.Flags().StringSliceVar(&opts.namespaces, "namespaces", nil, `Namespaces from which to retrieve suggested indexes. A namespace consists of one database and one collection resource written as &#x60;.&#x60;: &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each namespace. Omit this parameter to return results for all namespaces.`)
	cmd.Flags().Int64Var(&opts.nExamples, "nExamples", 5, `Maximum number of example queries that benefit from the suggested index.`)
	cmd.Flags().Int64Var(&opts.nIndexes, "nIndexes", 0, `Number that indicates the maximum indexes to suggest.`)
	cmd.Flags().Int64Var(&opts.since, "since", 0, `Date and time from which the query retrieves the suggested indexes. This parameter expresses its value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).

- If you don&#39;t specify the **duration** parameter, the endpoint returns data covering from the **since** value and the current time.
- If you specify neither the **duration** nor the **since** parameters, the endpoint returns data from the previous 24 hours.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("processId")
	return cmd
}

func performanceAdvisorBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "performanceAdvisor",
		Short: `Returns suggested indexes and slow query data for a database deployment. Also enables or disables MongoDB Cloud-managed slow operation thresholds. To view field values in a sample query, you must have the Project Data Access Read Only role or higher. Otherwise, MongoDB Cloud returns redacted data rather than the field values.`,
	}
	cmd.AddCommand(
		disableSlowOperationThresholdingBuilder(),
		enableSlowOperationThresholdingBuilder(),
		listSlowQueriesBuilder(),
		listSlowQueryNamespacesBuilder(),
		listSuggestedIndexesBuilder(),
	)
	return cmd
}

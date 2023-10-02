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
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"text/template"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

type getOrganizationEventOpts struct {
	client     *admin.APIClient
	orgId      string
	eventId    string
	includeRaw bool
	format     string
	tmpl       *template.Template
}

func (opts *getOrganizationEventOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(); err != nil {
		return err
	}

	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}
	if opts.orgId == "" {
		return errors.New(`required flag(s) "orgId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.orgId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.orgId)
	}

	if opts.format != "" {
		opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n")
	}

	return err
}

func (opts *getOrganizationEventOpts) run(ctx context.Context, _ io.Reader, w io.Writer) error {

	params := &admin.GetOrganizationEventApiParams{
		OrgId:      opts.orgId,
		EventId:    opts.eventId,
		IncludeRaw: &opts.includeRaw,
	}

	resp, _, err := opts.client.EventsApi.GetOrganizationEventWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	if opts.format == "" {
		_, err = fmt.Fprintln(w, string(prettyJSON))
		return err
	}

	var parsedJSON interface{}
	if err = json.Unmarshal([]byte(prettyJSON), &parsedJSON); err != nil {
		return err
	}

	err = opts.tmpl.Execute(w, parsedJSON)
	return err
}

func getOrganizationEventBuilder() *cobra.Command {
	opts := getOrganizationEventOpts{}
	cmd := &cobra.Command{
		Use:   "getOrganizationEvent",
		Short: "Return One Event from One Organization",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)
	cmd.Flags().StringVar(&opts.eventId, "eventId", "", `Unique 24-hexadecimal digit string that identifies the event that you want to return. Use the [/events](#tag/Events/operation/listOrganizationEvents) endpoint to retrieve all events to which the authenticated user has access.`)
	cmd.Flags().BoolVar(&opts.includeRaw, "includeRaw", false, `Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.`)

	_ = cmd.MarkFlagRequired("eventId")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type getProjectEventOpts struct {
	client     *admin.APIClient
	groupId    string
	eventId    string
	includeRaw bool
	format     string
	tmpl       *template.Template
}

func (opts *getProjectEventOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(); err != nil {
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
		opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n")
	}

	return err
}

func (opts *getProjectEventOpts) run(ctx context.Context, _ io.Reader, w io.Writer) error {

	params := &admin.GetProjectEventApiParams{
		GroupId:    opts.groupId,
		EventId:    opts.eventId,
		IncludeRaw: &opts.includeRaw,
	}

	resp, _, err := opts.client.EventsApi.GetProjectEventWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	if opts.format == "" {
		_, err = fmt.Fprintln(w, string(prettyJSON))
		return err
	}

	var parsedJSON interface{}
	if err = json.Unmarshal([]byte(prettyJSON), &parsedJSON); err != nil {
		return err
	}

	err = opts.tmpl.Execute(w, parsedJSON)
	return err
}

func getProjectEventBuilder() *cobra.Command {
	opts := getProjectEventOpts{}
	cmd := &cobra.Command{
		Use:   "getProjectEvent",
		Short: "Return One Event from One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.eventId, "eventId", "", `Unique 24-hexadecimal digit string that identifies the event that you want to return. Use the [/events](#tag/Events/operation/listProjectEvents) endpoint to retrieve all events to which the authenticated user has access.`)
	cmd.Flags().BoolVar(&opts.includeRaw, "includeRaw", false, `Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.`)

	_ = cmd.MarkFlagRequired("eventId")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type listOrganizationEventsOpts struct {
	client       *admin.APIClient
	orgId        string
	includeCount bool
	itemsPerPage int
	pageNum      int
	eventType    []string
	includeRaw   bool
	maxDate      string
	minDate      string
	format       string
	tmpl         *template.Template
}

func (opts *listOrganizationEventsOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(); err != nil {
		return err
	}

	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}
	if opts.orgId == "" {
		return errors.New(`required flag(s) "orgId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.orgId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.orgId)
	}

	if opts.format != "" {
		opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n")
	}

	return err
}

func (opts *listOrganizationEventsOpts) run(ctx context.Context, _ io.Reader, w io.Writer) error {

	var maxDate *time.Time
	var errMaxDate error
	if opts.maxDate != "" {
		*maxDate, errMaxDate = time.Parse(time.RFC3339, opts.maxDate)
		if errMaxDate != nil {
			return errMaxDate
		}
	}

	var minDate *time.Time
	var errMinDate error
	if opts.minDate != "" {
		*minDate, errMinDate = time.Parse(time.RFC3339, opts.minDate)
		if errMinDate != nil {
			return errMinDate
		}
	}

	params := &admin.ListOrganizationEventsApiParams{
		OrgId:        opts.orgId,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
		EventType:    &opts.eventType,
		IncludeRaw:   &opts.includeRaw,
		MaxDate:      maxDate,
		MinDate:      minDate,
	}

	resp, _, err := opts.client.EventsApi.ListOrganizationEventsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	if opts.format == "" {
		_, err = fmt.Fprintln(w, string(prettyJSON))
		return err
	}

	var parsedJSON interface{}
	if err = json.Unmarshal([]byte(prettyJSON), &parsedJSON); err != nil {
		return err
	}

	err = opts.tmpl.Execute(w, parsedJSON)
	return err
}

func listOrganizationEventsBuilder() *cobra.Command {
	opts := listOrganizationEventsOpts{}
	cmd := &cobra.Command{
		Use:   "listOrganizationEvents",
		Short: "Return All Events from One Organization",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)
	cmd.Flags().StringSliceVar(&opts.eventType, "eventType", nil, `Category of incident recorded at this moment in time.

**IMPORTANT**: The complete list of event type values changes frequently.`)
	cmd.Flags().BoolVar(&opts.includeRaw, "includeRaw", false, `Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.`)
	cmd.Flags().StringVar(&opts.maxDate, "maxDate", "", `Date and time from when MongoDB Cloud stops returning events. This parameter uses the &lt;a href&#x3D;&quot;https://en.wikipedia.org/wiki/ISO_8601&quot; target&#x3D;&quot;_blank&quot; rel&#x3D;&quot;noopener noreferrer&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC.`)
	cmd.Flags().StringVar(&opts.minDate, "minDate", "", `Date and time from when MongoDB Cloud starts returning events. This parameter uses the &lt;a href&#x3D;&quot;https://en.wikipedia.org/wiki/ISO_8601&quot; target&#x3D;&quot;_blank&quot; rel&#x3D;&quot;noopener noreferrer&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC.`)

	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type listProjectEventsOpts struct {
	client            *admin.APIClient
	groupId           string
	includeCount      bool
	itemsPerPage      int
	pageNum           int
	clusterNames      []string
	eventType         []string
	excludedEventType []string
	includeRaw        bool
	maxDate           string
	minDate           string
	format            string
	tmpl              *template.Template
}

func (opts *listProjectEventsOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(); err != nil {
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
		opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n")
	}

	return err
}

func (opts *listProjectEventsOpts) run(ctx context.Context, _ io.Reader, w io.Writer) error {

	var maxDate *time.Time
	var errMaxDate error
	if opts.maxDate != "" {
		*maxDate, errMaxDate = time.Parse(time.RFC3339, opts.maxDate)
		if errMaxDate != nil {
			return errMaxDate
		}
	}

	var minDate *time.Time
	var errMinDate error
	if opts.minDate != "" {
		*minDate, errMinDate = time.Parse(time.RFC3339, opts.minDate)
		if errMinDate != nil {
			return errMinDate
		}
	}

	params := &admin.ListProjectEventsApiParams{
		GroupId:           opts.groupId,
		IncludeCount:      &opts.includeCount,
		ItemsPerPage:      &opts.itemsPerPage,
		PageNum:           &opts.pageNum,
		ClusterNames:      &opts.clusterNames,
		EventType:         &opts.eventType,
		ExcludedEventType: &opts.excludedEventType,
		IncludeRaw:        &opts.includeRaw,
		MaxDate:           maxDate,
		MinDate:           minDate,
	}

	resp, _, err := opts.client.EventsApi.ListProjectEventsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	if opts.format == "" {
		_, err = fmt.Fprintln(w, string(prettyJSON))
		return err
	}

	var parsedJSON interface{}
	if err = json.Unmarshal([]byte(prettyJSON), &parsedJSON); err != nil {
		return err
	}

	err = opts.tmpl.Execute(w, parsedJSON)
	return err
}

func listProjectEventsBuilder() *cobra.Command {
	opts := listProjectEventsOpts{}
	cmd := &cobra.Command{
		Use:   "listProjectEvents",
		Short: "Return All Events from One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)
	cmd.Flags().StringSliceVar(&opts.clusterNames, "clusterNames", nil, `Human-readable label that identifies the cluster.`)
	cmd.Flags().StringSliceVar(&opts.eventType, "eventType", nil, `Category of incident recorded at this moment in time.

**IMPORTANT**: The complete list of event type values changes frequently.`)
	cmd.Flags().StringSliceVar(&opts.excludedEventType, "excludedEventType", nil, `Category of event that you would like to exclude from query results, such as CLUSTER_CREATED

**IMPORTANT**: Event type names change frequently. Verify that you specify the event type correctly by checking the complete list of event types.`)
	cmd.Flags().BoolVar(&opts.includeRaw, "includeRaw", false, `Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.`)
	cmd.Flags().StringVar(&opts.maxDate, "maxDate", "", `Date and time from when MongoDB Cloud stops returning events. This parameter uses the &lt;a href&#x3D;&quot;https://en.wikipedia.org/wiki/ISO_8601&quot; target&#x3D;&quot;_blank&quot; rel&#x3D;&quot;noopener noreferrer&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC.`)
	cmd.Flags().StringVar(&opts.minDate, "minDate", "", `Date and time from when MongoDB Cloud starts returning events. This parameter uses the &lt;a href&#x3D;&quot;https://en.wikipedia.org/wiki/ISO_8601&quot; target&#x3D;&quot;_blank&quot; rel&#x3D;&quot;noopener noreferrer&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC.`)

	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

func eventsBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "events",
		Short: `Returns events. This collection remains under revision and may change.`,
	}
	cmd.AddCommand(
		getOrganizationEventBuilder(),
		getProjectEventBuilder(),
		listOrganizationEventsBuilder(),
		listProjectEventsBuilder(),
	)
	return cmd
}

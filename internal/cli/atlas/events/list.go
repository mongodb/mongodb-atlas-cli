// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package events

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	store "github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20231115004/admin"
)

type EventListOpts struct {
	cli.ListOpts
	EventType []string
	MinDate   string
	MaxDate   string
}

type ListOpts struct {
	EventListOpts
	cli.OutputOpts
	orgID     string
	projectID string
	store     store.EventLister
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var listTemplate = `ID	TYPE	CREATED{{range valueOrEmptySlice .Results}}
{{.Id}}	{{.EventTypeName}}	{{.Created}}{{end}}
`

func (opts *ListOpts) Run() error {
	var r interface{}
	var err error

	if opts.orgID != "" {
		listEventsAPIParams := opts.NewOrgListOptions()
		r, err = opts.store.OrganizationEvents(&listEventsAPIParams)
	} else if opts.projectID != "" {
		listEventsAPIParams := opts.NewProjectListOptions()
		r, err = opts.store.ProjectEvents(&listEventsAPIParams)
	}
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *ListOpts) NewOrgListOptions() admin.ListOrganizationEventsApiParams {
	var eventType *[]string
	if len(opts.EventType) > 0 {
		eventType = &opts.EventType
	}
	p := admin.ListOrganizationEventsApiParams{
		OrgId:     opts.orgID,
		EventType: eventType,
		MaxDate:   pointer.StringToTimePointer(opts.MaxDate),
		MinDate:   pointer.StringToTimePointer(opts.MinDate),
	}
	if opts.ItemsPerPage > 0 {
		p.ItemsPerPage = &opts.ItemsPerPage
	}
	if opts.PageNum > 0 {
		p.PageNum = &opts.PageNum
	}
	return p
}

func (opts *ListOpts) NewProjectListOptions() admin.ListProjectEventsApiParams {
	var eventType *[]string
	if len(opts.EventType) > 0 {
		eventType = &opts.EventType
	}
	p := admin.ListProjectEventsApiParams{
		GroupId:   opts.projectID,
		EventType: eventType,
		MaxDate:   pointer.StringToTimePointer(opts.MaxDate),
		MinDate:   pointer.StringToTimePointer(opts.MinDate),
	}
	if opts.ItemsPerPage > 0 {
		p.ItemsPerPage = &opts.ItemsPerPage
	}
	if opts.PageNum > 0 {
		p.PageNum = &opts.PageNum
	}
	return p
}

// ListBuilder
//
//	atlas event(s) list
//
// [--projectId projectId]
// [--orgId orgId]
// [--page N]
// [--limit N]
// [--minDate minDate]
// [--maxDate maxDate].
func ListBuilder() *cobra.Command {
	opts := &ListOpts{}
	opts.Template = listTemplate
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Return all events for an organization or project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Deprecated: `  
  To return project events prefer
  atlas events projects list [--projectId <projectId>]

  To return organization events prefer
  atlas events organizations list [--orgId <orgId>]
`,
		Aliases: []string{"ls"},
		Args:    require.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if opts.orgID != "" && opts.projectID != "" {
				return fmt.Errorf("both --%s and --%s set", flag.ProjectID, flag.OrgID)
			}
			if opts.orgID == "" && opts.projectID == "" {
				return fmt.Errorf("--%s or --%s must be set", flag.ProjectID, flag.OrgID)
			}
			opts.OutWriter = cmd.OutOrStdout()
			return opts.initStore(cmd.Context())()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().IntVar(&opts.PageNum, flag.Page, 0, usage.Page)
	cmd.Flags().IntVar(&opts.ItemsPerPage, flag.Limit, 0, usage.Limit)

	cmd.Flags().StringSliceVar(&opts.EventType, flag.TypeFlag, nil, usage.Event)
	cmd.Flags().StringVar(&opts.MaxDate, flag.MaxDate, "", usage.MaxDate)
	cmd.Flags().StringVar(&opts.MinDate, flag.MinDate, "", usage.MinDate)

	cmd.Flags().StringVar(&opts.projectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVar(&opts.orgID, flag.OrgID, "", usage.OrgID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}

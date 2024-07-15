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
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/convert"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
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
	var r any
	var err error

	if opts.orgID != "" {
		listEventsAPIParams, lErr := opts.NewOrgListOptions()
		if lErr != nil {
			return lErr
		}

		if r, err = opts.store.OrganizationEvents(listEventsAPIParams); err != nil {
			return err
		}
	} else if opts.projectID != "" {
		listEventsAPIParams, lErr := opts.NewProjectListOptions()
		if lErr != nil {
			return lErr
		}
		if r, err = opts.store.ProjectEvents(listEventsAPIParams); err != nil {
			return err
		}
	}

	return opts.Print(r)
}

func (opts *ListOpts) NewOrgListOptions() (*admin.ListOrganizationEventsApiParams, error) {
	var eventType *[]string
	var err error

	if len(opts.EventType) > 0 {
		eventType = &opts.EventType
	}
	p := &admin.ListOrganizationEventsApiParams{
		OrgId:     opts.orgID,
		EventType: eventType,
	}
	if p.MaxDate, err = parseDate(opts.MaxDate); err != nil {
		return p, err
	}
	if p.MinDate, err = parseDate(opts.MinDate); err != nil {
		return p, err
	}
	if opts.ItemsPerPage > 0 {
		p.ItemsPerPage = &opts.ItemsPerPage
	}
	if opts.PageNum > 0 {
		p.PageNum = &opts.PageNum
	}
	if opts.OmitCount {
		p.IncludeCount = pointer.Get(false)
	}
	return p, nil
}

func (opts *ListOpts) NewProjectListOptions() (*admin.ListProjectEventsApiParams, error) {
	var eventType *[]string
	var err error
	if len(opts.EventType) > 0 {
		eventType = &opts.EventType
	}
	p := &admin.ListProjectEventsApiParams{
		GroupId:   opts.projectID,
		EventType: eventType,
	}

	if p.MaxDate, err = parseDate(opts.MaxDate); err != nil {
		return p, err
	}

	if p.MinDate, err = parseDate(opts.MinDate); err != nil {
		return p, err
	}

	if opts.ItemsPerPage > 0 {
		p.ItemsPerPage = &opts.ItemsPerPage
	}
	if opts.PageNum > 0 {
		p.PageNum = &opts.PageNum
	}
	if opts.OmitCount {
		p.IncludeCount = pointer.Get(false)
	}

	return p, nil
}

func parseDate(date string) (*time.Time, error) {
	if date == "" {
		return nil, nil
	}

	parsedDate, err := convert.ParseTimestamp(date)
	if err != nil {
		return nil, err
	}
	return &parsedDate, nil
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
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			if opts.orgID != "" && opts.projectID != "" {
				return fmt.Errorf("both --%s and --%s set", flag.ProjectID, flag.OrgID)
			}
			if opts.orgID == "" && opts.projectID == "" {
				return fmt.Errorf("--%s or --%s must be set", flag.ProjectID, flag.OrgID)
			}
			opts.OutWriter = cmd.OutOrStdout()

			return opts.initStore(cmd.Context())()
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().IntVar(&opts.PageNum, flag.Page, cli.DefaultPage, usage.Page)
	cmd.Flags().IntVar(&opts.ItemsPerPage, flag.Limit, cli.DefaultPageLimit, usage.Limit)
	cmd.Flags().BoolVar(&opts.OmitCount, flag.OmitCount, false, usage.OmitCount)

	cmd.Flags().StringSliceVar(&opts.EventType, flag.TypeFlag, nil, usage.Event)
	cmd.Flags().StringVar(&opts.MaxDate, flag.MaxDate, "", usage.MaxDate)
	cmd.Flags().StringVar(&opts.MinDate, flag.MinDate, "", usage.MinDate)

	cmd.Flags().StringVar(&opts.projectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVar(&opts.orgID, flag.OrgID, "", usage.OrgID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}

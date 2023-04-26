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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	store "github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas/mongodbatlasv2"
	"k8s.io/utils/pointer"
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

var listTemplate = `ID	TYPE	CREATED{{range .Results}}
{{.Id}}	{{.EventType}}	{{.Created}}{{end}}
`

func (opts *ListOpts) Run() error {
	var r interface{}
	var err error
	minDate, _ := time.Parse(time.RFC3339, opts.MinDate)
	maxDate, _ := time.Parse(time.RFC3339, opts.MaxDate)

	if opts.orgID != "" {
		// TODO Support multiple event types by API
		eventType, _ := mongodbatlasv2.NewEventTypeForOrgFromValue(opts.EventType[0])
		// TODO Use APIparams objects directly in the CLI
		listEventsAPIParams := mongodbatlasv2.ListOrganizationEventsApiParams{
			OrgId:        opts.orgID,
			ItemsPerPage: pointer.Int32(int32(opts.ItemsPerPage)),
			PageNum:      pointer.Int32(int32(opts.PageNum)),
			EventType:    eventType,
			IncludeRaw:   new(bool),
			MaxDate:      &minDate,
			MinDate:      &maxDate,
		}
		r, err = opts.store.OrganizationEvents(&listEventsAPIParams)
	} else if opts.projectID != "" {
		// TODO CLOUDP-17348 event type is array but we expect single event
		var eventType *mongodbatlasv2.EventTypeForNdsGroup
		if len(opts.EventType) > 1 {
			eventType, _ = mongodbatlasv2.NewEventTypeForNdsGroupFromValue(opts.EventType[0])
		}
		// TODO  CLOUDP-173460 Use APIparams objects directly in SDK (without need for store)
		listEventsAPIParams := mongodbatlasv2.ListProjectEventsApiParams{
			GroupId:      opts.projectID,
			ItemsPerPage: pointer.Int32(int32(opts.ItemsPerPage)),
			PageNum:      pointer.Int32(int32(opts.PageNum)),
			EventType:    eventType,
			IncludeRaw:   new(bool),
			MaxDate:      &minDate,
			MinDate:      &maxDate,
		}
		r, err = opts.store.ProjectEvents(&listEventsAPIParams)
	}
	if err != nil {
		return err
	}

	return opts.Print(r)
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
  mongocli atlas|ops-manager|cloud-manager events projects list [--projectId <projectId>]

  To return organization events prefer
  mongocli atlas|ops-manager|cloud-manager events organizations list [--orgId <orgId>]
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

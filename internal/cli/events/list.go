// Copyright 2020 MongoDB Inc
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
package events

import (
	"fmt"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type ListOpts struct {
	cli.ListOpts
	cli.OutputOpts
	orgID     string
	projectID string
	eventType []string
	minDate   string
	maxDate   string
	store     store.EventLister
}

func (opts *ListOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

var listTemplate = `ID	TYPE	CREATED{{range .Results}}
{{.ID}}	{{.EventTypeName}}	{{.Created}}{{end}}
`

func (opts *ListOpts) Run() error {
	listOpts := opts.newEventListOptions()

	var r *atlas.EventResponse
	var err error

	if opts.orgID != "" {
		r, err = opts.store.OrganizationEvents(opts.orgID, listOpts)
	} else if opts.projectID != "" {
		r, err = opts.store.ProjectEvents(opts.projectID, listOpts)
	}
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *ListOpts) newEventListOptions() *atlas.EventListOptions {
	return &atlas.EventListOptions{
		ListOptions: atlas.ListOptions{
			PageNum:      opts.PageNum,
			ItemsPerPage: opts.ItemsPerPage,
		},
		EventType: opts.eventType,
		MinDate:   opts.minDate,
		MaxDate:   opts.maxDate,
	}
}

// mongocli atlas event(s) list [--projectId projectId] [--orgId orgId] [--page N] [--limit N] [--minDate minDate] [--maxDate maxDate]
func ListBuilder() *cobra.Command {
	opts := &ListOpts{}
	opts.Template = listTemplate
	cmd := &cobra.Command{
		Use:     "list",
		Short:   listEvents,
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
			return opts.initStore()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().IntVar(&opts.PageNum, flag.Page, 0, usage.Page)
	cmd.Flags().IntVar(&opts.ItemsPerPage, flag.Limit, 0, usage.Limit)

	cmd.Flags().StringSliceVar(&opts.eventType, flag.Type, nil, usage.Event)
	cmd.Flags().StringVar(&opts.maxDate, flag.MaxDate, "", usage.MaxDate)
	cmd.Flags().StringVar(&opts.minDate, flag.MinDate, "", usage.MinDate)

	cmd.Flags().StringVar(&opts.projectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVar(&opts.orgID, flag.OrgID, "", usage.OrgID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}

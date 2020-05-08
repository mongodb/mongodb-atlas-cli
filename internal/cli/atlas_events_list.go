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
package cli

import (
	"fmt"

	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type atlasEventsListOpts struct {
	listOpts
	orgID     string
	projectID string
	eventType []string
	minDate   string
	maxDate   string
	store     store.EventLister
}

func (opts *atlasEventsListOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *atlasEventsListOpts) Run() error {
	listOpts := opts.newEventListOptions()

	var result *atlas.EventResponse
	var err error

	if opts.orgID != "" {
		result, err = opts.store.OrganizationEvents(opts.orgID, listOpts)
	} else if opts.projectID != "" {
		result, err = opts.store.ProjectEvents(opts.projectID, listOpts)
	}

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func (opts *atlasEventsListOpts) newEventListOptions() *atlas.EventListOptions {
	return &atlas.EventListOptions{
		ListOptions: atlas.ListOptions{
			PageNum:      opts.pageNum,
			ItemsPerPage: opts.itemsPerPage,
		},
		EventType: opts.eventType,
		MinDate:   opts.minDate,
		MaxDate:   opts.maxDate,
	}
}

// mongocli atlas event(s) list [--projectId projectId] [--orgId orgId] [--page N] [--limit N] [--minDate minDate] [--maxDate maxDate]
func AtlasEventsListBuilder() *cobra.Command {
	opts := &atlasEventsListOpts{}
	cmd := &cobra.Command{
		Use:     "list",
		Short:   description.ListEvents,
		Aliases: []string{"ls"},
		Args:    cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if opts.orgID != "" && opts.projectID != "" {
				return fmt.Errorf("both --%s and --%s set", flags.ProjectID, flags.OrgID)
			}
			if opts.orgID == "" && opts.projectID == "" {
				return fmt.Errorf("--%s or --%s must be set", flags.ProjectID, flags.OrgID)
			}
			return opts.initStore()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().IntVar(&opts.pageNum, flags.Page, 0, usage.Page)
	cmd.Flags().IntVar(&opts.itemsPerPage, flags.Limit, 0, usage.Limit)

	cmd.Flags().StringSliceVar(&opts.eventType, flags.Type, nil, usage.Event)
	cmd.Flags().StringVar(&opts.maxDate, flags.MaxDate, "", usage.MaxDate)
	cmd.Flags().StringVar(&opts.minDate, flags.MinDate, "", usage.MinDate)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVar(&opts.orgID, flags.OrgID, "", usage.OrgID)

	return cmd
}

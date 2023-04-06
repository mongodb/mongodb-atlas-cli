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
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type orgListOpts struct {
	cli.OutputOpts
	EventListOpts
	cli.GlobalOpts
	store store.OrganizationEventLister
}

func (opts *orgListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *orgListOpts) Run() error {
	listOpts := opts.newEventListOptions()

	var r *atlas.EventResponse
	var err error
	r, err = opts.store.OrganizationEvents(opts.ConfigOrgID(), listOpts)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// OrgListBuilder
//
//	mongocli atlas event(s) list
//
// [--orgId orgId]
// [--page N]
// [--limit N]
// [--minDate minDate]
// [--maxDate maxDate].
func OrgListBuilder() *cobra.Command {
	opts := &orgListOpts{}
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Return all events for the specified organization.",
		Aliases: []string{"ls"},
		Args:    require.NoArgs,
		Example: fmt.Sprintf(`  # Return a JSON-formatted list of events for the organization with the ID 5dd5a6b6f10fab1d71a58495:
  %s events organizations list --orgId 5dd5a6b6f10fab1d71a58495 --output json`, cli.ExampleAtlasEntryPoint()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
			)
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

	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}

func OrgsBuilder() *cobra.Command {
	const use = "organizations"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Organization operations.",
		Long:    "List organization events.",
		Aliases: cli.GenerateAliases(use, "orgs", "org"),
	}
	cmd.AddCommand(
		OrgListBuilder(),
	)
	return cmd
}

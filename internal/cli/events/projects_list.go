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

package events

import (
	"context"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type projectListOpts struct {
	EventListOpts
	cli.GlobalOpts
	cli.OutputOpts
	store store.ProjectEventLister
}

func (opts *projectListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *projectListOpts) Run() error {
	listEventsAPIParams, err := opts.NewProjectListOptions()
	if err != nil {
		return err
	}

	r, err := opts.store.ProjectEvents(listEventsAPIParams)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *projectListOpts) NewProjectListOptions() (*admin.ListProjectEventsApiParams, error) {
	var eventType *[]string
	var err error
	if len(opts.EventType) > 0 {
		eventType = &opts.EventType
	}
	p := &admin.ListProjectEventsApiParams{
		GroupId:   opts.ConfigProjectID(),
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

// ProjectListBuilder
//
//	atlas event(s) list
//
// [--page N]
// [--limit N]
// [--minDate minDate]
// [--maxDate maxDate].
func ProjectListBuilder() *cobra.Command {
	opts := &projectListOpts{}
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Return all events for the specified project.",
		Aliases: []string{"ls"},
		Args:    require.NoArgs,
		Example: `  # Return a JSON-formatted list of events for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas events projects list --Id 5e2211c17a3e5a48f5497de3 --output json

  # Return a JSON-formatted list of events between 2024-03-18T14:40:03-0000 and 2024-03-18T15:00:03-0000 and for the project with the ID 5e2211c17a3e5a48f5497de3
  atlas events projects list --output json --projectId 5e2211c17a3e5a48f5497de3  --minDate 2024-03-18T14:40:03-0000 --maxDate 2024-03-18T15:00:03-0000`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
			)
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

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}

func ProjectsBuilder() *cobra.Command {
	const use = "projects"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Project operations.",
		Long:    "List projects events.",
		Aliases: cli.GenerateAliases(use),
	}
	cmd.AddCommand(
		ProjectListBuilder(),
	)

	return cmd
}

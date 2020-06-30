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

package alerts

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type GlobalListOpts struct {
	cli.ListOpts
	status string
	store  store.GlobalAlertLister
}

func (opts *GlobalListOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *GlobalListOpts) Run() error {
	alertOpts := opts.newAlertsListOptions()

	result, err := opts.store.GlobalAlerts(alertOpts)
	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func (opts *GlobalListOpts) newAlertsListOptions() *atlas.AlertsListOptions {
	return &atlas.AlertsListOptions{
		Status:      opts.status,
		ListOptions: *opts.NewListOptions(),
	}
}

// mongocli om|cm alert(s) global list [--status status]
func GlobalListBuilder() *cobra.Command {
	opts := &GlobalListOpts{}
	cmd := &cobra.Command{
		Use:     "list",
		Short:   description.ListGlobalAlerts,
		Aliases: []string{"ls"},
		Args:    cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().IntVar(&opts.PageNum, flag.Page, 0, usage.Page)
	cmd.Flags().IntVar(&opts.ItemsPerPage, flag.Limit, 0, usage.Limit)
	cmd.Flags().StringVar(&opts.status, flag.Status, "", usage.Status)

	return cmd
}

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
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type atlasAlertsListOpts struct {
	globalOpts
	listOpts
	status string
	store  store.AlertLister
}

func (opts *atlasAlertsListOpts) initStore() error {
	var err error
	opts.store, err = store.New()
	return err
}

func (opts *atlasAlertsListOpts) Run() error {

	listOpts := opts.newAlertsListOptions()
	result, err := opts.store.Alerts(opts.ProjectID(), listOpts)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func (opts *atlasAlertsListOpts) newAlertsListOptions() *atlas.AlertsListOptions {
	o := new(atlas.AlertsListOptions)
	o.Status = opts.status
	o.ItemsPerPage = opts.itemsPerPage
	o.PageNum = opts.pageNum

	return o
}

// mongocli atlas alerts list [--status status] [--projectId projectId] [--page N] [--limit N]
func AtlasAlertsListBuilder() *cobra.Command {
	opts := new(atlasAlertsListOpts)
	cmd := &cobra.Command{
		Use:     "list",
		Short:   description.ListAlerts,
		Aliases: []string{"ls"},
		Args:    cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().IntVar(&opts.pageNum, flags.Page, 0, usage.Page)
	cmd.Flags().IntVar(&opts.itemsPerPage, flags.Limit, 0, usage.Limit)
	cmd.Flags().StringVar(&opts.status, flags.Status, "", usage.Status)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}

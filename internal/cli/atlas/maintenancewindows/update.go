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

package maintenancewindows

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	dayOfWeek           string
	hourOfDay  string
	startASAP bool
	store        store.OnlineArchiveUpdater
}

func (opts *UpdateOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

var updateTemplate = "Maintenance window '{{.ID}}' updated.\n"

func (opts *UpdateOpts) Run() error {
	archive := opts.newOnlineArchive()
	r, err := opts.store.UpdateOnlineArchive(opts.ConfigProjectID(), opts.clusterName, archive)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *UpdateOpts) newOnlineArchive() *atlas.OnlineArchive {
	archive := &atlas.OnlineArchive{
		ID: opts.id,
		Criteria: &atlas.OnlineArchiveCriteria{
			ExpireAfterDays: opts.archiveAfter,
		},
	}
	return archive
}

// mongocli atlas maintenanceWindow(s) update(s) <ID> --dayOfWeek dayOfWeek --hourOfDay hourOfDay --startASAP [--projectId projectId]
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	cmd := &cobra.Command{
		Use:   "update <ID>",
		Short: maintenanceWindowsArchive,
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.dayOfWeek, flag.DayOfWeek, "", usage.DayOfWeek)
	cmd.Flags().StringVar(&opts.hourOfDay, flag.HourOfDay, "", usage.HourOfDay)
	cmd.Flags().BoolVar(&opts.startASAP, flag.StartASAP, false, usage.StartASAP)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}

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

package maintenance

import (
	"fmt"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type ClearOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	dayOfWeek int
	hourOfDay int
	startASAP bool
	store     store.MaintenanceWindowClearer
}

func (opts *ClearOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

var clearTemplate = "Maintenance window removed.\n"

func (opts *ClearOpts) Run() error {
	err := opts.store.ClearMaintenanceWindow(opts.ConfigProjectID())
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	fmt.Print(clearTemplate)
	return nil
}


// mongocli atlas maintenanceWindow(s) update(s) --dayOfWeek dayOfWeek --hourOfDay hourOfDay --startASAP [--projectId projectId]
func ClearBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	cmd := &cobra.Command{
		Use:   "update",
		Short: updateMaintenanceWindows,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), clearTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}

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
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=defer_mock_test.go -package=maintenance . Deferrer

type Deferrer interface {
	DeferMaintenanceWindow(string) error
}

type DeferOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	store Deferrer
}

func (opts *DeferOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var deferTemplate = "Maintenance window deferred.\n"

func (opts *DeferOpts) Run() error {
	err := opts.store.DeferMaintenanceWindow(opts.ConfigProjectID())
	if err != nil {
		return err
	}
	return opts.Print(nil)
}

// atlas maintenanceWindow(s) defer [--projectId projectId].
func DeferBuilder() *cobra.Command {
	opts := &DeferOpts{}
	cmd := &cobra.Command{
		Use:   "defer",
		Short: "Defer scheduled maintenance for your project for one week.",
		Long:  longDesc + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Annotations: map[string]string{
			"output": deferTemplate,
		},
		Example: `  # Defer scheduled maintenance for one week for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas maintenanceWindows defer --projectId 5e2211c17a3e5a48f5497de3 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), deferTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}

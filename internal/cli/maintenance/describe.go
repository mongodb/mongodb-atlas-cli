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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312004/admin"
)

var describeTemplate = `DAY OF THE WEEK	HOUR OF DAY	START ASAP
{{.DayOfWeek}}	{{.HourOfDay}}	{{.StartASAP}}
`

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=describe_mock_test.go -package=maintenance . Describer

type Describer interface {
	MaintenanceWindow(string) (*atlasv2.GroupMaintenanceWindow, error)
}

type DescribeOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	store Describer
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.MaintenanceWindow(opts.ConfigProjectID())
	if err != nil {
		return err
	}
	return opts.Print(r)
}

// atlas maintenanceWindow(s) describe|get [--projectId projectId].
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:         "describe",
		Aliases:     []string{"get"},
		Short:       "Return the maintenance window details for your project.",
		Long:        longDesc + fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Annotations: map[string]string{"output": describeTemplate},
		Example: `  # Return the maintenance window for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas maintenanceWindows describe --projectId 5e2211c17a3e5a48f5497de3 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
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

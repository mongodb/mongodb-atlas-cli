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

package settings

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=describe_mock_test.go -package=settings . AlertConfigurationDescriber

type AlertConfigurationDescriber interface {
	AlertConfiguration(string, string) (*atlasv2.GroupAlertsConfig, error)
}

type describeOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	alertID string
	store   AlertConfigurationDescriber
}

func (opts *describeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var settingsTemplate = `ID	TYPE	ENABLED
{{.Id}}	{{.EventTypeName}}	{{.Enabled}}
`

func (opts *describeOpts) Run() error {
	r, err := opts.store.AlertConfiguration(opts.ConfigProjectID(), opts.alertID)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func describeBuilder() *cobra.Command {
	opts := new(describeOpts)
	cmd := &cobra.Command{
		Use:   "describe <alertConfigId>",
		Short: "Return the details for the specified alert settings for your project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Example: `  #  Return the JSON-formatted details for the alert settings with the ID 5d1113b25a115342acc2d1aa in the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas alerts settings describe 5d1113b25a115342acc2d1aa --projectId 5e2211c17a3e5a48f5497de3 --output json`,
		Annotations: map[string]string{
			"alertConfigIdDesc": "Unique identifier of the alert settings you want to describe.",
			"output":            settingsTemplate,
		},
		Aliases: []string{"get"},
		Args:    require.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), settingsTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.alertID = args[0]
			return opts.Run()
		},
	}

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}

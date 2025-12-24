// Copyright 2021 MongoDB Inc
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

package regionalmodes

import (
	"context"
	"fmt"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312011/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=disable_mock_test.go -package=regionalmodes . RegionalizedPrivateEndpointSettingUpdater

type RegionalizedPrivateEndpointSettingUpdater interface {
	UpdateRegionalizedPrivateEndpointSetting(string, bool) (*atlasv2.ProjectSettingItem, error)
}

type DisableOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	store RegionalizedPrivateEndpointSettingUpdater
}

var disableTemplate = "Regionalized private endpoint setting disabled.\n"

func (opts *DisableOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DisableOpts) Run() error {
	r, err := opts.store.UpdateRegionalizedPrivateEndpointSetting(opts.ConfigProjectID(), false)
	if err != nil {
		return err
	}
	return opts.Print(r)
}

// atlas privateEndpoint(s) regionalMode(s) disable [--projectId projectId].
func DisableBuilder() *cobra.Command {
	opts := &DisableOpts{}
	cmd := &cobra.Command{
		Use:   "disable",
		Short: "Disable the regionalized private endpoint setting for your project.",
		Long: `This disables the ability to create multiple private resources per region in all cloud service providers for this project.

` + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Annotations: map[string]string{
			"output": disableTemplate,
		},
		Example: `  # Disable the regionalied private endpoint setting in the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas privateEndpoints regionalModes disable --projectId 5e2211c17a3e5a48f5497de3 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), disableTemplate),
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

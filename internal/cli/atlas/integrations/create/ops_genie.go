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

package create

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

const opsGenieType = "OPS_GENIE"

type OpsGenieOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	apiKey string
	region string
	store  store.IntegrationCreator
}

func (opts *OpsGenieOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var createTemplateOpsGenie = "Ops Genie integration configured.\n"

func (opts *OpsGenieOpts) Run() error {
	r, err := opts.store.CreateIntegration(opts.ConfigProjectID(), opsGenieType, opts.newOpsGenieIntegration())
	if err != nil {
		return err
	}
	return opts.Print(r)
}

func (opts *OpsGenieOpts) newOpsGenieIntegration() *atlas.ThirdPartyIntegration {
	return &atlas.ThirdPartyIntegration{
		Type:   opsGenieType,
		Region: opts.region,
		APIKey: opts.apiKey,
	}
}

// mongocli atlas integration(s) create OPS_GENIE --apiKey apiKey --region region [--projectId projectId].
func OpsGenieBuilder() *cobra.Command {
	opts := &OpsGenieOpts{}
	cmd := &cobra.Command{
		Use:     opsGenieType,
		Aliases: []string{"ops_genie", "opsGenie"},
		Short:   "Create or update an Opsgenie integration for your project.",
		Long:    "The requesting API key must have the Organization Owner or Project Owner role to configure an integration with Opsgenie.",
		Example: fmt.Sprintf(`  # Integrate Opsgenie with Atlas for the project with the ID 5e2211c17a3e5a48f5497de3:
  %s integrations create OPS_GENIE --apiKey a1a23bcdef45ghijk6789 --projectId 5e2211c17a3e5a48f5497de3 --output json`, cli.ExampleAtlasEntryPoint()),
		Args: require.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplateOpsGenie),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.region, flag.Region, "US", usage.APIRegion)
	cmd.Flags().StringVar(&opts.apiKey, flag.APIKey, "", usage.OpsgenieAPIKey)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.APIKey)

	return cmd
}

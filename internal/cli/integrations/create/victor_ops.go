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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

var victorOpsIntegrationType = "VICTOR_OPS"

type VictorOpsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	apiKey     string
	routingKey string
	store      store.IntegrationCreator
}

func (opts *VictorOpsOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var createTemplateVictorOps = "Victor Ops integration configured.\n"

func (opts *VictorOpsOpts) Run() error {
	r, err := opts.store.CreateIntegration(opts.ConfigProjectID(), victorOpsIntegrationType, opts.newVictorOpsIntegration())
	if err != nil {
		return err
	}
	return opts.Print(r)
}

func (opts *VictorOpsOpts) newVictorOpsIntegration() *atlasv2.ThirdPartyIntegration {
	return &atlasv2.ThirdPartyIntegration{
		Type:       &victorOpsIntegrationType,
		ApiKey:     &opts.apiKey,
		RoutingKey: &opts.routingKey,
	}
}

// atlas integration(s) create VICTOR_OPS --apiKey apiKey --routingKey routingKey [--projectId projectId].
func VictorOpsBuilder() *cobra.Command {
	opts := &VictorOpsOpts{}
	cmd := &cobra.Command{
		Use:     victorOpsIntegrationType,
		Aliases: []string{"victor_ops", "victorOps"},
		Short:   "Create or update a Splunk On-Call integration for your project.",
		Long: `VictorOps is now Splunk On-Call.
		
The requesting API key must have the Organization Owner or Project Owner role to configure an integration with Splunk On-Call.

` + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Annotations: map[string]string{
			"output": createTemplateVictorOps,
		},
		Example: `  # Integrate Splunk On-Call with Atlas using the routing key operations for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas integrations create VICTOR_OPS --apiKey a1a23bcdef45ghijk6789 --routingKey operations --projectId 5e2211c17a3e5a48f5497de3 --output json`,
		Args: require.NoArgs,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplateVictorOps),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.apiKey, flag.APIKey, "", usage.VictorOpsAPIKey)
	cmd.Flags().StringVar(&opts.routingKey, flag.RoutingKey, "", usage.RoutingKey)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.APIKey)
	_ = cmd.MarkFlagRequired(flag.RoutingKey)

	return cmd
}

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

var webhookIntegrationType = "WEBHOOK"

type WebhookOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	url    string
	secret string
	store  store.IntegrationCreator
}

func (opts *WebhookOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var createTemplateWebhook = "Webhook integration configured.\n"

func (opts *WebhookOpts) Run() error {
	r, err := opts.store.CreateIntegration(opts.ConfigProjectID(), webhookIntegrationType, opts.newWebhookIntegration())
	if err != nil {
		return err
	}
	return opts.Print(r)
}

func (opts *WebhookOpts) newWebhookIntegration() *atlasv2.ThirdPartyIntegration {
	return &atlasv2.ThirdPartyIntegration{
		Type:   &webhookIntegrationType,
		Url:    &opts.url,
		Secret: &opts.secret,
	}
}

// atlas integration(s) create WEBHOOK --url url --secret secret [--projectId projectId].
func WebhookBuilder() *cobra.Command {
	opts := &WebhookOpts{}
	cmd := &cobra.Command{
		Use:     webhookIntegrationType,
		Aliases: []string{"webhook"},
		Short:   "Create or update a webhook integration for your project.",
		Long: `The requesting API key must have the Organization Owner or Project Owner role to configure a webhook integration.

` + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Annotations: map[string]string{
			"output": createTemplateWebhook,
		},
		Example: `  # Integrate a webhook with Atlas that uses the secret mySecret for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas integrations create WEBHOOK --url http://9b4ac7aa.abc.io/payload --secret mySecret --projectId 5e2211c17a3e5a48f5497de3 --output json`,
		Args: require.NoArgs,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplateWebhook),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.url, flag.URL, "", usage.URL)
	cmd.Flags().StringVar(&opts.secret, flag.Secret, "", usage.Secret)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.URL)
	_ = cmd.MarkFlagRequired(flag.Secret)

	return cmd
}

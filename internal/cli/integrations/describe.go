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

package integrations

import (
	"context"
	"fmt"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	integrationType string
	store           store.IntegrationDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var describeTemplateSlack = `TYPE	API TOKEN	TEAM	CHANNEL
{{.Type}}	{{.ApiToken}}	{{.TeamName}}	{{ .ChannelName }}
`
var describeTemplateDatadogOpsGenie = `TYPE	API KEY	REGION
{{.Type}}	{{.ApiKey}}	{{.Region}}
`
var describeTemplateMicrosoftTeams = `TYPE	WEBHOOK URL
{{.Type}}	{{.MicrosoftTeamsWebhookUrl}}
`
var describeTemplateNewRelic = `TYPE	ACCOUNT ID	LICENSE KEY	WRITE TOKEN	READ TOKEN
{{.Type}}	{{.AccountId}}	{{.LicenseKey}}	{{.WriteToken}}	{{.ReadToken}}
`
var describeTemplatePagerDuty = `TYPE	SERVICE KEY
{{.Type}}	{{.ServiceKey}}
`
var describeTemplateVictorOps = `TYPE	API KEY	ROUTING KEY
{{.Type}}	{{.ApiKey}}	{{.RoutingKey}}
`
var describeTemplateWebhook = `TYPE	URL	SECRET
{{.Type}}	{{.Url}}	{{.Secret}}
`

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.Integration(opts.ConfigProjectID(), opts.integrationType)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *DescribeOpts) template() string {
	switch opts.integrationType {
	case "DATADOG":
		return describeTemplateDatadogOpsGenie
	case "MICROSOFT_TEAMS":
		return describeTemplateMicrosoftTeams
	case "NEW_RELIC":
		return describeTemplateNewRelic
	case "PAGER_DUTY":
		return describeTemplatePagerDuty
	case "VICTOR_OPS":
		return describeTemplateVictorOps
	case "OPS_GENIE":
		return describeTemplateDatadogOpsGenie
	case "WEBHOOK":
		return describeTemplateWebhook
	default:
		return describeTemplateSlack
	}
}

// atlas integration(s) describe <TYPE> [--projectId projectId].
func DescribeBuilder() *cobra.Command {
	opts := &DescribeOpts{}
	cmd := &cobra.Command{
		Use:   "describe <integrationType>",
		Short: "Return the details for the specified third-party integration for your project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:  require.ExactValidArgs(1),
		Example: `  # Return the JSON-formatted details for the Datadog integration for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas integrations describe DATADOG --projectId 5e2211c17a3e5a48f5497de3 --output json`,
		ValidArgs: []string{"PAGER_DUTY", "MICROSOFT_TEAMS", "SLACK", "DATADOG", "NEW_RELIC", "OPS_GENIE", "VICTOR_OPS", "WEBHOOK", "PROMETHEUS"},
		Annotations: map[string]string{
			"integrationTypeDesc": "Human-readable label that identifies the integrated service. Valid values are PAGER_DUTY, MICROSOFT_TEAMS, SLACK, DATADOG, NEW_RELIC, OPS_GENIE, VICTOR_OPS, WEBHOOK, PROMETHEUS.",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.integrationType = strings.ToUpper(args[0])
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), opts.template()),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}

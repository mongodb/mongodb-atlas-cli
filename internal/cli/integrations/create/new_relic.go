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

var newRelicIntegrationType = "NEW_RELIC"

type NewRelicOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	licenseKey string
	accountID  string
	writeToken string
	readToken  string
	store      store.IntegrationCreator
}

func (opts *NewRelicOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var createTemplateNewRelic = "New Relic integration configured.\n"

func (opts *NewRelicOpts) Run() error {
	r, err := opts.store.CreateIntegration(opts.ConfigProjectID(), newRelicIntegrationType, opts.newNewRelicIntegration())
	if err != nil {
		return err
	}
	return opts.Print(r)
}

func (opts *NewRelicOpts) newNewRelicIntegration() *atlasv2.ThirdPartyIntegration {
	return &atlasv2.ThirdPartyIntegration{
		Type:       &newRelicIntegrationType,
		LicenseKey: &opts.licenseKey,
		AccountId:  &opts.accountID,
		WriteToken: &opts.writeToken,
		ReadToken:  &opts.readToken,
	}
}

// atlas integration(s) create NEW_RELIC --licenceKey licenceKey --accountId accountId --writeToken writeToken --readToken readToken [--projectId projectId].
func NewRelicBuilder() *cobra.Command {
	opts := &NewRelicOpts{}
	cmd := &cobra.Command{
		Use:     "NEW_RELIC",
		Aliases: []string{"new_relic", "newRelic"},
		Short:   "Create or update the New Relic integration.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output": createTemplateNewRelic,
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplateNewRelic),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
		Deprecated: "The NEW_RELIC integration is deprecated and no longer supported",
	}

	cmd.Flags().StringVar(&opts.licenseKey, flag.LicenceKey, "", usage.LicenceKey)
	cmd.Flags().StringVar(&opts.accountID, flag.AccountID, "", usage.NewRelicAccountID)
	cmd.Flags().StringVar(&opts.writeToken, flag.WriteToken, "", usage.WriteToken)
	cmd.Flags().StringVar(&opts.readToken, flag.ReadToken, "", usage.ReadToken)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.LicenceKey)
	_ = cmd.MarkFlagRequired(flag.AccountID)
	_ = cmd.MarkFlagRequired(flag.WriteToken)
	_ = cmd.MarkFlagRequired(flag.ReadToken)

	return cmd
}

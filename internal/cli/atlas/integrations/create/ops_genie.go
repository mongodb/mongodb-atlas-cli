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
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const OpsGenieType = "OPS_GENIE"

type OpsGenieOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	apiKey string
	region string
	store  store.IntegrationCreator
}

func (opts *OpsGenieOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

var createTemplateOpsGenie = "Ops Genie integration configured.\n"

func (opts *OpsGenieOpts) Run() error {
	r, err := opts.store.CreateIntegration(opts.ConfigProjectID(), OpsGenieType, opts.newThirdPartyIntegration())
	if err != nil {
		return err
	}
	return opts.Print(r)
}

func (opts *OpsGenieOpts) newThirdPartyIntegration() *atlas.ThirdPartyIntegration {
	return &atlas.ThirdPartyIntegration{
		Type:   OpsGenieType,
		Region: opts.region,
		APIKey: opts.apiKey,
	}
}

// mongocli atlas integration(s) create OPS_GENIE --apiKey apiKey --region region [--projectId projectId]
func OpsGenieBuilder() *cobra.Command {
	opts := &OpsGenieOpts{}
	cmd := &cobra.Command{
		Use:     OpsGenieType,
		Aliases: []string{"ops_genie", "opsGenie", "og", "OG"},
		Short:   opsGenie,
		Args:    cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), createTemplateOpsGenie),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.region, flag.Region, "", usage.IntegrationRegion)
	cmd.Flags().StringVar(&opts.apiKey, flag.APIKey, "", usage.APIKey)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.APIKey)

	return cmd
}

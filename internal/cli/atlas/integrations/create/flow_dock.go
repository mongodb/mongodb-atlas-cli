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

const flowDockType = "FLOWDOCK"

type FlowDockOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	apiToken string
	flowName string
	orgName  string
	store    store.IntegrationCreator
}

func (opts *FlowDockOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

var createTemplateFlowDock = "Flow Dock integration configured.\n"

func (opts *FlowDockOpts) Run() error {
	r, err := opts.store.CreateIntegration(opts.ConfigProjectID(), flowDockType, opts.newFlowDockIntegration())
	if err != nil {
		return err
	}
	return opts.Print(r)
}

func (opts *FlowDockOpts) newFlowDockIntegration() *atlas.ThirdPartyIntegration {
	return &atlas.ThirdPartyIntegration{
		Type:     flowDockType,
		OrgName:  opts.orgName,
		FlowName: opts.flowName,
		APIToken: opts.apiToken,
	}
}

// mongocli atlas integration(s) create FLOWDOCK --apiToken apiToken --orgName orgName --flowName --flowName [--projectId projectId]
func FlowDockBuilder() *cobra.Command {
	opts := &FlowDockOpts{}
	cmd := &cobra.Command{
		Use:     flowDockType,
		Aliases: []string{"flowdock"},
		Short:   slack,
		Args:    cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), createTemplateFlowDock),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.flowName, flag.FlowName, "", usage.FlowName)
	cmd.Flags().StringVar(&opts.apiToken, flag.APIToken, "", usage.IntegrationAPIToken)
	cmd.Flags().StringVar(&opts.orgName, flag.OrgName, "", usage.OrgName)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.FlowName)
	_ = cmd.MarkFlagRequired(flag.APIToken)
	_ = cmd.MarkFlagRequired(flag.OrgName)

	return cmd
}

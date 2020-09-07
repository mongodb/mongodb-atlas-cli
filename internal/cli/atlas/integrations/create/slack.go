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

const slackType = "SLACK"

type SlackOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	apiToken    string
	teamName    string
	channelName string
	store       store.IntegrationCreator
}

func (opts *SlackOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

var createTemplateSlack = "Slack integration configured.\n"

func (opts *SlackOpts) Run() error {
	r, err := opts.store.CreateIntegration(opts.ConfigProjectID(), slackType, opts.newOpsGenieIntegration())
	if err != nil {
		return err
	}
	return opts.Print(r)
}

func (opts *SlackOpts) newOpsGenieIntegration() *atlas.ThirdPartyIntegration {
	return &atlas.ThirdPartyIntegration{
		Type:        slackType,
		ChannelName: opts.channelName,
		TeamName:    opts.teamName,
		APIToken:    opts.apiToken,
	}
}

// mongocli atlas integration(s) create slack --apiKey apiKey --region region [--projectId projectId]
func SlackBuilder() *cobra.Command {
	opts := &SlackOpts{}
	cmd := &cobra.Command{
		Use:     slackType,
		Aliases: []string{"slack"},
		Short:   slack,
		Args:    cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), createTemplateSlack),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.channelName, flag.ChannelName, "", usage.ChannelName)
	cmd.Flags().StringVar(&opts.apiToken, flag.APIToken, "", usage.SlackIntegrationAPIToken)
	cmd.Flags().StringVar(&opts.teamName, flag.TeamName, "", usage.TeamName)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.APIToken)
	_ = cmd.MarkFlagRequired(flag.TeamName)

	return cmd
}

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

package apikeys

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/output"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/ops-manager/opsmngr"
)

const createTemplate = "API Key '{{.Key}}' created.\n"

type CreateOpts struct {
	cli.GlobalOpts
	desc  string
	store store.AgentAPIKeyCreator
}

func (opts *CreateOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *CreateOpts) newAgentAPIKeysRequest() *opsmngr.AgentAPIKeysRequest {
	return &opsmngr.AgentAPIKeysRequest{
		Desc: opts.desc,
	}
}

func (opts *CreateOpts) Run() error {
	p, err := opts.store.CreateAgentAPIKey(opts.ConfigProjectID(), opts.newAgentAPIKeysRequest())

	if err != nil {
		return err
	}

	return output.Print(config.Default(), createTemplate, p)
}

// mongocli iam organizations|orgs apiKey(s)|apikeys create [--desc description][--projectId projectId]
func CreateBuilder() *cobra.Command {
	opts := new(CreateOpts)
	cmd := &cobra.Command{
		Use:   "create",
		Short: CreateAgentAPIKey,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.init)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.desc, flag.Description, "", usage.APIKeyDescription)

	cmd.Flags().StringVar(&opts.OrgID, flag.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flag.Description)

	return cmd
}

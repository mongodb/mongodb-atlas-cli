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

package livemigration

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas/mongodbatlas"
)

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store      		store.LinkTokenCreator
	orgId      		string
	accessListIPs []string
}

func (opts *CreateOpts) initStore() error {
	var err error
	opts.store, err = store.New(store.AuthenticatedPreset(config.Default()))
	return err
}

// TODO: What do we want to see here
var createTemplate = "Link Token '{{.Token}}' created.\n"

func (opts *CreateOpts) Run() error {
	createRequest := opts.newTokenCreateRequest()

	r, err := opts.store.CreateLinkToken(opts.orgId, createRequest)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newTokenCreateRequest() *mongodbatlas.TokenCreateRequest {
	return &mongodbatlas.TokenCreateRequest{
		AccessListIPs: opts.accessListIPs,
	}
}

// mongocli atlas cloudMigration|liveMigration|lm link create [--accessListIps accessListIps] [--orgId orgId]
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create one new link-token.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.orgId, flag.OrgID, "", usage.OrgID)
	cmd.Flags().StringArrayVar(&opts.accessListIPs, flag.AccessListIPs, []string{}, usage.AccessListCIDREntries)

	_ = cmd.MarkFlagRequired(flag.OrgID)

	return cmd
}

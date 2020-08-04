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

package whitelist

import (
	"fmt"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/output"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const createTemplate = "Created new whitelist entry(s).\n"

type CreateOpts struct {
	cli.GlobalOpts
	apyKey string
	ips    []string
	cidrs  []string
	store  store.OrganizationAPIKeyWhitelistCreator
}

func (opts *CreateOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *CreateOpts) newWhitelistAPIKeysReq() ([]*atlas.WhitelistAPIKeysReq, error) {
	var whitelistRep []*atlas.WhitelistAPIKeysReq
	if len(opts.ips) == 0 && len(opts.cidrs) == 0 {
		return nil, fmt.Errorf("at least one between --ip and --cidr must be used")
	}
	for _, v := range opts.ips {
		whitelist := atlas.WhitelistAPIKeysReq{
			IPAddress: v,
		}
		whitelistRep = append(whitelistRep, &whitelist)
	}

	for _, v := range opts.cidrs {
		whitelist := atlas.WhitelistAPIKeysReq{
			CidrBlock: v,
		}
		whitelistRep = append(whitelistRep, &whitelist)
	}

	return whitelistRep, nil
}

func (opts *CreateOpts) Run() error {
	whitelistReq, err := opts.newWhitelistAPIKeysReq()
	if err != nil {
		return err
	}

	p, err := opts.store.CreateOrganizationAPIKeyWhite(opts.ConfigOrgID(), opts.apyKey, whitelistReq)

	if err != nil {
		return err
	}

	return output.Print(config.Default(), createTemplate, p)
}

// mongocli iam organizations|orgs apiKey(s)|apikeys whitelist|ipwhitelist create [--keyId keyId] [--orgId orgId] [--ip ip] [--cidr cidr]
func CreateBuilder() *cobra.Command {
	opts := new(CreateOpts)
	cmd := &cobra.Command{
		Use:   "create",
		Short: description.CreateOrganizationsAPIKey,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunEOrg(opts.init)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.apyKey, flag.APIKey, "", usage.APIKey)
	cmd.Flags().StringSliceVar(&opts.cidrs, flag.CIDR, []string{}, usage.WhitelistCIDREntry)
	cmd.Flags().StringSliceVar(&opts.ips, flag.IP, []string{}, usage.WhitelistIPEntry)

	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)

	_ = cmd.MarkFlagRequired(flag.APIKey)

	return cmd
}

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

package globalwhitelist

import (
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/output"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/ops-manager/opsmngr"
)

const createTemplate = "Created new whitelist entry(s).\n"

type CreateOpts struct {
	description string
	cidrs       []string
	store       store.GlobalAPIKeyWhitelistCreator
}

func (opts *CreateOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *CreateOpts) newWhitelistAPIKeysReq() []*opsmngr.WhitelistAPIKeysReq {
	whitelistRep := make([]*opsmngr.WhitelistAPIKeysReq, len(opts.cidrs))
	for i := range opts.cidrs {
		whitelist := &opsmngr.WhitelistAPIKeysReq{
			CidrBlock:   opts.cidrs[i],
			Description: opts.description,
		}
		whitelistRep[i] = whitelist
	}

	return whitelistRep
}

func (opts *CreateOpts) Run() error {
	whitelistReq := opts.newWhitelistAPIKeysReq()
	p, err := opts.store.CreateGlobalAPIKeyWhite(whitelistReq)

	if err != nil {
		return err
	}

	return output.Print(config.Default(), createTemplate, p)
}

// mongocli iam globalWhitelist(s) create [--cidr cidr][--desc description]
func CreateBuilder() *cobra.Command {
	opts := new(CreateOpts)
	cmd := &cobra.Command{
		Use:   "create",
		Short: createWhitelist,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringSliceVar(&opts.cidrs, flag.CIDR, []string{}, usage.WhitelistCIDREntry)
	cmd.Flags().StringVar(&opts.description, flag.Description, "", usage.WhitelistIPEntry)

	_ = cmd.MarkFlagRequired(flag.CIDR)
	_ = cmd.MarkFlagRequired(flag.Description)
	return cmd
}

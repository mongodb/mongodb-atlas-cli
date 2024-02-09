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

package globalaccesslists

import (
	"context"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/ops-manager/opsmngr"
)

const createTemplate = "Global access list entry '{{.ID}}' created.\n"

type CreateOpts struct {
	cli.OutputOpts
	description string
	cidr        string
	store       store.GlobalAPIKeyWhitelistCreator
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *CreateOpts) newWhitelistAPIKeysReq() *opsmngr.WhitelistAPIKeysReq {
	entry := &opsmngr.WhitelistAPIKeysReq{
		CidrBlock:   opts.cidr,
		Description: opts.description,
	}
	return entry
}

func (opts *CreateOpts) Run() error {
	entry := opts.newWhitelistAPIKeysReq()
	r, err := opts.store.CreateGlobalAPIKeyWhitelist(entry)

	if err != nil {
		return err
	}

	return opts.Print(r)
}

// mongocli iam globalAccessList(s) create [--cidr cidr][--desc description].
func CreateBuilder() *cobra.Command {
	opts := new(CreateOpts)
	opts.Template = createTemplate
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create an IP access list entry for your global API key.",
		Example: `  # Create an access list entry for your global API key to allow access from 192.0.2.0/24:
  mongocli iam globalAccessLists create --cidr 192.0.2.0/24 --desc "My Global IP" --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			return opts.initStore(cmd.Context())()
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.cidr, flag.CIDR, "", usage.AccessListCIDREntry)
	cmd.Flags().StringVar(&opts.description, flag.Description, "", usage.APIAccessListIPEntry)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.CIDR)
	_ = cmd.MarkFlagRequired(flag.Description)
	return cmd
}

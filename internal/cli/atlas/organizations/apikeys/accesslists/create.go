// Copyright 2023 MongoDB Inc
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

package accesslists

import (
	"context"
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	storeHelper "github.com/mongodb/mongodb-atlas-cli/internal/store"
	store "github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20231001002/admin"
)

const createTemplate = "Created new access list entry(s).\n"

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	apyKey    string
	ips       []string
	cidrs     []string
	currentIP bool
	store     store.OrganizationAPIKeyAccessListCreator
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *CreateOpts) newAccessListAPIKeysReq() (*[]admin.UserAccessList, error) {
	req := make([]admin.UserAccessList, 0, len(opts.ips)+len(opts.cidrs))
	if len(opts.ips) == 0 && len(opts.cidrs) == 0 {
		return nil, fmt.Errorf("either --%s, --%s must be set", flag.IP, flag.CIDR)
	}
	for i := range opts.ips {
		req = append(req, admin.UserAccessList{
			IpAddress: &opts.ips[i],
		})
	}

	for i := range opts.cidrs {
		req = append(req, admin.UserAccessList{
			CidrBlock: &opts.cidrs[i],
		})
	}

	return &req, nil
}

func (opts *CreateOpts) Run() error {
	req, err := opts.newAccessListAPIKeysReq()
	if err != nil {
		return err
	}

	params := &admin.CreateApiKeyAccessListApiParams{
		OrgId:          opts.ConfigOrgID(),
		ApiUserId:      opts.apyKey,
		UserAccessList: req,
	}

	result, err := opts.store.CreateOrganizationAPIKeyAccessList(params)
	if err != nil {
		return err
	}
	return opts.Print(result)
}

// atlas organizations|orgs apiKey(s)|apikeys accessList create [--apiKey keyId] [--orgId orgId] [--ip ip] [--cidr cidr].
func CreateBuilder() *cobra.Command {
	opts := new(CreateOpts)
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create an IP access list entry for your API Key.",
		Long: fmt.Sprintf(`To view possible values for the apiKey option, run %s organizations apiKeys list.

`+fmt.Sprintf(usage.RequiredRole, "Read Write"), cli.ExampleAtlasEntryPoint()),
		Example: fmt.Sprintf(`  # Create access list entries for two IP addresses for the API key with the ID 5f24084d8dbffa3ad3f21234 in the organization with the ID 5a1b39eec902201990f12345:
  %s organizations apiKeys accessLists create --apiKey 5f24084d8dbffa3ad3f21234 --cidr 192.0.2.0/24,198.51.100.0/24 --orgId 5a1b39eec902201990f12345 --output json`, cli.ExampleAtlasEntryPoint()),
		Annotations: map[string]string{
			"output": createTemplate,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.currentIP {
				publicIP := storeHelper.IPAddress()
				if publicIP == "" {
					return errors.New("unable to find your public IP address. Specify the public IP address for this command")
				}
				opts.ips = append(opts.ips, publicIP)
			}
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.apyKey, flag.APIKey, "", usage.APIKey)
	cmd.Flags().StringSliceVar(&opts.cidrs, flag.CIDR, []string{}, usage.AccessListCIDREntry)
	cmd.Flags().StringSliceVar(&opts.ips, flag.IP, []string{}, usage.APIAccessListIPEntry)

	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	cmd.Flags().BoolVar(&opts.currentIP, flag.CurrentIP, false, usage.CurrentIP)

	_ = cmd.MarkFlagRequired(flag.APIKey)

	return cmd
}

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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
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

func (opts *CreateOpts) newAccessListAPIKeysReq() (*[]admin.UserAccessListRequest, error) {
	req := make([]admin.UserAccessListRequest, 0, len(opts.ips)+len(opts.cidrs))
	if len(opts.ips) == 0 && len(opts.cidrs) == 0 {
		return nil, fmt.Errorf("either --%s, --%s must be set", flag.IP, flag.CIDR)
	}
	for i := range opts.ips {
		req = append(req, admin.UserAccessListRequest{
			IpAddress: &opts.ips[i],
		})
	}

	for i := range opts.cidrs {
		req = append(req, admin.UserAccessListRequest{
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
		OrgId:                 opts.ConfigOrgID(),
		ApiUserId:             opts.apyKey,
		UserAccessListRequest: req,
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
		Long: `To view possible values for the apiKey option, run atlas organizations apiKeys list.

` + fmt.Sprintf(usage.RequiredRole, "Read Write"),
		Example: `  # Create access list entries for two IP addresses for the API key with the ID 5f24084d8dbffa3ad3f21234 in the organization with the ID 5a1b39eec902201990f12345:
  atlas organizations apiKeys accessLists create --apiKey 5f24084d8dbffa3ad3f21234 --cidr 192.0.2.0/24,198.51.100.0/24 --orgId 5a1b39eec902201990f12345 --output json`,
		Annotations: map[string]string{
			"output": createTemplate,
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			if opts.currentIP {
				publicIP := store.IPAddress()
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

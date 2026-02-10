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

package link

import (
	"context"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312013/admin"
)

var createTemplate = "Link-token '{{.LinkToken}}' successfully created.\n"

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=create_mock_test.go -package=link . TokenCreator

type TokenCreator interface {
	CreateLinkToken(string, *atlasv2.TargetOrgRequest) (*atlasv2.TargetOrg, error)
}

type CreateOpts struct {
	cli.OrgOpts
	cli.OutputOpts
	store        TokenCreator
	accessListIP []string
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *CreateOpts) Run() error {
	createRequest := opts.newTokenCreateRequest()

	r, err := opts.store.CreateLinkToken(opts.ConfigOrgID(), createRequest)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newTokenCreateRequest() *atlasv2.TargetOrgRequest {
	return &atlasv2.TargetOrgRequest{
		AccessListIps: &opts.accessListIP,
	}
}

// atlas liveMigrations|lm link create [--accessListIp accessListIp] [--orgId orgId].
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new link-token for a push live migration.",
		Long:  `To migrate using scripts, use mongomirror instead of the Atlas CLI. To learn more about mongomirror, see https://www.mongodb.com/docs/atlas/reference/mongomirror/.`,
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
			return opts.Run()
		},
	}
	opts.AddOrgOptFlags(cmd)
	cmd.Flags().StringSliceVar(&opts.accessListIP, flag.AccessListIP, []string{}, usage.LinkTokenAccessListCIDREntries)
	opts.AddOutputOptFlags(cmd)

	return cmd
}

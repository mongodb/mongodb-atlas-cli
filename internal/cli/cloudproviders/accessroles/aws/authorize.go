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

package aws

import (
	"context"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312010/admin"
)

const authorizeTemplate = "AWS IAM role '{{.RoleId}} successfully authorized.\n"

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=authorize_mock_test.go -package=aws . CloudProviderAccessRoleAuthorizer

// CloudProviderAccessRoleAuthorizer encapsulates the logic to manage different cloud providers.
type CloudProviderAccessRoleAuthorizer interface {
	AuthorizeCloudProviderAccessRole(string, string, *atlasv2.CloudProviderAccessRoleRequestUpdate) (*atlasv2.CloudProviderAccessRole, error)
}

type AuthorizeOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	store             CloudProviderAccessRoleAuthorizer
	roleID            string
	IAMAssumedRoleARN string
}

func (opts *AuthorizeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *AuthorizeOpts) Run() error {
	r, err := opts.store.AuthorizeCloudProviderAccessRole(opts.ConfigProjectID(), opts.roleID, opts.newCloudProviderAuthorizationRequest())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *AuthorizeOpts) newCloudProviderAuthorizationRequest() *atlasv2.CloudProviderAccessRoleRequestUpdate {
	return &atlasv2.CloudProviderAccessRoleRequestUpdate{
		ProviderName:      provider,
		IamAssumedRoleArn: &opts.IAMAssumedRoleARN,
	}
}

// atlas cloudProvider aws accessRoles authorize <roleId> --iamAssumedRoleArn iamAssumedRoleArn [--projectId projectId].
func AuthorizeBuilder() *cobra.Command {
	opts := &AuthorizeOpts{}
	cmd := &cobra.Command{
		Use:   "authorize <roleId>",
		Short: "Authorize an AWS IAM role.",
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"roleIdDesc": "Unique ID of the role to authorize.",
			"output":     authorizeTemplate,
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), authorizeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.roleID = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.IAMAssumedRoleARN, flag.IAMAssumedRoleARN, "", usage.IAMAssumedRoleARN)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagFilename(flag.IAMAssumedRoleARN)
	return cmd
}

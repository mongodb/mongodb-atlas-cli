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

package aws

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const enableTemplate = "AWS IAM role '{{.RoleID}} successfully enabled.\n"

type EnableOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store             store.CloudProviderAccessRoleEnabler
	roleID            string
	IAMAssumedRoleARN string
}

func (opts *EnableOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *EnableOpts) Run() error {
	r, err := opts.store.EnableCloudProviderAccessRole(opts.ConfigProjectID(), opts.roleID, opts.newCloudProviderAuthorizationRequest())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *EnableOpts) newCloudProviderAuthorizationRequest() *atlas.CloudProviderAuthorizationRequest {
	return &atlas.CloudProviderAuthorizationRequest{
		ProviderName:      provider,
		IAMAssumedRoleARN: opts.IAMAssumedRoleARN,
	}
}

// mongocli atlas cloudProvider aws accessRoles enable --roleId roleId --iamAssumedRoleArn iamAssumedRoleArn [--projectId projectId]
func EnableBuilder() *cobra.Command {
	opts := &EnableOpts{}
	cmd := &cobra.Command{
		Use:   "enable",
		Short: enable,
		Args:  require.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), enableTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.roleID, flag.RoleID, "", usage.RoleID)
	cmd.Flags().StringVar(&opts.IAMAssumedRoleARN, flag.IAMAssumedRoleARN, "", usage.IAMAssumedRoleARN)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagFilename(flag.RoleID)
	_ = cmd.MarkFlagFilename(flag.IAMAssumedRoleARN)
	return cmd
}

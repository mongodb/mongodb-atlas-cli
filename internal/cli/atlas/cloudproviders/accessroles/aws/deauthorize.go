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
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const disableTemplate = "AWS IAM role successfully deauthorized.\n"

type DeauthorizeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store  store.CloudProviderAccessRoleDeauthorizer
	roleID string
}

func (opts *DeauthorizeOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *DeauthorizeOpts) Run() error {
	err := opts.store.DeauthorizeCloudProviderAccessRoles(opts.newCloudProviderDeauthorizationRequest())
	if err != nil {
		return err
	}

	return opts.Print(nil)
}

func (opts *DeauthorizeOpts) newCloudProviderDeauthorizationRequest() *atlas.CloudProviderDeauthorizationRequest {
	return &atlas.CloudProviderDeauthorizationRequest{
		ProviderName: provider,
		GroupID:      opts.ConfigProjectID(),
		RoleID:       opts.roleID,
	}
}

// mongocli atlas cloudProvider aws accessRoles deauthorize <roleId> [--projectId projectId]
func DeauthorizeBuilder() *cobra.Command {
	opts := &DeauthorizeOpts{}
	cmd := &cobra.Command{
		Use:   "deauthorize",
		Short: deauthorize,
		Args:  require.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), disableTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.roleID = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}

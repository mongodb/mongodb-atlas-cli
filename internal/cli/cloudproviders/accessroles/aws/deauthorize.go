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
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

const (
	deauthorizeSuccess  = "AWS IAM role successfully deauthorized.\n"
	deauthorizeFail     = "AWS IAM role not deauthorized.\n"
	confirmationMessage = "Are you sure you want to deauthorize: %s"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=deauthorize_mock_test.go -package=aws . CloudProviderAccessRoleDeauthorizer

type CloudProviderAccessRoleDeauthorizer interface {
	DeauthorizeCloudProviderAccessRoles(string, string, string) error
}

type DeauthorizeOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	*cli.DeleteOpts
	store CloudProviderAccessRoleDeauthorizer
}

func (opts *DeauthorizeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeauthorizeOpts) Run() error {
	if !opts.Confirm {
		return opts.Print(deauthorizeFail)
	}

	err := opts.store.DeauthorizeCloudProviderAccessRoles(opts.ConfigProjectID(), provider, opts.Entry)
	if err != nil {
		return err
	}

	return opts.Print(deauthorizeSuccess)
}

// atlas cloudProvider aws accessRoles deauthorize <roleId> [--projectId projectId].
func DeauthorizeBuilder() *cobra.Command {
	opts := &DeauthorizeOpts{
		DeleteOpts: cli.NewDeleteOpts(deauthorizeSuccess, deauthorizeFail),
	}
	cmd := &cobra.Command{
		Use:   "deauthorize <roleId>",
		Short: "Deauthorize an AWS IAM role.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"roleIdDesc": "Unique ID of the role to authorize.",
			"output":     deauthorizeSuccess,
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(opts.ValidateProjectID, opts.initStore(cmd.Context()))
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.Entry = args[0]
			if err := opts.PromptWithMessage(confirmationMessage); err != nil {
				return err
			}
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	opts.AddProjectOptsFlags(cmd)

	return cmd
}

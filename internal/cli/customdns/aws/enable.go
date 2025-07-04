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
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=enable_mock_test.go -package=aws . CustomDNSEnabler

type CustomDNSEnabler interface {
	EnableCustomDNS(string) (*atlasv2.AWSCustomDNSEnabled, error)
}

type EnableOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	store CustomDNSEnabler
}

var enableTemplate = "DNS configuration enabled.\n"

func (opts *EnableOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *EnableOpts) Run() error {
	r, err := opts.store.EnableCustomDNS(opts.ConfigProjectID())
	if err != nil {
		return err
	}
	return opts.Print(r)
}

// atlas customDns aws enable [--projectId projectId].
func EnableBuilder() *cobra.Command {
	opts := &EnableOpts{}
	cmd := &cobra.Command{
		Use:   "enable",
		Short: "Enable the custom DNS configuration of an Atlas cluster deployed to AWS in the specified project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Annotations: map[string]string{
			"output": enableTemplate,
		},
		Example: `  # Enable the custom DNS configuration deployed to AWS in the project with ID 618d48e05277a606ed2496fe:		
  atlas customDns aws enable --projectId 618d48e05277a606ed2496fe `,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), enableTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}

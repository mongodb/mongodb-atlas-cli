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

	"github.com/andreaangiolillo/mongocli-test/internal/cli"
	"github.com/andreaangiolillo/mongocli-test/internal/config"
	"github.com/andreaangiolillo/mongocli-test/internal/flag"
	"github.com/andreaangiolillo/mongocli-test/internal/store"
	"github.com/andreaangiolillo/mongocli-test/internal/usage"
	"github.com/spf13/cobra"
)

type EnableOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store store.CustomDNSEnabler
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

// mongocli atlas customDns aws enable [--projectId projectId].
func EnableBuilder() *cobra.Command {
	opts := &EnableOpts{}
	cmd := &cobra.Command{
		Use:   "enable",
		Short: "Enable the custom DNS configuration of an Atlas cluster deployed to AWS in the specified project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Annotations: map[string]string{
			"output": enableTemplate,
		},
		Example: fmt.Sprintf(`  # Enable the custom DNS configuration deployed to AWS in the project with ID 618d48e05277a606ed2496fe:		
  %s customDns aws enable --projectId 618d48e05277a606ed2496fe `, cli.ExampleAtlasEntryPoint()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), enableTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}

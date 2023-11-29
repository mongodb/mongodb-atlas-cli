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

package ondemand

import (
	"context"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	store "github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	*cli.DeleteOpts
	store store.CompliancePolicyOnDemandPolicyDeleter
}

const deleteTemplate = `ID	RETENTION
{{if .OnDemandPolicyItem}}{{.OnDemandPolicyItem.Id}}	{{.OnDemandPolicyItem.RetentionValue}} {{.OnDemandPolicyItem.RetentionUnit}}{{end}}
`
const confirmationMessage = "Are you sure you want to delete the ondemand policy item of the backup compliance policy"

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	raw, err := opts.store.DeleteOnDemandPolicy(opts.ConfigProjectID())
	if err != nil {
		return err
	}

	return opts.Print(raw)
}

func DeleteBuilder() *cobra.Command {
	opts := new(DeleteOpts)
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes the ondemand policy item of the backup compliance policy for your project.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), deleteTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PromptWithMessage(confirmationMessage); err != nil {
				return err
			}

			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}

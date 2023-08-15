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

package compliancepolicy

import (
	"context"
	"errors"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201004/admin"
)

type CopyProtectionOpts struct {
	cli.GlobalOpts
	cli.WatchOpts
	policy *atlasv2.DataProtectionSettings
	store  store.CompliancePolicy
	enable bool
}

const (
	enable  = "enable"
	disable = "disable"
	active  = "ACTIVE"
)

var copyProtectionTemplate = `Copy protection has been set to: {{.CopyProtectionEnabled}}
`

func (opts *CopyProtectionOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *CopyProtectionOpts) PreRun() error {
	currentPolicy, err := opts.store.DescribeCompliancePolicy(opts.ConfigProjectID())
	if err != nil {
		return err
	}
	opts.policy = currentPolicy
	return nil
}

func (opts *CopyProtectionOpts) copyProtectionWatcher() (bool, error) {
	res, err := opts.store.DescribeCompliancePolicy(opts.ConfigProjectID())
	opts.policy = res
	if err != nil {
		return false, err
	}
	if res.GetState() == "" {
		return false, errors.New("could not access State field")
	}
	return (res.GetState() == active), nil
}

func (opts *CopyProtectionOpts) Run() error {
	opts.policy.SetCopyProtectionEnabled(opts.enable)
	_, err := opts.store.UpdateCompliancePolicy(opts.ConfigProjectID(), opts.policy)
	if err != nil {
		return err
	}
	if err := opts.Watch(opts.copyProtectionWatcher); err != nil {
		return err
	}

	return opts.Print(opts.policy)
}

func CopyProtectionBuilder() *cobra.Command {
	opts := new(CopyProtectionOpts)
	use := "copyProtection"
	cmd := &cobra.Command{
		Use:       use,
		Aliases:   cli.GenerateAliases(use),
		Args:      require.ExactValidArgs(1),
		ValidArgs: []string{enable, disable},
		Short:     "Enable or disable copyprotection of the backup compliance policy for your project.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if args[0] == enable {
				opts.enable = true
			} else {
				opts.enable = false
			}
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), copyProtectionTemplate),
				opts.PreRun,
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

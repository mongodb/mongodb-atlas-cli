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

package copyprotection

import (
	"context"
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type DisableOpts struct {
	cli.GlobalOpts
	cli.WatchOpts
	policy *atlasv2.DataProtectionSettings20231001
	store  store.CompliancePolicyCopyProtectionDisabler
}

var disableWatchTemplate = `Copy protection has been disabled.
`

var disableTemplate = `Copy protection is being disabled.
`

func (opts *DisableOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DisableOpts) watcher() (any, bool, error) {
	res, err := opts.store.DescribeCompliancePolicy(opts.ConfigProjectID())
	if err != nil {
		return nil, false, err
	}
	opts.policy = res
	if res.GetState() == "" {
		return nil, false, errors.New("could not access State field")
	}
	return nil, res.GetState() == active, nil
}

func (opts *DisableOpts) Run() error {
	res, err := opts.store.DisableCopyProtection(opts.ConfigProjectID())
	if err != nil {
		return fmt.Errorf("couldn't disable copy protection: %w", err)
	}
	opts.policy = res
	if opts.EnableWatch {
		if _, err := opts.Watch(opts.watcher); err != nil {
			return err
		}
		opts.Template = disableWatchTemplate
	}
	return opts.Print(opts.policy)
}

func DisableBuilder() *cobra.Command {
	opts := new(DisableOpts)
	use := "disable"
	cmd := &cobra.Command{
		Use:   use,
		Short: "Disable copy protection of the backup compliance policy for your project.",
		Args:  require.NoArgs,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), disableTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	cmd.Flags().BoolVarP(&opts.EnableWatch, flag.EnableWatch, flag.EnableWatchShort, false, usage.EnableWatchDefault)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}

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

package pointintimerestore

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

type EnableOpts struct {
	cli.GlobalOpts
	cli.WatchOpts
	policy            *atlasv2.DataProtectionSettings20231001
	store             store.CompliancePolicyPointInTimeRestoresEnabler
	restoreWindowDays int
}

var enableWatchTemplate = `Point-in-Time restores have been enabled with a restore window of {{.RestoreWindowDays}} days.
`

var enableTemplate = `Point-in-Time restores is being enabled with a restore window of {{.RestoreWindowDays}} days.
`

func (opts *EnableOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *EnableOpts) watcher() (any, bool, error) {
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

func (opts *EnableOpts) Run() error {
	res, err := opts.store.EnablePointInTimeRestore(opts.ConfigProjectID(), opts.restoreWindowDays)
	if err != nil {
		return fmt.Errorf("couldn't enable point in time restore: %w", err)
	}
	opts.policy = res
	if opts.EnableWatch {
		if _, err := opts.Watch(opts.watcher); err != nil {
			return err
		}
		opts.Template = enableWatchTemplate
	}
	return opts.Print(opts.policy)
}

func (opts *EnableOpts) validateRestoreWindowDays() error {
	if opts.restoreWindowDays <= 0 {
		return errors.New("restoreWindowDays must be a positive, non-zero integer")
	}
	return nil
}

func EnableBuilder() *cobra.Command {
	opts := new(EnableOpts)
	use := "enable"
	cmd := &cobra.Command{
		Use:   use,
		Args:  require.NoArgs,
		Short: "Enable Point-in-Time restores of the backup compliance policy for your project.",
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.validateRestoreWindowDays,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), enableTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	cmd.Flags().BoolVarP(&opts.EnableWatch, flag.EnableWatch, flag.EnableWatchShort, false, usage.EnableWatchDefault)
	cmd.Flags().IntVar(&opts.restoreWindowDays, flag.RestoreWindowDays, 0, usage.RestoreWindowDays)
	_ = cmd.MarkFlagRequired(flag.RestoreWindowDays)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}

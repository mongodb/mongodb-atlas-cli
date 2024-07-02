// Copyright 2024 MongoDB Inc
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

package auditing

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	enabled                   bool
	auditAuthorizationSuccess bool
	filename                  string
	auditFilter               string
	fs                        afero.Fs
	store                     store.AuditingUpdater
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var updateTemplate = "Auditing configuration successfully updated.\n"

func (opts *UpdateOpts) Run() error {
	auditLog, err := opts.newAuditLog()
	if err != nil {
		return err
	}
	r, err := opts.store.UpdateAuditingConfig(opts.ConfigProjectID(), auditLog)
	if err != nil {
		return err
	}
	return opts.Print(r)
}

func (opts *UpdateOpts) newAuditLog() (*atlasv2.AuditLog, error) {
	if opts.filename != "" {
		fileContent, err := file.LoadFile(opts.fs, opts.filename)
		if err != nil {
			return nil, err
		}
		opts.auditFilter = string(fileContent)
	}

	return &atlasv2.AuditLog{
		Enabled:                   &opts.enabled,
		AuditAuthorizationSuccess: &opts.auditAuthorizationSuccess,
		AuditFilter:               &opts.auditFilter,
	}, nil
}

// UpdateBuilder atlas auditing update [--file filter.json] [--auditFilter "{"atype": "authenticate"}"] [--auditAuthorizationSuccess] [--enabled] --projectId projectId.
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Updates the auditing configuration for the specified project",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:  require.NoArgs,
		Annotations: map[string]string{
			"output": updateTemplate,
		},
		Example: `  # Audit all authentication events for known users:
  atlas auditing update --auditFilter '{"atype": "authenticate"}'

  # Audit all authentication events for known user via a configuration file:
  atlas auditing update -f filter.json
`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVar(&opts.auditFilter, flag.AuditFilter, "", usage.AuditFilter)
	cmd.Flags().BoolVar(&opts.enabled, flag.Enabled, false, usage.EnabledAuditing)
	cmd.Flags().StringVarP(&opts.filename, flag.File, flag.FileShort, "", usage.AuditingFilename)
	cmd.Flags().BoolVar(&opts.auditAuthorizationSuccess, flag.AuditAuthorizationSuccess, false, usage.AuditAuthorizationSuccess)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagFilename(flag.File)

	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.AuditFilter)
	cmd.MarkFlagsOneRequired(flag.File, flag.AuditFilter)

	return cmd
}

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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	enabled                   bool
	auditAuthorizationSuccess bool
	filename                  string
	auditFilter               string
	store                     store.AuditingDescriber
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var updateTemplate = `AUDIT AUTHORIZATION SUCCESS	AUDIT FILTER	CONFIGURATION TYPE	ENABLED
{{.AuditAuthorizationSuccess}}	{{.AuditFilter}}	{{.ConfigurationType}}	{{.Enabled}}
`

func (opts *UpdateOpts) Run() error {
	print(opts.auditFilter)

	r, err := opts.store.Auditing(opts.ConfigProjectID())
	if err != nil {
		return err
	}
	return opts.Print(r)
}

// UpdateBuilder atlas auditing update [--file filter.json] [--auditFilter "{"atype": "authenticate"}"] [--auditAuthorizationSuccess] [--enabled] --projectId projectId.
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Updates the auditing configuration for the specified project",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:  require.NoArgs,
		Annotations: map[string]string{
			"output": updateTemplate,
		},
		Example: `  # Audit all authentication events for known users:
  atlas auditing update --auditFilter "{"atype": "authenticate"}"
  # Audit all authentication events for known user via a configuration file:
  atlas auditing update -f filter.json
`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
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
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagFilename(flag.File)

	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.AuditFilter)

	return cmd
}

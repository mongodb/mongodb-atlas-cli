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

package auditing

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=describe_mock_test.go -package=auditing . Describer

type Describer interface {
	Auditing(string) (*atlasv2.AuditLog, error)
}

type DescribeOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	store Describer
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var describeTemplate = `AUDIT AUTHORIZATION SUCCESS	AUDIT FILTER	CONFIGURATION TYPE	ENABLED
{{.AuditAuthorizationSuccess}}	{{.AuditFilter}}	{{.ConfigurationType}}	{{.Enabled}}
`

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.Auditing(opts.ConfigProjectID())
	if err != nil {
		return err
	}
	return opts.Print(r)
}

// DescribeBuilder atlas auditing describe --projectId projectId.
func DescribeBuilder() *cobra.Command {
	opts := &DescribeOpts{}
	cmd := &cobra.Command{
		Use:     "describe",
		Aliases: []string{"get"},
		Short:   "Returns the auditing configuration for the specified project.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output": describeTemplate,
		},
		Example: `  # Return the JSON-formatted details for the default project:
  atlas auditing describe --output json`,
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

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}

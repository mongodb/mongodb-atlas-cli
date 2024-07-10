// Copyright 2021 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package customdbroles

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

const describeTemplate = `NAME	ACTION	DB	COLLECTION	CLUSTER {{- $roleName := .RoleName }} {{range valueOrEmptySlice .Actions}} 
{{- $actionName := .Action }} {{- range valueOrEmptySlice .Resources}}
{{ $roleName }}	{{ $actionName }}{{if .Db }}	{{ .Db }}{{else}}	N/A{{end}}{{if .Collection }}	{{ .Collection }}{{else if .Cluster}}	N/A{{else}}	ALL COLLECTIONS{{end}}{{if .Cluster}}	{{ .Cluster }}{{else}}	N/A	{{end}}{{end}}{{end}}
`

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store    store.DatabaseRoleDescriber
	roleName string
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.DatabaseRole(opts.ConfigProjectID(), opts.roleName)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas dbroles(s) describe <roleName> --projectId projectId.
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:   "describe <roleName>",
		Short: "Return a single custom database role for your project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"roleNameDesc": "Name of the custom role to retrieve.",
		},
		Aliases: []string{"get"},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.roleName = args[0]
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
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}

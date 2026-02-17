// Copyright 2021 MongoDB Inc
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

package customdbroles

import (
	"context"
	"fmt"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312014/admin"
)

const listTemplate = `NAME	ACTION	INHERITED ROLES	DB	COLLECTION	CLUSTER{{range valueOrEmptySlice .}}{{- $roleName := .RoleName }} {{range valueOrEmptySlice .Actions}} 
{{- $actionName := .Action }} {{- range valueOrEmptySlice .Resources}}
{{ $roleName }}	{{ $actionName }}	N/A{{if .Db }}	{{ .Db }}{{else}}	N/A{{end}}{{if .Collection }}	{{ .Collection }}{{else if .Cluster}}	N/A{{else}}	ALL COLLECTIONS{{end}}{{if .Cluster}}	{{ .Cluster }}{{else}}	N/A	{{end}}{{end}}{{end}}{{range valueOrEmptySlice .InheritedRoles}}
{{ $roleName }}	N/A	{{ .Role }}	{{ .Db}}	N/A	N/A{{end}}{{end}}
`

const deprecatedFlagMessage = "--pageNum and --ItemsPerPage are not supported by customdbroles list"

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=list_mock_test.go -package=customdbroles . DatabaseRoleLister

type DatabaseRoleLister interface {
	DatabaseRoles(string) ([]atlasv2.UserCustomDBRole, error)
}

type ListOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	cli.ListOpts
	store DatabaseRoleLister
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ListOpts) Run() error {
	r, err := opts.store.DatabaseRoles(opts.ConfigProjectID())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas dbroles(s) list --projectId projectId [--page N] [--limit N].
func ListBuilder() *cobra.Command {
	opts := new(ListOpts)
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List custom database roles for your project.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Aliases: []string{"ls"},
		Args:    require.NoArgs,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	opts.AddListOptsFlagsWithoutOmitCount(cmd)
	_ = cmd.Flags().MarkDeprecated(flag.Page, deprecatedFlagMessage)
	_ = cmd.Flags().MarkDeprecated(flag.Limit, deprecatedFlagMessage)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}

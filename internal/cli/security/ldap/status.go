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

package ldap

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

type StatusOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	id    string
	store store.LDAPConfigurationDescriber
}

func (opts *StatusOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var verifyStatusTemplate = `REQUEST ID	PROJECT ID	STATUS
{{.RequestId}}	{{.GroupId}}	{{.Status}}
`

func (opts *StatusOpts) Run() error {
	r, err := opts.store.GetStatusLDAPConfiguration(opts.ConfigProjectID(), opts.id)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas ldap verify status <ID> [--projectId projectId].
func StatusBuilder() *cobra.Command {
	opts := &StatusOpts{}
	cmd := &cobra.Command{
		Use:   "status <requestId>",
		Args:  require.ExactValidArgs(1),
		Short: "Get the status of an LDAP configuration request.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Annotations: map[string]string{
			"requestIdDesc": "ID of the request to verify an LDAP configuration.",
			"output":        verifyStatusTemplate,
			"toc":           "true",
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), verifyStatusTemplate))
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	cmd.AddCommand(
		WatchBuilder(),
	)

	return cmd
}

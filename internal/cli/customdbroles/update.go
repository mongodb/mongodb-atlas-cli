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
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/convert"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const updateTemplate = "Custom database role '{{.RoleName}}' successfully updated.\n"

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	roleName       string
	action         []string
	inheritedRoles []string
	append         bool
	store          store.DatabaseRoleUpdater
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *UpdateOpts) Run() error {
	var role *atlasv2.UserCustomDBRole
	if opts.append {
		var err error
		if role, err = opts.store.DatabaseRole(opts.ConfigProjectID(), opts.roleName); err != nil {
			return err
		}
	}

	out, err := opts.store.UpdateDatabaseRole(opts.ConfigProjectID(), opts.roleName, opts.newCustomDBRole(role))
	if err != nil {
		return err
	}

	return opts.Print(out)
}

func (opts *UpdateOpts) newCustomDBRole(existingRole *atlasv2.UserCustomDBRole) *atlasv2.UserCustomDBRole {
	inheritedRoles := convert.BuildAtlasInheritedRoles(opts.inheritedRoles)
	actions := joinActions(convert.BuildAtlasActions(opts.action))

	if opts.append {
		actions = appendActions(existingRole.GetActions(), actions)
		inheritedRoles = append(inheritedRoles, existingRole.GetInheritedRoles()...)
	}

	return &atlasv2.UserCustomDBRole{
		Actions:        &actions,
		InheritedRoles: &inheritedRoles,
	}
}

func (opts *UpdateOpts) validate() error {
	if len(opts.action) == 0 && len(opts.inheritedRoles) == 0 {
		return errors.New("you must provide either actions or inherited roles")
	}
	return nil
}

// atlas dbrole(s) update create <roleName> --privilege action[@dbName.collection] --inheritedRole role@db --append.
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	cmd := &cobra.Command{
		Use:   "update <roleName>",
		Short: "Update a custom database role for your project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Annotations: map[string]string{
			"roleNameDesc": "Name of the custom role to update.",
			"output":       updateTemplate,
		},
		Args: require.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
				opts.validate,
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.roleName = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringSliceVar(&opts.inheritedRoles, flag.InheritedRole, []string{}, usage.InheritedRoles+usage.UpdateWarning)
	cmd.Flags().StringSliceVar(&opts.action, flag.Privilege, []string{}, usage.PrivilegeAction+usage.UpdateWarning)
	cmd.Flags().BoolVar(&opts.append, flag.Append, false, usage.Append)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}

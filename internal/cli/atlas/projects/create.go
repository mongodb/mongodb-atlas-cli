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

package projects

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	store "github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
)

const atlasCreateTemplate = "Project '{{.Id}}' created.\n"

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	name                        string
	projectOwnerID              string
	regionUsageRestrictions     bool
	withoutDefaultAlertSettings bool
	store                       store.ProjectCreator
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		if opts.ConfigOrgID() == "" {
			return cli.ErrMissingOrgID
		}

		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *CreateOpts) Run() error {
	r, err := opts.store.CreateProject(opts.newCreateProjectOptions())

	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newCreateProjectGroup() atlasv2.Group {
	var defaultAlertSettings *bool
	if opts.withoutDefaultAlertSettings {
		f := false
		defaultAlertSettings = &f
	}
	restrictions := opts.newRegionUsageRestrictions()
	return atlasv2.Group{
		Name:                      opts.name,
		OrgId:                     opts.ConfigOrgID(),
		WithDefaultAlertsSettings: defaultAlertSettings,
		RegionUsageRestrictions:   restrictions,
	}
}

func (opts *CreateOpts) newRegionUsageRestrictions() *string {
	if opts.regionUsageRestrictions {
		govRegionOnly := "GOV_REGIONS_ONLY"
		return &govRegionOnly
	}
	return nil
}

func (opts *CreateOpts) newCreateProjectOptions() *atlasv2.CreateProjectApiParams {
	return &atlasv2.CreateProjectApiParams{ProjectOwnerId: &opts.projectOwnerID, Group: pointer.Get(opts.newCreateProjectGroup())}
}

// atlas project(s) create <name> [--orgId orgId] [--ownerID ownerID] [--withoutDefaultAlertSettings].
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	opts.Template = atlasCreateTemplate
	cmd := &cobra.Command{
		Use:   "create <projectName>",
		Short: "Create a project in your organization.",
		Long: `Projects group clusters into logical collections that support an application environment, workload, or both. Each project can have its own users, teams, security, and alert settings.

` + fmt.Sprintf(usage.RequiredRole, "Project Data Access Read/Write"),
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"projectNameDesc": "Label that identifies the project.",
			"output":          atlasCreateTemplate,
		},
		Example: fmt.Sprintf(`  # Create a project in the organization with the ID 5e2211c17a3e5a48f5497de3 using default alert settings:
  %s projects create --orgId 5e2211c17a3e5a48f5497de3 --output json`, cli.ExampleAtlasEntryPoint()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			return opts.PreRunE(opts.initStore(cmd.Context()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)
	cmd.Flags().StringVar(&opts.projectOwnerID, flag.OwnerID, "", usage.ProjectOwnerID)
	cmd.Flags().BoolVar(&opts.regionUsageRestrictions, flag.GovCloudRegionsOnly, false, usage.GovCloudRegionsOnly)
	cmd.Flags().BoolVar(&opts.withoutDefaultAlertSettings, flag.WithoutDefaultAlertSettings, false, usage.WithoutDefaultAlertSettings)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}

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
	"sort"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const atlasCreateTemplate = "Project '{{.Id}}' created.\n"

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	name                        string
	projectOwnerID              string
	regionUsageRestrictions     bool
	withoutDefaultAlertSettings bool
	tag                         map[string]string
	store                       store.ProjectCreator
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
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

func (opts *CreateOpts) newCreateProjectGroupTags() *[]atlasv2.ResourceTag {
	if len(opts.tag) == 0 {
		return nil
	}

	tags := make([]atlasv2.ResourceTag, 0)

	keys := make([]string, 0)
	for k := range opts.tag {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := opts.tag[key]
		if key == "" || value == "" {
			continue
		}

		resourceTag := *atlasv2.NewResourceTag(key, value)

		tags = append(tags, resourceTag)
	}

	return &tags
}

func (opts *CreateOpts) newCreateProjectGroup() *atlasv2.Group {
	return &atlasv2.Group{
		Name:                      opts.name,
		OrgId:                     opts.ConfigOrgID(),
		WithDefaultAlertsSettings: opts.defaultAlertSettings(),
		RegionUsageRestrictions:   opts.newRegionUsageRestrictions(),
		Tags:                      opts.newCreateProjectGroupTags(),
	}
}

func (opts *CreateOpts) defaultAlertSettings() *bool {
	var defaultAlertSettings *bool
	if opts.withoutDefaultAlertSettings {
		f := false
		defaultAlertSettings = &f
	}
	return defaultAlertSettings
}

func (opts *CreateOpts) newRegionUsageRestrictions() *string {
	if opts.regionUsageRestrictions {
		govRegionOnly := "GOV_REGIONS_ONLY"
		return &govRegionOnly
	}
	return nil
}

func (opts *CreateOpts) newCreateProjectOptions() *atlasv2.CreateProjectApiParams {
	return &atlasv2.CreateProjectApiParams{
		ProjectOwnerId: &opts.projectOwnerID,
		Group:          opts.newCreateProjectGroup(),
	}
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
		Example: `  # Create a project in the organization with the ID 5e2211c17a3e5a48f5497de3 using default alert settings:
  atlas projects create my-project --orgId 5e2211c17a3e5a48f5497de3 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.initStore(cmd.Context()),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)
	cmd.Flags().StringVar(&opts.projectOwnerID, flag.OwnerID, "", usage.ProjectOwnerID)
	cmd.Flags().BoolVar(&opts.regionUsageRestrictions, flag.GovCloudRegionsOnly, false, usage.GovCloudRegionsOnly)
	cmd.Flags().BoolVar(&opts.withoutDefaultAlertSettings, flag.WithoutDefaultAlertSettings, false, usage.WithoutDefaultAlertSettings)
	cmd.Flags().StringToStringVar(&opts.tag, flag.Tag, nil, usage.ProjectTag)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}

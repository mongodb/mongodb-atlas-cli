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

package projects

import (
	"context"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const (
	atlasCreateTemplate = "Project '{{.ID}}' created.\n"
)

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	name                        string
	projectOwnerID              string
	withoutDefaultAlertSettings bool
	serviceVersion              *semver.Version
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
	var defaultAlertSettings *bool
	if opts.withoutDefaultAlertSettings {
		f := false
		defaultAlertSettings = &f
	}

	r, err := opts.store.CreateProject(opts.name, opts.ConfigOrgID(), defaultAlertSettings, opts.newCreateProjectOptions())

	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newCreateProjectOptions() *atlas.CreateProjectOptions {
	return &atlas.CreateProjectOptions{ProjectOwnerID: opts.projectOwnerID}
}

func (opts *CreateOpts) validateOwnerID() error {
	if opts.projectOwnerID == "" || opts.serviceVersion == nil {
		return nil
	}

	constrain, err := semver.NewConstraint(">= 6.0")
	if err != nil {
		return err
	}

	if !constrain.Check(opts.serviceVersion) {
		return fmt.Errorf("%s is available only for Atlas, Cloud Manager and Ops Manager >= 6.0", flag.OwnerID)
	}

	return nil
}

func (opts *CreateOpts) validateWithoutDefaultAlertSettings() error {
	if !opts.withoutDefaultAlertSettings || opts.serviceVersion == nil {
		return nil
	}

	constrain, err := semver.NewConstraint(">= 6.0")
	if err != nil {
		return err
	}

	if !constrain.Check(opts.serviceVersion) {
		return fmt.Errorf("%s is available only for Atlas, Cloud Manager and Ops Manager >= 6.0", flag.WithoutDefaultAlertSettings)
	}

	return nil
}

func (opts *CreateOpts) initServiceVersion() error {
	if config.Service() != config.OpsManagerService {
		return nil
	}
	v, err := opts.store.ServiceVersion()
	if err != nil {
		return err
	}

	sv, err := cli.ParseServiceVersion(v)
	if err != nil {
		return err
	}

	opts.serviceVersion = sv
	return nil
}

// mongocli iam project(s) create <name> [--orgId orgId] [--ownerID ownerID] [--withoutDefaultAlertSettings].
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
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			if !config.IsCloud() {
				opts.Template += "Agent API Key: '{{.AgentAPIKey}}'\n"
			}
			return opts.PreRunE(opts.initStore(cmd.Context()), opts.initServiceVersion, opts.validateOwnerID, opts.validateWithoutDefaultAlertSettings)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)
	cmd.Flags().StringVar(&opts.projectOwnerID, flag.OwnerID, "", usage.ProjectOwnerID)
	cmd.Flags().BoolVar(&opts.withoutDefaultAlertSettings, flag.WithoutDefaultAlertSettings, false, usage.WithoutDefaultAlertSettings)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}

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

package featurepolicies

import (
	"context"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/ops-manager/opsmngr"
)

var updateTemplate = "Feature control policies updated.\n"

type UpdateOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	store    store.FeatureControlPoliciesUpdater
	name     string
	systemID string
	policy   []string
	filename string
	fs       afero.Fs
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *UpdateOpts) Run() error {
	p, err := opts.newFeatureControl()
	if err != nil {
		return err
	}
	r, err := opts.store.UpdateFeatureControlPolicy(opts.ConfigProjectID(), p)
	if err != nil {
		return err
	}
	return opts.Print(r)
}

func (opts *UpdateOpts) newFeatureControl() (*opsmngr.FeaturePolicy, error) {
	policy := new(opsmngr.FeaturePolicy)
	if opts.filename != "" {
		if err := file.Load(opts.fs, opts.filename, policy); err != nil {
			return nil, err
		}
	} else {
		policy.ExternalManagementSystem = &opsmngr.ExternalManagementSystem{
			Name:     opts.name,
			SystemID: opts.systemID,
		}
		policy.Policies = opts.newPolicies()
	}
	return policy, nil
}

func (opts *UpdateOpts) newPolicies() []*opsmngr.Policy {
	policies := make([]*opsmngr.Policy, len(opts.policy))
	for i := range opts.policy {
		policy := &opsmngr.Policy{
			Policy: opts.policy[i],
		}
		policies[i] = policy
	}
	return policies
}

// UpdateBuilder mongocli ops-manager featurePolicy(ies) update --name name --policy policy --systemId systemId [--projectId projectId].
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update feature control policies for your project.",
		Long:  "Feature Control Policies allow you to enable or disable certain MongoDB features based on your site-specific needs.",
		Example: `Disable user management for a project:
  $ mongocli ops-manager featurePolicies update --projectId <projectId> --name Operator --policy DISABLE_USER_MANAGEMENT

  Update policies from a JSON configuration file:
  $ mongocli atlas featurePolicies update --projectId <projectId> --file <path/to/file.json>
`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if opts.filename == "" {
				_ = cmd.MarkFlagRequired(flag.Name)
			}
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringSliceVar(&opts.policy, flag.Policy, []string{}, usage.Policy)
	cmd.Flags().StringVar(&opts.name, flag.Name, "", usage.ExternalSystemName)
	cmd.Flags().StringVar(&opts.systemID, flag.SystemID, "", usage.SystemID)
	cmd.Flags().StringVarP(&opts.filename, flag.File, flag.FileShort, "", usage.PoliciesFilename)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagFilename(flag.File)

	return cmd
}

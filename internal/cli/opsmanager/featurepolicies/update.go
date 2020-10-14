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
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/admin/backupstore"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/ops-manager/opsmngr"
)

var updateTemplate = "Oplog configuration '{{.ID}}' updated.\n"

type UpdateOpts struct {
	cli.OutputOpts
	backupstore.AdminOpts
	cli.GlobalOpts
	store    store.FeatureControlPoliciesUpdater
	name     string
	systemID string
	policy   []string
}

func (opts *UpdateOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *UpdateOpts) Run() error {
	r, err := opts.store.UpdateFeatureControlPolicy(opts.ConfigProjectID(), opts.newFeatureControl())
	if err != nil {
		return err
	}
	return opts.Print(r)
}

func (opts *UpdateOpts) newFeatureControl() *opsmngr.FeaturePolicy {
	return &opsmngr.FeaturePolicy{
		ID: opts.ConfigProjectID(),
		ExternalManagementSystem: &opsmngr.ExternalManagementSystem{
			Name:     opts.name,
			SystemID: opts.systemID,
		},
		Policies: opts.newPolicies(),
	}
}

func (opts *UpdateOpts) newPolicies() []*opsmngr.Policy {
	policies := make([]*opsmngr.Policy, len(opts.policy))
	for _, value := range opts.policy {
		policy := &opsmngr.Policy{
			Policy: value,
		}
		policies = append(policies, policy)
	}
	return policies
}

// mongocli ops-manager featurePolicy(ies) update --name name --policy policy --systemId systemId [--projectId projectId]

func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	opts.Template = updateTemplate
	cmd := &cobra.Command{
		Use:   "update",
		Short: update,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringSliceVar(&opts.policy, flag.Policy, []string{}, usage.Policy)
	cmd.Flags().StringVar(&opts.name, flag.Name, "", usage.ExternalSystemName)
	cmd.Flags().StringVar(&opts.systemID, flag.SystemID, "", usage.SystemID)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.Name)

	return cmd
}

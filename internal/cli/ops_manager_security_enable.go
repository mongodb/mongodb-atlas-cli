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

package cli

import (
	"fmt"

	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/convert"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

const (
	cr     = "MONGODB-CR"
	sha1   = "SCRAM-SHA-1"
	sha256 = "SCRAM-SHA-256"
)

type opsManagerSecurityEnableOpts struct {
	globalOpts
	mechanisms []string
	store      store.AutomationPatcher
}

func (opts *opsManagerSecurityEnableOpts) initStore() error {
	var err error
	opts.store, err = store.New()
	return err
}

func (opts *opsManagerSecurityEnableOpts) Run() error {
	current, err := opts.store.GetAutomationConfig(opts.ProjectID())

	if err != nil {
		return err
	}

	if err := convert.EnableMechanism(current, opts.mechanisms); err != nil {
		return err
	}
	if err := opts.store.UpdateAutomationConfig(opts.ProjectID(), current); err != nil {
		return err
	}

	fmt.Print(deploymentStatus(config.OpsManagerURL(), opts.ProjectID()))

	return nil
}

// mongocli ops-manager security enable[MONGODB-CR|SCRAM-SHA-256]  [--projectId projectId]
func OpsManagerSecurityEnableBuilder() *cobra.Command {
	opts := &opsManagerSecurityEnableOpts{}
	cmd := &cobra.Command{
		Use:       fmt.Sprintf("enable [%s|%s]", cr, sha256),
		Short:     description.EnableSecurity,
		Args:      cobra.OnlyValidArgs,
		ValidArgs: []string{cr, sha1, sha256},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.mechanisms = args
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}

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

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mcli/internal/config"
	"github.com/mongodb/mcli/internal/convert"
	"github.com/mongodb/mcli/internal/flags"
	"github.com/mongodb/mcli/internal/messages"
	"github.com/mongodb/mcli/internal/search"
	"github.com/mongodb/mcli/internal/store"
	"github.com/mongodb/mcli/internal/usage"
	"github.com/spf13/cobra"
)

type cmClustersShutdownOpts struct {
	*globalOpts
	name    string
	confirm bool
	store   store.AutomationStore
}

func (opts *cmClustersShutdownOpts) init() error {
	if opts.ProjectID() == "" {
		return errMissingProjectID
	}

	s, err := store.New()

	if err != nil {
		return err
	}

	opts.store = s
	return nil
}

func (opts *cmClustersShutdownOpts) Run() error {
	current, err := opts.store.GetAutomationConfig(opts.ProjectID())

	if err != nil {
		return err
	}

	if !search.ClusterExists(current, opts.name) {
		return fmt.Errorf("cluster '%s' doesn't exist", opts.name)
	}

	convert.Shutdown(current, opts.name)

	if err = opts.store.UpdateAutomationConfig(opts.ProjectID(), current); err != nil {
		return err
	}

	fmt.Print(messages.DeploymentStatus(config.OpsManagerURL(), opts.ProjectID()))

	return nil
}

func (opts *cmClustersShutdownOpts) Confirm() error {
	if opts.confirm {
		return nil
	}
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("Are you sure you want to shutdown: %s", opts.name),
	}
	return survey.AskOne(prompt, &opts.confirm)
}

// mcli cloud-manager cluster(s) shutdown [name] --projectId projectId [--force]
func CloudManagerClustersShutdownBuilder() *cobra.Command {
	opts := &cmClustersShutdownOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:   "shutdown [name]",
		Short: "Shutdown a Cloud Manager cluster.",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.init(); err != nil {
				return err
			}
			opts.name = args[0]
			return opts.Confirm()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.confirm, flags.Force, false, usage.Force)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}

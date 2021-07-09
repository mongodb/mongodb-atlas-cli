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

package clusters

import (
	"errors"
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/ops-manager/atmcfg"
)

type ShutdownOpts struct {
	cli.GlobalOpts
	clusterName string
	confirm     bool
	processes   []string
	store       store.AutomationPatcher
}

func (opts *ShutdownOpts) initStore() error {
	var err error
	opts.store, err = store.New(store.AuthenticatedPreset(config.Default()))
	return err
}

func (opts *ShutdownOpts) Run() error {
	if !opts.confirm {
		return nil
	}
	current, err := opts.store.GetAutomationConfig(opts.ConfigProjectID())
	if err != nil {
		return err
	}

	err = atmcfg.ShutdownProcessesByClusterName(current, opts.clusterName, opts.processes)
	if err != nil {
		return err
	}

	if err := opts.store.UpdateAutomationConfig(opts.ConfigProjectID(), current); err != nil {
		return err
	}

	fmt.Print(cli.DeploymentStatus(config.OpsManagerURL(), opts.ConfigProjectID()))

	return nil
}

func (opts *ShutdownOpts) Confirm() error {
	if opts.confirm {
		return nil
	}

	shutdownProcess := opts.clusterName

	if shutdownProcess == "" {
		shutdownProcess = strings.Join(opts.processes, ", ")
	}
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("Are you sure you want to shutdown: %s", shutdownProcess),
	}
	return survey.AskOne(prompt, &opts.confirm)
}

func (opts *ShutdownOpts) validateInputs() error {
	if opts.clusterName == "" && len(opts.processes) == 0 {
		return errors.New("you have to provide the Cluster Name or use --process")
	}

	return nil
}

// mongocli cloud-manager cluster(s) shutdown <clusterName> --projectId projectId --processName hostname:port,hostname:port[--force].
func ShutdownBuilder() *cobra.Command {
	opts := &ShutdownOpts{}
	cmd := &cobra.Command{
		Use:   "shutdown <clusterName>",
		Short: "Shutdown a cluster or a list of processes for your project.",
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"args":     "clusterName",
			"nameDesc": "Name of the cluster that you want to shutdown.",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(opts.ValidateProjectID, opts.validateInputs, opts.initStore); err != nil {
				return err
			}
			opts.clusterName = args[0]
			return opts.Confirm()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.confirm, flag.Force, false, usage.Force)
	cmd.Flags().StringSliceVar(&opts.processes, flag.ProcessName, []string{}, usage.ProcessName)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}

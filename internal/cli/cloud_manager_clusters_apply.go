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
	"github.com/mongodb/mongocli/internal/file"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/messages"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

type cmClustersApplyOpts struct {
	*globalOpts
	filename string
	fs       afero.Fs
	store    store.AutomationPatcher
}

func (opts *cmClustersApplyOpts) init() error {
	if opts.ProjectID() == "" {
		return errMissingProjectID
	}

	var err error
	opts.store, err = store.New()
	return err
}

func (opts *cmClustersApplyOpts) Run() error {
	newConfig := new(convert.ClusterConfig)
	err := file.Load(opts.fs, opts.filename, newConfig)
	if err != nil {
		return err
	}
	current, err := opts.store.GetAutomationConfig(opts.ProjectID())

	if err != nil {
		return err
	}

	err = newConfig.PatchAutomationConfig(current)

	if err != nil {
		return err
	}

	if err = opts.store.UpdateAutomationConfig(opts.ProjectID(), current); err != nil {
		return err
	}

	fmt.Print(messages.DeploymentStatus(config.OpsManagerURL(), opts.ProjectID()))

	return nil
}

// mongocli cloud-manager cluster(s) apply --projectId projectId --file myfile.yaml
func CloudManagerClustersApplyBuilder() *cobra.Command {
	opts := &cmClustersApplyOpts{
		globalOpts: newGlobalOpts(),
		fs:         afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "apply",
		Short: description.ApplyOMCluster,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.filename, flags.File, flags.FileShort, "", "Filename to use to change the automation config")

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flags.File)

	return cmd
}

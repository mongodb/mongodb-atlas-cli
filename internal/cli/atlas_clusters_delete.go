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
	"github.com/mongodb/mcli/internal/flags"
	"github.com/mongodb/mcli/internal/store"
	"github.com/mongodb/mcli/internal/usage"
	"github.com/spf13/cobra"
)

type atlasClustersDeleteOpts struct {
	*globalOpts
	name    string
	confirm bool
	store   store.ClusterDeleter
}

func (opts *atlasClustersDeleteOpts) init() error {
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

func (opts *atlasClustersDeleteOpts) Run() error {
	if !opts.confirm {
		fmt.Println("Cluster not deleted")
		return nil
	}
	if err := opts.store.DeleteCluster(opts.ProjectID(), opts.name); err != nil {
		return err
	}

	fmt.Printf("Cluster '%s' deleted\n", opts.name)

	return nil
}

func (opts *atlasClustersDeleteOpts) Confirm() error {
	if opts.confirm {
		return nil
	}
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("Are you sure you want to delete: %s", opts.name),
	}
	return survey.AskOne(prompt, &opts.confirm)
}

// mcli atlas cluster(s) delete name --projectId projectId [--confirm]
func AtlasClustersDeleteBuilder() *cobra.Command {
	opts := &atlasClustersDeleteOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:     "delete [name]",
		Short:   "Delete an Atlas cluster.",
		Aliases: []string{"rm"},
		Args:    cobra.ExactArgs(1),
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

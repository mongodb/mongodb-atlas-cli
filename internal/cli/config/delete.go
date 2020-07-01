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

package config

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongocli/internal/prompt"

	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/spf13/cobra"
)

type RenameOpts struct {
	name string
}

func (opts *RenameOpts) Run() error {
	profile := config.GetConfigDescription(opts.name)
	if len(profile) == 0 {
		return fmt.Errorf("profile %v does not exist", opts.name)
	}

	shouldReplace := false
	p := prompt.NewDeleteConfirm(opts.name)
	if err := survey.AskOne(p, &shouldReplace); err != nil {
		return err
	}

	if !shouldReplace {
		return nil
	}

	config.SetName(&opts.name)
	if err := config.Delete(); err != nil {
		return err
	}

	return nil
}

func DeleteBuilder() *cobra.Command {
	opts := &RenameOpts{}
	cmd := &cobra.Command{
		Use:     "delete <name>",
		Aliases: []string{"rm"},
		Short:   description.ConfigDeleteDescription,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}

	return cmd
}

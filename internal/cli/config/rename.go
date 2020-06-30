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
	oldName string
	newName string
}

func (opts *RenameOpts) Run() error {
	oldProfile := config.GetConfigDescription(opts.oldName, false)
	if len(oldProfile) == 0 {
		return fmt.Errorf("profile %v does not exist", opts.oldName)
	}

	shouldReplace := false
	newProfile := config.GetConfigDescription(opts.newName, false)
	if len(newProfile) > 0 {
		p := prompt.NewProfileReplaceConfirm(opts.newName)
		if err := survey.AskOne(p, &shouldReplace); err != nil {
			return err
		}

		if !shouldReplace {
			return nil
		}
	}

	if err := config.DeleteConfig(opts.oldName); err != nil {
		return err
	}

	if err := config.DeleteConfig(opts.newName); err != nil {
		return err
	}

	if err := config.SetConfig(opts.newName, oldProfile); err != nil {
		return err
	}

	return nil
}

func RenameBuilder() *cobra.Command {
	opts := &RenameOpts{}
	cmd := &cobra.Command{
		Use:   "rename <oldName> <newName>",
		Short: description.ConfigRenameDescription,
		Long:  fmt.Sprintf(description.ConfigRenameLong),
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.oldName = args[0]
			opts.newName = args[1]
			return opts.Run()
		},
	}

	return cmd
}

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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/prompt"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"github.com/spf13/cobra"
)

type RenameOpts struct {
	oldName string
	newName string
}

func (opts *RenameOpts) Run() error {
	if !config.Exists(opts.oldName) {
		return fmt.Errorf("profile %s does not exist", opts.oldName)
	}

	if config.Exists(opts.newName) {
		replaceExistingProfile := false
		p := prompt.NewProfileReplaceConfirm(opts.newName)
		if err := telemetry.TrackAskOne(p, &replaceExistingProfile); err != nil {
			return err
		}

		if !replaceExistingProfile {
			fmt.Printf("Profile was not renamed.\n")
			return nil
		}
	}

	if err := config.SetName(opts.oldName); err != nil {
		return err
	}

	if err := config.Rename(opts.newName); err != nil {
		return err
	}

	fmt.Printf("The profile %s was renamed to %s.\n", opts.oldName, opts.newName)
	return nil
}

func RenameBuilder() *cobra.Command {
	const argsN = 2
	o := &RenameOpts{}
	cmd := &cobra.Command{
		Use:     "rename <oldProfileName> <newProfileName>",
		Aliases: []string{"mv"},
		Short:   "Rename a profile.",
		Example: `  # Rename a profile called myProfile to testProfile:
  atlas config rename myProfile testProfile`,
		Annotations: map[string]string{
			"oldProfileNameDesc": "Name of the profile to rename.",
			"newProfileNameDesc": "New name of the profile.",
		},
		Args: require.ExactArgs(argsN),
		RunE: func(_ *cobra.Command, args []string) error {
			o.oldName = args[0]
			o.newName = args[1]
			return o.Run()
		},
	}

	return cmd
}

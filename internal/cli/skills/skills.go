// Copyright 2026 MongoDB Inc
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

package skills

import (
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

// Builder builds the top-level skills command.
func Builder() *cobra.Command {
	const use = "skills"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   "Discover and install MongoDB agent skills for Claude Code or Codex.",
		Long:    "List MongoDB agent skills from the upstream repository and install them for Claude Code or Codex.",
	}
	hideAtlasProfilePersistentFlag(cmd)

	cmd.AddCommand(
		ListBuilder(),
		InstallBuilder(),
	)

	return cmd
}

func hideAtlasProfileFlag(cmd *cobra.Command) {
	cmd.Flags().StringP(flag.Profile, flag.ProfileShort, "", usage.ProfileAtlasCLI)
	_ = cmd.Flags().MarkHidden(flag.Profile)
}

func hideAtlasProfilePersistentFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP(flag.Profile, flag.ProfileShort, "", usage.ProfileAtlasCLI)
	_ = cmd.PersistentFlags().MarkHidden(flag.Profile)
}

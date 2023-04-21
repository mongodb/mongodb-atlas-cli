// Copyright 2023 MongoDB Inc
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

package accesslists

import (
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
	const use = "accessLists"
	deprecated := append([]string{"whitelists"}, cli.GenerateAliases("whitelists")...)
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Manage the IP access list for your API Key.",
		Aliases: cli.GenerateAliases(use, deprecated...),
	}

	cmd.AddCommand(
		ListBuilder(),
		CreateBuilder(),
		DeleteBuilder(),
	)

	return cmd
}

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

package organizations

import (
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/cli/organizations/apikeys"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/cli/organizations/invitations"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/cli/organizations/users"
	"github.com/spf13/cobra"
)

const use = "organizations"

func Builder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Manage your Atlas organizations.",
		Long:    "Create, list and manage your MongoDB organizations.",
		Aliases: cli.GenerateAliases(use, "orgs", "org"),
	}
	cmd.AddCommand(
		CreateAtlasBuilder(),
		ListBuilder(),
		DescribeBuilder(),
		DeleteBuilder(),
		apikeys.Builder(),
		users.Builder(),
		invitations.Builder(),
	)
	return cmd
}

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

package privateendpoints

import (
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/privateendpoints/aws"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/privateendpoints/azure"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/privateendpoints/datalake"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/privateendpoints/gcp"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/privateendpoints/interfaces"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/privateendpoints/onlinearchive"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/privateendpoints/regionalmodes"
	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
	const use = "privateEndpoints"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   "Manage Atlas private endpoints.",
	}
	cmd.AddCommand(
		ListBuilder(),
		DescribeBuilder(),
		CreateBuilder(),
		DeleteBuilder(),
		WatchBuilder(),
		interfaces.Builder(),
		aws.Builder(),
		azure.Builder(),
		gcp.Builder(),
		regionalmodes.Builder(),
		datalake.Builder(),
		onlinearchive.Builder(),
	)

	return cmd
}

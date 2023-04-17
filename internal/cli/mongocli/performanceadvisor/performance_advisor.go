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

package performanceadvisor

import (
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/mongocli/performanceadvisor/namespaces"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/mongocli/performanceadvisor/slowoperationthreshold"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/mongocli/performanceadvisor/slowquerylogs"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/mongocli/performanceadvisor/suggestedindexes"
	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
	const use = "performanceAdvisor"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   "Learn more about slow queries and get suggestions to improve database performance.",
	}
	cmd.AddCommand(
		namespaces.Builder(),
		slowquerylogs.Builder(),
		suggestedindexes.Builder(),
		slowoperationthreshold.Builder())

	return cmd
}

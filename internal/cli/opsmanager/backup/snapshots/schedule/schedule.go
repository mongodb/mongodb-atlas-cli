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

<<<<<<< HEAD:internal/cli/opsmanager/agents/versions/versions.go
package versions
=======
package schedule
>>>>>>> origin/master:internal/cli/opsmanager/backup/snapshots/schedule/schedule.go

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
<<<<<<< HEAD:internal/cli/opsmanager/agents/versions/versions.go
	const use = "versions"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   short,
	}
	cmd.AddCommand()
=======
	const use = "schedule"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   shots,
	}

	cmd.AddCommand(
		DescribeBuilder(),
		UpdateBuilder())

>>>>>>> origin/master:internal/cli/opsmanager/backup/snapshots/schedule/schedule.go
	return cmd
}

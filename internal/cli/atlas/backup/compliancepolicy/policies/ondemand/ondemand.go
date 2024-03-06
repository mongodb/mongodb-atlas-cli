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

package ondemand

import (
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/spf13/cobra"
)

func baseCommand() *cobra.Command {
	const use = "ondemand"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   "Manage the on-demand policy item of the backup compliance policy for your project.",
	}

	return cmd
}

func Builder() *cobra.Command {
	cmd := baseCommand()

	cmd.AddCommand(
		CreateBuilder(),
		UpdateBuilder(),
		DescribeBuilder(),
		// delete command not available as once set,
		// an on-demand policy can only be updated.
	)

	return cmd
}

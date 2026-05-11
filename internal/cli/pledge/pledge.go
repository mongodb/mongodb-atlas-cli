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

package pledge

import (
	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pledge",
		Short: "Manage session permission pledges.",
		Long: `A pledge restricts the current shell session to a subset of Atlas CLI operations.
All atlas invocations in the same session — including subshells — inherit the restriction.
Pledges can only be narrowed, never widened.`,
	}

	cmd.AddCommand(
		SetBuilder(),
		ShowBuilder(),
		AllowBuilder(),
	)

	return cmd
}

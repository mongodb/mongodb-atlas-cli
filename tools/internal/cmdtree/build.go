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

// Package cmdtree reports the atlas CLI command tree, the function cobra runs
// for each command, and the Atlas SDK functions that function can reach.
package cmdtree

import "github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/root"

// packagesPattern selects the source the call graph is built from.
const packagesPattern = "github.com/mongodb/mongodb-atlas-cli/atlascli/internal/..."

// Build returns the atlas CLI command tree with the Atlas SDK functions behind
// each command filled in.
func Build() (Command, error) {
	tree := Walk(root.Builder())

	var symbols []string
	each(&tree, func(cmd *Command) {
		if cmd.Run != "" {
			symbols = append(symbols, cmd.Run)
		}
	})

	calls, err := SDKCalls([]string{packagesPattern}, symbols)
	if err != nil {
		return Command{}, err
	}

	each(&tree, func(cmd *Command) {
		cmd.ApiOperationIds = calls[cmd.Run]
		cmd.Category = category(cmd)
		cmd.LegacyCategory = legacyCategory(cmd)
	})
	return tree, nil
}

// each visits cmd and every command below it.
func each(cmd *Command, visit func(*Command)) {
	visit(cmd)
	for i := range cmd.Children {
		each(&cmd.Children[i], visit)
	}
}

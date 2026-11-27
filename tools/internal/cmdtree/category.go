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

package cmdtree

import "strings"

// Category says how much a command does, judged by the Atlas SDK operations it
// reaches.
type Category string

const (
	// CategoryMeta is a command that calls no Atlas SDK operation.
	CategoryMeta Category = "Meta"
	// CategoryWorkflow is a command that calls exactly one.
	CategoryWorkflow Category = "Workflow"
	// CategoryDirect is a command that calls more than one.
	CategoryDirect Category = "Direct"
)

// category classifies cmd. A command that only groups subcommands runs nothing,
// so it is left uncategorised rather than reported as Meta.
func category(cmd *Command) Category {
	if cmd.Run == "" {
		return ""
	}
	switch len(cmd.ApiOperationIds) {
	case 0:
		return CategoryMeta
	case 1:
		return CategoryWorkflow
	default:
		return CategoryDirect
	}
}

// LegacyCategory says how far a command is from the generated atlas api
// passthrough.
type LegacyCategory string

const (
	// LegacyL1 is the generated atlas api subtree.
	LegacyL1 LegacyCategory = "L1"
	// LegacyL2 is a command that maps to exactly one Atlas SDK operation.
	LegacyL2 LegacyCategory = "L2"
	// LegacyL3 is everything else.
	LegacyL3 LegacyCategory = "L3"
)

// apiCommandPath is the root of the generated subtree.
const apiCommandPath = "atlas api"

// legacyCategory classifies cmd. The api subtree is checked first, so an api
// command stays L1 even though it reaches no Atlas SDK operation: it goes
// through the hand-written executor instead.
func legacyCategory(cmd *Command) LegacyCategory {
	if isAPISubtree(cmd.Path) {
		return LegacyL1
	}
	if len(cmd.ApiOperationIds) == 1 {
		return LegacyL2
	}
	return LegacyL3
}

// isAPISubtree reports whether path is the api command or one below it.
func isAPISubtree(path string) bool {
	return path == apiCommandPath || strings.HasPrefix(path, apiCommandPath+" ")
}

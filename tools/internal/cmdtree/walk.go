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

import (
	"reflect"
	"runtime"
	"sort"

	"github.com/spf13/cobra"
)

// Command is one node of the atlas CLI command tree. A command that groups
// subcommands carries Children; a command that runs carries the function cobra
// runs for it and the Atlas SDK operations that function can reach. A command
// can be both, so the fields are independent rather than exclusive.
type Command struct {
	Path            string         `json:"path"`
	Use             string         `json:"use"`
	Short           string         `json:"short,omitempty"`
	Hidden          bool           `json:"hidden,omitempty"`
	Run             string         `json:"run,omitempty"`
	ApiOperationIds []ApiOperation `json:"ApiOperationIds,omitempty"`
	Category        Category       `json:"category,omitempty"`
	LegacyCategory  LegacyCategory `json:"legacy_category,omitempty"`
	Flags           []Flag         `json:"flags,omitempty"`
	Children        []Command      `json:"children,omitempty"`
}

// Walk returns the command tree rooted at root. Taking a *cobra.Command rather
// than building one keeps it usable on a hand-built tree in tests.
func Walk(root *cobra.Command) Command {
	return Command{
		Path:     root.CommandPath(),
		Use:      root.Use,
		Short:    root.Short,
		Hidden:   root.Hidden,
		Run:      canonical(runSymbol(root)),
		Flags:    flags(root),
		Children: children(root),
	}
}

// children returns the subcommands of cmd, sorted by name so the output does
// not depend on the order they were registered in.
func children(cmd *cobra.Command) []Command {
	subcommands := cmd.Commands()
	sort.Slice(subcommands, func(i, j int) bool { return subcommands[i].Name() < subcommands[j].Name() })

	children := make([]Command, 0, len(subcommands))
	for _, subcommand := range subcommands {
		children = append(children, Walk(subcommand))
	}
	return children
}

// runSymbol names the function cobra runs for cmd, preferring RunE over Run.
// Commands that only group subcommands have neither and return "".
func runSymbol(cmd *cobra.Command) string {
	if cmd.RunE != nil {
		return funcName(reflect.ValueOf(cmd.RunE).Pointer())
	}
	if cmd.Run != nil {
		return funcName(reflect.ValueOf(cmd.Run).Pointer())
	}
	return ""
}

func funcName(pc uintptr) string {
	f := runtime.FuncForPC(pc)
	if f == nil {
		return ""
	}
	return f.Name()
}

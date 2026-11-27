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

// Cmdtree prints the atlas CLI command tree as JSON on stdout, with the
// function cobra runs for each command and the Atlas SDK functions that
// function can reach.
//
// The tree comes from the running program, so commands built at runtime (the
// generated 'atlas api' subtree, for one) are included. The SDK calls come from
// a call graph over the same source, because no breakpoint can tell you what a
// command calls.
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/internal/cmdtree"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "cmdtree:", err)
		os.Exit(1)
	}
}

func run() error {
	tree, err := cmdtree.Build()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(tree)
}

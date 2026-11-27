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
	"fmt"
	"sort"
	"strings"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

// sdkPrefix is the module path of the Atlas Go SDK. Calls into it end the walk.
const sdkPrefix = "go.mongodb.org/atlas-sdk/"

// SDKCalls reports, for each function in symbols, the Atlas SDK operations
// reachable from it. A symbol absent from the program is an error: silently
// returning nothing for it would be indistinguishable from a command that
// genuinely calls no operations.
func SDKCalls(patterns, symbols []string) (map[string][]ApiOperation, error) {
	pkgs, err := packages.Load(&packages.Config{Mode: packages.LoadSyntax}, patterns...)
	if err != nil {
		return nil, err
	}
	if n := packages.PrintErrors(pkgs); n > 0 {
		return nil, fmt.Errorf("loading %v: %d package errors", patterns, n)
	}

	prog, _ := ssautil.AllPackages(pkgs, ssa.InstantiateGenerics)
	prog.Build()

	byName := make(map[string]*ssa.Function)
	for fn := range ssautil.AllFunctions(prog) {
		byName[canonical(fn.String())] = fn
	}

	reacher := newReacher(prog)

	calls := make(map[string][]ApiOperation)
	var missing []string
	for symbol := range unique(symbols) {
		fn, ok := byName[symbol]
		if !ok {
			missing = append(missing, symbol)
			continue
		}
		calls[symbol] = reacher.SDKCallers(fn)
	}
	if len(missing) > 0 {
		sort.Strings(missing)
		return nil, fmt.Errorf("not found in the program: %s", strings.Join(missing, ", "))
	}
	return calls, nil
}

// isSDK reports whether fn belongs to the Atlas SDK.
func isSDK(fn *ssa.Function) bool {
	if fn == nil || fn.Pkg == nil || fn.Pkg.Pkg == nil {
		return false
	}
	return strings.HasPrefix(fn.Pkg.Pkg.Path(), sdkPrefix)
}

func unique(values []string) map[string]bool {
	set := make(map[string]bool, len(values))
	for _, value := range values {
		set[value] = true
	}
	return set
}

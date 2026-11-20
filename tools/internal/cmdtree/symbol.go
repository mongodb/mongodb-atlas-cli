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

const (
	pointerReceiver = "(*" // ssa opens a pointer receiver with this
	valueReceiver   = "("  // and a value receiver with this
	methodSep       = ")." // both use this to separate the receiver from the method
	qualifiedSep    = "/"  // separates a package path from its last element
	typeSep         = "."  // separates a package from its type, and a type from its method
	anonSep         = "$"  // ssa marks wrappers and anonymous functions with this
)

// canonical maps a function name from either the Go runtime
// ("pkg/path.(*Type).Method-fm", "pkg/path.Func.func1") or from x/tools/ssa
// ("(*pkg/path.Type).Method", "pkg/path.Func$1", "pkg/path.Func$bound") onto a
// single readable form: the runtime one.
//
// The two spellings disagree on where the package path goes for methods and on
// how anonymous functions and wrappers are suffixed. Wrapper suffixes are
// dropped rather than kept, because a wrapper has no body of its own: the
// function that actually runs is the one it forwards to, and that is the one
// worth following through the call graph.
func canonical(name string) string {
	name = strings.TrimSuffix(name, "-fm")
	name = strings.TrimSuffix(name, anonSep+"bound")
	name = strings.TrimSuffix(name, anonSep+"thunk")
	name = anonFuncs(name)

	if rest, ok := strings.CutPrefix(name, pointerReceiver); ok {
		if recv, method, ok := splitMethod(rest); ok {
			if pkg, typ, ok := splitQualified(recv); ok {
				return pkg + typeSep + pointerReceiver + typ + methodSep + method
			}
		}
	}
	if rest, ok := strings.CutPrefix(name, valueReceiver); ok {
		if recv, method, ok := splitMethod(rest); ok {
			if pkg, typ, ok := splitQualified(recv); ok {
				return pkg + typeSep + typ + typeSep + method
			}
		}
	}
	return name
}

// splitMethod splits "pkg/path.Type).Method" into "pkg/path.Type" and "Method".
func splitMethod(s string) (recv, method string, ok bool) {
	i := strings.Index(s, methodSep)
	if i < 0 {
		return "", "", false
	}
	return s[:i], s[i+len(methodSep):], true
}

// splitQualified splits "github.com/a/b.Type" into its package path
// "github.com/a/b" and type name "Type".
func splitQualified(s string) (pkg, typ string, ok bool) {
	slash := strings.LastIndex(s, qualifiedSep)
	dot := strings.LastIndex(s[slash+1:], typeSep)
	if dot < 0 {
		return "", "", false
	}
	dot += slash + 1
	return s[:dot], s[dot+1:], true
}

// anonFuncs rewrites ssa's "$1" and "$1$1" anonymous function suffixes into the
// runtime's ".func1" and ".func1.1".
func anonFuncs(name string) string {
	if !strings.Contains(name, anonSep) {
		return name
	}
	parts := strings.Split(name, anonSep)
	var b strings.Builder
	b.WriteString(parts[0])
	for i, part := range parts[1:] {
		if i == 0 {
			b.WriteString(".func")
		} else {
			b.WriteString(typeSep)
		}
		b.WriteString(part)
	}
	return b.String()
}

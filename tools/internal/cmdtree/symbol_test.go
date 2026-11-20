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

import "testing"

// The spellings the Go runtime and x/tools/ssa use for the same functions.
const (
	packageFunc      = "github.com/a/b/pkg.Build"
	runtimeRunE      = "github.com/a/b/pkg.(*Opts).RunE"
	ssaRunE          = "(*github.com/a/b/pkg.Opts).RunE"
	ssaValueRunE     = "(github.com/a/b/pkg.Opts).RunE"
	valueRunE        = "github.com/a/b/pkg.Opts.RunE"
	runtimeClosure   = "github.com/a/b/pkg.Build.func1"
	ssaClosure       = "github.com/a/b/pkg.Build$1"
	ssaNestedClosure = "github.com/a/b/pkg.Build$1$2"
	nestedClosure    = "github.com/a/b/pkg.Build.func1.2"
)

func TestCanonical(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"package function is spelled the same either way", packageFunc, packageFunc},
		{"runtime pointer method", runtimeRunE, runtimeRunE},
		{"ssa pointer method", ssaRunE, runtimeRunE},
		{"ssa value method", ssaValueRunE, valueRunE},
		{"runtime method value wrapper", runtimeRunE + "-fm", runtimeRunE},
		{"ssa bound method value wrapper", ssaRunE + "$bound", runtimeRunE},
		{"ssa thunk wrapper", ssaRunE + "$thunk", runtimeRunE},
		{"runtime closure", runtimeClosure, runtimeClosure},
		{"ssa closure", ssaClosure, runtimeClosure},
		{"ssa nested closure", ssaNestedClosure, nestedClosure},
		{"no package path", "main.main", "main.main"},
		{"unqualified receiver is left alone", "(*Broken).Run", "(*Broken).Run"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canonical(tt.in); got != tt.want {
				t.Errorf("canonical(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

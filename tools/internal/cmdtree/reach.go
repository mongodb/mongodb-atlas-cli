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
	"go/types"
	"sort"
	"strings"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

// target is a function a call can reach, with the values bound to its free
// variables when it came from a closure.
type target struct {
	fn       *ssa.Function
	bindings []ssa.Value
}

// frame is a function to walk, together with the values bound to its
// parameters and free variables. Those bindings are what make the walk context
// sensitive.
type frame struct {
	fn  *ssa.Function
	env map[ssa.Value][]target
}

// reacher walks a program from an entry function, resolving every call with
// the arguments passed at that call site.
type reacher struct {
	methods map[string][]*ssa.Function
}

// newReacher indexes the methods of prog so that a call through an interface
// can be resolved to the types that implement it.
func newReacher(prog *ssa.Program) *reacher {
	methods := map[string][]*ssa.Function{}
	for fn := range ssautil.AllFunctions(prog) {
		if fn.Synthetic != "" || fn.Object() == nil || fn.Signature == nil || fn.Signature.Recv() == nil {
			continue
		}
		methods[fn.Name()] = append(methods[fn.Name()], fn)
	}
	return &reacher{methods: methods}
}

// SDKCallers returns the Atlas SDK operations reachable from entry.
//
// A whole-program call graph cannot answer this question here. atlas-cli's
// shared helpers take a callback, and cli.DeleteOpts.Delete is one of them: it
// receives a func() error and calls it. A call graph links that call to every
// function value of the same shape in the program, so every command using the
// helper appears to call every other command's SDK function. Following the
// argument passed at the call site is what keeps a command's calls its own.
func (r *reacher) SDKCallers(entry *ssa.Function) []ApiOperation {
	found := map[ApiOperation]bool{}
	seen := map[string]bool{}
	queue := []frame{{fn: entry}}

	for len(queue) > 0 {
		f := queue[0]
		queue = queue[1:]

		key := f.key()
		if seen[key] {
			continue
		}
		seen[key] = true

		for _, block := range f.fn.Blocks {
			for _, instr := range block.Instrs {
				call, ok := instr.(ssa.CallInstruction)
				if !ok {
					continue
				}
				for _, t := range r.callees(call.Common(), f.env) {
					if isSDK(t.fn) {
						if operation, ok := apiOperation(canonical(t.fn.String())); ok {
							found[operation] = true
						}
						continue
					}
					if len(t.fn.Blocks) == 0 {
						continue // declared in a dependency, so there is nothing to walk
					}
					queue = append(queue, r.step(call.Common(), t, f.env))
				}
			}
		}
	}

	operations := make([]ApiOperation, 0, len(found))
	for operation := range found {
		operations = append(operations, operation)
	}
	sort.Slice(operations, func(i, j int) bool {
		if operations[i].Version != operations[j].Version {
			return operations[i].Version < operations[j].Version
		}
		return operations[i].OperationID < operations[j].OperationID
	})
	return operations
}

// step binds the callee's parameters and free variables to the values the
// caller passed.
func (r *reacher) step(call *ssa.CallCommon, t target, env map[ssa.Value][]target) frame {
	next := map[ssa.Value][]target{}

	// A variadic function has one slice parameter where the call site has the
	// elements, so only the parameters before it line up with the arguments.
	fixed := t.fn.Params
	if t.fn.Signature.Variadic() && len(fixed) > 0 {
		fixed = fixed[:len(fixed)-1]
	}
	for i, param := range fixed {
		if i >= len(call.Args) {
			break
		}
		if targets := r.resolve(call.Args[i], env); len(targets) > 0 {
			next[param] = targets
		}
	}
	for i, free := range t.fn.FreeVars {
		if i >= len(t.bindings) {
			break
		}
		if targets := r.resolve(t.bindings[i], env); len(targets) > 0 {
			next[free] = targets
		}
	}
	return frame{fn: t.fn, env: next}
}

// callees returns the functions call can reach.
func (r *reacher) callees(call *ssa.CallCommon, env map[ssa.Value][]target) []target {
	if closure, ok := call.Value.(*ssa.MakeClosure); ok {
		if fn, ok := closure.Fn.(*ssa.Function); ok {
			return []target{{fn: fn, bindings: closure.Bindings}}
		}
	}
	if fn := call.StaticCallee(); fn != nil {
		return []target{{fn: fn}}
	}
	if call.IsInvoke() {
		return r.implementations(call)
	}
	return r.resolve(call.Value, env)
}

// implementations returns the methods that satisfy an interface call.
func (r *reacher) implementations(call *ssa.CallCommon) []target {
	iface, ok := call.Value.Type().Underlying().(*types.Interface)
	if !ok {
		return nil
	}
	methods := r.methods[call.Method.Name()]
	targets := make([]target, 0, len(methods))
	for _, fn := range methods {
		if !types.Implements(fn.Signature.Recv().Type(), iface) {
			continue
		}
		targets = append(targets, target{fn: fn})
	}
	return targets
}

// resolve returns the functions a value can hold. A value that is not bound
// and not a literal function is unknown, so the walk stops there rather than
// guessing at every function of a matching type.
func (r *reacher) resolve(v ssa.Value, env map[ssa.Value][]target) []target {
	return r.resolveSeen(v, env, map[ssa.Value]bool{})
}

// resolveSeen carries the values already visited, because a phi node can refer
// back to itself.
func (r *reacher) resolveSeen(v ssa.Value, env map[ssa.Value][]target, seen map[ssa.Value]bool) []target {
	if targets, ok := env[v]; ok {
		return targets
	}
	if seen[v] {
		return nil
	}
	seen[v] = true

	switch v := v.(type) {
	case *ssa.Function:
		return []target{{fn: v}}
	case *ssa.MakeClosure:
		fn, ok := v.Fn.(*ssa.Function)
		if !ok {
			return nil
		}
		return []target{{fn: fn, bindings: v.Bindings}}
	case *ssa.ChangeType:
		return r.resolveSeen(v.X, env, seen)
	case *ssa.Convert:
		return r.resolveSeen(v.X, env, seen)
	case *ssa.MakeInterface:
		return r.resolveSeen(v.X, env, seen)
	case *ssa.Extract:
		// A comma-ok type assertion yields a tuple, and the type switch calls
		// the value taken out of it.
		return r.resolveSeen(v.Tuple, env, seen)
	case *ssa.TypeAssert:
		// cli.DeleteOpts.Delete takes the callback as any and type switches on
		// it, so the function value arrives through a type assertion.
		return r.resolveSeen(v.X, env, seen)
	case *ssa.Phi:
		targets := make([]target, 0, len(v.Edges))
		for _, edge := range v.Edges {
			targets = append(targets, r.resolveSeen(edge, env, seen)...)
		}
		return targets
	}
	return nil
}

// key identifies a frame by its function and its bindings, so the same
// function reached with different arguments is walked once per argument set.
func (f frame) key() string {
	bindings := make([]string, 0, len(f.env))
	for value, targets := range f.env {
		for _, t := range targets {
			bindings = append(bindings, value.String()+" -> "+t.fn.String())
		}
	}
	sort.Strings(bindings)

	var b strings.Builder
	b.WriteString(f.fn.String())
	for _, binding := range bindings {
		b.WriteString(" | ")
		b.WriteString(binding)
	}
	return b.String()
}

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

// Package agent defines the interface for AI agent hook adapters.
package agent

// State reports the installation status of a hook.
type State string

const (
	StateInstalled   State = "installed"
	StateUninstalled State = "not installed"
	StateUnknown     State = "unknown"
)

// InstallOpts carries parameters for hook installation.
type InstallOpts struct {
	// Profile is the pledge profile to pledge at session start (e.g., "readonly").
	Profile string
	// ProjectLevel if true, installs the hook in the current directory's .claude/ instead of ~/.claude/.
	ProjectLevel bool
}

// Agent is implemented by each supported AI coding agent.
type Agent interface {
	// Name returns the canonical agent name (e.g., "claude-code").
	Name() string
	// Install adds the pledge hook to the agent's configuration.
	Install(opts InstallOpts) error
	// Uninstall removes the _atlas_managed hook entries.
	Uninstall() error
	// Status reports whether the hook is currently installed.
	Status() State
}

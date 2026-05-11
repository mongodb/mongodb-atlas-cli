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

package agent

import "errors"

// ErrCodexUnsupported is returned by the Codex adapter until Codex publishes
// session-hook documentation.
var ErrCodexUnsupported = errors.New(
	"codex hook support is not yet available; " +
		"session hooks are not yet documented for Codex. " +
		"Track progress at https://github.com/mongodb/mongodb-atlas-cli/issues",
)

type codexAgent struct{}

// NewCodex returns the Codex agent stub.
func NewCodex() Agent { return &codexAgent{} }

func (a *codexAgent) Name() string     { return "codex" }
func (a *codexAgent) Status() State    { return StateUnknown }
func (a *codexAgent) Install(_ InstallOpts) error { return ErrCodexUnsupported }
func (a *codexAgent) Uninstall() error { return ErrCodexUnsupported }

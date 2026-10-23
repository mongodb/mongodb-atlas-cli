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

package skills

import (
	"context"
	"fmt"
	"strings"
)

const jsonOutputFormat = "json"

type fakeAgentCommandRunner struct {
	agent       string
	path        string
	lookPathErr error
	lookedUp    string
	responses   map[string][]string
	calls       []string
}

func (runner *fakeAgentCommandRunner) LookPath(command string) (string, error) {
	runner.lookedUp = command
	return runner.path, runner.lookPathErr
}

func (runner *fakeAgentCommandRunner) Run(_ context.Context, executable string, args ...string) ([]byte, error) {
	if executable != runner.path {
		return nil, fmt.Errorf("unexpected %s executable %q", runner.agent, executable)
	}
	key := strings.Join(args, " ")
	runner.calls = append(runner.calls, key)
	responses, ok := runner.responses[key]
	if !ok || len(responses) == 0 {
		return nil, fmt.Errorf("unexpected %s command %q", runner.agent, key)
	}
	response := responses[0]
	runner.responses[key] = responses[1:]
	return []byte(response), nil
}

func (runner *fakeAgentCommandRunner) addResponse(command, output string) {
	runner.responses[command] = append(runner.responses[command], output)
}

func newFakeAgentCommandRunner(agent string) *fakeAgentCommandRunner {
	return &fakeAgentCommandRunner{
		agent:     agent,
		path:      "/mock/bin/" + agent,
		responses: make(map[string][]string),
	}
}

func newFakeClaudeCommandRunner() *fakeAgentCommandRunner {
	return newFakeAgentCommandRunner(claudeAgent)
}

func newFakeCodexCommandRunner() *fakeAgentCommandRunner {
	return newFakeAgentCommandRunner(codexAgent)
}

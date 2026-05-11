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

package agent_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/hook/agent"
)

func shellConfigPath(t *testing.T) string {
	t.Helper()
	return filepath.Join(t.TempDir(), ".bashrc")
}

func TestShellInstallOnEmptyFile(t *testing.T) {
	path := shellConfigPath(t)
	a := agent.NewShell(path)
	if err := a.Install(agent.InstallOpts{Profile: "readonly"}); err != nil {
		t.Fatal(err)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(content), "atlas pledge set readonly") {
		t.Error("expected pledge snippet in shell config after install")
	}
}

func TestShellInstallIdempotent(t *testing.T) {
	path := shellConfigPath(t)
	a := agent.NewShell(path)
	opts := agent.InstallOpts{Profile: "readonly"}
	if err := a.Install(opts); err != nil {
		t.Fatal(err)
	}
	if err := a.Install(opts); err != nil {
		t.Fatal(err)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	count := strings.Count(string(content), "atlas pledge set readonly")
	if count != 1 {
		t.Errorf("expected exactly 1 pledge line after idempotent install, got %d", count)
	}
}

func TestShellUninstallPreservesUnrelated(t *testing.T) {
	path := shellConfigPath(t)
	initial := "export FOO=bar\n"
	if err := os.WriteFile(path, []byte(initial), 0o600); err != nil {
		t.Fatal(err)
	}

	a := agent.NewShell(path)
	if err := a.Install(agent.InstallOpts{Profile: "readonly"}); err != nil {
		t.Fatal(err)
	}
	if err := a.Uninstall(); err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	s := string(content)
	if strings.Contains(s, "atlas pledge") {
		t.Error("pledge snippet should be gone after uninstall")
	}
	if !strings.Contains(s, "FOO=bar") {
		t.Error("unrelated content should survive uninstall")
	}
}

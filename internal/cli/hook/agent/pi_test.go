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

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func newTestPi(t *testing.T) *piAgent {
	t.Helper()
	dir := t.TempDir()
	return &piAgent{extensionPath: filepath.Join(dir, ".pi", "agent", "extensions", "atlas-pledge.ts")}
}

func TestPiInstall_MissingFile(t *testing.T) {
	a := newTestPi(t)

	if err := a.Install(InstallOpts{Profile: "readonly"}); err != nil {
		t.Fatalf("Install: %v", err)
	}

	content, err := os.ReadFile(a.extensionPath)
	if err != nil {
		t.Fatalf("reading extension file: %v", err)
	}
	if !strings.Contains(string(content), "readonly") {
		t.Error("profile not substituted in extension content")
	}
	if strings.Contains(string(content), "{{.Profile}}") {
		t.Error("template placeholder not replaced")
	}
}

func TestPiInstall_CustomProfile(t *testing.T) {
	a := newTestPi(t)

	if err := a.Install(InstallOpts{Profile: "read-write"}); err != nil {
		t.Fatalf("Install: %v", err)
	}

	content, _ := os.ReadFile(a.extensionPath)
	if !strings.Contains(string(content), "read-write") {
		t.Error("custom profile not substituted")
	}
}

func TestPiInstall_Idempotent(t *testing.T) {
	a := newTestPi(t)

	if err := a.Install(InstallOpts{Profile: "readonly"}); err != nil {
		t.Fatalf("first Install: %v", err)
	}
	info1, _ := os.Stat(a.extensionPath)

	if err := a.Install(InstallOpts{Profile: "readonly"}); err != nil {
		t.Fatalf("second Install: %v", err)
	}
	info2, _ := os.Stat(a.extensionPath)

	// File should not have been modified (same mod time).
	if info1.ModTime() != info2.ModTime() {
		t.Error("idempotent install modified the file")
	}
}

func TestPiInstall_OverwritesOnContentChange(t *testing.T) {
	a := newTestPi(t)

	if err := a.Install(InstallOpts{Profile: "readonly"}); err != nil {
		t.Fatalf("first Install: %v", err)
	}
	if err := a.Install(InstallOpts{Profile: "read-write"}); err != nil {
		t.Fatalf("second Install with different profile: %v", err)
	}

	content, _ := os.ReadFile(a.extensionPath)
	if !strings.Contains(string(content), "read-write") {
		t.Error("extension not updated to new profile")
	}
}

func TestPiUninstall_RemovesFile(t *testing.T) {
	a := newTestPi(t)

	_ = a.Install(InstallOpts{Profile: "readonly"})

	if err := a.Uninstall(); err != nil {
		t.Fatalf("Uninstall: %v", err)
	}
	if _, err := os.Stat(a.extensionPath); !os.IsNotExist(err) {
		t.Error("extension file still exists after uninstall")
	}
}

func TestPiUninstall_NoopWhenAbsent(t *testing.T) {
	a := newTestPi(t)
	if err := a.Uninstall(); err != nil {
		t.Fatalf("Uninstall on missing file: %v", err)
	}
}

func TestPiStatus(t *testing.T) {
	a := newTestPi(t)

	if got := a.Status(); got != StateUninstalled {
		t.Errorf("expected StateUninstalled before install, got %q", got)
	}

	_ = a.Install(InstallOpts{Profile: "readonly"})

	if got := a.Status(); got != StateInstalled {
		t.Errorf("expected StateInstalled after install, got %q", got)
	}
}

// Copyright 2021 MongoDB Inc
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
//go:build e2e || brew

package brew_test

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
)

const (
	profileString = "PROFILE NAME"
	errorMessage  = "Error: this action requires authentication"
)

func TestAtlasCLIConfig(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tempDirEnv := "XDG_CONFIG_HOME=" + os.TempDir() // make sure no config.toml is detected

	t.Run("config ls", func(t *testing.T) {
		cmd := exec.Command(cliPath, "config", "ls")
		cmd.Env = append(os.Environ(), tempDirEnv)
		resp, err := e2e.RunAndGetStdOut(cmd)
		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		got := strings.TrimSpace(string(resp))
		want := profileString

		if got != want {
			t.Errorf("want %q; got %q\n", want, got)
		}
	})

	t.Run("projects ls", func(t *testing.T) {
		cmd := exec.Command(cliPath, "projects", "ls")
		cmd.Env = append(os.Environ(), tempDirEnv)
		resp, err := cmd.CombinedOutput()
		if err == nil {
			t.Fatalf("expected error, resp: %v", string(resp))
		}
		got := string(resp)

		if !strings.Contains(got, errorMessage) {
			t.Errorf("want %q; got %q\n", errorMessage, got)
		}
	})

	t.Run("help", func(t *testing.T) {
		cmd := exec.Command(cliPath, "help")
		cmd.Env = append(os.Environ(), tempDirEnv)
		if resp, err := e2e.RunAndGetStdOut(cmd); err != nil {
			t.Fatalf("unexpected error, resp: %v", string(resp))
		}
	})
}

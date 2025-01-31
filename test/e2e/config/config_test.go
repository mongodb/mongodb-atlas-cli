// Copyright 2020 MongoDB Inc
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

//go:build e2e || config

package config_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/Netflix/go-expect"
	pseudotty "github.com/creack/pty"
	"github.com/hinshun/vt10x"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/require"
)

const (
	configEntity    = "config"
	existingProfile = "e2e"
)

func TestConfig(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)
	t.Run("config", func(t *testing.T) {
		key := os.Getenv("MCLI_PRIVATE_API_KEY")
		_ = os.Unsetenv("MCLI_PRIVATE_API_KEY")
		t.Cleanup(func() {
			_ = os.Setenv("MCLI_PRIVATE_API_KEY", key)
		})
		pty, tty, err := pseudotty.Open()
		if err != nil {
			t.Fatalf("failed to open pseudotty: %v", err)
		}

		term := vt10x.New(vt10x.WithWriter(tty))
		// To debug add os.Stdout to expect.WithStdout
		c, err := expect.NewConsole(expect.WithStdin(pty), expect.WithStdout(term), expect.WithCloser(pty, tty))
		if err != nil {
			t.Fatalf("failed to create console: %v", err)
		}
		defer c.Close()

		cmd := exec.Command(cliPath, configEntity, "init", "-P", "e2e-expect")
		cmd.Stdin = c.Tty()
		cmd.Stdout = c.Tty()
		cmd.Stderr = c.Tty()
		cmd.Env = os.Environ()

		if err = cmd.Start(); err != nil {
			t.Fatal(err)
		}

		if _, err = c.ExpectString("Public API Key"); err != nil {
			t.Fatal(err)
		}
		if _, err = c.SendLine("qwerty"); err != nil {
			t.Fatalf("SendLine() = %v", err)
		}

		if _, err = c.ExpectString("Private API Key"); err != nil {
			t.Fatal(err)
		}
		if _, err = c.SendLine("qwerty"); err != nil {
			t.Fatalf("SendLine() = %v", err)
		}
		if _, err = c.ExpectString("Do you want to enter the Organization ID manually"); err != nil {
			t.Fatal(err)
		}
		if _, err = c.SendLine("y"); err != nil {
			t.Fatalf("SendLine() = %v", err)
		}
		if _, err = c.ExpectString("Default Org ID:"); err != nil {
			t.Fatal(err)
		}
		if _, err = c.SendLine("5e429f2e06822c6eac4d59c9"); err != nil {
			t.Fatalf("SendLine() = %v", err)
		}
		if _, err = c.ExpectString("Do you want to enter the Project ID manually"); err != nil {
			t.Fatal(err)
		}
		if _, err = c.SendLine("y"); err != nil {
			t.Fatalf("SendLine() = %v", err)
		}
		if _, err = c.ExpectString("Default Project ID:"); err != nil {
			t.Fatal(err)
		}
		if _, err = c.SendLine("5e429f2e06822c6eac4d59c9"); err != nil {
			t.Fatalf("SendLine() = %v", err)
		}
		if _, err = c.ExpectString("Default Output Format"); err != nil {
			t.Fatal(err)
		}
		if _, err = c.SendLine(""); err != nil {
			t.Fatalf("SendLine() = %v", err)
		}
		if err = cmd.Wait(); err != nil {
			t.Fatalf("unexpected error: %v, resp", err)
		}
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath, configEntity, "ls")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		if !strings.Contains(string(resp), existingProfile) {
			t.Errorf("expected %q to contain %q\n", string(resp), existingProfile)
		}
	})
	t.Run("Describe", func(t *testing.T) {
		// This depends on a ORG_ID ENV
		cmd := exec.Command(
			cliPath,
			configEntity,
			"describe",
			"e2e",
			"-o=json",
		)
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		var config map[string]any
		if err := json.Unmarshal(resp, &config); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if _, ok := config["org_id"]; !ok {
			t.Errorf("expected %v, to have key %s\n", config, "org_id")
		}
		if _, ok := config["service"]; !ok {
			t.Errorf("expected %v, to have key %s\n", config, "service")
		}
	})
	t.Run("Rename", func(t *testing.T) {
		cmd := exec.Command(
			cliPath,
			configEntity,
			"rename",
			"e2e",
			"renamed",
		)
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		const expected = "The profile e2e was renamed to renamed.\n"
		if string(resp) != expected {
			t.Errorf("expected %s, got %s\n", expected, string(resp))
		}
	})
	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath, configEntity, "delete", "renamed", "--force")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		const expected = "Profile 'renamed' deleted\n"
		if string(resp) != expected {
			t.Errorf("expected %s, got %s\n", expected, string(resp))
		}
	})
}

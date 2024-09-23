// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build e2e || (atlas && plugin && run)

package atlas_test

import (
	"fmt"
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

func TestPluginRun(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)

	cmd := exec.Command(cliPath,
		"plugin",
		"install",
		"mongodb/atlas-cli-plugin-example")
	resp, err := e2e.RunAndGetStdOut(cmd)
	require.NoError(t, err, string(resp))

	t.Run("Hello", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"example",
			"hello")
		resp, err := e2e.RunAndGetStdOut(cmd)
		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		helloWorlString := "Hello World!"
		if !strings.Contains(string(resp), helloWorlString) {
			t.Errorf("expected %q to contain %q\n", string(resp), helloWorlString)
		}
	})

	t.Run("Echo", func(t *testing.T) {
		echoString := "this string will be the output of the echo command"
		cmd := exec.Command(cliPath,
			"example",
			"echo",
			echoString)
		resp, err := e2e.RunAndGetStdOut(cmd)
		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		if !strings.Contains(string(resp), echoString) {
			t.Errorf("expected %q to contain %q\n", string(resp), echoString)
		}
	})

	t.Run("Printenv", func(t *testing.T) {
		var sb strings.Builder

		for _, env := range os.Environ() {
			sb.WriteString(fmt.Sprintf("\t- %s\n", env))
		}

		cmd := exec.Command(cliPath,
			"example",
			"printenv")
		resp, err := e2e.RunAndGetStdOut(cmd)
		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		if !strings.Contains(string(resp), sb.String()) {
			t.Errorf("expected %q to contain %q\n", string(resp), sb.String())
		}
	})

	t.Run("Stdinreader", func(t *testing.T) {
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

		cmd := exec.Command(cliPath,
			"example",
			"stdinreader")
		cmd.Stdin = c.Tty()
		cmd.Stdout = c.Tty()

		if err = cmd.Start(); err != nil {
			t.Fatal(err)
		}
		if _, err = c.ExpectString("Please enter your name: "); err != nil {
			t.Fatal(err)
		}
		testName := "testName"
		if _, err = c.SendLine(testName); err != nil {
			t.Fatalf("SendLine() = %v", err)
		}
		if err = cmd.Wait(); err != nil {
			t.Fatalf("unexpected error: %v, resp", err)
		}
	})
}

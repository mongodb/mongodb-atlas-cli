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

package require

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
)

func emptyRun(*cobra.Command, []string) {}

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(root, args...)
	return output, err
}

func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, buf.String(), err
}

func TestNoArgs(t *testing.T) {
	c := &cobra.Command{Use: "c", Args: NoArgs, Run: emptyRun}
	output, err := executeCommand(c)
	if output != "" {
		t.Errorf("Unexpected output: %v", output)
	}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestExactArgs(t *testing.T) {
	c := &cobra.Command{Use: "c", Args: ExactArgs(3), Run: emptyRun}
	output, err := executeCommand(c, "a", "b", "c")
	if output != "" {
		t.Errorf("Unexpected output: %v", output)
	}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestMaximumNArgs(t *testing.T) {
	c := &cobra.Command{Use: "c", Args: MaximumNArgs(3), Run: emptyRun}
	output, err := executeCommand(c, "a", "b")
	if output != "" {
		t.Errorf("Unexpected output: %v", output)
	}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestMinimumNArgs(t *testing.T) {
	c := &cobra.Command{Use: "c", Args: MinimumNArgs(2), Run: emptyRun}
	output, err := executeCommand(c, "a", "b", "c")
	if output != "" {
		t.Errorf("Unexpected output: %v", output)
	}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

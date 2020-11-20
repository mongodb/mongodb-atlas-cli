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

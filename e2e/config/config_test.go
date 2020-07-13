package config_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestProjects(t *testing.T) {
	cliPath, err := filepath.Abs("../../bin/mongocli")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, err = os.Stat(cliPath)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	const configEntity = "config"
	const existingProfile = "e2e"
	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath, configEntity, "ls")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		if !strings.Contains(string(resp), existingProfile) {
			t.Errorf("expected '%s; to contain '%s'\n", string(resp), existingProfile)
		}
	})
	t.Run("Describe", func(t *testing.T) {
		// This depends on a ORG_ID ENV
		cmd := exec.Command(cliPath, configEntity, "describe", "e2e")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		const expected = `service = cloud
org_id = 5e429e7706822c6eac4d5971
public_api_key = redacted
`
		if string(resp) != expected {
			t.Errorf("expected %s, got %s\n", expected, string(resp))
		}
	})
	t.Run("Rename", func(t *testing.T) {
		cmd := exec.Command(cliPath, configEntity, "rename", "e2e", "renamed")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		const expected = "The profile e2e was renamed to renamed.\n"
		if string(resp) != expected {
			t.Errorf("expected %s, got %s\n", expected, string(resp))
		}
	})
}

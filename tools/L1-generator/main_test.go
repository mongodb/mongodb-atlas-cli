package main

import (
	"os"
	"testing"

	"github.com/bradleyjkemp/cupaloy/v2"
	"github.com/spf13/afero"
)

func testSpec(t *testing.T, specPath string) {
	t.Helper()
	realFs := afero.NewOsFs()

	specBytes, err := afero.ReadFile(realFs, specPath)
	if err != nil {
		t.Errorf("failed to load '%s', error: %s", specPath, err)
		t.FailNow()
	}

	fs := afero.NewMemMapFs()
	_ = afero.WriteFile(fs, "spec.yml", specBytes, os.ModeType)

	if err := convertSpecToL1Commands(fs, "spec.yml", "commands.go"); err != nil {
		t.Errorf("failed to convert spec into commmands, error: %s", err)
		t.FailNow()
	}

	resultBytes, err := afero.ReadFile(fs, "commands.go")
	if err != nil {
		t.Errorf("failed to read result commands file, error: %s", err)
		t.FailNow()
	}

	resultString := string(resultBytes)
	if err := cupaloy.Snapshot(resultString); err != nil {
		t.Errorf("unexpected result %s", err)
		t.FailNow()
	}
}

// To update snapshots run: UPDATE_SNAPSHOTS=true go test ./...
func TestSnapshotFixiture00(t *testing.T) {
	testSpec(t, "./fixtures/00-spec.yaml")
}

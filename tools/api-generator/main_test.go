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

package main

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/bradleyjkemp/cupaloy/v2"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func testSpec(t *testing.T, name, specPath string) {
	t.Helper()
	snapshotter := cupaloy.New(cupaloy.SnapshotFileExtension(".snapshot"))

	realFs := afero.NewOsFs()

	specBytes, err := afero.ReadFile(realFs, specPath)
	if err != nil {
		t.Errorf("failed to load '%s', error: %s", specPath, err)
		t.FailNow()
	}

	fs := afero.NewMemMapFs()
	_ = afero.WriteFile(fs, "spec.yml", specBytes, os.ModeType)

	if err := convertSpecToAPICommands(context.Background(), fs, "spec.yml", "commands.go"); err != nil {
		t.Errorf("failed to convert spec into commmands, error: %s", err)
		t.FailNow()
	}

	resultBytes, err := afero.ReadFile(fs, "commands.go")
	if err != nil {
		t.Errorf("failed to read result commands file, error: %s", err)
		t.FailNow()
	}

	resultString := string(resultBytes)
	if err := snapshotter.SnapshotWithName(name, resultString); err != nil {
		t.Errorf("unexpected result %s", err)
		t.FailNow()
	}
}

// To update snapshots run: UPDATE_SNAPSHOTS=true go test ./...
func TestSnapshots(t *testing.T) {
	const FixtureDirectory = "./fixtures/"
	files, err := os.ReadDir(FixtureDirectory)
	require.NoError(t, err, "failed to load fixtures")

	for _, file := range files {
		fileName := file.Name()
		fullPath := filepath.Join(FixtureDirectory, fileName)
		t.Run(fileName, func(t *testing.T) {
			testSpec(t, fileName, fullPath)
		})
	}
}

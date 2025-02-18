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
	"bytes"
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/bradleyjkemp/cupaloy/v2"
)

func testSpec(t *testing.T, name, specPath string) {
	t.Helper()
	snapshotter := cupaloy.New(cupaloy.SnapshotFileExtension(".snapshot"))

	specFile, err := os.OpenFile(specPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed to load '%s', error: %s", specPath, err)
	}
	defer specFile.Close()

	overlayPath := specPath + ".overlay.yaml"
	var overlays []io.Reader
	if _, err = os.Stat(overlayPath); err == nil {
		overlayFile, err := os.OpenFile(overlayPath, os.O_RDONLY, os.ModePerm)
		if err != nil {
			t.Fatalf("failed to load '%s', error: %s", overlayPath, err)
		}
		defer overlayFile.Close()
		overlays = []io.Reader{overlayFile}
	}

	buf := &bytes.Buffer{}
	if err := convertSpecToAPICommands(context.Background(), specFile, overlays, buf); err != nil {
		t.Fatalf("failed to convert spec into commmands, error: %s", err)
	}

	if err := snapshotter.SnapshotWithName(name, buf.String()); err != nil {
		t.Fatalf("unexpected result %s", err)
	}
}

// To update snapshots run: UPDATE_SNAPSHOTS=true go test ./...
func TestSnapshots(t *testing.T) {
	const FixtureDirectory = "./fixtures/"
	files, err := os.ReadDir(FixtureDirectory)
	if err != nil {
		t.Fatalf("failed to load fixtures: %s", err)
	}

	for _, file := range files {
		fileName := file.Name()
		if strings.HasSuffix(fileName, ".overlay.yaml") {
			continue
		}
		fullPath := filepath.Join(FixtureDirectory, fileName)
		t.Run(fileName, func(t *testing.T) {
			testSpec(t, fileName, fullPath)
		})
	}
}

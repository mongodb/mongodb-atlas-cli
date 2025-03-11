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
	"fmt"
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
	overlayPath := specPath + ".overlay.yaml"

	outputFunctions := map[OutputType]func(ctx context.Context, spec io.Reader, w io.Writer) error{
		Commands: convertSpecToAPICommands,
		Metadata: convertSpecToMetadata,
	}

	for outputType, outputTypeFunc := range outputFunctions {
		specFile, err := os.OpenFile(specPath, os.O_RDONLY, os.ModePerm)
		if err != nil {
			t.Fatalf("failed to load '%s', error: %s", specPath, err)
		}
		t.Cleanup(func() {
			specFile.Close()
		})

		spec, err := applyOverlays(specFile, overlayPath)
		if err != nil {
			t.Fatalf("failed to apply overlays %q, error: %s", specPath, err)
		}

		buf := &bytes.Buffer{}
		if err := outputTypeFunc(context.Background(), spec, buf); err != nil {
			t.Fatalf("failed to convert spec into %s, error: %s", outputType, err)
		}

		if err := snapshotter.SnapshotWithName(fmt.Sprintf("%s-%s", name, outputType), buf.String()); err != nil {
			t.Fatalf("unexpected result %s", err)
		}
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

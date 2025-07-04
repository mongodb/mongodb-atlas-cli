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
	"testing"
	"time"

	"github.com/bradleyjkemp/cupaloy/v2"
)

func testSpec(t *testing.T, name, specPath string) {
	t.Helper()

	snapshotTime := time.Date(2025, time.May, 1, 0, 0, 0, 0, time.UTC)
	snapshotter := cupaloy.New(cupaloy.SnapshotSubdirectory("testdata/.snapshots"), cupaloy.SnapshotFileExtension(".snapshot"))

	outputFunctions := map[OutputType]func(ctx context.Context, now time.Time, r io.Reader, w io.Writer) error{
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

		buf := &bytes.Buffer{}
		if err := outputTypeFunc(t.Context(), snapshotTime, specFile, buf); err != nil {
			t.Fatalf("failed to convert spec into %s, error: %s", outputType, err)
		}

		if err := snapshotter.SnapshotWithName(fmt.Sprintf("%s-%s", name, outputType), buf.String()); err != nil {
			t.Fatalf("unexpected result %s", err)
		}
	}
}

// To update snapshots run: UPDATE_SNAPSHOTS=true go test ./...
func TestSnapshots(t *testing.T) {
	const FixtureDirectory = "./testdata/fixtures/"
	files, err := os.ReadDir(FixtureDirectory)
	if err != nil {
		t.Fatalf("failed to load fixtures: %s", err)
	}

	for _, file := range files {
		fileName := file.Name()
		fullPath := filepath.Join(FixtureDirectory, fileName)
		t.Run(fileName, func(t *testing.T) {
			testSpec(t, fileName, fullPath)
		})
	}
}

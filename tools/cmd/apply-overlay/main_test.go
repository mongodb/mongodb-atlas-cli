// Copyright 2025 MongoDB Inc
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
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/bradleyjkemp/cupaloy/v2"
)

func TestRun(t *testing.T) {
	snapshotter := cupaloy.New(cupaloy.SnapshotFileExtension(".snapshot"))

	const FixtureDirectory = "./fixtures/"
	files, err := os.ReadDir(FixtureDirectory)
	if err != nil {
		t.Fatalf("failed to load fixtures: %s", err)
	}

	overlayFiles, err := filepath.Glob(FixtureDirectory + "*.overlay-*.yaml")
	if err != nil {
		t.Fatalf("failed to find overlays: %s", err)
	}

	for _, file := range files {
		fileName := file.Name()
		fullPath := filepath.Join(FixtureDirectory, fileName)

		if slices.Contains(overlayFiles, fullPath) {
			continue
		}

		t.Run(fileName, func(t *testing.T) {
			buf := &bytes.Buffer{}
			if err := run(fullPath, fullPath+".overlay-*.yaml", buf); err != nil {
				t.Fatalf("failed to run: %s", err)
			}

			if err := snapshotter.SnapshotWithName(fileName, buf.String()); err != nil {
				t.Fatalf("unexpected result %s", err)
			}
		})
	}
}

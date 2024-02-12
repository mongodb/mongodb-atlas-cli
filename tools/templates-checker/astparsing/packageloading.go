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

package astparsing

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
)

func LoadPackagesRecursive(path string) ([]*packages.Package, error) {
	conf := &packages.Config{
		Mode: packages.NeedTypes |
			packages.NeedSyntax |
			packages.NeedTypesInfo,
	}

	// Calculate the absolute path, if it was already absolute nothing happens
	sourcePath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	// Extract all go files per directory
	filesPerDirectory, err := getSourceFilesPerDirectory(sourcePath)
	if err != nil {
		return nil, err
	}

	pkgs := make([]*packages.Package, 0)

	// Load every package in every directory
	for directory, files := range filesPerDirectory {
		// Show progress to the user, loading all packages can take a while
		fmt.Printf("Loading dir: %v\n", directory)

		// Load the packages from the directory (should only load one, but multiple will work too)
		pkgsInDir, err := packages.Load(conf, files...)

		if err != nil {
			return nil, err
		}

		pkgs = append(pkgs, pkgsInDir...)
	}

	return pkgs, nil
}

// Returns all .go files per directory (recursively)
// filenames ending on _test.go are excluded
//
// Example:
// - root/child1/ : [ "file1.go" ]
// - root/child2/ : [ "file2.go", "file3.go" ].
func getSourceFilesPerDirectory(sourcePath string) (sourceFiles map[string][]string, err error) {
	sourceFiles = make(map[string][]string)

	err = filepath.WalkDir(sourcePath, func(s string, _ fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Get the base of the file, example "file1.go"
		fileName := filepath.Base(s)

		// Only process the file if:
		// - it's a go file
		// - it's not a go test file
		if strings.HasSuffix(fileName, ".go") && !strings.HasSuffix(fileName, "_test.go") {
			dir := filepath.Dir(s)

			// Create the directory slice if needed
			if sourceFiles[dir] == nil {
				sourceFiles[dir] = make([]string, 0)
			}

			// Extend the slice with the current file
			sourceFiles[dir] = append(sourceFiles[dir], s)
		}
		return nil
	})

	return
}

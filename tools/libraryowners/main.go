// Copyright 2022 MongoDB Inc
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

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/mod/modfile"
)

const (
	libraryOwnersPath = "build/ci/library_owners.json"
	goModpath         = "go.mod"
)

func newGoMod() (*modfile.File, error) {
	goModFile, err := os.ReadFile(goModpath)
	if err != nil {
		return nil, err
	}

	f, err := modfile.Parse(goModpath, goModFile, nil)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func newLibraryOwners() (map[string]string, error) {
	libraryOwnersFile, err := os.ReadFile(libraryOwnersPath)
	if err != nil {
		return nil, err
	}

	if len(libraryOwnersFile) == 0 {
		return nil, fmt.Errorf("'%s' is empty", libraryOwnersPath)
	}

	var libraryOwnersContent map[string]string
	if err = json.Unmarshal(libraryOwnersFile, &libraryOwnersContent); err != nil {
		return nil, err
	}

	return libraryOwnersContent, nil
}

func validateNewLibs(libOwners map[string]string, goMod *modfile.File) error {
	var addedLibs []string
	for _, library := range goMod.Require {
		if library.Indirect {
			continue
		}
		if val, ok := libOwners[library.Mod.Path]; !ok || val == "" {
			addedLibs = append(addedLibs, library.Mod.Path)
		}
	}
	if len(addedLibs) != 0 {
		return fmt.Errorf("%q doesn't have an owner, please, add one to %q", addedLibs, libraryOwnersPath)
	}
	return nil
}

func validateRemovedLibs(libOwners map[string]string, goMod *modfile.File) error {
	var removedLibs []string
	for library := range libOwners {
		var found bool
		for _, dep := range goMod.Require {
			if library == dep.Mod.Path {
				found = true
				break
			}
		}
		if !found {
			removedLibs = append(removedLibs, library)
		}
		found = false
	}
	if len(removedLibs) != 0 {
		return fmt.Errorf("%q are not defined, please remove from %q", removedLibs, libraryOwnersPath)
	}
	return nil
}

func main() {
	libraryOwners, err := newLibraryOwners()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error while parsing '%s': %v\n", libraryOwnersPath, err)
		os.Exit(1)
	}

	goMod, err := newGoMod()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error while parsing '%s': %v\n", goModpath, err)
		os.Exit(1)
	}

	if err := validateNewLibs(libraryOwners, goMod); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error during the validation of '%s': %v\n", libraryOwnersPath, err)
		os.Exit(1)
	}

	if err := validateRemovedLibs(libraryOwners, goMod); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error during the validation of '%s': %v\n", libraryOwnersPath, err)
		os.Exit(1)
	}
}

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

func validate(libraryOwner map[string]string, goMod *modfile.File) error {
	for _, library := range goMod.Require {
		if library.Indirect {
			continue
		}

		if val, ok := libraryOwner[library.Mod.Path]; !ok {
			return fmt.Errorf("'%s' is not inside '%s'. Please, remove this dependency from '%s'", library.Mod.Path, libraryOwnersPath, goModpath)
		} else if val == "" {
			return fmt.Errorf("'%s' doesn't have a owner. Please, add a owner to %s in '%s", library.Mod.Path, library.Mod.Path, libraryOwnersPath)
		}
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

	if err := validate(libraryOwners, goMod); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error during the validation of '%s': %v\n", libraryOwnersPath, err)
		os.Exit(1)
	}
}

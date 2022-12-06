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
	"log"
	"os"

	"golang.org/x/mod/modfile"
)

const libraryOwnersPath = "library_owners.json"
const goModpath = "go.mod"

func newGoMod() *modfile.File {
	goModFile, err := os.ReadFile(goModpath)
	if err != nil {
		log.Fatalln(err)
	}

	f, err := modfile.Parse(goModpath, goModFile, nil)
	if err != nil {
		log.Fatalln(err)
	}
	return f
}

func newLibraryOwners() map[string]interface{} {
	libraryOwnersFile, err := os.ReadFile(libraryOwnersPath)
	if err != nil {
		log.Fatalln(err)
	}

	if len(libraryOwnersFile) == 0 {
		log.Fatalf("\n'%s' is empty", libraryOwnersPath)
	}

	var libraryOwnersContent map[string]interface{}
	err = json.Unmarshal(libraryOwnersFile, &libraryOwnersContent)
	if err != nil {
		log.Fatal("\nError during Unmarshal(): ", err)
	}
	return libraryOwnersContent
}

func validate(libraryOwner map[string]interface{}, goMod *modfile.File) {
	for _, library := range goMod.Require {
		if library.Indirect {
			continue
		}

		if val, ok := libraryOwner[library.Mod.Path]; !ok {
			log.Fatalf("\n'%s' is not inside '%s'. Please, add this dependency to '%s'.", library.Mod.Path, libraryOwnersPath, libraryOwnersPath)
		} else if val == "" {
			log.Fatalf("\n'%s' doesn't have a owner. Please, add a owner to %s in '%s'.", library.Mod.Path, library.Mod.Path, libraryOwnersPath)
		}
	}
}

func main() {
	validate(newLibraryOwners(), newGoMod())
}

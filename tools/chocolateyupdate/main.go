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
	"flag"
	"fmt"
	"log"
	"os"

	exec "golang.org/x/sys/execabs"
)

func update(path, version, secret string) error {
	packageName := fmt.Sprintf("atlascli.%s.nupkg", version)
	cmd := exec.Command("choco", "push", packageName, "--api-key", secret)
	cmd.Dir = path
	err := cmd.Start()
	return err
}

func main() {
	var version string
	const packagePath = "build/package/chocolatey/temp"
	secret := os.Getenv("SECRET_API_KEY")

	flag.StringVar(&version, "version", "", "Atlas CLI version")
	flag.Parse()

	if version == "" {
		log.Fatalln("You must specify Atlas CLI version")
	}

	err := update(packagePath, version, secret)
	if err != nil {
		log.Fatal(err)
	}
}

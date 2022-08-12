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
	"log"
	"os"

	exec "golang.org/x/sys/execabs"
)

func update(path, name, secret string) error {
	cmd := exec.Command("choco", "push", name, "--api-key", secret)
	cmd.Dir = path
	err := cmd.Start()
	return err
}

func main() {
	var packagePath, packageName string
	secret := os.Getenv("SECRET_API_KEY")

	flag.StringVar(&packagePath, "path", "", "Chocolatey package path")
	flag.StringVar(&packageName, "name", "", "Chocolatey package name")
	flag.Parse()

	if packagePath == "" {
		log.Fatalln("You must specify Chocolatey package path")
	}
	if packageName == "" {
		log.Fatalln("You must specify Chocolatey package name")
	}

	err := update(packagePath, packageName, secret)
	if err != nil {
		log.Fatal(err)
	}
}

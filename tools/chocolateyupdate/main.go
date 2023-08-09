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
	"os/exec"
)

func update(path, secret string, sourcePath string) error {
	cmd := exec.Command("choco", "push", path, "--api-key", secret, "--source", sourcePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}

func main() {
	var packagePath string
	var sourcePath string
	secret := os.Getenv("SECRET_API_KEY")

	flag.StringVar(&packagePath, "path", "", "Chocolatey package path")
	flag.StringVar(&sourcePath, "source", "https://push.chocolatey.org/", "Chocolatey source path")
	flag.Parse()

	if packagePath == "" {
		log.Fatalln("You must specify Chocolatey package path")
	}

	err := update(packagePath, secret, sourcePath)
	if err != nil {
		log.Fatal(err)
	}
}

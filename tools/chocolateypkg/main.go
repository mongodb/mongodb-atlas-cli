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
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"text/template"

	exec "golang.org/x/sys/execabs"
)

type NuspecDetails struct {
	Version string
}

type InstallScriptDetails struct {
	URL      string
	CheckSum string
}

func createDirectory(path, name string) error {
	dirLocation := fmt.Sprintf("%s/%s", path, name)
	err := os.MkdirAll(dirLocation, os.ModePerm)
	return err
}

func createFile(name string) (f *os.File, err error) {
	f, err = os.Create(name)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func replaceNuspec(path, version string) error {
	nuspecPath := fmt.Sprintf("%s/atlascli.nuspec", path)
	newVersion := NuspecDetails{version}

	p, err := os.ReadFile(nuspecPath)
	if err != nil {
		return err
	}
	tmpl, err := template.New("NuspecTemplate").Parse(string(p))
	if err != nil {
		return err
	}
	var generatedNuspec bytes.Buffer
	err = tmpl.Execute(&generatedNuspec, newVersion)
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/temp/atlascli.nuspec", path)
	f, err := createFile(filePath)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		fileErr := f.Close()
		if fileErr != nil {
			log.Fatal(fileErr)
		}
	}(f)
	_, err = f.Write(generatedNuspec.Bytes())
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	return nil
}

func generateSha256(f *os.File) (string, error) {
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func replaceInstallScript(path, msiPath, url string) error {
	scriptPath := fmt.Sprintf("%s/tools/chocolateyinstall.ps1", path)

	f, err := os.Open(msiPath)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		fileErr := f.Close()
		if fileErr != nil {
			log.Fatal("Error when closing file.")
		}
	}(f)

	checkSum, err := generateSha256(f)
	if err != nil {
		return err
	}
	newInstallDetails := InstallScriptDetails{
		URL:      url,
		CheckSum: checkSum,
	}

	p, err := os.ReadFile(scriptPath)
	if err != nil {
		return err
	}
	tmpl, err := template.New("InstallScriptTemplate").Parse(string(p))
	if err != nil {
		return err
	}
	var generatedScript bytes.Buffer
	err = tmpl.Execute(&generatedScript, newInstallDetails)
	if err != nil {
		return err
	}

	newDirectoryPath := fmt.Sprintf("%s/temp", path)
	err = createDirectory(newDirectoryPath, "tools")
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/temp/tools/chocolateyinstall.ps1", path)
	f, err = createFile(filePath)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		fileErr := f.Close()
		if fileErr != nil {
			log.Fatal(fileErr)
		}
	}(f)
	_, err = f.Write(generatedScript.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func main() {
	args := os.Args[1:]
	const argumentsLength = 3
	if len(args) < argumentsLength {
		log.Fatal("Exactly 3 arguments must be present to run Chocolatey update script.")
	}

	const path = "build/package/chocolatey"

	err := createDirectory(path, "temp")
	checkError(err)

	version := args[0]
	err = replaceNuspec(path, version)
	checkError(err)

	msiPath := args[1]
	downloadURL := args[2]
	err = replaceInstallScript(path, msiPath, downloadURL)
	checkError(err)

	const chocoCommand = "pack"
	cmd := exec.Command("choco", chocoCommand)
	cmd.Dir = path + "/temp"
	err = cmd.Start()
	checkError(err)
}

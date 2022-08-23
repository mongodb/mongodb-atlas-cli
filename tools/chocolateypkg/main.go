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
	"context"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"path"
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

func createDirectory(dir, name string) error {
	dirLocation := path.Join(dir, name)
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

func replaceNuspec(dir, version string) error {
	nuspecPath := path.Join(dir, "mongodb-atlas.nuspec")
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

	filePath := path.Join(dir, "temp", "mongodb-atlas.nuspec")
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

func generateSha256(url string) (string, error) {
	r, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, http.NoBody)
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	h := sha256.New()
	if _, err := io.Copy(h, resp.Body); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func replaceInstallScript(dir, url string) error {
	scriptPath := path.Join(dir, "tools", "chocolateyinstall.ps1")

	checkSum, err := generateSha256(url)
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

	newDirectoryPath := path.Join(dir, "temp")
	err = createDirectory(newDirectoryPath, "tools")
	if err != nil {
		return err
	}

	filePath := path.Join(dir, "temp", "tools", "chocolateyinstall.ps1")
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
	_, err = f.Write(generatedScript.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var version, downloadURL, srcPath string

	flag.StringVar(&version, "version", "", "Atlas CLI version")
	flag.StringVar(&srcPath, "srcPath", "", "Path to templates")
	flag.StringVar(&downloadURL, "url", "", "URL to download Atlas CLI installer")
	flag.Parse()

	if version == "" {
		log.Fatalln("You must specify Atlas CLI version")
	}
	if downloadURL == "" {
		log.Fatalln("You must specify download URL")
	}

	err := createDirectory(srcPath, "temp")
	checkError(err)

	err = replaceNuspec(srcPath, version)
	checkError(err)

	err = replaceInstallScript(srcPath, downloadURL)
	checkError(err)

	const chocoCommand = "pack"
	cmd := exec.Command("choco", chocoCommand)
	cmd.Dir = path.Join(srcPath, "temp")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	checkError(err)
}

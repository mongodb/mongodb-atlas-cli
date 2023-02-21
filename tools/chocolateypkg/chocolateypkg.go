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

/*
chocolateypkg generates chocolatey package information.

Usage:

	chocolateypkg [flags]

The flags are:

	-version
		Atlas CLI version to package
	-out
		Output folder for files
	-url
		Atlas CLI installer URL to download
*/
package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	_ "embed"
	"encoding/hex"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"text/template"
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
	return os.MkdirAll(dirLocation, os.ModePerm)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//go:embed mongodb-atlas.nuspec
var atlasNuSpec string

func replaceNuspec(dir, version string) error {
	newVersion := NuspecDetails{version}
	tmpl, err := template.New("NuspecTemplate").Parse(atlasNuSpec)
	if err != nil {
		return err
	}
	var generatedNuspec bytes.Buffer
	if err = tmpl.Execute(&generatedNuspec, newVersion); err != nil {
		return err
	}

	filePath := path.Join(dir, "mongodb-atlas.nuspec")
	f, err := os.Create(filePath)
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
	return err
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

//go:embed chocolateyinstall.ps1
var installScript string

func replaceInstallScript(dir, url string) error {
	checkSum, err := generateSha256(url)
	if err != nil {
		return err
	}
	newInstallDetails := InstallScriptDetails{
		URL:      url,
		CheckSum: checkSum,
	}

	tmpl, err := template.New("InstallScriptTemplate").Parse(installScript)
	if err != nil {
		return err
	}
	var generatedScript bytes.Buffer
	if err = tmpl.Execute(&generatedScript, newInstallDetails); err != nil {
		return err
	}

	if err = createDirectory(dir, "tools"); err != nil {
		return err
	}

	filePath := path.Join(dir, "tools", "chocolateyinstall.ps1")
	f, err := os.Create(filePath)
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

	return err
}

func main() {
	var (
		version     string
		downloadURL string
		outPath     string
	)

	flag.StringVar(&version, "version", "", "Atlas CLI version to package")
	flag.StringVar(&outPath, "out", "dist", "Output folder for files")
	flag.StringVar(&downloadURL, "url", "", "Atlas CLI installer URL to download")
	flag.Parse()

	if version == "" {
		log.Fatalln("You must specify Atlas CLI version")
	}
	if downloadURL == "" {
		log.Fatalln("You must specify download URL")
	}

	err := replaceNuspec(outPath, version)
	checkError(err)

	err = replaceInstallScript(outPath, downloadURL)
	checkError(err)
	log.Println("Success!")
}

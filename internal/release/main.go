// Copyright 2020 MongoDB Inc
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
	"time"
)

type DownloadArchive struct {
	PreviousReleasesLink string     `json:"previous_releases_link"`
	ReleaseDate          time.Time  `json:"release_date"`
	ReleaseNotesLink     string     `json:"release_notes_link"`
	TutorialLink         string     `json:"tutorial_link"`
	Version              string     `json:"version"`
	ManualLink           string     `json:"manual_link"`
	Platform             []Platform `json:"platform"`
}

type Platform struct {
	Arch          string  `json:"arch"`
	OS            string  `json:"os"`
	PackageFormat string  `json:"package_format"`
	Packages      Package `json:"packages"`
}

type Package struct {
	Title string `json:"title"`
	Links []Link `json:"links"`
}

type Link struct {
	DownloadLink string `json:"download_link"`
	Name         string `json:"name"`
}

func newPlatform(version, arch, system, distro, format string) *Platform { //nolint:unparam // Arch is x86_64 for now
	p := &Platform{}
	p.Arch = arch
	p.OS = distro
	p.PackageFormat = format
	p.Packages = Package{
		Title: "MongoDB CLI",
		Links: []Link{
			{
				DownloadLink: fmt.Sprintf("https://fastdl.mongodb.org/mongocli/mongocli_%s_%s_%s.%s", version, system, arch, format),
				Name:         format,
			},
		},
	}
	return p
}

func main() {
	version := os.Args[1]
	feedFilename := fmt.Sprintf("mongocli_%s.json", version)
	fmt.Printf("Generating JSON: %s\n", feedFilename)
	err := generateFile(feedFilename, version)

	if err != nil {
		fmt.Printf("error encoding file: %v\n", err)

		os.Exit(1)
	}
	fmt.Printf("File %s has been generated\n", feedFilename)
}

func generateFile(name, version string) error {
	feedFile, err := os.Create(name)
	if err != nil {
		return err
	}
	downloadArchive := &DownloadArchive{
		ReleaseDate:          time.Now().UTC(),
		ManualLink:           fmt.Sprintf("https://docs.mongodb.com/mongocli/v%s/", version),
		PreviousReleasesLink: "https://github.com/mongodb/mongocli/releases",
		ReleaseNotesLink:     fmt.Sprintf("https://docs.mongodb.com/mongocli/v%s/release-notes/", version),
		TutorialLink:         fmt.Sprintf("https://docs.mongodb.com/mongocli/v%s/quick-start/", version),
		Platform: []Platform{
			*newPlatform(version, "x86_64", "linux", "Debian 9 / Ubuntu 16.04 + 18.04", "deb"),
			*newPlatform(version, "x86_64", "linux", "Red Hat + CentOS 6, 7, 8 / SUSE 12 + 15 / Amazon Linux", "rpm"),
			*newPlatform(version, "x86_64", "windows", "Microsoft Windows", "zip"),
			*newPlatform(version, "x86_64", "macos", "macOS", "zip"),
			*newPlatform(version, "x86_64", "linux", "Linux (x86_64)", "tar.gz"),
		},
	}
	defer feedFile.Close()
	jsonEncoder := json.NewEncoder(feedFile)
	jsonEncoder.SetIndent("", "  ")
	return jsonEncoder.Encode(downloadArchive)
}

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
	"time"

	"github.com/mongodb/mongocli/internal/release"
)

func generateFile(name, version string) error {
	packageName := "mongodb-atlas"
	feedFile, err := os.Create(name)
	if err != nil {
		return err
	}
	defer feedFile.Close()
	downloadArchive := &release.DownloadArchive{
		ReleaseDate:          time.Now().UTC(),
		Version:              version,
		ManualLink:           fmt.Sprintf("https://docs.mongodb.com/mongocli/v%s/", version),
		PreviousReleasesLink: "https://github.com/mongodb/mongocli/releases",
		ReleaseNotesLink:     fmt.Sprintf("https://docs.mongodb.com/mongocli/v%s/release-notes/", version),
		TutorialLink:         fmt.Sprintf("https://docs.mongodb.com/mongocli/v%s/quick-start/", version),
		Platform: []release.Platform{
			*release.NewPlatform(packageName, version, "x86_64", "linux", "Linux (x86_64)", []string{"tar.gz"}),
			*release.NewPlatform(packageName, version, "x86_64", "linux", "Debian 9, 10 / Ubuntu 18.04, 20.04", []string{"deb"}),
			*release.NewPlatform(packageName, version, "x86_64", "linux", "Red Hat + CentOS 6, 7, 8 / SUSE 12 + 15 / Amazon Linux", []string{"rpm"}),
			*release.NewPlatform(packageName, version, "x86_64", "windows", "Microsoft Windows", []string{"zip", "msi"}),
			*release.NewPlatform(packageName, version, "x86_64", "macos", "macOS (x86_64)", []string{"zip"}),
			*release.NewPlatform(packageName, version, "arm64", "macos", "macOS (arm64)", []string{"zip"}),
		},
	}
	jsonEncoder := json.NewEncoder(feedFile)
	jsonEncoder.SetIndent("", "  ")
	return jsonEncoder.Encode(downloadArchive)
}

func main() {
	version := os.Args[1]
	feedFilename := "mongodb-atlas.json"
	fmt.Printf("Generating JSON: %s\n", feedFilename)
	err := generateFile(feedFilename, version)

	if err != nil {
		fmt.Printf("error encoding file: %v\n", err)

		os.Exit(1)
	}
	fmt.Printf("File %s has been generated\n", feedFilename)
}

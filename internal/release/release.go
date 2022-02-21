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

package release

import (
	"fmt"
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
	Arch          string  `json:"arch,omitempty"`
	OS            string  `json:"os"`
	PackageFormat string  `json:"package_format,omitempty"`
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

func NewPlatform(tool, version, arch, system, distro string, formats []string) *Platform {
	p := &Platform{}
	p.Arch = arch
	p.OS = distro

	links := make([]Link, len(formats))
	for i, f := range formats {
		links[i] = Link{
			DownloadLink: fmt.Sprintf("https://fastdl.mongodb.org/mongocli/%s_%s_%s_%s.%s", tool, version, system, arch, f),
			Name:         f,
		}
	}

	p.Packages = Package{
		Title: "MongoDB CLI",
		Links: links,
	}
	return p
}

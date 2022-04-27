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

package telemetry

import (
	"encoding/json"
	"errors"
	"os"
	"path"

	"github.com/mongodb/mongocli/internal/config"
	"github.com/spf13/afero"
)

const (
	dirPermissions   = 0700
	filePermissions  = 0600
	cacheFilename    = "telemetry"
	maxCacheFileSize = 100 * 1024 * 1024 // 100MB
)

var fs = afero.NewOsFs()

type Event struct {
	Timestamp  string            `json:"timestamp"`
	Source     string            `json:"source"`
	Name       string            `json:"name"`
	Properties map[string]string `json:"properties"`
}

func (event *Event) Cache() error {
	if !config.TelemetryEnabled() {
		return nil
	}
	file, err := openCacheFile()
	if err != nil {
		return err
	}
	defer file.Close()
	data, _ := json.MarshalIndent(event, "", "\t")
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func openCacheFile() (afero.File, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return nil, err
	}
	cacheDir = path.Join(cacheDir, "atlascli")
	exists, err := afero.DirExists(fs, cacheDir)
	if err != nil {
		return nil, err
	}
	if !exists {
		if mkdirError := fs.MkdirAll(cacheDir, dirPermissions); mkdirError != nil {
			return nil, mkdirError
		}
	}
	filename := path.Join(cacheDir, cacheFilename)
	exists, err = afero.Exists(fs, filename)
	if err != nil {
		return nil, err
	}
	if exists {
		info, statError := fs.Stat(filename)
		if statError != nil {
			return nil, statError
		}
		size := info.Size()
		if size > maxCacheFileSize {
			return nil, errors.New("telemetry cache file too large")
		}
	}
	file, err := fs.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, filePermissions)
	return file, err
}

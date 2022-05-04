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
	"path/filepath"
	"strings"
	"time"

	"github.com/mongodb/mongocli/internal/config"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

const (
	cacheFilename   = "telemetry"
	dirPermissions  = 0700
	filePermissions = 0600
)

var fs = afero.NewOsFs()
var maxCacheFileSize int64 = 100_000_000 // 100MB
var startTime = time.Now()

type Event struct {
	Timestamp  string                 `json:"timestamp"`
	Source     string                 `json:"source"`
	Name       string                 `json:"name"`
	Properties map[string]interface{} `json:"properties"`
}

func TrackCommand(cmd *cobra.Command) {
	if !config.TelemetryEnabled() {
		return
	}
	track(cmd)
}

func newEvent(cmd *cobra.Command) Event {
	now := time.Now()
	cmdPath := cmd.CommandPath()
	command := strings.ReplaceAll(cmdPath, " ", "-")

	duration := now.Sub(startTime)

	var properties = map[string]interface{}{
		"command":  command,
		"duration": duration.Milliseconds(),
		"result":   "SUCCESS",
	}
	var event = Event{
		Timestamp:  now.Format(time.RFC3339Nano),
		Source:     config.ToolName,
		Name:       config.ToolName + "-event",
		Properties: properties,
	}

	return event
}

func track(cmd *cobra.Command) {
	event := newEvent(cmd)

	cacheDir, err := os.UserCacheDir()
	if err != nil {
		logError(err)
		return
	}
	cacheDir = filepath.Join(cacheDir, config.ToolName)
	err = save(event, cacheDir)
	if err != nil {
		logError(err)
		return
	}
}

func openCacheFile(cacheDir string) (afero.File, error) {
	exists, err := afero.DirExists(fs, cacheDir)
	if err != nil {
		return nil, err
	}
	if !exists {
		if mkdirError := fs.MkdirAll(cacheDir, dirPermissions); mkdirError != nil {
			return nil, mkdirError
		}
	}
	filename := filepath.Join(cacheDir, cacheFilename)
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

func save(event Event, cacheDir string) error {
	file, err := openCacheFile(cacheDir)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	return err
}

func logError(err error) {
	// No-op function until logging is implemented (CLOUDP-110988)
	_ = err
}

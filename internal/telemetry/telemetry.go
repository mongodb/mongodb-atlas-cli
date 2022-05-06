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
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/mongodb/mongocli/internal/config"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

const (
	cacheFilename           = "telemetry"
	dirPermissions          = 0700
	filePermissions         = 0600
	defaultMaxCacheFileSize = 100_000_000 // 100MB
)

var contextKey = telemetryContextKey{}

type telemetryContextKey struct{}

type telemetryContextValue struct {
	startTime time.Time
}

func NewContext() context.Context {
	return context.WithValue(context.Background(), contextKey, telemetryContextValue{
		startTime: time.Now(),
	})
}

type tracker struct {
	fs               afero.Fs
	maxCacheFileSize int64
	cacheDir         string
}

func newTracker() (*tracker, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return nil, err
	}

	cacheDir = filepath.Join(cacheDir, config.ToolName)

	return &tracker{
		fs:               afero.NewOsFs(),
		maxCacheFileSize: defaultMaxCacheFileSize,
		cacheDir:         cacheDir,
	}, nil
}

func TrackCommand(cmd *cobra.Command) {
	if !config.TelemetryEnabled() {
		return
	}
	t, err := newTracker()
	if err != nil {
		logError(err)
		return
	}
	if err = t.track(cmd, nil); err != nil {
		logError(err)
	}
}

func TrackCommandError(cmd *cobra.Command, e error) {
	if !config.TelemetryEnabled() {
		return
	}
	t, err := newTracker()
	if err != nil {
		logError(err)
		return
	}
	if err = t.track(cmd, e); err != nil {
		logError(err)
	}
}

func (t *tracker) track(cmd *cobra.Command, e error) error {
	options := []eventOpt{withCommandPath(cmd), withDuration(cmd), withFlags(cmd), withProfile(), withVersion(), withOS(), withAuthMethod(), withService(), withProjectID(cmd), withOrgID(cmd), withTerminal(), withInstaller(t.fs)}

	if e != nil {
		options = append(options, withError(e))
	}

	event := newEvent(options...)

	return t.save(event)
}

func (t *tracker) openCacheFile() (afero.File, error) {
	exists, err := afero.DirExists(t.fs, t.cacheDir)
	if err != nil {
		return nil, err
	}
	if !exists {
		if mkdirError := t.fs.MkdirAll(t.cacheDir, dirPermissions); mkdirError != nil {
			return nil, mkdirError
		}
	}
	filename := filepath.Join(t.cacheDir, cacheFilename)
	exists, err = afero.Exists(t.fs, filename)
	if err != nil {
		return nil, err
	}
	if exists {
		info, statError := t.fs.Stat(filename)
		if statError != nil {
			return nil, statError
		}
		size := info.Size()
		if size > t.maxCacheFileSize {
			return nil, errors.New("telemetry cache file too large")
		}
	}
	file, err := t.fs.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, filePermissions)
	return file, err
}

func (t *tracker) save(event Event) error {
	file, err := t.openCacheFile()
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

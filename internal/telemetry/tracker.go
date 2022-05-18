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
	"strconv"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/spf13/afero"
)

const (
	cacheFilename           = "telemetry"
	dirPermissions          = 0700
	filePermissions         = 0600
	defaultMaxCacheFileSize = 10_000_000 // 10MB
	defaultMaxBatchSize     = 100
)

type tracker struct {
	fs               afero.Fs
	maxCacheFileSize int64
	cacheDir         string
	store            store.EventsSender
	maxBatchSize     int
}

func newTracker(ctx context.Context) (*tracker, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return nil, err
	}

	cacheDir = filepath.Join(cacheDir, config.ToolName)

	telemetryStore, err := store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx), store.Telemetry())
	if err != nil {
		return nil, err
	}

	return &tracker{
		fs:               afero.NewOsFs(),
		maxCacheFileSize: defaultMaxCacheFileSize,
		cacheDir:         cacheDir,
		store:            telemetryStore,
		maxBatchSize:     defaultMaxBatchSize,
	}, nil
}

func (t *tracker) track(data TrackOptions) error {
	options := []eventOpt{withCommandPath(data.Cmd), withDuration(data.Cmd), withFlags(data.Cmd), withProfile(), withVersion(), withOS(), withAuthMethod(), withService(), withProjectID(data.Cmd), withOrgID(data.Cmd), withTerminal(), withInstaller(t.fs), withExtraProps(data.extraProps)}

	if data.Err != nil {
		options = append(options, withError(data.Err))
	}
	event := newEvent(options...)
	err := t.save(event)
	if err != nil {
		// If the event cannot be cached, at least make an effort to send it
		return t.store.SendEvents(&[]Event{event})
	}
	events, err := t.read(t.maxBatchSize)
	if err != nil {
		return err
	}
	err = t.store.SendEvents(events)
	if err != nil {
		return err
	}
	return t.remove(t.maxBatchSize)
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

// Append a single event to the cache file.
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

// Read the first n events from the cache file.
func (t *tracker) read(n int) (*[]Event, error) {
	events := make([]Event, 0, n)
	filename := filepath.Join(t.cacheDir, cacheFilename)
	exists, err := afero.Exists(t.fs, filename)
	if err != nil {
		return nil, err
	}
	if exists {
		// Read the first n events...
		file, err := t.fs.Open(filename)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		decoder := json.NewDecoder(file)
		for i := 0; i < n && decoder.More(); i++ {
			var event Event
			err = decoder.Decode(&event) // Reads the next JSON-encoded value
			if err != nil {
				return nil, err
			}
			events = append(events, event)
		}
	}
	return &events, nil
}

// Remove the first n events from the cache file.
func (t *tracker) remove(n int) error {
	filename := filepath.Join(t.cacheDir, cacheFilename)
	exists, err := afero.Exists(t.fs, filename)
	if !exists || err != nil {
		// Nothing to do
		return nil
	}
	file, err := t.fs.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	tmpFilename := filepath.Join(t.cacheDir, timeString())
	tmpFile, err := t.fs.OpenFile(tmpFilename, os.O_WRONLY|os.O_CREATE, filePermissions)
	if err != nil {
		return err
	}
	defer tmpFile.Close()
	decoder := json.NewDecoder(file)
	for i := 0; decoder.More(); i++ {
		var event Event
		err = decoder.Decode(&event) // Reads the next JSON-encoded value
		if err != nil {
			return err
		}
		if i < n {
			continue
		}
		data, e := json.Marshal(event)
		if e != nil {
			return e
		}
		_, e = tmpFile.Write(data)
		if e != nil {
			return e
		}
	}
	err = t.fs.Rename(tmpFilename, filename)
	if err != nil {
		return err
	}
	return nil
}

// Used to generate a unique temporary filename.
func timeString() string {
	base := 32
	nanos := time.Now().UnixNano()
	return strconv.FormatInt(nanos, base)
}

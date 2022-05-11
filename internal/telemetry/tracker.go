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
	"net/http"
	"os"
	"path/filepath"

	"github.com/mongodb/mongocli/internal/store"

	"github.com/mongodb/mongocli/internal/config"
	"github.com/spf13/afero"
)

const (
	cacheFilename           = "telemetry"
	dirPermissions          = 0700
	filePermissions         = 0600
	defaultMaxCacheFileSize = 100_000_000 // 100MB
	urlPath                 = "api/private/v1.0/telemetry/events"
)

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

func (t *tracker) track(data TrackOptions) error {
	options := []eventOpt{withCommandPath(data.Cmd), withDuration(data.Cmd), withFlags(data.Cmd), withProfile(), withVersion(), withOS(), withAuthMethod(), withService(), withProjectID(data.Cmd), withOrgID(data.Cmd), withTerminal(), withInstaller(t.fs), withExtraProps(data.extraProps)}

	if data.Err != nil {
		options = append(options, withError(data.Err))
	}

	event := newEvent(options...)
	err := t.send(data.Cmd.Context(), &[]Event{event})
	if err != nil {
		// Could not send the event, so log the error and cache the event
		logError(err)
		return t.save(event)
	}

	return nil
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

func (t *tracker) send(ctx context.Context, events *[]Event) error {
	if config.Service() != config.CloudService {
		// Only send events to Atlas - not to AtlasGov or OpsManager or CloudManager
		return nil
	}
	s, err := store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
	if err != nil {
		return err
	}
	client, err := s.GetAtlasClient()
	if err != nil {
		return err
	}
	request, err := client.NewRequest(ctx, http.MethodPost, urlPath, &events)
	if err != nil {
		return err
	}
	_, err = client.Do(ctx, request, nil)
	return err
}

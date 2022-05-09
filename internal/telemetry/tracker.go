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
	"fmt"
	"os"
	"path/filepath"

	"github.com/mongodb/mongocli/internal/validate"

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

func (t *tracker) track(cmd *cobra.Command, e error) error {
	options := []eventOpt{withCommandPath(cmd), withDuration(cmd), withFlags(cmd), withProfile(), withVersion(), withOS(), withAuthMethod(), withService(), withProjectID(cmd), withOrgID(cmd), withTerminal(), withInstaller(t.fs)}

	if e != nil {
		options = append(options, withError(e))
	}

	event := newEvent(options...)

	// TODO: If no profile, then just save to cache and return
	// TODO: Else send each event in the cache (in batches?), delete the cache, and send this event

	fmt.Printf("*** config.Name: %s\n", config.Name())
	err := validate.Credentials()
	if err != nil {
		// Either there is no profile, or the profile has an invalid token, or it has neither token nor API keys.
		// Effectively, no profile is in effect to make any endpoint calls, so cache the event...
		// TODO: Will this ever be reached? Without credentials the command will fail...
		fmt.Println("*** No credentials - caching event...")
		return t.save(event)
	}

	err = t.send(event)
	if err != nil {
		logError(err)
		// TODO: If we cannot send an event, then it must be cached...
	}

	// return nil
	fmt.Println("*** Sending events not yet implemented - caching event...")
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

func (t *tracker) send(event Event) error {
	// TODO: Find url to send the event to (Atlas or AtlasGov)
	// If not Atlas, then simply return (ie.don't even send to AtlasGov)
	fmt.Printf("*** (TODO) Sending event to Atlas telemetry endpoint: %+v\n", event)
	return nil
}

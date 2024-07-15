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

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

const (
	cacheFilename           = "telemetry"
	dirPermissions          = 0700
	filePermissions         = 0600
	defaultMaxCacheFileSize = 500_000 // 500KB
)

type tracker struct {
	fs               afero.Fs
	maxCacheFileSize int64
	cacheDir         string
	unauthStore      store.UnauthEventsSender
	store            store.EventsSender
	storeSet         bool
	cmd              *cobra.Command
	args             []string
	installer        *string
}

func newTracker(ctx context.Context, cmd *cobra.Command, args []string) (*tracker, error) {
	var err error

	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return nil, err
	}

	cacheDir = filepath.Join(cacheDir, config.AtlasCLI)

	t := &tracker{
		fs:               afero.NewOsFs(),
		maxCacheFileSize: defaultMaxCacheFileSize,
		cacheDir:         cacheDir,
		cmd:              cmd,
		args:             args,
		installer:        readInstaller(),
	}

	t.storeSet = true
	if t.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx), store.Telemetry()); err != nil {
		_, _ = log.Debugf("telemetry: failed to set store: %v\n", err)
		t.storeSet = false
	}

	if !t.storeSet {
		o := []store.Option{store.UnauthenticatedPreset(config.Default()), store.WithContext(ctx), store.Telemetry(), store.Service(config.CloudService)}

		if t.unauthStore, err = store.New(o...); err != nil {
			_, _ = log.Debugf("telemetry: failed to set unauth store: %v\n", err)
		}
	}

	return t, nil
}

func (t *tracker) defaultCommandOptions() []EventOpt {
	return []EventOpt{
		withCommandPath(t.cmd),
		withHelpCommand(t.cmd, t.args),
		withFlags(t.cmd),
		withProfile(config.Default()),
		withVersion(),
		withOS(),
		withAuthMethod(config.Default()),
		withService(config.Default()),
		withProjectID(t.cmd, config.Default()),
		withOrgID(t.cmd, config.Default()),
		withTerminal(t.cmd),
		withInstaller(t.installer),
		withUserAgent(),
		withAnonymousID(),
		withCI(),
		withCLIUserType(),
		withSkipUpdateCheck(config.Default()),
	}
}

func (t *tracker) trackCommand(data TrackOptions, opt ...EventOpt) error {
	o := append(t.defaultCommandOptions(), withDuration(t.cmd))
	if data.Signal != "" {
		o = append(o, withSignal(data.Signal))
	}
	if data.Err != nil {
		o = append(o, withError(data.Err))
	}
	o = append(o, opt...)
	event := newEvent(o...)
	events, err := t.read()
	if err != nil {
		_, _ = log.Debugf("telemetry: failed to read cache: %v\n", err)
	}
	events = append(events, event)
	_, _ = log.Debugf("telemetry: events: %v\n", events)
	if !t.storeSet {
		err = t.unauthStore.SendUnauthEvents(events)
	} else {
		err = t.store.SendEvents(events)
	}
	if err != nil {
		return t.save(event)
	}
	return t.remove()
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
	data = append(data, '\n')
	_, err = file.Write(data)
	return err
}

// Read all events in the cache file.
func (t *tracker) read() ([]Event, error) {
	initialSize := 100
	events := make([]Event, 0, initialSize)
	filename := filepath.Join(t.cacheDir, cacheFilename)
	exists, err := afero.Exists(t.fs, filename)
	if err != nil {
		return events, err
	}
	if exists {
		file, err := t.fs.Open(filename)
		if err != nil {
			return events, err
		}
		defer file.Close()
		decoder := json.NewDecoder(file)
		for decoder.More() {
			var event Event
			if err := decoder.Decode(&event); err != nil {
				return events, err
			}
			events = append(events, event)
		}
	}
	return events, nil
}

// Removes the cache file.
func (t *tracker) remove() error {
	filename := filepath.Join(t.cacheDir, cacheFilename)
	exists, err := afero.Exists(t.fs, filename)
	if exists && err == nil {
		return t.fs.Remove(filename)
	}
	return err
}

func castBool(i any) bool {
	b, ok := i.(bool)
	if ok {
		return b
	}

	p, ok := i.(*bool)

	var ret bool
	if ok && i != nil {
		ret = *p
	}

	return ret
}

func castString(i any) string {
	s, ok := i.(string)
	if ok {
		return s
	}

	p, ok := i.(*string)

	var ret string
	if ok && i != nil {
		ret = *p
	}

	return ret
}

func (t *tracker) trackSurvey(p survey.Prompt, response any, e error) error {
	o := t.defaultCommandOptions()

	if e != nil {
		o = append(o, withError(e))
	}

	switch v := p.(type) {
	case *survey.Confirm:
		choice := "false"
		if castBool(response) {
			choice = "true"
		}
		o = append(o, withPrompt(v.Message, "confirm"), withDefault(castBool(response) == v.Default), withChoice(choice))
	case *survey.Input:
		o = append(o, withPrompt(v.Message, "input"), withDefault(castString(response) == v.Default), withEmpty(castString(response) == ""))
	case *survey.Password:
		o = append(o, withPrompt(v.Message, "password"), withEmpty(castString(response) == ""))
	case *survey.Select:
		o = append(o, withPrompt(v.Message, "select"), withDefault(castString(response) == v.Default), withEmpty(castString(response) == ""), withChoice(castString(response)))
	default:
		return errors.New("unknown survey prompt")
	}

	event := newEvent(o...)

	// all sent at once via TrackCommand
	// assuming there is always a TrackCommand after many TrackAsk
	return t.save(event)
}

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
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/spf13/afero"
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
	store            store.EventsSender
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
	}, nil
}

func (t *tracker) trackCommand(data TrackOptions) error {
	options := []eventOpt{withCommandPath(data.Cmd), withDuration(data.Cmd), withFlags(data.Cmd), withProfile(), withVersion(), withOS(), withAuthMethod(), withService(), withProjectID(data.Cmd), withOrgID(data.Cmd), withTerminal(), withInstaller(t.fs), withExtraProps(data.extraProps)}

	if data.Err != nil {
		options = append(options, withError(data.Err))
	}

	event := newEvent(options...)
	err := t.store.SendEvents(&[]Event{event})
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

func castBool(i interface{}) bool {
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

func castString(i interface{}) string {
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

func (t *tracker) trackSurvey(p survey.Prompt, response interface{}, e error) error {
	options := []eventOpt{}

	if e != nil {
		options = append(options, withError(e))
	}

	switch v := p.(type) {
	case *survey.Confirm:
		options = append(options, withPrompt(v.Message, "confirm"), withDefault(castBool(response) == v.Default))
	case *survey.Input:
		options = append(options, withPrompt(v.Message, "input"), withDefault(castString(response) == v.Default), withEmpty(castString(response) == ""))
	case *survey.Password:
		options = append(options, withPrompt(v.Message, "password"), withEmpty(castString(response) == ""))
	case *survey.Select:
		options = append(options, withPrompt(v.Message, "select"), withDefault(castString(response) == v.Default), withEmpty(castString(response) == ""))
	default:
		return errors.New("unknown survey prompt")
	}

	event := newEvent(options...)

	return t.save(event)
}

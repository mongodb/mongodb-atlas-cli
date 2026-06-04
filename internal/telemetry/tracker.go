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

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

const (
	cacheFilename           = "telemetry"
	backoffFilename         = "telemetry_backoff"
	dirPermissions          = 0700
	filePermissions         = 0600
	defaultMaxCacheFileSize = 500_000     // 500KB
	maxBatchSize            = 32          // backend rate limit per request
	backoffDuration         = time.Minute // minimum gap between send attempts
)

const (
	cmdEventType    = "cmd"
	promptEventType = "prompt"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=tracker_mock_test.go -package=telemetry -source=tracker.go

type EventsSender interface {
	SendEvents(body any) error
}

type UnauthEventsSender interface {
	SendUnauthEvents(body any) error
}

type tracker struct {
	fs               afero.Fs
	maxCacheFileSize int64
	cacheDir         string
	unauthStore      UnauthEventsSender
	store            EventsSender
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
		withEventType(cmdEventType),
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
		withOutput(t.cmd, config.Default()),
		withTerminal(t.cmd),
		withInstaller(t.installer),
		withUserAgent(),
		withAnonymousID(),
		withCI(),
		withAgent(),
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

	// Skip sending if we're in a backoff window; just cache the event.
	if t.isBackedOff() {
		_, _ = log.Debugf("telemetry: in backoff window, caching event for later\n")
		return t.save(event)
	}

	cachedEvents, err := t.read()
	if err != nil {
		_, _ = log.Debugf("telemetry: failed to read cache: %v\n", err)
	}

	cachedEvents = append(cachedEvents, event)

	// Cap batch at maxBatchSize; backend enforces a per-request rate limit.
	// Three-index slice prevents batch and remaining from sharing a backing array,
	// so serialization of batch cannot silently corrupt remaining.
	batch := cachedEvents
	var remaining []Event
	if len(cachedEvents) > maxBatchSize {
		batch = cachedEvents[:maxBatchSize:maxBatchSize]
		remaining = cachedEvents[maxBatchSize:]
	}

	_, _ = log.Debugf("telemetry: sending %d events (remaining cached: %d)\n", len(batch), len(remaining))

	var sendErr error
	if !t.storeSet {
		sendErr = t.unauthStore.SendUnauthEvents(batch)
	} else {
		sendErr = t.store.SendEvents(batch)
	}

	if sendErr != nil {
		// On any error: do not retry. Create backoff file so rapid subsequent
		// invocations skip the send entirely instead of hammering the backend.
		_, _ = log.Debugf("telemetry: send error (%v), backing off for %v\n", sendErr, backoffDuration)
		_ = t.saveBackoff() // best-effort: failure just means no backoff protection
		// Old cached events remain on disk (remove() not called).
		// Append only the new event so it's preserved for the next run.
		return t.save(event)
	}

	// Batch sent successfully: clear backoff, clear cache, persist remaining events.
	_ = t.removeBackoff() // best-effort: stale file only wastes one invocation
	if err := t.remove(); err != nil {
		return err
	}
	return t.saveAll(remaining)
}

func (t *tracker) openCacheFile() (afero.File, error) {
	if mkdirError := t.fs.MkdirAll(t.cacheDir, dirPermissions); mkdirError != nil {
		return nil, mkdirError
	}
	filename := filepath.Join(t.cacheDir, cacheFilename)
	exists, err := afero.Exists(t.fs, filename)
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

// Append multiple events to the cache file.
func (t *tracker) saveAll(events []Event) error {
	for _, e := range events {
		if err := t.save(e); err != nil {
			return err
		}
	}
	return nil
}

// saveBackoff creates the backoff sentinel file. Its mtime marks when the
// backoff window started; no content is written.
func (t *tracker) saveBackoff() error {
	if err := t.fs.MkdirAll(t.cacheDir, dirPermissions); err != nil {
		return err
	}
	filename := filepath.Join(t.cacheDir, backoffFilename)
	file, err := t.fs.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, filePermissions)
	if err != nil {
		return err
	}
	return file.Close()
}

// isBackedOff returns true when the backoff file exists and backoffDuration has
// not yet elapsed since its creation. Deletes the file once the window expires.
func (t *tracker) isBackedOff() bool {
	filename := filepath.Join(t.cacheDir, backoffFilename)
	info, err := t.fs.Stat(filename)
	if err != nil {
		return false
	}
	if time.Since(info.ModTime()) > backoffDuration {
		_ = t.fs.Remove(filename)
		return false
	}
	return true
}

// removeBackoff deletes the backoff file.
func (t *tracker) removeBackoff() error {
	filename := filepath.Join(t.cacheDir, backoffFilename)
	err := t.fs.Remove(filename)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return err
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
	o = append(o, withEventType(promptEventType))

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

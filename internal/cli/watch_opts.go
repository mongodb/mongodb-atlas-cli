// Copyright 2020 MongoDB Inc
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

package cli

import (
	"cmp"
	"errors"
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/watchers"
)

type WatchOpts struct {
	OutputOpts
	s              *spinner.Spinner
	EnableWatch    bool
	DefaultWait    time.Duration
	Timeout        uint
	IsRetryableErr func(err error) bool
}

const (
	defaultWait = 4 * time.Second
	speed       = 100 * time.Millisecond
)

type Watcher func() (any, bool, error)

// Watch allow to init the OutputOpts in a functional way.
func (opts *WatchOpts) Watch(f Watcher) (any, error) {
	if f == nil {
		return nil, errors.New("no watcher provided")
	}
	opts.start()
	for {
		value, done, err := opts.exponentialBackoff(f)
		if err != nil || done {
			opts.stop()
			return value, err
		}
		if !opts.IsTerminal() {
			if _, err = fmt.Fprint(opts.ConfigWriter(), "."); err != nil {
				return nil, err
			}
		}
		time.Sleep(opts.GetDefaultWait())
	}
}

var backoffCoefficients = []float32{0.5, 1, 2}

func (opts *WatchOpts) exponentialBackoff(f Watcher) (any, bool, error) {
	if opts.IsRetryableErr == nil {
		return f()
	}

	for _, coefficient := range backoffCoefficients {
		if value, done, err := f(); err == nil || !opts.IsRetryableErr(err) {
			return value, done, err
		}
		time.Sleep(time.Duration(coefficient) * opts.GetDefaultWait())
	}
	// Should only happen after trying three times (>14 seconds)
	return f()
}

func (opts *WatchOpts) WatchWatcher(w *watchers.Watcher) error {
	if opts.EnableWatch {
		opts.start()
		err := w.Watch()
		opts.stop()
		return err
	}

	return nil
}

func (opts *WatchOpts) start() {
	if opts.IsTerminal() {
		opts.s = spinner.New(spinner.CharSets[9], speed)
		opts.s.Start()
	}
}

func (opts *WatchOpts) stop() {
	if opts.IsTerminal() {
		opts.s.Stop()
	}
}

func (opts *WatchOpts) GetDefaultWait() time.Duration {
	return cmp.Or(opts.DefaultWait, defaultWait)
}

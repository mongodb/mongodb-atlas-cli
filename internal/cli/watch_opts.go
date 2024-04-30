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
	"errors"
	"fmt"
	"time"

	"github.com/andreaangiolillo/mongocli-test/internal/watchers"
	"github.com/briandowns/spinner"
)

type WatchOpts struct {
	OutputOpts
	s              *spinner.Spinner
	EnableWatch    bool
	Timeout        uint
	IsRetryableErr func(err error) bool
}

const (
	defaultWait = 4 * time.Second
	speed       = 100 * time.Millisecond
)

type Watcher func() (bool, error)

// Watch allow to init the OutputOpts in a functional way.
func (opts *WatchOpts) Watch(f Watcher) error {
	if f == nil {
		return errors.New("no watcher provided")
	}
	opts.start()
	for {
		done, err := opts.exponentialBackoff(f)
		if err != nil || done {
			opts.stop()
			return err
		}
		if !opts.IsTerminal() {
			if _, err = fmt.Fprint(opts.ConfigWriter(), "."); err != nil {
				return err
			}
		}
		time.Sleep(defaultWait)
	}
}

var backoffTimes = []time.Duration{2 * time.Second, 4 * time.Second, 8 * time.Second}

func (opts *WatchOpts) exponentialBackoff(f Watcher) (bool, error) {
	if opts.IsRetryableErr == nil {
		return f()
	}

	for _, backoffTime := range backoffTimes {
		if done, err := f(); err == nil || !opts.IsRetryableErr(err) {
			return done, err
		}
		time.Sleep(backoffTime)
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

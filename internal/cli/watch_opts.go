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

	"github.com/briandowns/spinner"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type WatchOpts struct {
	OutputOpts
	s *spinner.Spinner
	n int
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
	opts.n = 0
	for {
		done, err := f()
		if err != nil {
			if opts.exponentialBackoff(err) {
				continue
			}
			opts.stop()
			return err
		}
		if done {
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

func (opts *WatchOpts) exponentialBackoff(err error) bool {
	backoffTimes := []time.Duration{1 * time.Second, 2 * time.Second, 4 * time.Second}
	var atlasErr *atlas.ErrorResponse
	errorCode := "CLUSTER_NOT_FOUND"
	if opts.n < len(backoffTimes) && errors.As(err, &atlasErr) && atlasErr.ErrorCode == errorCode {
		time.Sleep(backoffTimes[opts.n])
		opts.n++
		return true
	}
	return false
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

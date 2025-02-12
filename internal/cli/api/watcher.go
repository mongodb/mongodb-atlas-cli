// Copyright 2025 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"context"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/api"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/terminal"
)

const (
	spinnerSpeed  = 100 * time.Millisecond
	watchInterval = time.Second
)

// Watcher provides the functionality to wait until a certain operation is in the expected state
//
// The main goal of this watcher is to:
// - Call the API layer watcher
// - Handle sleeps and cancellations
// - Handle an optional spinner to Stdout if the terminal is not piped.
type Watcher struct {
	apiWatcher api.Watcher
	spinner    *spinner.Spinner
}

// Construct a new CLI layer watcher
//
// See api.NewWatcher for docs on arguments.
func NewWatcher(executor api.CommandExecutor, requestParams map[string][]string, responseBody []byte, props api.WatcherProperties) (*Watcher, error) {
	apiWatcher, err := api.NewWatcher(executor, requestParams, responseBody, props)
	if err != nil {
		return nil, err
	}

	// We only want to spin when Stderr is not piped
	// We're showing the spinner on Stderr because then the output of Stdout is still usable in scripts
	var s *spinner.Spinner
	if terminal.IsTerminal(os.Stderr) {
		s = spinner.New(spinner.CharSets[9], spinnerSpeed, spinner.WithWriter(os.Stderr))
	}

	return &Watcher{
		apiWatcher: *apiWatcher,
		spinner:    s,
	}, nil
}

func (w *Watcher) Wait(ctx context.Context) error {
	// Take care of the spinner if we're in terminal mode
	w.startSpinner()
	defer w.stopSpinner()

	// Keep calling WatchOne until one of the following events happens
	// - watcher completes successfully
	// - watcher returns an error
	// - the context is cancelled (example user presses crtl+c)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			done, err := w.apiWatcher.WatchOne(ctx)
			if err != nil {
				return err
			}

			if done {
				return nil
			}

			time.Sleep(watchInterval)
		}
	}
}

func (w *Watcher) startSpinner() {
	if w.spinner != nil {
		w.spinner.Start()
	}
}

func (w *Watcher) stopSpinner() {
	if w.spinner != nil {
		w.spinner.Stop()
	}
}

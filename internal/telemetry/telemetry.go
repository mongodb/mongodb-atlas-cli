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
	"strings"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/spf13/cobra"
)

var contextKey = telemetryContextKey{}

type telemetryContextKey struct{}

type telemetryContextValue struct {
	startTime time.Time
}

func NewContext() context.Context {
	return context.WithValue(context.Background(), contextKey, telemetryContextValue{
		startTime: time.Now(),
	})
}

type TrackOptions struct {
	Cmd        *cobra.Command
	Err        error
	extraProps map[string]interface{}
}

func TrackCommand(opt TrackOptions, args ...string) {
	if !config.TelemetryEnabled() {
		return
	}
	t, err := newTracker(opt.Cmd.Context())
	if err != nil {
		logError(err)
		return
	}

	checkHelp(&opt, args...)

	if err = t.track(opt); err != nil {
		logError(err)
	}
}

func logError(err error) {
	// No-op function until logging is implemented (CLOUDP-110988)
	_ = err
}

func checkHelp(opt *TrackOptions, args ...string) {
	if opt.Cmd.Name() != "help" {
		return
	}
	cmd, _, err := opt.Cmd.Root().Find(args)
	if err != nil {
		return
	}
	opt.extraProps = map[string]interface{}{
		"help_command": strings.ReplaceAll(cmd.CommandPath(), " ", "-"),
	}
}

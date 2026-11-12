// Copyright 2026 MongoDB Inc
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

package plugin

import (
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"github.com/spf13/cobra"
)

// TelemetryForForwardedArgs reports the real command path and flags for plugin
// commands that disable flag parsing and forward their arguments verbatim to the
// plugin binary (first-class and regular plugins). cobra cannot resolve the
// subcommand or flags itself in that case, so they are recovered from the raw
// arguments.
func TelemetryForForwardedArgs(cmd *cobra.Command, args []string) telemetry.EventOpt {
	commandPath, flags := forwardedCommandAndFlags(cmd, args)
	return func(event telemetry.Event) {
		event.Properties["command"] = strings.ReplaceAll(commandPath, " ", "-")
		if len(flags) > 0 {
			event.Properties["flags"] = flags
		}
	}
}

// forwardedCommandAndFlags splits the args a plugin command forwards verbatim to
// the plugin binary into the subcommand path and the set flag names. It assumes
// subcommands precede flags (the CLI convention); a subcommand placed after a
// flag is not recovered. Flag values are not distinguished from subcommands, so
// a positional argument is reported as part of the subcommand path.
func forwardedCommandAndFlags(cmd *cobra.Command, args []string) (commandPath string, flags []string) {
	commandPath = cmd.CommandPath()

	collectingFlags := false
	for _, arg := range args {
		switch {
		case arg == "--":
			// The terminator ends both subcommands and flags.
			return commandPath, flags
		case strings.HasPrefix(arg, "-"):
			collectingFlags = true
			if name := strings.TrimLeft(arg, "-"); name != "" {
				if i := strings.IndexByte(name, '='); i >= 0 {
					name = name[:i]
				}
				flags = append(flags, name)
			}
		case !collectingFlags:
			commandPath += " " + arg
		}
	}

	return commandPath, flags
}

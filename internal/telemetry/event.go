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
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/homebrew"
	"github.com/mongodb/mongocli/internal/version"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Event struct {
	Timestamp  time.Time              `json:"timestamp"`
	Source     string                 `json:"source"`
	Name       string                 `json:"name"`
	Properties map[string]interface{} `json:"properties"`
}

type eventOpt func(Event)

func withProfile() eventOpt { // either "default" or base64 hash
	return func(event Event) {
		if config.Name() == config.DefaultProfile {
			event.Properties["profile"] = config.DefaultProfile
			return
		}

		h := sha256.Sum256([]byte(config.Name()))
		event.Properties["profile"] = base64.StdEncoding.EncodeToString(h[:])
	}
}

func withCommandPath(cmd *cobra.Command) eventOpt {
	return func(event Event) {
		cmdPath := cmd.CommandPath()
		event.Properties["command"] = strings.ReplaceAll(cmdPath, " ", "-")
	}
}

func withDuration(cmd *cobra.Command) eventOpt {
	return func(event Event) {
		ctxValue, found := cmd.Context().Value(contextKey).(telemetryContextValue)
		if !found {
			logError(errors.New("telemetry context not found"))
			return
		}

		event.Properties["duration"] = event.Timestamp.Sub(ctxValue.startTime).Milliseconds()
	}
}

func withFlags(cmd *cobra.Command) eventOpt {
	return func(event Event) {
		setFlags := make([]string, 0, cmd.Flags().NFlag())
		cmd.Flags().Visit(func(f *pflag.Flag) {
			setFlags = append(setFlags, f.Name)
		})

		if len(setFlags) > 0 {
			event.Properties["flags"] = setFlags
		}
	}
}

func withVersion() eventOpt {
	return func(event Event) {
		event.Properties["version"] = version.Version
		event.Properties["git-commit"] = version.GitCommit
	}
}

func withOS() eventOpt {
	return func(event Event) {
		event.Properties["os"] = runtime.GOOS
		event.Properties["arch"] = runtime.GOARCH
	}
}

func withAuthMethod() eventOpt {
	return func(event Event) {
		if config.PublicAPIKey() != "" && config.PrivateAPIKey() != "" {
			event.Properties["auth_method"] = "api_key"
			return
		}

		event.Properties["auth_method"] = "oauth"
	}
}

func withService() eventOpt {
	return func(event Event) {
		event.Properties["service"] = config.Service()
		if config.OpsManagerURL() != "" {
			event.Properties["ops_manager_url"] = config.OpsManagerURL()
		}
	}
}

func withProjectID(cmd *cobra.Command) eventOpt {
	return func(event Event) {
		fromFlag, _ := cmd.Flags().GetString(flag.ProjectID)

		if fromFlag != "" {
			event.Properties["project_id"] = fromFlag
			return
		}

		if config.ProjectID() != "" {
			event.Properties["project_id"] = config.ProjectID()
		}
	}
}

func withOrgID(cmd *cobra.Command) eventOpt {
	return func(event Event) {
		fromFlag, _ := cmd.Flags().GetString(flag.OrgID)

		if fromFlag != "" {
			event.Properties["org_id"] = fromFlag
			return
		}

		if config.OrgID() != "" {
			event.Properties["org_id"] = config.OrgID()
		}
	}
}

func withTerminal() eventOpt {
	return func(event Event) {
		if cli.IsTerminal(os.Stdout) {
			event.Properties["terminal"] = "tty"
			return
		}

		if cli.IsCygwinTerminal(os.Stdout) {
			event.Properties["terminal"] = "cygwin"
		}
	}
}

func withInstaller(fs afero.Fs) eventOpt {
	return func(event Event) {
		c, err := homebrew.NewChecker(fs)
		if err != nil {
			logError(err)
			return
		}
		if c.IsHomebrew() {
			event.Properties["installer"] = "brew"
		}
	}
}

func withError(err error) eventOpt {
	return func(event Event) {
		event.Properties["result"] = "ERROR"

		errorMessage := strings.Split(err.Error(), "\n")[0] //only first line

		event.Properties["error"] = errorMessage
	}
}

func newEvent(opts ...eventOpt) Event {
	var event = Event{
		Timestamp: time.Now(),
		Source:    config.ToolName,
		Name:      config.ToolName + "-event",
		Properties: map[string]interface{}{
			"result": "SUCCESS",
		},
	}

	for _, fn := range opts {
		fn(event)
	}

	return event
}

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
	"io"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/denisbrodbeck/machineid"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/terminal"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Event struct {
	Timestamp  time.Time      `json:"timestamp"`
	Source     string         `json:"source"`
	Properties map[string]any `json:"properties"`
}

type EventOpt func(Event)

func withHelpCommand(cmd *cobra.Command, args []string) EventOpt {
	return func(event Event) {
		if cmd.Name() != "help" {
			return
		}

		helpCmd, _, err := cmd.Root().Find(args)
		if err != nil {
			_, _ = log.Debugf("telemetry: failed to find help command: %v\n", err)
			return
		}

		event.Properties["help_command"] = strings.ReplaceAll(helpCmd.CommandPath(), " ", "-")
	}
}

type ConfigNameGetter interface {
	Name() string
}

func withProfile(c ConfigNameGetter) EventOpt { // either "default" or base64 hash
	return func(event Event) {
		if c.Name() == config.DefaultProfile {
			event.Properties["profile"] = config.DefaultProfile
			return
		}

		h := sha256.Sum256([]byte(config.Name()))
		event.Properties["profile"] = base64.StdEncoding.EncodeToString(h[:])
	}
}

func withPrompt(p, k string) EventOpt {
	return func(event Event) {
		event.Properties["prompt"] = sanitizePrompt(p)
		event.Properties["prompt_type"] = k
	}
}

func withChoice(c string) EventOpt {
	return func(event Event) {
		event.Properties["choice"] = sanitizeSelectOption(c)
	}
}

func sanitizeSelectOption(v string) string {
	parenthesesRegex := regexp.MustCompile(`^.*\(([^()]*)\)$`)

	return parenthesesRegex.ReplaceAllString(v, "$1")
}

func sanitizePrompt(q string) string {
	bracketsRegex := regexp.MustCompile(`\[[^]\[]*]`)

	return bracketsRegex.ReplaceAllString(q, "[]")
}

func withDefault(d bool) EventOpt {
	return func(event Event) {
		event.Properties["default"] = d
	}
}

func withEmpty(e bool) EventOpt {
	return func(event Event) {
		event.Properties["empty"] = e
	}
}

type CmdName interface {
	CommandPath() string
	CalledAs() string
}

func withCommandPath(cmd CmdName) EventOpt {
	return func(event Event) {
		cmdPath := cmd.CommandPath()
		event.Properties["command"] = strings.ReplaceAll(cmdPath, " ", "-")
		if cmd.CalledAs() != "" {
			event.Properties["alias"] = cmd.CalledAs()
		}
	}
}

func withDuration(cmd *cobra.Command) EventOpt {
	return func(event Event) {
		if cmd.Context() == nil {
			_, _ = log.Debugln("telemetry: context not found")
			return
		}

		ctxValue, found := cmd.Context().Value(contextKey).(telemetryContextValue)
		if !found {
			_, _ = log.Debugln("telemetry: context not found")
			return
		}

		event.Properties["duration"] = event.Timestamp.Sub(ctxValue.startTime).Milliseconds()
	}
}

func withCI() EventOpt {
	return func(event Event) {
		_, ok := os.LookupEnv("CI")
		event.Properties["ci"] = ok
	}
}

func withAnonymousID() EventOpt {
	return func(event Event) {
		id, err := machineid.ProtectedID(config.AtlasCLI)
		if err != nil {
			event.Properties["device_id_err"] = err.Error()
			_, _ = log.Debugf("error generating machine id: %v\n", err)
			return
		}
		event.Properties["device_id"] = id
	}
}

func WithDeploymentType(t string) EventOpt {
	return func(event Event) {
		event.Properties["deployment_type"] = t
	}
}

type CmdFlags interface {
	Flags() *pflag.FlagSet
}

func withFlags(cmd CmdFlags) EventOpt {
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

func withVersion() EventOpt {
	return func(event Event) {
		event.Properties["version"] = version.Version
		event.Properties["git_commit"] = version.GitCommit
	}
}

func withOS() EventOpt {
	return func(event Event) {
		event.Properties["os"] = runtime.GOOS
		event.Properties["arch"] = runtime.GOARCH
	}
}

type Authenticator interface {
	PublicAPIKey() string
	PrivateAPIKey() string
	AccessToken() string
}

func withAuthMethod(c Authenticator) EventOpt {
	return func(event Event) {
		if c.PublicAPIKey() != "" && c.PrivateAPIKey() != "" {
			event.Properties["auth_method"] = "api_key"
			return
		} else if c.AccessToken() != "" {
			event.Properties["auth_method"] = "oauth"
		}
	}
}

type ServiceGetter interface {
	Service() string
	OpsManagerURL() string
}

func withService(c ServiceGetter) EventOpt {
	return func(event Event) {
		event.Properties["service"] = c.Service()
		if c.OpsManagerURL() != "" {
			event.Properties["ops_manager_url"] = c.OpsManagerURL()
		}
	}
}

type SkipUpdateGetter interface {
	SkipUpdateCheck() bool
}

func withSkipUpdateCheck(c SkipUpdateGetter) EventOpt {
	return func(event Event) {
		event.Properties["skip_update_check"] = c.SkipUpdateCheck()
	}
}

type ProjectIDGetter interface {
	ProjectID() string
}

func withProjectID(cmd CmdFlags, c ProjectIDGetter) EventOpt {
	return func(event Event) {
		fromFlag, _ := cmd.Flags().GetString(flag.ProjectID)

		if fromFlag != "" {
			event.Properties["project_id"] = fromFlag
			return
		}

		if c.ProjectID() != "" {
			event.Properties["project_id"] = c.ProjectID()
		}
	}
}

func withCLIUserType() EventOpt {
	return func(event Event) {
		event.Properties["cli_user_type"] = config.CLIUserType
	}
}

type OrgIDGetter interface {
	OrgID() string
}

func withOrgID(cmd CmdFlags, c OrgIDGetter) EventOpt {
	return func(event Event) {
		fromFlag, _ := cmd.Flags().GetString(flag.OrgID)

		if fromFlag != "" {
			event.Properties["org_id"] = fromFlag
			return
		}

		if c.OrgID() != "" {
			event.Properties["org_id"] = c.OrgID()
		}
	}
}

type Printer interface {
	OutOrStdout() io.Writer
}

func withTerminal(cmd Printer) EventOpt {
	return func(event Event) {
		if terminal.IsCygwinTerminal(cmd.OutOrStdout()) {
			event.Properties["terminal"] = "cygwin"
		}

		if terminal.IsTerminal(cmd.OutOrStdout()) {
			event.Properties["terminal"] = "tty"
			return
		}
	}
}

func withInstaller(installer *string) EventOpt {
	return func(event Event) {
		if installer != nil {
			event.Properties["installer"] = *installer
		}
	}
}

func withError(err error) EventOpt {
	return func(event Event) {
		event.Properties["result"] = "ERROR"

		errorMessage := strings.Split(err.Error(), "\n")[0] // only first line

		event.Properties["error"] = errorMessage
	}
}

func withSignal(s string) EventOpt {
	return func(event Event) {
		event.Properties["signal"] = s
	}
}

func withUserAgent() EventOpt {
	return func(event Event) {
		event.Properties["UserAgent"] = config.UserAgent
		event.Properties["HostName"] = config.HostName
	}
}

func newEvent(opts ...EventOpt) Event {
	var event = Event{
		Timestamp: time.Now(),
		Source:    config.AtlasCLI,
		Properties: map[string]any{
			"result": "SUCCESS",
		},
	}

	for _, fn := range opts {
		fn(event)
	}

	return event
}

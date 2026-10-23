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

package root

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/accesslists"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/accesslogs"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/alerts"
	apiCmd "github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/api"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/auditing"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/auth"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/backup"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/cloudproviders"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/clusters"
	cliconfig "github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/customdbroles"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/customdns"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/datafederation"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/datalake"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/dbusers"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/events"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/federatedauthentication"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/integrations"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/livemigrations"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/logs"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/maintenance"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/metrics"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/networking"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/organizations"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/performanceadvisor"
	pluginCmd "github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/plugin"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/privateendpoints"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/processes"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/projects"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/security"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/setup"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/skills"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/streams"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/teams"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/users"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/homebrew"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/latestrelease"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/prerun"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/sighandle"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/terminal"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20250312027/detectaiagent"
)

const atlas = "atlas"

const (
	agentSkillsWarningCacheFile       = "mongodb/atlas-cli/agent-skills-warning"
	agentSkillsWarningDirectoryMode   = os.FileMode(0o700)
	agentSkillsWarningFilePermissions = os.FileMode(0o600)
)

type Notifier struct {
	currentVersion string
	finder         latestrelease.VersionFinder
	filesystem     afero.Fs
	writer         io.Writer
}

func handleSignal() {
	sighandle.Notify(func(sig os.Signal) {
		telemetry.FinishTrackingCommand(telemetry.TrackOptions{
			Err:    errors.New(sig.String()),
			Signal: sig.String(),
		})
		os.Exit(1)
	}, os.Interrupt, syscall.SIGTERM)
}

func isAgentSkillsCommand(cmd *cobra.Command) bool {
	return cmd.CommandPath() == atlas+" skills" || strings.HasPrefix(cmd.CommandPath(), atlas+" skills ")
}

func notifyAgentSkillsIfApplicable(writer io.Writer, filesystem afero.Fs, homeDir, claudeConfigDir, codexHomeDir, cacheDir string, agentDetected bool) error {
	if !agentDetected {
		return nil
	}
	installed, err := skills.HasInstalledMongoDBSkills(filesystem, homeDir, claudeConfigDir, codexHomeDir)
	if err != nil {
		return err
	}
	if installed {
		return nil
	}
	warningPath := filepath.Join(cacheDir, filepath.FromSlash(agentSkillsWarningCacheFile))
	if err := filesystem.MkdirAll(filepath.Dir(warningPath), agentSkillsWarningDirectoryMode); err != nil {
		return err
	}
	if _, err := filesystem.Stat(warningPath); err == nil {
		return nil
	} else if !errors.Is(err, os.ErrNotExist) {
		return err
	}
	_, err = fmt.Fprintln(writer, `Warning: MongoDB agent skills are not installed for Claude Code or Codex.
To install all standalone skills for Claude Code, run:
  atlas skills install --all --agent claude
To install all standalone skills for Codex, run:
  atlas skills install --all --agent codex
To install the Atlas plugin for Claude Code, run:
  atlas skills install --plugin mongodb-atlas --agent claude
To install the Atlas plugin for Codex, run:
  atlas skills install --plugin mongodb-atlas --agent codex`)
	if err != nil {
		return err
	}
	return afero.WriteFile(filesystem, warningPath, []byte("shown"), agentSkillsWarningFilePermissions)
}

func isShellCompletionCommand(cmd *cobra.Command) bool {
	path := cmd.CommandPath()
	return strings.HasPrefix(path, atlas+" completion") || strings.HasPrefix(path, atlas+" __complete")
}

func maybeWarnAboutAgentSkills(cmd *cobra.Command) {
	if isShellCompletionCommand(cmd) {
		return
	}
	if ctx := cmd.Context(); ctx != nil && ctx.Err() != nil {
		return
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return
	}
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return
	}
	_, agentDetected := detectaiagent.Detect()
	if err := notifyAgentSkillsIfApplicable(cmd.ErrOrStderr(), afero.NewOsFs(), homeDir, os.Getenv("CLAUDE_CONFIG_DIR"), os.Getenv("CODEX_HOME"), cacheDir, agentDetected); err != nil {
		_, _ = log.Debugf("agent skills discovery warning failed: %v\n", err)
	}
}

// Builder constructs the Atlas command tree and optionally registers installed Atlas CLI plugin commands.
func Builder(registerPluginCommands bool) *cobra.Command {
	var (
		profile    string
		debugLevel bool
	)
	opts := &cli.RefresherOpts{}
	rootCmd := &cobra.Command{
		Version: version.Version,
		Use:     atlas,
		Short:   "CLI tool to manage MongoDB Atlas.",
		Long: `The Atlas CLI is a command line interface built specifically for MongoDB Atlas. You can manage your Atlas database deployments and Atlas Search from the terminal with short, intuitive commands.
		
Use the --help flag with any command for more info on that command.`,
		Example: `  # Display the help menu for the config command:
  atlas config --help
`,
		SilenceUsage:  true,
		SilenceErrors: true,
		Annotations: map[string]string{
			"toc": "true",
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			log.SetWriter(cmd.ErrOrStderr())
			if debugLevel {
				log.SetLevel(log.DebugLevel)
			}
			if isAgentSkillsCommand(cmd) {
				return nil
			}

			if err := config.InitProfile(profile); err != nil {
				return err
			}

			telemetry.StartTrackingCommand(cmd, args)
			handleSignal()

			if shouldSetService(cmd) {
				config.SetService(config.CloudService)
			}

			return prerun.ExecuteE(opts.InitFlow(config.Default()))
		},
		// PersistentPostRun only runs if the command is successful
		PersistentPostRun: func(cmd *cobra.Command, _ []string) {
			maybeWarnAboutAgentSkills(cmd)
			if isAgentSkillsCommand(cmd) {
				return
			}
			// we don't run the release alert feature on the completion command
			if strings.HasPrefix(cmd.CommandPath(), fmt.Sprintf("%s %s", atlas, "completion")) {
				return
			}

			w := cmd.ErrOrStderr()
			fs := afero.NewOsFs()
			f, _ := latestrelease.NewVersionFinder(fs, version.NewReleaseVersionDescriber())

			notifier := &Notifier{
				currentVersion: latestrelease.VersionFromTag(version.Version),
				finder:         f,
				filesystem:     fs,
				writer:         w,
			}

			if check, isHb := notifier.shouldCheck(); check && cmd.Context().Err() == nil {
				_ = notifier.notifyIfApplicable(isHb)
			}
			telemetry.FinishTrackingCommand(telemetry.TrackOptions{})
		},
	}
	rootCmd.SetVersionTemplate(formattedVersion())
	rootCmd.SetHelpTemplate(cli.HelpTemplate)

	// hidden shortcuts
	loginCmd := auth.LoginBuilder()
	loginCmd.Hidden = true
	logoutCmd := auth.LogoutBuilder()
	logoutCmd.Hidden = true
	whoCmd := auth.WhoAmIBuilder()
	whoCmd.Hidden = true
	registerCmd := auth.RegisterBuilder()
	registerCmd.Hidden = true

	rootCmd.AddCommand(
		cliconfig.Builder(),
		auth.Builder(),
		setup.Builder(),
		projects.Builder(),
		organizations.Builder(),
		users.Builder(),
		teams.Builder(),
		clusters.Builder(),
		dbusers.Builder(),
		customdbroles.Builder(),
		accesslists.Builder(),
		datalake.Builder(),
		alerts.Builder(),
		backup.Builder(),
		events.Builder(),
		metrics.Builder(),
		performanceadvisor.Builder(),
		logs.Builder(),
		processes.Builder(),
		privateendpoints.Builder(),
		networking.Builder(),
		security.Builder(),
		skills.Builder(),
		integrations.Builder(),
		maintenance.Builder(),
		customdns.Builder(),
		cloudproviders.Builder(),
		streams.Builder(),
		livemigrations.Builder(),
		accesslogs.Builder(),
		loginCmd,
		logoutCmd,
		whoCmd,
		registerCmd,
		datafederation.Builder(),
		auditing.Builder(),
		deployments.Builder(),
		federatedauthentication.Builder(),
		apiCmd.Builder(),
	)

	if registerPluginCommands {
		pluginCmd.RegisterCommands(rootCmd)
	}

	rootCmd.PersistentFlags().StringVarP(&profile, flag.Profile, flag.ProfileShort, "", usage.ProfileAtlasCLI)
	rootCmd.PersistentFlags().BoolVarP(&debugLevel, flag.Debug, flag.DebugShort, false, usage.Debug)
	_ = rootCmd.PersistentFlags().MarkHidden(flag.Debug)

	_ = rootCmd.RegisterFlagCompletionFunc(flag.Profile, func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return config.List(), cobra.ShellCompDirectiveDefault
	})
	return rootCmd
}

const verTemplate = `atlascli version: %s
git version: %s
Go version: %s
   os: %s
   arch: %s
   compiler: %s
`

func shouldSetService(cmd *cobra.Command) bool {
	if config.Service() != "" {
		return false
	}

	if strings.HasPrefix(cmd.CommandPath(), fmt.Sprintf("%s %s", atlas, "config")) { // user wants to set credentials
		return false
	}

	if strings.HasPrefix(cmd.CommandPath(), fmt.Sprintf("%s %s", atlas, "completion")) {
		return false
	}

	// Setting the service would create the selected profile in memory and defeat logout's profile existence check.
	if strings.HasPrefix(cmd.CommandPath(), fmt.Sprintf("%s %s %s", atlas, "auth", "logout")) ||
		strings.HasPrefix(cmd.CommandPath(), fmt.Sprintf("%s %s", atlas, "logout")) {
		return false
	}

	return true
}

func formattedVersion() string {
	return fmt.Sprintf(verTemplate,
		version.Version,
		version.GitCommit,
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
		runtime.Compiler)
}

func (n *Notifier) shouldCheck() (shouldCheck, isHb bool) {
	shouldCheck = !config.SkipUpdateCheck() && terminal.IsTerminal(n.writer)
	isHb = false

	if !shouldCheck {
		return shouldCheck, isHb
	}

	c, _ := homebrew.NewChecker(n.filesystem)
	isHb = c.IsHomebrew()

	return shouldCheck, isHb
}

func (n *Notifier) notifyIfApplicable(isHb bool) error {
	release, err := n.finder.Find()
	if err != nil || release == nil {
		return err
	}

	// homebrew is an external dependency we give them 24h to have the cli available there
	if isHb && !isAtLeast24HoursPast(release.PublishedAt) {
		return nil
	}

	var upgradeInstructions string
	if isHb {
		upgradeInstructions = `To upgrade, run "brew update && brew upgrade mongodb-atlas-cli"`
	} else {
		upgradeInstructions = "To upgrade, see: https://dochub.mongodb.org/core/install-atlas-cli"
	}
	_, err = fmt.Fprintf(n.writer, `
A new version of atlascli is available %q!
%s

To disable this alert, run "atlas config set skip_update_check true"
`, release.Version, upgradeInstructions)
	return err
}

func isAtLeast24HoursPast(t time.Time) bool {
	return !t.IsZero() && time.Since(t) >= time.Hour*24
}

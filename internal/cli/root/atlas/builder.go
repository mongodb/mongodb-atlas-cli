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

package atlas

import (
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/andreangiolillo/mongocli-test/internal/cli"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/accesslists"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/accesslogs"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/alerts"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/auditing"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/backup"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/cloudproviders"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/clusters"
	atlasconfig "github.com/andreangiolillo/mongocli-test/internal/cli/atlas/config"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/customdbroles"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/customdns"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/datafederation"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/datalake"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/datalakepipelines"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/dbusers"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/deployments"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/events"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/integrations"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/kubernetes"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/livemigrations"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/logs"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/maintenance"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/metrics"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/networking"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/organizations"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/performanceadvisor"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/privateendpoints"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/processes"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/projects"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/security"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/serverless"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/setup"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/streams"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/teams"
	"github.com/andreangiolillo/mongocli-test/internal/cli/atlas/users"
	"github.com/andreangiolillo/mongocli-test/internal/cli/auth"
	"github.com/andreangiolillo/mongocli-test/internal/config"
	"github.com/andreangiolillo/mongocli-test/internal/flag"
	"github.com/andreangiolillo/mongocli-test/internal/homebrew"
	"github.com/andreangiolillo/mongocli-test/internal/latestrelease"
	"github.com/andreangiolillo/mongocli-test/internal/log"
	"github.com/andreangiolillo/mongocli-test/internal/prerun"
	"github.com/andreangiolillo/mongocli-test/internal/sighandle"
	"github.com/andreangiolillo/mongocli-test/internal/telemetry"
	"github.com/andreangiolillo/mongocli-test/internal/terminal"
	"github.com/andreangiolillo/mongocli-test/internal/usage"
	"github.com/andreangiolillo/mongocli-test/internal/validate"
	"github.com/andreangiolillo/mongocli-test/internal/version"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

const atlas = "atlas"

type Notifier struct {
	currentVersion string
	finder         latestrelease.VersionFinder
	filesystem     afero.Fs
	writer         io.Writer
}

type AuthRequirements int64

const (
	// command does not require authentication.
	NoAuth AuthRequirements = 0
	// command requires authentication.
	RequiredAuth AuthRequirements = 1
	// command can work with or without authentication,
	// and if access token token is found, try to refresh it.
	OptionalAuth AuthRequirements = 2
)

func handleSignal() {
	sighandle.Notify(func(sig os.Signal) {
		telemetry.FinishTrackingCommand(telemetry.TrackOptions{
			Err:    errors.New(sig.String()),
			Signal: sig.String(),
		})
		os.Exit(1)
	}, os.Interrupt, syscall.SIGTERM)
}

// Builder conditionally adds children commands as needed.
func Builder() *cobra.Command {
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
		SilenceUsage: true,
		Annotations: map[string]string{
			"toc": "true",
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			log.SetWriter(cmd.ErrOrStderr())
			if debugLevel {
				log.SetLevel(log.DebugLevel)
			}

			if err := cli.InitProfile(profile); err != nil {
				return err
			}

			telemetry.StartTrackingCommand(cmd, args)

			handleSignal()

			if shouldSetService(cmd) {
				config.SetService(config.CloudService)
			}
			if authReq := shouldCheckCredentials(cmd); authReq != NoAuth {
				if err := prerun.ExecuteE(
					opts.InitFlow(config.Default()),
					func() error {
						if err := opts.RefreshAccessToken(cmd.Context()); err != nil {
							if authReq == RequiredAuth {
								return err
							}
							_, _ = log.Warningf("Could not refresh access token: %s\n", err.Error())
						}
						return nil
					},
				); err != nil {
					return err
				}

				if authReq == RequiredAuth {
					return validate.Credentials()
				}
			}

			return nil
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			// we don't run the release alert feature on the completion command
			if strings.HasPrefix(cmd.CommandPath(), fmt.Sprintf("%s %s", atlas, "completion")) {
				return
			}

			w := cmd.ErrOrStderr()
			fs := afero.NewOsFs()
			f, _ := latestrelease.NewVersionFinder(fs, version.NewReleaseVersionDescriber())

			notifier := &Notifier{
				currentVersion: latestrelease.VersionFromTag(version.Version, config.ToolName),
				finder:         f,
				filesystem:     fs,
				writer:         w,
			}

			if check, isHb := notifier.shouldCheck(); check {
				_ = notifier.notifyIfApplicable(isHb)
			}
			telemetry.FinishTrackingCommand(telemetry.TrackOptions{})
		},
	}
	rootCmd.SetVersionTemplate(formattedVersion())

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
		atlasconfig.Builder(),
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
		datalakepipelines.Builder(),
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
		integrations.Builder(),
		maintenance.Builder(),
		customdns.Builder(),
		cloudproviders.Builder(),
		serverless.Builder(),
		streams.Builder(),
		livemigrations.Builder(),
		accesslogs.Builder(),
		loginCmd,
		logoutCmd,
		whoCmd,
		registerCmd,
		kubernetes.Builder(),
		datafederation.Builder(),
		auditing.Builder(),
		deployments.Builder(),
	)

	rootCmd.PersistentFlags().StringVarP(&profile, flag.Profile, flag.ProfileShort, "", usage.ProfileAtlasCLI)
	rootCmd.PersistentFlags().BoolVarP(&debugLevel, flag.Debug, flag.DebugShort, false, usage.Debug)
	_ = rootCmd.PersistentFlags().MarkHidden(flag.Debug)

	_ = rootCmd.RegisterFlagCompletionFunc(flag.Profile, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return config.List(), cobra.ShellCompDirectiveDefault
	})
	return rootCmd
}

const verTemplate = `%s version: %s
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

	return true
}

func shouldCheckCredentials(cmd *cobra.Command) AuthRequirements {
	searchByName := []string{
		"__complete",
		"help",
	}
	for _, n := range searchByName {
		if cmd.Name() == n {
			return NoAuth
		}
	}
	customRequirements := map[string]AuthRequirements{
		fmt.Sprintf("%s %s", atlas, "completion"):  NoAuth,       // completion commands do not require credentials
		fmt.Sprintf("%s %s", atlas, "config"):      NoAuth,       // user wants to set credentials
		fmt.Sprintf("%s %s", atlas, "auth"):        NoAuth,       // user wants to set credentials
		fmt.Sprintf("%s %s", atlas, "login"):       NoAuth,       // user wants to set credentials
		fmt.Sprintf("%s %s", atlas, "setup"):       NoAuth,       // user wants to set credentials
		fmt.Sprintf("%s %s", atlas, "register"):    NoAuth,       // user wants to set credentials
		fmt.Sprintf("%s %s", atlas, "quickstart"):  NoAuth,       // command supports login
		fmt.Sprintf("%s %s", atlas, "deployments"): OptionalAuth, // command supports local and Atlas
	}
	for p, r := range customRequirements {
		if strings.HasPrefix(cmd.CommandPath(), p) {
			return r
		}
	}
	return RequiredAuth
}

func formattedVersion() string {
	return fmt.Sprintf(verTemplate,
		config.ToolName,
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
		upgradeInstructions = fmt.Sprintf(`To upgrade, run "brew update && brew upgrade %s".`, homebrew.FormulaName(config.ToolName))
	} else {
		upgradeInstructions = "To upgrade, see: https://dochub.mongodb.org/core/install-atlas-cli."
	}

	newVersionTemplate := `
A new version of %s is available '%s'!
%s

To disable this alert, run "%s config set skip_update_check true".
`
	_, err = fmt.Fprintf(n.writer, newVersionTemplate, config.ToolName, release.Version, upgradeInstructions, config.BinName())
	return err
}

func isAtLeast24HoursPast(t time.Time) bool {
	return !t.IsZero() && time.Since(t) >= time.Hour*24
}

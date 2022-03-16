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
	"fmt"
	"io"
	"runtime"
	"strings"
	"time"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/alerts"
	"github.com/mongodb/mongocli/internal/cli/atlas/accesslists"
	"github.com/mongodb/mongocli/internal/cli/atlas/accesslogs"
	"github.com/mongodb/mongocli/internal/cli/atlas/backup"
	"github.com/mongodb/mongocli/internal/cli/atlas/cloudproviders"
	"github.com/mongodb/mongocli/internal/cli/atlas/clusters"
	atlasConfig "github.com/mongodb/mongocli/internal/cli/atlas/config"
	"github.com/mongodb/mongocli/internal/cli/atlas/customdbroles"
	"github.com/mongodb/mongocli/internal/cli/atlas/customdns"
	"github.com/mongodb/mongocli/internal/cli/atlas/datalake"
	"github.com/mongodb/mongocli/internal/cli/atlas/dbusers"
	"github.com/mongodb/mongocli/internal/cli/atlas/integrations"
	"github.com/mongodb/mongocli/internal/cli/atlas/livemigrations"
	"github.com/mongodb/mongocli/internal/cli/atlas/logs"
	"github.com/mongodb/mongocli/internal/cli/atlas/maintenance"
	"github.com/mongodb/mongocli/internal/cli/atlas/metrics"
	"github.com/mongodb/mongocli/internal/cli/atlas/networking"
	"github.com/mongodb/mongocli/internal/cli/atlas/privateendpoints"
	"github.com/mongodb/mongocli/internal/cli/atlas/processes"
	"github.com/mongodb/mongocli/internal/cli/atlas/quickstart"
	"github.com/mongodb/mongocli/internal/cli/atlas/security"
	"github.com/mongodb/mongocli/internal/cli/atlas/serverless"
	"github.com/mongodb/mongocli/internal/cli/auth"
	"github.com/mongodb/mongocli/internal/cli/events"
	"github.com/mongodb/mongocli/internal/cli/figautocomplete"
	"github.com/mongodb/mongocli/internal/cli/iam/organizations"
	"github.com/mongodb/mongocli/internal/cli/iam/projects"
	"github.com/mongodb/mongocli/internal/cli/iam/teams"
	"github.com/mongodb/mongocli/internal/cli/iam/users"
	"github.com/mongodb/mongocli/internal/cli/performanceadvisor"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/homebrew"
	"github.com/mongodb/mongocli/internal/latestrelease"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/mongodb/mongocli/internal/validate"
	"github.com/mongodb/mongocli/internal/version"
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

// Builder conditionally adds children commands as needed.
func Builder(profile *string) *cobra.Command {
	rootCmd := &cobra.Command{
		Version: version.Version,
		Use:     atlas,
		Aliases: []string{config.ToolName},
		Short:   "CLI tool to manage MongoDB Atlas",
		Long:    fmt.Sprintf("Use %s command help for information on a specific command", atlas),
		Example: `
  Display the help menu for the config command
  $ atlas config --help`,
		SilenceUsage: true,
		Annotations: map[string]string{
			"toc": "true",
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if config.Service() == "" {
				config.SetService(config.CloudService)
			}

			if cmd.Name() == figautocomplete.CmdUse { // figautocomplete command does not require credentials
				return nil
			}

			if cmd.Name() == "quickstart" { // quickstart has its own check
				return nil
			}

			if strings.HasPrefix(cmd.CommandPath(), fmt.Sprintf("%s %s", atlas, "config")) { // user wants to set credentials
				return nil
			}

			if strings.HasPrefix(cmd.CommandPath(), fmt.Sprintf("%s %s", atlas, "auth")) { // user wants to set credentials
				return nil
			}

			return validate.Credentials()
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

	rootCmd.AddCommand(
		atlasConfig.Builder(),
		auth.Builder(),
		quickstart.Builder(),
		projects.Builder(),
		organizations.AtlasCLIBuilder(),
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
		integrations.Builder(),
		maintenance.Builder(),
		customdns.Builder(),
		cloudproviders.Builder(),
		serverless.Builder(),
		livemigrations.Builder(),
		accesslogs.Builder(),
		loginCmd,
		logoutCmd,
		whoCmd,
		figautocomplete.Builder(),
	)

	rootCmd.PersistentFlags().StringVarP(profile, flag.Profile, flag.ProfileShort, "", usage.Profile)

	return rootCmd
}

const verTemplate = `%s version: %s
git version: %s
Go version: %s
   os: %s
   arch: %s
   compiler: %s
`

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
	shouldCheck = !config.SkipUpdateCheck() && cli.IsTerminal(n.writer)
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
		upgradeInstructions = fmt.Sprintf(`To upgrade, see: https://dochub.mongodb.org/core/%s-install.`, config.ToolName)
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

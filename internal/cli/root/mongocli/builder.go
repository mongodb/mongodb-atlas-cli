// Copyright 2021 MongoDB Inc
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

package mongocli

import (
	"fmt"
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/atlas"
	"github.com/mongodb/mongocli/internal/cli/auth"
	"github.com/mongodb/mongocli/internal/cli/cloudmanager"
	cliconfig "github.com/mongodb/mongocli/internal/cli/config"
	"github.com/mongodb/mongocli/internal/cli/figautocomplete"
	"github.com/mongodb/mongocli/internal/cli/iam"
	"github.com/mongodb/mongocli/internal/cli/opsmanager"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/homebrew"
	"github.com/mongodb/mongocli/internal/latestrelease"
	"github.com/mongodb/mongocli/internal/search"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/mongodb/mongocli/internal/version"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"io"
	"runtime"
)

type Notifier struct {
}

// Builder conditionally adds children commands as needed.
// This is important in particular for Atlas as it dynamically sets flags for cluster creation and
// this can be slow to timeout on environments with limited internet access (Ops Manager).
func Builder(profile *string, argsWithoutProg []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Version: version.Version,
		Use:     config.ToolName,
		Short:   "CLI tool to manage your MongoDB Cloud",
		Long:    fmt.Sprintf("Use %s command help for information on a specific command", config.ToolName),
		Example: `
  Display the help menu for the config command
  $ mongocli config --help`,
		SilenceUsage: true,
		Annotations: map[string]string{
			"toc": "true",
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			w := cmd.ErrOrStderr()
			if shouldSkipPrintNewVersion(w) {
				return
			}
			c, _ := homebrew.NewChecker(afero.NewOsFs())
			c.IsHomebrew()
			// p := NewPrinter(w, config.ToolName, config.BinName(), c.IsHomebrew())
			checker := latestrelease.NewChecker(version.Version, config.ToolName)
			// shouldCheck && isLatests{
			// print new vercions
			//}
			_ = checker.CheckAvailable()

		},
	}
	rootCmd.SetVersionTemplate(formattedVersion())
	hasArgs := len(argsWithoutProg) != 0

	if hasArgs && (argsWithoutProg[0] == "--version" || argsWithoutProg[0] == "-v") {
		return rootCmd
	}
	rootCmd.AddCommand(cliconfig.Builder())

	shouldIncludeAtlas := []string{
		atlas.Use,
		"help",
		"--help",
		"-h",
		"completion",
		"__complete",
		"fig-autocomplete",
	}
	if !hasArgs || search.StringInSlice(shouldIncludeAtlas, argsWithoutProg[0]) {
		rootCmd.AddCommand(atlas.Builder())
	}
	// hidden shortcuts
	loginCmd := auth.LoginBuilder()
	loginCmd.Hidden = true
	logoutCmd := auth.LogoutBuilder()
	logoutCmd.Hidden = true
	whoCmd := auth.WhoAmIBuilder()
	whoCmd.Hidden = true

	rootCmd.AddCommand(
		cloudmanager.Builder(),
		opsmanager.Builder(),
		iam.Builder(),
		auth.Builder(),
		loginCmd,
		logoutCmd,
		whoCmd,
		figautocomplete.Builder(),
	)

	rootCmd.PersistentFlags().StringVarP(profile, flag.Profile, flag.ProfileShort, "", usage.Profile)

	return rootCmd
}

type Printer interface {
	PrintNewVersionAvailable(latestVersion, homebrewCommand string) error
}

func NewPrinter(w io.Writer, t, b string) Printer {
	return &printer{
		writer: w,
		tool:   t,
		bin:    b,
	}
}

type printer struct {
	writer io.Writer
	tool   string
	bin    string
}

func (p *printer) PrintNewVersionAvailable(latestVersion, formulaName string) error {
	var upgradeInstructions string
	if formulaName != "" {
		upgradeInstructions = fmt.Sprintf(`To upgrade, run "brew update && brew upgrade %s".`, formulaName)
	} else {
		upgradeInstructions = fmt.Sprintf(`To upgrade, see: https://dochub.mongodb.org/core/%s-install.`, p.tool)
	}

	newVersionTemplate := `
A new version of %s is available '%s'!
%s

To disable this alert, run "%s config set skip_update_check true".
`
	_, err := fmt.Fprintf(p.writer, newVersionTemplate, p.tool, latestVersion, upgradeInstructions, p.bin)
	return err
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

func shouldSkipPrintNewVersion(w io.Writer) bool {
	return config.SkipUpdateCheck() || !cli.IsTerminal(w)
}

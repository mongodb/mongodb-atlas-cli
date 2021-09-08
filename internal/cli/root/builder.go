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

package root

import (
	"fmt"
	"io"
	"runtime"

	"github.com/Masterminds/semver/v3"
	"github.com/mongodb/mongocli/internal/cli/atlas"
	"github.com/mongodb/mongocli/internal/cli/cloudmanager"
	cliconfig "github.com/mongodb/mongocli/internal/cli/config"
	"github.com/mongodb/mongocli/internal/cli/iam"
	"github.com/mongodb/mongocli/internal/cli/opsmanager"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/search"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/mongodb/mongocli/internal/version"
	"github.com/spf13/cobra"
)

type BuilderOpts struct {
	store store.VersionDescriber
}

// rootBuilder conditionally adds children commands as needed.
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
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			opts := &BuilderOpts{}
			err := opts.initStore()
			if err != nil {
				return err
			}
			return opts.printNewVersionAvailable(cmd.ErrOrStderr())
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
	}
	if !hasArgs || search.StringInSlice(shouldIncludeAtlas, argsWithoutProg[0]) {
		rootCmd.AddCommand(atlas.Builder())
	}
	rootCmd.AddCommand(
		cloudmanager.Builder(),
		opsmanager.Builder(),
		iam.Builder(),
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

func (opts *BuilderOpts) hasNewVersionAvailable() (newVersionAvailable bool, newVersion string, err error) {
	if version.Version == "" {
		return false, "", nil
	}

	svCurrentVersion, err := semver.NewVersion(version.Version)
	if err != nil {
		return false, "", err
	}

	latestVersion, err := opts.store.LatestVersion()
	if err != nil {
		return false, "", err
	}

	svLatestVersion, err := semver.NewVersion(latestVersion)
	if err != nil {
		return false, "", err
	}

	if svCurrentVersion.Compare(svLatestVersion) < 0 {
		return true, latestVersion, nil
	}

	return false, "", nil
}

func (opts *BuilderOpts) printNewVersionAvailable(w io.Writer) error {
	if config.SkipUpdateCheck() {
		return nil
	}
	newVersionAvailable, latestVersion, err := opts.hasNewVersionAvailable()
	if err != nil {
		return err
	}
	if newVersionAvailable {
		newVersionTemplate := `
A new MongoCLI version is available '%s'!
To upgrade, check https://docs.mongodb.com/mongocli/stable/install/ (run "mongocli config set skip_update_check true" to disable the alert.)
`
		_, err = fmt.Fprintf(w, newVersionTemplate, latestVersion)
		return err
	}
	return nil
}

func (opts *BuilderOpts) initStore() error {
	var err error
	opts.store, err = store.New(store.UnauthenticatedPreset(config.Default()))
	return err
}

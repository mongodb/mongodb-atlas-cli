// Copyright 2020 MongoDB Inc
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

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"

	survey "github.com/AlecAivazis/survey/v2/core"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/root/mongocli"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/spf13/cobra"
)

var (
	profile string
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(ctx context.Context) {
	rootCmd := mongocli.Builder(&profile, os.Args[1:])
	// append here to avoid a recursive link on generated docs
	rootCmd.Long += `

To learn more, see our documentation: https://www.mongodb.com/docs/mongocli/stable/`
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}

func updateMongoCLIConfigPath() {
	mongoCLIConfigHome, err := config.MongoCLIConfigHome()
	if err != nil {
		return
	}

	mongoCLIConfigPath := path.Join(mongoCLIConfigHome, "config.toml")
	f, err := os.Open(mongoCLIConfigPath) // if config.toml is already there, exit
	if err == nil {
		f.Close()
		return
	}

	oldMongoCLIConfigHome, err := config.OldMongoCLIConfigHome() //nolint:staticcheck // Deprecated before fully removing support in the future
	if err != nil {
		return
	}

	oldMongoCLIConfigPath := path.Join(oldMongoCLIConfigHome, "mongocli.toml")
	in, err := os.Open(oldMongoCLIConfigPath)
	if err != nil {
		return
	}
	defer in.Close()

	_, _ = fmt.Fprintf(os.Stderr, `MongoCLI uses a new config path. Copying mongocli.toml content to: %s
`, mongoCLIConfigPath)

	// check if new config home already exists and create if not
	if _, err = os.Stat(mongoCLIConfigHome); err != nil {
		defaultPermissions := 0700
		if err = os.Mkdir(mongoCLIConfigHome, os.FileMode(defaultPermissions)); err != nil {
			log.Printf("There was an error generating %s: %v", mongoCLIConfigHome, err)
			return
		}
	}

	out, err := os.Create(mongoCLIConfigPath)
	if err != nil {
		return
	}
	defer out.Close()

	if _, err = io.Copy(out, in); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "There was an error generating %s: %v", mongoCLIConfigPath, err)
		return
	}
	defer os.Remove(oldMongoCLIConfigPath)

	_, _ = fmt.Fprintf(os.Stderr, `MongoCLI configuration moved to: %s
`, mongoCLIConfigPath)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if err := config.LoadMongoCLIConfig(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	if err := cli.InitProfile(profile); err != nil {
		log.Fatalf("Error loading profile: %v", err)
	}
}

func main() {
	cobra.EnableCommandSorting = false
	if term := os.Getenv("TERM"); strings.HasSuffix(term, "-m") {
		survey.DisableColor = true
	}
	cobra.OnInitialize(updateMongoCLIConfigPath, initConfig)

	Execute(context.Background())
}

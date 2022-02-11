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

package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/mongodb/mongocli/internal/cli/root/atlas"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	profile string
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(ctx context.Context) {
	rootCmd := atlas.Builder(&profile)
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if err := config.Load(); err != nil {
		// we use mongocli.toml to generate atlasCLI config
		var e viper.ConfigFileNotFoundError
		if !errors.As(err, &e) || !createConfigFromMongoCLIConfig() { // search mongoCLI config only if atlasCLI config doesn't exist
			printError(err)
		} else if err := config.Load(); err != nil {
			printError(err)
		}
	}

	if profile != "" {
		config.SetName(profile)
	} else if profile = config.GetString(flag.Profile); profile != "" {
		config.SetName(profile)
	} else if availableProfiles := config.List(); len(availableProfiles) == 1 {
		config.SetName(availableProfiles[0])
	}
}

// createConfigFromMongoCLIConfig creates atlasCLI config file from mongoCLI config file.
func createConfigFromMongoCLIConfig() bool {
	mongoCLIConfigPath := mongoCLIConfigPath()
	if mongoCLIConfigPath == "" {
		return false
	}

	atlasCLIConfigPath := copyConfig(mongoCLIConfigPath)
	if atlasCLIConfigPath == "" {
		return false
	}

	_, _ = fmt.Fprintf(os.Stderr, `we have used %s to generate %s

`, mongoCLIConfigPath, atlasCLIConfigPath)

	return true
}

// copyConfig copies config in oldConfigPath to the correct config location.
func copyConfig(oldConfigPath string) string {
	in, err := os.Open(oldConfigPath)
	if err != nil {
		return ""
	}
	defer in.Close()

	configHomePath, err := config.ConfigurationHomePath(config.ToolName)
	if err != nil {
		return ""
	}

	_, err = os.Stat(configHomePath) // check if the dir is already there
	if err != nil {
		defaultPermissions := 0700
		if err = os.Mkdir(configHomePath, os.FileMode(defaultPermissions)); err != nil {
			return ""
		}
	}

	newConfigPath := fmt.Sprintf("%s/%s", configHomePath, "config.toml")
	out, err := os.Create(newConfigPath)
	if err != nil {
		return ""
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return ""
	}

	return newConfigPath
}

func printError(err error) {
	var e viper.ConfigFileNotFoundError
	if !errors.As(err, &e) {
		log.Fatalf("Error loading config: %v", err)
	}
}

func mongoCLIConfigPath() string {
	configDir, err := config.ConfigurationHomePath("mongocli")
	if err != nil {
		return ""
	}

	configPath := fmt.Sprintf("%s/mongocli.toml", configDir)
	_, err = os.Stat(configPath)
	if err != nil {
		return ""
	}

	return configPath
}

func main() {
	cobra.EnableCommandSorting = false
	cobra.OnInitialize(initConfig)

	Execute(context.Background())
}

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
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2/core"
	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/commonerrors"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/root"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config/migrations"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func execute(ctx context.Context, rootCmd *cobra.Command) {
	// append here to avoid a recursive link on generated docs
	rootCmd.Long += `

To learn more, see our documentation: https://www.mongodb.com/docs/atlas/cli/stable/connect-atlas-cli/`
	if cmd, err := rootCmd.ExecuteContextC(ctx); err != nil {
		err := commonerrors.Check(err)
		rootCmd.PrintErrln(rootCmd.ErrPrefix(), err)
		if !telemetry.StartedTrackingCommand() {
			telemetry.StartTrackingCommand(cmd, os.Args[1:])
		}

		telemetry.FinishTrackingCommand(telemetry.TrackOptions{
			Err: err,
		})
		os.Exit(1)
	}
}

// loadConfig reads in config file and ENV variables if set.
func loadConfig() (*config.Profile, error) {
	// Migrate config to the latest version.
	migrator := migrations.NewDefaultMigrator()
	if err := migrator.Migrate(); err != nil {
		return nil, fmt.Errorf("error migrating config: %w", err)
	}

	configStore, initErr := config.NewDefaultStore()

	if initErr != nil {
		return nil, fmt.Errorf("error loading config: %w. Please run `atlas auth login` to reconfigure your profile", initErr)
	}

	if !configStore.IsSecure() {
		fmt.Fprintf(os.Stderr, "Warning: Secure storage is not available, falling back to insecure storage\n")
	}

	profile := config.NewProfile(config.DefaultProfile, configStore)
	config.SetProfile(profile)

	return profile, nil
}

func trackInitError(e error, rootCmd *cobra.Command) {
	if e == nil {
		return
	}
	if cmd, args, err := rootCmd.Find(os.Args[1:]); err == nil {
		if !telemetry.StartedTrackingCommand() {
			telemetry.StartTrackingCommand(cmd, args)
		}
		telemetry.FinishTrackingCommand(telemetry.TrackOptions{
			Err: e,
		})
	}
	log.Print(e)
}

func initTrack(rootCmd *cobra.Command) {
	cmd, args, _ := rootCmd.Find(os.Args[1:])
	telemetry.StartTrackingCommand(cmd, args)
}

func main() {
	cobra.EnableCommandSorting = false
	if term := os.Getenv("TERM"); strings.HasSuffix(term, "-m") {
		core.DisableColor = true
	}

	// Load config
	profile, loadProfileErr := loadConfig()

	rootCmd := root.Builder()
	initTrack(rootCmd)
	trackInitError(loadProfileErr, rootCmd)

	// Initialize context, attach
	// - telemetry
	// - profile
	ctx := telemetry.NewContext()
	config.WithProfile(ctx, profile)

	execute(ctx, rootCmd)
}

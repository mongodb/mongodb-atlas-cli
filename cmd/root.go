package cmd

import (
	"fmt"
	"log"

	"github.com/10gen/mcli/internal/cli"
	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/version"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Version: version.Version,
		Use:     config.Name,
		Short:   "CLI tool to manage your MongoDB Cloud",
		Long:    fmt.Sprintf("Use %s command help for information on a specific command", config.Name),
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// config commands
	rootCmd.AddCommand(cli.ConfigBuilder())
	// Atlas commands
	rootCmd.AddCommand(cli.AtlasBuilder())
	// C/OM commands
	rootCmd.AddCommand(cli.CloudManagerBuilder())
	// IAM commands
	rootCmd.AddCommand(cli.IAMBuilder())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if err := config.Load(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
}

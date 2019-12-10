package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	orgID   string
	profile string

	rootCmd = &cobra.Command{
		Version: Version,
		Use:     toolName,
		Short:   "CLI tool to manage your MongoDB Cloud",
		Long:    fmt.Sprintf("Use %s command help for information on a specific command", toolName),
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	exitOnErr(err)
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&profile, "profile", "p", "default", "profile")
}

func createConfigFile() {
	// TODO: viper to release patch for this
	configFile := fmt.Sprintf("%s/%s.toml", configDir(), toolName)

	_, err := os.OpenFile(configFile, os.O_RDONLY|os.O_CREATE, 0600)
	exitOnErr(err)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	configDir := configDir()
	viper.SetEnvPrefix(toolName)
	viper.AutomaticEnv()
	viper.SetConfigType(configType)
	viper.SetConfigName(toolName)
	viper.AddConfigPath(configDir) // path to look for the config file in

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			createConfigFile()
		} else {
			exitOnErr(err)
		}
	}
}

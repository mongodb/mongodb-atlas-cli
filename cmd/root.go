package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configName = "mcli"
)

var (
	version string
	orgID   string
	profile string

	rootCmd = &cobra.Command{
		Version: version,
		Use:     "mcli",
		Short:   "CLI tool to manage your mongoDB cloud",
		Long:    "Use mcli command help for information on a specific  command",
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
	configFile := fmt.Sprintf("%s/%s.toml", configDir(), configName)

	_, err := os.OpenFile(configFile, os.O_RDONLY|os.O_CREATE, 0600)
	exitOnErr(err)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	configDir := configDir()
	viper.SetEnvPrefix(configName)
	viper.AutomaticEnv()
	viper.SetConfigType("toml")
	viper.SetConfigName(configName)
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

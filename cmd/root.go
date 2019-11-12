package cmd

import (
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile   string
	version   string
	orgID     string
	projectID string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: version,
	Use:     "mpc",
	Short:   "MPC cli tool to manage your mongo cloud",
	Long:    "Use mpc command help for information on a  specific  command.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	exitOnErr(err)
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mpc.json)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		exitOnErr(err)
		_, err2 := os.OpenFile(home+"/.mpc.json", os.O_RDONLY|os.O_CREATE, 0600)
		exitOnErr(err2)

		viper.SetEnvPrefix("mpc")
		viper.AutomaticEnv()
		viper.SetConfigType("json")
		viper.SetConfigName(".mpc")
		// Search config in home directory with name ".mpc" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)

		if ok {
			log.Println("Config file not found :(")
		}
	}
}

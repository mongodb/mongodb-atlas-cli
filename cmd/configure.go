package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure the tool",
	Long:  `Set up authentication settings as well as the Base URL of the private cloud deployment.`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter baseURL: ")
		baseURL, _ := reader.ReadString('\n')

		viper.Set("base_url", strings.TrimSpace(baseURL))

		fmt.Print("Enter Public Key: ")
		username, _ := reader.ReadString('\n')

		viper.Set("public_key", strings.TrimSpace(username))

		fmt.Print("Enter Private Key: ")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		exitOnErr(err)
		password := string(bytePassword)
		viper.Set("private_key", strings.TrimSpace(password))

		err2 := viper.WriteConfig()
		exitOnErr(err2)

		fmt.Println("\nDone!")
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}

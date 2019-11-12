package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/mongodb-labs/pcgc/pkg/opsmanager"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

// projectsCmd represents the automation command
var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "Admin operations",
}

var whitelistIP string

// automationStatusCmd represents  status command
var adminCreate = &cobra.Command{
	Use:   "create",
	Short: "Create the first admin",
	Long:  "Create the first admin, this cna only be called once per installation",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter Username: ")
		username, _ := reader.ReadString('\n')

		fmt.Print("Enter First Name: ")
		firstName, _ := reader.ReadString('\n')

		fmt.Print("Enter Last Name: ")
		lastName, _ := reader.ReadString('\n')

		fmt.Print("Enter Password: ")
		bytePassword, err := terminal.ReadPassword(syscall.Stdin)
		exitOnErr(err)
		password := string(bytePassword)

		user := opsmanager.User{
			Username:  strings.TrimSpace(username),
			Password:  strings.TrimSpace(password),
			FirstName: strings.TrimSpace(firstName),
			LastName:  strings.TrimSpace(lastName),
		}
		userCreation, err := newDefaultClient().CreateFirstUser(user, whitelistIP)

		exitOnErr(err)

		prettyJSON(userCreation)
	},
}

func init() {
	adminCreate.Flags().StringVar(&whitelistIP, "whitelist-ip", "", "IPs")
	rootCmd.AddCommand(adminCmd)
	adminCmd.AddCommand(adminCreate)
}

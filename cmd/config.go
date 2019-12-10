package cmd

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	service string
	// configureCmd represents the configure command
	// mcli config [--service atlas|ops-manager|cloud-manager] [--profile profileName]
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "Configure the tool",
		Long:  `Set up authentication settings as well as the Base URL of the private cloud deployment.`,
		Run: func(cmd *cobra.Command, args []string) {
			viper.Set(fmt.Sprintf("%s.service", profile), service)

			if service == "ops-manager" {
				baseURLPrompt := promptui.Prompt{
					Label: "Base URL",
				}

				baseURL, err := baseURLPrompt.Run()

				exitOnErr(err)

				viper.Set(fmt.Sprintf("%s.base_url", profile), strings.TrimSpace(baseURL))
			}

			usernamePrompt := promptui.Prompt{
				Label: "Public Key",
			}

			username, err := usernamePrompt.Run()

			exitOnErr(err)

			viper.Set(fmt.Sprintf("%s.public_key", profile), strings.TrimSpace(username))

			prompt := promptui.Prompt{
				Label: "Enter Private Key",
				Mask:  '*',
			}

			password, err := prompt.Run()

			exitOnErr(err)

			viper.Set(fmt.Sprintf("%s.private_key", profile), strings.TrimSpace(password))

			err2 := viper.WriteConfig()
			exitOnErr(err2)

			fmt.Println("\nDone!")
		},
	}
)

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringVarP(&service, "service", "s", "cloud", "Service provider, Atlas, Cloud manager ot Ops manager")
}

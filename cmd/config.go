package cmd

import (
	"fmt"
	"strings"

	"github.com/10gen/mcli/internal/cli"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	service string
	// configureCmd represents the configure command
	// mcli config [--service cloud|ops-manager|cloud-manager] [--profile profileName]
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "Configure the tool",
		Long:  `Set up authentication settings as well as the Base URL of the private cloud deployment.`,
		Run: func(cmd *cobra.Command, args []string) {
			config := &cli.Configuration{
				Profile: profile,
			}

			config.SetService(service)

			if service == cli.OpsManagerService {
				baseURLPrompt := promptui.Prompt{
					Label: "Base URL",
				}

				baseURL, err := baseURLPrompt.Run()

				exitOnErr(err)
				config.SetOpsManagerURL(strings.TrimSpace(baseURL))
			}

			publicAPIKeyPrompt := promptui.Prompt{
				Label: "Public Key",
			}

			publicAPIKey, err := publicAPIKeyPrompt.Run()
			exitOnErr(err)

			config.SetPublicAPIKey(strings.TrimSpace(publicAPIKey))

			privateAPIKeyPrompt := promptui.Prompt{
				Label: "Enter Private Key",
				Mask:  '*',
			}

			privateAPIKey, err := privateAPIKeyPrompt.Run()
			exitOnErr(err)

			config.SetPrivateAPIKey(strings.TrimSpace(privateAPIKey))

			err = config.Save()
			exitOnErr(err)

			fmt.Println("\nDone!")
		},
	}
)

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringVarP(&service, "service", "s", "cloud", "Service provider, Atlas, Cloud manager or Ops manager")
}

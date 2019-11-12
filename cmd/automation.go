package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/mongodb-labs/pcgc/pkg/opsmanager"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// projectsCmd represents the automation command
var automationCmd = &cobra.Command{
	Use:   "automation",
	Short: "Automation operations",
	Long:  "Manage projects automation configs.",
}

// automationStatusCmd represents  status command
var automationStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Automation status",

	Run: func(cmd *cobra.Command, args []string) {
		automationStatus, err := newAuthenticatedClient().GetAutomationStatus(projectID)

		exitOnErr(err)

		prettyJSON(automationStatus)
	},
}

// automationStatusCmd represents  status command
var automationRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Automation retrieve",

	Run: func(cmd *cobra.Command, args []string) {
		automationStatus, err := newAuthenticatedClient().GetAutomationConfig(projectID)

		exitOnErr(err)

		prettyJSON(automationStatus)
	},
}

// automationStatusCmd represents  status command
var automationUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Automation update from a file",
	Long:  "Apply a new automation config to your group",

	Run: func(cmd *cobra.Command, args []string) {
		file, err := ioutil.ReadFile(file)
		exitOnErr(err)
		data := opsmanager.AutomationConfig{}

		err2 := json.Unmarshal(file, &data)
		exitOnErr(err2)

		_, err3 := newAuthenticatedClient().UpdateAutomationConfig(projectID, data)
		exitOnErr(err3)

		fmt.Println("Applying new configuration...")
	},
}

func aliasProjectIDToGroupID(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "group-id":
		name = "project-id"
	}
	return pflag.NormalizedName(name)
}

var file string

func init() {
	automationStatusCmd.Flags().StringVar(&projectID, "project-id", "", "Project ID, group-id can also be used")
	_ = automationStatusCmd.MarkFlagRequired("project-id")
	automationStatusCmd.Flags().SetNormalizeFunc(aliasProjectIDToGroupID)

	automationRetrieveCmd.Flags().StringVar(&projectID, "project-id", "", "Project ID, group-id can also be used")
	_ = automationRetrieveCmd.MarkFlagRequired("project-id")
	automationRetrieveCmd.Flags().SetNormalizeFunc(aliasProjectIDToGroupID)

	automationUpdateCmd.Flags().StringVar(&projectID, "project-id", "", "Project ID, group-id can also be used")
	automationUpdateCmd.Flags().StringVarP(&file, "file", "f", "", "File to read the config")
	_ = automationUpdateCmd.MarkFlagRequired("project-id")
	_ = automationUpdateCmd.MarkFlagRequired("file")
	automationUpdateCmd.Flags().SetNormalizeFunc(aliasProjectIDToGroupID)

	rootCmd.AddCommand(automationCmd)
	automationCmd.AddCommand(automationStatusCmd)
	automationCmd.AddCommand(automationRetrieveCmd)
	automationCmd.AddCommand(automationUpdateCmd)
}

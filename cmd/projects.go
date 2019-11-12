package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:     "projects",
	Short:   "Projects operations",
	Long:    "Create, List and manage your MongoDB private cloud projects.",
	Aliases: []string{"groups"},
}

// listCmd represents the list command
var listProjectsCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects",
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := newAuthenticatedClient().GetAllProjects()

		if err != nil {
			fmt.Println("Error:", err)
		}

		prettyJSON(projects)
	},
}

// createCmd represents the create command
var createProjectCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a project",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		newProject, err := newAuthenticatedClient().CreateOneProject(args[0], orgID)

		exitOnErr(err)
		prettyJSON(newProject)
	},
}

func init() {
	createProjectCmd.Flags().StringVar(&orgID, "org-id", "", "Organization ID for the group")
	rootCmd.AddCommand(projectsCmd)
	projectsCmd.AddCommand(listProjectsCmd)
	projectsCmd.AddCommand(createProjectCmd)
}

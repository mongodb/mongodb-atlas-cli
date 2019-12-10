package cmd

import (
	"context"
	"errors"

	"github.com/mongodb-labs/pcgc/cloudmanager"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:     "projects",
	Short:   "Projects operations",
	Long:    "Create, List and manage your MongoDB private cloud projects.",
	Aliases: []string{"project", "group", "groups"},
}

// listCmd represents the list command
var listProjectsCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := newAuthenticatedClient(profile)
		exitOnErr(err)
		viper.GetString("default.service")
		projects, _, err := client.(*cloudmanager.Client).Projects.GetAllProjects(context.Background())

		exitOnErr(err)

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
		client, err := newAuthenticatedClient(profile)

		exitOnErr(err)
		project := &cloudmanager.Project{Name: "Test"}
		newProject, _, err := client.(*cloudmanager.Client).Projects.Create(context.Background(), project)
		exitOnErr(err)
		prettyJSON(newProject)
	},
}

func init() {
	createProjectCmd.Flags().StringVar(&orgID, "orgId", "", "Organization ID for the group")
	rootCmd.AddCommand(projectsCmd)
	projectsCmd.AddCommand(listProjectsCmd)
	projectsCmd.AddCommand(createProjectCmd)
}

package cmd

import (
	"errors"

	"github.com/10gen/mcli/internal/cli"
	"github.com/spf13/cobra"
)

var (
	// projectsCmd represents the projects command
	projectsCmd = &cobra.Command{
		Use:     "projects",
		Short:   "Projects operations",
		Long:    "Create, list and manage your MongoDB Cloud projects.",
		Aliases: []string{"project"},
	}

	// listCmd represents the list command
	listProjectsCmd = &cobra.Command{
		Use:   "list",
		Short: "List projects",
		Run: func(cmd *cobra.Command, args []string) {
			config := &cli.Configuration{Profile: profile}
			service := &cli.Projects{Configuration: config}
			projects, _, err := service.ListProjects()
			exitOnErr(err)
			prettyJSON(projects)
		},
	}

	// createCmd represents the create command
	createProjectCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a project",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("requires a name argument")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			config := &cli.Configuration{Profile: profile}
			service := &cli.Projects{Configuration: config}
			project, _, err := service.CreateProject(args[0], orgID)
			exitOnErr(err)
			prettyJSON(project)
		},
	}
)

func init() {
	createProjectCmd.Flags().StringVar(&orgID, "orgId", "", "Organization ID for the project")
	rootCmd.AddCommand(projectsCmd)
	projectsCmd.AddCommand(listProjectsCmd)
	projectsCmd.AddCommand(createProjectCmd)
}

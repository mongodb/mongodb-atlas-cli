package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/mongodb-labs/pcgc/cloudmanager"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
			projects, _, err := listProjects()
			exitOnErr(err)
			prettyJSON(projects)
		},
	}

	// createCmd represents the create command
	createProjectCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a project",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires a name argument")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			project, _, err := createProject(args[0])
			exitOnErr(err)
			prettyJSON(project)
		},
	}
)

// listProjects encapsulate the logic to manage different cloud providers
func listProjects() (interface{}, *atlas.Response, error) {
	client, err := newAuthenticatedClient(profile)
	exitOnErr(err)
	service := viper.GetString(fmt.Sprintf("%s.service", profile))
	switch service {
	case "cloud":
		return client.(*atlas.Client).Projects.GetAllProjects(context.Background())
	case "cloud-manager", "ops-manager":
		return client.(*cloudmanager.Client).Projects.GetAllProjects(context.Background())
	default:
		return nil, nil, fmt.Errorf("unsupported service: %s", service)
	}
}

// createProject encapsulate the logic to manage different cloud providers
func createProject(name string) (interface{}, *atlas.Response, error) {
	client, err := newAuthenticatedClient(profile)
	exitOnErr(err)
	service := viper.GetString(fmt.Sprintf("%s.service", profile))
	switch service {
	case "cloud":
		project := &atlas.Project{Name: name, OrgID: orgID}
		return client.(*atlas.Client).Projects.Create(context.Background(), project)
	case "cloud-manager", "ops-manager":
		project := &cloudmanager.Project{Name: name, OrgID: orgID}
		return client.(*cloudmanager.Client).Projects.Create(context.Background(), project)
	default:
		return nil, nil, fmt.Errorf("unsupported service: %s", service)
	}
}

func init() {
	createProjectCmd.Flags().StringVar(&orgID, "orgId", "", "Organization ID for the project")
	rootCmd.AddCommand(projectsCmd)
	projectsCmd.AddCommand(listProjectsCmd)
	projectsCmd.AddCommand(createProjectCmd)
}

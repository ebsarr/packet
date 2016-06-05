package cmd

import (
	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage your projects.",
	// Long: ``,
}

var listProjectsCmd = &cobra.Command{
	Use:   "listall",
	Short: "Print out all projects",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ListProjects()
		return err
	},
}

var listProjectCmd = &cobra.Command{
	Use:   "list",
	Short: "Print out project associate with the given ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		projectID := GetProjectID(cmd)
		err := ListProject(projectID)
		return err
	},
}

var createProjectCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project",
	RunE: func(cmd *cobra.Command, args []string) error {
		name := cmd.Flag("name").Value.String()
		paymentID := cmd.Flag("payment-id").Value.String()
		err := CreateProject(name, paymentID)
		return err
	},
}

var deleteProjectCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a project",
	RunE: func(cmd *cobra.Command, args []string) error {
		projectID := GetProjectID(cmd)
		err := DeleteProject(projectID)
		return err
	},
}

var updateProjectCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a project",
	RunE: func(cmd *cobra.Command, args []string) error {
		projectID := GetProjectID(cmd)
		name := cmd.Flag("name").Value.String()
		paymentID := cmd.Flag("payment-id").Value.String()
		err := UpdateProject(projectID, name, paymentID)
		return err
	},
}

var listProjectEventsCmd = &cobra.Command{
	Use:   "list-events",
	Short: "View events by project ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		projectID := GetProjectID(cmd)
		err := ListProjectEvents(projectID)
		return err
	},
}

func init() {
	projectCmd.AddCommand(listProjectsCmd, listProjectCmd, createProjectCmd, deleteProjectCmd, updateProjectCmd, listProjectEventsCmd)
	RootCmd.AddCommand(projectCmd)

	// Flags for command: packet project list
	listProjectCmd.Flags().String("project-id", "", "Project ID")

	// Flags for command: packet project create
	createProjectCmd.Flags().String("name", "", "Project name")
	createProjectCmd.Flags().String("payment-id", "", "ID of the payment method to associate to this project")

	// Flags for command: packet project delete
	deleteProjectCmd.Flags().String("project-id", "", "Project ID")

	// Flags for command: packet project update
	updateProjectCmd.Flags().String("project-id", "", "Project ID")
	updateProjectCmd.Flags().String("name", "", "Project name")
	updateProjectCmd.Flags().String("payment-id", "", "ID of the payment method to associate to this project")
	
	// Flags for command: packet project list-events
	listProjectEventsCmd.Flags().String("project-id", "", "Project ID")
}

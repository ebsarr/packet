// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
		projectID := cmd.Flag("project-id").Value.String()
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
		projectID := cmd.Flag("project-id").Value.String()
		err := DeleteProject(projectID)
		return err
	},
}

var updateProjectCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a project",
	RunE: func(cmd *cobra.Command, args []string) error {
		projectID := cmd.Flag("project-id").Value.String()
		name := cmd.Flag("name").Value.String()
		paymentID := cmd.Flag("payment-id").Value.String()
		err := UpdateProject(projectID, name, paymentID)
		return err
	},
}

func init() {
	projectCmd.AddCommand(listProjectsCmd, listProjectCmd, createProjectCmd, deleteProjectCmd, updateProjectCmd)
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
}

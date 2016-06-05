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

// storageCmd represents the storage command
var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Manage your storages",
}

var listStoragesCmd = &cobra.Command{
	Use:   "listall",
	Short: "View your volumes",
	RunE: func(cmd *cobra.Command, args []string) error {
		projectID := GetProjectID(cmd)
		err := ListStorages(projectID)
		return err
	},
}

var createStorageCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a volume",
	RunE: func(cmd *cobra.Command, args []string) error {
		projectID := GetProjectID(cmd)
		description := cmd.Flag("desc").Value.String()
		plan := cmd.Flag("plan").Value.String()
		facility := cmd.Flag("facility").Value.String()
		size, err := cmd.Flags().GetInt("size")
		if err != nil {
			return err
		}
		e := CreateStorage(projectID, description, plan, facility, size)
		return e
	},
}

var listStorageCmd = &cobra.Command{
	Use:   "list",
	Short: "View volume by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		storageID := cmd.Flag("storage-id").Value.String()
		err := ListStorage(storageID)
		return err
	},
}

var updateStorageCmd = &cobra.Command{
	Use:   "update",
	Short: "Update volume",
	RunE: func(cmd *cobra.Command, args []string) error {
		storageID := cmd.Flag("storage-id").Value.String()
		description := cmd.Flag("desc").Value.String()
		size, err := cmd.Flags().GetInt("size")
		if err != nil {
			return err
		}
		locked, err := cmd.Flags().GetBool("lock")
		if err != nil {
			return err
		}
		e := UpdateStorage(storageID, description, size, locked)
		return e
	},
}

var deleteStorageCmd = &cobra.Command{
	Use:	"delete",
	Short:	"Delete storage",
	RunE:	func(cmd *cobra.Command, args []string) error {
		storageID := cmd.Flag("storage-id").Value.String()
		err := DeleteStorage(storageID)
		return err
	},
}

func init() {
	storageCmd.AddCommand(listStoragesCmd, createStorageCmd, listStorageCmd, updateStorageCmd, deleteStorageCmd)
	RootCmd.AddCommand(storageCmd)

	// Flags for command: packet storage listall
	listStoragesCmd.Flags().String("project-id", "", "Project ID")

	// Flags for command: packet storage create
	createStorageCmd.Flags().String("project-id", "", "Project ID")
	createStorageCmd.Flags().String("desc", "", "Description")
	createStorageCmd.Flags().String("plan", "storage_1", "storage_1 || storage_2")
	createStorageCmd.Flags().String("facility", "ewr1", "ewr1 || sjc1 || ams1")
	createStorageCmd.Flags().Int("size", 120, "Volume size")

	// Flags for command: packet storage list
	listStorageCmd.Flags().String("storage-id", "", "Storage ID")

	// Flags for command: packet storage update
	updateStorageCmd.Flags().String("storage-id", "", "Storage ID")
	updateStorageCmd.Flags().String("desc", "", "Description")
	updateStorageCmd.Flags().Int("size", 120, "Volume size")
	updateStorageCmd.Flags().Bool("lock", false, "Update and lock")
	
	// Flags for command: packet storage delete
	deleteStorageCmd.Flags().String("storage-id", "", "Storage ID")

}

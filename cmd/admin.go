package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

// adminCmd represents the admin command
var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "Manage projects, ssh keys, etc...",
}

// Profile commands

// addProfileCmd represents the configure command
var addProfileCmd = &cobra.Command{
	Use:   "add-profile",
	Short: "Set default configs for the packet cli.",
	Long: `Set default configs for the packet cli.

The following configurations are supported:
- default API key
  This default key will be used if "--key" flag is missing in command.
- default project ID
  This ID will be used if "--project-id" flag is missing in command.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := Configure()
		return err
	},
}

// list-profsCmd represents the list-profs command
var listProfilesCmd = &cobra.Command{
	Use:   "list-profiles",
	Short: "List configured profiles",
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flag("name").Value.String()
		fmt.Printf("%-10s\t%-32s\t%s\n", "NAME", "APIKEY", "DEFAULT PROJECT")
		fmt.Printf("%-10s\t%-32s\t%s\n", "----", "------", "---------------")
		confs, _ := ReadConfigs()
		if name != "" {
			if conf, found := confs.Profiles[name]; found {
				fmt.Printf("%-10s\t%s\n", name, conf)
			}
		}
		if confs != nil {
			for profile, conf := range confs.Profiles {
				fmt.Printf("%-10s\t%s\n", profile, conf)
			}
		}
	},
}

var deleteProfileCmd = &cobra.Command{
	Use:   "delete-profile",
	Short: "Delete a profile",
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flag("name").Value.String()
		confs, _ := ReadConfigs()
		if name != "" {
			if _, ok := confs.Profiles[name]; ok {
				delete(confs.Profiles, name)
				writeConfigFile(confs)
			}
		}
	},
}

// Project commands

var listProjectsCmd = &cobra.Command{
	Use:   "list-projects",
	Short: "Retrieve all projects",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ListProjects()
		return err
	},
}

var listProjectCmd = &cobra.Command{
	Use:   "list-project",
	Short: "Retrieve a project by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		projectID := GetProjectID(cmd)
		err := ListProject(projectID)
		return err
	},
}

var createProjectCmd = &cobra.Command{
	Use:   "create-project",
	Short: "Create a new project",
	RunE: func(cmd *cobra.Command, args []string) error {
		name := cmd.Flag("name").Value.String()
		paymentID := cmd.Flag("payment-id").Value.String()
		err := CreateProject(name, paymentID)
		return err
	},
}

var deleteProjectCmd = &cobra.Command{
	Use:   "delete-project",
	Short: "Delete a project",
	RunE: func(cmd *cobra.Command, args []string) error {
		projectID := GetProjectID(cmd)
		err := DeleteProject(projectID)
		return err
	},
}

var updateProjectCmd = &cobra.Command{
	Use:   "update-project",
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
	Use:   "list-project-events",
	Short: "View events by project ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		projectID := GetProjectID(cmd)
		err := ListProjectEvents(projectID)
		return err
	},
}

// ssh key management commands

var listSSHKeysCmd = &cobra.Command{
	Use:   "list-sshkeys",
	Short: "View all configured SSH keys",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ListSSHKeys()
		return err
	},
}

var listSSHKeyCmd = &cobra.Command{
	Use:   "list-sshkey",
	Short: "View configured SSH key associated with the given ID.",
	RunE: func(cmd *cobra.Command, args []string) error {
		sshKeyID := cmd.Flag("key-id").Value.String()
		err := ListSSHKey(sshKeyID)
		return err
	},
}

var createSSHKeyCmd = &cobra.Command{
	Use:   "create-sshkey",
	Short: "Configure a new SSH key",
	RunE: func(cmd *cobra.Command, args []string) error {
		keyFile := cmd.Flag("file").Value.String()
		key, err := ioutil.ReadFile(keyFile)
		if err != nil {
			return err
		}
		label := cmd.Flag("label").Value.String()
		err = CreateSSHKey(label, string(key))
		return err
	},
}

var deleteSSHKeyCmd = &cobra.Command{
	Use:   "delete-sshkey",
	Short: "Delete SSH key associated with the given ID.",
	RunE: func(cmd *cobra.Command, args []string) error {
		sshKeyID := cmd.Flag("key-id").Value.String()
		err := DeleteSSHKey(sshKeyID)
		return err
	},
}

var updateSSHKeyCmd = &cobra.Command{
	Use:   "update-sshkey",
	Short: "Update a SSH key: change the key or its label",
	RunE: func(cmd *cobra.Command, args []string) error {
		sshKeyID := cmd.Flag("key-id").Value.String()
		keyFile := cmd.Flag("file").Value.String()
		key, err := ioutil.ReadFile(keyFile)
		if err != nil {
			return err
		}
		label := cmd.Flag("label").Value.String()
		err = UpdateSSHKey(sshKeyID, label, string(key))
		return err
	},
}

// listing os command

// listOSCmd represents the OS command
var listOSCmd = &cobra.Command{
	Use:   "list-os",
	Short: "View available operating systems",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ListOS()
		return err
	},
}

// Facilities

// listFacilitiesCmd represents the list-facilities command
var listFacilitiesCmd = &cobra.Command{
	Use:   "list-facilities",
	Short: "View a list of facilities(packet DCs)",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ListFacilities()
		return err
	},
}

// Plans

// planCmd represents the plan command
var listPlansCmd = &cobra.Command{
	Use:   "list-plans",
	Short: "View available plans.",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ListPlans()
		return err
	},
}

// Spot market

// spotPricesCmd represents the spot-prices command
var spotPricesCmd = &cobra.Command{
	Use:   "spot-prices",
	Short: "View spot market prices",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := SpotMarketPrices()
		return err
	},
}

func init() {
	adminCmd.AddCommand(addProfileCmd, listProfilesCmd, deleteProfileCmd, listProjectsCmd, listProjectCmd, createProjectCmd, updateProjectCmd, deleteProjectCmd, listProjectEventsCmd, listSSHKeysCmd, listSSHKeyCmd, createSSHKeyCmd, deleteSSHKeyCmd, updateSSHKeyCmd, listOSCmd, listFacilitiesCmd, listPlansCmd, spotPricesCmd)
	RootCmd.AddCommand(adminCmd)

	listProfilesCmd.Flags().StringP("name", "n", "", "Profile name")
	deleteProfileCmd.Flags().StringP("name", "n", "", "Profile name")

	// Flags for command: packet admin list-project
	listProjectCmd.Flags().String("project-id", "", "Project ID")

	// Flags for command: packet admin create-project
	createProjectCmd.Flags().String("name", "", "Project name")
	createProjectCmd.Flags().String("payment-id", "", "ID of the payment method to associate to this project")

	// Flags for command: packet admin delete-project
	deleteProjectCmd.Flags().String("project-id", "", "Project ID")

	// Flags for command: packet admin update-project
	updateProjectCmd.Flags().String("project-id", "", "Project ID")
	updateProjectCmd.Flags().String("name", "", "Project name")
	updateProjectCmd.Flags().String("payment-id", "", "ID of the payment method to associate to this project")

	// Flags for command: packet admin list-project-events
	listProjectEventsCmd.Flags().String("project-id", "", "Project ID")

	// Flags for command: packet admin list-sshkeys
	listSSHKeyCmd.Flags().String("key-id", "", "SSH key ID")

	//Flags for command: packet admin create-sshkey
	createSSHKeyCmd.Flags().String("label", "", "Label to assign to the key")
	createSSHKeyCmd.Flags().StringP("file", "f", "", "Read public key from file.")

	// Flags for command: packet admin delete-sshkey
	deleteSSHKeyCmd.Flags().String("key-id", "", "SSH key ID")

	// Flags for command: packet admin update-sshkey
	updateSSHKeyCmd.Flags().String("key-id", "", "SSH key ID")
	updateSSHKeyCmd.Flags().String("label", "", "Label to assign to the key")
	updateSSHKeyCmd.Flags().StringP("file", "f", "", "Read public key from file.")
}

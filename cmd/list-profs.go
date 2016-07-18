package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// profileCmd represents the device command
var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Manage your profiles",
	//Long: ``,
}

// list-profsCmd represents the list-profs command
var listProfsCmd = &cobra.Command{
	Use:   "list",
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
		for profile, conf := range confs.Profiles {
			fmt.Printf("%-10s\t%s\n", profile, conf)
		}
	},
}

var deleteProfCmd = &cobra.Command{
	Use:   "delete",
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

func init() {
	profileCmd.AddCommand(listProfsCmd, deleteProfCmd)
	RootCmd.AddCommand(profileCmd)

	listProfsCmd.Flags().StringP("name", "n", "", "Profile name")
	deleteProfCmd.Flags().StringP("name", "n", "", "Profile name")
}

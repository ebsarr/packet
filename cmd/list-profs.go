package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// list-profsCmd represents the list-profs command
var listProfsCmd = &cobra.Command{
	Use:   "list-profiles",
	Short: "List configured profiles",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%-10s\t%-32s\t%s\n", "NAME", "APIKEY", "DEFAULT PROJECT")
		fmt.Printf("%-10s\t%-32s\t%s\n", "----", "------", "---------------")
		confs, err := ReadConfigs()
		if err != nil {
			return nil
		}
		for profile, conf := range confs.Profiles {
			fmt.Printf("%-10s\t%-32s\t%s\n", profile, conf.APIKEY, conf.DefaultProjectID)
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(listProfsCmd)
}

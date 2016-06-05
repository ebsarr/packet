package cmd

import (
	"github.com/spf13/cobra"
)

// listOSCmd represents the OS command
var listOSCmd = &cobra.Command{
	Use:   "list-os",
	Short: "View available operating systems",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ListOS()
		return err
	},
}

func init() {
	RootCmd.AddCommand(listOSCmd)
}

package cmd

import (
	"github.com/spf13/cobra"
)

// planCmd represents the plan command
var listPlansCmd = &cobra.Command{
	Use:   "list-plans",
	Short: "View available plans.",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ListPlans()
		return err
	},
}

func init() {
	RootCmd.AddCommand(listPlansCmd)
}

package cmd

import (
	"github.com/spf13/cobra"
)

// planCmd represents the plan command
var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "View available plans.",
	// Long: ``,
}

var listPlanCmd = &cobra.Command{
	Use:   "list",
	Short: "Print out available plans.",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ListPlans()
		return err
	},
}

func init() {
	planCmd.AddCommand(listPlanCmd)
	RootCmd.AddCommand(planCmd)
}

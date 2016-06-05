package cmd

import (
	"github.com/spf13/cobra"
)

// listFacilitiesCmd represents the list-facilities command
var listFacilitiesCmd = &cobra.Command{
	Use:   "list-facilities",
	Short: "View a list of facilities(packet DCs)",
	RunE:	func(cmd *cobra.Command, args []string) error {
		err := ListFacilities()
		return err
	},
}

func init() {
	RootCmd.AddCommand(listFacilitiesCmd)
}

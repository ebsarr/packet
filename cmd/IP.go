package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// IPCmd represents the IP command
var IPCmd = &cobra.Command{
	Use:   "IP",
	Short: "Manage device IP addresses",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("IP called")
	},
}

func init() {
	RootCmd.AddCommand(IPCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// IPCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// IPCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

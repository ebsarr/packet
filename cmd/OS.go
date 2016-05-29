package cmd

import (
	"github.com/spf13/cobra"
)

// OSCmd represents the OS command
var OSCmd = &cobra.Command{
	Use:   "OS",
	Short: "View available operating systems",
	// Long: ``,
	// Run: func(cmd *cobra.Command, args []string) {
	// TODO: Work your own magic here
	//fmt.Println("OS called")
	// },
}

// OSListCmd represents the OS list command
var OSListCmd = &cobra.Command{
	Use:   "list",
	Short: "Print out available operating systems",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ListOS()
		return err
	},
}

func init() {
	OSCmd.AddCommand(OSListCmd)
	RootCmd.AddCommand(OSCmd)
}

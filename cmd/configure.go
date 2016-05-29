package cmd

import (
	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Enable default configs for the packet cli.",
	// Long: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := Configure()
		return err
	},
}

func init() {
	RootCmd.AddCommand(configureCmd)
}

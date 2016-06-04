package cmd

import (
	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
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

func init() {
	RootCmd.AddCommand(configureCmd)
}

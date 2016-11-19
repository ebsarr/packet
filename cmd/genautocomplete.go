package cmd

import (
	"github.com/spf13/cobra"
)

// genautocompleteCmd represents the genautocomplete command
var genautocompleteCmd = &cobra.Command{
	Use:    "genautocomplete",
	Short:  "Generate a auto-completion script for the packet command",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		RootCmd.GenBashCompletionFile("packet-autocomplete.sh")
	},
}

func init() {
	RootCmd.AddCommand(genautocompleteCmd)
}

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// gendocCmd represents the gendoc command
var gendocCmd = &cobra.Command{
	Use:    "gendoc",
	Short:  "Generate a markdown documentation file for this tool",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		doc.GenMarkdownTree(RootCmd, "./doc")
	},
}

func init() {
	RootCmd.AddCommand(gendocCmd)
}

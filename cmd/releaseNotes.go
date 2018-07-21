package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// releaseNotesCmd represents the releaseNotes command
var releaseNotesCmd = &cobra.Command{
	Use:    "release-notes",
	Short:  "show release notes",
	Hidden: true,
	Long:   "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s v%s %v/%v\n", cmd.Parent().CommandPath(), version, runtime.GOOS, runtime.GOARCH)
		fmt.Println()
		fmt.Println("RELEASE NOTES")
		fmt.Println(releaseNotes)
	},
}

func init() {
	RootCmd.AddCommand(releaseNotesCmd)
}

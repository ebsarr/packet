package cmd

import (
	"github.com/spf13/cobra"
)

// eventCmd represents the event command
var eventCmd = &cobra.Command{
	Use:   "event",
	Short: "View operation logs on a device",
}

var listEventsCmd = &cobra.Command{
	Use:	"listall",
	Short:	"View events by device ID",
	RunE:	func(cmd *cobra.Command, args []string) error {
		deviceID := cmd.Flag("device-id").Value.String()
		err := ListEvents(deviceID)
		return err
	},
}

var listEventCmd = &cobra.Command{
	Use:	"list",
	Short:	"View event by ID",
	RunE:	func(cmd *cobra.Command, args []string) error {
		eventID := cmd.Flag("event-id").Value.String()
		err := ListEvent(eventID)
		return err
	},
}

func init() {
	eventCmd.AddCommand(listEventsCmd, listEventCmd)
	RootCmd.AddCommand(eventCmd)

	// Flags for command: packet event listall
	listEventsCmd.Flags().String("device-id", "", "Device ID")
	
	// Flags for command: packet event list
	listEventCmd.Flags().String("event-id", "", "Event ID")
}

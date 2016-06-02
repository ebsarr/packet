package cmd

import (
	"github.com/spf13/cobra"
)

// IPCmd represents the IP command
var IPCmd = &cobra.Command{
	Use:   "IP",
	Short: "Manage device IP addresses",
}

var listIPCmd = &cobra.Command{
	Use:   "list",
	Short: "Prints out IP address by ID",
	RunE:	func(cmd *cobra.Command, args []string) error {
		addressID := cmd.Flag("address-id").Value.String()
		err := ListIPAddress(addressID)
		return err
	},
}

var assignIPCmd = &cobra.Command{
	Use:   "assign",
	Short: "Assign an IP address to a device by ID",
	RunE:	func(cmd *cobra.Command, args []string) error {
		address := cmd.Flag("address").Value.String()
		deviceID := cmd.Flag("device-id").Value.String()
		err := AssignIPAddress(deviceID, address)
		return err
	},
}

var unAssignIPCmd = &cobra.Command{
	Use:   "unassign",
	Short: "Unassign an IP address from a device",
	RunE:	func(cmd *cobra.Command, args []string) error {
		addressID := cmd.Flag("address-id").Value.String()
		err := UnAssignIPAddress(addressID)
		return err
	},
}

func init() {
	IPCmd.AddCommand(listIPCmd, assignIPCmd, unAssignIPCmd)
	RootCmd.AddCommand(IPCmd)

	// Flags for command: packet IP list
	listIPCmd.Flags().String("address-id", "", "IP address ID")
	
	// Flags for command: packet IP assign
	assignIPCmd.Flags().String("address", "", "IP address.(format: x.x.x.x/y)")
	assignIPCmd.Flags().String("device-id", "", "ID of device to assign to")
	
	// Flags for command: packet IP unassign
	unAssignIPCmd.Flags().String("address-id", "", "IP address ID")
}

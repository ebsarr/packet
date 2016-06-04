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
	RunE: func(cmd *cobra.Command, args []string) error {
		addressID := cmd.Flag("address-id").Value.String()
		err := ListIPAddress(addressID)
		return err
	},
}

var assignIPCmd = &cobra.Command{
	Use:   "assign",
	Short: "Assign an IP address to a device by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		address := cmd.Flag("address").Value.String()
		deviceID := cmd.Flag("device-id").Value.String()
		err := AssignIPAddress(deviceID, address)
		return err
	},
}

var unAssignIPCmd = &cobra.Command{
	Use:   "unassign",
	Short: "Unassign an IP address from a device",
	RunE: func(cmd *cobra.Command, args []string) error {
		addressID := cmd.Flag("address-id").Value.String()
		err := UnAssignIPAddress(addressID)
		return err
	},
}

var listReservationsCmd = &cobra.Command{
	Use:   "list-reservations",
	Short: "Print out a list of IP resevations for a single project",
	RunE: func(cmd *cobra.Command, args []string) error {
		projectID := GetProjectID(cmd)
		err := ListIPReservations(projectID)
		return err
	},
}

var requestMoreIPReservationsCmd = &cobra.Command{
	Use:   "request-more",
	Short: "Request more IP space for a project in order to have additional IP addresses to assign to devices",
	RunE: func(cmd *cobra.Command, args []string) error {
		projectID := GetProjectID(cmd)
		ipType := cmd.Flag("type").Value.String()
		quantity, err := cmd.Flags().GetInt("quantity")
		if err != nil {
			return err
		}
		comments := cmd.Flag("comments").Value.String()

		e := RequestMoreIPReservations(projectID, ipType, comments, quantity)
		return e
	},
}

var listReservationCmd = &cobra.Command{
	Use:   "list-reservation",
	Short: "Print out a single IP reservation object by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		reservationID := cmd.Flag("reservation-id").Value.String()
		err := ListIPReservations(reservationID)
		return err
	},
}

var removeReservationCmd = &cobra.Command{
	Use:   "remove-reservation",
	Short: "Remove an IP reservation from the project",
	RunE: func(cmd *cobra.Command, args []string) error {
		reservationID := cmd.Flag("reservation-id").Value.String()
		err := RemoveIPReservation(reservationID)
		return err
	},
}

func init() {
	IPCmd.AddCommand(listIPCmd, assignIPCmd, unAssignIPCmd, listReservationsCmd, listReservationCmd, requestMoreIPReservationsCmd, removeReservationCmd)
	RootCmd.AddCommand(IPCmd)

	// Flags for command: packet IP list
	listIPCmd.Flags().String("address-id", "", "IP address ID")

	// Flags for command: packet IP assign
	assignIPCmd.Flags().String("address", "", "IP address.(format: x.x.x.x/y)")
	assignIPCmd.Flags().String("device-id", "", "ID of device to assign to")

	// Flags for command: packet IP unassign
	unAssignIPCmd.Flags().String("address-id", "", "IP address ID")

	// Flags for command: packet IP list-reservations
	listReservationsCmd.Flags().String("project-id", "", "Project ID")

	// Flags for command: packet IP list-reservation
	listReservationCmd.Flags().String("reservation-id", "", "Reservation ID")

	// Flags for command: packet IP remove-reservation
	removeReservationCmd.Flags().String("reservation-id", "", "Reservation ID")

	// Flags for command: packet IP request-more
	requestMoreIPReservationsCmd.Flags().String("project-id", "", "Project ID")
	requestMoreIPReservationsCmd.Flags().String("type", "public_ipv4", "public_ipv4 || global_ipv4")
	requestMoreIPReservationsCmd.Flags().String("comments", "", "Comment to Packet team")
	requestMoreIPReservationsCmd.Flags().Int("quantity", 1, "How many IPv4 you want to request. Options: 1, 2, 4, 8, 16, 32, 64, 128, 256")
}

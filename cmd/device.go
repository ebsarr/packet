package cmd

import (
	"github.com/spf13/cobra"
)

var silent bool

// deviceCmd represents the device command
var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "Manage your devices",
	//Long: ``,
}

var listDevicesCmd = &cobra.Command{
	Use:   "listall",
	Short: "Print out all devices in a project",
	RunE: func(cmd *cobra.Command, args []string) error {
		projectID := GetProjectID(cmd)
		err := ListDevices(projectID)
		return err
	},
}

var listDeviceCmd = &cobra.Command{
	Use:   "list",
	Short: "Print out device info for the given device ID.",
	RunE: func(cmd *cobra.Command, args []string) error {
		deviceID := cmd.Flag("device-id").Value.String()
		err := ListDevice(deviceID)
		return err
	},
}

var createDeviceCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new device",
	RunE: func(cmd *cobra.Command, args []string) error {
		projectID := GetProjectID(cmd)
		hostname := cmd.Flag("hostname").Value.String()
		plan := cmd.Flag("plan").Value.String()
		facility := cmd.Flag("facility").Value.String()
		osType := cmd.Flag("os-type").Value.String()
		billing := cmd.Flag("billing").Value.String()
		// tags := cmd.Flag("tags").Value.String()
		if silent {
			err := CreateDevice(projectID, hostname, plan, facility, osType, billing, []string{})
			return err
		}
		err := CreateDeviceVerbose(projectID, hostname, plan, facility, osType, billing, []string{})
		return err
	},
}

var deleteDeviceCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete device",
	RunE: func(cmd *cobra.Command, args []string) error {
		deviceID := cmd.Flag("device-id").Value.String()
		err := DeleteDevice(deviceID)
		return err
	},
}

var lockDeviceCmd = &cobra.Command{
	Use:   "lock",
	Short: "Lock device",
	RunE: func(cmd *cobra.Command, args []string) error {
		deviceID := cmd.Flag("device-id").Value.String()
		err := LockDevice(deviceID)
		return err
	},
}

var unlockDeviceCmd = &cobra.Command{
	Use:   "unlock",
	Short: "Unlock device",
	RunE: func(cmd *cobra.Command, args []string) error {
		deviceID := cmd.Flag("device-id").Value.String()
		err := UnlockDevice(deviceID)
		return err
	},
}

var powerOnDeviceCmd = &cobra.Command{
	Use:   "power-on",
	Short: "Power on device",
	RunE: func(cmd *cobra.Command, args []string) error {
		deviceID := cmd.Flag("device-id").Value.String()
		err := PowerOnDevice(deviceID)
		return err
	},
}

var powerOffDeviceCmd = &cobra.Command{
	Use:   "power-off",
	Short: "Power off device",
	RunE: func(cmd *cobra.Command, args []string) error {
		deviceID := cmd.Flag("device-id").Value.String()
		err := PowerOffDevice(deviceID)
		return err
	},
}

var rebootDeviceCmd = &cobra.Command{
	Use:   "reboot",
	Short: "Reboot device",
	RunE: func(cmd *cobra.Command, args []string) error {
		deviceID := cmd.Flag("device-id").Value.String()
		err := RebootDevice(deviceID)
		return err
	},
}

func init() {
	// Subcommands
	deviceCmd.AddCommand(listDevicesCmd, listDeviceCmd, createDeviceCmd, deleteDeviceCmd, lockDeviceCmd, unlockDeviceCmd, powerOnDeviceCmd, powerOffDeviceCmd, rebootDeviceCmd)

	// add command to root
	RootCmd.AddCommand(deviceCmd)

	// Flags for command: packet device listall
	listDevicesCmd.Flags().String("project-id", "", "Specify the project ID.")

	// Flags for command: packet device list
	listDeviceCmd.Flags().String("device-id", "", "Specify ID of device to display.")

	// Flags for command: packet device create
	createDeviceCmd.Flags().String("project-id", "", "The project ID.")
	createDeviceCmd.Flags().String("hostname", "", "Hostname to assign to the created device.")
	createDeviceCmd.Flags().String("plan", "baremetal_0", "Server type to create the device.")
	createDeviceCmd.Flags().String("facility", "", "DC location. Available values are sjc1: Sunnyvale CA, ewr1: Parsippany NJ, ams1: Amsterdam NL")
	createDeviceCmd.Flags().String("os-type", "centos_7", "Operating system to deploy to the server.")
	createDeviceCmd.Flags().String("billing", "hourly", "Choose \"hourly\" or \"monthly\" billing.")
	createDeviceCmd.Flags().BoolVarP(&silent, "silent", "s", false, "Omit provisioning logs")

	// Flags for other device commands that require the device ID.
	deleteDeviceCmd.Flags().String("device-id", "", "Device ID")
	lockDeviceCmd.Flags().String("device-id", "", "Device ID")
	unlockDeviceCmd.Flags().String("device-id", "", "Device ID")
	powerOnDeviceCmd.Flags().String("device-id", "", "Device ID")
	powerOffDeviceCmd.Flags().String("device-id", "", "Device ID")
	rebootDeviceCmd.Flags().String("device-id", "", "Device ID")
}

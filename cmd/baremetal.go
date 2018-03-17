package cmd

import (
	"errors"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var silent, spotInstance, alwaysPXE bool
var spotPriceMax float64

// baremetalCmd represents the baremetal command
var baremetalCmd = &cobra.Command{
	Use:   "baremetal",
	Short: "Manage server devices.",
	// Long: ``,
}

var listDevicesCmd = &cobra.Command{
	Use:   "list-devices",
	Short: "Retrieve all devices in a project",
	RunE: func(cmd *cobra.Command, args []string) error {
		projectID := GetProjectID(cmd)
		err := ListDevices(projectID)
		return err
	},
}

var listDeviceCmd = &cobra.Command{
	Use:   "list-device",
	Short: "Retrieve a device",
	RunE: func(cmd *cobra.Command, args []string) error {
		deviceID := cmd.Flag("device-id").Value.String()
		err := ListDevice(deviceID)
		return err
	},
}

var createDeviceCmd = &cobra.Command{
	Use:   "create-device",
	Short: "Create a new device",
	RunE: func(cmd *cobra.Command, args []string) error {
		var userData string
		projectID := GetProjectID(cmd)
		hostname := cmd.Flag("hostname").Value.String()
		plan := cmd.Flag("plan").Value.String()
		facility := cmd.Flag("facility").Value.String()
		osType := cmd.Flag("os-type").Value.String()
		billing := cmd.Flag("billing").Value.String()
		ipxeScriptURL := cmd.Flag("ipxe-script-url").Value.String()
		// for getting userdata, --userfile has higher priority.
		userData = cmd.Flag("userdata").Value.String()
		userDataFile := cmd.Flag("userfile").Value.String()
		if userDataFile == "" {
			userDataFile = cmd.Flag("file").Value.String()
		}
		if userDataFile != "" {
			data, err := ioutil.ReadFile(userDataFile)
			if err != nil {
				return err
			}
			userData = string(data)
		}

		// tags := cmd.Flag("tags").Value.String()
		if silent {
			err := CreateDevice(projectID, hostname, plan, facility, osType, billing, userData, ipxeScriptURL, []string{}, spotInstance, alwaysPXE, spotPriceMax)
			return err
		}
		err := CreateDeviceVerbose(projectID, hostname, plan, facility, osType, billing, userData, ipxeScriptURL, []string{}, spotInstance, alwaysPXE, spotPriceMax)
		return err
	},
}

var updateDeviceCmd = &cobra.Command{
	Use:   "update-device",
	Short: "Update a device",
	RunE: func(cmd *cobra.Command, args []string) error {
		var userData, userDataFile, hostname, description, ipxeScriptURL string
		var locked, alwaysPXE bool
		// var userDataFile string
		deviceID := cmd.Flag("device-id").Value.String()

		// Retrieve current device info
		client, err := NewPacketClient()
		if err != nil {
			return err
		}
		d, _, err := client.Devices.Get(deviceID)
		if err != nil {
			return err
		}

		if cmd.Flag("hostname").Changed {
			hostname = cmd.Flag("hostname").Value.String()
		} else {
			hostname = d.Hostname
		}

		if cmd.Flag("description").Changed {
			description = cmd.Flag("description").Value.String()
		}

		alwaysPXEFlag := cmd.Flag("always-pxe").Value.String()
		if alwaysPXEFlag != "" {
			if alwaysPXEFlag == "true" {
				alwaysPXE = true
			} else if alwaysPXEFlag == "false" {
				alwaysPXE = false
			} else {
				return errors.New("Bad value; --always-pxe only accepts true || false")
			}
		} else {
			alwaysPXE = d.AlwaysPXE
		}

		if cmd.Flag("ipxe-script-url").Changed {
			ipxeScriptURL = cmd.Flag("ipxe-script-url").Value.String()
		} else {
			ipxeScriptURL = d.IPXEScriptURL
		}

		lockFlag := cmd.Flag("lock").Value.String()
		if lockFlag != "" {
			if lockFlag == "true" {
				locked = true
			} else if lockFlag == "false" {
				locked = false
			} else {
				return errors.New("Bad value; --lock only accepts true || false")
			}
		} else {
			locked = d.Locked
		}

		// for getting userdata, --userfile has higher priority.
		if !cmd.Flag("userfile").Changed {
			if cmd.Flag("userdata").Changed {
				userData = cmd.Flag("userdata").Value.String()
			} else {
				userData = d.UserData
			}
		} else {
			userDataFile = cmd.Flag("userfile").Value.String()
			data, err := ioutil.ReadFile(userDataFile)
			if err != nil {
				return err
			}
			userData = string(data)
		}

		err = UpdateDevice(deviceID, hostname, description, userData, ipxeScriptURL, []string{}, locked, alwaysPXE)
		return err
	},
}

var deleteDeviceCmd = &cobra.Command{
	Use:   "delete-device",
	Short: "Delete a device",
	RunE: func(cmd *cobra.Command, args []string) error {
		deviceID := cmd.Flag("device-id").Value.String()
		err := DeleteDevice(deviceID)
		return err
	},
}

var lockDeviceCmd = &cobra.Command{
	Use:   "lock-device",
	Short: "Lock a device",
	RunE: func(cmd *cobra.Command, args []string) error {
		deviceID := cmd.Flag("device-id").Value.String()
		err := LockDevice(deviceID)
		return err
	},
}

var unlockDeviceCmd = &cobra.Command{
	Use:   "unlock-device",
	Short: "Unlock a device",
	RunE: func(cmd *cobra.Command, args []string) error {
		deviceID := cmd.Flag("device-id").Value.String()
		err := UnlockDevice(deviceID)
		return err
	},
}

var powerOnDeviceCmd = &cobra.Command{
	Use:   "poweron-device",
	Short: "Power on a device",
	RunE: func(cmd *cobra.Command, args []string) error {
		deviceID := cmd.Flag("device-id").Value.String()
		err := PowerOnDevice(deviceID)
		return err
	},
}

var powerOffDeviceCmd = &cobra.Command{
	Use:   "poweroff-device",
	Short: "Power off a device",
	RunE: func(cmd *cobra.Command, args []string) error {
		deviceID := cmd.Flag("device-id").Value.String()
		err := PowerOffDevice(deviceID)
		return err
	},
}

var rebootDeviceCmd = &cobra.Command{
	Use:   "reboot-device",
	Short: "Reboot a device",
	RunE: func(cmd *cobra.Command, args []string) error {
		deviceID := cmd.Flag("device-id").Value.String()
		err := RebootDevice(deviceID)
		return err
	},
}

var listDeviceEventsCmd = &cobra.Command{
	Use:   "list-events",
	Short: "View events by device ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		deviceID := cmd.Flag("device-id").Value.String()
		err := ListDeviceEvents(deviceID)
		return err
	},
}

func init() {
	// subcommands
	baremetalCmd.AddCommand(listDevicesCmd, listDeviceCmd, createDeviceCmd, updateDeviceCmd, deleteDeviceCmd, lockDeviceCmd, unlockDeviceCmd, powerOnDeviceCmd, powerOffDeviceCmd, rebootDeviceCmd, listDeviceEventsCmd)

	// add to root command
	RootCmd.AddCommand(baremetalCmd)

	// Flags for command: packet baremetal list-devices
	listDevicesCmd.Flags().String("project-id", "", "Specify the project ID.")

	// Flags for command: packet baremetal list-device
	listDeviceCmd.Flags().String("device-id", "", "Specify ID of device to display.")

	// Flags for command: packet baremetal create
	createDeviceCmd.Flags().String("project-id", "", "The project ID.")
	createDeviceCmd.Flags().String("hostname", "", "Hostname of the device")
	createDeviceCmd.Flags().String("plan", "baremetal_0", "Server type to create the device")
	createDeviceCmd.Flags().String("facility", "", "DC location. Use \"packet admin list-facilities\" to see available facilities")
	createDeviceCmd.Flags().String("os-type", "centos_7", "Operating system to deploy to the server")
	createDeviceCmd.Flags().String("billing", "hourly", "Choose \"hourly\" or \"monthly\" billing")
	createDeviceCmd.Flags().StringP("file", "f", "", "Read userdata from a file. This option works but is deprecated; use \"--userfile\" instead")
	createDeviceCmd.Flags().String("userfile", "", "Read userdata from a `[file]`")
	createDeviceCmd.Flags().String("userdata", "", "userdata string. This options will be disgarded if \"--userfile\" is present")
	createDeviceCmd.Flags().BoolVarP(&silent, "silent", "s", false, "Omit provisioning logs")
	createDeviceCmd.Flags().BoolVarP(&spotInstance, "spot-instance", "", false, "Create as a spot instance")
	createDeviceCmd.Flags().BoolVarP(&alwaysPXE, "always-pxe", "", false, "Set PXE boot to `true`")
	createDeviceCmd.Flags().String("ipxe-script-url", "", "Script URL")
	createDeviceCmd.Flags().Float64VarP(&spotPriceMax, "spot-price-max", "", 0.0, "Spot market price bid")

	// Flags for command: packet baremetal update-device
	updateDeviceCmd.Flags().String("device-id", "", "Device ID")
	updateDeviceCmd.Flags().String("hostname", "", "Hostname of the device")
	updateDeviceCmd.Flags().String("description", "", "Description")
	updateDeviceCmd.Flags().String("ipxe-script-url", "", "Script URL")
	updateDeviceCmd.Flags().String("userfile", "", "Read userdata from a `[file]`")
	updateDeviceCmd.Flags().String("userdata", "", "userdata string. This options will be disgarded if \"--userfile\" is present")
	updateDeviceCmd.Flags().String("always-pxe", "", "PXE boot: [true || false]")
	updateDeviceCmd.Flags().String("lock", "", "Lock device: [true || false]")

	// Flags for other device commands that require the device ID.
	deleteDeviceCmd.Flags().String("device-id", "", "Device ID")
	lockDeviceCmd.Flags().String("device-id", "", "Device ID")
	unlockDeviceCmd.Flags().String("device-id", "", "Device ID")
	powerOnDeviceCmd.Flags().String("device-id", "", "Device ID")
	powerOffDeviceCmd.Flags().String("device-id", "", "Device ID")
	rebootDeviceCmd.Flags().String("device-id", "", "Device ID")
	listDeviceEventsCmd.Flags().String("device-id", "", "Device ID")

}

package cmd

import (
	"github.com/spf13/cobra"
)

// storageCmd represents the storage command
var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Manage your storages",
}

var listStoragesCmd = &cobra.Command{
	Use:   "list-volumes",
	Short: "Retrieve all volumes",
	RunE: func(cmd *cobra.Command, args []string) error {
		projectID := GetProjectID(cmd)
		err := ListStorages(projectID)
		return err
	},
}

var createStorageCmd = &cobra.Command{
	Use:   "create-volume",
	Short: "Create a volume",
	RunE: func(cmd *cobra.Command, args []string) error {
		projectID := GetProjectID(cmd)
		description := cmd.Flag("desc").Value.String()
		plan := cmd.Flag("plan").Value.String()
		facility := cmd.Flag("facility").Value.String()
		size, err := cmd.Flags().GetInt("size")
		if err != nil {
			return err
		}
		snapshotFrequency := cmd.Flag("frequency").Value.String()
		snapshotCount, err := cmd.Flags().GetInt("count")
		if err != nil {
			return err
		}
		e := CreateStorage(projectID, description, plan, facility, snapshotFrequency, size, snapshotCount)
		return e
	},
}

var listStorageCmd = &cobra.Command{
	Use:   "list-volume",
	Short: "Retrieve a volume by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		storageID := cmd.Flag("storage-id").Value.String()
		err := ListStorage(storageID)
		return err
	},
}

var updateStorageCmd = &cobra.Command{
	Use:   "update-volume",
	Short: "Update a volume",
	RunE: func(cmd *cobra.Command, args []string) error {
		storageID := cmd.Flag("storage-id").Value.String()
		description := cmd.Flag("desc").Value.String()
		size, err := cmd.Flags().GetInt("size")
		if err != nil {
			return err
		}
		locked, err := cmd.Flags().GetBool("lock")
		if err != nil {
			return err
		}
		e := UpdateStorage(storageID, description, size, locked)
		return e
	},
}

var deleteStorageCmd = &cobra.Command{
	Use:   "delete-volume",
	Short: "Delete storage",
	RunE: func(cmd *cobra.Command, args []string) error {
		storageID := cmd.Flag("storage-id").Value.String()
		err := DeleteStorage(storageID)
		return err
	},
}

var createSnapshotPolicyCmd = &cobra.Command{
	Use:   "create-snapshot-policy",
	Short: "Create a snapshot policy",
	RunE: func(cmd *cobra.Command, args []string) error {
		storageID := cmd.Flag("storage-id").Value.String()
		snapshotFrequency := cmd.Flag("frequency").Value.String()
		snapshotCount, err := cmd.Flags().GetInt("count")
		if err != nil {
			return err
		}
		e := CreateSnapshotPolicy(storageID, snapshotFrequency, snapshotCount)
		return e
	},
}

var updateSnapshotPolicyCmd = &cobra.Command{
	Use:   "update-snapshot-policy",
	Short: "Update a snapshot policy",
	RunE: func(cmd *cobra.Command, args []string) error {
		policyID := cmd.Flag("policy-id").Value.String()
		snapshotFrequency := cmd.Flag("frequency").Value.String()
		snapshotCount, err := cmd.Flags().GetInt("count")
		if err != nil {
			return err
		}
		e := UpdateSnapshotPolicy(policyID, snapshotFrequency, snapshotCount)
		return e
	},
}

var deleteSnapshotPolicyCmd = &cobra.Command{
	Use:   "delete-snapshot-policy",
	Short: "Delete a snapshot policy",
	RunE: func(cmd *cobra.Command, args []string) error {
		policyID := cmd.Flag("policy-id").Value.String()
		err := DeleteSnapshotPolicy(policyID)
		return err
	},
}

var listSnapshotsCmd = &cobra.Command{
	Use:   "list-snapshots",
	Short: "View a list of the current volume’s snapshots",
	RunE: func(cmd *cobra.Command, args []string) error {
		policyID := cmd.Flag("policy-id").Value.String()
		err := ListSnapshots(policyID)
		return err
	},
}

var createSnapshotCmd = &cobra.Command{
	Use:   "create-snapshot",
	Short: "Create a snapshot of your volume",
	RunE: func(cmd *cobra.Command, args []string) error {
		policyID := cmd.Flag("policy-id").Value.String()
		err := CreateSnapshot(policyID)
		return err
	},
}

var deleteSnapshotCmd = &cobra.Command{
	Use:   "delete-snapshot",
	Short: "Delete a snapshot of your volume",
	RunE: func(cmd *cobra.Command, args []string) error {
		storageID := cmd.Flag("storage-id").Value.String()
		snapshotID := cmd.Flag("snapshot-id").Value.String()
		err := DeleteSnapshot(storageID, snapshotID)
		return err
	},
}

var listStorageEventsCmd = &cobra.Command{
	Use:   "list-volume-events",
	Short: "View a list of the current volume's events",
	RunE: func(cmd *cobra.Command, args []string) error {
		storageID := cmd.Flag("storage-id").Value.String()
		snapshotID := cmd.Flag("snapshot-id").Value.String()
		err := ListStorageEvents(storageID, snapshotID)
		return err
	},
}

var attachStorageCmd = &cobra.Command{
	Use:   "attach-volume",
	Short: "Attach a volume to a device",
	RunE: func(cmd *cobra.Command, args []string) error {
		storageID := cmd.Flag("storage-id").Value.String()
		snapshotID := cmd.Flag("snapshot-id").Value.String()
		deviceID := cmd.Flag("device-id").Value.String()
		err := AttachStorage(storageID, snapshotID, deviceID)
		return err
	},
}

var detachStorageCmd = &cobra.Command{
	Use:   "detach-volume",
	Short: "Detach a volume from a device",
	RunE: func(cmd *cobra.Command, args []string) error {
		attachmentID := cmd.Flag("attachement-id").Value.String()
		err := DetachStorage(attachmentID)
		return err
	},
}

func init() {
	storageCmd.AddCommand(listStoragesCmd, createStorageCmd, listStorageCmd, updateStorageCmd, deleteStorageCmd, createSnapshotPolicyCmd, updateSnapshotPolicyCmd, deleteSnapshotPolicyCmd, listSnapshotsCmd, createSnapshotCmd, deleteSnapshotCmd, listStorageEventsCmd, attachStorageCmd, detachStorageCmd)
	RootCmd.AddCommand(storageCmd)

	// Flags for command: packet storage listall
	listStoragesCmd.Flags().String("project-id", "", "Project ID")

	// Flags for command: packet storage create
	createStorageCmd.Flags().String("project-id", "", "Project ID")
	createStorageCmd.Flags().String("desc", "", "Description")
	createStorageCmd.Flags().String("plan", "storage_1", "storage_1 || storage_2")
	createStorageCmd.Flags().String("facility", "ewr1", "ewr1 || sjc1 || ams1")
	createStorageCmd.Flags().Int("size", 120, "Volume size")
	createStorageCmd.Flags().String("frequency", "15min", "Snapshot frequency")
	createStorageCmd.Flags().Int("count", 4, "Snapshots count")

	// Flags for command: packet storage list
	listStorageCmd.Flags().String("storage-id", "", "Storage ID")

	// Flags for command: packet storage update
	updateStorageCmd.Flags().String("storage-id", "", "Storage ID")
	updateStorageCmd.Flags().String("desc", "", "Description")
	updateStorageCmd.Flags().Int("size", 120, "Volume size")
	updateStorageCmd.Flags().Bool("lock", false, "Update and lock")

	// Flags for command: packet storage delete
	deleteStorageCmd.Flags().String("storage-id", "", "Storage ID")

	// Flags for command: packet storage create-snapshot-policy
	createSnapshotPolicyCmd.Flags().String("storage-id", "", "Storage ID")
	createSnapshotPolicyCmd.Flags().String("frequency", "", "Snapshot frequency")
	createSnapshotPolicyCmd.Flags().Int("count", 1, "Snapshots count")

	// Flags for command: packet storage update-snapshot-policy
	updateSnapshotPolicyCmd.Flags().String("policy-id", "", "Snapshot policy ID")
	updateSnapshotPolicyCmd.Flags().String("frequency", "15min", "Snapshot frequency")
	updateSnapshotPolicyCmd.Flags().Int("count", 4, "Snapshots count")

	// Flags for command: packet storage delete-snapshot-policy
	deleteSnapshotPolicyCmd.Flags().String("policy-id", "", "Snapshot policy ID")

	// Flags for command: packet storage list-snapshots
	listSnapshotsCmd.Flags().String("policy-id", "", "Snapshot policy ID")

	// Flags for command: packet storage create-snapshot
	createSnapshotCmd.Flags().String("policy-id", "", "Snapshot policy ID")

	// Flags for command: packet storage delete-snapshot
	deleteSnapshotCmd.Flags().String("storage-id", "", "Storage ID")
	deleteSnapshotCmd.Flags().String("snapshot-id", "", "Snapshot policy ID")

	// Flags for command: packet storage list-events
	listStorageEventsCmd.Flags().String("storage-id", "", "Storage ID")
	listStorageEventsCmd.Flags().String("snapshot-id", "", "Snapshot policy ID")

	// Flags for command: packet storage attach
	attachStorageCmd.Flags().String("storage-id", "", "Storage ID")
	attachStorageCmd.Flags().String("snapshot-id", "", "Snapshot policy ID")
	attachStorageCmd.Flags().String("device-id", "", "Device ID")

	// Flags for command: packet storage detach
	detachStorageCmd.Flags().String("attachement-id", "", "Attachment ID")
}

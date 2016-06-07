package cmd

import (
	"fmt"
	"time"

	"github.com/packethost/packngo"

	"github.com/ebsarr/packet/extpackngo"
)

// Ifs to Facility API

// ListFacilities returns a list of packet DCs
func ListFacilities() error {
	client, err := NewPacketClient()
	if err != nil {
		return err
	}

	facilities, _, err := client.Facilities.List()
	if err != nil {
		return err
	}

	e := MarshallAndPrint(facilities)
	return e
}

// IFs to Projects API

// ListProjects prints out all projects of the user.
func ListProjects() error {
	client, err := NewPacketClient()
	if err != nil {
		return err
	}

	projects, _, err := client.Projects.List()
	if err != nil {
		return err
	}

	e := MarshallAndPrint(projects)
	return e
}

// ListProject prints out the project associated with a given project id
func ListProject(projectID string) error {
	client, err := NewPacketClient()
	if err != nil {
		return err
	}

	p, _, err := client.Projects.Get(projectID)
	if err != nil {
		return err
	}

	e := MarshallAndPrint(p)
	return e
}

// CreateProject creates a new project with the given project name
func CreateProject(name, paymentID string) error {
	client, err := NewPacketClient()
	if err != nil {
		return err
	}

	req := packngo.ProjectCreateRequest{
		Name:          name,
		PaymentMethod: paymentID,
	}

	p, _, err := client.Projects.Create(&req)
	if err != nil {
		return err
	}

	e := MarshallAndPrint(p)
	return e
}

// DeleteProject deletes the project associated with the given project id.
func DeleteProject(id string) error {
	client, err := NewPacketClient()
	if err != nil {
		return err
	}

	_, e := client.Projects.Delete(id)
	return e
}

// UpdateProject updates the project associated with the given project id either
// by changing the name or the payment method.
func UpdateProject(id, name, paymentID string) error {
	client, err := NewPacketClient()
	if err != nil {
		return err
	}

	req := packngo.ProjectUpdateRequest{
		ID:            id,
		Name:          name,
		PaymentMethod: paymentID,
	}

	p, _, err := client.Projects.Update(&req)
	if err != nil {
		return err
	}

	e := MarshallAndPrint(p)
	return e
}

// IFs to Device, Devices API

// ListDevices prints out all devices associated with the given project id.
func ListDevices(projectID string) error {
	client, err := NewPacketClient()
	if err != nil {
		return err
	}

	d, _, err := client.Devices.List(projectID)
	if err != nil {
		return err
	}

	e := MarshallAndPrint(d)
	return e
}

// ListDevice prints out the device associated with the given device id.
func ListDevice(deviceID string) error {
	client, err := NewPacketClient()
	if err != nil {
		return err
	}

	d, _, err := client.Devices.Get(deviceID)
	if err != nil {
		return err
	}

	e := MarshallAndPrint(d)
	return e
}

// CreateDevice creates a new device
func CreateDevice(projectID, hostname, plan, facility, operatingSystem, billingCycle string, tags []string) error {
	client, err := NewPacketClient()
	if err != nil {
		return err
	}

	req := packngo.DeviceCreateRequest{
		HostName:     hostname,
		Plan:         plan,
		Facility:     facility,
		OS:           operatingSystem,
		BillingCycle: billingCycle,
		ProjectID:    projectID,
		UserData:     "",
		Tags:         tags,
	}

	d, _, err := client.Devices.Create(&req)
	if err != nil {
		return err
	}

	e := MarshallAndPrint(d)
	return e
}

// CreateDeviceVerbose creates a new device and logs events till the device is provisionned
func CreateDeviceVerbose(projectID, hostname, plan, facility, operatingSystem, billingCycle string, tags []string) error {
	client, err := NewPacketClient()
	if err != nil {
		return err
	}

	req := packngo.DeviceCreateRequest{
		HostName:     hostname,
		Plan:         plan,
		Facility:     facility,
		OS:           operatingSystem,
		BillingCycle: billingCycle,
		ProjectID:    projectID,
		UserData:     "",
		Tags:         tags,
	}

	d, _, err := client.Devices.Create(&req)
	if err != nil {
		return err
	}

	// print events till device is provisionned
	finalEvent := "Provision complete! Your device is ready to go."
	lastEvent := ""

	extclient, err := NewExtPacketClient()
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Provisioning of device successfully started...")

	for {
		events, _, err := extclient.Events.ListDeviceEvents(d.ID)
		if err != nil {
			return err
		}

		currentEventO := events[0]

		if currentEventO.Body != lastEvent {
			fmt.Printf(" [ %s ] %s\n", currentEventO.Create, currentEventO.Body)
			lastEvent = currentEventO.Body
		}

		if currentEventO.Body == finalEvent {
			fmt.Println()
			break
		}

		time.Sleep(10 * time.Second)
	}

	return ListDevice(d.ID)
}

// DeleteDevice deletes the device associated with the given device id.
func DeleteDevice(deviceID string) error {
	client, err := NewPacketClient()
	if err != nil {
		return err
	}

	_, e := client.Devices.Delete(deviceID)
	return e
}

// LockDevice locks the device associated with the given device id.
func LockDevice(deviceID string) error {
	client, err := NewPacketClient()
	if err != nil {
		return err
	}

	_, e := client.Devices.Lock(deviceID)
	return e
}

// UnlockDevice unlocks the device associated with the given device id.
func UnlockDevice(deviceID string) error {
	client, err := NewPacketClient()
	if err != nil {
		return err
	}

	_, e := client.Devices.Unlock(deviceID)
	return e
}

// PowerOnDevice powers on the device associated with the given device id.
func PowerOnDevice(deviceID string) error {
	client, err := NewPacketClient()
	if err != nil {
		return err
	}

	_, e := client.Devices.PowerOn(deviceID)
	return e
}

// PowerOffDevice powers off the device associated with the given device id.
func PowerOffDevice(deviceID string) error {
	client, err := NewPacketClient()
	if err != nil {
		return err
	}

	_, e := client.Devices.PowerOff(deviceID)
	return e
}

// RebootDevice reboots the device associated with the given device id.
func RebootDevice(deviceID string) error {
	client, err := NewPacketClient()
	if err != nil {
		return err
	}

	_, e := client.Devices.Reboot(deviceID)
	return e
}

// IFs to Plan API

// ListPlans prints out the available plans(server types).
func ListPlans() error {
	client, err := NewPacketClient()
	if err != nil {
		return err
	}

	p, _, err := client.Plans.List()
	if err != nil {
		return err
	}

	e := MarshallAndPrint(p)
	return e
}

// IFs to OS API

// ListOS prints out the available operating systems.
func ListOS() error {
	client, err := NewPacketClient()
	if err != nil {
		return err
	}

	o, _, err := client.OperatingSystems.List()
	if err != nil {
		return err
	}

	e := MarshallAndPrint(o)
	return e
}

// IFs to SSH API

// ListSSHKeys prints out all ssh keys generated by the user.
func ListSSHKeys() error {
	client, err := NewPacketClient()
	if err != nil {
		return err
	}

	k, _, err := client.SSHKeys.List()
	if err != nil {
		return err
	}

	e := MarshallAndPrint(k)
	return e
}

// ListSSHKey prints out the ssh key associated with the given key id.
func ListSSHKey(keyID string) error {
	client, err := NewPacketClient()
	if err != nil {
		return err
	}

	k, _, err := client.SSHKeys.Get(keyID)
	if err != nil {
		return err
	}

	e := MarshallAndPrint(k)
	return e
}

// CreateSSHKey creates a new ssh key.
func CreateSSHKey(label, key string) error {
	client, err := NewPacketClient()
	if err != nil {
		return err
	}

	req := packngo.SSHKeyCreateRequest{
		Key:   key,
		Label: label,
	}

	k, _, err := client.SSHKeys.Create(&req)
	if err != nil {
		return err
	}

	e := MarshallAndPrint(k)
	return e
}

// DeleteSSHKey deletes the ssh key associated with the given key id.
func DeleteSSHKey(keyID string) error {
	client, err := NewPacketClient()
	if err != nil {
		return err
	}

	_, e := client.SSHKeys.Delete(keyID)
	return e
}

// UpdateSSHKey updates the ssh key associated with the given key id.
func UpdateSSHKey(keyID, label, key string) error {
	client, err := NewPacketClient()
	if err != nil {
		return err
	}

	req := packngo.SSHKeyUpdateRequest{
		ID:    keyID,
		Label: label,
		Key:   key,
	}

	k, _, err := client.SSHKeys.Update(&req)
	if err != nil {
		return err
	}

	e := MarshallAndPrint(k)
	return e
}

// IFs to Event API

// ListDeviceEvents prints out events by device ID
func ListDeviceEvents(id string) error {
	client, err := NewExtPacketClient()
	if err != nil {
		return err
	}

	events, _, err := client.Events.ListDeviceEvents(id)
	if err != nil {
		return err
	}

	e := MarshallAndPrint(events)
	return e
}

// ListProjectEvents prints out events by device ID
func ListProjectEvents(id string) error {
	client, err := NewExtPacketClient()
	if err != nil {
		return err
	}

	events, _, err := client.Events.ListProjectEvents(id)
	if err != nil {
		return err
	}

	e := MarshallAndPrint(events)
	return e
}

// ListStorageEvents prints out events by device ID
func ListStorageEvents(storageID, snapshotID string) error {
	client, err := NewExtPacketClient()
	if err != nil {
		return err
	}

	events, _, err := client.Events.ListStorageEvents(storageID, snapshotID)
	if err != nil {
		return err
	}

	e := MarshallAndPrint(events)
	return e
}

// Extention of the Device API to assign IP address

// ListIPAddress prints out ip address by ID
func ListIPAddress(ipAddressID string) error {
	client, err := NewExtPacketClient()
	if err != nil {
		return err
	}

	ip, _, err := client.IPs.Get(ipAddressID)
	if err != nil {
		return err
	}

	e := MarshallAndPrint(ip)
	return e
}

// AssignIPAddress assigns an IP address to a device by ID
func AssignIPAddress(deviceID, ipAddress string) error {
	client, err := NewExtPacketClient()
	if err != nil {
		return err
	}

	req := extpackngo.IPAddressAssignRequest{
		Address: ipAddress,
	}

	ip, _, err := client.IPs.Assign(deviceID, &req)
	if err != nil {
		return err
	}

	e := MarshallAndPrint(ip)
	return e
}

// UnAssignIPAddress unassigns and IP address from a device
func UnAssignIPAddress(ipAddressID string) error {
	client, err := NewExtPacketClient()
	if err != nil {
		return err
	}

	_, e := client.IPs.Unassign(ipAddressID)
	return e
}

// ListIPReservations provides a list of IP resevations for a single project
func ListIPReservations(projectID string) error {
	client, err := NewExtPacketClient()
	if err != nil {
		return err
	}

	reservations, _, err := client.IPReservations.List(projectID)
	if err != nil {
		return err
	}

	e := MarshallAndPrint(reservations)
	return e
}

// RequestMoreIPReservations requests more IP space for a project in order to have additional IP addresses to assign to devices
func RequestMoreIPReservations(projectID, ipType, comments string, quantity int) error {
	client, err := NewExtPacketClient()
	if err != nil {
		return err
	}

	req := extpackngo.IPReservationRequest{
		Type:     ipType,
		Quantity: quantity,
		Comments: comments,
	}

	_, e := client.IPReservations.RequestMore(projectID, &req)
	return e
}

// ListIPReservation returns a single IP reservation object
func ListIPReservation(id string) error {
	client, err := NewExtPacketClient()
	if err != nil {
		return err
	}

	reservation, _, err := client.IPReservations.Get(id)
	if err != nil {
		return err
	}

	e := MarshallAndPrint(reservation)
	return e
}

// RemoveIPReservation removes an IP reservation from the project
func RemoveIPReservation(id string) error {
	client, err := NewExtPacketClient()
	if err != nil {
		return err
	}

	_, e := client.IPReservations.Remove(id)
	return e
}

// Ifs to VOLUMES API

// ListStorages returns a list of the current projects’s volumes
func ListStorages(projectID string) error {
	client, err := NewExtPacketClient()
	if err != nil {
		return err
	}

	storages, _, err := client.Storages.List(projectID)
	if err != nil {
		return err
	}

	e := MarshallAndPrint(storages)
	return e
}

// CreateStorage creates a new volume
func CreateStorage(projectID, description, plan, facility, frequency string, size, count int) error {
	client, err := NewExtPacketClient()
	if err != nil {
		return err
	}

	// Create a snapshot policy
	sp := &extpackngo.SnapshotPolicy{
		SnapshotFrequency:	frequency,
		SnapshotCount:	count,
	}
	request := &extpackngo.StorageCreateRequest{
		Description:      description,
		Plan:             plan,
		Size:             size,
		Facility:         facility,
		SnapshotPolicies: sp,
	}

	storage, _, err := client.Storages.Create(projectID, request)
	if err != nil {
		return err
	}

	e := MarshallAndPrint(storage)
	return e
}

// ListStorage returns a volume by ID
func ListStorage(storageID string) error {
	client, err := NewExtPacketClient()
	if err != nil {
		return err
	}

	storage, _, err := client.Storages.Get(storageID)
	if err != nil {
		return err
	}

	e := MarshallAndPrint(storage)
	return e
}

// UpdateStorage updates a volume
func UpdateStorage(storageID, description string, size int, locked bool) error {
	client, err := NewExtPacketClient()
	if err != nil {
		return err
	}

	request := &extpackngo.StorageUpdateRequest{
		Description: description,
		Size:        size,
		Locked:      locked,
	}

	_, e := client.Storages.Update(storageID, request)
	return e
}

// DeleteStorage deletes a volume
func DeleteStorage(storageID string) error {
	client, err := NewExtPacketClient()
	if err != nil {
		return err
	}

	_, e := client.Storages.Delete(storageID)
	return e
}

// CreateSnapshotPolicy creates a snapshot policy
func CreateSnapshotPolicy(storageID, frequency string, count int) error {
	client, err := NewExtPacketClient()
	if err != nil {
		return err
	}

	request := &extpackngo.CreateSnapshotPolicyRequest{
		SnapshotCount:     count,
		SnapshotFrequency: frequency,
	}

	_, e := client.Storages.CreateSnapshotPolicy(storageID, request)
	return e
}

// UpdateSnapshotPolicy updates a snapshot policy
func UpdateSnapshotPolicy(snapshotPolicyID, frequency string, count int) error {
	client, err := NewExtPacketClient()
	if err != nil {
		return err
	}

	request := &extpackngo.UpdateSnapshotPolicyRequest{
		SnapshotCount:     count,
		SnapshotFrequency: frequency,
	}

	_, e := client.Storages.UpdateSnapshotPolicy(snapshotPolicyID, request)
	return e
}

// DeleteSnapshotPolicy deletes a snapshot policy
func DeleteSnapshotPolicy(snapshotPolicyID string) error {
	client, err := NewExtPacketClient()
	if err != nil {
		return err
	}

	_, e := client.Storages.DeleteSnapshotPolicy(snapshotPolicyID)
	return e
}

// ListSnapshots returns a list of the current volume’s snapshots
func ListSnapshots(snapshotPolicyID string) error {
	client, err := NewExtPacketClient()
	if err != nil {
		return err
	}

	snapshots, _, err := client.Storages.ListSnapshots(snapshotPolicyID)
	if err != nil {
		return err
	}

	e := MarshallAndPrint(snapshots)
	return e
}

// CreateSnapshot creates a new snapshot of your volume
func CreateSnapshot(snapshotPolicyID string) error {
	client, err := NewExtPacketClient()
	if err != nil {
		return err
	}

	request := &extpackngo.CreateSnapShotRequest{}

	_, e := client.Storages.CreateSnapshot(snapshotPolicyID, request)
	return e
}

// DeleteSnapshot deletes a snapshot
func DeleteSnapshot(storageID, snapshotID string) error {
	client, err := NewExtPacketClient()
	if err != nil {
		return err
	}

	_, e := client.Storages.DeleteSnapshot(storageID, snapshotID)
	return e
}

// AttachStorage attaches your volume to a device
func AttachStorage(storageID, snapshotID, deviceID string) error {
	client, err := NewExtPacketClient()
	if err != nil {
		return err
	}

	request := &extpackngo.AttachStorageRequest{
		DeviceID: deviceID,
	}

	_, e := client.Storages.Attach(storageID, snapshotID, request)
	return e
}

// DetachStorage detaches your volume from a device
func DetachStorage(attachmentID string) error {
	client, err := NewExtPacketClient()
	if err != nil {
		return err
	}

	_, e := client.Storages.Detach(attachmentID)
	return e
}

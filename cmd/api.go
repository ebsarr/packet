package cmd

import (
	"github.com/packethost/packngo"
)

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
func ListProject(id string) error {
	client, err := NewPacketClient()
	if err != nil {
		return err
	}

	p, _, err := client.Projects.Get(id)
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

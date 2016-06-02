package extpackngo

import (
	"fmt"

	"github.com/packethost/packngo"
)

// IPBasePath ...
const IPBasePath = "/ips"

// IPService interface defines available event methods
type IPService interface {
	Assign(deviceID string, assignRequest *IPAddressAssignRequest) (*IPAddress, *Response, error)
	Unassign(ipAddressID string) (*Response, error)
	Get(ipAddressID string) (*IPAddress, *Response, error)
}

// IPAddress represents a ip address
type IPAddress struct {
	ID            string            `json:"id"`
	Address       string            `json:"address"`
	Network       string            `json:"network"`
	AddressFamily int               `json:"address_family"`
	Netmask       string            `json:"netmask"`
	Public        bool              `json:"public"`
	Cidr          int               `json:"cidr"`
	AssignedTo    map[string]string `json:"assigned_to"`
	Href          string            `json:"href"`
}

// IPAddressAssignRequest represents the body if a ip assign request
type IPAddressAssignRequest struct {
	Address string `json:"address"`
}

func (i IPAddress) String() string {
	return packngo.Stringify(i)
}

// IPServiceOp implements IPService
type IPServiceOp struct {
	client *Client
}

// Get returns IpAddress by ID
func (i *IPServiceOp) Get(ipAddressID string) (*IPAddress, *Response, error) {
	path := fmt.Sprintf("%s/%s", IPBasePath, ipAddressID)

	req, err := i.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	ip := new(IPAddress)
	resp, err := i.client.Do(req, ip)
	if err != nil {
		return nil, resp, err
	}

	return ip, resp, err
}

// Unassign unassigns an IP address record. This will remove the relationship between an IP
// and the device and will make the IP address available to be assigned to another device.
func (i *IPServiceOp) Unassign(ipAddressID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", IPBasePath, ipAddressID)

	req, err := i.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := i.client.Do(req, nil)
	return resp, err
}

// Assign assigns an IP address to a device. The IP address must be in one of the IP ranges assigned to the device’s project.
func (i *IPServiceOp) Assign(deviceID string, assignRequest *IPAddressAssignRequest) (*IPAddress, *Response, error) {
	path := fmt.Sprintf("devices/%s%s", deviceID, IPBasePath)

	req, err := i.client.NewRequest("POST", path, assignRequest)

	ip := new(IPAddress)
	resp, err := i.client.Do(req, ip)
	if err != nil {
		return nil, resp, err
	}

	return ip, resp, err
}

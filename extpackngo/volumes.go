package extpackngo

import (
	"fmt"

	"github.com/packethost/packngo"
)

const storageBasePath = "/storage"

// StorageService interface defines available storage methods
type StorageService interface {
	List(projectID string) ([]Storage, *Response, error)
	Create(projectID string, request *StorageCreateRequest) (*Storage, *Response, error)
	Get(storageID string) (*Storage, *Response, error)
	Update(storageID string, request *StorageUpdateRequest) (*Response, error)
	Delete(storageID string) (*Response, error)
	CreateSnapshotPolicy(storageID string, request *CreateSnapshotPolicyRequest) (*Response, error)
	UpdateSnapshotPolicy(snapshotPolicyID string, request *UpdateSnapshotPolicyRequest) (*Response, error)
	DeleteSnapshotPolicy(snapshotPolicyID string) (*Response, error)
	ListSnapshots(snapshotPolicyID string) ([]Snapshot, *Response, error)
	CreateSnapshot(snapshotPolicyID string, request *CreateSnapShotRequest) (*Response, error)
	DeleteSnapshot(storageID, snapshotID string) (*Response, error)
	Attach(storageID, snapshotID string, request *AttachStorageRequest) (*Response, error)
	Detach(attachmentID string) (*Response, error)
}

// StorageServiceOP implements the StorageService interface
type StorageServiceOP struct {
	client *Client
}

// Storage represents a packet block storage
type Storage struct {
	ID               string              `json:"id"`
	Name             string              `json:"name"`
	Description      string              `json:"description"`
	Size             int                 `json:"size"`
	Locked           bool                `json:"locked"`
	BillingCycle     string              `json:"billing_cycle"`
	State            string              `json:"state"`
	Create           string              `json:"created_at"`
	Update           string              `json:"updated_at"`
	Project          map[string]string   `json:"project"`
	Facility         map[string]string   `json:"facility"`
	SnapshotPolicies []map[string]string `json:"snapshot_policies"`
	Attachements     []map[string]string `json:"attachments"`
	Plan             packngo.Plan        `json:"plan"`
	Href             string              `json:"href"`
}

// StorageCreateRequest represents the body of a storage create request
type StorageCreateRequest struct {
	Description      string           `json:"description"`
	Plan             string           `json:"plan"`
	Size             int              `json:"size"`
	Facility         string           `json:"facility"`
	SnapshotPolicies *SnapshotPolicy   `json:"snapshot_policies"`
}

// StorageUpdateRequest represents the body of a storage update request
type StorageUpdateRequest struct {
	Description string `json:"description"`
	Size        int    `json:"size"`
	Locked      bool   `json:"size"`
}

// SnapshotPolicy represents a snapshot policy
type SnapshotPolicy struct {
	SnapshotCount     int    `json:"snapshot_count"`
	SnapshotFrequency string `json:"snapshot_frequency"`
}

// CreateSnapshotPolicyRequest represents the body of a create snapshot policy request
type CreateSnapshotPolicyRequest struct {
	SnapshotCount     int    `json:"snapshot_count"`
	SnapshotFrequency string `json:"snapshot_frequency"`
}

// UpdateSnapshotPolicyRequest represents the body of a update snapshot policy request
type UpdateSnapshotPolicyRequest struct {
	SnapshotCount     int    `json:"snapshot_count"`
	SnapshotFrequency string `json:"snapshot_frequency"`
}

// Snapshot reprensents a snapshot
type Snapshot struct {
	ID     string            `json:"id"`
	Status string            `json:"status"`
	Create string            `json:"created_at"`
	Volume map[string]string `json:"volume"`
}

type snapshotsRoot struct {
	Snapshots []Snapshot `json:"snapshots"`
}

// CreateSnapShotRequest represents the body of a snapshot create request
type CreateSnapShotRequest struct {
}

// AttachStorageRequest represents the body of a attach storage request
type AttachStorageRequest struct {
	DeviceID string `json:"device_id"`
}

type volumesRoot struct {
	Volumes []Storage `json:"volumes"`
}

// List returns a list of the current projects’s volumes
func (s *StorageServiceOP) List(projectID string) ([]Storage, *Response, error) {
	path := fmt.Sprintf("projects/%s%s", projectID, storageBasePath)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(volumesRoot)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Volumes, resp, err
}

// Create creates a new volume
func (s *StorageServiceOP) Create(projectID string, request *StorageCreateRequest) (*Storage, *Response, error) {
	path := fmt.Sprintf("projects/%s%s", projectID, storageBasePath)

	req, err := s.client.NewRequest("POST", path, request)
	if err != nil {
		return nil, nil, err
	}

	storage := new(Storage)
	resp, err := s.client.Do(req, storage)
	if err != nil {
		return nil, nil, err
	}

	return storage, resp, err
}

// Get returns a volume by ID
func (s *StorageServiceOP) Get(storageID string) (*Storage, *Response, error) {
	path := fmt.Sprintf("%s/%s", storageBasePath, storageID)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	storage := new(Storage)
	resp, err := s.client.Do(req, storage)
	if err != nil {
		return nil, nil, err
	}

	return storage, resp, err
}

// Update updates a volume
func (s *StorageServiceOP) Update(storateID string, request *StorageUpdateRequest) (*Response, error) {
	path := fmt.Sprintf("%s/%s", storageBasePath, storateID)

	req, err := s.client.NewRequest("POST", path, request)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Delete deletes a volume
func (s *StorageServiceOP) Delete(storageID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", storageBasePath, storageID)

	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// CreateSnapshotPolicy creates a snapshot policy
func (s *StorageServiceOP) CreateSnapshotPolicy(storageID string, request *CreateSnapshotPolicyRequest) (*Response, error) {
	path := fmt.Sprintf("%s/%s", storageBasePath, storageID)

	req, err := s.client.NewRequest("POST", path, request)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// UpdateSnapshotPolicy updates a snapshot policy
func (s *StorageServiceOP) UpdateSnapshotPolicy(snapshotPolicyID string, request *UpdateSnapshotPolicyRequest) (*Response, error) {
	path := fmt.Sprintf("%s/snapshot-policies/%s", storageBasePath, snapshotPolicyID)

	req, err := s.client.NewRequest("POST", path, request)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// DeleteSnapshotPolicy deletes a snapshot policy
func (s *StorageServiceOP) DeleteSnapshotPolicy(snapshotPolicyID string) (*Response, error) {
	path := fmt.Sprintf("%s/snapshot-policies/%s", storageBasePath, snapshotPolicyID)

	req, err := s.client.NewRequest("POST", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// ListSnapshots returns a list of the current volume’s snapshots
func (s *StorageServiceOP) ListSnapshots(snapshotPolicyID string) ([]Snapshot, *Response, error) {
	path := fmt.Sprintf("%s/snapshot-policies/%s", storageBasePath, snapshotPolicyID)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(snapshotsRoot)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, nil, err
	}

	return root.Snapshots, resp, err
}

// CreateSnapshot creates a new snapshot of your volume
func (s *StorageServiceOP) CreateSnapshot(snapshotPolicyID string, request *CreateSnapShotRequest) (*Response, error) {
	path := fmt.Sprintf("%s/snapshot-policies/%s", storageBasePath, snapshotPolicyID)

	req, err := s.client.NewRequest("POST", path, request)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// DeleteSnapshot deletes a snapshot
func (s *StorageServiceOP) DeleteSnapshot(storageID, snapshotID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s/snapshots/%s", storageBasePath, storageID, snapshotID)

	req, err := s.client.NewRequest("POST", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Attach attaches your volume to a device
func (s *StorageServiceOP) Attach(storageID, snapshotID string, request *AttachStorageRequest) (*Response, error) {
	path := fmt.Sprintf("%s/%s/snapshots/%s", storageBasePath, storageID, snapshotID)

	req, err := s.client.NewRequest("POST", path, request)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Detach detaches your volume from a device
func (s *StorageServiceOP) Detach(attachmentID string) (*Response, error) {
	path := fmt.Sprintf("%s/attachments/%s", storageBasePath, attachmentID)

	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, err
}

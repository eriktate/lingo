package lingo

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// DiskClient implements the Disker interface and provides all of the
// functionality for managing disks on Linodes.
type DiskClient struct {
	api APIClient
}

// NewDiskClient returns a new DiskClient given an APIClient.
func NewDiskClient(api APIClient) DiskClient {
	return DiskClient{api: api}
}

// ListDisks retrieves all of the Disks associatd with the given Linode ID.
func (c DiskClient) ListDisks(linodeID uint) ([]Disk, error) {
	data, err := c.api.Get(fmt.Sprintf("linode/instances/%d/disks", linodeID))
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request for ListDisks")
	}

	var results Results
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal ListDisks response")
	}

	var disks []Disk
	if err := json.Unmarshal(results.Data, &disks); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal ListDisks data")
	}

	return disks, nil
}

// ViewDisk retrieves a single Disk associated with the given Linode ID and Disk ID.
func (c DiskClient) ViewDisk(linodeID, diskID uint) (Disk, error) {
	var disk Disk
	data, err := c.api.Get(fmt.Sprintf("linode/instances/%d/disks/%d", linodeID, diskID))
	if err != nil {
		return disk, errors.Wrap(err, "failed to make request for ViewDisk")
	}

	if err := json.Unmarshal(data, &disk); err != nil {
		return disk, errors.Wrap(err, "failed to unmarshal ViewDisk data")
	}

	return disk, nil
}

// CreateDisk creates a new disk attached to a Linode.
func (c DiskClient) CreateDisk(req CreateDiskRequest) (Disk, error) {
	var disk Disk
	payload, err := json.Marshal(req)
	if err != nil {
		return disk, errors.Wrap(err, "failed to marshal request for CreateDisk")
	}

	data, err := c.api.Post(fmt.Sprintf("linode/instances/%d/disks", req.LinodeID), payload)
	if err != nil {
		return disk, errors.Wrap(err, "failed to make request for CreateDisk")
	}

	if err := json.Unmarshal(data, &disk); err != nil {
		return disk, errors.Wrap(err, "failed to unmarshal CreateDisk response")
	}

	return disk, nil
}

// UpdateDisk updates an existing disk attached to a Linode.
func (c DiskClient) UpdateDisk(req UpdateDiskRequest) (Disk, error) {
	var disk Disk
	payload, err := json.Marshal(req)
	if err != nil {
		return disk, errors.Wrap(err, "failed to marshal request for UpdateDisk")
	}

	data, err := c.api.Put(fmt.Sprintf("linode/instances/%d/disks", req.LinodeID), payload)
	if err != nil {
		return disk, errors.Wrap(err, "failed to make request for UpdateDisk")
	}

	if err := json.Unmarshal(data, &disk); err != nil {
		return disk, errors.Wrap(err, "failed to unmarshal UpdateDisk response")
	}

	return disk, nil
}

// DeleteDisk deletes a specific Disk from a specific Linode.
func (c DiskClient) DeleteDisk(linodeID, diskID uint) error {
	_, err := c.api.Delete(fmt.Sprintf("linode/instances/%d/disks/%d", linodeID, diskID))
	if err != nil {
		return errors.Wrap(err, "failed to make request for DeleteDisk")
	}

	return nil
}

// ResetDiskRootPassword resets the root password on the specified Disk.
func (c DiskClient) ResetDiskRootPassword(req UpdateDiskRequest) (Disk, error) {
	var disk Disk
	payload, err := json.Marshal(req)
	if err != nil {
		return disk, errors.Wrap(err, "failed to marshal request for ResetDiskRootPassword")
	}

	data, err := c.api.Post(fmt.Sprintf("linode/instances/%d/disks/%d/password", req.LinodeID, req.ID), payload)
	if err != nil {
		return disk, errors.Wrap(err, "failed to make request for ResetDiskRootPassword")
	}

	if err := json.Unmarshal(data, &disk); err != nil {
		return disk, errors.Wrap(err, "failed to unmarshal ResetDiskRootPassword response")
	}

	return disk, nil
}

// ResizeDisk resizes the specified Disk to the given size in MB.
func (c DiskClient) ResizeDisk(linodeID, diskID, size uint) (Disk, error) {
	var disk Disk
	req := struct {
		Size uint `json:"size"`
	}{size}

	payload, err := json.Marshal(req)
	if err != nil {
		return disk, errors.Wrap(err, "failed to marshal request for ResizeDisk")
	}

	data, err := c.api.Post(fmt.Sprintf("linode/instances/%d/disks/%d/resize", linodeID, diskID), payload)
	if err != nil {
		return disk, errors.Wrap(err, "failed to make request for ResizeDisk")
	}

	if err := json.Unmarshal(data, &disk); err != nil {
		return disk, errors.Wrap(err, "failed to unmarshal ResizeDisk response")
	}

	return disk, nil
}

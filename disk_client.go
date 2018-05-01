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

// CreateDisk creates a new disk attached to a Linode.
func (c DiskClient) CreateDisk(req NewDisk) (Disk, error) {
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

// GetDisks retrieves all of the Disks associatd with the given Linode ID.
func (c DiskClient) GetDisks(linodeID uint) ([]Disk, error) {
	data, err := c.api.Get(fmt.Sprintf("linode/instances/%d/disks", linodeID))
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request for GetDisks")
	}

	var results Results
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal GetDisks response")
	}

	var disks []Disk
	if err := json.Unmarshal(results.Data, &disks); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal GetDisks data")
	}

	return disks, nil
}

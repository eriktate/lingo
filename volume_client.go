package lingo

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type VolumeClient struct {
	api APIClient
}

func NewVolumeClient(api APIClient) VolumeClient {
	return VolumeClient{api: api}
}

func (c VolumeClient) ListVolumes() ([]Volume, error) {
	data, err := c.api.Get("volumes")
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request for ListVolumes")
	}

	var results Results
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, errors.Wrap(err, "failed to decode ListVolumes response")
	}

	// TODO: Do something with paging here?
	var volumes []Volume
	if err := json.Unmarshal(results.Data, &volumes); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal ListVolumes data")
	}

	return volumes, nil
}

// ViewVolume retrieves a slice of machine volume available in Linode.
func (c VolumeClient) ViewVolume(id uint) (Volume, error) {
	var volume Volume
	data, err := c.api.Get(fmt.Sprintf("volumes/%d", id))
	if err != nil {
		return volume, errors.Wrap(err, "failed to make request for ViewVolume")
	}

	if err := json.Unmarshal(data, &volume); err != nil {
		return volume, errors.Wrap(err, "failed to unmarshal ViewVolume data")
	}

	return volume, nil
}

// CreateVolume creates a new volume.
func (c VolumeClient) CreateVolume(req CreateVolumeRequest) (Volume, error) {
	var volume Volume
	payload, err := json.Marshal(req)
	if err != nil {
		return volume, errors.Wrap(err, "failed to marshal request for CreateVolume")
	}

	data, err := c.api.Post("volumes", payload)
	if err != nil {
		return volume, errors.Wrap(err, "failed to make request for CreateVolume")
	}

	if err := json.Unmarshal(data, &volume); err != nil {
		return volume, errors.Wrap(err, "failed to decode CreateVolume response")
	}

	return volume, nil
}

// UpdateVolume updates an existing machine volume.
func (c VolumeClient) UpdateVolume(req UpdateVolumeRequest) (Volume, error) {
	var volume Volume
	payload, err := json.Marshal(req)
	if err != nil {
		return volume, errors.Wrap(err, "failed to marshal request for UpdateVolume")
	}

	data, err := c.api.Put(fmt.Sprintf("volumes/%d", req.ID), payload)
	if err != nil {
		return volume, errors.Wrap(err, "failed to make request for UpdateVolume")
	}

	if err := json.Unmarshal(data, &volume); err != nil {
		return volume, errors.Wrap(err, "failed to decode UpdateVolume response")
	}

	return volume, nil
}

// DeleteVolume retrieves a slice of machine volumes available in Linode.
func (c VolumeClient) DeleteVolume(id uint) error {
	if _, err := c.api.Delete(fmt.Sprintf("volumes/%d", id)); err != nil {
		return errors.Wrap(err, "failed to make request for DeleteVolume")
	}

	return nil
}

// AttachVolume attaches an existing volume to a Linode instance.
func (c VolumeClient) AttachVolume(req AttachVolumeRequest) error {
	if _, err := c.api.Post(fmt.Sprintf("volumes/%d/attach", req.ID), nil); err != nil {
		return errors.Wrap(err, "failed to make request for AttachVolume")
	}

	return nil
}

// CloneVolume clones an existing volume.
func (c VolumeClient) CloneVolume(req UpdateVolumeRequest) error {
	if _, err := c.api.Post(fmt.Sprintf("volumes/%d/clone", req.ID), nil); err != nil {
		return errors.Wrap(err, "failed to make request for CloneVolume")
	}

	return nil
}

// DetatchVolume detatches an existing volume from it's Linode instance.
func (c VolumeClient) DetatchVolume(id uint) error {
	if _, err := c.api.Post(fmt.Sprintf("volumes/%d/detatch", id), nil); err != nil {
		return errors.Wrap(err, "failed to make request for DetatchVolume")
	}

	return nil
}

// ResizeVolume resizes an existing volume to the size represented in GB.
func (c VolumeClient) ResizeVolume(id, size uint) error {
	req := struct {
		Size uint `json:"size"`
	}{size}

	payload, err := json.Marshal(&req)
	if err != nil {
		return errors.Wrap(err, "failed to marshal request for ResizeVolume")
	}

	if _, err := c.api.Post(fmt.Sprintf("volumes/%d/resize", id), payload); err != nil {
		return errors.Wrap(err, "failed to make request for DetatchVolume")
	}

	return nil
}

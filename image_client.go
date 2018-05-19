package lingo

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// An ImageClient is an API struct that can make requests related to Linode Images.
type ImageClient struct {
	api APIClient
}

// NewImageClient returns a new ImageClient given a valid APIClient.
func NewImageClient(api APIClient) ImageClient {
	return ImageClient{api: api}
}

// ListImages retrieves a slice of machine images available in Linode.
func (c ImageClient) ListImages() ([]Image, error) {
	data, err := c.api.Get("images")
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request for ListImages")
	}

	var results Results
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, errors.Wrap(err, "failed to decode ListImages response")
	}

	// TODO: Do something with paging here?
	var images []Image
	if err := json.Unmarshal(results.Data, &images); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal ListImages data")
	}

	return images, nil
}

// ViewImage retrieves a slice of machine images available in Linode.
func (c ImageClient) ViewImage(id string) (Image, error) {
	var image Image
	data, err := c.api.Get("images/" + id)
	if err != nil {
		return image, errors.Wrap(err, "failed to make request for ViewImage")
	}

	if err := json.Unmarshal(data, &image); err != nil {
		return image, errors.Wrap(err, "failed to unmarshal ViewImage data")
	}

	return image, nil
}

// CreateImage creates a new machine image from an existing Linode disk.
func (c ImageClient) CreateImage(req CreateImageRequest) (Image, error) {
	var image Image
	payload, err := json.Marshal(req)
	if err != nil {
		return image, errors.Wrap(err, "failed to marshal request for CreateImage")
	}

	data, err := c.api.Post("images", payload)
	if err != nil {
		return image, errors.Wrap(err, "failed to make request for CreateImage")
	}

	if err := json.Unmarshal(data, &image); err != nil {
		return image, errors.Wrap(err, "failed to decode CreateImage response")
	}

	return image, nil
}

// UpdateImage updates an existing machine image.
func (c ImageClient) UpdateImage(req UpdateImageRequest) (Image, error) {
	var image Image
	payload, err := json.Marshal(req)
	if err != nil {
		return image, errors.Wrap(err, "failed to marshal request for UpdateImage")
	}

	data, err := c.api.Put(fmt.Sprintf("images/%s", req.ID), payload)
	if err != nil {
		return image, errors.Wrap(err, "failed to make request for UpdateImage")
	}

	if err := json.Unmarshal(data, &image); err != nil {
		return image, errors.Wrap(err, "failed to decode UpdateImage response")
	}

	return image, nil
}

// DeleteImage retrieves a slice of machine images available in Linode.
func (c ImageClient) DeleteImage(id string) error {
	if _, err := c.api.Delete("images/" + id); err != nil {
		return errors.Wrap(err, "failed to make request for DeleteImage")
	}

	return nil
}

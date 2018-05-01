package lingo

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type ImageClient struct {
	api APIClient
}

func NewImageClient(api APIClient) ImageClient {
	return ImageClient{api: api}
}

// GetImages retrieves a slice of machine images available in Linode.
func (c ImageClient) GetImages() ([]Image, error) {
	data, err := c.api.Get("images")
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request for GetImages")
	}

	var results Results
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, errors.Wrap(err, "failed to decode GetImages response")
	}

	// TODO: Do something with paging here?
	var images []Image
	if err := json.Unmarshal(results.Data, &images); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal GetImages data")
	}

	return images, nil
}

// CreateImage creates a new machine image from an existing Linode disk.
func (c ImageClient) CreateImage(req NewImage) (Image, error) {
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

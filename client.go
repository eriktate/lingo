package lingo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// Base URI for the linode API.
const baseURI = "https://api.linode.com/v4/"

// A Client is capable of making API calls to the Linode API.
type Client struct {
	apiKey string
	h      *http.Client
}

// Results is the envelope format for most GET requests.
type Results struct {
	Data    json.RawMessage `json:"data"`
	Page    uint            `json:"page"`
	Pages   uint            `json:"pages"`
	Results uint            `json:"results"`
}

// NewClient returns a new Linode client struct loaded with the given
// API key.
func NewClient(apiKey string) Client {
	// TODO: Build a Client struct here instead of using the default.
	return Client{
		apiKey: apiKey,
		h:      http.DefaultClient,
	}
}

// GetImages retrieves a slice of machine images available in Linode.
func (c Client) GetImages() ([]Image, error) {
	req, err := c.makeGetRequest("images")
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request for GetImages")
	}

	res, err := c.h.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to complete GetImages request")
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.Wrap(err, "failed to GetImages")
	}

	var results Results
	if err := json.NewDecoder(res.Body).Decode(&results); err != nil {
		return nil, errors.Wrap(err, "failed to decode GetImages response")
	}

	// TODO: Do something with paging here?
	var images []Image
	if err := json.Unmarshal(results.Data, &images); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal GetImages data")
	}

	return images, nil
}

func (c Client) GetRegions() ([]Region, error) {
	req, err := c.makeGetRequest("regions")
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request for GetRegions")
	}

	res, err := c.h.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to complete GetRegions request")
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.Wrap(err, "failed to GetRegions")
	}

	var results Results
	if err := json.NewDecoder(res.Body).Decode(&results); err != nil {
		return nil, errors.Wrap(err, "failed to decode GetRegions response")
	}

	var regions []Region
	if err := json.Unmarshal(results.Data, &regions); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal GetRegions data")
	}

	return regions, nil
}

func (c Client) GetRegion(id string) (Region, error) {
	var region Region

	req, err := c.makeGetRequest("regions/" + id)
	if err != nil {
		return region, errors.Wrap(err, "failed to create request for GetRegion")
	}

	res, err := c.h.Do(req)
	if err != nil {
		return region, errors.Wrap(err, "failed to complete GetRegion request")
	}

	if res.StatusCode != http.StatusOK {
		return region, errors.Wrap(err, "failed to GetRegion")
	}

	if err := json.NewDecoder(res.Body).Decode(&region); err != nil {
		return region, errors.Wrap(err, "failed to decode GetRegion response")
	}

	return region, nil
}

func (c Client) makeGetRequest(path string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", baseURI, path), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	return req, nil
}

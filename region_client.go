package lingo

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type RegionClient struct {
	api APIClient
}

func NewRegionClient(api APIClient) RegionClient {
	return RegionClient{api: api}
}

func (c RegionClient) GetRegions() ([]Region, error) {
	data, err := c.api.Get("regions")
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request for GetRegions")
	}

	var results Results
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal GetRegions response")
	}

	var regions []Region
	if err := json.Unmarshal(results.Data, &regions); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal GetRegions data")
	}

	return regions, nil
}

func (c RegionClient) GetRegion(id string) (Region, error) {
	var region Region

	data, err := c.api.Get("regions/" + id)
	if err != nil {
		return region, errors.Wrap(err, "failed to make request for GetRegion")
	}

	if err := json.Unmarshal(data, &region); err != nil {
		return region, errors.Wrap(err, "failed to unmarshal GetRegion response")
	}

	return region, nil
}

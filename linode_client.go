package lingo

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type LinodeClient struct {
	apiKey string
	api    APIClient
}

func (c LinodeClient) CreateLinode(linode NewLinode) (Linode, error) {
	var created Linode

	payload, err := json.Marshal(&linode)
	if err != nil {
		return created, errors.Wrap(err, "failed to marshal request for CreateLinode")
	}

	data, err := c.api.Post("linode/instances", payload)
	if err != nil {
		return created, errors.Wrap(err, "failed to make request for CreateLinode")
	}

	if err := json.Unmarshal(data, &created); err != nil {
		return created, errors.Wrap(err, "failed to decode CreateLinode response")
	}

	return created, nil
}

func (c LinodeClient) GetLinodes() ([]Linode, error) {
	data, err := c.api.Get("linode/instances")
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request for GetLinodes")
	}

	var results Results
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal GetLinodes response")
	}

	var linodes []Linode
	if err := json.Unmarshal(results.Data, &linodes); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal GetLinodes data")
	}

	return linodes, nil
}

func (c LinodeClient) GetLinode(id uint) (Linode, error) {
	var linode Linode

	data, err := c.api.Get(fmt.Sprintf("linode/instances/%d", id))
	if err != nil {
		return linode, errors.Wrap(err, "failed to make request for GetLinode")
	}

	if err := json.Unmarshal(data, &linode); err != nil {
		return linode, errors.Wrap(err, "failed to decode GetLinode response")
	}

	return linode, nil
}

func (c LinodeClient) DeleteLinode(id uint) error {
	_, err := c.api.Delete(fmt.Sprintf("linode/instances/%d", id))
	if err != nil {
		return errors.Wrap(err, "failed to make request for DeleteLinode")
	}

	return nil
}

func (c LinodeClient) BootLinode(id uint) error {
	if _, err := c.api.Post(fmt.Sprintf("linode/instances/%d/boot", id), nil); err != nil {
		return errors.Wrap(err, "failed to make request for BootLinode")
	}

	return nil
}

func (c LinodeClient) BootLinodeWithConfig(id, configID uint) error {
	config := struct {
		ConfigID uint `json:"config_id"`
	}{configID}

	payload, err := json.Marshal(config)
	if err != nil {
		return errors.Wrap(err, "failed to marshal request for BootLinodeWithConfig")
	}

	if _, err := c.api.Post(fmt.Sprintf("linode/instances/%d/boot", id), payload); err != nil {
		return errors.Wrap(err, "failed to make request for BootLinode")
	}

	return nil
}

func (c LinodeClient) RebootLinode(id uint) error {
	if _, err := c.api.Post(fmt.Sprintf("linode/instances/%d/reboot", id), nil); err != nil {
		return errors.Wrap(err, "failed to make request for RebootLinode")
	}

	return nil
}

func (c LinodeClient) RebootLinodeWithConfig(id, configID uint) error {
	config := struct {
		ConfigID uint `json:"config_id"`
	}{configID}

	payload, err := json.Marshal(config)
	if err != nil {
		return errors.Wrap(err, "failed to marshal request for RebootLinodeWithConfig")
	}

	if _, err := c.api.Post(fmt.Sprintf("linode/instances/%d/reboot", id), payload); err != nil {
		return errors.Wrap(err, "failed to make request for RebootLinode")
	}

	return nil
}

func (c LinodeClient) ShutdownLinode(id uint) error {
	if _, err := c.api.Post(fmt.Sprintf("linode/instances/%d/shutdown", id), nil); err != nil {
		return errors.Wrap(err, "failed to make request for ShutdownLinode")
	}

	return nil
}

func (c LinodeClient) GetTypes() ([]LinodeType, error) {
	data, err := c.api.Get("linode/types")
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request for GetTypes")
	}

	var results Results
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal GetTypes response")
	}

	var types []LinodeType
	if err := json.Unmarshal(results.Data, &types); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal GetTypes data")
	}

	return types, nil
}

func (c LinodeClient) GetType(id string) (LinodeType, error) {
	var linodeType LinodeType

	data, err := c.api.Get("linode/types/" + id)
	if err != nil {
		return linodeType, errors.Wrap(err, "failed to make request for GetType")
	}

	if err := json.Unmarshal(data, &linodeType); err != nil {
		return linodeType, errors.Wrap(err, "failed to decode GetType response")
	}

	return linodeType, nil
}

func (c LinodeClient) ResizeLinode(id uint, typeID string) error {
	typePayload := struct {
		Type string `json:"type"`
	}{typeID}

	payload, err := json.Marshal(&typePayload)
	if err != nil {
		return errors.Wrap(err, "failed to marshal request for ResizeLinode")
	}

	if _, err := c.api.Post(fmt.Sprintf("linode/instances/%d/resize", id), payload); err != nil {
		return errors.Wrap(err, "failed to create request for GetTypes")
	}

	return nil
}

func (c LinodeClient) Mutate(id uint, typeID string) error {
	if _, err := c.api.Post(fmt.Sprintf("linode/instances/%d/mutate", id), nil); err != nil {
		return errors.Wrap(err, "failed to create request for Mutate")
	}

	return nil
}

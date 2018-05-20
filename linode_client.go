package lingo

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type LinodeClient struct {
	api APIClient
}

func NewLinodeClient(api APIClient) LinodeClient {
	return LinodeClient{api: api}
}

func (c LinodeClient) CreateLinode(linode CreateLinodeRequest) (Linode, error) {
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

func (c LinodeClient) UpdateLinode(req UpdateLinodeRequest) (Linode, error) {
	var updated Linode

	payload, err := json.Marshal(&req)
	if err != nil {
		return updated, errors.Wrap(err, "failed to marshal request for UpdateLinode")
	}

	data, err := c.api.Put(fmt.Sprintf("linode/instances/%d", req.ID), payload)
	if err != nil {
		return updated, errors.Wrap(err, "failed to make request for UpdateLinode")
	}

	if err := json.Unmarshal(data, &updated); err != nil {
		return updated, errors.Wrap(err, "failed to decode UpdateLinode response")
	}

	return updated, nil
}

func (c LinodeClient) ListLinodes() ([]Linode, error) {
	data, err := c.api.Get("linode/instances")
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request for ListLinodes")
	}

	var results Results
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal ListLinodes response")
	}

	var linodes []Linode
	if err := json.Unmarshal(results.Data, &linodes); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal ListLinodes data")
	}

	return linodes, nil
}

func (c LinodeClient) ViewLinode(id uint) (Linode, error) {
	var linode Linode

	data, err := c.api.Get(fmt.Sprintf("linode/instances/%d", id))
	if err != nil {
		return linode, errors.Wrap(err, "failed to make request for ViewLinode")
	}

	if err := json.Unmarshal(data, &linode); err != nil {
		return linode, errors.Wrap(err, "failed to decode ViewLinode response")
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

func (c LinodeClient) ListTypes() ([]LinodeType, error) {
	data, err := c.api.Get("linode/types")
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request for ListTypes")
	}

	var results Results
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal ListTypes response")
	}

	var types []LinodeType
	if err := json.Unmarshal(results.Data, &types); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal ListTypes data")
	}

	return types, nil
}

func (c LinodeClient) ViewType(id string) (LinodeType, error) {
	var linodeType LinodeType

	data, err := c.api.Get("linode/types/" + id)
	if err != nil {
		return linodeType, errors.Wrap(err, "failed to make request for ViewType")
	}

	if err := json.Unmarshal(data, &linodeType); err != nil {
		return linodeType, errors.Wrap(err, "failed to decode ViewType response")
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
		return errors.Wrap(err, "failed to make request for ResizeLinode")
	}

	return nil
}

func (c LinodeClient) Upgrade(id uint, typeID string) error {
	if _, err := c.api.Post(fmt.Sprintf("linode/instances/%d/mutate", id), nil); err != nil {
		return errors.Wrap(err, "failed to create request for Mutate")
	}

	return nil
}

func (c LinodeClient) CloneLinode(req CloneLinodeRequest) (Linode, error) {
	var clone Linode

	payload, err := json.Marshal(req)
	if err != nil {
		return clone, errors.Wrap(err, "failed to marshal request for CloneLinode")
	}

	data, err := c.api.Post(fmt.Sprintf("linode/instances/%d/clone", req.ID), payload)
	if err != nil {
		return clone, errors.Wrap(err, "failed to make request for CloneLinode")
	}

	if err := json.Unmarshal(data, &clone); err != nil {
		return clone, errors.Wrap(err, "failed to unmarshal CloneLinode data")
	}

	return clone, nil
}

func (c LinodeClient) RebuildLinode(req RebuildLinodeRequest) (Linode, error) {
	var linode Linode

	payload, err := json.Marshal(req)
	if err != nil {
		return linode, errors.Wrap(err, "failed to marshal request for RebuildLinode")
	}

	data, err := c.api.Post(fmt.Sprintf("linode/instances/%d/rebuild", req.ID), payload)
	if err != nil {
		return linode, errors.Wrap(err, "failed to make request for RebuildLinode")
	}

	if err := json.Unmarshal(data, &linode); err != nil {
		return linode, errors.Wrap(err, "failed to unmarshal RebuildLinode data")
	}

	return linode, nil
}

func (c LinodeClient) ListLinodeVolumes(id uint) ([]Volume, error) {
	data, err := c.api.Get(fmt.Sprintf("linode/%d/volumes", id))
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request for ListLinodeVolumes")
	}

	var results Results
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal ListLinodeVolumes response")
	}

	var volumes []Volume
	if err := json.Unmarshal(results.Data, &volumes); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal ListLinodeVolumes data")
	}

	return volumes, nil
}

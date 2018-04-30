package lingo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

type LinodeClient struct {
	apiKey string
	h      HTTPClient
}

func (c LinodeClient) CreateLinode(linode NewLinode) (Linode, error) {
	var created Linode

	data, err := json.Marshal(&linode)
	if err != nil {
		return created, errors.Wrap(err, "failed to marshal request for CreateLinode")
	}

	req, err := c.h.makePostRequest("linode/instances", data)
	if err != nil {
		return created, errors.Wrap(err, "failed to create request for CreateLinode")
	}

	res, err := c.h.Do(req)
	if err != nil {
		return created, errors.Wrap(err, "failed to complete CreateLinode request")
	}

	if res.StatusCode != http.StatusOK {
		errorText, _ := ioutil.ReadAll(res.Body)
		log.Println(string(errorText))
		return created, errors.New("failed to CreateLinode")
	}

	if err := json.NewDecoder(res.Body).Decode(&created); err != nil {
		return created, errors.Wrap(err, "failed to decode CreateLinode response")
	}

	return created, nil
}

func (c LinodeClient) GetLinodes() ([]Linode, error) {
	req, err := c.makeGetRequest("linode/instances")
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request for GetLinodes")
	}

	res, err := c.h.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to complete GetLinodes request")
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.Wrap(err, "failed to GetLinodes")
	}

	var results Results
	if err := json.NewDecoder(res.Body).Decode(&results); err != nil {
		return nil, errors.Wrap(err, "failed to decode GetLinodes response")
	}

	var linodes []Linode
	if err := json.Unmarshal(results.Data, &linodes); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal GetLinodes data")
	}

	return linodes, nil
}

func (c LinodeClient) GetLinode(id uint) (Linode, error) {
	var linode Linode

	req, err := c.makeGetRequest(fmt.Sprintf("linode/instances/%d", id))
	if err != nil {
		return linode, errors.Wrap(err, "failed to create request for GetLinode")
	}

	res, err := c.h.Do(req)
	if err != nil {
		return linode, errors.Wrap(err, "failed to complete GetLinode request")
	}

	if res.StatusCode != http.StatusOK {
		return linode, errors.Wrap(err, "failed to GetLinode")
	}

	if err := json.NewDecoder(res.Body).Decode(&linode); err != nil {
		return linode, errors.Wrap(err, "failed to decode GetLinode response")
	}

	return linode, nil
}

func (c LinodeClient) DeleteLinode(id uint) error {
	req, err := c.makeDeleteRequest(fmt.Sprintf("linode/instances/%d", id))
	if err != nil {
		return errors.Wrap(err, "failed to create request for DeleteLinode")
	}

	res, err := c.h.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to complete DeleteLinode request")
	}

	if res.StatusCode != http.StatusOK {
		return errors.Wrap(err, "failed to DeleteLinode")
	}

	return nil
}

func (c LinodeClient) BootLinode(id uint) error {
	req, err := c.makePostRequest(fmt.Sprintf("linode/instances/%d/boot", id), nil)
	if err != nil {
		return errors.Wrap(err, "failed to create request for BootLinode")
	}

	res, err := c.h.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to complete BootLinode request")
	}

	if res.StatusCode != http.StatusOK {
		return errors.Wrap(err, "failed to BootLinode")
	}

	return nil
}

func (c LinodeClient) BootLinodeWithConfig(id, configID uint) error {
	config := struct {
		ConfigID uint `json:"config_id"`
	}{configID}

	data, err := json.Marshal(config)
	if err != nil {
		return errors.Wrap(err, "failed to marshal request for BootLinodeWithConfig")
	}

	req, err := c.makePostRequest(fmt.Sprintf("linode/instances/%d/boot", id), data)
	if err != nil {
		return errors.Wrap(err, "failed to create request for BootLinode")
	}

	res, err := c.h.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to complete BootLinode request")
	}

	if res.StatusCode != http.StatusOK {
		return errors.Wrap(err, "failed to BootLinode")
	}

	return nil
}

func (c LinodeClient) RebootLinode(id uint) error {
	req, err := c.makePostRequest(fmt.Sprintf("linode/instances/%d/reboot", id), nil)
	if err != nil {
		return errors.Wrap(err, "failed to create request for RebootLinode")
	}

	res, err := c.h.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to complete RebootLinode request")
	}

	if res.StatusCode != http.StatusOK {
		return errors.Wrap(err, "failed to RebootLinode")
	}

	return nil
}

func (c LinodeClient) RebootLinodeWithConfig(id, configID uint) error {
	config := struct {
		ConfigID uint `json:"config_id"`
	}{configID}

	data, err := json.Marshal(config)
	if err != nil {
		return errors.Wrap(err, "failed to marshal request for RebootLinodeWithConfig")
	}

	req, err := c.makePostRequest(fmt.Sprintf("linode/instances/%d/reboot", id), data)
	if err != nil {
		return errors.Wrap(err, "failed to create request for RebootLinode")
	}

	res, err := c.h.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to complete RebootLinode request")
	}

	if res.StatusCode != http.StatusOK {
		return errors.Wrap(err, "failed to RebootLinode")
	}

	return nil
}

func (c LinodeClient) ShutdownLinode(id uint) error {
	req, err := c.makePostRequest(fmt.Sprintf("linode/instances/%d/shutdown", id), nil)
	if err != nil {
		return errors.Wrap(err, "failed to create request for ShutdownLinode")
	}

	res, err := c.h.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to complete ShutdownLinode request")
	}

	if res.StatusCode != http.StatusOK {
		return errors.Wrap(err, "failed to ShutdownLinode")
	}

	return nil
}

func (c LinodeClient) GetTypes() ([]LinodeType, error) {
	req, err := c.makeGetRequest("linode/types")
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request for GetTypes")
	}

	res, err := c.h.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to complete GetTypes request")
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.Wrap(err, "failed to GetTypes")
	}

	var results Results
	if err := json.NewDecoder(res.Body).Decode(&results); err != nil {
		return nil, errors.Wrap(err, "failed to decode GetTypes response")
	}

	var types []LinodeType
	if err := json.Unmarshal(results.Data, &types); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal GetTypes data")
	}

	return types, nil
}

func (c LinodeClient) GetType(id string) (LinodeType, error) {
	var linodeType LinodeType

	req, err := c.makeGetRequest("linode/types/" + id)
	if err != nil {
		return linodeType, errors.Wrap(err, "failed to create request for GetType")
	}

	res, err := c.h.Do(req)
	if err != nil {
		return linodeType, errors.Wrap(err, "failed to complete GetType request")
	}

	if res.StatusCode != http.StatusOK {
		return linodeType, errors.Wrap(err, "failed to GetType")
	}

	if err := json.NewDecoder(res.Body).Decode(&linodeType); err != nil {
		return linodeType, errors.Wrap(err, "failed to decode GetType response")
	}

	return linodeType, nil
}

func (c LinodeClient) ResizeLinode(id uint, typeID string) error {
	typePayload := struct {
		Type string `json:"type"`
	}{typeID}

	data, err := json.Marshal(&typePayload)
	if err != nil {
		return errors.Wrap(err, "failed to marshal request for ResizeLinode")
	}

	req, err := c.makePostRequest(fmt.Sprintf("linode/instances/%d/resize", id), data)
	if err != nil {
		return errors.Wrap(err, "failed to create request for GetTypes")
	}

	res, err := c.h.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to complete GetTypes request")
	}

	if res.StatusCode != http.StatusOK {
		return errors.Wrap(err, "failed to GetTypes")
	}

	var results Results
	if err := json.NewDecoder(res.Body).Decode(&results); err != nil {
		return errors.Wrap(err, "failed to decode GetTypes response")
	}

	var types []LinodeType
	if err := json.Unmarshal(results.Data, &types); err != nil {
		return errors.Wrap(err, "failed to unmarshal GetTypes data")
	}

	return nil
}

func (c LinodeClient) Mutate(id uint, typeID string) error {
	req, err := c.makePostRequest(fmt.Sprintf("linode/instances/%d/mutate", id), nil)
	if err != nil {
		return errors.Wrap(err, "failed to create request for Mutate")
	}

	res, err := c.h.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to complete Mutate request")
	}

	if res.StatusCode != http.StatusOK {
		return errors.Wrap(err, "failed to Mutate")
	}

	return nil
}

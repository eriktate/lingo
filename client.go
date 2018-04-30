package lingo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

func (c Client) GetNodeBalancers() ([]NodeBalancer, error) {
	req, err := c.makeGetRequest("nodebalancers")
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request for GetNodeBalancers")
	}

	res, err := c.h.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to complete GetNodeBalancers request")
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.Wrap(err, "failed to GetNodeBalancers")
	}

	var results Results
	if err := json.NewDecoder(res.Body).Decode(&results); err != nil {
		return nil, errors.Wrap(err, "failed to decode GetNodeBalancers response")
	}

	var balancers []NodeBalancer
	if err := json.Unmarshal(results.Data, &balancers); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal GetNodeBalancers data")
	}

	return balancers, nil
}

func (c Client) GetNodeBalancer(id uint) (NodeBalancer, error) {
	var balancer NodeBalancer
	req, err := c.makeGetRequest(fmt.Sprintf("nodebalancers/%d", id))
	if err != nil {
		return balancer, errors.Wrap(err, "failed to create request for GetNodeBalancer")
	}

	res, err := c.h.Do(req)
	if err != nil {
		return balancer, errors.Wrap(err, "failed to complete GetNodeBalancer request")
	}

	if res.StatusCode != http.StatusOK {
		return balancer, errors.Wrap(err, "failed to GetNodeBalancer")
	}

	if err := json.NewDecoder(res.Body).Decode(&balancer); err != nil {
		return balancer, errors.Wrap(err, "failed to decode GetNodeBalancer response")
	}

	return balancer, nil
}

func (c Client) DeleteNodeBalancer(id uint) error {
	req, err := c.makeDeleteRequest(fmt.Sprintf("nodebalancers/%d", id))
	if err != nil {
		return errors.Wrap(err, "failed to create request for DeleteNodeBalancer")
	}

	res, err := c.h.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to complete DeleteNodeBalancer request")
	}

	if res.StatusCode != http.StatusOK {
		return errors.Wrap(err, "failed to DeleteNodeBalancer")
	}

	return nil
}

func (c Client) CreateNodeBalancer(request CreateBalancerRequest) (NodeBalancer, error) {
	var created NodeBalancer

	data, err := json.Marshal(&request)
	if err != nil {
		return created, errors.Wrap(err, "failed to marshal request for CreateNodeBalancer")
	}

	req, err := c.makePostRequest("nodebalancers", data)
	if err != nil {
		return created, errors.Wrap(err, "failed to create request for CreateNodeBalancer")
	}

	res, err := c.h.Do(req)
	if err != nil {
		return created, errors.Wrap(err, "failed to complete CreateNodeBalancer request")
	}

	if res.StatusCode != http.StatusOK {
		errorText, _ := ioutil.ReadAll(res.Body)
		log.Println(string(errorText))
		return created, errors.New("failed to CreateNodeBalancer")
	}

	if err := json.NewDecoder(res.Body).Decode(&created); err != nil {
		return created, errors.Wrap(err, "failed to decode CreateNodeBalancer response")
	}

	return created, nil
}

func (c Client) UpdateNodeBalancer(request UpdateBalancerRequest) (NodeBalancer, error) {
	var created NodeBalancer

	data, err := json.Marshal(&request)
	if err != nil {
		return created, errors.Wrap(err, "failed to marshal request for UpdateNodeBalancer")
	}

	req, err := c.makePutRequest(fmt.Sprintf("nodebalancers/%d", request.ID), data)
	if err != nil {
		return created, errors.Wrap(err, "failed to create request for UpdateNodeBalancer")
	}

	res, err := c.h.Do(req)
	if err != nil {
		return created, errors.Wrap(err, "failed to complete UpdateNodeBalancer request")
	}

	if res.StatusCode != http.StatusOK {
		errorText, _ := ioutil.ReadAll(res.Body)
		log.Println(string(errorText))
		return created, errors.New("failed to UpdateNodeBalancer")
	}

	if err := json.NewDecoder(res.Body).Decode(&created); err != nil {
		return created, errors.Wrap(err, "failed to decode UpdateNodeBalancer response")
	}

	return created, nil
}

func (c Client) CreateLinode(linode NewLinode) (Linode, error) {
	var created Linode

	data, err := json.Marshal(&linode)
	if err != nil {
		return created, errors.Wrap(err, "failed to marshal request for CreateLinode")
	}

	req, err := c.makePostRequest("linode/instances", data)
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

func (c Client) GetLinodes() ([]Linode, error) {
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

func (c Client) GetLinode(id uint) (Linode, error) {
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

func (c Client) DeleteLinode(id uint) error {
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

func (c Client) BootLinode(id uint) error {
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

func (c Client) BootLinodeWithConfig(id, configID uint) error {
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

func (c Client) RebootLinode(id uint) error {
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

func (c Client) RebootLinodeWithConfig(id, configID uint) error {
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

func (c Client) ShutdownLinode(id uint) error {
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

func (c Client) makeGetRequest(path string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", baseURI, path), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	return req, nil
}

func (c Client) makePostRequest(path string, data []byte) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", baseURI, path), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c Client) makePutRequest(path string, data []byte) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", baseURI, path), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c Client) makeDeleteRequest(path string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s%s", baseURI, path), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	return req, nil
}

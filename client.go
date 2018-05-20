package lingo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// Base URI for the linode API.
const baseURI = "https://api.linode.com/v4/"

// An APIClient is capable of making API calls to the Linode API.
type APIClient struct {
	apiKey  string
	backoff *backoffConfig
	h       *http.Client
}

// Results is the envelope format for GET requests that return paged data.
type Results struct {
	Data    json.RawMessage `json:"data"`
	Page    uint            `json:"page"`
	Pages   uint            `json:"pages"`
	Results uint            `json:"results"`
}

// NewAPIClient returns a new Linode client struct loaded with the given
// API key.
func NewAPIClient(apiKey string, backoff bool) APIClient {
	// TODO: Build a Client struct here instead of using the default.
	return APIClient{
		apiKey: apiKey,
		h:      http.DefaultClient,
	}
}

// TODO: Might be better to return the http.Response pointer, but until it's necessary just return the body.
func (c APIClient) Get(path string) ([]byte, error) {
	req, err := c.makeGetRequest(path)
	if err != nil {
		return nil, err
	}

	return c.do(req)
}

func (c APIClient) Post(path string, payload []byte) ([]byte, error) {
	req, err := c.makePostRequest(path, payload)
	if err != nil {
		return nil, err
	}

	return c.do(req)
}

func (c APIClient) Put(path string, payload []byte) ([]byte, error) {
	req, err := c.makePutRequest(path, payload)
	if err != nil {
		return nil, err
	}

	return c.do(req)
}

func (c APIClient) Delete(path string) ([]byte, error) {
	req, err := c.makeDeleteRequest(path)
	if err != nil {
		return nil, err
	}

	return c.do(req)
}

func (c APIClient) do(req *http.Request) ([]byte, error) {
	if c.backoff != nil {
		defer c.backoff.Reset()
	}

	res, err := c.h.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		var errs Errors
		if err := json.Unmarshal(data, &errs); err != nil {
			return nil, errors.Errorf("request failed, could not unmarshal error response. Raw error: %s", string(data))
		}

		if errs.IsBusy() && c.backoff != nil {
			if err := c.backoff.Retry(); err != nil {
				return nil, errors.Wrap(errs, err.Error())
			}

			return c.do(req)
		}

		return nil, err
	}

	return data, nil
}

func (c APIClient) makeGetRequest(path string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", baseURI, path), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	return req, nil
}

func (c APIClient) makePostRequest(path string, data []byte) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", baseURI, path), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c APIClient) makePutRequest(path string, data []byte) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", baseURI, path), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c APIClient) makeDeleteRequest(path string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s%s", baseURI, path), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	return req, nil
}

// Lingo is an aggregation of all Linode client implementations. With a Lingo struct, you can make
// any type of Linode API call.
type Lingo struct {
	LinodeClient
	BalancerClient
	ImageClient
	RegionClient
	DomainClient
	VolumeClient
	DiskClient
}

// NewLingo returns a new Lingo struct given a Linode API key.
func NewLingo(apiKey string, backoff bool) Lingo {
	api := NewAPIClient(apiKey, backoff)

	return Lingo{
		LinodeClient:   NewLinodeClient(api),
		BalancerClient: NewBalancerClient(api),
		ImageClient:    NewImageClient(api),
		RegionClient:   NewRegionClient(api),
		DomainClient:   NewDomainClient(api),
		VolumeClient:   NewVolumeClient(api),
		DiskClient:     NewDiskClient(api),
	}
}

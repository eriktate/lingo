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

// An APIClient is capable of making API calls to the Linode API.
type APIClient struct {
	apiKey string
	h      *http.Client
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
func NewAPIClient(apiKey string) APIClient {
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
	res, err := c.h.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// TODO: Figure out how to properly return this error and remove log.
	if res.StatusCode != http.StatusOK {
		log.Println(string(data))
		return nil, errors.New("request failed")
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
}

func NewLingo(apiKey string) Lingo {
	api := NewAPIClient(apiKey)

	return Lingo{
		LinodeClient:   NewLinodeClient(api),
		BalancerClient: NewBalancerClient(api),
		ImageClient:    NewImageClient(api),
		RegionClient:   NewRegionClient(api),
	}
}

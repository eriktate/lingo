package lingo

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type BalancerClient struct {
	api APIClient
}

func NewBalancerClient(api APIClient) BalancerClient {
	return BalancerClient{api: api}
}

func (c BalancerClient) ListNodeBalancers() ([]NodeBalancer, error) {
	data, err := c.api.Get("nodebalancers")
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request for ListNodeBalancer")
	}

	var results Results
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal ListNodeBalancer response")
	}

	var balancers []NodeBalancer
	if err := json.Unmarshal(results.Data, &balancers); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal ListNodeBalancer data")
	}

	return balancers, nil
}

func (c BalancerClient) ViewNodeBalancer(id uint) (NodeBalancer, error) {
	var balancer NodeBalancer
	data, err := c.api.Get(fmt.Sprintf("nodebalancers/%d", id))
	if err != nil {
		return balancer, errors.Wrap(err, "failed to make request for ViewNodeBalancer")
	}

	if err := json.Unmarshal(data, &balancer); err != nil {
		return balancer, errors.Wrap(err, "failed to decode ViewNodeBalancer response")
	}

	return balancer, nil
}

func (c BalancerClient) DeleteNodeBalancer(id uint) error {
	if _, err := c.api.Delete(fmt.Sprintf("nodebalancers/%d", id)); err != nil {
		return errors.Wrap(err, "failed to make request for DeleteNodeBalancer")
	}

	return nil
}

func (c BalancerClient) CreateNodeBalancer(request CreateBalancerRequest) (NodeBalancer, error) {
	var created NodeBalancer

	payload, err := json.Marshal(&request)
	if err != nil {
		return created, errors.Wrap(err, "failed to marshal request for CreateNodeBalancer")
	}

	data, err := c.api.Post("nodebalancers", payload)
	if err != nil {
		return created, errors.Wrap(err, "failed to make request for CreateNodeBalancer")
	}

	if err := json.Unmarshal(data, &created); err != nil {
		return created, errors.Wrap(err, "failed to unmarshal CreateNodeBalancer response")
	}

	return created, nil
}

func (c BalancerClient) UpdateNodeBalancer(request UpdateBalancerRequest) (NodeBalancer, error) {
	var created NodeBalancer

	payload, err := json.Marshal(&request)
	if err != nil {
		return created, errors.Wrap(err, "failed to marshal request for UpdateNodeBalancer")
	}

	data, err := c.api.Put(fmt.Sprintf("nodebalancers/%d", request.ID), payload)
	if err != nil {
		return created, errors.Wrap(err, "failed to make request for UpdateNodeBalancer")
	}

	if err := json.Unmarshal(data, &created); err != nil {
		return created, errors.Wrap(err, "failed to unmarshal UpdateNodeBalancer response")
	}

	return created, nil
}

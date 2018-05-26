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

func (c BalancerClient) DeleteNodeBalancer(id uint) error {
	if _, err := c.api.Delete(fmt.Sprintf("nodebalancers/%d", id)); err != nil {
		return errors.Wrap(err, "failed to make request for DeleteNodeBalancer")
	}

	return nil
}

func (c BalancerClient) ListNodeBalancerConfigs(balancerID uint) ([]BalancerConfig, error) {
	data, err := c.api.Get(fmt.Sprintf("nodebalancers/%d/configs", balancerID))
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request for ListNodeBalancerConfigs")
	}

	var results Results
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal ListNodeBalancerConfigs response")
	}

	var configs []BalancerConfig
	if err := json.Unmarshal(results.Data, &configs); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal ListNodeBalancerConfigs data")
	}

	return configs, nil
}

func (c BalancerClient) ViewNodeBalancerConfig(balancerID, configID uint) (BalancerConfig, error) {
	var config BalancerConfig
	data, err := c.api.Get(fmt.Sprintf("nodebalancers/%d/configs/%d", balancerID, configID))
	if err != nil {
		return config, errors.Wrap(err, "failed to make request for ViewNodeBalancerConfig")
	}

	if err := json.Unmarshal(data, &config); err != nil {
		return config, errors.Wrap(err, "failed to decode ViewNodeBalancerConfig response")
	}

	return config, nil
}

func (c BalancerClient) CreateNodeBalancerConfig(req CreateBalancerConfigRequest) (BalancerConfig, error) {
	var created BalancerConfig

	payload, err := json.Marshal(&req)
	if err != nil {
		return created, errors.Wrap(err, "failed to marshal request for CreateNodeBalancerConfig")
	}

	data, err := c.api.Post(fmt.Sprintf("nodebalancers/%d/configs", req.NodeBalancerID), payload)
	if err != nil {
		return created, errors.Wrap(err, "failed to make request for CreateNodeBalancerConfig")
	}

	if err := json.Unmarshal(data, &created); err != nil {
		return created, errors.Wrap(err, "failed to unmarshal CreateNodeBalancerConfig response")
	}

	return created, nil
}

func (c BalancerClient) UpdateNodeBalancerConfig(req UpdateBalancerConfigRequest) (NodeBalancer, error) {
	var updated NodeBalancer

	payload, err := json.Marshal(&req)
	if err != nil {
		return updated, errors.Wrap(err, "failed to marshal request for UpdateNodeBalancerConfig")
	}

	data, err := c.api.Put(fmt.Sprintf("nodebalancers/%d/configs/%d", req.NodeBalancerID, req.ID), payload)
	if err != nil {
		return updated, errors.Wrap(err, "failed to make request for UpdateNodeBalancerConfig")
	}

	if err := json.Unmarshal(data, &updated); err != nil {
		return updated, errors.Wrap(err, "failed to unmarshal UpdateNodeBalancerConfig response")
	}

	return updated, nil
}

func (c BalancerClient) DeleteNodeBalancerConfig(balancerID, configID uint) error {
	if _, err := c.api.Delete(fmt.Sprintf("nodebalancers/%d/configs/%d", balancerID, configID)); err != nil {
		return errors.Wrap(err, "failed to make request for DeleteNodeBalancerConfig")
	}

	return nil
}

func (c BalancerClient) ListNodes(balancerID, configID uint) ([]Node, error) {
	data, err := c.api.Get(fmt.Sprintf("nodebalancers/%d/configs/%d/nodes", balancerID, configID))
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request for ListNodes")
	}

	var results Results
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal ListNodes response")
	}

	var nodes []Node
	if err := json.Unmarshal(results.Data, &nodes); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal ListNodes data")
	}

	return nodes, nil
}

func (c BalancerClient) ViewNode(balancerID, configID, nodeID uint) (Node, error) {
	var node Node
	data, err := c.api.Get(fmt.Sprintf("nodebalancers/%d/configs/%d/nodes/%d", balancerID, configID, nodeID))
	if err != nil {
		return node, errors.Wrap(err, "failed to make request for ViewNode")
	}

	if err := json.Unmarshal(data, &node); err != nil {
		return node, errors.Wrap(err, "failed to decode ViewNode response")
	}

	return node, nil
}

func (c BalancerClient) CreateNode(req CreateNodeRequest) (Node, error) {
	var created Node

	payload, err := json.Marshal(&req)
	if err != nil {
		return created, errors.Wrap(err, "failed to marshal request for CreateNode")
	}

	data, err := c.api.Post(fmt.Sprintf("nodebalancers/%d/configs/%d/nodes", req.BalancerID, req.ConfigID), payload)
	if err != nil {
		return created, errors.Wrap(err, "failed to make request for CreateNode")
	}

	if err := json.Unmarshal(data, &created); err != nil {
		return created, errors.Wrap(err, "failed to unmarshal CreateNode response")
	}

	return created, nil
}

func (c BalancerClient) UpdateNode(req UpdateNodeRequest) (Node, error) {
	var updated Node

	payload, err := json.Marshal(&req)
	if err != nil {
		return updated, errors.Wrap(err, "failed to marshal request for UpdateNode")
	}

	data, err := c.api.Put(fmt.Sprintf("nodebalancers/%d/configs/%d/nodes/%d", req.BalancerID, req.ConfigID, req.ID), payload)
	if err != nil {
		return updated, errors.Wrap(err, "failed to make request for UpdateNode")
	}

	if err := json.Unmarshal(data, &updated); err != nil {
		return updated, errors.Wrap(err, "failed to unmarshal UpdateNode response")
	}

	return updated, nil
}

func (c BalancerClient) DeleteNode(balancerID, configID, nodeID uint) error {
	if _, err := c.api.Delete(fmt.Sprintf("nodebalancers/%d/configs/%d/nodes/%d", balancerID, configID, nodeID)); err != nil {
		return errors.Wrap(err, "failed to make request for DeleteNode")
	}

	return nil
}

package lingo_test

import (
	"os"
	"testing"

	"github.com/eriktate/lingo"
)

func Test_CRUDBalancers(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey, nil)
	client := lingo.NewBalancerClient(api)

	createRequest1 := lingo.CreateBalancerRequest{
		Region:             "us-east-1a",
		Label:              "a_test_balancer",
		ClientConnThrottle: 10,
	}

	createRequest2 := lingo.CreateBalancerRequest{
		Region:             "us-east-1a",
		Label:              "another_test_balancer",
		ClientConnThrottle: 10,
	}

	balancer1, err := client.CreateNodeBalancer(createRequest1)
	if err != nil {
		t.Fatalf("Failed to create balancer1: %s", err)
	}

	balancer2, err := client.CreateNodeBalancer(createRequest2)
	if err != nil {
		t.Fatalf("Failed to create balancer2: %s", err)
	}

	if _, err = client.ListNodeBalancers(); err != nil {
		t.Fatalf("Failed to fetch balancers: %s", err)
	}

	fetch1, err := client.ViewNodeBalancer(balancer1.ID)
	if err != nil {
		t.Fatalf("Failed to fetch balancer1: %s", err)
	}

	updateRequest := lingo.UpdateBalancerRequest{
		ID:                 balancer2.ID,
		Label:              "some_other_balancer",
		ClientConnThrottle: 5,
	}

	if _, err := client.UpdateNodeBalancer(updateRequest); err != nil {
		t.Fatalf("Failed to update balancer2: %s", err)
	}

	fetch2, err := client.ViewNodeBalancer(balancer2.ID)
	if err != nil {
		t.Fatalf("Failed to fetch balancer2: %s", err)
	}

	if fetch2.Label != updateRequest.Label {
		t.Fatal("Update was not applied")
	}

	if err := client.DeleteNodeBalancer(fetch1.ID); err != nil {
		t.Fatalf("Failed to delete balancer1: %s", err)
	}

	if err := client.DeleteNodeBalancer(fetch2.ID); err != nil {
		t.Fatalf("Failed to delete balancer2: %s", err)
	}
}

func Test_CRUDBalancerConfig(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey, nil)
	client := lingo.NewBalancerClient(api)

	createRequest := lingo.CreateBalancerRequest{
		Region:             "us-east-1a",
		Label:              "a_test_balancer",
		ClientConnThrottle: 10,
	}

	nb, err := client.CreateNodeBalancer(createRequest)
	if err != nil {
		t.Fatalf("Failed to create node balancer: %s", err)
	}

	createConfig := lingo.CreateBalancerConfigRequest{
		NodeBalancerID: nb.ID,
		Protocol:       lingo.ProtocolHTTP,
		Algorithm:      lingo.AlgoRoundRobin,
		Stickiness:     lingo.StickyNone,
		Check:          lingo.CheckNone,
		Port:           80,
		CipherSuite:    lingo.CipherSuiteRecommended,
	}

	conf, err := client.CreateNodeBalancerConfig(createConfig)
	if err != nil {
		t.Fatalf("Failed to create config: %s", err)
	}

	updateConfig := lingo.UpdateBalancerConfigRequest{
		NodeBalancerID: createConfig.NodeBalancerID,
		ID:             conf.ID,
		Stickiness:     lingo.StickyHTTPCookie,
	}

	if _, err := client.UpdateNodeBalancerConfig(updateConfig); err != nil {
		t.Fatalf("Failed to update node balancer: %s", err)
	}

	getConfig, err := client.ViewNodeBalancerConfig(conf.NodeBalancerID, conf.ID)
	if err != nil {
		t.Fatalf("Failed to view node balancer: %s", err)
	}

	if getConfig.Stickiness != updateConfig.Stickiness {
		t.Fatalf("Update failed to apply. Expectin %s, but got %s", updateConfig.Stickiness, getConfig.Stickiness)
	}

	configs, err := client.ListNodeBalancerConfigs(nb.ID)
	if err != nil {
		t.Fatalf("Failed to list configs: %s", err)
	}

	if len(configs) == 0 {
		t.Fatal("Listing configs returned no results")
	}

	if err := client.DeleteNodeBalancerConfig(nb.ID, conf.ID); err != nil {
		t.Fatalf("Failed to delete config: %s", err)
	}

	if err := client.DeleteNodeBalancer(nb.ID); err != nil {
		t.Fatalf("Failed to cleanup node balancer: %s", err)
	}
}

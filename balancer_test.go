package lingo_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/eriktate/lingo"
)

func Test_CRUDBalancers(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey, nil)
	client := lingo.NewBalancerClient(api)

	existing, err := client.ListNodeBalancers()
	if err != nil {
		t.Fatalf("Failed to list domains: %s", err)
	}

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

	balancers, err := client.ListNodeBalancers()
	if err != nil {
		t.Fatalf("Failed to fetch balancers: %s", err)
	}

	expected := len(existing) + 2
	if len(balancers) != expected {
		t.Fatalf("Something strange happened. Expected to list %d balancers, but got %d", expected, len(balancers))
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
		BalancerID:  nb.ID,
		Protocol:    lingo.ProtocolHTTP,
		Algorithm:   lingo.AlgoRoundRobin,
		Stickiness:  lingo.StickyNone,
		Check:       lingo.CheckNone,
		Port:        80,
		CipherSuite: lingo.CipherSuiteRecommended,
	}

	conf, err := client.CreateNodeBalancerConfig(createConfig)
	if err != nil {
		t.Fatalf("Failed to create config: %s", err)
	}

	updateConfig := lingo.UpdateBalancerConfigRequest{
		BalancerID: createConfig.BalancerID,
		ID:         conf.ID,
		Stickiness: lingo.StickyHTTPCookie,
	}

	log.Printf("Update config: %+v", updateConfig)
	if _, err := client.UpdateNodeBalancerConfig(updateConfig); err != nil {
		t.Fatalf("Failed to update node balancer: %s", err)
	}

	getConfig, err := client.ViewNodeBalancerConfig(conf.BalancerID, conf.ID)
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

func Test_CRUDNode(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey, nil)
	client := lingo.NewBalancerClient(api)
	linodeClient := lingo.NewLinodeClient(api)

	createLinode := lingo.CreateLinodeRequest{
		Region:   "us-east-1a",
		Type:     "g5-nanode-1",
		Image:    "linode/debian9",
		RootPass: "test123",
	}

	testLinode, err := linodeClient.CreateLinode(createLinode)
	if err != nil {
		t.Fatalf("Failed to create linode: %s", err)
	}

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
		BalancerID:  nb.ID,
		Protocol:    lingo.ProtocolHTTP,
		Algorithm:   lingo.AlgoRoundRobin,
		Stickiness:  lingo.StickyNone,
		Check:       lingo.CheckNone,
		Port:        80,
		CipherSuite: lingo.CipherSuiteRecommended,
	}

	conf, err := client.CreateNodeBalancerConfig(createConfig)
	if err != nil {
		t.Fatalf("Failed to create config: %s", err)
	}

	log.Printf("Linode IP addresses: %+v", testLinode.IPv4)
	createNode := lingo.CreateNodeRequest{
		BalancerID: nb.ID,
		ConfigID:   conf.ID,
		Label:      "test_node",
		Address:    fmt.Sprintf("%s:80", testLinode.IPv4[0]),
		Mode:       lingo.NodeModeReject,
	}

	node, err := client.CreateNode(createNode)
	if err != nil {
		t.Fatalf("Failed to create node: %s", err)
	}

	updateNode := lingo.UpdateNodeRequest{
		BalancerID: nb.ID,
		ConfigID:   conf.ID,
		ID:         node.ID,
		Mode:       lingo.NodeModeAccept,
	}

	if _, err := client.UpdateNode(updateNode); err != nil {
		t.Fatalf("Failed to update node: %s", err)
	}

	getNode, err := client.ViewNode(nb.ID, conf.ID, node.ID)
	if err != nil {
		t.Fatalf("Failed to retrieve node: %s", err)
	}

	if getNode.Mode != updateNode.Mode {
		t.Fatalf("Update failed to applay. Expected mode to be %s, but got %s", updateNode.Mode, getNode.Mode)
	}

	nodes, err := client.ListNodes(nb.ID, conf.ID)
	if err != nil {
		t.Fatalf("Failed to list nodes: %s", err)
	}

	if len(nodes) == 0 {
		t.Fatal("ListNodes returned no results")
	}

	if err := client.DeleteNode(nb.ID, conf.ID, node.ID); err != nil {
		t.Fatalf("Failed to delete node: %s", err)
	}

	if err := client.DeleteNodeBalancer(nb.ID); err != nil {
		t.Fatalf("Failed to cleanup node balancer: %s", err)
	}
}

package lingo_test

import (
	"os"
	"testing"

	"github.com/eriktate/lingo"
)

func Test_Integration_Balancers(t *testing.T) {
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

	if _, err = client.GetNodeBalancers(); err != nil {
		t.Fatalf("Failed to fetch balancers: %s", err)
	}

	fetch1, err := client.GetNodeBalancer(balancer1.ID)
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

	fetch2, err := client.GetNodeBalancer(balancer2.ID)
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

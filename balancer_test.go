package lingo_test

import (
	"os"
	"testing"

	"github.com/eriktate/lingo"
)

func Test_Integration_Balancers(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	client := lingo.NewClient(apiKey)

	createRequest1 := lingo.CreateNodeBalancerRequest{
		Region:             "us-east-1a",
		Label:              "a_test_balancer",
		ClientConnThrottle: 10,
	}

	createRequest2 := lingo.CreateNodeBalancerRequest{
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

	_, err = client.GetNodeBalancers()
	if err != nil {
		t.Fatalf("Failed to fetch balancers: %s", err)
	}

	fetch1, err := client.GetNodeBalancer(balancer1.ID)
	if err != nil {
		t.Fatalf("Failed to fetch balancer1: %s", err)
	}

	fetch2, err := client.GetNodeBalancer(balancer2.ID)
	if err != nil {
		t.Fatalf("Failed to fetch balancer2: %s", err)
	}

	if err := client.DeleteNodeBalancer(fetch1.ID); err != nil {
		t.Fatalf("Failed to delete balancer1: %s", err)
	}

	if err := client.DeleteNodeBalancer(fetch2.ID); err != nil {
		t.Fatalf("Failed to delete balancer2: %s", err)
	}
}

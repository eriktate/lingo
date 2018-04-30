package lingo_test

import (
	"os"
	"testing"

	"github.com/eriktate/lingo"
)

func Test_Integration_Linodes(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	client := lingo.NewClient(apiKey)

	createLinode1 := lingo.NewLinode{
		Region:   "us-east-1a",
		Type:     "g5-nanode-1",
		Image:    "linode/debian9",
		RootPass: "test123",
	}

	createLinode2 := lingo.NewLinode{
		Region:   "us-west",
		Type:     "g5-nanode-1",
		Image:    "linode/debian9",
		RootPass: "test123",
	}

	created1, err := client.CreateLinode(createLinode1)
	if err != nil {
		t.Fatalf("Failed to create linode1: %s", err)
	}

	created2, err := client.CreateLinode(createLinode2)
	if err != nil {
		t.Fatalf("Failed to create linode2: %s", err)
	}

	_, err = client.GetLinode(created1.ID)
	if err != nil {
		t.Fatalf("Failed to fetch linode1: %s", err)
	}

	linodes, err := client.GetLinodes()
	if err != nil {
		t.Fatalf("Failed to fetch linodes: %s", err)
	}

	if len(linodes) == 0 {
		t.Fatalf("Failed to retrieve any linodes")
	}

	if err := client.DeleteLinode(created1.ID); err != nil {
		t.Fatalf("Failed to delete linode1: %s", err)
	}

	if err := client.DeleteLinode(created2.ID); err != nil {
		t.Fatalf("Failed to delete linode2: %s", err)
	}
}

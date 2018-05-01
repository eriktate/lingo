package lingo_test

import (
	"os"
	"testing"

	"github.com/eriktate/lingo"
)

func Test_Integration_Linodes(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey)
	client := lingo.NewLinodeClient(api)

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

	if _, err := client.GetTypes(); err != nil {
		t.Fatalf("Failed to fetch linode types: %s", err)
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

func Test_GetTypes(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey)
	client := lingo.NewLinodeClient(api)

	types, err := client.GetTypes()
	if err != nil {
		t.Fatalf("Failed to GetTypes: %s", err)
	}

	if len(types) > 0 {
		ltype, err := client.GetType(types[0].ID)
		if err != nil {
			t.Fatalf("Failed to GetType: %s", err)
		}

		if ltype.Label != types[0].Label {
			t.Fatal("Type doesn't match Types slice")
		}
	}
}

func Test_BootLinode(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey)
	client := lingo.NewLinodeClient(api)

	createLinode := lingo.NewLinode{
		Region:   "us-east-1a",
		Type:     "g5-nanode-1",
		Image:    "linode/debian9",
		RootPass: "test123",
	}

	testLinode, err := client.CreateLinode(createLinode)
	if err != nil {
		t.Fatalf("Failed to create linode: %s", err)
	}

	if err := client.BootLinode(testLinode.ID); err != nil {
		t.Fatalf("Failed to boot linode: %s", err)
	}

	if err := client.RebootLinode(testLinode.ID); err != nil {
		t.Fatalf("Failed to reboot linode: %s", err)
	}

	if err := client.ShutdownLinode(testLinode.ID); err != nil {
		t.Fatalf("Failed to shutdown linode: %s", err)
	}

	if err := client.DeleteLinode(testLinode.ID); err != nil {
		t.Fatalf("Failed to cleanup: %s", err)
	}
}

func Test_ResizeLinode(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey)
	client := lingo.NewLinodeClient(api)

	newType := "g5-standard-1"
	createLinode := lingo.NewLinode{
		Region:   "us-east-1a",
		Type:     "g5-nanode-1",
		Image:    "linode/debian9",
		RootPass: "test123",
	}

	testLinode, err := client.CreateLinode(createLinode)
	if err != nil {
		t.Fatalf("Failed to create linode: %s", err)
	}

	if err := client.ResizeLinode(testLinode.ID, newType); err != nil {
		t.Fatalf("Failed to resize linode: %s", err)
	}

	if err := client.DeleteLinode(testLinode.ID); err != nil {
		t.Fatalf("Failed to cleanup: %s", err)
	}
}

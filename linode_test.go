package lingo_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/eriktate/lingo"
)

func Test_Integration_Linodes(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey, nil)
	client := lingo.NewLinodeClient(api)

	createLinode1 := lingo.CreateLinodeRequest{
		Region:   "us-east-1a",
		Type:     "g5-nanode-1",
		Image:    "linode/debian9",
		RootPass: "test123",
	}

	createLinode2 := lingo.CreateLinodeRequest{
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

	_, err = client.ViewLinode(created1.ID)
	if err != nil {
		t.Fatalf("Failed to fetch linode1: %s", err)
	}

	if _, err := client.ListTypes(); err != nil {
		t.Fatalf("Failed to fetch linode types: %s", err)
	}

	linodes, err := client.ListLinodes()
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

func Test_ListTypes(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey, nil)
	client := lingo.NewLinodeClient(api)

	types, err := client.ListTypes()
	if err != nil {
		t.Fatalf("Failed to ListTypes: %s", err)
	}

	if len(types) > 0 {
		ltype, err := client.ViewType(types[0].ID)
		if err != nil {
			t.Fatalf("Failed to ViewType: %s", err)
		}

		if ltype.Label != types[0].Label {
			t.Fatal("Type doesn't match Types slice")
		}
	}
}

func Test_BootLinode(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey, nil)
	client := lingo.NewLinodeClient(api)

	createLinode := lingo.CreateLinodeRequest{
		Region:   "us-east-1a",
		Type:     "g5-nanode-1",
		Image:    "linode/debian9",
		RootPass: "test123",
		Booted:   true,
	}

	testLinode, err := client.CreateLinode(createLinode)
	if err != nil {
		t.Fatalf("Failed to create linode: %s", err)
	}

	waitUntilRunning(client, testLinode.ID)

	if err := client.ShutdownLinode(testLinode.ID); err != nil {
		t.Fatalf("Failed to shutdown linode: %s", err)
	}

	waitUntilOffline(client, testLinode.ID)

	if err := client.BootLinode(testLinode.ID); err != nil {
		t.Fatalf("Failed to boot linode: %s", err)
	}

	waitUntilRunning(client, testLinode.ID)

	if err := client.RebootLinode(testLinode.ID); err != nil {
		t.Fatalf("Failed to reboot linode: %s", err)
	}

	if err := client.DeleteLinode(testLinode.ID); err != nil {
		t.Fatalf("Failed to cleanup: %s", err)
	}
}

func Test_ResizeLinode(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey, nil)
	client := lingo.NewLinodeClient(api)

	newType := "g5-standard-1"
	createLinode := lingo.CreateLinodeRequest{
		Region:   "us-east-1a",
		Type:     "g5-nanode-1",
		Image:    "linode/debian9",
		RootPass: "test123",
		Booted:   true,
	}

	testLinode, err := client.CreateLinode(createLinode)
	if err != nil {
		t.Fatalf("Failed to create linode: %s", err)
	}

	waitUntilRunning(client, testLinode.ID)

	if err := client.ResizeLinode(testLinode.ID, newType); err != nil {
		t.Fatalf("Failed to resize linode: %s", err)
	}

	if err := client.DeleteLinode(testLinode.ID); err != nil {
		t.Fatalf("Failed to cleanup: %s", err)
	}
}

func Test_CloneLinode(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey, nil)
	client := lingo.NewLinodeClient(api)

	createLinode := lingo.CreateLinodeRequest{
		Region:   "us-east-1a",
		Type:     "g5-nanode-1",
		Image:    "linode/debian9",
		RootPass: "test123",
		Booted:   true,
	}

	log.Println("Creating linode to clone...")
	testLinode, err := client.CreateLinode(createLinode)
	if err != nil {
		t.Fatalf("Failed to create linode: %s", err)
	}

	cloneRequest := lingo.CloneLinodeRequest{
		ID:     testLinode.ID,
		Region: testLinode.Region,
		Type:   testLinode.Type,
	}

	waitUntilRunning(client, testLinode.ID)

	log.Println("Cloning linode...")
	clone, err := client.CloneLinode(cloneRequest)
	if err != nil {
		t.Fatalf("Failed to create linode: %s", err)
	}

	if err := client.DeleteLinode(testLinode.ID); err != nil {
		t.Fatalf("Failed to clean up")
	}

	if err := client.DeleteLinode(clone.ID); err != nil {
		t.Fatalf("Failed to clean up")
	}
}

func Test_RebuildLinode(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey, nil)
	client := lingo.NewLinodeClient(api)

	createLinode := lingo.CreateLinodeRequest{
		Region:   "us-east-1a",
		Type:     "g5-nanode-1",
		Image:    "linode/debian9",
		RootPass: "test123",
		Booted:   true,
	}

	log.Println("Creating linode to rebuild...")
	testLinode, err := client.CreateLinode(createLinode)
	if err != nil {
		t.Fatalf("Failed to create linode: %s", err)
	}

	rebuildRequest := lingo.RebuildLinodeRequest{
		ID:       testLinode.ID,
		Image:    "linode/centos7",
		RootPass: "test123",
	}

	waitUntilRunning(client, testLinode.ID)

	log.Println("Rebuilding linode...")
	if _, err := client.RebuildLinode(rebuildRequest); err != nil {
		t.Fatalf("Failed to rebuild linode: %s", err)
	}

	if err := client.DeleteLinode(testLinode.ID); err != nil {
		t.Fatalf("Failed to clean up")
	}

}
func waitUntilRunning(client lingo.LinodeClient, id uint) error {
	return waitUntil(client, id, lingo.StatusRunning)
}

func waitUntilOffline(client lingo.LinodeClient, id uint) error {
	return waitUntil(client, id, lingo.StatusOffline)
}

func waitUntil(client lingo.LinodeClient, id uint, status lingo.Status) error {
	linode, err := client.ViewLinode(id)
	if err != nil {
		return err
	}

	if linode.Status != status {
		log.Printf("Waiting for %s...", status)
		time.Sleep(5 * time.Second)
		waitUntil(client, id, status)
	}

	return nil
}

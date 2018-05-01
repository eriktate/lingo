package lingo_test

import (
	"log"
	"os"
	"testing"

	"github.com/eriktate/lingo"
)

func Test_Disks(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey)
	client := lingo.NewDiskClient(api)
	linodeClient := lingo.NewLinodeClient(api)

	createLinode := lingo.NewLinode{
		Region:   "us-east-1a",
		Type:     "g5-standard-2",
		Image:    "linode/debian9",
		RootPass: "test123",
		Booted:   true,
	}

	log.Println("Creating linode to add disks...")
	testLinode, err := linodeClient.CreateLinode(createLinode)
	if err != nil {
		t.Fatalf("Failed to create linode: %s", err)
	}

	newDisk1 := lingo.NewDisk{
		LinodeID:   testLinode.ID,
		Size:       128,
		Label:      "test_disk",
		FileSystem: lingo.FileSystemExt4,
	}

	newDisk2 := lingo.NewDisk{
		LinodeID: testLinode.ID,
		Size:     512,
		Image:    "linode/debian9",
		RootPass: "test321",
	}

	waitUntilRunning(linodeClient, testLinode.ID)

	if _, err := client.CreateDisk(newDisk1); err != nil {
		t.Fatalf("Failed to create disk: %s", err)
	}

	if _, err := client.CreateDisk(newDisk2); err != nil {
		t.Fatalf("Failed to create disk: %s", err)
	}

	disks, err := client.GetDisks(testLinode.ID)
	if err != nil {
		t.Fatalf("Failed to get disks: %s", err)
	}

	log.Printf("Disks: %+v", disks)
}

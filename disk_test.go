package lingo_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/eriktate/lingo"
)

func Test_Disks(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey, nil)
	client := lingo.NewDiskClient(api)
	linodeClient := lingo.NewLinodeClient(api)

	createLinode := lingo.CreateLinodeRequest{
		Region:   "us-east-1a",
		Type:     "g5-standard-2",
		Image:    "linode/debian9",
		RootPass: "test123",
		Booted:   true,
	}

	// log.Println("Creating linode to add disks...")
	testLinode, err := linodeClient.CreateLinode(createLinode)
	if err != nil {
		t.Fatalf("Failed to create linode: %s", err)
	}

	newDisk1 := lingo.CreateDiskRequest{
		LinodeID:   testLinode.ID,
		Size:       128,
		Label:      "test_disk",
		FileSystem: lingo.FileSystemExt4,
	}

	newDisk2 := lingo.CreateDiskRequest{
		LinodeID: testLinode.ID,
		Size:     512,
		Image:    "linode/debian9",
		RootPass: "test321",
	}

	waitUntilRunning(linodeClient, testLinode.ID)

	disks, err := client.ListDisks(testLinode.ID)
	if err != nil {
		t.Fatalf("Failed to get disks: %s", err)
	}

	disk1, err := client.CreateDisk(newDisk1)
	if err != nil {
		t.Fatalf("Failed to create disk: %s", err)
	}

	disk2, err := client.CreateDisk(newDisk2)
	if err != nil {
		t.Fatalf("Failed to create disk: %s", err)
	}

	log.Printf("Disks: %+v", disks)

	updateReq := lingo.UpdateDiskRequest{
		LinodeID: testLinode.ID,
		ID:       disk1.ID,
		Label:    "This is a meh label",
	}

	if _, err := client.UpdateDisk(updateReq); err != nil {
		t.Fatalf("Failed to update disk: %s", err)
	}

	if err := client.DeleteDisk(testLinode.ID, disk1.ID); err != nil {
		t.Fatalf("Failed to clean up disk1: %s", err)
	}

	if err := client.DeleteDisk(testLinode.ID, disk2.ID); err != nil {
		t.Fatalf("Failed to clean up disk2: %s", err)
	}
}

func Test_ResizeDisk(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey, nil)
	client := lingo.NewDiskClient(api)
	linodeClient := lingo.NewLinodeClient(api)

	createLinode := lingo.CreateLinodeRequest{
		Region:   "us-east-1a",
		Type:     "g5-standard-2",
		Image:    "linode/debian9",
		RootPass: "test123",
		Booted:   true,
	}

	// log.Println("Creating linode to add disks...")
	testLinode, err := linodeClient.CreateLinode(createLinode)
	if err != nil {
		t.Fatalf("Failed to create linode: %s", err)
	}

	waitUntilRunning(linodeClient, testLinode.ID)

	disks, err := client.ListDisks(testLinode.ID)
	if err != nil {
		t.Fatalf("Failed to get disks: %s", err)
	}

	if err := linodeClient.ShutdownLinode(testLinode.ID); err != nil {
		t.Fatalf("Failed to shutdown linode: %s", err)
	}

	waitUntilOffline(linodeClient, testLinode.ID)

	// Find the largest disk so we can shrink it.
	var largest lingo.Disk
	if len(disks) > 0 {
		largest = disks[0]
		for _, d := range disks {
			if d.Size > largest.Size {
				largest = d
			}
		}
	}

	newSize := uint(20000)
	if _, err := client.ResizeDisk(testLinode.ID, largest.ID, newSize); err != nil {
		t.Fatalf("Failed to resize disk: %s", err)
	}

	disk, err := waitUntilResize(client, testLinode.ID, largest.ID, largest.Size)
	if err != nil {
		t.Fatalf("Something went wrong while resizing disk: %s", err)
	}

	if disk.Size != newSize {
		t.Fatalf("New size of %d does not match requested size of %d", disk.Size, newSize)
	}
}

func waitUntilResize(client lingo.DiskClient, linodeID, diskID, previousSize uint) (lingo.Disk, error) {
	disk, err := client.ViewDisk(linodeID, diskID)
	if err != nil {
		return disk, err
	}

	if disk.Size == previousSize {
		log.Printf("Waiting for disk resize...")
		time.Sleep(5 * time.Second)
		waitUntilResize(client, linodeID, diskID, previousSize)
	}

	return disk, nil
}

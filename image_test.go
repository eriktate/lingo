package lingo_test

import (
	"os"
	"testing"

	"github.com/eriktate/lingo"
)

func Test_ListImages(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey, nil)
	client := lingo.NewImageClient(api)

	if _, err := client.ListImages(); err != nil {
		t.Fatalf("Failed to GetImages: %s", err)
	}
}

// func Test_ViewImage(t *testing.T) {
// 	apiKey := os.Getenv("LINODE_API_KEY")
// 	api := lingo.NewAPIClient(apiKey)
// 	client := lingo.NewImageClient(api)
// }

func Test_Image(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey, nil)
	client := lingo.NewImageClient(api)
	linodeClient := lingo.NewLinodeClient(api)
	diskClient := lingo.NewDiskClient(api)

	existing, err := client.ListImages()
	if err != nil {
		t.Fatalf("Failed to fetch existing images: %s", err)
	}

	createLinode := lingo.CreateLinodeRequest{
		Region:   "us-east-1a",
		Type:     "g5-nanode-1",
		Image:    "linode/debian9",
		RootPass: "test123",
		Booted:   true,
	}

	linode, err := linodeClient.CreateLinode(createLinode)
	if err != nil {
		t.Fatalf("Failed to create test linode: %s", err)
	}

	waitUntilRunning(linodeClient, linode.ID)

	disks, err := diskClient.ListDisks(linode.ID)
	if err != nil {
		t.Fatalf("Failed to fetch disks: %s", err)
	}

	var diskID uint
	if len(disks) > 0 {
		diskID = disks[0].ID
	}

	imageReq := lingo.CreateImageRequest{
		DiskID:      diskID,
		Label:       "test-image",
		Description: "This is a test",
	}

	image, err := client.CreateImage(imageReq)
	if err != nil {
		t.Fatalf("Failed to create image: %s", err)
	}

	updateReq := lingo.UpdateImageRequest{
		ID:          image.ID,
		Description: "This is ALSO a test",
	}

	if _, err := client.UpdateImage(updateReq); err != nil {
		t.Fatalf("Failed to update image: %s", err)
	}

	getImage, err := client.ViewImage(image.ID)
	if err != nil {
		t.Fatalf("Failed to view image: %s", err)
	}

	if getImage.Description == image.Description {
		t.Fatalf("Updates were not applied to image. Expected %s, but got %s", updateReq.Description, getImage.Description)
	}

	images, err := client.ListImages()
	if err != nil {
		t.Fatalf("Failed to list images: %s", err)
	}

	expected := (len(existing) + 1)
	if len(images) != expected {
		t.Fatalf("Image listing returned unexpected results. Expected %d images, but got %d", expected, len(images))
	}

	if err := client.DeleteImage(image.ID); err != nil {
		t.Fatalf("Failed to delete image: %s", err)
	}

	if err := linodeClient.DeleteLinode(linode.ID); err != nil {
		t.Fatalf("Failed to cleanup linode: %s", err)
	}
}

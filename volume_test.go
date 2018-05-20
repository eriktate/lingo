package lingo_test

import (
	"os"
	"testing"

	"github.com/eriktate/lingo"
)

func Test_Volume(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey)
	client := lingo.NewVolumeClient(api)

	existing, err := client.ListVolumes()
	if err != nil {
		t.Fatalf("Failed to fetch existing volumes: %s", err)
	}

	createReq := lingo.CreateVolumeRequest{
		Label:  "test",
		Size:   20,
		Region: "us-east",
	}

	volume, err := client.CreateVolume(createReq)
	if err != nil {
		t.Fatalf("Failed to create volume: %s", err)
	}

	updateReq := lingo.UpdateVolumeRequest{
		ID:    volume.ID,
		Label: "updated-test",
	}

	if _, err := client.UpdateVolume(updateReq); err != nil {
		t.Fatalf("Failed to update volume: %s", err)
	}

	newSize := uint(40)
	if err := client.ResizeVolume(volume.ID, newSize); err != nil {
		t.Fatalf("Failed to resize volume: %s", err)
	}

	getVolume, err := client.ViewVolume(volume.ID)
	if err != nil {
		t.Fatalf("Failed to view volume: %s", err)
	}

	if getVolume.Label != updateReq.Label {
		t.Fatalf("Update not applied. Expected %s but got %s", updateReq.Label, getVolume.Label)
	}

	if getVolume.Size != newSize {
		t.Fatalf("Resize not applied. Expected %d but got %d", newSize, getVolume.Size)
	}

	// TODO: Not quite sure how to test this. Cloning is apparently not very instant, and you
	// don't get an ID back.
	// cloneReq := lingo.UpdateVolumeRequest{
	// 	ID:    volume.ID,
	// 	Label: "cloned-test",
	// }

	// if err := client.CloneVolume(cloneReq); err != nil {
	// 	t.Fatalf("Failed to clone volume: %s", err)
	// }

	volumes, err := client.ListVolumes()
	if err != nil {
		t.Fatalf("Failed to fetch volumes: %s", err)
	}

	expected := len(existing) + 1
	if len(volumes) != expected {
		t.Fatalf("Something went wrong. Total number of volumes is %d, but expected %d", len(volumes), expected)
	}

	if err := client.DeleteVolume(volume.ID); err != nil {
		t.Fatalf("Failed to delete volume: %s", err)
	}

	// if err := client.DeleteVolume(clone.ID); err != nil {
	// 	t.Fatalf("Failed to delete volume clone: %s", err)
	// }
}

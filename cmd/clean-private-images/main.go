package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/eriktate/lingo"
)

func main() {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey)
	client := lingo.NewImageClient(api)

	images, err := client.ListImages()
	if err != nil {
		log.Fatalf("Failed to get images: %s", err)
	}

	for _, image := range images {
		if strings.HasPrefix(image.ID, "private/") {
			if err := client.DeleteImage(image.ID); err != nil {
				log.Fatalf("Failed to delete private image with ID %s: %s", image.ID, err)
			}
		}
	}

	// Verify new image list
	cleanedImages, err := client.ListImages()
	if err != nil {
		log.Fatalf("Failed to verify images: %s", err)
	}

	data, err := json.Marshal(&cleanedImages)
	if err != nil {
		log.Fatalf("Failed to marshal images: %s", err)
	}

	log.Println("Image List JSON:")
	fmt.Println(string(data))
}

package lingo_test

import (
	"os"
	"testing"

	"github.com/eriktate/lingo"
)

func Test_GetImages(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey)
	client := lingo.NewImageClient(api)

	if _, err := client.GetImages(); err != nil {
		t.Fatalf("Failed to GetImages: %s", err)
	}
}

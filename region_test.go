package lingo_test

import (
	"os"
	"testing"

	"github.com/eriktate/lingo"
)

func Test_Integration_Regions(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	client := lingo.NewClient(apiKey)

	if _, err := client.GetRegions(); err != nil {
		t.Fatalf("Failed to GetRegions: %s", err)
	}

	if _, err := client.GetRegion("ap-northeast"); err != nil {
		t.Fatalf("Failed to GetRegion: %s", err)
	}
}

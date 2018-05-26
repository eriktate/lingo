package lingo_test

import (
	"os"
	"testing"

	"github.com/eriktate/lingo"
)

func Test_Regions(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey, nil)
	client := lingo.NewRegionClient(api)

	if _, err := client.ListRegions(); err != nil {
		t.Fatalf("Failed to list regions: %s", err)
	}

	if _, err := client.ViewRegion("ap-northeast"); err != nil {
		t.Fatalf("Failed to get region: %s", err)
	}
}

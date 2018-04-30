package main

import (
	"log"
	"os"

	"github.com/eriktate/lingo"
)

func main() {
	if err := realMain(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func realMain() error {
	apiKey := os.Getenv("LINODE_API_KEY")
	client := lingo.NewClient(apiKey)

	images, err := client.GetImages()
	if err != nil {
		return err
	}

	log.Printf("Images: %+v", images)

	regions, err := client.GetRegions()
	if err != nil {
		return err
	}

	log.Printf("Regions: %+v", regions)

	region, err := client.GetRegion("ap-northeast")
	if err != nil {
		return err
	}

	log.Printf("Region: %+v", region)
	return nil
}

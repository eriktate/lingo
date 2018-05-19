package lingo_test

import (
	"os"
	"testing"

	"github.com/eriktate/lingo"
)

func Test_CRUDDomain(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey)
	client := lingo.NewDomainClient(api)

	newDomain1 := lingo.Domain{
		Domain: "testdomain.io",
		Type:   lingo.DomainTypeMaster,
		SOA:    "test@otherdomain.com",
	}

	newDomain2 := lingo.Domain{
		Domain: "newdomain.io",
		Type:   lingo.DomainTypeSlave,
		SOA:    "test@otherdomain.com",
	}

	domain1, err := client.CreateDomain(newDomain1)
	if err != nil {
		t.Fatalf("Failed to create domain 1: %s", err)
	}

	domain2, err := client.CreateDomain(newDomain2)
	if err != nil {
		t.Fatalf("Failed to create domain 2: %s", err)
	}

	updateDomain := domain1
	updateDomain.Description = "UPDATED"
	if _, err := client.UpdateDomain(updateDomain); err != nil {
		t.Fatalf("Failed to update domain: %s", err)
	}

	getDomain, err := client.ViewDomain(domain1.ID)
	if err != nil {
		t.Fatalf("Failed to view domain: %s", err)
	}

	if getDomain.Description != updateDomain.Description {
		t.Fatal("Update of domain didn't actually occur")
	}

	if err := client.DeleteDomain(domain1.ID); err != nil {
		t.Fatalf("Failed to delete domain 1: %s", err)
	}

	if err := client.DeleteDomain(domain2.ID); err != nil {
		t.Fatalf("Failed to delete domain 2: %s", err)
	}
}

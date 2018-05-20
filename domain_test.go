package lingo_test

import (
	"os"
	"testing"

	"github.com/eriktate/lingo"
)

func Test_CRUDDomain(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey, nil)
	client := lingo.NewDomainClient(api)

	existing, err := client.ListDomains()
	if err != nil {
		t.Fatalf("Failed to list domains: %s", err)
	}

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

	domains, err := client.ListDomains()
	if err != nil {
		t.Fatalf("Failed to list domains: %s", err)
	}

	expected := len(existing) + 2
	if len(domains) != expected {
		t.Fatalf("Something strange happened. Expected to list %d domains, but got %d", expected, len(domains))
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

func Test_CRUDDomainRecord(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey, nil)
	client := lingo.NewDomainClient(api)

	newDomain := lingo.Domain{
		Domain: "testdomain.io",
		Type:   lingo.DomainTypeMaster,
		SOA:    "test@otherdomain.com",
	}

	domain, err := client.CreateDomain(newDomain)
	if err != nil {
		t.Fatalf("Failed to create domain: %s", err)
	}

	newRecord1 := lingo.DomainRecord{
		Type:   lingo.CNAME,
		Name:   "www.testdomain.io",
		Target: "testdomain.io",
	}

	newRecord2 := lingo.DomainRecord{
		Type:   lingo.A,
		Target: "127.0.0.1",
	}

	record1, err := client.CreateDomainRecord(domain.ID, newRecord1)
	if err != nil {
		t.Fatalf("Failed to create domain record 1: %s", err)
	}

	record2, err := client.CreateDomainRecord(domain.ID, newRecord2)
	if err != nil {
		t.Fatalf("Failed to create domain record 2: %s", err)
	}

	updateRecord := record1
	updateRecord.Name = "UPDATED"
	if _, err := client.UpdateDomainRecord(domain.ID, updateRecord); err != nil {
		t.Fatalf("Failed to update domain record: %s", err)
	}

	getRecord, err := client.ViewDomainRecord(domain.ID, record1.ID)
	if err != nil {
		t.Fatalf("Failed to view domain record: %s", err)
	}

	if getRecord.Name != updateRecord.Name {
		t.Fatal("Update of domain record didn't actually occur")
	}

	if err := client.DeleteDomainRecord(domain.ID, record1.ID); err != nil {
		t.Fatalf("Failed to delete domain record 1: %s", err)
	}

	if err := client.DeleteDomainRecord(domain.ID, record2.ID); err != nil {
		t.Fatalf("Failed to delete domain record 2: %s", err)
	}

	if err := client.DeleteDomain(domain.ID); err != nil {
		t.Fatalf("Failed to cleanup domain: %s", err)
	}
}

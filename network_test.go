package lingo_test

import (
	"os"
	"testing"

	"github.com/eriktate/lingo"
)

func Test_CRUDNetwork(t *testing.T) {
	apiKey := os.Getenv("LINODE_API_KEY")
	api := lingo.NewAPIClient(apiKey, nil)
	client := lingo.NewNetworkClient(api)
	linodeClient := lingo.NewLinodeClient(api)

	existing, err := client.ListAddresses()
	if err != nil {
		t.Fatalf("Failed to list existing addresses: %s", err)
	}

	createLinode := lingo.CreateLinodeRequest{
		Region:   "us-east-1a",
		Type:     "g5-standard-2",
		Image:    "linode/debian9",
		RootPass: "test123",
		Booted:   true,
	}

	testLinode, err := linodeClient.CreateLinode(createLinode)
	if err != nil {
		t.Fatalf("Failed to create linode: %s", err)
	}

	allocateRequest := lingo.AllocateAddressRequest{
		LinodeID: testLinode.ID,
		Type:     lingo.IPv4,
		Public:   false,
	}

	if _, err := client.AllocateAddress(allocateRequest); err != nil {
		t.Fatalf("Failed to allocate address: %s", err)
	}

	// TODO: Figure out a way to test this.
	// rdnsRequest := lingo.UpdateRDNSRequest{
	// 	Address: testLinode.IPv4[0],
	// 	RDNS:    "test.example.org",
	// }

	// if _, err := client.UpdateAddressRDNS(rdnsRequest); err != nil {
	// 	t.Fatalf("Failed to update rdns: %s", err)
	// }

	if _, err := client.ViewAddress(testLinode.IPv4[0]); err != nil {
		t.Fatalf("Failed to view address: %s", err)
	}

	// if getAddr.RDNS != rdnsRequest.RDNS {
	// 	t.Fatalf("Update failed to apply. Expected rdns to be %s, but got %s", rdnsRequest.RDNS, getAddr.RDNS)
	// }

	addrs, err := client.ListAddresses()
	if err != nil {
		t.Fatalf("Failed to list addresses: %s", err)
	}

	expected := len(existing) + 1
	if len(addrs) != expected {
		t.Fatalf("Something strange happened. Expected to list %d addresses, but got %d", expected, len(addrs))
	}
}

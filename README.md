# lingo
[![GoDoc](https://godoc.org/github.com/eriktate/lingo?status.svg)](http://godoc.org/github.com/eriktate/lingo)
Go client library for the Linode API.

## WIP
This project is under very active development and is definitely *NOT* safe for production use! It's probably safe to use in hobby projects if the API you need happens to be finished.

## Getting Started
The library is built on top of a core API client that knows how to make HTTP requests to the Linode API. The rest is divided into many small, friendlier clients, each one concerning itself with a particular resource in the Linode ecosystem. The `Lingo` struct is just an aggregation of all of the smaller clients, making it a one stop shop for any API call you might want to make.

The simplest way to get started looks like this:
```go
import (
	"log"

	"github.com/eriktate/lingo"
)

func main() {
	apiKey := "your-linode-api-key-here"
	// In place of nil you can specify a backoff configuration.
	linode := lingo.NewLingo(apiKey, nil)

	// Most API functions take a parameter struct.
	createLinodeRequest := lingo.CreateLinodeRequest{
		Region:   "us-east-1a",
		Type:     "g5-nanode-1",
		Image:    "linode/debian9",
		RootPass: "test123",
	}

	newLinode, err := linode.CreateLinode(createLinodeRequest)
	if err != nil {
		log.Fatalf("Something went wrong while creating linode: %s", err)
	}

	log.Printf("New linode created with ID: %d", newLinode.ID)
}

```

If you only want to work with a particular API, you can do so:
```go
import (
	"log"

	"github.com/eriktate/lingo"
)

func main() {
	apiKey := "your-linode-api-key-here"
	api := lingo.NewAPIClient(apiKey, nil)
	domain := lingo.DomainClient(api)

	createDomainRequest := lingo.CreateDomainRequest{
		Domain: "testdomain.io",
		Type:   lingo.DomainTypeMaster,
		SOA:    "test@testdomain.io",
	}

	newDomain, err := domain.CreateDomain(createDomainRequest)
	if err != nil {
		log.Fatalf("Something went wrong while creating domain: %s", err)
	}

	log.Printf("New domain created with ID: %d", newDomain.ID)
}
```

## Completed APIs
- Domain
- Image
- Linode Types
- Regions
- Volume
- Networking

## Partial APIs
- Linode Instance
- NodeBalancer
- Disk

## TODO APIs
- LongView
- StackScripts (?)
- Profile (?)
- Account (?)

## Related Projects
I'm currently working on building a linode provider for terraform using this client library. You can check that out at https://github.com/eriktate/terraform-provider-linode

package lingo

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// A NetworkClient is an API struct that can make requests related to Linode network configurations.
type NetworkClient struct {
	api APIClient
}

// NewNetworkClient returns a new NetworkClient given a valid APIClient.
func NewNetworkClient(api APIClient) NetworkClient {
	return NetworkClient{api: api}
}

// ListAddresses retrieves a slice of network addresses currently in use.
func (c NetworkClient) ListAddresses() ([]Address, error) {
	data, err := c.api.Get("networking/ips")
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request for ListAddresses")
	}

	var results Results
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, errors.Wrap(err, "failed to decode ListAddresses response")
	}

	// TODO: Do something with paging here?
	var addresses []Address
	if err := json.Unmarshal(results.Data, &addresses); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal ListAddresses data")
	}

	return addresses, nil
}

// ViewAddress retrieves a slice of machine images available in Linode.
func (c NetworkClient) ViewAddress(address string) (Address, error) {
	var ip Address
	data, err := c.api.Get(fmt.Sprintf("networking/ips/%s", address))
	if err != nil {
		return ip, errors.Wrap(err, "failed to make request for ViewAddress")
	}

	if err := json.Unmarshal(data, &ip); err != nil {
		return ip, errors.Wrap(err, "failed to unmarshal ViewAddress data")
	}

	return ip, nil
}

// AllocateAddress allocates a new IP address to a specified Linode instance.
func (c NetworkClient) AllocateAddress(req AllocateAddressRequest) (Address, error) {
	var ip Address
	payload, err := json.Marshal(req)
	if err != nil {
		return ip, errors.Wrap(err, "failed to marshal request for AllocateAddress")
	}

	data, err := c.api.Post("networking/ips", payload)
	if err != nil {
		return ip, errors.Wrap(err, "failed to make request for AllocateAddress")
	}

	if err := json.Unmarshal(data, &ip); err != nil {
		return ip, errors.Wrap(err, "failed to decode AllocateAddress response")
	}

	return ip, nil
}

// AssignAddress assigns a set of existing IP addresses to a set of existing Linode instances.
func (c NetworkClient) AssignAddress(req AssignAddressRequest) error {
	payload, err := json.Marshal(req)
	if err != nil {
		return errors.Wrap(err, "failed to marshal request for AssignAddress")
	}

	if _, err := c.api.Post("networking/ipv4/assign", payload); err != nil {
		return errors.Wrap(err, "failed to make request for AssignAddress")
	}

	return nil
}

// UpdateAddressRDNS updates the RDNS configuration of an existing IP address.
func (c NetworkClient) UpdateAddressRDNS(req UpdateRDNSRequest) (Address, error) {
	var ip Address
	payload, err := json.Marshal(req)
	if err != nil {
		return ip, errors.Wrap(err, "failed to marshal request for UpdateAddressRDNS")
	}

	data, err := c.api.Put(fmt.Sprintf("networking/ips/%s", req.Address), payload)
	if err != nil {
		return ip, errors.Wrap(err, "failed to make request for UpdateAddressRDNS")
	}

	if err := json.Unmarshal(data, &ip); err != nil {
		return ip, errors.Wrap(err, "failed to decode UpdateAddressRDNS response")
	}

	return ip, nil
}

// ConfigureSharing configures IP sharing for the specified IPs and Linode instance.
func (c NetworkClient) ConfigureSharing(req SharingRequest) error {
	payload, err := json.Marshal(req)
	if err != nil {
		return errors.Wrap(err, "failed to marshal request for ConfigureSharing")
	}

	if _, err := c.api.Post("networking/ipv4/share", payload); err != nil {
		return errors.Wrap(err, "failed to make request for ConfigureSharing")
	}

	return nil
}

// ListIPv6Pools retrieves a slice of IPv6 pools currently in use.
func (c NetworkClient) ListIPv6Pools() ([]IPv6Pool, error) {
	data, err := c.api.Get("networking/ipv6/pools")
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request for ListIPv6Pools")
	}

	var results Results
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, errors.Wrap(err, "failed to decode ListIPv6Pools response")
	}

	// TODO: Do something with paging here?
	var pools []IPv6Pool
	if err := json.Unmarshal(results.Data, &pools); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal ListIPv6Pools data")
	}

	return pools, nil
}

// ListIPv6Ranges retrieves a slice of IPv6 ranges currently in use.
func (c NetworkClient) ListIPv6Ranges() ([]IPv6Range, error) {
	data, err := c.api.Get("networking/ipv6/ranges")
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request for ListIPv6Ranges")
	}

	var results Results
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, errors.Wrap(err, "failed to decode ListIPv6Ranges response")
	}

	// TODO: Do something with paging here?
	var ranges []IPv6Range
	if err := json.Unmarshal(results.Data, &ranges); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal ListIPv6Ranges data")
	}

	return ranges, nil
}

package lingo

// An AddressType is an enumeration of possible network address types.
type AddressType string

// Enum values for AddressType.
const (
	IPv4  = AddressType("ipv4")
	IPv6  = AddressType("ipv6")
	Pool  = AddressType("ipv6/pool")
	Range = AddressType("ipv6/range")
)

// AllocateAddressRequest is a parameter struct for allocating new IP addresses.
type AllocateAddressRequest struct {
	LinodeID uint        `json:"linode_id"`
	Type     AddressType `json:"type"`
	Public   bool        `json:"public"`
}

// UpdateRDNSRequest is a parameter struct for updating the RDNS config for an IP address.
type UpdateRDNSRequest struct {
	Address string `json:"-"`
	RDNS    string `json:"rdns"`
}

// AssignAddressRequest is a parameter struct for assigning IP addresses to Linode instances.
type AssignAddressRequest struct {
	Region      string       `json:"region"`
	Assignments []Assignment `json:"assignments"`
}

// An Assignment represents an assignment relationship between an address and a Linode instance.
type Assignment struct {
	LinodeID uint   `json:"linode_id"`
	Address  string `json:"address"`
}

// SharingRequest is a parameter struct for sharing IP addresses.
type SharingRequest struct {
	LinodeID uint     `json:"linode_id"`
	IPs      []string `json:"ips"`
}

// An Address represents a Linode network address.
type Address struct {
	Address    string      `json:"address"`
	Gateway    string      `json:"gateway"`
	SubnetMask string      `json:"subnet_mask"`
	Prefix     uint        `json:"prefix"`
	Type       AddressType `json:"type"`
	Public     bool        `json:"public"`
	RDNS       string      `json:"rdns"`
	LinodeID   uint        `json:"linode_id"`
	Region     string      `json:"region"`
}

// An IPv6Range represents a range.
type IPv6Range struct {
	Range  string `json:"range"`
	Region string `json:"region"`
}

// An IPv6Pool represents a pool.
type IPv6Pool IPv6Range

// A Networker works with Linode network configurations.
type Networker interface {
	ListAddresses() ([]Address, error)
	ViewAddress(address string) (Address, error)
	AllocateAddress(req AllocateAddressRequest) (Address, error)
	UpdateAddressRDNS(req UpdateRDNSRequest) (Address, error)
	AssignAddress(req AssignAddressRequest) error
	ConfigureSharing(req SharingRequest) error
	ListIPv6Pools() ([]IPv6Range, error)
	ListIPv6Range() ([]IPv6Pool, error)
}

package lingo

// A DomainType is an enumeration of possible Linode Domain types.
type DomainType string

// Enum values for DomainType.
const (
	DomainTypeMaster = DomainType("master")
	DomainTypeSlave  = DomainType("slave")
)

// A DomainStatus is an enumeration of possible Linode Domain statuses.
type DomainStatus string

// Enum values for DomainStatus.
const (
	DomainStatusDisabled  = DomainStatus("disabled")
	DomainStatusActive    = DomainStatus("active")
	DomainStatusEditMode  = DomainStatus("edit_mode")
	DomainStatusHasErrors = DomainStatus("has_errors")
)

// A Domain represents a Linode Domain.
// TODO: Write custom unmarshaler so that domains that don't fit the proper regex error.
type Domain struct {
	ID          uint         `json:"id"`
	Domain      string       `json:"domain"`
	Type        DomainType   `json:"type"`
	Status      DomainStatus `json:"status,omitempty"`
	Description string       `json:"description,omitempty"`
	TTLSec      uint         `json:"ttl_sec,omitempty"`
	RetrySec    uint         `json:"retry_sec,omitempty"`
	MasterIPs   []string     `json:"master_ips,omitempty"`
	AxfrIPs     []string     `json:"axfr_ips,omitempty"`
	ExpireSec   uint         `json:"expire_sec,omitempty"`
	RefreshSec  uint         `json:"refresh_sec,omitempty"`
	SOA         string       `json:"soa_email,omitempty"`
}

// A DomainRecordType is an enumeration of possible Linode Domain Record types.
type DomainRecordType string

// Enum values for DomainRecordType.
const (
	CNAME = DomainRecordType("CNAME")
	A     = DomainRecordType("A")
	AAAA  = DomainRecordType("AAAA")
	NS    = DomainRecordType("NS")
	MX    = DomainRecordType("MX")
	TXT   = DomainRecordType("TXT")
	SRV   = DomainRecordType("SRV")
	PTR   = DomainRecordType("PTR")
	CAA   = DomainRecordType("CAA")
)

// A DomainRecord represetns a Linode Domain Record.
type DomainRecord struct {
	ID       uint             `json:"id"`
	Name     string           `json:"name"`
	Target   string           `json:"target"`
	Priority uint8            `json:"priority"`
	Type     DomainRecordType `json:"type"`
	Port     uint             `json:"port"`
	Service  string           `json:"service,omitempty"`
	Protocol string           `json:"protocol,omitempty"`
	TTLSec   uint             `json:"ttl_sec"`
	Tag      string           `json:"tag,omitempty"`
}

// A Domainer works with Linode Domains and Domain Records.
type Domainer interface {
	ListDomains() ([]Domain, error)
	ViewDomain(id uint) (Domain, error)
	CreateDomain(domain Domain) (Domain, error)
	UpdateDomain(domain Domain) (Domain, error)
	DeleteDomain(id uint) error

	ListDomainRecord(domainID uint) ([]DomainRecord, error)
	ViewDomainRecord(domainID, recordID uint) (DomainRecord, error)
	CreateDomainRecord(domainID uint, record DomainRecord) (DomainRecord, error)
	UpdateDomainRecord(domainID uint, record DomainRecord) (DomainRecord, error)
	DeleteDomainRecord(domainID, recordID uint) error
}

// ValidateDomainType validates whether or not a test string is a DomainType enum.
func ValidateDomainType(test string) bool {
	switch DomainType(test) {
	case DomainTypeMaster, DomainTypeSlave:
		return true
	default:
		return false
	}
}

// ValidateDomainStatus validates whether or not a test string is a DomainStatus enum.
func ValidateDomainStatus(test string) bool {
	switch DomainStatus(test) {
	case DomainStatusActive, DomainStatusDisabled, DomainStatusEditMode, DomainStatusHasErrors:
		return true
	default:
		return false
	}
}

// ValidateDomainRecordType validates whether or not a test string is a DomainRecordType enum.
func ValidateDomainRecordType(test string) bool {
	switch DomainRecordType(test) {
	case CNAME, A, AAAA, NS, MX, TXT, SRV, PTR, CAA:
		return true
	default:
		return false
	}
}

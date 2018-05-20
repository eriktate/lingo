package lingo

import "encoding/json"

// Status is an enum of possible instances statuses.
type Status string

const (
	StatusOffline      = Status("offline")
	StatusBooting      = Status("booting")
	StatusRunning      = Status("running")
	StatusShuttingDown = Status("shutting_down")
	StatusRebooting    = Status("rebooting")
	StatusProvisioning = Status("provisioning")
	StatusDeleting     = Status("deleting")
	StatusMigrating    = Status("migrating")
)

// Hypervisor is an enum of possible hypervisors to be used for an instance.
type Hypervisor string

const (
	HypervisorKVM = Hypervisor("kvm")
	HypervisorXen = Hypervisor("xen")
)

// An Alerts struct specifies what the alerting bounds for particular metrics should be.
type Alerts struct {
	CPU           uint `json:"cpu"`
	IO            uint `json:"io"`
	NetworkIn     uint `json:"network_in"`
	NetworkOut    uint `json:"network_out"`
	TransferQuota uint `json:"transfer_quota"`
}

// A Specs struct represents the hardware specification for a given instance.
type Specs struct {
	Disk     uint `json:"disk"`
	Memory   uint `json:"memory"`
	Vcpus    uint `json:"vcpus"`
	Transfer uint `json:"transfer"`
}

// A Linode represents a Linode instance.
// TODO: Add Backup field
type Linode struct {
	ID         uint       `json:"id"`
	Alerts     Alerts     `json:"alerts"`
	Region     string     `json:"region"`
	Image      string     `json:"image,omitempty"`
	IPv4       []string   `json:"ipv4"`
	IPv6       string     `json:"ipv6"`
	Label      string     `json:"label,omitempty"`
	Type       string     `json:"type"`
	Status     Status     `json:"status"`
	Hypervisor Hypervisor `json:"hypervisor"`
	Specs      Specs      `json:"specs"`
	Created    Time       `json:"created"`
	Updatd     Time       `json:"updated"`
}

// CreateLinodeRequest is a paremeter struct h
type CreateLinodeRequest struct {
	Region          string          `json:"region"`
	Type            string          `json:"type"`
	Label           string          `json:"label,omitempty"`
	RootPass        string          `json:"root_pass,omitempty"`
	AuthorizedKeys  []string        `json:"authorized_keys,omitempty"`
	StackScriptID   uint            `json:"stackscript_id,omitempty"`
	StackscriptData json.RawMessage `json:"stackscript_data,omitempty"`
	BackupID        uint            `json:"backup_id,omitempty"`
	Image           string          `json:"image,omitempty"`
	BackupsEnabled  bool            `json:"backups_enabled"`
	Booted          bool            `json:"booted"`
	SwapSize        uint            `json:"swap_size,omitempty"`
}

// UpdateLinodeRequest is a parameter struct for specifying how to update an existing instance.
type UpdateLinodeRequest struct {
	ID     uint   `json:"-"`
	Label  string `json:"label,omitempty"`
	Alerts Alerts `json:"alerts,omitempty"`
}

// CloneLinodeRequest is a parameter struct for specifying a clone to be created.
type CloneLinodeRequest struct {
	ID             uint     `json:"-"`
	Region         string   `json:"region"`
	Type           string   `json:"type"`
	Label          string   `json:"label,omitempty"`
	LinodeID       uint     `json:"linode_id,omitempty"`
	BackupsEnabled bool     `json:"backups_enabled"`
	Disks          []string `json:"disks,omitempty"`
	Configs        []string `json:"configs,omitempty"`
}

// RebuildLinodeRequest is a parameter struct for specifying how a rebuild should be executed.
type RebuildLinodeRequest struct {
	ID              uint            `json:"-"`
	Image           string          `json:"image"`
	RootPass        string          `json:"root_pass"`
	AuthorizedKeys  []string        `json:"authorized_keys,omitempty"`
	StackScriptID   uint            `json:"stackscript_id,omitempty"`
	StackscriptData json.RawMessage `json:"stackscript_data,omitempty"`
	Booted          bool            `json:"booted"`
}

// A Class is an enum of possible instance classes.
type Class string

const (
	ClassNanode   = Class("nanode")
	ClassStandard = Class("standard")
	ClassHighmem  = Class("highmem")
)

// A Price indicates the monthly and hourly costs for a particular instance type.
type Price struct {
	Hourly  float32 `json:"hourly"`
	Monthly float32 `json:"monthly"`
}

// Addons represent what addons are included with an instance type.
type Addons struct {
	Backups struct {
		Price Price `json:"price"`
	} `json:"backups"`
}

// A LinodeType represents a Linode instance type.
type LinodeType struct {
	ID         string `json:"id"`
	Disk       int    `json:"disk"`
	Class      Class  `json:"class"`
	Price      Price  `json:"price"`
	Label      string `json:"label"`
	Addons     Addons `json:"addons"`
	NetworkOut uint   `json:"network_out"`
	Memory     uint   `json:"memory"`
	Transfer   uint   `json:"transfer"`
	Vcpus      uint   `json:"vcpus"`
}

// A Linoder works with Linode instances.
type Linoder interface {
	ListLinodes() ([]Linode, error)
	ViewLinode(id uint) (Linode, error)
	CreateLinode(req CreateLinodeRequest) (Linode, error)
	UpdateLinode(req UpdateLinodeRequest)
	DeleteLinode(id uint) error
	BootLinode(id uint) error
	BootLinodeWithConfig(id, configID uint) error
	RebootLinode(id uint) error
	RebootLinodeWithConfig(id, configID uint) error
	ShutdownLinode(id uint) error
	ResizeLinode(id uint, typeID string) error
	Upgrade(id uint, typeID string) error
	CloneLinode(req CloneLinodeRequest) (Linode, error)
	RebuildLinode(req RebuildLinodeRequest) error
	ListLinodeVolumes(id uint) ([]Volume, error)
	ListTypes() ([]LinodeType, error)
	ViewType(id string) (LinodeType, error)
}

package lingo

import "encoding/json"

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

type Hypervisor string

const (
	HypervisorKVM = Hypervisor("kvm")
	HypervisorXen = Hypervisor("xen")
)

type Alerts struct {
	CPU           uint `json:"cpu"`
	IO            uint `json:"io"`
	NetworkIn     uint `json:"network_in"`
	NetworkOut    uint `json:"network_out"`
	TransferQuota uint `json:"transfer_quota"`
}

type Specs struct {
	Disk     uint `json:"disk"`
	Memory   uint `json:"memory"`
	Vcpus    uint `json:"vcpus"`
	Transfer uint `json:"transfer"`
}

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

type UpdateLinodeRequest struct {
	ID     uint   `json:"-"`
	Label  string `json:"label,omitempty"`
	Alerts Alerts `json:"alerts,omitempty"`
}

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

type RebuildLinodeRequest struct {
	ID              uint            `json:"-"`
	Image           string          `json:"image"`
	RootPass        string          `json:"root_pass"`
	AuthorizedKeys  []string        `json:"authorized_keys,omitempty"`
	StackScriptID   uint            `json:"stackscript_id,omitempty"`
	StackscriptData json.RawMessage `json:"stackscript_data,omitempty"`
	Booted          bool            `json:"booted"`
}

type Class string

const (
	ClassNanode   = Class("nanode")
	ClassStandard = Class("standard")
	ClassHighmem  = Class("highmem")
)

type Price struct {
	Hourly  float32 `json:"hourly"`
	Monthly float32 `json:"monthly"`
}

type Addons struct {
	Backups struct {
		Price Price `json:"price"`
	} `json:"backups"`
}

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
	GetTypes() ([]LinodeType, error)
	GetType(id string) (LinodeType, error)
	ResizeLinode(id uint, typeID string) error
	Mutate(id uint) error
	CloneLinode(req CloneLinodeRequest) (Linode, error)
	RebuildLinode(req RebuildLinodeRequest) error
	GetLinodeVolumes(id uint) ([]Volume, error)
}

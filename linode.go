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
	Group      string     `json:"group,omitempty"`
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

type NewLinode struct {
	Region          string          `json:"region"`
	Type            string          `json:"type"`
	Label           string          `json:"label,omitempty"`
	Group           string          `json:"group,omitempty"`
	RootPass        string          `json:"root_pass,omitempty"`
	AuthorizedKeys  []string        `json:"authorized_keys,omitempty"`
	StackScriptID   uint            `json:"stackscript_id,omitempty"`
	StackscriptData json.RawMessage `json:"stackscript_data,omitempty"`
	BackupID        uint            `json:"backup_id,omitempty"`
	Image           string          `json:"image,omitempty"`
	BackupsEnabled  bool            `json:"backups_enabled"`
	Booted          bool            `json:"booted"`
}

type Linoder interface {
	GetLinodes() ([]Linode, error)
	CreateLinode(linodw NewLinode) (Linode, error)
}

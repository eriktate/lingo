package lingo

import "encoding/json"

type DiskStatus string

const (
	DiskStatusReady    = DiskStatus("ready")
	DiskStatusNotReady = DiskStatus("not ready")
	DiskStatusUpdated  = DiskStatus("updated")
)

type FileSystem string

const (
	FileSystemRaw    = FileSystem("raw")
	FileSystemSwap   = FileSystem("swap")
	FileSystemExt3   = FileSystem("ext3")
	FileSystemExt4   = FileSystem("ext4")
	FileSystemInitrd = FileSystem("initrd")
)

type Disk struct {
	ID         uint       `json:"id"`
	Label      string     `json:"label"`
	Status     DiskStatus `json:"status"`
	Size       uint       `json:"size"`
	FileSystem FileSystem `json:"filesystem"`
	Created    Time       `json:"created"`
	Updated    Time       `json:"updated"`
}

type NewDisk struct {
	LinodeID        uint            `json:"-"`
	Size            uint            `json:"size"`
	Image           string          `json:"image,omitempty"`
	RootPass        string          `json:"root_pass,omitempty"`
	AuthorizedKeys  []string        `json:"authorized_keys,omitempty"`
	Label           string          `json:"label,omitempty"`
	FileSystem      FileSystem      `json:"filesystem,omitempty"`
	ReadOnly        bool            `json:"read_only"`
	StackscriptID   uint            `json:"stackscript_id,omitempty"`
	StackscriptData json.RawMessage `json:"stackscript_data,omitempty"`
}

type Disker interface {
	GetDisks(linodeID uint) ([]Disk, error)
	CreateDisk(newDisk NewDisk) (Disk, error)
}

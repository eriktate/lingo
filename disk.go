package lingo

import "encoding/json"

type DiskStatus string

const (
	DiskStatusReady    = DiskStatus("ready")
	DiskStatusNotReady = DiskStatus("not ready")
	DiskStatusUpdated  = DiskStatus("updated")
)

// A FileSystem is an enumeration of potential file systems a disk can be created with.
type FileSystem string

// The FileSystem enum values.
const (
	FileSystemRaw    = FileSystem("raw")
	FileSystemSwap   = FileSystem("swap")
	FileSystemExt3   = FileSystem("ext3")
	FileSystemExt4   = FileSystem("ext4")
	FileSystemInitrd = FileSystem("initrd")
)

// A Disk represents a Linode Disk.
type Disk struct {
	ID         uint       `json:"id"`
	Label      string     `json:"label"`
	Status     DiskStatus `json:"status"`
	Size       uint       `json:"size"`
	FileSystem FileSystem `json:"filesystem"`
	Created    Time       `json:"created"`
	Updated    Time       `json:"updated"`
}

// A CreateDiskRequest contains the information necessary to build a new Linode Disk.
type CreateDiskRequest struct {
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

// An UpdateDiskRequest wraps up the data that can be updated on a Linode Disk.
type UpdateDiskRequest struct {
	ID         uint       `json:"-"`
	LinodeID   uint       `json:"-"`
	Label      string     `json:"label,omitempty"`
	FileSystem FileSystem `json:"filesystem,omitempty"`
}

// A Disker describes all of the functions necessary to fulfill the Linode Disk API.
type Disker interface {
	ListDisks(linodeID uint) ([]Disk, error)
	ViewDisk(linodeID, diskID uint) (Disk, error)
	CreateDisk(req CreateDiskRequest) (Disk, error)
	UpdateDisk(req UpdateDiskRequest) (Disk, error)
	DeleteDisk(linodeID, diskID uint) error
	ResetDiskRootPassword(linodeID, diskID uint, password string) (Disk, error)
	ResizeDisk(linodeID, diskID, size uint) (Disk, error)
}

// ValidateFileSystem validates whether or not a test string is a FileSystem.
func ValidateFileSystem(test string) bool {
	switch FileSystem(test) {
	case FileSystemRaw:
		fallthrough
	case FileSystemSwap:
		fallthrough
	case FileSystemExt3:
		fallthrough
	case FileSystemExt4:
		fallthrough
	case FileSystemInitrd:
		return true
	default:
		return false
	}
}

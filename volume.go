package lingo

// VolumeStatus is an enum of possible volume statuses.
type VolumeStatus string

// VolumeStatus values
const (
	VolumeStatusCreating       = VolumeStatus("creating")
	VolumeStatusActive         = VolumeStatus("active")
	VolumeStatusResizing       = VolumeStatus("resizing")
	VolumeStatusOffline        = VolumeStatus("offline")
	VolumeStatusDeleting       = VolumeStatus("deleting")
	VolumeStatusDeleted        = VolumeStatus("deleted")
	VolumeStatusContactSupport = VolumeStatus("contact_support")
)

// A Volume represents a Linode volume.
type Volume struct {
	ID             uint         `json:"id"`
	Label          string       `json:"label"`
	Status         VolumeStatus `json:"status"`
	Size           uint         `json:"size"`
	Region         string       `json:"string"`
	Created        Time         `json:"created"`
	Updated        Time         `json:"updated"`
	LinodeID       uint         `json:"linode_id"`
	FilesystemPath string       `json:"filesystem_path"`
}

// A CreateVolumeRequest is a parameter struct for specifying a new volume.
type CreateVolumeRequest struct {
	Label    string `json:"label"`
	Size     uint   `json:"size"`
	Region   string `json:"region,omitempty"`
	LinodeID uint   `json:"linode_id,omitempty"`
	ConfigID uint   `json:"config_id,omitempty"`
}

// An UpdateVolumeRequest is a parameter struct for specifying how an existing volume should be updated.
type UpdateVolumeRequest struct {
	ID    uint   `json:"-"`
	Label string `json:"label"`
}

// An AttachVolumeRequest is a paremeter struct for specifying how an existing volume should be attached
// to a Linode instance.
type AttachVolumeRequest struct {
	ID       uint `json:"-"`
	LinodeID uint `json:"linode_id"`
	ConfigID uint `json:"config_id,omitempty"`
}

// A Volumer works with Linode volumes.
type Volumer interface {
	ListVolumes() ([]Volume, error)
	ViewVolume(id uint) (Volume, error)
	CreateVolume(req CreateVolumeRequest) (Volume, error)
	UpdateVolume(req UpdateVolumeRequest) (Volume, error)
	DeleteVolume(id uint) error
	AttachVolume(req AttachVolumeRequest) error
	CloneVolume(req UpdateVolumeRequest) error
	DetachVolume(id uint) error
	ResizeVolume(id, size uint) error
}

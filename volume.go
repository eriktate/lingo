package lingo

// VolumeStatus aliases a string for use as a pseudo enum.
type VolumeStatus string

// Possible values for VolumeStatus.
const (
	VolumeStatusCreating = VolumeStatus("creating")
	VolumeStatusActive   = VolumeStatus("active")
	VolumeStatusResizing = VolumeStatus("resizing")
	VolumeStatusOffline  = VolumeStatus("offline")
)

// A Volume
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

type Volumer interface {
	// TODO: Add some things here
}

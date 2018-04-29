package lingo

// An ImageType is an enum of possible Linode image types.
type ImageType string

// The possible image types are "manual" and "automatic"
const (
	ImageTypeManual    = ImageType("manual")
	ImageTypeAutomatic = ImageType("automatic")
)

// An Image represents a Linode machine image result.
type Image struct {
	ID          string    `json:"id"`
	Label       string    `json:"label"`
	Description string    `json:"description"`
	Type        ImageType `json:"type"`
	IsPublic    bool      `json:"is_public"`
	Size        int       `json:"size"`
	Vendor      string    `json:"vendor"`
	Deprecated  bool      `json:"deprecated"`
	CreatedBy   string    `json:"created_by"`
	Created     string    `json:"created"`
}

// A NewImageRequest contains the fields necessary to build a new image.
type NewImageRequest struct {
	DiskID      string `json:"disk_id"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

// An Imager works with Linode machine images.
type Imager interface {
	GetImages() ([]Image, error)
	GetImagesByLabel(label string) ([]Image, error)
	CreateImage(image NewImageRequest) error
}

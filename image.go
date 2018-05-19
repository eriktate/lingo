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
	Created     Time      `json:"created"`
}

// A NewImage packages up the fields required for creating a new Image.
type NewImage struct {
	DiskID      uint   `json:"disk_id"`
	Label       string `json:"label,omitempty"`
	Description string `json:"description,omitempty"`
}

// A CreateImageRequest contains the fields necessary to build a new image.
type CreateImageRequest struct {
	DiskID      uint   `json:"disk_id"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

// An UpdateImageRequest contains the fields necessary to update an existing image.
type UpdateImageRequest struct {
	ID          string `json:"-"`
	Label       string `json:"label,omitempty"`
	Description string `json:"description,omitempty"`
}

// An Imager works with Linode machine images.
type Imager interface {
	ListImages() ([]Image, error)
	ViewImage() (Image, error)
	CreateImage(req CreateImageRequest) (Image, error)
	UpdateImage(req UpdateImageRequest) error
	DeleteImage(id string) error
}

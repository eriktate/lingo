package lingo

// A Region represents a Linode deployment region.
type Region struct {
	ID      string `json:"id"`
	Country string `json:"country"`
}

// A Regioner works with Linode regions.
type Regioner interface {
	ListRegions() ([]Region, error)
	ViewRegion(id string) (Region, error)
}

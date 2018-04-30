package lingo

type Region struct {
	ID      string `json:"id"`
	Country string `json:"country"`
}

type Regioner interface {
	GetRegions() ([]Region, error)
	GetRegion(id string) (Region, error)
}

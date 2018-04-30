package lingo

type Transfer struct {
	In    float32 `json:"in"`
	Out   float32 `json:"out"`
	Total float32 `json:"total"`
}

type NodeBalancer struct {
	ID                 uint   `json:"id"`
	Label              string `json:"label"`
	Hostname           string `json:"hostname"`
	ClientConnThrottle uint   `json:"client_conn_throttle"`
	Region             string `json:"region"`
	IPV4               string `json:"ipv4"`
	IPV6               string `json:"ipv6"`
	Created            Time   `json:"created"`
	Updated            Time   `json:"updated"`
	Transfer           Transfer
}

type CreateNodeBalancerRequest struct {
	Region             string `json:"region"`
	Label              string `json:"label,omitempty"`
	ClientConnThrottle uint   `json:"client_conn_throttle,omitempty"`
}

type Balancer interface {
	GetNodeBalancers() ([]NodeBalancer, error)
	GetNodeBalancer(id string) (NodeBalancer, error)
	CreateNodeBalancer(req CreateNodeBalancerRequest) (NodeBalancer, error)
	DeleteNodeBalancer(id string) error
}

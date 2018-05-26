package lingo

// Transer wraps the data representing network metrics for a Node Balancer.
type Transfer struct {
	In    float32 `json:"in"`
	Out   float32 `json:"out"`
	Total float32 `json:"total"`
}

// A NodeBalancer represents a Linode node balancer.
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

// A CreateBalancerRequest contains the fields necessary to spin up a new node balancer.
type CreateBalancerRequest struct {
	Region             string `json:"region"`
	Label              string `json:"label,omitempty"`
	ClientConnThrottle uint   `json:"client_conn_throttle,omitempty"`
}

// An UpdateBalancerRequest contains the fields necessary to update an existing node balancer.
// TODO: Not sure if I prefer the approach of including the ID in this
// struct, or providing to functions that need it.
type UpdateBalancerRequest struct {
	ID                 uint   `json:"-"`
	Label              string `json:"label,omitempty"`
	ClientConnThrottle uint   `json:"client_conn_throttle,omitempty"`
}

// A Protocol is an enumeration of possible node balancer protocols.
type Protocol string

// Enum values for Protocol.
const (
	ProtocolHTTP  = Protocol("http")
	ProtocolHTTPS = Protocol("https")
	ProtocolTCP   = Protocol("tcp")
)

// An Algo is an enumeration of possible node balancer routing algorithms.
type Algo string

// Enum values for Algo.
const (
	AlgoRoundRobin = Algo("roundrobin")
	AlgoLeastConn  = Algo("leastconn")
	AlgoSource     = Algo("source")
)

// Sticky is an enumeration of possible node balancer stickiness settings.
type Sticky string

// Enum values for Sticky.
const (
	StickyNone       = Sticky("none")
	StickyTable      = Sticky("table")
	StickyHTTPCookie = Sticky("http_cookie")
)

// A Check is an enumeration of possible node balancer check settings.
type Check string

// Enum values for Check.
const (
	CheckNone       = Check("none")
	CheckConnection = Check("connection")
	CheckHTTP       = Check("http")
	CheckHTTPBody   = Check("http_body")
)

// A CipherSuite is an enumeration of possible ciphers to use for SSL connections served by a node balancer.
type CipherSuite string

// Enum values for CipherSuite.
const (
	CipherSuiteRecommended = CipherSuite("recommended")
	CipherSuiteLegacy      = CipherSuite("legacy")
)

// A CreateBalancerConfigRequest is a parameter struct for creating new node balancer configs.
type CreateBalancerConfigRequest struct {
	ID             uint        `json:"-"`
	NodeBalancerID uint        `json:"-"`
	CheckPath      string      `json:"check_path"`
	Protocol       Protocol    `json:"protocol"`
	Algorithm      Algo        `json:"algorithm"`
	Stickiness     Sticky      `json:"stickiness"`
	Check          Check       `json:"check"`
	CheckInterval  uint        `json:"check_interval"`
	CheckTimeout   uint        `json:"check_timeout"`
	CheckAttempts  uint        `json:"check_attempts"`
	Port           uint        `json:"port"`
	CheckBody      string      `json:"check_body"`
	CheckPassive   bool        `json:"check_passive"`
	CipherSuite    CipherSuite `json:"cipher_suite"`
	SSLCert        string      `json:"ssl_cert,omitempty"`
	SSLKey         string      `json:"ssl_key,omitempty"`
}

// An UpdateBalancerConfigRequest is a parameter struct for updating an existing node balancer config.
type UpdateBalancerConfigRequest CreateBalancerConfigRequest

// A BalancerConfig represents a configuration for a node balancer.
type BalancerConfig struct {
	ID             uint        `json:"id"`
	NodeBalancerID uint        `json:"-"`
	CheckPath      string      `json:"check_path"`
	Protocol       Protocol    `json:"protocol"`
	Algorithm      Algo        `json:"algorithm"`
	Stickiness     Sticky      `json:"stickiness"`
	Check          Check       `json:"check"`
	CheckInterval  uint        `json:"check_interval"`
	CheckTimeout   uint        `json:"check_timeout"`
	CheckAttempts  uint        `json:"check_attempts"`
	Port           uint        `json:"port"`
	CheckBody      string      `json:"check_body"`
	CheckPassive   bool        `json:"check_passive"`
	CipherSuite    CipherSuite `json:"cipher_suite"`
	SSLCert        string      `json:"ssl_cert,omitempty"`
	SSLKey         string      `json:"ssl_key,omitempty"`
	SSLFingerPrint string      `json:"ssl_fingerprint,omitempty"`
	SSLCommonName  string      `json:"ssl_commonname,omitempty"`
	// TODO: Not doing embedded struct definitions anywhere else...maybe break this out.
	NodeStatus struct {
		Up   uint `json:"up"`
		Down uint `json:"down"`
	} `json:"node_status"`
}

// A NodeStatus is an enumeration of possible node statuses.
type NodeStatus string

// Enum values for NodeStatus.
const (
	NodeStatusUnknown = NodeStatus("unknown")
	NodeStatusUp      = NodeStatus("UP")
	NodeStatusDown    = NodeStatus("DOWN")
)

// A NodeMode is an enumeration of possible node modes.
type NodeMode string

// Enum values for NodeMode.
const (
	NodeModeAccept = NodeMode("accept")
	NodeModeReject = NodeMode("reject")
	NodeModeDrain  = NodeMode("drain")
)

// A Node represents an individual node behind a node balancer.
type Node struct {
	ID         uint       `json:"id"`
	BalancerID uint       `json:"nodebalancer_id"`
	ConfigID   uint       `json:"config_id"`
	Address    string     `json:"address"`
	Label      string     `json:"label"`
	Status     NodeStatus `json:"status"`
	Weight     uint       `json:"weight"`
	Mode       NodeMode   `json:"mode"`
}

// A CreateNodeRequest is a parameter struct for creating a new node within a node balancer.
type CreateNodeRequest struct {
	BalancerID uint     `json:"-"`
	ConfigID   uint     `json:"-"`
	Label      string   `json:"label"`
	Weight     uint     `json:"weight,omitempty"`
	Mode       NodeMode `json:"mode,omitempty"`
}

// An UpdateNodeRequest is a parameter struct for updating an existing node within a node balancer.
type UpdateNodeRequest struct {
	ID         uint     `json:"-"`
	BalancerID uint     `json:"-"`
	ConfigID   uint     `json:"-"`
	Address    string   `json:"address,omitempty"`
	Label      string   `json:"label,omitempty"`
	Weight     uint     `json:"weight,omitempty"`
	Mode       NodeMode `json:"mode,omitempty"`
}

// A Balancer works with Linode node balancers.
type Balancer interface {
	ListNodeBalancers() ([]NodeBalancer, error)
	ViewNodeBalancer(id string) (NodeBalancer, error)
	CreateNodeBalancer(req CreateBalancerRequest) (NodeBalancer, error)
	UpdateNodeBalancer(req UpdateBalancerRequest) (NodeBalancer, error)
	DeleteNodeBalancer(id string) error

	ListNodeBalancerConfigs(balancerID uint) ([]BalancerConfig, error)
	ViewNodeBalancerConfig(balancerID, configID uint) (BalancerConfig, error)
	CreateNodeBalancerConfig(req CreateBalancerConfigRequest) (BalancerConfig, error)
	UpdateNodeBalancerConfig(req UpdateBalancerConfigRequest) (BalancerConfig, error)
	DeleteNodeBalancerConfig(balancerID, configID uint) error

	ListNodes(balancerID uint) ([]Node, error)
	ViewNode(balancerID, nodeID uint) (Node, error)
	CreateNode(req CreateNodeRequest) (Node, error)
	UpdateNode(req UpdateNodeRequest) (Node, error)
	DeleteNode(balancerID, nodeID uint) error
}

package main

import (
	"encoding/json"

	"github.com/lib/pq"

	"github.com/jinzhu/gorm"
)

//FactHandler :
type FactHandler int

//FactPort :
const FactPort = 8200

// AppGroup Type Constants
const (
	_ = iota
	Logical
	Development
	Training
	QA
	Staging
	Production
	Other
)

// Server Type Constants
const (
	Hypervisor = "hypervisor"
	BareMetal  = "baremetal"
	VM         = "virtualmachine"
	Container  = "container"
	OtherOS    = "other"
	App        = "application"
)

// Service Type Constants
const (
	WebServices = "webservices" // such as httpd
	Database    = "database"    // such as MySQL

)

// Service Port Type Constants
const (
	External = "external"
	Internal = "internal"
)

// Server Boot Script Fn
const (
	PreBootScript  = 1
	PostBootScript = 2
)

// IPAddress Types
const (
	Management  = "management.v4"
	Transaction = "transaction.v4"
	All         = "*.v4"
)

// Volume Types
const (
	RootVolume  = "root"
	LogVolume   = "log"
	DataVolume  = "data"
	OtherVolume = "other"
)

//LogicalAppGroup :
type LogicalAppGroup AppGroupAbstract

//ConcreteAppGroup :
type ConcreteAppGroup AppGroupAbstract

//AppGroupBase :
type AppGroupBase struct {
	AppGroupName    string `json:"name,omitempty"`
	AppGroupID      string `json:"GUID,omitempty"` // UUID.String() Value
	DefaultAppGroup bool   `json:"default"`
	UserOrg         string `json:"userOrg,omitempty"`
	UserGroup       string `json:"userGroup,omitempty"`
	UserDomain      string `json:"userDomain,omitempty"`
}

// AppGroupAbstract : Holds the logical structure of the application group (say, SAP Landscape)
type AppGroupAbstract struct {
	ID string `json:"id,omitempty"`
	AppGroupBase
	AppGroupType string            `json:"type,omitempty"` // Refer to AppGroupType Constants
	AppGroupTags map[string]string `json:"tags,omitempty"`
	Applications []struct {
		AppSequence int
		Application *AppAbstract
	} `json:"apps,omitempty"`
}

type AppGroupModel struct {
	gorm.Model
	AppGroupBase
	AppGroupType int               // Refer to AppGroupType Constants
	AppGroupTags map[string]string `gorm:"-"`
	AppTags      json.RawMessage   `sql:"type:json"`
	// Applications []uint8         `sql:"type:integer[]"`
	Applications pq.Int64Array `sql:"type:integer[]"`
}

type AppBase struct {
	ApplicationName string
	ApplicationID   string // UUID.String() Value
	ApplicationURL  string // url.URL format: [scheme:][//[userinfo@]host][/]path[?query][#fragment]
}

// AppAbstract : Holds the logical structure for the application component
type AppAbstract struct {
	ID string
	AppBase
	AppTags    map[string]string
	AppObjects map[string]interface{} // This creates a simple map of all the Application Objects at Runtime. Notation: AppObject[server.<servername>] = <server object>
	Servers    []struct {
		ServerBootSeq int
		ServerBootFn  func(map[string]interface{}, int) error
		Server        *ServerAbstract
	}
}

type AppModel struct {
	gorm.Model
	AppBase
}

// ServerAbstract : Holds the structure for server/s in the application component
type ServerAbstract struct {
	ServerID       string // UUID.String() Value
	ServerName     string
	ServerType     string // Refer to the Server Type Constants
	ServerTags     map[string]string
	ServerURL      string // url.URL format: [scheme:][//[userinfo@]host][/]path[?query][#fragment]
	ServerHostName string
	OSInfo         struct {
		OSDistribution string
		OSKernelLevel  string
		OSPatchLevel   string
	}
	IPv4Addresses map[string]string // Key = "management.v4" or "transaction.v4" or "*.v4" or "iscsi.v4". Output = IPv4Address
	Volumes       map[string]string // Key = "root", "log", "data" etc.
	OtherInfo     struct {
		CPU      int
		GBMemory int
	}
	ClusterInfo struct {
		ClusterName    string
		ClusterMembers []*ServerAbstract
	}
	Services []ServiceAbstract
}

// ServiceAbstract : Holds the structure for services running on the servers in the application component
type ServiceAbstract struct {
	ServiceID    string // UUID.String() Value
	ServiceName  string
	ServiceType  string // Refer to the Server Type Constants
	ServiceTags  map[string]string
	ServiceURL   string // url.URL format: [scheme:][//[userinfo@]host][/]path[?query][#fragment]
	ServicePorts []struct {
		PortNumber string
		PortType   string // External or Internal
	}
	ServicePathInHost string
	ServiceCommands   struct {
		CmdToStart   string
		CmdToStop    string
		CmdToRestart string
	}
}

// Credentials :
type Credential struct {
	gorm.Model    `json:"-"`
	CredGUID      string                 `json:"credguid,omitempty"` // UUID.String() Value
	CredType      string                 `json:"credtype,omitempty"`
	UserDomain    string                 `json:"userDomain,omitempty"`
	UserOrg       string                 `json:"userOrg,omitempty"`
	UserGroup     string                 `json:"userGroup,omitempty"`
	UserID        string                 `json:"userID,omitempty"`
	PasswordHash  []byte                 `json:"passwordhash,omitempty"`
	Password      string                 `json:"passwordstring,omitempty"`
	LoginURL      string                 `json:"loginurl,omitempty"` // url.URL format: [scheme:][//[userinfo@]host][/]path[?query][#fragment]
	SSHPrivateKey []byte                 `json:"sshprivatekey,omitempty"`
	OtherCredInfo map[string]interface{} `json:"othercredinfomap,omitempty"` //holds all other information as key/value pair
}

type AppCred struct {
	AppGUID string        `json:"appguid,omitempty"`
	Creds   []*Credential `json:"creds,omitempty"`
}

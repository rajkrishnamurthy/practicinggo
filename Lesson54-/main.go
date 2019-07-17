package main

import (
	"encoding/json"
	"time"
)

// Configuration : Mapping the File to the Configuration Spec
type Configuration struct {
	ConfigurationName    string
	ConfigurationVersion string
	ConfigurationScope
	ConfigurationCategory              string
	ConfigurationSourceObjectType      string      // Refer ObjecType Constants
	ConfigurationSourceObject          interface{} // This should contain a pointer to the source object such as task, rule
	ConfigurationTargetObjectType      string      // Refer ObjecType Constants
	ConfigurationTargetObject          interface{} // This should contain a pointer to the target object such as report
	ConfigurationTargetMatchExpression string      // Regex expression that will match the name
	ConfigurationSpecs                 []ConfigurationSpec
	ConfigurationCreateDTM             time.Time
	ConfigurationLastUpdatedDTM        time.Time
}

// ConfigurationScope : Defines the scope for the Configuration Specification
type ConfigurationScope struct {
	ConfigurationScopeType string // Refer SCOPE_TYPE constants
	Domain                 string // * if SCOPE_TYPE=global. If not, provide the specific Domain value
	Org                    string // * if SCOPE_TYPE=global. If not, provide the specific Org value
	Group                  string // * if SCOPE_TYPE=global. If not, provide the specific Group value
	User                   string // * if SCOPE_TYPE=global. If need to personalize for a given user, provide user id.
}

// ConfigurationSpec : Parent Configuration Object
type ConfigurationSpec struct {
	ConfigurationGUID        string                                  // Unique Identifier
	ConfigurationDataType    string                                  // Refer DataType Constants
	ConfigurationDataMap     map[string]*ConfigurationDataMapDetails // The key to the DataMap is the "FieldName" which will have to be unique
	ConfigurationDataFile    json.RawMessage                         // This is alternative implementation of the map. User can directly store encoded json map
	ConfigurationDataMapHash string                                  // hash value of the ConfigurationDataMap
}

// ConfigurationDataMapDetails : This refers to the field mapping. Only applicable for ConfigurationCategory = CATEGORY_DATA
type ConfigurationDataMapDetails struct {
	// FieldName        string
	FieldDisplayName string
	IsFieldIndexed   bool
	IsFieldVisible   bool
	FieldMapPath     string // This contains the map path such as JSON PATH or CSV file header
	FieldOrder       int    // Order in which the field should appear
	FieldExpression  string // This is valid only for Computed Types.
}

// ConfigurationScopeType
const (
	SCOPE_TYPE_GLOBAL = 10
	SCOPE_TYPE_DOMAIN = 20
	SCOPE_TYPE_USER   = 30
)

// ConfigurationCategory
const (
	CATEGORY_DATA   = "data"
	CATEGORY_ACTION = "action"
)

// ConfigurationDataType
const (
	TYPE_DATA            = "data"           // This refer to the data configuration
	TYPE_METADATA        = "metadata"       // This refers to the metadata that is preconfigured and generated by the system
	TYPE_ACTION_RESPONSE = "actionresponse" // This refers to the data coming from the action created through the cnextensionservice such as creating a service now ticket and getting the acknowledgement response
	TYPE_ACTION_INBOUND  = "actioninbound"  // This refers to the data coming asynchronously from external sources thorugh cnextensionservice
	TYPE_COMPUTED        = "computed"       // This refers to the data that is computed based on other fields

)

// ObjectTypes
const (
	OBJECT_TASK    = "task"
	OBJECT_RULE    = "rule"
	OBJECT_CONTROL = "control"
	OBJECT_REPORT  = "report"
)

func main() {

}

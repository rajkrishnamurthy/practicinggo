package main

import (
	"context"
	"encoding/json"
	"time"
)

//WorkflowConfig : Identifies a workflow configuration group
type WorkflowConfig struct {
	WorkflowConfigGUID        string              // UUID string that uniquely identifies the workflow configuration
	WorkflowConfigName        string              // Give it a name
	WorkflowConfigDescription string              // Give it a description
	WorkflowConfigLabels      map[string][]string // Key:Value/s pair for the Workflow configuration
	WorkflowConfigType        int                 // Refer to the constants for the Workflow configuration Type
	// WorkflowConfigStates      []*WorkflowStateConfig
	WorkflowStateTransitions []WorkflowStateTransition
}

// //WorkflowStateConfig : This defines the states and the state transitions for the workflow
// type WorkflowStateConfig struct {
// 	// *WorkflowConfig
// 	WorkflowStateConfigGUID  string // UUID string that uniquely identifies the workflow state transition
// 	WorkflowStateTransitions []WorkflowStateTransition
// }

// WorkflowStateTransition : Define the transition
type WorkflowStateTransition struct {
	FromState       string
	ToState         string
	TriggeringEvent string
	Policy          *Policy
	StateEnter      *WorkflowExecutionUnitInterface // if ok, err := .Execute; ok {<State entry success. Move forward>} else {<Abort>}
	StateExit       *WorkflowExecutionUnitInterface // // if ok, err := .Execute; ok {<State exit success. Move forward>} else {<Abort>}
	ExecutionUnit   *WorkflowExecutionUnitInterface // // if ok, err := .Execute; ok {<Execution success. Move forward>} else {<Abort>}
} // This will get loaded from a json object at run time

// WorkflowExecutionUnitInterface : All objects that participate in the workflow execution such as flow, command, rule, REST calls etc. should support this interface methods
type WorkflowExecutionUnitInterface interface {
	Execute(context.Context, *WorkflowStateInstance) (bool, error)
	CancelExecution(context.Context, *WorkflowStateInstance) (bool, error)
}

// GenericWorkflowExecutionUnit : All things that will execute for the worflow (at design time) will go here
type GenericWorkflowExecutionUnit struct {
	Name                     string
	Labels                   map[string][]string
	ImplementationObjectType string      // Refer to the constants for the Workflow Execution Unit
	ImplementationObject     interface{} // This can attach any allowable Implementation Object
}

//Policy : Normalizes the Policy Information
type Policy struct {
	PolicyGUID     string // UUID string that uniquely identifies the flow
	PolicyName     string
	PolicyLabels   map[string][]string
	MatchingPolicy *MatchingPolicy // Refer to constants on the Matching Elements Types: Domain, Org, Group, etc.
	SLAPolicy      *SLAPolicy
}

//MatchElement : Matching Criteria
type MatchElement struct {
	Scope              string              // Refer to Scope Constants
	Elements           []string            // Specifics such as an array of Domain Names, or Org Names etc. Scope cannot be "." or "*"
	Context            string              // Refer to Context Constants
	Labels             map[string][]string // These labels will get carried forward into all matching elements
	SelectorExpression interface{}         // To be determined

}

//MatchingPolicy : Defines the Matching Policy for User identification and Assignment
type MatchingPolicy map[int]*MatchElement

//SLAPolicy : Defines the SLAs and the remediation actions if breached
type SLAPolicy struct {
	ExpectedDurationInHours int // time.Duration.Hours()
	GraceDurationInHours    int // time.Duration.Hours()
	// Warning = Status > ExpectedDuration
	// Breach = Status > Expected + Grace
	RemediationAction WorkflowExecutionUnitInterface // if ok, err := .Execute; ok {<Remediation success. Move forward>} else {<Abort>}
}

// WorkflowStateInstance : The real implementation details of the Workflow Configuration
type WorkflowStateInstance struct {
	WorkflowStateInstanceGUID     string
	WorkflowStateInstanceSeq      int // Each workflow state instance can have multiple sub-states. Refer above
	ctx                           context.Context
	CurrentState                  string    // state name. Refer configuration
	TriggeredEvent                string    // Which event triggered the movement into the current state
	StateStartDTM                 time.Time // Time that the state started (i.e, transitioned into)
	StateEndDTM                   time.Time // Time that the state started (i.e, transitioned into)
	Status                        string    // Refer to the Constants for State's States. These transitions are hardcoded
	MatchingUsers                 []*UserEntity
	AssignedUsers                 []*UserEntity
	InputPayloadIntoExecutionUnit []byte // This can be considered proof of work
	OutputFromExecutionUnit       []byte
	PreviousState                 *WorkflowStateInstance
}

// UserEntity : Describes the attribute for a User Object in the Rule Request Workflow system
type UserEntity struct {
	DomainID string
	OrgID    string
	GroupID  string
	UserID   string
	Role     string
	Labels   map[string][]string //Carried forward from Matching Elements
}

// WorkflowStateInstanceHistory : Audit records
type WorkflowStateInstanceHistory struct {
	WorkflowConfigGUID                string
	WorkflowStateConfigGUID           string
	WorkflowStateInstanceGUID         string
	WorkflowStateInstanceSeq          int       // Each workflow state instance can have multiple sub-states. Refer above
	RecordedDTM                       time.Time // Time that this record was inserted
	RecordedBy                        *UserSessionInfo
	CurrentState                      string          // state name. Refer configuration
	TriggeredEvent                    string          // Which event triggered the movement into the current state
	StateStartDTM                     time.Time       // Time that the state started (i.e, transitioned into)
	StateEndDTM                       time.Time       // Time that the state started (i.e, transitioned into)
	Status                            string          // Refer to the Constants for State's States. These transitions are hardcoded
	MatchingUsers                     json.RawMessage // JSON representation of []UserEntity
	AssignedUsers                     json.RawMessage // JSON representation of []UserEntity
	InputPayloadIntoExecutionUnit     []byte          // This can be considered proof of work
	OutputFromExecutionUnit           []byte
	PreviousWorkflowStateInstanceGUID string
}

// UserSessionInfo : Contains information on the operator making state changes
type UserSessionInfo struct {
	AuthToken string
	UserID    string
}

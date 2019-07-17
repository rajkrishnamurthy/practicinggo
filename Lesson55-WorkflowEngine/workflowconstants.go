package main

// Constants for Flow Type
const (
	RuleDevFlow = iota + 1
)

// Constants for System Defined Flow Phases
const (
	Initial = "intial"
	Final   = "final"
)

// Constants for State's State :(
const (
	ToDo = iota + 1 // Not yet started
	Waiting
	Doing
	Cancelled
	Done
)

// Scope for Domains, Orgs, Groups, Users and Roles
const (
	ScopeNA         = ""
	ScopeCurrent    = "."
	ScopeAll        = "*"
	ContinubeDomain = "continube"
)

// Contexts for Domains, Orgs, Groups, Users and Roles
const (
	ContextPostIncludeAll  = "*"
	ContextPostIncludeOnly = "+"
	ContextPostExcludeOnly = "-"
)

// Generic Constants for User/System Events for State Transition
// The Specific event names will be Phase + Generic Event
const (
	Start = "start"
	// CheckOut = "checkout"
	// CheckIn  = "checkin"
	Next   = "next"
	Prev   = "prev"
	Yield  = "yield"
	Remove = "remove"
	Reset  = "reset"
)

// Constants for Matching Elements
const (
	MatchDomain = iota + 1
	MatchOrg
	MatchGroup
	MatchUser
	MatchRole
)

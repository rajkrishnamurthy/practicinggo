package main

import "time"

// Configuration Methods:
// ----------------
func initWorkflowConfig() *WorkflowConfig {
	return &WorkflowConfig{}
}

func (*WorkflowConfig) initWorkflowStateConfig() *WorkflowStateConfig {
	return nil
}

func (workflowstateconfig *WorkflowStateConfig) constructWorkflowTransition() (*WorkflowStateTransition, error) {
	return nil, nil
}

func (*WorkflowStateTransition) setExecutionUnit(WorkflowExecutionUnitInterface) (bool, error) {
	return true, nil
}

func (*WorkflowStateTransition) setSLAPolicy(SLAPolicy) (bool, error) {
	return true, nil
}

func (*WorkflowStateTransition) setMatchingPolicy(MatchingPolicy) (bool, error) {
	return true, nil
}

func (*WorkflowStateTransition) addMatchingPolicy(*MatchElement) (bool, error) {
	return true, nil
}

func (workflowconfig *WorkflowConfig) setWorkflowStates(workflowtransitions []WorkflowStateTransition) (*WorkflowStateConfig, error) {
	return nil, nil
}

func (workflowconfig *WorkflowConfig) addWorkflowState(workflowtransitions WorkflowStateTransition) (*WorkflowStateConfig, error) {
	return nil, nil
}

// Execution Methods
// ----------------
func (*WorkflowConfig) instantiateWorkflowConfig() *WorkflowStateInstance {
	return &WorkflowStateInstance{}
}

func (*WorkflowStateInstance) getMatchedUsers(State string) (users []UserEntity, err error) {
	return nil, nil
}

func (*WorkflowStateInstance) getAssignedUsers(State string) (users []UserEntity, err error) {
	return nil, nil
}

func (*WorkflowStateInstance) assignUsers(State string, users []UserEntity) (err error) {
	return nil
}

func (*WorkflowStateInstance) transitionState(currentState string, event string, payload []byte, users []UserEntity, sessionInfo *UserSessionInfo) (outputPayload []byte, ok bool, err error) {
	return nil, false, nil
}

func (*WorkflowStateInstance) moveNext(currentState string, payload []byte, users []UserEntity, sessionInfo *UserSessionInfo) (outputPayload []byte, ok bool, err error) {
	// event := "next"
	// outputPayload, ok, err := workflowstateinstance.transitionState(currentState, event, payload, users, sessionInfo)

	return nil, false, nil
}

func (*WorkflowStateInstance) getExecutionStatus(State string) (status string, err error) {
	return "pass", nil
}

// Reporting Methods
// ----------------

func (*WorkflowStateInstance) getSLAInfo(State string) (slapolicy *SLAPolicy, startDTM, endDTM time.Time) {
	return nil, time.Now(), time.Now()
}

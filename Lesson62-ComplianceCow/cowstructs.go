package main

import (
	"encoding/json"
	"time"
)

type cowPlan struct {
	ID              string // Unique GUID for the Plan
	Domain          string // Domain for which the compliance cow plan is set. This will be pre configured
	Org             string
	Group           string
	CNPlanGUID      string              // This refers to the PLAN GUID in the ContiNube system
	PlanControls    []cowControl        // This refers to the Controls configured under the Plan in the ContiNube System by default. Can we overriden in Compliane Cow. Refer to the struct below.
	Version         string              // Version number maintained in Compliance Cow agains the PlanGUID
	cowPlanWorkflow *cowControlWorkflow // This will point to the general workflow configuration for the Plan
}

type cowPlanInstance struct {
	*cowPlan
	ID                       string              // Unique GUID for the PlanInstance
	PlanInstance             int                 // This is a running sequence number, every time PCI-DSS asssessment is "started"
	CNPlanExecutionGUID      string              // This refers to the Plan Execution GUID in the ContiNube system
	ControlsScope            []cowControl        // The user can select one or more controls for assesssment
	PlanInstanceStarted      time.Time           // This is the time when the user started the assesssment
	PlanInstanceEnded        time.Time           // This is the time when the user created some 'finality' for the assessment
	CNPlanExecutionStartTime time.Time           // Start of the Plan Execution from the ContiNube subsystem
	CNPlanExecutionEndTime   time.Time           // Start of the Plan Execution from the ContiNube subsystem
	PlanInstanceStatus       string              // To be determined. Possible states: start, in-progress, aborted, completed
	InstanceInitiatedBy      cowUser             // Who started this instance?
	cowPlanInstanceWorkflow  *cowControlWorkflow // This will point to the specific workflow configuration for the Plan Instance
}

type cowUser struct {
	ID                string // Unique GUID for the cow user
	Domain            string
	Org               string
	Group             string
	Roles             []*cowRole
	LastTransactionID string // GUID association in the Transaction Audit Table
	LastStatus        string // Active, Disabled, Suspended, Removed
	LastUpdated       time.Time
	ReportingTo       []*cowUser // Who does this person report to?
	UserTags          map[string][]string
}

type cowRole struct {
	ID string // Unique GUID for the cowRole
	// Same as ContiNube Role Object
	// Need to bring in Role, Privilege objects into Compliance Cow
	RoleTags map[string][]string
}

type cowControl struct {
	CNPlanControlID    string // This refers to the ContiNube Control GUID
	CNParentControlID  string // This refers to the ContiNube Control's Parent GUID (in ContiNube)
	ControlName        string // This fetches the Control Name from ContiNube as default but can be overriden in Compliance Cow
	ControlDescription string // This fetches the Control Description from ContiNube as default but can be overriden in Compliance Cow
	ControlDisplayable string // This fetches the Control Displayable field from ContiNube as default but can be overriden in Compliance Cow
	ControlAlias       string // This fetches the Control Alias from ContiNube as default but can be overriden in Compliance Cow
	ControlPriority    string // High, Medium, Low
	ControlStage       string // Internal, External, QSA
	ControlStatus      string // Status for Internal or External Stages: Pending, In-Progress, Submitted, Approved, Rejected; Status for QSA: In-Place, In-Place-With-CCW, Not-In-Place, Not Tested, Not Applicable
	ControlRisk        cowControlRisk
	ControlType        string              // Preventative, Detective and Compensating Control
	ControlTags        map[string][]string // This fetches the Control Tags from ContiNube as default but can be overriden in Compliance Cow
	CNDependsOn        []string            // This fetches the Control DependsOn field from ContiNube as default but can be overriden in Compliance Cow
}

type cowControlNotes struct {
	*cowPlan                            // This points to the Plan.i.e, TEMPLATE. Note: You can only have either the Plan or the Plan Instance
	*cowPlanInstance                    // This points to the notes for a specific instance. If NIL refers to TEMPLATE NOTES
	Control                 *cowControl // The Notes applies to a specific control
	Notes                   *cowGeneralNotes
	LastStatus              string // Status of Notes; In-Progress, Active, Closed
	LastUpdated             time.Time
	AssignedTo              []cowUser           // Assigned to specific users. Default will be ALL USERs
	ControlNotesTags        map[string][]string // This fetches the Control Tags from ContiNube as default but can be overriden in Compliance Cow
	cowControlNotesWorkflow *cowControlWorkflow
}

type cowControlChecklist struct {
	*cowPlan                                // This points to the Plan.i.e, TEMPLATE. Note: You can only have either the Plan or the Plan Instance
	*cowPlanInstance                        // This points to the notes for a specific instance. If NIL refers to TEMPLATE CheckList
	Control                     *cowControl // The Notes applies to a specific control
	Topic                       string      // Topic for the Notes
	Description                 string
	Creator                     cowUser
	CreatedOn                   time.Time
	Priority                    int // byte operated. High, Medium, Low
	DueDate                     time.Time
	LastStatus                  string // Status of Notes; In-Progress, Active, Closed
	LastUpdated                 time.Time
	Owner                       cowUser             // Who is the owner of the checklist item. By default, the creator is the owner
	AssignedTo                  []cowUser           // Assigned to specific users. Default will be ALL USERs
	ControlCheckListTags        map[string][]string // This fetches the Control Tags from ContiNube as default but can be overriden in Compliance Cow
	cowControlChecklistWorkflow *cowControlWorkflow
}

type cowControlAttachments struct {
	*cowPlanInstance                         // This points to the notes for a specific Control instance. Cannot be NIL i.e, there is no cncept of template for Attachments
	Control                      *cowControl // The Notes applies to a specific control
	AttachmentName               string      // Topic for the Notes
	AttachmentFileLocation       string      // cowStorage Location where the file is storage. cowStorage + cowMinio are clones of cnStorage and cnMinio services in Compliance Cow
	Creator                      cowUser
	CreatedOn                    time.Time
	Owner                        cowUser             // Who is the owner of the checklist item. By default, the creator is the owner
	AssignedTo                   []cowUser           // Assigned to specific users. Default will be ALL USERs
	ControlAttachmentTags        map[string][]string // This fetches the Control Tags from ContiNube as default but can be overriden in Compliance Cow
	cowControlAttachmentWorkflow *cowControlWorkflow
}

type cowControlEvidences struct {
	*cowPlanInstance             // This points to the notes for a specific Control instance. Cannot be NIL i.e, there is no cncept of template for Attachments
	Control          *cowControl // The Notes applies to a specific control
	Evidences        []cowEvidence
}

type cowControlExceptions struct {
	*cowPlanInstance               // This points to the notes for a specific Control instance. Cannot be NIL i.e, there is no cncept of template for Attachments
	Control          *cowControl   // The Notes applies to a specific control
	Exceptions       []cowEvidence // Typically there will be only one section
}

type cowEvidence struct {
	ID                      string      // GUID assigned to each Evidence
	SelectedVersion         *cowVersion // This points to the selected version of Control Evidence records.
	CNSynthesizerFilterGUID string      // This points to the Synthesizer card in ContiNube for Evidence
	CNSQLFilter             string      // SQL that can be applied on top of Synthesizer output
	EvidenceTopic           string
	EvidenceDescription     string
	EvidenceType            string // Automated, Manual
	DisplayOrder            int    // The Display Order for the evidences
	Creator                 cowUser
	CreatedOn               time.Time
	Owner                   cowUser   // Who is the owner of the checklist item. By default, the creator is the owner
	AssignedTo              []cowUser // Assigned to specific users. Default will be ALL USERs
	Priority                int       // byte operated. High, Medium, Low
	DueDate                 time.Time
	EvidenceNotes           *cowControlNotes // When notes are made at the overall Evidence Section level
	LastStatus              string           // Status of Notes; In-Progress, Active, Closed
	LastUpdated             time.Time
	EvidenceVersion         *cowVersion // We will maintain the last 5 versions of the evidences SUBMITTED
	EvidenceWorkflow        *cowControlWorkflow
	EvidenceTemplateVersion *cowVersion
	EvidenceTags            map[string][]string // User defined tags for each evidence section and/or exceptions
}

type cowRecordset struct {
	EvidenceGUID       string            // Which evidence does this specific recordset belong to
	Records            []json.RawMessage // Every working record is stored as an []bytes. Also, every record from CNSynthesizer should contain a record GUID
	SelectedRecordsSQL string            // Translating selected records to SQL, if applicable
	LastTransactionID  string            // GUID association in the Transaction Audit Table
	AssignedTo         []cowUser         // Assigned to specific users. Default will be ALL USERs
	Priority           int               // byte operated. High, Medium, Low
	DueDate            time.Time
	EvidenceNotes      *cowGeneralNotes // When notes are made at the overall Evidence Section level
	LastStatus         string           // Status of Notes; In-Progress, Active, Closed
	LastUpdated        time.Time
	RecordsetWorkflow  *cowControlWorkflow
	RecordsetTags      map[string][]string // User defined tags for each evidence section and/or exceptions
}

type cowControlConfig struct {
}

type cowControlSurveyConfig struct {
}

type cowControlWorkflow struct {
}

type cowTransactionAudit struct {
	ID                           string //Unique GUID for the transaction
	TransactionType              int    // Bit Operated for the Type of Transaciton: plan, planinstance, user, role, privilege, workflow, control, control-notes, control-attachments, control-action-checklist, control-actions, control-exception, control-evidence, control-evidence-recordset,control-evidence-action, control-survey
	LogLevel                     int    // To seperate detailed logs. Levels 1 - 5
	TargetObjectID               string // GUID of the target object being logged
	TransactionCorrelationID     string
	TransactionStatus            int // successful, failure
	TransactionDoneBy            cowUser
	TransactionDTM               time.Time
	TransactionRequestData       []byte // Raw data of the transaction
	TransactionRequesDataType    string
	TransactionResponseData      []byte // Raw data of the transaction
	TransactionResponseDataType  string
	TransactionExceptionData     []byte // Raw data of the transaction
	TransactionExceptionDataType string
	UserNotification             bool   // Should we show the transaction record to the user
	GitCommitPath                string // Is this redundant?
	GitCommitGUID                string // Commit ID we get from Git Commit
}

type cowVersion struct {
	ID                  string
	Repo                string
	Branch              string
	Type                string // Indicates the type of the file: template, evidence, exception
	WorkingFileLocation string // Working copy maintained in cowStorage/cowMinio. Any save goes into working file
	TargetFileInRepo    string // Any commit goes into GIT
	CommitHash          string
	PreviousVersion     *cowVersion // Should we avoid these attributes and use git directly?
	NextVersion         *cowVersion // Should we avoid these attributes and use git directly?

	// How will we manage conflicts?
}

type cowGeneralNotes struct {
	Topic       string // Topic for the Notes
	Notes       string
	NoteType    string // Status, General
	Creator     cowUser
	CreatedOn   time.Time
	Priority    int // byte operated. High, Medium, Low
	DueDate     time.Time
	LastUpdated time.Time
}

type cowControlRisk struct {
}

type cowIssue struct {
}

type cowObservation struct {
}

package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_Employee - A SoftLayer_User_Employee models a single SoftLayer employee for the
// purposes of ticket updates created by SoftLayer employees. SoftLayer portal and API users cannot see
// individual employee names in ticket responses. SoftLayer employees can be assigned to customer
// accounts as a personal support representative. Employee names and email will be available if an
// employee is assigned to the account.
type SoftLayer_User_Employee struct {

	// Email - A SoftLayer employee's email address. Email addresses are only visible to
	// [[SoftLayer_Account|SoftLayer Accounts]] that are assigned to an employee
	Email string `json:"email"`

	// LastName - A SoftLayer employee's last name. Last names are only visible to
	// [[SoftLayer_Account|SoftLayer Accounts]] that are assigned to an employee
	LastName string `json:"lastName"`

	// EmployeeDepartmentId - A SoftLayer employee's [[SoftLayer_User_Employee_Department|department]] id.
	EmployeeDepartmentId int `json:"employeeDepartmentId"`

	// FirstName - A SoftLayer employee's first name. First names are only visible to
	// [[SoftLayer_Account|SoftLayer Accounts]] that are assigned to an employee
	FirstName string `json:"firstName"`

	// OfficePhone - <nil>
	OfficePhone string `json:"officePhone"`

	// DisplayName - <nil>
	DisplayName string `json:"displayName"`

	// Username - A representation of a SoftLayer employee's username. In all cases this should simply
	// state "Employee".
	Username string `json:"username"`
}

func (softlayer_user_employee *SoftLayer_User_Employee) String() string {
	return "SoftLayer_User_Employee"
}

// SoftLayer_User_Employee_Extended is SoftLayer_User_Employee with all maskable types.
type SoftLayer_User_Employee_Extended struct {
	SoftLayer_User_Employee

	// MetricTrackingObject - <nil>
	MetricTrackingObject *SoftLayer_Metric_Tracking_Object `json:"metricTrackingObject"`

	// EmployeeDepartment - The department that a SoftLayer employee belongs to.
	EmployeeDepartment *SoftLayer_User_Employee_Department `json:"employeeDepartment"`

	// LayoutProfiles - <nil>
	LayoutProfiles []*SoftLayer_Layout_Profile `json:"layoutProfiles"`

	// Roles - <nil>
	Roles []*SoftLayer_User_Permission_Role `json:"roles"`

	// TicketActivities - <nil>
	TicketActivities []*SoftLayer_Ticket_Activity `json:"ticketActivities"`

	// TicketAttachmentReferences - <nil>
	TicketAttachmentReferences []*SoftLayer_Ticket_Attachment `json:"ticketAttachmentReferences"`

	// ActionCount - no documentation
	ActionCount uint64 `json:"actionCount"`

	// Actions - <nil>
	Actions []*SoftLayer_User_Permission_Action `json:"actions"`

	// ChatTranscript - <nil>
	ChatTranscript []*SoftLayer_Ticket_Chat `json:"chatTranscript"`

	// TicketAttachmentReferenceCount - no documentation
	TicketAttachmentReferenceCount uint64 `json:"ticketAttachmentReferenceCount"`

	// ChatTranscriptCount - no documentation
	ChatTranscriptCount uint64 `json:"chatTranscriptCount"`

	// LayoutProfileCount - no documentation
	LayoutProfileCount uint64 `json:"layoutProfileCount"`

	// RoleCount - no documentation
	RoleCount uint64 `json:"roleCount"`

	// TicketActivityCount - no documentation
	TicketActivityCount uint64 `json:"ticketActivityCount"`
}

func (softlayer_user_employee *SoftLayer_User_Employee_Extended) String() string {
	return "SoftLayer_User_Employee"
}

package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_Employee - A SoftLayer_User_Employee models a single SoftLayer employee for the
// purposes of ticket updates created by SoftLayer employees. SoftLayer portal and API users cannot see
// individual employee names in ticket responses. SoftLayer employees can be assigned to customer
// accounts as a personal support representative. Employee names and email will be available if an
// employee is assigned to the account.
type SoftLayer_User_Employee struct {

	// EmployeeDepartmentId - A SoftLayer employee's [[SoftLayer_User_Employee_Department|department]] id.
	EmployeeDepartmentId int `json:"employeeDepartmentId,omitempty"`

	// FirstName - A SoftLayer employee's first name. First names are only visible to
	// [[SoftLayer_Account|SoftLayer Accounts]] that are assigned to an employee
	FirstName string `json:"firstName,omitempty"`

	// Email - A SoftLayer employee's email address. Email addresses are only visible to
	// [[SoftLayer_Account|SoftLayer Accounts]] that are assigned to an employee
	Email string `json:"email,omitempty"`

	// DisplayName - <nil>
	DisplayName string `json:"displayName,omitempty"`

	// OfficePhone - <nil>
	OfficePhone string `json:"officePhone,omitempty"`

	// Username - A representation of a SoftLayer employee's username. In all cases this should simply
	// state "Employee".
	Username string `json:"username,omitempty"`

	// LastName - A SoftLayer employee's last name. Last names are only visible to
	// [[SoftLayer_Account|SoftLayer Accounts]] that are assigned to an employee
	LastName string `json:"lastName,omitempty"`

	// MetricTrackingObject - <nil>
	MetricTrackingObject *SoftLayer_Metric_Tracking_Object `json:"metricTrackingObject,omitempty"`

	// Roles - <nil>
	Roles []*SoftLayer_User_Permission_Role `json:"roles,omitempty"`

	// TicketActivities - <nil>
	TicketActivities []*SoftLayer_Ticket_Activity `json:"ticketActivities,omitempty"`

	// LayoutProfileCount - no documentation
	LayoutProfileCount uint64 `json:"layoutProfileCount,omitempty"`

	// RoleCount - no documentation
	RoleCount uint64 `json:"roleCount,omitempty"`

	// TicketAttachmentReferenceCount - no documentation
	TicketAttachmentReferenceCount uint64 `json:"ticketAttachmentReferenceCount,omitempty"`

	// LayoutProfiles - <nil>
	LayoutProfiles []*SoftLayer_Layout_Profile `json:"layoutProfiles,omitempty"`

	// TicketAttachmentReferences - <nil>
	TicketAttachmentReferences []*SoftLayer_Ticket_Attachment `json:"ticketAttachmentReferences,omitempty"`

	// ChatTranscriptCount - no documentation
	ChatTranscriptCount uint64 `json:"chatTranscriptCount,omitempty"`

	// TicketActivityCount - no documentation
	TicketActivityCount uint64 `json:"ticketActivityCount,omitempty"`

	// Actions - <nil>
	Actions []*SoftLayer_User_Permission_Action `json:"actions,omitempty"`

	// ChatTranscript - <nil>
	ChatTranscript []*SoftLayer_Ticket_Chat `json:"chatTranscript,omitempty"`

	// EmployeeDepartment - The department that a SoftLayer employee belongs to.
	EmployeeDepartment *SoftLayer_User_Employee_Department `json:"employeeDepartment,omitempty"`

	// ActionCount - no documentation
	ActionCount uint64 `json:"actionCount,omitempty"`
}

func (softlayer_user_employee *SoftLayer_User_Employee) String() string {
	return "SoftLayer_User_Employee"
}

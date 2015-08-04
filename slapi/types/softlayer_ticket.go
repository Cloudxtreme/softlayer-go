package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Ticket - The SoftLayer_Ticket data type models a single SoftLayer customer support or
// notification ticket. Each ticket object contains references to it's updates, the user it's assigned
// to, the SoftLayer department and employee that it's assigned to, and any hardware objects or
// attached files associated with the ticket. Tickets are described in further detail on the
// [[SoftLayer_Ticket]] service page. To create a support ticket execute the
// [[SoftLayer_Ticket::createStandardTicket|createStandardTicket]] or
// [[SoftLayer_Ticket::createAdministrativeTicket|createAdministrativeTicket]] methods in the
// SoftLayer_Ticket service. To create an upgrade ticket for the SoftLayer sales group execute the
// [[SoftLayer_Ticket::createUpgradeTicket|createUpgradeTicket]].
type SoftLayer_Ticket struct {

	// StatusId - no documentation
	StatusId int `json:"statusId,omitempty"`

	// ServerAdministrationFlag - Whether a ticket is a standard or an administrative support ticket.
	// Administrative support tickets typically incur a $3 USD charge.
	ServerAdministrationFlag int `json:"serverAdministrationFlag,omitempty"`

	// ServerAdministrationBillingAmount - The amount of money in US Dollars that a ticket has charged to
	// an account. A ticket's administrative billing amount is a one time charge and only applies to
	// administrative support tickets.
	ServerAdministrationBillingAmount int `json:"serverAdministrationBillingAmount,omitempty"`

	// ServerAdministrationBillingInvoiceId - The internal identifier of the invoice associated with a
	// ticket's administrative charge. Only tickets with an administrative charge have an associated
	// invoice.
	ServerAdministrationBillingInvoiceId int `json:"serverAdministrationBillingInvoiceId,omitempty"`

	// SubjectId - An internal identifier of the pre-set subject that a ticket is associated with. Standard
	// support tickets have a subject set while administrative tickets have a null subject. A standard
	// support ticket's title is the name of it's associated subject.
	SubjectId int `json:"subjectId,omitempty"`

	// Title - A ticket's title. This is typically a brief summary of the issue described in the ticket.
	Title string `json:"title,omitempty"`

	// NotifyUserOnUpdateFlag - Whether or not the user who owns a ticket is notified via email when a
	// ticket is updated.
	NotifyUserOnUpdateFlag bool `json:"notifyUserOnUpdateFlag,omitempty"`

	// TotalUpdateCount - <nil>
	TotalUpdateCount int `json:"totalUpdateCount,omitempty"`

	// ChangeOwnerFlag - <nil>
	ChangeOwnerFlag bool `json:"changeOwnerFlag,omitempty"`

	// Id - A ticket's internal identifier. Each ticket is defined by a unique identifier.
	Id int `json:"id,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// AssignedUserId - An internal identifier of the portal user that a ticket is assigned to.
	AssignedUserId int `json:"assignedUserId,omitempty"`

	// UserEditableFlag - no documentation
	UserEditableFlag bool `json:"userEditableFlag,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// AccountId - An internal identifier of the SoftLayer customer account that a ticket is associated
	// with.
	AccountId int `json:"accountId,omitempty"`

	// LocationId - The internal identifier of the location associated with a ticket.
	LocationId int `json:"locationId,omitempty"`

	// ResponsibleBrandId - <nil>
	ResponsibleBrandId int `json:"responsibleBrandId,omitempty"`

	// ServerAdministrationRefundInvoiceId - The internal identifier of the refund invoice associated with
	// a ticket. Only tickets with an account refund associated with them have an associated refund
	// invoice.
	ServerAdministrationRefundInvoiceId int `json:"serverAdministrationRefundInvoiceId,omitempty"`

	// FinalComments - Feedback left by a portal or API user on their experiences in a ticket. Final
	// comments may be created after a ticket is closed.
	FinalComments string `json:"finalComments,omitempty"`

	// GroupId - The internal identifier of the SoftLayer department that a ticket is assigned to.
	GroupId int `json:"groupId,omitempty"`

	// Priority - <nil>
	Priority int `json:"priority,omitempty"`

	// OriginatingIpAddress - no documentation
	OriginatingIpAddress string `json:"originatingIpAddress,omitempty"`

	// ServiceProviderId - <nil>
	ServiceProviderId int `json:"serviceProviderId,omitempty"`

	// BillableFlag - Whether a ticket has a one-time charge associated with it. Standard tickets are free
	// while administrative tickets typically cost $3
	BillableFlag bool `json:"billableFlag,omitempty"`

	// LastEditDate - The date that a ticket was last modified. A modification does not necessarily mean
	// that an update was added.
	LastEditDate *time.Time `json:"lastEditDate,omitempty"`

	// LastEditType - The type of user who last edited or updated a ticket. This is either or
	LastEditType string `json:"lastEditType,omitempty"`

	// ServiceProviderResourceId - A ticket's internal identifier at its service provider. Each ticket is
	// defined by a unique identifier.
	ServiceProviderResourceId int `json:"serviceProviderResourceId,omitempty"`
}

func (softlayer_ticket *SoftLayer_Ticket) String() string {
	return "SoftLayer_Ticket"
}

// SoftLayer_Ticket_Extended is SoftLayer_Ticket with all maskable types.
type SoftLayer_Ticket_Extended struct {
	SoftLayer_Ticket

	// InvoiceItemCount - A count of the invoice items associated with a ticket. Ticket based invoice items
	// only exist when a ticket incurs a fee that has been invoiced.
	InvoiceItemCount uint64 `json:"invoiceItemCount,omitempty"`

	// FirstAttachedResource - The first physical or virtual server attached to a ticket.
	FirstAttachedResource *SoftLayer_Ticket_Attachment `json:"firstAttachedResource,omitempty"`

	// ServiceProvider - <nil>
	ServiceProvider *SoftLayer_Service_Provider `json:"serviceProvider,omitempty"`

	// LastUpdate - no documentation
	LastUpdate *SoftLayer_Ticket_Update `json:"lastUpdate,omitempty"`

	// AwaitingUserResponseFlag - no documentation
	AwaitingUserResponseFlag bool `json:"awaitingUserResponseFlag,omitempty"`

	// LastActivity - <nil>
	LastActivity *SoftLayer_Ticket_Activity `json:"lastActivity,omitempty"`

	// Location - A ticket's associated location within the SoftLayer location hierarchy.
	Location *SoftLayer_Location `json:"location,omitempty"`

	// AssignedAgentCount - no documentation
	AssignedAgentCount uint64 `json:"assignedAgentCount,omitempty"`

	// AttachedAdditionalEmails - The list of additional emails to notify when a ticket update is made.
	AttachedAdditionalEmails []*SoftLayer_User_Customer_AdditionalEmail `json:"attachedAdditionalEmails,omitempty"`

	// LastViewedDate - A timestamp of the last time the Ticket was viewed by the active user.
	LastViewedDate *time.Time `json:"lastViewedDate,omitempty"`

	// ServerAdministrationBillingInvoice - The invoice associated with a ticket. Only tickets with an
	// associated administrative charge have an invoice.
	ServerAdministrationBillingInvoice *SoftLayer_Billing_Invoice `json:"serverAdministrationBillingInvoice,omitempty"`

	// UpdateCount - no documentation
	UpdateCount uint64 `json:"updateCount,omitempty"`

	// TagReferenceCount - no documentation
	TagReferenceCount uint64 `json:"tagReferenceCount,omitempty"`

	// AttachedResources - <nil>
	AttachedResources []*SoftLayer_Ticket_Attachment `json:"attachedResources,omitempty"`

	// Updates - no documentation
	Updates []*SoftLayer_Ticket_Update `json:"updates,omitempty"`

	// EmployeeAttachmentCount - no documentation
	EmployeeAttachmentCount uint64 `json:"employeeAttachmentCount,omitempty"`

	// ScheduledActionCount - no documentation
	ScheduledActionCount uint64 `json:"scheduledActionCount,omitempty"`

	// AssignedUser - no documentation
	AssignedUser *SoftLayer_User_Customer `json:"assignedUser,omitempty"`

	// Group - The SoftLayer department that a ticket is assigned to.
	Group *SoftLayer_Ticket_Group `json:"group,omitempty"`

	// StateCount - no documentation
	StateCount uint64 `json:"stateCount,omitempty"`

	// AttachedHardware - The hardware associated with a ticket. This is used in cases where a ticket is
	// directly associated with one or more pieces of hardware.
	AttachedHardware []*SoftLayer_Hardware `json:"attachedHardware,omitempty"`

	// EmployeeAttachments - <nil>
	EmployeeAttachments []*SoftLayer_User_Employee `json:"employeeAttachments,omitempty"`

	// LastEditor - <nil>
	LastEditor *SoftLayer_User_Interface `json:"lastEditor,omitempty"`

	// Account - The SoftLayer customer account associated with a ticket.
	Account *SoftLayer_Account `json:"account,omitempty"`

	// NewUpdatesFlag - True if there are new, unread updates to this ticket for the current user, False
	// otherwise.
	NewUpdatesFlag bool `json:"newUpdatesFlag,omitempty"`

	// AttachedFiles - no documentation
	AttachedFiles []*SoftLayer_Ticket_Attachment_File `json:"attachedFiles,omitempty"`

	// AttachedHardwareCount - <nil>
	AttachedHardwareCount uint `json:"attachedHardwareCount,omitempty"`

	// Subject - A ticket's subject. Only standard support tickets have an associated subject. A standard
	// support ticket's title corresponds with it's subject's name.
	Subject *SoftLayer_Ticket_Subject `json:"subject,omitempty"`

	// TagReferences - <nil>
	TagReferences []*SoftLayer_Tag_Reference `json:"tagReferences,omitempty"`

	// AttachedVirtualGuestCount - A count of the virtual guests associated with a ticket. This is used in
	// cases where a ticket is directly associated with one or more virtualized guests installations or
	// Virtual Servers.
	AttachedVirtualGuestCount uint64 `json:"attachedVirtualGuestCount,omitempty"`

	// AttachedVirtualGuests - The virtual guests associated with a ticket. This is used in cases where a
	// ticket is directly associated with one or more virtualized guests installations or Virtual Servers.
	AttachedVirtualGuests []*SoftLayer_Virtual_Guest `json:"attachedVirtualGuests,omitempty"`

	// ScheduledActions - <nil>
	ScheduledActions []*SoftLayer_Provisioning_Version1_Transaction `json:"scheduledActions,omitempty"`

	// Status - no documentation
	Status *SoftLayer_Ticket_Status `json:"status,omitempty"`

	// InvoiceItems - The invoice items associated with a ticket. Ticket based invoice items only exist
	// when a ticket incurs a fee that has been invoiced.
	InvoiceItems []*SoftLayer_Billing_Invoice_Item `json:"invoiceItems,omitempty"`

	// State - <nil>
	State []*SoftLayer_Ticket_State `json:"state,omitempty"`

	// AttachedAdditionalEmailCount - A count of the list of additional emails to notify when a ticket
	// update is made.
	AttachedAdditionalEmailCount uint64 `json:"attachedAdditionalEmailCount,omitempty"`

	// AttachedFileCount - no documentation
	AttachedFileCount uint64 `json:"attachedFileCount,omitempty"`

	// AssignedAgents - <nil>
	AssignedAgents []*SoftLayer_User_Customer `json:"assignedAgents,omitempty"`

	// CancellationRequest - no documentation
	CancellationRequest *SoftLayer_Billing_Item_Cancellation_Request `json:"cancellationRequest,omitempty"`

	// FirstUpdate - The first update made to a ticket. This is typically the contents of a ticket when
	// it's created.
	FirstUpdate *SoftLayer_Ticket_Update `json:"firstUpdate,omitempty"`

	// ServerAdministrationRefundInvoice - The refund invoice associated with a ticket. Only tickets with a
	// refund applied in them have an associated refund invoice.
	ServerAdministrationRefundInvoice *SoftLayer_Billing_Invoice `json:"serverAdministrationRefundInvoice,omitempty"`

	// AttachedResourceCount - no documentation
	AttachedResourceCount uint64 `json:"attachedResourceCount,omitempty"`
}

func (softlayer_ticket *SoftLayer_Ticket_Extended) String() string {
	return "SoftLayer_Ticket"
}

package sl

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

	// Account - The SoftLayer customer account associated with a ticket.
	Account *SoftLayer_Account `json:"account"`

	// AccountId - An internal identifier of the SoftLayer customer account that a ticket is associated
	// with.
	AccountId int `json:"accountId"`

	// AssignedAgentCount - no documentation
	AssignedAgentCount uint64 `json:"assignedAgentCount"`

	// AssignedAgents - <nil>
	AssignedAgents []*SoftLayer_User_Customer `json:"assignedAgents"`

	// AssignedUser - no documentation
	AssignedUser *SoftLayer_User_Customer `json:"assignedUser"`

	// AssignedUserId - An internal identifier of the portal user that a ticket is assigned to.
	AssignedUserId int `json:"assignedUserId"`

	// AttachedAdditionalEmailCount - A count of the list of additional emails to notify when a ticket
	// update is made.
	AttachedAdditionalEmailCount uint64 `json:"attachedAdditionalEmailCount"`

	// AttachedAdditionalEmails - The list of additional emails to notify when a ticket update is made.
	AttachedAdditionalEmails []*SoftLayer_User_Customer_AdditionalEmail `json:"attachedAdditionalEmails"`

	// AttachedFileCount - no documentation
	AttachedFileCount uint64 `json:"attachedFileCount"`

	// AttachedFiles - no documentation
	AttachedFiles []*SoftLayer_Ticket_Attachment_File `json:"attachedFiles"`

	// AttachedHardware - The hardware associated with a ticket. This is used in cases where a ticket is
	// directly associated with one or more pieces of hardware.
	AttachedHardware []*SoftLayer_Hardware `json:"attachedHardware"`

	// AttachedHardwareCount - <nil>
	AttachedHardwareCount uint `json:"attachedHardwareCount"`

	// AttachedResourceCount - no documentation
	AttachedResourceCount uint64 `json:"attachedResourceCount"`

	// AttachedResources - <nil>
	AttachedResources []*SoftLayer_Ticket_Attachment `json:"attachedResources"`

	// AttachedVirtualGuestCount - A count of the virtual guests associated with a ticket. This is used in
	// cases where a ticket is directly associated with one or more virtualized guests installations or
	// Virtual Servers.
	AttachedVirtualGuestCount uint64 `json:"attachedVirtualGuestCount"`

	// AttachedVirtualGuests - The virtual guests associated with a ticket. This is used in cases where a
	// ticket is directly associated with one or more virtualized guests installations or Virtual Servers.
	AttachedVirtualGuests []*SoftLayer_Virtual_Guest `json:"attachedVirtualGuests"`

	// AwaitingUserResponseFlag - no documentation
	AwaitingUserResponseFlag bool `json:"awaitingUserResponseFlag"`

	// BillableFlag - Whether a ticket has a one-time charge associated with it. Standard tickets are free
	// while administrative tickets typically cost $3
	BillableFlag bool `json:"billableFlag"`

	// CancellationRequest - no documentation
	CancellationRequest *SoftLayer_Billing_Item_Cancellation_Request `json:"cancellationRequest"`

	// ChangeOwnerFlag - <nil>
	ChangeOwnerFlag bool `json:"changeOwnerFlag"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// EmployeeAttachmentCount - no documentation
	EmployeeAttachmentCount uint64 `json:"employeeAttachmentCount"`

	// EmployeeAttachments - <nil>
	EmployeeAttachments []*SoftLayer_User_Employee `json:"employeeAttachments"`

	// FinalComments - Feedback left by a portal or API user on their experiences in a ticket. Final
	// comments may be created after a ticket is closed.
	FinalComments string `json:"finalComments"`

	// FirstAttachedResource - The first physical or virtual server attached to a ticket.
	FirstAttachedResource *SoftLayer_Ticket_Attachment `json:"firstAttachedResource"`

	// FirstUpdate - The first update made to a ticket. This is typically the contents of a ticket when
	// it's created.
	FirstUpdate *SoftLayer_Ticket_Update `json:"firstUpdate"`

	// Group - The SoftLayer department that a ticket is assigned to.
	Group *SoftLayer_Ticket_Group `json:"group"`

	// GroupId - The internal identifier of the SoftLayer department that a ticket is assigned to.
	GroupId int `json:"groupId"`

	// Id - A ticket's internal identifier. Each ticket is defined by a unique identifier.
	Id int `json:"id"`

	// InvoiceItemCount - A count of the invoice items associated with a ticket. Ticket based invoice items
	// only exist when a ticket incurs a fee that has been invoiced.
	InvoiceItemCount uint64 `json:"invoiceItemCount"`

	// InvoiceItems - The invoice items associated with a ticket. Ticket based invoice items only exist
	// when a ticket incurs a fee that has been invoiced.
	InvoiceItems []*SoftLayer_Billing_Invoice_Item `json:"invoiceItems"`

	// LastActivity - <nil>
	LastActivity *SoftLayer_Ticket_Activity `json:"lastActivity"`

	// LastEditDate - The date that a ticket was last modified. A modification does not necessarily mean
	// that an update was added.
	LastEditDate *time.Time `json:"lastEditDate"`

	// LastEditType - The type of user who last edited or updated a ticket. This is either or
	LastEditType string `json:"lastEditType"`

	// LastEditor - <nil>
	LastEditor *SoftLayer_User_Interface `json:"lastEditor"`

	// LastUpdate - no documentation
	LastUpdate *SoftLayer_Ticket_Update `json:"lastUpdate"`

	// LastViewedDate - A timestamp of the last time the Ticket was viewed by the active user.
	LastViewedDate *time.Time `json:"lastViewedDate"`

	// Location - A ticket's associated location within the SoftLayer location hierarchy.
	Location *SoftLayer_Location `json:"location"`

	// LocationId - The internal identifier of the location associated with a ticket.
	LocationId int `json:"locationId"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// NewUpdatesFlag - True if there are new, unread updates to this ticket for the current user, False
	// otherwise.
	NewUpdatesFlag bool `json:"newUpdatesFlag"`

	// NotifyUserOnUpdateFlag - Whether or not the user who owns a ticket is notified via email when a
	// ticket is updated.
	NotifyUserOnUpdateFlag bool `json:"notifyUserOnUpdateFlag"`

	// OriginatingIpAddress - no documentation
	OriginatingIpAddress string `json:"originatingIpAddress"`

	// Priority - <nil>
	Priority int `json:"priority"`

	// ResponsibleBrandId - <nil>
	ResponsibleBrandId int `json:"responsibleBrandId"`

	// ScheduledActionCount - no documentation
	ScheduledActionCount uint64 `json:"scheduledActionCount"`

	// ScheduledActions - <nil>
	ScheduledActions []*SoftLayer_Provisioning_Version1_Transaction `json:"scheduledActions"`

	// ServerAdministrationBillingAmount - The amount of money in US Dollars that a ticket has charged to
	// an account. A ticket's administrative billing amount is a one time charge and only applies to
	// administrative support tickets.
	ServerAdministrationBillingAmount int `json:"serverAdministrationBillingAmount"`

	// ServerAdministrationBillingInvoice - The invoice associated with a ticket. Only tickets with an
	// associated administrative charge have an invoice.
	ServerAdministrationBillingInvoice *SoftLayer_Billing_Invoice `json:"serverAdministrationBillingInvoice"`

	// ServerAdministrationBillingInvoiceId - The internal identifier of the invoice associated with a
	// ticket's administrative charge. Only tickets with an administrative charge have an associated
	// invoice.
	ServerAdministrationBillingInvoiceId int `json:"serverAdministrationBillingInvoiceId"`

	// ServerAdministrationFlag - Whether a ticket is a standard or an administrative support ticket.
	// Administrative support tickets typically incur a $3 USD charge.
	ServerAdministrationFlag int `json:"serverAdministrationFlag"`

	// ServerAdministrationRefundInvoice - The refund invoice associated with a ticket. Only tickets with a
	// refund applied in them have an associated refund invoice.
	ServerAdministrationRefundInvoice *SoftLayer_Billing_Invoice `json:"serverAdministrationRefundInvoice"`

	// ServerAdministrationRefundInvoiceId - The internal identifier of the refund invoice associated with
	// a ticket. Only tickets with an account refund associated with them have an associated refund
	// invoice.
	ServerAdministrationRefundInvoiceId int `json:"serverAdministrationRefundInvoiceId"`

	// ServiceProvider - <nil>
	ServiceProvider *SoftLayer_Service_Provider `json:"serviceProvider"`

	// ServiceProviderId - <nil>
	ServiceProviderId int `json:"serviceProviderId"`

	// ServiceProviderResourceId - A ticket's internal identifier at its service provider. Each ticket is
	// defined by a unique identifier.
	ServiceProviderResourceId int `json:"serviceProviderResourceId"`

	// State - <nil>
	State []*SoftLayer_Ticket_State `json:"state"`

	// StateCount - no documentation
	StateCount uint64 `json:"stateCount"`

	// Status - no documentation
	Status *SoftLayer_Ticket_Status `json:"status"`

	// StatusId - no documentation
	StatusId int `json:"statusId"`

	// Subject - A ticket's subject. Only standard support tickets have an associated subject. A standard
	// support ticket's title corresponds with it's subject's name.
	Subject *SoftLayer_Ticket_Subject `json:"subject"`

	// SubjectId - An internal identifier of the pre-set subject that a ticket is associated with. Standard
	// support tickets have a subject set while administrative tickets have a null subject. A standard
	// support ticket's title is the name of it's associated subject.
	SubjectId int `json:"subjectId"`

	// TagReferenceCount - no documentation
	TagReferenceCount uint64 `json:"tagReferenceCount"`

	// TagReferences - <nil>
	TagReferences []*SoftLayer_Tag_Reference `json:"tagReferences"`

	// Title - A ticket's title. This is typically a brief summary of the issue described in the ticket.
	Title string `json:"title"`

	// TotalUpdateCount - <nil>
	TotalUpdateCount int `json:"totalUpdateCount"`

	// UpdateCount - no documentation
	UpdateCount uint64 `json:"updateCount"`

	// Updates - no documentation
	Updates []*SoftLayer_Ticket_Update `json:"updates"`

	// UserEditableFlag - no documentation
	UserEditableFlag bool `json:"userEditableFlag"`
}

// AddAssignedAgent - <nil>
func (softlayer_ticket *SoftLayer_Ticket) AddAssignedAgent(agentId int) error {
	return nil
}

// AddAttachedAdditionalEmails - Creates new additional emails for assigned user if new emails are
// provided. Attaches any newly created additional emails to ticket.
func (softlayer_ticket *SoftLayer_Ticket) AddAttachedAdditionalEmails(emails []string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AddAttachedFile - Attach the given file to a SoftLayer ticket. A file attachment is a convenient way
// to submit non-textual error reports to SoftLayer employees in a ticket. File attachments to tickets
// must have a unique name.
func (softlayer_ticket *SoftLayer_Ticket) AddAttachedFile(fileAttachment SoftLayer_Container_Utility_File_Attachment) (*SoftLayer_Ticket_Attachment_File, error) {
	var returnValue *SoftLayer_Ticket_Attachment_File
	return returnValue, nil
}

// AddAttachedHardware - Attach the given hardware to a SoftLayer ticket. A hardware attachment
// provides an easy way for SoftLayer's employees to quickly look up your hardware records in the case
// of hardware-specific issues.
func (softlayer_ticket *SoftLayer_Ticket) AddAttachedHardware(hardwareId int) (*SoftLayer_Ticket_Attachment_Hardware, error) {
	var returnValue *SoftLayer_Ticket_Attachment_Hardware
	return returnValue, nil
}

// AddAttachedVirtualGuest - Attach the given CloudLayer Computing Instance to a SoftLayer ticket. An
// attachment provides an easy way for SoftLayer's employees to quickly look up your records in the
// case of specific issues.
func (softlayer_ticket *SoftLayer_Ticket) AddAttachedVirtualGuest(guestId int) (*SoftLayer_Ticket_Attachment_Virtual_Guest, error) {
	var returnValue *SoftLayer_Ticket_Attachment_Virtual_Guest
	return returnValue, nil
}

// AddFinalComments - As part of the customer service process SoftLayer has provided a quick feedback
// mechanism for its customers to rate their overall experience with SoftLayer after a ticket is
// closed. addFinalComments() sets these comments for a ticket update made by a SoftLayer employee.
// Final comments may only be set on closed tickets, can only be set once, and may not exceed 4000
// characters in length. Once the comments are set ''addFinalComments()'' returns a boolean true.
func (softlayer_ticket *SoftLayer_Ticket) AddFinalComments(finalComments string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AddScheduledAlert - <nil>
func (softlayer_ticket *SoftLayer_Ticket) AddScheduledAlert(activationTime string) error {
	return nil
}

// AddScheduledAutoClose - <nil>
func (softlayer_ticket *SoftLayer_Ticket) AddScheduledAutoClose(activationTime string) error {
	return nil
}

// AddUpdate - Add an update to a ticket. A ticket update's entry has a maximum length of 4000
// characters, so ''addUpdate()'' splits the ''entry'' property in the ''templateObject'' parameter
// into 3900 character blocks and creates one entry per 3900 character block. Once complete
// ''addUpdate()'' emails the ticket's owner and additional email addresses with an update message if
// the ticket's ''notifyUserOnUpdateFlag'' is set. If the ticket is a Legal or Abuse ticket, then the
// account's abuse email is also notified when the updates are processed. Finally, ''addUpdate()''
// returns an array of the newly created ticket updates.
func (softlayer_ticket *SoftLayer_Ticket) AddUpdate(templateObject SoftLayer_Ticket_Update, attachedFiles []SoftLayer_Container_Utility_File_Attachment) ([]*SoftLayer_Ticket_Update, error) {
	var returnValue []*SoftLayer_Ticket_Update
	return returnValue, nil
}

// CreateAdministrativeTicket - Create an administrative support ticket. Use an administrative ticket
// if you require SoftLayer's assistance managing your server or content. If you are experiencing an
// issue with SoftLayer's hardware, network, or services then please open a standard support ticket.
// Support tickets may only be created in the open state. The SoftLayer API defaults new ticket
// properties ''userEditableFlag'' to true, ''accountId'' to the id of the account that your API user
// belongs to, and ''statusId'' to 1001 (or "open"). You may not assign your new to ticket to users
// that your API user does not have access to. Once your ticket is created it is placed in a queue for
// SoftLayer employees to work. As they update the ticket new [[SoftLayer_Ticket_Update]] entries are
// added to the ticket object. Administrative support tickets add a one-time $3USD charge to your
// account.
func (softlayer_ticket *SoftLayer_Ticket) CreateAdministrativeTicket(templateObject SoftLayer_Ticket, contents string, attachmentId int, rootPassword string, controlPanelPassword string, accessPort string, attachedFiles []SoftLayer_Container_Utility_File_Attachment, attachmentType string) (*SoftLayer_Ticket, error) {
	var returnValue *SoftLayer_Ticket
	return returnValue, nil
}

// CreateCancelServerTicket - A cancel server request creates a ticket to cancel the resource on next
// bill date. The hardware ID parameter is required to determine which server is to be cancelled.
// Hourly bare metal servers will be cancelled on next bill date. The reason parameter could be from
// the list below: * "No longer needed" * "Business closing down" * "Server / Upgrade Costs" *
// "Migrating to larger server" * "Migrating to smaller server" * "Migrating to a different SoftLayer
// datacenter" * "Network performance / latency" * "Support response / timing" * "Sales process /
// upgrades" * "Moving to competitor" The content parameter describes further the reason for cancelling
// the server.
func (softlayer_ticket *SoftLayer_Ticket) CreateCancelServerTicket(attachmentId int, reason string, content string, cancelAssociatedItems bool, attachmentType string) (*SoftLayer_Ticket, error) {
	var returnValue *SoftLayer_Ticket
	return returnValue, nil
}

// CreateCancelServiceTicket - A cancel service request creates a sales ticket. The hardware ID
// parameter is required to determine which server is to be cancelled. The reason parameter could be
// from the list below: * "No longer needed" * "Business closing down" * "Server / Upgrade Costs" *
// "Migrating to larger server" * "Migrating to smaller server" * "Migrating to a different SoftLayer
// datacenter" * "Network performance / latency" * "Support response / timing" * "Sales process /
// upgrades" * "Moving to competitor" The content parameter describes further the reason for cancelling
// service.
func (softlayer_ticket *SoftLayer_Ticket) CreateCancelServiceTicket(attachmentId int, reason string, content string, attachmentType string) (*SoftLayer_Ticket, error) {
	var returnValue *SoftLayer_Ticket
	return returnValue, nil
}

// CreateStandardTicket - Create a standard support ticket. Use a standard support ticket if you need
// to work out a problem related to SoftLayer's hardware, network, or services. If you require
// SoftLayer's assistance managing your server or content then please open an administrative ticket.
// Support tickets may only be created in the open state. The SoftLayer API defaults new ticket
// properties ''userEditableFlag'' to true, ''accountId'' to the id of the account that your API user
// belongs to, and ''statusId'' to 1001 (or "open"). You may not assign your new to ticket to users
// that your API user does not have access to. Once your ticket is created it is placed in a queue for
// SoftLayer employees to work. As they update the ticket new [[SoftLayer_Ticket_Update]] entries are
// added to the ticket object.
func (softlayer_ticket *SoftLayer_Ticket) CreateStandardTicket(templateObject SoftLayer_Ticket, contents string, attachmentId int, rootPassword string, controlPanelPassword string, accessPort string, attachedFiles []SoftLayer_Container_Utility_File_Attachment, attachmentType string) (*SoftLayer_Ticket, error) {
	var returnValue *SoftLayer_Ticket
	return returnValue, nil
}

// CreateUpgradeTicket - Create a ticket for the SoftLayer sales team to perform a hardware or service
// upgrade. Our sales team will work with you on upgrade feasibility and pricing and then send the
// upgrade ticket to the proper department to perform the actual upgrade. Service affecting upgrades,
// such as server hardware or CloudLayer Computing Instance upgrades that require the server powered
// down must have a two hour maintenance specified for our datacenter engineers to perform your
// upgrade. Account level upgrades, such as adding VPN users, CDNLayer accounts, and monitoring
// services are processed much faster and do not require a maintenance window.
func (softlayer_ticket *SoftLayer_Ticket) CreateUpgradeTicket(attachmentId int, genericUpgrade string, upgradeMaintenanceWindow string, details string, attachmentType string) (*SoftLayer_Ticket, error) {
	var returnValue *SoftLayer_Ticket
	return returnValue, nil
}

// Edit - Edit a SoftLayer ticket. The edit method is two-fold. You may either edit a ticket itself,
// add an update to a ticket, attach up to two files to a ticket, or perform all of these tasks. The
// SoftLayer API ignores changes made to the ''userEditableFlag'' and ''accountId'' properties. You may
// not assign a ticket to a user that your API account does not have access to. You may not enter a
// custom title for standard support tickets, buy may do so when editing an administrative ticket.
// Finally, you may not close a ticket using this method. Please contact SoftLayer if you need a ticket
// closed. If you need to only add an update to a ticket then please use the
// [[SoftLayer_Ticket::addUpdate|addUpdate]] method in this service. Likewise if you need to only
// attach a file to a ticket then use the [[SoftLayer_Ticket::addAttachedFile|addAttachedFile]] method.
// The edit method exists as a convenience if you need to perform all these tasks at once.
func (softlayer_ticket *SoftLayer_Ticket) Edit(templateObject SoftLayer_Ticket, contents string, attachedFiles []SoftLayer_Container_Utility_File_Attachment) (*SoftLayer_Ticket, error) {
	var returnValue *SoftLayer_Ticket
	return returnValue, nil
}

// GetAllTicketGroups - getAllTicketGroups() retrieves a list of all groups that a ticket may be
// assigned to. Ticket groups represent the internal department at SoftLayer who a ticket is assigned
// to. Every SoftLayer ticket has groupId and ticketGroup properties that correspond to one of the
// groups returned by getAllTicketGroups().
func (softlayer_ticket *SoftLayer_Ticket) GetAllTicketGroups() ([]*SoftLayer_Ticket_Group, error) {
	var returnValue []*SoftLayer_Ticket_Group
	return returnValue, nil
}

// GetAllTicketStatuses - getAllTicketStatuses() retrieves a list of all statuses that a ticket may
// exist in. Ticket status represent the current state of a ticket, usually "open", "assigned", and
// "closed". Every SoftLayer ticket has statusId and status properties that correspond to one of the
// statuses returned by getAllTicketStatuses().
func (softlayer_ticket *SoftLayer_Ticket) GetAllTicketStatuses() ([]*SoftLayer_Ticket_Status, error) {
	var returnValue []*SoftLayer_Ticket_Status
	return returnValue, nil
}

// GetAttachedFile - Retrieve the file attached to a SoftLayer ticket by it's given identifier. To
// retrieve a list of files attached to a ticket either call the SoftLayer_Ticket::getAttachedFiles
// method or call SoftLayer_Ticket::getObject with ''attachedFiles'' defined in an object mask.
func (softlayer_ticket *SoftLayer_Ticket) GetAttachedFile(attachmentId int) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Ticket object whose ID number corresponds to the ID
// number of the init parameter passed to the SoftLayer_Ticket service. You can only retrieve tickets
// that are associated with your SoftLayer customer account.
func (softlayer_ticket *SoftLayer_Ticket) GetObject() (*SoftLayer_Ticket, error) {
	var returnValue *SoftLayer_Ticket
	return returnValue, nil
}

// GetTicketsClosedSinceDate - no documentation
func (softlayer_ticket *SoftLayer_Ticket) GetTicketsClosedSinceDate(closeDate time.Time) ([]*SoftLayer_Ticket, error) {
	var returnValue []*SoftLayer_Ticket
	return returnValue, nil
}

// MarkAsViewed - Mark a ticket as viewed. All currently posted updates will be marked as viewed. The
// lastViewedDate property will be updated to the current time.
func (softlayer_ticket *SoftLayer_Ticket) MarkAsViewed() error {
	return nil
}

// RemoveAssignedAgent - <nil>
func (softlayer_ticket *SoftLayer_Ticket) RemoveAssignedAgent(agentId int) error {
	return nil
}

// RemoveAttachedAdditionalEmails - removeAttachedAdditionalEmails() removes the specified email
// addresses from a ticket's notification list. If one of the provided email addresses is not attached
// to the ticket then ''removeAttachedAdditiaonalEmails()'' ignores it and continues to the next one.
// Once the email addresses are removed ''removeAttachedAdditiaonalEmails()'' returns a boolean true.
func (softlayer_ticket *SoftLayer_Ticket) RemoveAttachedAdditionalEmails(emails []string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveAttachedHardware - detach the given hardware from a SoftLayer ticket. Removing a hardware
// attachment may delay ticket processing time if the hardware removed is relevant to the ticket's
// issue. Return a boolean true upon successful hardware detachment.
func (softlayer_ticket *SoftLayer_Ticket) RemoveAttachedHardware(hardwareId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveAttachedVirtualGuest - Detach the given CloudLayer Computing Instance from a SoftLayer ticket.
// Removing an attachment may delay ticket processing time if the instance removed is relevant to the
// ticket's issue. Return a boolean true upon successful detachment.
func (softlayer_ticket *SoftLayer_Ticket) RemoveAttachedVirtualGuest(guestId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveScheduledAlert - <nil>
func (softlayer_ticket *SoftLayer_Ticket) RemoveScheduledAlert() error {
	return nil
}

// RemoveScheduledAutoClose - <nil>
func (softlayer_ticket *SoftLayer_Ticket) RemoveScheduledAutoClose() error {
	return nil
}

// SetTags - <nil>
func (softlayer_ticket *SoftLayer_Ticket) SetTags(tags string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SurveyEligible - Use [[SoftLayer_Ticket_Survey::getPreference]] method.
func (softlayer_ticket *SoftLayer_Ticket) SurveyEligible() (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// UpdateAttachedAdditionalEmails - Creates new additional emails for assigned user if new emails are
// provided. Attaches any newly created additional emails to ticket. Remove any additional emails from
// a ticket that are not provided as part of $emails
func (softlayer_ticket *SoftLayer_Ticket) UpdateAttachedAdditionalEmails(emails []string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

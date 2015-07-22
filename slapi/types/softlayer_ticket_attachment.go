package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Ticket_Attachment - SoftLayer tickets have the ability to be associated with specific
// pieces of hardware in a customer's inventory. Attaching hardware to a ticket can greatly increase
// response time from SoftLayer for issues that are related to one or more specific servers on a
// customer's account. The SoftLayer_Ticket_Attachment_Hardware data type models the relationship
// between a piece of hardware and a ticket. Only one attachment record may exist per hardware item per
// ticket.
type SoftLayer_Ticket_Attachment struct {

	// AssignedAgent - <nil>
	AssignedAgent *SoftLayer_User_Customer `json:"assignedAgent"`

	// AttachmentId - The internal identifier of an item that is attached to a ticket.
	AttachmentId int `json:"attachmentId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`

	// ScheduledAction - <nil>
	ScheduledAction *SoftLayer_Provisioning_Version1_Transaction `json:"scheduledAction"`

	// Ticket - no documentation
	Ticket *SoftLayer_Ticket `json:"ticket"`

	// TicketId - The internal identifier of the ticket that an item is attached to.
	TicketId int `json:"ticketId"`
}
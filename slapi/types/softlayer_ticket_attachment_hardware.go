package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Ticket_Attachment_Hardware - SoftLayer tickets have the ability to be associated with
// specific pieces of hardware in a customer's inventory. Attaching hardware to a ticket can greatly
// increase response time from SoftLayer for issues that are related to one or more specific servers on
// a customer's account. The SoftLayer_Ticket_Attachment_Hardware data type models the relationship
// between a piece of hardware and a ticket. Only one attachment record may exist per hardware item per
// ticket.
type SoftLayer_Ticket_Attachment_Hardware struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// TicketId - The internal identifier of the ticket that an item is attached to.
	TicketId int `json:"ticketId,omitempty"`

	// HardwareId - The internal identifier of a piece of hardware that is attached to a ticket.
	HardwareId int `json:"hardwareId,omitempty"`

	// AttachmentId - The internal identifier of an item that is attached to a ticket.
	AttachmentId int `json:"attachmentId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`

	// Resource - no documentation
	Resource *SoftLayer_Hardware `json:"resource,omitempty"`

	// AssignedAgent - <nil>
	AssignedAgent *SoftLayer_User_Customer `json:"assignedAgent,omitempty"`

	// ScheduledAction - <nil>
	ScheduledAction *SoftLayer_Provisioning_Version1_Transaction `json:"scheduledAction,omitempty"`

	// Ticket - no documentation
	Ticket *SoftLayer_Ticket `json:"ticket,omitempty"`
}

func (softlayer_ticket_attachment_hardware *SoftLayer_Ticket_Attachment_Hardware) String() string {
	return "SoftLayer_Ticket_Attachment_Hardware"
}

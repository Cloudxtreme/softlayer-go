package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Ticket_Attachment_Virtual_Guest - SoftLayer tickets have the ability to be associated with
// specific pieces of hardware in a customer's inventory. Attaching hardware to a ticket can greatly
// increase response time from SoftLayer for issues that are related to one or more specific servers on
// a customer's account. The SoftLayer_Ticket_Attachment_Hardware data type models the relationship
// between a piece of hardware and a ticket. Only one attachment record may exist per hardware item per
// ticket.
type SoftLayer_Ticket_Attachment_Virtual_Guest struct {

	// VirtualGuestId - The internal identifier of the virtualized guest or CloudLayer Computing Instance
	// that is attached to a ticket.
	VirtualGuestId int `json:"virtualGuestId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// TicketId - The internal identifier of the ticket that an item is attached to.
	TicketId int `json:"ticketId,omitempty"`

	// AttachmentId - The internal identifier of an item that is attached to a ticket.
	AttachmentId int `json:"attachmentId,omitempty"`
}

func (softlayer_ticket_attachment_virtual_guest *SoftLayer_Ticket_Attachment_Virtual_Guest) String() string {
	return "SoftLayer_Ticket_Attachment_Virtual_Guest"
}

// SoftLayer_Ticket_Attachment_Virtual_Guest_Extended is SoftLayer_Ticket_Attachment_Virtual_Guest with all maskable types.
type SoftLayer_Ticket_Attachment_Virtual_Guest_Extended struct {
	SoftLayer_Ticket_Attachment_Virtual_Guest

	// AssignedAgent - <nil>
	AssignedAgent *SoftLayer_User_Customer `json:"assignedAgent,omitempty"`

	// Resource - The virtualized guest or CloudLayer Computing Instance that is attached to a ticket.
	Resource *SoftLayer_Virtual_Guest `json:"resource,omitempty"`

	// VirtualGuest - The virtualized guest or CloudLayer Computing Instance that is attached to a ticket.
	VirtualGuest *SoftLayer_Virtual_Guest `json:"virtualGuest,omitempty"`

	// Ticket - no documentation
	Ticket *SoftLayer_Ticket `json:"ticket,omitempty"`

	// ScheduledAction - <nil>
	ScheduledAction *SoftLayer_Provisioning_Version1_Transaction `json:"scheduledAction,omitempty"`
}

func (softlayer_ticket_attachment_virtual_guest *SoftLayer_Ticket_Attachment_Virtual_Guest_Extended) String() string {
	return "SoftLayer_Ticket_Attachment_Virtual_Guest"
}

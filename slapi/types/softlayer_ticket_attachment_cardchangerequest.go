package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Ticket_Attachment_CardChangeRequest - This datatype contains tickets referenced from card
// change request
type SoftLayer_Ticket_Attachment_CardChangeRequest struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// TicketId - The internal identifier of the ticket that an item is attached to.
	TicketId int `json:"ticketId,omitempty"`

	// AttachmentId - The internal identifier of an item that is attached to a ticket.
	AttachmentId int `json:"attachmentId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Resource - The card change request that is attached to a ticket.
	Resource *SoftLayer_Billing_Payment_Card_ChangeRequest `json:"resource,omitempty"`

	// AssignedAgent - <nil>
	AssignedAgent *SoftLayer_User_Customer `json:"assignedAgent,omitempty"`

	// ScheduledAction - <nil>
	ScheduledAction *SoftLayer_Provisioning_Version1_Transaction `json:"scheduledAction,omitempty"`

	// Ticket - no documentation
	Ticket *SoftLayer_Ticket `json:"ticket,omitempty"`
}

func (softlayer_ticket_attachment_cardchangerequest *SoftLayer_Ticket_Attachment_CardChangeRequest) String() string {
	return "SoftLayer_Ticket_Attachment_CardChangeRequest"
}

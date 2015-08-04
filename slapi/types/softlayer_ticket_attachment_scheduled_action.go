package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Ticket_Attachment_Scheduled_Action - <nil>
type SoftLayer_Ticket_Attachment_Scheduled_Action struct {

	// TicketId - The internal identifier of the ticket that an item is attached to.
	TicketId int `json:"ticketId,omitempty"`

	// AttachmentId - The internal identifier of an item that is attached to a ticket.
	AttachmentId int `json:"attachmentId,omitempty"`

	// RunDate - The internal identifier of a scheduled action transaction that is attached to a ticket.
	RunDate *time.Time `json:"runDate,omitempty"`

	// TransactionId - The internal identifier of a scheduled action transaction that is attached to a
	// ticket.
	TransactionId int `json:"transactionId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`
}

func (softlayer_ticket_attachment_scheduled_action *SoftLayer_Ticket_Attachment_Scheduled_Action) String() string {
	return "SoftLayer_Ticket_Attachment_Scheduled_Action"
}

// SoftLayer_Ticket_Attachment_Scheduled_Action_Extended is SoftLayer_Ticket_Attachment_Scheduled_Action with all maskable types.
type SoftLayer_Ticket_Attachment_Scheduled_Action_Extended struct {
	SoftLayer_Ticket_Attachment_Scheduled_Action

	// AssignedAgent - <nil>
	AssignedAgent *SoftLayer_User_Customer `json:"assignedAgent,omitempty"`

	// Ticket - no documentation
	Ticket *SoftLayer_Ticket `json:"ticket,omitempty"`

	// Resource - <nil>
	Resource *SoftLayer_Provisioning_Version1_Transaction `json:"resource,omitempty"`

	// Transaction - <nil>
	Transaction *SoftLayer_Provisioning_Version1_Transaction `json:"transaction,omitempty"`

	// ScheduledAction - <nil>
	ScheduledAction *SoftLayer_Provisioning_Version1_Transaction `json:"scheduledAction,omitempty"`
}

func (softlayer_ticket_attachment_scheduled_action *SoftLayer_Ticket_Attachment_Scheduled_Action_Extended) String() string {
	return "SoftLayer_Ticket_Attachment_Scheduled_Action"
}

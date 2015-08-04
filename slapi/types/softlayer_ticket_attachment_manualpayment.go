package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Ticket_Attachment_ManualPayment - This datatype contains tickets referenced from manual
// payments
type SoftLayer_Ticket_Attachment_ManualPayment struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// TicketId - The internal identifier of the ticket that an item is attached to.
	TicketId int `json:"ticketId,omitempty"`

	// AttachmentId - The internal identifier of an item that is attached to a ticket.
	AttachmentId int `json:"attachmentId,omitempty"`
}

func (softlayer_ticket_attachment_manualpayment *SoftLayer_Ticket_Attachment_ManualPayment) String() string {
	return "SoftLayer_Ticket_Attachment_ManualPayment"
}

// SoftLayer_Ticket_Attachment_ManualPayment_Extended is SoftLayer_Ticket_Attachment_ManualPayment with all maskable types.
type SoftLayer_Ticket_Attachment_ManualPayment_Extended struct {
	SoftLayer_Ticket_Attachment_ManualPayment

	// ScheduledAction - <nil>
	ScheduledAction *SoftLayer_Provisioning_Version1_Transaction `json:"scheduledAction,omitempty"`

	// Ticket - no documentation
	Ticket *SoftLayer_Ticket `json:"ticket,omitempty"`

	// AssignedAgent - <nil>
	AssignedAgent *SoftLayer_User_Customer `json:"assignedAgent,omitempty"`

	// Resource - no documentation
	Resource *SoftLayer_Billing_Payment_Card_ManualPayment `json:"resource,omitempty"`
}

func (softlayer_ticket_attachment_manualpayment *SoftLayer_Ticket_Attachment_ManualPayment_Extended) String() string {
	return "SoftLayer_Ticket_Attachment_ManualPayment"
}

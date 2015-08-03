package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Ticket_Attachment_Scheduled_Action - <nil>
type SoftLayer_Ticket_Attachment_Scheduled_Action struct {

	// RunDate - The internal identifier of a scheduled action transaction that is attached to a ticket.
	RunDate *time.Time `json:"runDate"`

	// TransactionId - The internal identifier of a scheduled action transaction that is attached to a
	// ticket.
	TransactionId int `json:"transactionId"`
}

func (softlayer_ticket_attachment_scheduled_action *SoftLayer_Ticket_Attachment_Scheduled_Action) String() string {
	return "SoftLayer_Ticket_Attachment_Scheduled_Action"
}

// SoftLayer_Ticket_Attachment_Scheduled_Action_Extended is SoftLayer_Ticket_Attachment_Scheduled_Action with all maskable types.
type SoftLayer_Ticket_Attachment_Scheduled_Action_Extended struct {
	SoftLayer_Ticket_Attachment_Scheduled_Action

	// Transaction - <nil>
	Transaction *SoftLayer_Provisioning_Version1_Transaction `json:"transaction"`

	// Resource - <nil>
	Resource *SoftLayer_Provisioning_Version1_Transaction `json:"resource"`
}

func (softlayer_ticket_attachment_scheduled_action *SoftLayer_Ticket_Attachment_Scheduled_Action_Extended) String() string {
	return "SoftLayer_Ticket_Attachment_Scheduled_Action"
}

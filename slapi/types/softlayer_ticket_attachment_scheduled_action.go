package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Ticket_Attachment_Scheduled_Action - <nil>
type SoftLayer_Ticket_Attachment_Scheduled_Action struct {

	// Resource - <nil>
	Resource *SoftLayer_Provisioning_Version1_Transaction `json:"resource"`

	// RunDate - The internal identifier of a scheduled action transaction that is attached to a ticket.
	RunDate *time.Time `json:"runDate"`

	// Transaction - <nil>
	Transaction *SoftLayer_Provisioning_Version1_Transaction `json:"transaction"`

	// TransactionId - The internal identifier of a scheduled action transaction that is attached to a
	// ticket.
	TransactionId int `json:"transactionId"`
}

func (softlayer_ticket_attachment_scheduled_action *SoftLayer_Ticket_Attachment_Scheduled_Action) String() string {
	return "SoftLayer_Ticket_Attachment_Scheduled_Action"
}

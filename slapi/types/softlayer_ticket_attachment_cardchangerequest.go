package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Ticket_Attachment_CardChangeRequest - This datatype contains tickets referenced from card
// change request
type SoftLayer_Ticket_Attachment_CardChangeRequest struct {

	// Resource - The card change request that is attached to a ticket.
	Resource *SoftLayer_Billing_Payment_Card_ChangeRequest `json:"resource"`
}

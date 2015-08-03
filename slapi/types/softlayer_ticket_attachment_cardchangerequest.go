package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Ticket_Attachment_CardChangeRequest - This datatype contains tickets referenced from card
// change request
type SoftLayer_Ticket_Attachment_CardChangeRequest struct {

	// Resource - The card change request that is attached to a ticket.
	Resource *SoftLayer_Billing_Payment_Card_ChangeRequest `json:"resource"`
}

func (softlayer_ticket_attachment_cardchangerequest *SoftLayer_Ticket_Attachment_CardChangeRequest) String() string {
	return "SoftLayer_Ticket_Attachment_CardChangeRequest"
}

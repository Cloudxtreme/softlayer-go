package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Ticket_Attachment_CardChangeRequest - This datatype contains tickets referenced from card
// change request
type SoftLayer_Ticket_Attachment_CardChangeRequest struct {
}

func (softlayer_ticket_attachment_cardchangerequest *SoftLayer_Ticket_Attachment_CardChangeRequest) String() string {
	return "SoftLayer_Ticket_Attachment_CardChangeRequest"
}

// SoftLayer_Ticket_Attachment_CardChangeRequest_Extended is SoftLayer_Ticket_Attachment_CardChangeRequest with all maskable types.
type SoftLayer_Ticket_Attachment_CardChangeRequest_Extended struct {
	SoftLayer_Ticket_Attachment_CardChangeRequest

	// Resource - The card change request that is attached to a ticket.
	Resource *SoftLayer_Billing_Payment_Card_ChangeRequest `json:"resource"`
}

func (softlayer_ticket_attachment_cardchangerequest *SoftLayer_Ticket_Attachment_CardChangeRequest_Extended) String() string {
	return "SoftLayer_Ticket_Attachment_CardChangeRequest"
}

package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Ticket_Attachment_ManualPayment - This datatype contains tickets referenced from manual
// payments
type SoftLayer_Ticket_Attachment_ManualPayment struct {
}

// SoftLayer_Ticket_Attachment_ManualPayment_Extended is SoftLayer_Ticket_Attachment_ManualPayment with all maskable types.
type SoftLayer_Ticket_Attachment_ManualPayment_Extended struct {
	SoftLayer_Ticket_Attachment_ManualPayment

	// Resource - no documentation
	Resource *SoftLayer_Billing_Payment_Card_ManualPayment `json:"resource"`
}

func (softlayer_ticket_attachment_manualpayment *SoftLayer_Ticket_Attachment_ManualPayment) String() string {
	return "SoftLayer_Ticket_Attachment_ManualPayment"
}

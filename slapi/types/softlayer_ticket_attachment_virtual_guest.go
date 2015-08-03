package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Ticket_Attachment_Virtual_Guest - SoftLayer tickets have the ability to be associated with
// specific pieces of hardware in a customer's inventory. Attaching hardware to a ticket can greatly
// increase response time from SoftLayer for issues that are related to one or more specific servers on
// a customer's account. The SoftLayer_Ticket_Attachment_Hardware data type models the relationship
// between a piece of hardware and a ticket. Only one attachment record may exist per hardware item per
// ticket.
type SoftLayer_Ticket_Attachment_Virtual_Guest struct {

	// Resource - The virtualized guest or CloudLayer Computing Instance that is attached to a ticket.
	Resource *SoftLayer_Virtual_Guest `json:"resource"`

	// VirtualGuest - The virtualized guest or CloudLayer Computing Instance that is attached to a ticket.
	VirtualGuest *SoftLayer_Virtual_Guest `json:"virtualGuest"`

	// VirtualGuestId - The internal identifier of the virtualized guest or CloudLayer Computing Instance
	// that is attached to a ticket.
	VirtualGuestId int `json:"virtualGuestId"`
}

func (softlayer_ticket_attachment_virtual_guest *SoftLayer_Ticket_Attachment_Virtual_Guest) String() string {
	return "SoftLayer_Ticket_Attachment_Virtual_Guest"
}

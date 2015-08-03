package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Ticket_Attachment_Hardware - SoftLayer tickets have the ability to be associated with
// specific pieces of hardware in a customer's inventory. Attaching hardware to a ticket can greatly
// increase response time from SoftLayer for issues that are related to one or more specific servers on
// a customer's account. The SoftLayer_Ticket_Attachment_Hardware data type models the relationship
// between a piece of hardware and a ticket. Only one attachment record may exist per hardware item per
// ticket.
type SoftLayer_Ticket_Attachment_Hardware struct {

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// HardwareId - The internal identifier of a piece of hardware that is attached to a ticket.
	HardwareId int `json:"hardwareId"`

	// Resource - no documentation
	Resource *SoftLayer_Hardware `json:"resource"`
}

func (softlayer_ticket_attachment_hardware *SoftLayer_Ticket_Attachment_Hardware) String() string {
	return "SoftLayer_Ticket_Attachment_Hardware"
}

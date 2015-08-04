package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Ticket_Update_Type - <nil>
type SoftLayer_Ticket_Update_Type struct {

	// Description - <nil>
	Description string `json:"description,omitempty"`

	// KeyName - <nil>
	KeyName string `json:"keyName,omitempty"`

	// Ticket - <nil>
	Ticket *SoftLayer_Ticket_Update `json:"ticket,omitempty"`
}

func (softlayer_ticket_update_type *SoftLayer_Ticket_Update_Type) String() string {
	return "SoftLayer_Ticket_Update_Type"
}

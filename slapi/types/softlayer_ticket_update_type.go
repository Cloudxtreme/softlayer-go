package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Ticket_Update_Type - <nil>
type SoftLayer_Ticket_Update_Type struct {

	// KeyName - <nil>
	KeyName string `json:"keyName,omitempty"`

	// Description - <nil>
	Description string `json:"description,omitempty"`

	// Ticket - <nil>
	Ticket *SoftLayer_Ticket_Update `json:"ticket,omitempty"`
}

func (softlayer_ticket_update_type *SoftLayer_Ticket_Update_Type) String() string {
	return "SoftLayer_Ticket_Update_Type"
}

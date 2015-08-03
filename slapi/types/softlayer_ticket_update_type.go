package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Ticket_Update_Type - <nil>
type SoftLayer_Ticket_Update_Type struct {

	// Description - <nil>
	Description string `json:"description"`

	// KeyName - <nil>
	KeyName string `json:"keyName"`
}

func (softlayer_ticket_update_type *SoftLayer_Ticket_Update_Type) String() string {
	return "SoftLayer_Ticket_Update_Type"
}

// SoftLayer_Ticket_Update_Type_Extended is SoftLayer_Ticket_Update_Type with all maskable types.
type SoftLayer_Ticket_Update_Type_Extended struct {
	SoftLayer_Ticket_Update_Type

	// Ticket - <nil>
	Ticket *SoftLayer_Ticket_Update `json:"ticket"`
}

func (softlayer_ticket_update_type *SoftLayer_Ticket_Update_Type_Extended) String() string {
	return "SoftLayer_Ticket_Update_Type"
}

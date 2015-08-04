package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Ticket_State - <nil>
type SoftLayer_Ticket_State struct {

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// StateTypeId - <nil>
	StateTypeId int `json:"stateTypeId,omitempty"`

	// TicketId - <nil>
	TicketId int `json:"ticketId,omitempty"`

	// StateType - <nil>
	StateType *SoftLayer_Ticket_State_Type `json:"stateType,omitempty"`

	// Ticket - <nil>
	Ticket *SoftLayer_Ticket `json:"ticket,omitempty"`
}

func (softlayer_ticket_state *SoftLayer_Ticket_State) String() string {
	return "SoftLayer_Ticket_State"
}

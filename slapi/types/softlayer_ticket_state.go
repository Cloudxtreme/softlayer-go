package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Ticket_State - <nil>
type SoftLayer_Ticket_State struct {

	// StateTypeId - <nil>
	StateTypeId int `json:"stateTypeId,omitempty"`

	// TicketId - <nil>
	TicketId int `json:"ticketId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`
}

func (softlayer_ticket_state *SoftLayer_Ticket_State) String() string {
	return "SoftLayer_Ticket_State"
}

// SoftLayer_Ticket_State_Extended is SoftLayer_Ticket_State with all maskable types.
type SoftLayer_Ticket_State_Extended struct {
	SoftLayer_Ticket_State

	// StateType - <nil>
	StateType *SoftLayer_Ticket_State_Type `json:"stateType,omitempty"`

	// Ticket - <nil>
	Ticket *SoftLayer_Ticket `json:"ticket,omitempty"`
}

func (softlayer_ticket_state *SoftLayer_Ticket_State_Extended) String() string {
	return "SoftLayer_Ticket_State"
}

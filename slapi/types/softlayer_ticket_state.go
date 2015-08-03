package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Ticket_State - <nil>
type SoftLayer_Ticket_State struct {

	// Id - <nil>
	Id int `json:"id"`

	// StateType - <nil>
	StateType *SoftLayer_Ticket_State_Type `json:"stateType"`

	// StateTypeId - <nil>
	StateTypeId int `json:"stateTypeId"`

	// Ticket - <nil>
	Ticket *SoftLayer_Ticket `json:"ticket"`

	// TicketId - <nil>
	TicketId int `json:"ticketId"`
}

func (softlayer_ticket_state *SoftLayer_Ticket_State) String() string {
	return "SoftLayer_Ticket_State"
}

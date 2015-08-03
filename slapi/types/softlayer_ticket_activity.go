package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Ticket_Activity - <nil>
type SoftLayer_Ticket_Activity struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// CreateTimestamp - <nil>
	CreateTimestamp *time.Time `json:"createTimestamp"`

	// Editor - <nil>
	Editor *SoftLayer_User_Interface `json:"editor"`

	// Id - <nil>
	Id int `json:"id"`

	// Ticket - <nil>
	Ticket *SoftLayer_Ticket `json:"ticket"`

	// TicketUpdate - <nil>
	TicketUpdate *SoftLayer_Ticket_Update `json:"ticketUpdate"`

	// Value - <nil>
	Value string `json:"value"`
}

func (softlayer_ticket_activity *SoftLayer_Ticket_Activity) String() string {
	return "SoftLayer_Ticket_Activity"
}

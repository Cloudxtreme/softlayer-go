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

	// Id - <nil>
	Id int `json:"id"`

	// Value - <nil>
	Value string `json:"value"`
}

// SoftLayer_Ticket_Activity_Extended is SoftLayer_Ticket_Activity with all maskable types.
type SoftLayer_Ticket_Activity_Extended struct {
	SoftLayer_Ticket_Activity

	// Editor - <nil>
	Editor *SoftLayer_User_Interface `json:"editor"`

	// Ticket - <nil>
	Ticket *SoftLayer_Ticket `json:"ticket"`

	// TicketUpdate - <nil>
	TicketUpdate *SoftLayer_Ticket_Update `json:"ticketUpdate"`
}

func (softlayer_ticket_activity *SoftLayer_Ticket_Activity) String() string {
	return "SoftLayer_Ticket_Activity"
}

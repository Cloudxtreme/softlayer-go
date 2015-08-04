package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Ticket_Activity - <nil>
type SoftLayer_Ticket_Activity struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// CreateTimestamp - <nil>
	CreateTimestamp *time.Time `json:"createTimestamp,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// Value - <nil>
	Value string `json:"value,omitempty"`

	// Editor - <nil>
	Editor *SoftLayer_User_Interface `json:"editor,omitempty"`

	// Ticket - <nil>
	Ticket *SoftLayer_Ticket `json:"ticket,omitempty"`

	// TicketUpdate - <nil>
	TicketUpdate *SoftLayer_Ticket_Update `json:"ticketUpdate,omitempty"`
}

func (softlayer_ticket_activity *SoftLayer_Ticket_Activity) String() string {
	return "SoftLayer_Ticket_Activity"
}

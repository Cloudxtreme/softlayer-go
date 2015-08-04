package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Ticket_Chat - <nil>
type SoftLayer_Ticket_Chat struct {

	// EndDate - <nil>
	EndDate *time.Time `json:"endDate,omitempty"`

	// StartDate - <nil>
	StartDate *time.Time `json:"startDate,omitempty"`

	// Transcript - <nil>
	Transcript string `json:"transcript,omitempty"`

	// CustomerId - <nil>
	CustomerId int `json:"customerId,omitempty"`

	// Agent - <nil>
	Agent *SoftLayer_User_Employee `json:"agent,omitempty"`

	// Customer - <nil>
	Customer *SoftLayer_User_Customer `json:"customer,omitempty"`

	// TicketUpdate - <nil>
	TicketUpdate *SoftLayer_Ticket_Update_Chat `json:"ticketUpdate,omitempty"`
}

func (softlayer_ticket_chat *SoftLayer_Ticket_Chat) String() string {
	return "SoftLayer_Ticket_Chat"
}

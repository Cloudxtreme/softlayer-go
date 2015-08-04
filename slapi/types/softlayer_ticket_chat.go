package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Ticket_Chat - <nil>
type SoftLayer_Ticket_Chat struct {

	// CustomerId - <nil>
	CustomerId int `json:"customerId,omitempty"`

	// EndDate - <nil>
	EndDate *time.Time `json:"endDate,omitempty"`

	// StartDate - <nil>
	StartDate *time.Time `json:"startDate,omitempty"`

	// Transcript - <nil>
	Transcript string `json:"transcript,omitempty"`
}

func (softlayer_ticket_chat *SoftLayer_Ticket_Chat) String() string {
	return "SoftLayer_Ticket_Chat"
}

// SoftLayer_Ticket_Chat_Extended is SoftLayer_Ticket_Chat with all maskable types.
type SoftLayer_Ticket_Chat_Extended struct {
	SoftLayer_Ticket_Chat

	// Agent - <nil>
	Agent *SoftLayer_User_Employee `json:"agent,omitempty"`

	// Customer - <nil>
	Customer *SoftLayer_User_Customer `json:"customer,omitempty"`

	// TicketUpdate - <nil>
	TicketUpdate *SoftLayer_Ticket_Update_Chat `json:"ticketUpdate,omitempty"`
}

func (softlayer_ticket_chat *SoftLayer_Ticket_Chat_Extended) String() string {
	return "SoftLayer_Ticket_Chat"
}

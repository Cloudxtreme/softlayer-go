package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Ticket_Chat - <nil>
type SoftLayer_Ticket_Chat struct {

	// Agent - <nil>
	Agent *SoftLayer_User_Employee `json:"agent"`

	// Customer - <nil>
	Customer *SoftLayer_User_Customer `json:"customer"`

	// CustomerId - <nil>
	CustomerId int `json:"customerId"`

	// EndDate - <nil>
	EndDate *time.Time `json:"endDate"`

	// StartDate - <nil>
	StartDate *time.Time `json:"startDate"`

	// TicketUpdate - <nil>
	TicketUpdate *SoftLayer_Ticket_Update_Chat `json:"ticketUpdate"`

	// Transcript - <nil>
	Transcript string `json:"transcript"`
}

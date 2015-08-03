package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Ticket_Survey_Preference - <nil>
type SoftLayer_Container_Ticket_Survey_Preference struct {

	// OptedOutDate - <nil>
	OptedOutDate *time.Time `json:"optedOutDate"`

	// Applicable - <nil>
	Applicable bool `json:"applicable"`

	// OptedOut - <nil>
	OptedOut bool `json:"optedOut"`
}

func (softlayer_container_ticket_survey_preference *SoftLayer_Container_Ticket_Survey_Preference) String() string {
	return "SoftLayer_Container_Ticket_Survey_Preference"
}

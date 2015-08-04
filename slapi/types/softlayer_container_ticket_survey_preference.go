package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Ticket_Survey_Preference - <nil>
type SoftLayer_Container_Ticket_Survey_Preference struct {

	// Applicable - <nil>
	Applicable bool `json:"applicable,omitempty"`

	// OptedOut - <nil>
	OptedOut bool `json:"optedOut,omitempty"`

	// OptedOutDate - <nil>
	OptedOutDate *time.Time `json:"optedOutDate,omitempty"`
}

func (softlayer_container_ticket_survey_preference *SoftLayer_Container_Ticket_Survey_Preference) String() string {
	return "SoftLayer_Container_Ticket_Survey_Preference"
}

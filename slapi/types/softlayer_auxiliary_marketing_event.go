package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Auxiliary_Marketing_Event - <nil>
type SoftLayer_Auxiliary_Marketing_Event struct {

	// Title - <nil>
	Title string `json:"title"`

	// Url - <nil>
	Url string `json:"url"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// EnabledFlag - <nil>
	EnabledFlag int `json:"enabledFlag"`

	// EndDate - <nil>
	EndDate *time.Time `json:"endDate"`

	// Location - <nil>
	Location string `json:"location"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// StartDate - <nil>
	StartDate *time.Time `json:"startDate"`
}

func (softlayer_auxiliary_marketing_event *SoftLayer_Auxiliary_Marketing_Event) String() string {
	return "SoftLayer_Auxiliary_Marketing_Event"
}

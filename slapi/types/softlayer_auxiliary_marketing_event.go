package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Auxiliary_Marketing_Event - <nil>
type SoftLayer_Auxiliary_Marketing_Event struct {

	// EnabledFlag - <nil>
	EnabledFlag int `json:"enabledFlag,omitempty"`

	// EndDate - <nil>
	EndDate *time.Time `json:"endDate,omitempty"`

	// Location - <nil>
	Location string `json:"location,omitempty"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// StartDate - <nil>
	StartDate *time.Time `json:"startDate,omitempty"`

	// Title - <nil>
	Title string `json:"title,omitempty"`

	// Url - <nil>
	Url string `json:"url,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`
}

func (softlayer_auxiliary_marketing_event *SoftLayer_Auxiliary_Marketing_Event) String() string {
	return "SoftLayer_Auxiliary_Marketing_Event"
}

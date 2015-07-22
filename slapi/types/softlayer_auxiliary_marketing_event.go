package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Auxiliary_Marketing_Event - <nil>
type SoftLayer_Auxiliary_Marketing_Event struct {

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

	// Title - <nil>
	Title string `json:"title"`

	// Url - <nil>
	Url string `json:"url"`
}

// GetMarketingEvents - This method will return a collection of SoftLayer_Auxiliary_Marketing_Event
// objects ordered in ascending order by start date.
func (softlayer_auxiliary_marketing_event *SoftLayer_Auxiliary_Marketing_Event) GetMarketingEvents(ctx *slapi.RequestContext) ([]*SoftLayer_Auxiliary_Marketing_Event, error) {
	var returnValue []*SoftLayer_Auxiliary_Marketing_Event
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_auxiliary_marketing_event *SoftLayer_Auxiliary_Marketing_Event) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Auxiliary_Marketing_Event, error) {
	var returnValue *SoftLayer_Auxiliary_Marketing_Event
	return returnValue, nil
}

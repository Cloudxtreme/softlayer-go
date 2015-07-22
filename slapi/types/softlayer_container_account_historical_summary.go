package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Account_Historical_Summary - Historical Summary Container for account resource
// details
type SoftLayer_Container_Account_Historical_Summary struct {

	// Details - no documentation
	Details []*SoftLayer_Container_Account_Historical_Summary_Detail `json:"details"`

	// EndDate - no documentation
	EndDate *time.Time `json:"endDate"`

	// StartDate - no documentation
	StartDate *time.Time `json:"startDate"`
}
package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Account_Historical_Summary_Detail - Historical Summary Details Container for a
// resource's data
type SoftLayer_Container_Account_Historical_Summary_Detail struct {

	// EndDate - no documentation
	EndDate *time.Time `json:"endDate"`

	// StartDate - no documentation
	StartDate *time.Time `json:"startDate"`
}

func (softlayer_container_account_historical_summary_detail *SoftLayer_Container_Account_Historical_Summary_Detail) String() string {
	return "SoftLayer_Container_Account_Historical_Summary_Detail"
}

package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Options - <nil>
type SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Options struct {

	// AggregatesOnly - <nil>
	AggregatesOnly bool `json:"aggregatesOnly"`

	// Category - <nil>
	Category string `json:"category"`

	// Days - <nil>
	Days int `json:"days"`

	// EndDate - <nil>
	EndDate *time.Time `json:"endDate"`

	// SelectedStatistics - <nil>
	SelectedStatistics []string `json:"selectedStatistics"`

	// StartDate - <nil>
	StartDate *time.Time `json:"startDate"`
}

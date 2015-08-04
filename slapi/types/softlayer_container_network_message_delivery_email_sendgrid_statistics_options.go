package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Options - <nil>
type SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Options struct {

	// Days - <nil>
	Days int `json:"days,omitempty"`

	// EndDate - <nil>
	EndDate *time.Time `json:"endDate,omitempty"`

	// SelectedStatistics - <nil>
	SelectedStatistics []string `json:"selectedStatistics,omitempty"`

	// StartDate - <nil>
	StartDate *time.Time `json:"startDate,omitempty"`

	// AggregatesOnly - <nil>
	AggregatesOnly bool `json:"aggregatesOnly,omitempty"`

	// Category - <nil>
	Category string `json:"category,omitempty"`
}

func (softlayer_container_network_message_delivery_email_sendgrid_statistics_options *SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Options) String() string {
	return "SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Options"
}

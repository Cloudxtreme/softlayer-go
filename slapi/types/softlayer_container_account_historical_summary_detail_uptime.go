package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Account_Historical_Summary_Detail_Uptime - Historical Summary Details Container
// for a host resource uptime
type SoftLayer_Container_Account_Historical_Summary_Detail_Uptime struct {

	// ConfigurationValue - no documentation
	ConfigurationValue *SoftLayer_Monitoring_Agent_Configuration_Value `json:"configurationValue,omitempty"`

	// Data - no documentation
	Data []*SoftLayer_Metric_Tracking_Object_Data `json:"data,omitempty"`

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`

	// EndDate - no documentation
	EndDate *time.Time `json:"endDate,omitempty"`

	// StartDate - no documentation
	StartDate *time.Time `json:"startDate,omitempty"`

	// CloudComputingInstance - no documentation
	CloudComputingInstance *SoftLayer_Virtual_Guest `json:"cloudComputingInstance,omitempty"`
}

func (softlayer_container_account_historical_summary_detail_uptime *SoftLayer_Container_Account_Historical_Summary_Detail_Uptime) String() string {
	return "SoftLayer_Container_Account_Historical_Summary_Detail_Uptime"
}

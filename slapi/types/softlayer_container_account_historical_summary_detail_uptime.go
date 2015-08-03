package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Account_Historical_Summary_Detail_Uptime - Historical Summary Details Container
// for a host resource uptime
type SoftLayer_Container_Account_Historical_Summary_Detail_Uptime struct {

	// ConfigurationValue - no documentation
	ConfigurationValue *SoftLayer_Monitoring_Agent_Configuration_Value `json:"configurationValue"`

	// Data - no documentation
	Data []*SoftLayer_Metric_Tracking_Object_Data `json:"data"`

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// CloudComputingInstance - no documentation
	CloudComputingInstance *SoftLayer_Virtual_Guest `json:"cloudComputingInstance"`
}

func (softlayer_container_account_historical_summary_detail_uptime *SoftLayer_Container_Account_Historical_Summary_Detail_Uptime) String() string {
	return "SoftLayer_Container_Account_Historical_Summary_Detail_Uptime"
}

package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Account_Historical_Summary_Detail_Uptime - Historical Summary Details Container
// for a host resource uptime
type SoftLayer_Container_Account_Historical_Summary_Detail_Uptime struct {

	// CloudComputingInstance - no documentation
	CloudComputingInstance *SoftLayer_Virtual_Guest `json:"cloudComputingInstance"`

	// ConfigurationValue - no documentation
	ConfigurationValue *SoftLayer_Monitoring_Agent_Configuration_Value `json:"configurationValue"`

	// Data - no documentation
	Data []*SoftLayer_Metric_Tracking_Object_Data `json:"data"`

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware"`
}
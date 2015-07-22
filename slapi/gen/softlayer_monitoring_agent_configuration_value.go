package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Monitoring_Agent_Configuration_Value - no documentation
type SoftLayer_Monitoring_Agent_Configuration_Value struct {

	// AgentId - Internal identifier of a monitoring agent that this configuration value belongs to
	AgentId int `json:"agentId"`

	// ConfigurationDefinitionId - Internal identifier of a monitoring configuration definition by which
	// this value is defined
	ConfigurationDefinitionId int `json:"configurationDefinitionId"`

	// Definition - <nil>
	Definition *SoftLayer_Configuration_Template_Section_Definition `json:"definition"`

	// Description - User-friendly description of a configuration value.
	Description string `json:"description"`

	// Id - Internal identifier of a monitoring configuration value
	Id int `json:"id"`

	// MetricDataType - The metric data type used to retrieve metric data currently being tracked.
	MetricDataType *SoftLayer_Container_Metric_Data_Type `json:"metricDataType"`

	// MonitoringAgent - <nil>
	MonitoringAgent *SoftLayer_Monitoring_Agent `json:"monitoringAgent"`

	// Profile - <nil>
	Profile *SoftLayer_Configuration_Template_Section_Profile `json:"profile"`

	// ProfileId - Internal identifier of a configuration profile. Configuration profile is associated with
	// a configuration section type of "Template section" A "Template section" defines skeleton
	// configuration definitions. For instance, if you want to monitor additional hard disks with
	// "CPU/Memory/Disk Monitoring Agent", you will have to add a new configuration profiles.
	ProfileId int `json:"profileId"`

	// Value - no documentation
	Value string `json:"value"`
}

// GetObject - <nil>
func (softlayer_monitoring_agent_configuration_value *SoftLayer_Monitoring_Agent_Configuration_Value) GetObject() (*SoftLayer_Monitoring_Agent_Configuration_Value, error) {
	var returnValue *SoftLayer_Monitoring_Agent_Configuration_Value
	return returnValue, nil
}

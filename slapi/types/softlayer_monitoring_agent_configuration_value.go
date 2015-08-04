package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Monitoring_Agent_Configuration_Value - no documentation
type SoftLayer_Monitoring_Agent_Configuration_Value struct {

	// Id - Internal identifier of a monitoring configuration value
	Id int `json:"id,omitempty"`

	// ProfileId - Internal identifier of a configuration profile. Configuration profile is associated with
	// a configuration section type of "Template section" A "Template section" defines skeleton
	// configuration definitions. For instance, if you want to monitor additional hard disks with
	// "CPU/Memory/Disk Monitoring Agent", you will have to add a new configuration profiles.
	ProfileId int `json:"profileId,omitempty"`

	// AgentId - Internal identifier of a monitoring agent that this configuration value belongs to
	AgentId int `json:"agentId,omitempty"`

	// ConfigurationDefinitionId - Internal identifier of a monitoring configuration definition by which
	// this value is defined
	ConfigurationDefinitionId int `json:"configurationDefinitionId,omitempty"`

	// Description - User-friendly description of a configuration value.
	Description string `json:"description,omitempty"`

	// Value - no documentation
	Value string `json:"value,omitempty"`
}

func (softlayer_monitoring_agent_configuration_value *SoftLayer_Monitoring_Agent_Configuration_Value) String() string {
	return "SoftLayer_Monitoring_Agent_Configuration_Value"
}

// SoftLayer_Monitoring_Agent_Configuration_Value_Extended is SoftLayer_Monitoring_Agent_Configuration_Value with all maskable types.
type SoftLayer_Monitoring_Agent_Configuration_Value_Extended struct {
	SoftLayer_Monitoring_Agent_Configuration_Value

	// Definition - <nil>
	Definition *SoftLayer_Configuration_Template_Section_Definition `json:"definition,omitempty"`

	// MetricDataType - The metric data type used to retrieve metric data currently being tracked.
	MetricDataType *SoftLayer_Container_Metric_Data_Type `json:"metricDataType,omitempty"`

	// Profile - <nil>
	Profile *SoftLayer_Configuration_Template_Section_Profile `json:"profile,omitempty"`

	// MonitoringAgent - <nil>
	MonitoringAgent *SoftLayer_Monitoring_Agent `json:"monitoringAgent,omitempty"`
}

func (softlayer_monitoring_agent_configuration_value *SoftLayer_Monitoring_Agent_Configuration_Value_Extended) String() string {
	return "SoftLayer_Monitoring_Agent_Configuration_Value"
}

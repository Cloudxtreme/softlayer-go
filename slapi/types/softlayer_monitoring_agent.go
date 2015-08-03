package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Monitoring_Agent - A monitoring agent object contains information describing the agent.
type SoftLayer_Monitoring_Agent struct {

	// AgentStatus - no documentation
	AgentStatus *SoftLayer_Monitoring_Agent_Status `json:"agentStatus"`

	// ConfigurationProfileCount - A count of all custom configuration profiles associated with the
	// corresponding agent
	ConfigurationProfileCount uint64 `json:"configurationProfileCount"`

	// ConfigurationProfiles - All custom configuration profiles associated with the corresponding agent
	ConfigurationProfiles []*SoftLayer_Configuration_Template_Section_Profile `json:"configurationProfiles"`

	// ConfigurationTemplate - A template of an agent's current configuration which contains information
	// about the structure of the configuration values.
	ConfigurationTemplate *SoftLayer_Configuration_Template `json:"configurationTemplate"`

	// ConfigurationTemplateId - Internal identifier of a configuration template that is used to configure
	// this agent
	ConfigurationTemplateId int `json:"configurationTemplateId"`

	// ConfigurationValueCount - A count of the values associated with the corresponding Agent
	// configuration.
	ConfigurationValueCount uint64 `json:"configurationValueCount"`

	// ConfigurationValues - The values associated with the corresponding Agent configuration.
	ConfigurationValues []*SoftLayer_Monitoring_Agent_Configuration_Value `json:"configurationValues"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// ProductItem - Contains general information relating to a single SoftLayer product.
	ProductItem *SoftLayer_Product_Item `json:"productItem"`

	// RemoteMonitoringAgentFlag - Indicates if this monitoring agent resides on your local box or on a
	// SoftLayer monitoring cluster.
	RemoteMonitoringAgentFlag bool `json:"remoteMonitoringAgentFlag"`

	// RobotId - Internal identifier of a monitoring robot that this agent belongs to
	RobotId int `json:"robotId"`

	// SoftwareDescription - A description for a specific installation of a Software Component
	SoftwareDescription *SoftLayer_Software_Description `json:"softwareDescription"`

	// StatusId - no documentation
	StatusId int `json:"statusId"`

	// StatusName - no documentation
	StatusName string `json:"statusName"`

	// VirtualGuest - Softlayer_Virtual_Guest object related to the monitoring agent, which this virtual
	// guest object and hardware is on the server of the running agent.
	VirtualGuest *SoftLayer_Virtual_Guest `json:"virtualGuest"`
}

func (softlayer_monitoring_agent *SoftLayer_Monitoring_Agent) String() string {
	return "SoftLayer_Monitoring_Agent"
}

package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Monitoring_Agent - A monitoring agent object contains information describing the agent.
type SoftLayer_Monitoring_Agent struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`

	// Name - no documentation
	Name string `json:"name"`

	// ConfigurationTemplateId - Internal identifier of a configuration template that is used to configure
	// this agent
	ConfigurationTemplateId int `json:"configurationTemplateId"`

	// RobotId - Internal identifier of a monitoring robot that this agent belongs to
	RobotId int `json:"robotId"`

	// StatusId - no documentation
	StatusId int `json:"statusId"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// RemoteMonitoringAgentFlag - Indicates if this monitoring agent resides on your local box or on a
	// SoftLayer monitoring cluster.
	RemoteMonitoringAgentFlag bool `json:"remoteMonitoringAgentFlag"`
}

func (softlayer_monitoring_agent *SoftLayer_Monitoring_Agent) String() string {
	return "SoftLayer_Monitoring_Agent"
}

// SoftLayer_Monitoring_Agent_Extended is SoftLayer_Monitoring_Agent with all maskable types.
type SoftLayer_Monitoring_Agent_Extended struct {
	SoftLayer_Monitoring_Agent

	// ConfigurationProfiles - All custom configuration profiles associated with the corresponding agent
	ConfigurationProfiles []*SoftLayer_Configuration_Template_Section_Profile `json:"configurationProfiles"`

	// ConfigurationProfileCount - A count of all custom configuration profiles associated with the
	// corresponding agent
	ConfigurationProfileCount uint64 `json:"configurationProfileCount"`

	// ConfigurationValueCount - A count of the values associated with the corresponding Agent
	// configuration.
	ConfigurationValueCount uint64 `json:"configurationValueCount"`

	// AgentStatus - no documentation
	AgentStatus *SoftLayer_Monitoring_Agent_Status `json:"agentStatus"`

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// StatusName - no documentation
	StatusName string `json:"statusName"`

	// VirtualGuest - Softlayer_Virtual_Guest object related to the monitoring agent, which this virtual
	// guest object and hardware is on the server of the running agent.
	VirtualGuest *SoftLayer_Virtual_Guest `json:"virtualGuest"`

	// ConfigurationTemplate - A template of an agent's current configuration which contains information
	// about the structure of the configuration values.
	ConfigurationTemplate *SoftLayer_Configuration_Template `json:"configurationTemplate"`

	// ProductItem - Contains general information relating to a single SoftLayer product.
	ProductItem *SoftLayer_Product_Item `json:"productItem"`

	// SoftwareDescription - A description for a specific installation of a Software Component
	SoftwareDescription *SoftLayer_Software_Description `json:"softwareDescription"`

	// ConfigurationValues - The values associated with the corresponding Agent configuration.
	ConfigurationValues []*SoftLayer_Monitoring_Agent_Configuration_Value `json:"configurationValues"`
}

func (softlayer_monitoring_agent *SoftLayer_Monitoring_Agent_Extended) String() string {
	return "SoftLayer_Monitoring_Agent"
}

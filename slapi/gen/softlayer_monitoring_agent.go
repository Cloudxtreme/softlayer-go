package sl

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

// Activate - This method activates a SoftLayer_Monitoring_Agent.
func (softlayer_monitoring_agent *SoftLayer_Monitoring_Agent) Activate() (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AddConfigurationProfile - This method is used to apply changes to a monitoring agent's configuration
// for SoftLayer_Configuration_Template_Section with the property sectionType that has a keyName of
// Configuration values that are passed in can be new or updated objects but must have a definitionId
// and profileId defined for both. Existing SoftLayer_Monitoring_Agent_Configuration_Value values can
// be retrieved as a property of the SoftLayer_Configuration_Template_Section_Definition's from the
// monitoring agent's configurationTemplate property. New values will follow the structure of
// SoftLayer_Monitoring_Agent_Configuration_Value. It returns a
// SoftLayer_Provisioning_Version1_Transaction object to track the progress of the update being
// applied. Some configuration sections act as a template which helps to create additional monitoring
// configurations. For instance, Core Resource monitoring agent lets you create monitoring
// configurations for different disk volumes or disk path.
func (softlayer_monitoring_agent *SoftLayer_Monitoring_Agent) AddConfigurationProfile(configurationValues []SoftLayer_Monitoring_Agent_Configuration_Value) (*SoftLayer_Provisioning_Version1_Transaction, error) {
	var returnValue *SoftLayer_Provisioning_Version1_Transaction
	return returnValue, nil
}

// ApplyConfigurationValues - This method creates a transaction used to apply changes to a monitoring
// agent's configuration for an array of SoftLayer_Configuration_Template_Section that have the
// property sectionType with a name of 'Fixed section'. Configuration values that are passed in can be
// new or updated objects but must have a configurationDefinitionId defined for both. Existing
// SoftLayer_Monitoring_Agent_Configuration_Value values can be retrieved as a property of the
// SoftLayer_Configuration_Template_Section_Definition from the monitoring agent's
// configurationTemplate property. New values will follow the structure of
// SoftLayer_Monitoring_Agent_Configuration_Value. This method returns a
// SoftLayer_Provisioning_Version1_Transaction object to track the progress of the update being
// applied.
func (softlayer_monitoring_agent *SoftLayer_Monitoring_Agent) ApplyConfigurationValues(configurationValues []SoftLayer_Monitoring_Agent_Configuration_Value) (*SoftLayer_Provisioning_Version1_Transaction, error) {
	var returnValue *SoftLayer_Provisioning_Version1_Transaction
	return returnValue, nil
}

// Deactivate - This method will deactivate the monitoring agent, preventing it from generating any
// further alarms.
func (softlayer_monitoring_agent *SoftLayer_Monitoring_Agent) Deactivate() (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteConfigurationProfile - This method will remove a
// SoftLayer_Configuration_Template_Section_Profile from a SoftLayer_Configuration_Template_Section by
// passing in the sectionId of the profile object and identifier of the profile. This will execute the
// action immediately on the server and the SoftLayer_Configuration_Template_Section returning a
// boolean true if successful.
func (softlayer_monitoring_agent *SoftLayer_Monitoring_Agent) DeleteConfigurationProfile(sectionId int, profileId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeployMonitoringAgent - Initialize a monitoring agent and deploy it with the
// SoftLayer_Configuration_Template with the same identifier as the $configurationTemplateId parameter.
// If the configuration template ID is not provided, the current configuration template will be used.
// When executing this method, the existing configuration values will be lost. If no configuration
// template identifier is provided, the current configuration template will be used. '''Warning'''
// Reporting data may be lost as a result of executing this method.
func (softlayer_monitoring_agent *SoftLayer_Monitoring_Agent) DeployMonitoringAgent(configurationTemplateId int) (*SoftLayer_Provisioning_Version1_Transaction, error) {
	var returnValue *SoftLayer_Provisioning_Version1_Transaction
	return returnValue, nil
}

// GetActiveAlarmSubscribers - This method retrieves an array of SoftLayer_Notification_User_Subscriber
// objects belonging to the SoftLayer_Monitoring_Agent which are able to receive alarm notifications.
func (softlayer_monitoring_agent *SoftLayer_Monitoring_Agent) GetActiveAlarmSubscribers() ([]*SoftLayer_Notification_User_Subscriber, error) {
	var returnValue []*SoftLayer_Notification_User_Subscriber
	return returnValue, nil
}

// GetAvailableConfigurationTemplates - This method returns an array of available
// SoftLayer_Configuration_Template objects for this monitoring agent.
func (softlayer_monitoring_agent *SoftLayer_Monitoring_Agent) GetAvailableConfigurationTemplates() ([]*SoftLayer_Configuration_Template, error) {
	var returnValue []*SoftLayer_Configuration_Template
	return returnValue, nil
}

// GetAvailableConfigurationValues - Returns an array of available configuration values that are
// specific to a server or a Virtual that this monitoring agent is running on. For example, invoking
// this method against "Network Traffic Monitoring Agent" will return all available network adapters on
// your system.
func (softlayer_monitoring_agent *SoftLayer_Monitoring_Agent) GetAvailableConfigurationValues(configurationDefinitionId int, configValues []SoftLayer_Monitoring_Agent_Configuration_Value) ([]*SoftLayer_Monitoring_Agent_Configuration_Value, error) {
	var returnValue []*SoftLayer_Monitoring_Agent_Configuration_Value
	return returnValue, nil
}

// GetEligibleAlarmSubscibers - This method returns an array of SoftLayer_User_Customer objects,
// representing those who are allowed to be used as alarm subscribers.
func (softlayer_monitoring_agent *SoftLayer_Monitoring_Agent) GetEligibleAlarmSubscibers() ([]*SoftLayer_User_Customer, error) {
	var returnValue []*SoftLayer_User_Customer
	return returnValue, nil
}

// GetGraph - This method returns a SoftLayer_Container_Bandwidth_GraphOutputs object containing a
// base64 PNG string graph of the provided configuration values for the given begin and end dates.
func (softlayer_monitoring_agent *SoftLayer_Monitoring_Agent) GetGraph(configurationValues []SoftLayer_Monitoring_Agent_Configuration_Value, beginDate time.Time, endDate time.Time) (*SoftLayer_Container_Monitoring_Graph_Outputs, error) {
	var returnValue *SoftLayer_Container_Monitoring_Graph_Outputs
	return returnValue, nil
}

// GetGraphData - This method returns the metric data for each of the configuration values provided
// during the given time range.
func (softlayer_monitoring_agent *SoftLayer_Monitoring_Agent) GetGraphData(metricDataTypes []SoftLayer_Container_Metric_Data_Type, startDate time.Time, endDate time.Time) ([]*SoftLayer_Metric_Tracking_Object_Data, error) {
	var returnValue []*SoftLayer_Metric_Tracking_Object_Data
	return returnValue, nil
}

// GetObject - This method retrieves a monitoring agent whose identifier corresponds to the value
// provided in the initialization parameter passed to the SoftLayer_Monitoring_Agent service.
func (softlayer_monitoring_agent *SoftLayer_Monitoring_Agent) GetObject() (*SoftLayer_Monitoring_Agent, error) {
	var returnValue *SoftLayer_Monitoring_Agent
	return returnValue, nil
}

// RemoveActiveAlarmSubscriber - Use of this method will allow removing active subscribers from the
// monitoring agent. The agent subscribers can be managed within the portal from the "Alarm
// Subscribers" tab of the monitoring agent configuration.
func (softlayer_monitoring_agent *SoftLayer_Monitoring_Agent) RemoveActiveAlarmSubscriber(userRecordId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveAllAlarmSubscribers - Use of this method will allow removing all subscribers from the
// monitoring agent. The agent subscribers can be managed within the portal from the "Alarm
// Subscribers" tab of the monitoring agent configuration.
func (softlayer_monitoring_agent *SoftLayer_Monitoring_Agent) RemoveAllAlarmSubscribers() (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RestartMonitoringAgent - This method restarts a monitoring agent and sets the agent's status to
func (softlayer_monitoring_agent *SoftLayer_Monitoring_Agent) RestartMonitoringAgent() (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SetActiveAlarmSubscriber - This method assigns a user to receive the alerts generated by this
// SoftLayer_Monitoring_Agent.
func (softlayer_monitoring_agent *SoftLayer_Monitoring_Agent) SetActiveAlarmSubscriber(userRecordId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

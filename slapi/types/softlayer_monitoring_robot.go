package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Monitoring_Robot - The SoftLayer_Monitoring_Robot data type contains general information
// relating to a monitoring robot.
type SoftLayer_Monitoring_Robot struct {

	// Account - The account associated with the corresponding robot.
	Account *SoftLayer_Account `json:"account"`

	// AccountId - Internal identifier of a SoftLayer account that this robot belongs to
	AccountId int `json:"accountId"`

	// Id - no documentation
	Id int `json:"id"`

	// MonitoringAgentCount - A count of the program (monitoring agent) that gets details of a system or
	// application and reporting of the metric data and triggers alarms for predefined events.
	MonitoringAgentCount uint64 `json:"monitoringAgentCount"`

	// MonitoringAgents - The program (monitoring agent) that gets details of a system or application and
	// reporting of the metric data and triggers alarms for predefined events.
	MonitoringAgents []*SoftLayer_Monitoring_Agent `json:"monitoringAgents"`

	// Name - no documentation
	Name string `json:"name"`

	// RobotStatus - no documentation
	RobotStatus *SoftLayer_Monitoring_Robot_Status `json:"robotStatus"`

	// SoftwareComponent - The SoftLayer_Software_Component that corresponds to the robot installation on
	// the server.
	SoftwareComponent *SoftLayer_Software_Component `json:"softwareComponent"`

	// StatusId - no documentation
	StatusId int `json:"statusId"`
}

// CheckConnection - Checks if a monitoring robot can communicate with SoftLayer monitoring management
// system via the private network. TCP port 48000 - 48002 must be open on your server or your virtual
// server in order for this test to succeed.
func (softlayer_monitoring_robot *SoftLayer_Monitoring_Robot) CheckConnection(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeployMonitoringAgents - <nil>
func (softlayer_monitoring_robot *SoftLayer_Monitoring_Robot) DeployMonitoringAgents(commonOptions *slapi.CommonOptions, configurationTemplateGroup SoftLayer_Monitoring_Agent_Configuration_Template_Group) (*SoftLayer_Provisioning_Version1_Transaction, error) {
	var returnValue *SoftLayer_Provisioning_Version1_Transaction
	return returnValue, nil
}

// GetAvailableConfigurationGroups - Returns available configuration template groups for this
// monitoring agent.
func (softlayer_monitoring_robot *SoftLayer_Monitoring_Robot) GetAvailableConfigurationGroups(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Monitoring_Agent_Configuration_Template_Group, error) {
	var returnValue []*SoftLayer_Monitoring_Agent_Configuration_Template_Group
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_monitoring_robot *SoftLayer_Monitoring_Robot) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Monitoring_Robot, error) {
	var returnValue *SoftLayer_Monitoring_Robot
	return returnValue, nil
}

// ResetStatus - If our monitoring management system is not able to connect to your monitoring robot,
// it sets the robot status to "Limited Connectivity". Robots in this status will not be process by our
// monitoring management system. You cannot manage monitoring agents either. Use this method to resets
// monitoring robot status to "Active" to indicate the connection issue is resolved.
func (softlayer_monitoring_robot *SoftLayer_Monitoring_Robot) ResetStatus(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

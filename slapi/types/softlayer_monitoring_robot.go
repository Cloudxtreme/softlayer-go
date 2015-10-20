package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Monitoring_Robot - The SoftLayer_Monitoring_Robot data type contains general information
// relating to a monitoring robot.
type SoftLayer_Monitoring_Robot struct {

	// StatusId - no documentation
	StatusId int `json:"statusId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// AccountId - Internal identifier of a SoftLayer account that this robot belongs to
	AccountId int `json:"accountId,omitempty"`

	// MonitoringAgentCount - A count of the program (monitoring agent) that gets details of a system or
	// application and reporting of the metric data and triggers alarms for predefined events.
	MonitoringAgentCount uint64 `json:"monitoringAgentCount,omitempty"`

	// Account - The account associated with the corresponding robot.
	Account *SoftLayer_Account `json:"account,omitempty"`

	// MonitoringAgents - The program (monitoring agent) that gets details of a system or application and
	// reporting of the metric data and triggers alarms for predefined events.
	MonitoringAgents []*SoftLayer_Monitoring_Agent `json:"monitoringAgents,omitempty"`

	// SoftwareComponent - The SoftLayer_Software_Component that corresponds to the robot installation on
	// the server.
	SoftwareComponent *SoftLayer_Software_Component `json:"softwareComponent,omitempty"`

	// RobotStatus - no documentation
	RobotStatus *SoftLayer_Monitoring_Robot_Status `json:"robotStatus,omitempty"`
}

func (softlayer_monitoring_robot *SoftLayer_Monitoring_Robot) String() string {
	return "SoftLayer_Monitoring_Robot"
}

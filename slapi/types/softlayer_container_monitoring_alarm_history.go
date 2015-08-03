package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Monitoring_Alarm_History - The SoftLayer_Container_Monitoring_Alarm_History data
// type contains information relating to SoftLayer monitoring alarm history.
type SoftLayer_Container_Monitoring_Alarm_History struct {

	// ClosedDate - no documentation
	ClosedDate *time.Time `json:"closedDate"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Message - no documentation
	Message string `json:"message"`

	// RobotId - no documentation
	RobotId int `json:"robotId"`

	// Severity - no documentation
	Severity string `json:"severity"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// AgentId - ID of the monitoring agent that triggered this alarm
	AgentId int `json:"agentId"`

	// AlarmId - no documentation
	AlarmId string `json:"alarmId"`
}

func (softlayer_container_monitoring_alarm_history *SoftLayer_Container_Monitoring_Alarm_History) String() string {
	return "SoftLayer_Container_Monitoring_Alarm_History"
}

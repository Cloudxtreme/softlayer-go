package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Monitoring_Alarm_History - The SoftLayer_Container_Monitoring_Alarm_History data
// type contains information relating to SoftLayer monitoring alarm history.
type SoftLayer_Container_Monitoring_Alarm_History struct {

	// Message - no documentation
	Message string `json:"message,omitempty"`

	// RobotId - no documentation
	RobotId int `json:"robotId,omitempty"`

	// Severity - no documentation
	Severity string `json:"severity,omitempty"`

	// AccountId - no documentation
	AccountId int `json:"accountId,omitempty"`

	// AgentId - ID of the monitoring agent that triggered this alarm
	AgentId int `json:"agentId,omitempty"`

	// AlarmId - no documentation
	AlarmId string `json:"alarmId,omitempty"`

	// ClosedDate - no documentation
	ClosedDate *time.Time `json:"closedDate,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`
}

func (softlayer_container_monitoring_alarm_history *SoftLayer_Container_Monitoring_Alarm_History) String() string {
	return "SoftLayer_Container_Monitoring_Alarm_History"
}

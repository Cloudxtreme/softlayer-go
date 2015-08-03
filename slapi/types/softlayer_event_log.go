package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Event_Log - The SoftLayer_Event_Log data type contains an event detail occurred upon
// various SoftLayer resources.
type SoftLayer_Event_Log struct {

	// Label - no documentation
	Label string `json:"label"`

	// ObjectId - no documentation
	ObjectId int `json:"objectId"`

	// ObjectName - Event object name such as "server", "dns" and so on.
	ObjectName string `json:"objectName"`

	// Resource - A resource object that is associated with the event
	Resource *SoftLayer_Entity `json:"resource"`

	// Username - no documentation
	Username string `json:"username"`

	// EventName - Event name such as "reboot", "cancel", "update host" and so on.
	EventName string `json:"eventName"`

	// IpAddress - no documentation
	IpAddress string `json:"ipAddress"`

	// TraceId - A unique trace id. Multiple event can be grouped by a trace id.
	TraceId string `json:"traceId"`

	// UserType - Type of user that triggered the event. User type can be or
	UserType string `json:"userType"`

	// EventCreateDate - no documentation
	EventCreateDate *time.Time `json:"eventCreateDate"`

	// MetaData - no documentation
	MetaData string `json:"metaData"`

	// UserId - no documentation
	UserId int `json:"userId"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`
}

func (softlayer_event_log *SoftLayer_Event_Log) String() string {
	return "SoftLayer_Event_Log"
}

// SoftLayer_Event_Log_Extended is SoftLayer_Event_Log with all maskable types.
type SoftLayer_Event_Log_Extended struct {
	SoftLayer_Event_Log

	// User - <nil>
	User *SoftLayer_User_Customer `json:"user"`
}

func (softlayer_event_log *SoftLayer_Event_Log_Extended) String() string {
	return "SoftLayer_Event_Log"
}

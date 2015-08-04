package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Event_Log - The SoftLayer_Event_Log data type contains an event detail occurred upon
// various SoftLayer resources.
type SoftLayer_Event_Log struct {

	// Resource - A resource object that is associated with the event
	Resource *SoftLayer_Entity `json:"resource,omitempty"`

	// IpAddress - no documentation
	IpAddress string `json:"ipAddress,omitempty"`

	// Label - no documentation
	Label string `json:"label,omitempty"`

	// ObjectId - no documentation
	ObjectId int `json:"objectId,omitempty"`

	// UserType - Type of user that triggered the event. User type can be or
	UserType string `json:"userType,omitempty"`

	// Username - no documentation
	Username string `json:"username,omitempty"`

	// EventCreateDate - no documentation
	EventCreateDate *time.Time `json:"eventCreateDate,omitempty"`

	// ObjectName - Event object name such as "server", "dns" and so on.
	ObjectName string `json:"objectName,omitempty"`

	// UserId - no documentation
	UserId int `json:"userId,omitempty"`

	// EventName - Event name such as "reboot", "cancel", "update host" and so on.
	EventName string `json:"eventName,omitempty"`

	// TraceId - A unique trace id. Multiple event can be grouped by a trace id.
	TraceId string `json:"traceId,omitempty"`

	// AccountId - no documentation
	AccountId int `json:"accountId,omitempty"`

	// MetaData - no documentation
	MetaData string `json:"metaData,omitempty"`
}

func (softlayer_event_log *SoftLayer_Event_Log) String() string {
	return "SoftLayer_Event_Log"
}

// SoftLayer_Event_Log_Extended is SoftLayer_Event_Log with all maskable types.
type SoftLayer_Event_Log_Extended struct {
	SoftLayer_Event_Log

	// User - <nil>
	User *SoftLayer_User_Customer `json:"user,omitempty"`
}

func (softlayer_event_log *SoftLayer_Event_Log_Extended) String() string {
	return "SoftLayer_Event_Log"
}

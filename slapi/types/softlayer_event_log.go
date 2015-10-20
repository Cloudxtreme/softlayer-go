package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Event_Log - The SoftLayer_Event_Log data type contains an event detail occurred upon
// various SoftLayer resources.
type SoftLayer_Event_Log struct {

	// ObjectName - Event object name such as "server", "dns" and so on.
	ObjectName string `json:"objectName,omitempty"`

	// AccountId - no documentation
	AccountId int `json:"accountId,omitempty"`

	// Label - no documentation
	Label string `json:"label,omitempty"`

	// MetaData - no documentation
	MetaData string `json:"metaData,omitempty"`

	// Resource - A resource object that is associated with the event
	Resource *SoftLayer_Entity `json:"resource,omitempty"`

	// EventName - Event name such as "reboot", "cancel", "update host" and so on.
	EventName string `json:"eventName,omitempty"`

	// ObjectId - no documentation
	ObjectId int `json:"objectId,omitempty"`

	// UserType - Type of user that triggered the event. User type can be or
	UserType string `json:"userType,omitempty"`

	// Username - no documentation
	Username string `json:"username,omitempty"`

	// EventCreateDate - no documentation
	EventCreateDate *time.Time `json:"eventCreateDate,omitempty"`

	// IpAddress - no documentation
	IpAddress string `json:"ipAddress,omitempty"`

	// TraceId - A unique trace id. Multiple event can be grouped by a trace id.
	TraceId string `json:"traceId,omitempty"`

	// UserId - no documentation
	UserId int `json:"userId,omitempty"`

	// User - <nil>
	User *SoftLayer_User_Customer `json:"user,omitempty"`
}

func (softlayer_event_log *SoftLayer_Event_Log) String() string {
	return "SoftLayer_Event_Log"
}

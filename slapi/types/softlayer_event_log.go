package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Event_Log - The SoftLayer_Event_Log data type contains an event detail occurred upon
// various SoftLayer resources.
type SoftLayer_Event_Log struct {

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// EventCreateDate - no documentation
	EventCreateDate *time.Time `json:"eventCreateDate"`

	// EventName - Event name such as "reboot", "cancel", "update host" and so on.
	EventName string `json:"eventName"`

	// IpAddress - no documentation
	IpAddress string `json:"ipAddress"`

	// Label - no documentation
	Label string `json:"label"`

	// MetaData - no documentation
	MetaData string `json:"metaData"`

	// ObjectId - no documentation
	ObjectId int `json:"objectId"`

	// ObjectName - Event object name such as "server", "dns" and so on.
	ObjectName string `json:"objectName"`

	// Resource - A resource object that is associated with the event
	Resource *SoftLayer_Entity `json:"resource"`

	// TraceId - A unique trace id. Multiple event can be grouped by a trace id.
	TraceId string `json:"traceId"`

	// User - <nil>
	User *SoftLayer_User_Customer `json:"user"`

	// UserId - no documentation
	UserId int `json:"userId"`

	// UserType - Type of user that triggered the event. User type can be or
	UserType string `json:"userType"`

	// Username - no documentation
	Username string `json:"username"`
}

func (softlayer_event_log *SoftLayer_Event_Log) String() string {
	return "SoftLayer_Event_Log"
}

// GetAllEventNames - no documentation
func (softlayer_event_log *SoftLayer_Event_Log) GetAllEventNames(ctx *slapi.RequestContext, objectName string) ([]string, error) {
	var returnValue []string
	return returnValue, nil
}

// GetAllEventObjectNames - no documentation
func (softlayer_event_log *SoftLayer_Event_Log) GetAllEventObjectNames(ctx *slapi.RequestContext) ([]string, error) {
	var returnValue []string
	return returnValue, nil
}

// GetAllObjects - <nil>
func (softlayer_event_log *SoftLayer_Event_Log) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_Event_Log, error) {
	var returnValue []*SoftLayer_Event_Log
	return returnValue, nil
}

// GetAllUserTypes - <nil>
func (softlayer_event_log *SoftLayer_Event_Log) GetAllUserTypes(ctx *slapi.RequestContext) ([]string, error) {
	var returnValue []string
	return returnValue, nil
}

package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Monitor_Version1_Query_Result - The monitoring result object is used to show the
// status of the actions taken by the monitoring system. In general, only the responseStatus variable
// is needed, as it holds the information on the status of the service.
type SoftLayer_Network_Monitor_Version1_Query_Result struct {

	// FinishTime - no documentation
	FinishTime *time.Time `json:"finishTime"`

	// QueryHost - References the queryHost that this response relates to.
	QueryHost *SoftLayer_Network_Monitor_Version1_Query_Host `json:"queryHost"`

	// ResponseStatus - The response status for this server. The response status meanings are: 0:
	// Down/Critical: Server is down and/or has passed the critical response threshold (extremely long ping
	// response, abnormal behavior, etc.) 1: Warning - Server may be recovering from a previous down state,
	// or may have taken too long to respond 2: Up 3: Not used 4: Unknown - An unknown error has occurred.
	// If the problem persists, contact support. 5: Unknown - An unknown error has occurred. If the problem
	// persists, contact support.
	ResponseStatus int `json:"responseStatus"`

	// ResponseTime - no documentation
	ResponseTime float32 `json:"responseTime"`
}

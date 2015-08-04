package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Network_Monitor_Version1_Query_Result - The monitoring result object is used to show the
// status of the actions taken by the monitoring system. In general, only the responseStatus variable
// is needed, as it holds the information on the status of the service.
type SoftLayer_Network_Monitor_Version1_Query_Result struct {

	// FinishTime - no documentation
	FinishTime *time.Time `json:"finishTime,omitempty"`

	// ResponseStatus - The response status for this server. The response status meanings are: 0:
	// Down/Critical: Server is down and/or has passed the critical response threshold (extremely long ping
	// response, abnormal behavior, etc.) 1: Warning - Server may be recovering from a previous down state,
	// or may have taken too long to respond 2: Up 3: Not used 4: Unknown - An unknown error has occurred.
	// If the problem persists, contact support. 5: Unknown - An unknown error has occurred. If the problem
	// persists, contact support.
	ResponseStatus int `json:"responseStatus,omitempty"`

	// ResponseTime - no documentation
	ResponseTime slapi.Float64 `json:"responseTime,omitempty"`
}

func (softlayer_network_monitor_version1_query_result *SoftLayer_Network_Monitor_Version1_Query_Result) String() string {
	return "SoftLayer_Network_Monitor_Version1_Query_Result"
}

// SoftLayer_Network_Monitor_Version1_Query_Result_Extended is SoftLayer_Network_Monitor_Version1_Query_Result with all maskable types.
type SoftLayer_Network_Monitor_Version1_Query_Result_Extended struct {
	SoftLayer_Network_Monitor_Version1_Query_Result

	// QueryHost - References the queryHost that this response relates to.
	QueryHost *SoftLayer_Network_Monitor_Version1_Query_Host `json:"queryHost,omitempty"`
}

func (softlayer_network_monitor_version1_query_result *SoftLayer_Network_Monitor_Version1_Query_Result_Extended) String() string {
	return "SoftLayer_Network_Monitor_Version1_Query_Result"
}

package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Bandwidth_GraphInputs - SoftLayer_Container_Bandwidth_GraphInputs models a
// single inbound object for a given bandwidth graph.
type SoftLayer_Container_Bandwidth_GraphInputs struct {

	// EndDate - This is a unix timestamp that represents the stop date/time for a graph.
	EndDate *time.Time `json:"endDate"`

	// NetworkInterfaceId - The front-end or back-end network uplink interface associated with this server.
	NetworkInterfaceId int `json:"networkInterfaceId"`

	// Pod - <nil>
	Pod int `json:"pod"`

	// ServerName - This is a human readable name for the server or rack being graphed.
	ServerName string `json:"serverName"`

	// StartDate - This is a unix timestamp that represents the begin date/time for a graph.
	StartDate *time.Time `json:"startDate"`
}

func (softlayer_container_bandwidth_graphinputs *SoftLayer_Container_Bandwidth_GraphInputs) String() string {
	return "SoftLayer_Container_Bandwidth_GraphInputs"
}

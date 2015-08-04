package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Ticket_GraphInputs - SoftLayer_Container_Ticket_GraphInputs models a single
// inbound object for a given ticket graph.
type SoftLayer_Container_Ticket_GraphInputs struct {

	// EndDate - This is a unix timestamp that represents the stop date/time for a graph.
	EndDate *time.Time `json:"endDate,omitempty"`

	// NetworkInterfaceId - The front-end or back-end network uplink interface associated with this server.
	NetworkInterfaceId int `json:"networkInterfaceId,omitempty"`

	// Pod - <nil>
	Pod int `json:"pod,omitempty"`

	// ServerName - This is a human readable name for the server or rack being graphed.
	ServerName string `json:"serverName,omitempty"`

	// StartDate - This is a unix timestamp that represents the begin date/time for a graph.
	StartDate *time.Time `json:"startDate,omitempty"`
}

func (softlayer_container_ticket_graphinputs *SoftLayer_Container_Ticket_GraphInputs) String() string {
	return "SoftLayer_Container_Ticket_GraphInputs"
}

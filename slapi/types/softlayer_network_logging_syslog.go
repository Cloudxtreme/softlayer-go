package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Logging_Syslog - The Syslog class holds a single line from the Networking Firewall
// "Syslog" record, for firewall detected and blocked attempts on a server.
type SoftLayer_Network_Logging_Syslog struct {

	// CreateDate - Timestamp for when the connection was blocked by the firewall
	CreateDate *time.Time `json:"createDate"`

	// DestinationIpAddress - The Destination IP Address of the blocked connection (your end)
	DestinationIpAddress string `json:"destinationIpAddress"`

	// DestinationPort - The Destination Port of the blocked connection (your end)
	DestinationPort int `json:"destinationPort"`

	// EventType - This tells you what kind of firewall event this log line is for: accept or deny.
	EventType string `json:"eventType"`

	// Message - no documentation
	Message string `json:"message"`

	// Protocol - Connection protocol used to make the call that was blocked (tcp, udp, etc)
	Protocol string `json:"protocol"`

	// SourceIpAddress - The Source IP Address of the call that was blocked (attacker's end)
	SourceIpAddress string `json:"sourceIpAddress"`

	// SourcePort - The Source Port where the blocked connection was established (attacker's end)
	SourcePort int `json:"sourcePort"`

	// TotalEvents - If this is an aggregation of syslog events, this property shows the total events.
	TotalEvents int `json:"totalEvents"`
}

func (softlayer_network_logging_syslog *SoftLayer_Network_Logging_Syslog) String() string {
	return "SoftLayer_Network_Logging_Syslog"
}

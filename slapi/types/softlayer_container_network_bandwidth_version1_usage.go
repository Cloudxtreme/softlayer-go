package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Container_Network_Bandwidth_Version1_Usage -
// SoftLayer_Container_Network_Bandwidth_Version1_Usage models an hourly bandwidth record.
type SoftLayer_Container_Network_Bandwidth_Version1_Usage struct {

	// IncomingAmount - The amount of incoming bandwidth that a server has used within the hour of the
	// recordedDate.
	IncomingAmount slapi.Float64 `json:"incomingAmount,omitempty"`

	// OutgoingAmount - The amount of outgoing bandwidth that a server has used within the hour of the
	// recordedDate.
	OutgoingAmount slapi.Float64 `json:"outgoingAmount,omitempty"`

	// RecordedDate - The date and time that the bandwidth was used by a piece of hardware
	RecordedDate *time.Time `json:"recordedDate,omitempty"`
}

func (softlayer_container_network_bandwidth_version1_usage *SoftLayer_Container_Network_Bandwidth_Version1_Usage) String() string {
	return "SoftLayer_Container_Network_Bandwidth_Version1_Usage"
}

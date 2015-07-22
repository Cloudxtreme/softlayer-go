package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Network_Bandwidth_Version1_Usage -
// SoftLayer_Container_Network_Bandwidth_Version1_Usage models an hourly bandwidth record.
type SoftLayer_Container_Network_Bandwidth_Version1_Usage struct {

	// IncomingAmount - The amount of incoming bandwidth that a server has used within the hour of the
	// recordedDate.
	IncomingAmount float32 `json:"incomingAmount"`

	// OutgoingAmount - The amount of outgoing bandwidth that a server has used within the hour of the
	// recordedDate.
	OutgoingAmount float32 `json:"outgoingAmount"`

	// RecordedDate - The date and time that the bandwidth was used by a piece of hardware
	RecordedDate *time.Time `json:"recordedDate"`
}
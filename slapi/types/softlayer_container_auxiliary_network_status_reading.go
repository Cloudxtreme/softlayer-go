package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Container_Auxiliary_Network_Status_Reading - The
// SoftLayer_Container_Auxiliary_Network_Status_Reading data type contains information relating to an
// object being monitored from outside the SoftLayer network. It is primarily used to check the status
// of our edge routers from multiple locations around the world.
type SoftLayer_Container_Auxiliary_Network_Status_Reading struct {

	// AveragePing - no documentation
	AveragePing slapi.Float64 `json:"averagePing,omitempty"`

	// LastCheckDate - no documentation
	LastCheckDate *time.Time `json:"lastCheckDate,omitempty"`

	// Latency - The total response time in seconds calculated during the last check.
	Latency slapi.Float64 `json:"latency,omitempty"`

	// StatusCode - no documentation
	StatusCode string `json:"statusCode,omitempty"`

	// StatusMessage - no documentation
	StatusMessage string `json:"statusMessage,omitempty"`

	// Frequency - no documentation
	Frequency int `json:"frequency,omitempty"`

	// LastDownDate - no documentation
	LastDownDate *time.Time `json:"lastDownDate,omitempty"`

	// MaximumPing - no documentation
	MaximumPing slapi.Float64 `json:"maximumPing,omitempty"`

	// PingLoss - no documentation
	PingLoss slapi.Float64 `json:"pingLoss,omitempty"`

	// Target - no documentation
	Target string `json:"target,omitempty"`

	// Fails - Number of failures since the target was last detected to be working properly.
	Fails int `json:"fails,omitempty"`

	// Label - no documentation
	Label string `json:"label,omitempty"`

	// Location - no documentation
	Location string `json:"location,omitempty"`

	// StartDate - no documentation
	StartDate *time.Time `json:"startDate,omitempty"`

	// MinimumPing - no documentation
	MinimumPing slapi.Float64 `json:"minimumPing,omitempty"`

	// TargetType - no documentation
	TargetType string `json:"targetType,omitempty"`
}

func (softlayer_container_auxiliary_network_status_reading *SoftLayer_Container_Auxiliary_Network_Status_Reading) String() string {
	return "SoftLayer_Container_Auxiliary_Network_Status_Reading"
}

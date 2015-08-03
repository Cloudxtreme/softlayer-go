package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Auxiliary_Network_Status_Reading - The
// SoftLayer_Container_Auxiliary_Network_Status_Reading data type contains information relating to an
// object being monitored from outside the SoftLayer network. It is primarily used to check the status
// of our edge routers from multiple locations around the world.
type SoftLayer_Container_Auxiliary_Network_Status_Reading struct {

	// Fails - Number of failures since the target was last detected to be working properly.
	Fails int `json:"fails"`

	// PingLoss - no documentation
	PingLoss float32 `json:"pingLoss"`

	// StatusMessage - no documentation
	StatusMessage string `json:"statusMessage"`

	// Target - no documentation
	Target string `json:"target"`

	// AveragePing - no documentation
	AveragePing float32 `json:"averagePing"`

	// LastDownDate - no documentation
	LastDownDate *time.Time `json:"lastDownDate"`

	// Latency - The total response time in seconds calculated during the last check.
	Latency float32 `json:"latency"`

	// Frequency - no documentation
	Frequency int `json:"frequency"`

	// Label - no documentation
	Label string `json:"label"`

	// LastCheckDate - no documentation
	LastCheckDate *time.Time `json:"lastCheckDate"`

	// Location - no documentation
	Location string `json:"location"`

	// MinimumPing - no documentation
	MinimumPing float32 `json:"minimumPing"`

	// StatusCode - no documentation
	StatusCode string `json:"statusCode"`

	// TargetType - no documentation
	TargetType string `json:"targetType"`

	// MaximumPing - no documentation
	MaximumPing float32 `json:"maximumPing"`

	// StartDate - no documentation
	StartDate *time.Time `json:"startDate"`
}

func (softlayer_container_auxiliary_network_status_reading *SoftLayer_Container_Auxiliary_Network_Status_Reading) String() string {
	return "SoftLayer_Container_Auxiliary_Network_Status_Reading"
}

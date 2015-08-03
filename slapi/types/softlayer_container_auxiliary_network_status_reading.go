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

	// StatusMessage - no documentation
	StatusMessage string `json:"statusMessage"`

	// AveragePing - no documentation
	AveragePing float32 `json:"averagePing"`

	// Location - no documentation
	Location string `json:"location"`

	// StartDate - no documentation
	StartDate *time.Time `json:"startDate"`

	// Target - no documentation
	Target string `json:"target"`

	// LastCheckDate - no documentation
	LastCheckDate *time.Time `json:"lastCheckDate"`

	// MinimumPing - no documentation
	MinimumPing float32 `json:"minimumPing"`

	// MaximumPing - no documentation
	MaximumPing float32 `json:"maximumPing"`

	// Frequency - no documentation
	Frequency int `json:"frequency"`

	// Label - no documentation
	Label string `json:"label"`

	// LastDownDate - no documentation
	LastDownDate *time.Time `json:"lastDownDate"`

	// Latency - The total response time in seconds calculated during the last check.
	Latency float32 `json:"latency"`

	// PingLoss - no documentation
	PingLoss float32 `json:"pingLoss"`

	// StatusCode - no documentation
	StatusCode string `json:"statusCode"`

	// TargetType - no documentation
	TargetType string `json:"targetType"`

	// Fails - Number of failures since the target was last detected to be working properly.
	Fails int `json:"fails"`
}

func (softlayer_container_auxiliary_network_status_reading *SoftLayer_Container_Auxiliary_Network_Status_Reading) String() string {
	return "SoftLayer_Container_Auxiliary_Network_Status_Reading"
}

package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Auxiliary_Network_Status_Reading - The
// SoftLayer_Container_Auxiliary_Network_Status_Reading data type contains information relating to an
// object being monitored from outside the SoftLayer network. It is primarily used to check the status
// of our edge routers from multiple locations around the world.
type SoftLayer_Container_Auxiliary_Network_Status_Reading struct {

	// AveragePing - no documentation
	AveragePing float32 `json:"averagePing"`

	// Fails - Number of failures since the target was last detected to be working properly.
	Fails int `json:"fails"`

	// Frequency - no documentation
	Frequency int `json:"frequency"`

	// Label - no documentation
	Label string `json:"label"`

	// LastCheckDate - no documentation
	LastCheckDate *time.Time `json:"lastCheckDate"`

	// LastDownDate - no documentation
	LastDownDate *time.Time `json:"lastDownDate"`

	// Latency - The total response time in seconds calculated during the last check.
	Latency float32 `json:"latency"`

	// Location - no documentation
	Location string `json:"location"`

	// MaximumPing - no documentation
	MaximumPing float32 `json:"maximumPing"`

	// MinimumPing - no documentation
	MinimumPing float32 `json:"minimumPing"`

	// PingLoss - no documentation
	PingLoss float32 `json:"pingLoss"`

	// StartDate - no documentation
	StartDate *time.Time `json:"startDate"`

	// StatusCode - no documentation
	StatusCode string `json:"statusCode"`

	// StatusMessage - no documentation
	StatusMessage string `json:"statusMessage"`

	// Target - no documentation
	Target string `json:"target"`

	// TargetType - no documentation
	TargetType string `json:"targetType"`
}

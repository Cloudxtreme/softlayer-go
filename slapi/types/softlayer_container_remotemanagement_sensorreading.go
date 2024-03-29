package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_RemoteManagement_SensorReading - The
// SoftLayer_Container_RemoteManagement_SensorReadings contains sensor information retrieved from a
// server's remote management card.
type SoftLayer_Container_RemoteManagement_SensorReading struct {

	// SensorUnits - no documentation
	SensorUnits string `json:"sensorUnits,omitempty"`

	// UpperCritical - no documentation
	UpperCritical string `json:"upperCritical,omitempty"`

	// UpperNonCritical - no documentation
	UpperNonCritical string `json:"upperNonCritical,omitempty"`

	// UpperNonRecoverable - no documentation
	UpperNonRecoverable string `json:"upperNonRecoverable,omitempty"`

	// LowerCritical - no documentation
	LowerCritical string `json:"lowerCritical,omitempty"`

	// LowerNonCritical - no documentation
	LowerNonCritical string `json:"lowerNonCritical,omitempty"`

	// SensorReading - no documentation
	SensorReading string `json:"sensorReading,omitempty"`

	// Status - no documentation
	Status string `json:"status,omitempty"`

	// LowerNonRecoverable - no documentation
	LowerNonRecoverable string `json:"lowerNonRecoverable,omitempty"`

	// SensorId - no documentation
	SensorId string `json:"sensorId,omitempty"`
}

func (softlayer_container_remotemanagement_sensorreading *SoftLayer_Container_RemoteManagement_SensorReading) String() string {
	return "SoftLayer_Container_RemoteManagement_SensorReading"
}

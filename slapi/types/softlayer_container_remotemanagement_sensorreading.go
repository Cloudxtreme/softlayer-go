package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_RemoteManagement_SensorReading - The
// SoftLayer_Container_RemoteManagement_SensorReadings contains sensor information retrieved from a
// server's remote management card.
type SoftLayer_Container_RemoteManagement_SensorReading struct {

	// LowerCritical - no documentation
	LowerCritical string `json:"lowerCritical"`

	// LowerNonCritical - no documentation
	LowerNonCritical string `json:"lowerNonCritical"`

	// LowerNonRecoverable - no documentation
	LowerNonRecoverable string `json:"lowerNonRecoverable"`

	// SensorId - no documentation
	SensorId string `json:"sensorId"`

	// SensorReading - no documentation
	SensorReading string `json:"sensorReading"`

	// SensorUnits - no documentation
	SensorUnits string `json:"sensorUnits"`

	// Status - no documentation
	Status string `json:"status"`

	// UpperCritical - no documentation
	UpperCritical string `json:"upperCritical"`

	// UpperNonCritical - no documentation
	UpperNonCritical string `json:"upperNonCritical"`

	// UpperNonRecoverable - no documentation
	UpperNonRecoverable string `json:"upperNonRecoverable"`
}

func (softlayer_container_remotemanagement_sensorreading *SoftLayer_Container_RemoteManagement_SensorReading) String() string {
	return "SoftLayer_Container_RemoteManagement_SensorReading"
}

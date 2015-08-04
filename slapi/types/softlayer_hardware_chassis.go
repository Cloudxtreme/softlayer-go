package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Chassis - Every piece of hardware in SoftLayer's datacenters, including customer
// servers, are housed in one of many hardware chassis. The SoftLayer_Hardware_Chassis data type
// defines these chassis.
type SoftLayer_Hardware_Chassis struct {

	// DriveCapacity - The number of hard drives that a hardware chassis can hold.
	DriveCapacity int `json:"driveCapacity,omitempty"`

	// FormFactorId - no documentation
	FormFactorId int `json:"formFactorId,omitempty"`

	// PowerCapacity - The number of power supplies that a hardware chassis can hold.
	PowerCapacity int `json:"powerCapacity,omitempty"`

	// Version - no documentation
	Version string `json:"version,omitempty"`

	// GpuCapacity - The number of GPUs that a hardware chassis can hold.
	GpuCapacity int `json:"gpuCapacity,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Manufacturer - no documentation
	Manufacturer string `json:"manufacturer,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// UnitSize - The physical size of a hardware chassis. Currently this relates to the 'U' size of a
	// chassis buy default.
	UnitSize int `json:"unitSize,omitempty"`

	// HardwareFunction - no documentation
	HardwareFunction *SoftLayer_Hardware_Function `json:"hardwareFunction,omitempty"`
}

func (softlayer_hardware_chassis *SoftLayer_Hardware_Chassis) String() string {
	return "SoftLayer_Hardware_Chassis"
}

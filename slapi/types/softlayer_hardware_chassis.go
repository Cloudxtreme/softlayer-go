package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Chassis - Every piece of hardware in SoftLayer's datacenters, including customer
// servers, are housed in one of many hardware chassis. The SoftLayer_Hardware_Chassis data type
// defines these chassis.
type SoftLayer_Hardware_Chassis struct {

	// UnitSize - The physical size of a hardware chassis. Currently this relates to the 'U' size of a
	// chassis buy default.
	UnitSize int `json:"unitSize"`

	// Version - no documentation
	Version string `json:"version"`

	// FormFactorId - no documentation
	FormFactorId int `json:"formFactorId"`

	// GpuCapacity - The number of GPUs that a hardware chassis can hold.
	GpuCapacity int `json:"gpuCapacity"`

	// Manufacturer - no documentation
	Manufacturer string `json:"manufacturer"`

	// Name - no documentation
	Name string `json:"name"`

	// DriveCapacity - The number of hard drives that a hardware chassis can hold.
	DriveCapacity int `json:"driveCapacity"`

	// Id - no documentation
	Id int `json:"id"`

	// PowerCapacity - The number of power supplies that a hardware chassis can hold.
	PowerCapacity int `json:"powerCapacity"`
}

// SoftLayer_Hardware_Chassis_Extended is SoftLayer_Hardware_Chassis with all maskable types.
type SoftLayer_Hardware_Chassis_Extended struct {
	SoftLayer_Hardware_Chassis

	// HardwareFunction - no documentation
	HardwareFunction *SoftLayer_Hardware_Function `json:"hardwareFunction"`
}

func (softlayer_hardware_chassis *SoftLayer_Hardware_Chassis) String() string {
	return "SoftLayer_Hardware_Chassis"
}

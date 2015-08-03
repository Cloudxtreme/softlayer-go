package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Chassis - Every piece of hardware in SoftLayer's datacenters, including customer
// servers, are housed in one of many hardware chassis. The SoftLayer_Hardware_Chassis data type
// defines these chassis.
type SoftLayer_Hardware_Chassis struct {

	// DriveCapacity - The number of hard drives that a hardware chassis can hold.
	DriveCapacity int `json:"driveCapacity"`

	// GpuCapacity - The number of GPUs that a hardware chassis can hold.
	GpuCapacity int `json:"gpuCapacity"`

	// Id - no documentation
	Id int `json:"id"`

	// Manufacturer - no documentation
	Manufacturer string `json:"manufacturer"`

	// Version - no documentation
	Version string `json:"version"`

	// FormFactorId - no documentation
	FormFactorId int `json:"formFactorId"`

	// Name - no documentation
	Name string `json:"name"`

	// PowerCapacity - The number of power supplies that a hardware chassis can hold.
	PowerCapacity int `json:"powerCapacity"`

	// UnitSize - The physical size of a hardware chassis. Currently this relates to the 'U' size of a
	// chassis buy default.
	UnitSize int `json:"unitSize"`
}

func (softlayer_hardware_chassis *SoftLayer_Hardware_Chassis) String() string {
	return "SoftLayer_Hardware_Chassis"
}

// SoftLayer_Hardware_Chassis_Extended is SoftLayer_Hardware_Chassis with all maskable types.
type SoftLayer_Hardware_Chassis_Extended struct {
	SoftLayer_Hardware_Chassis

	// HardwareFunction - no documentation
	HardwareFunction *SoftLayer_Hardware_Function `json:"hardwareFunction"`
}

func (softlayer_hardware_chassis *SoftLayer_Hardware_Chassis_Extended) String() string {
	return "SoftLayer_Hardware_Chassis"
}

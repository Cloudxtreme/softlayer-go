package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_HardDrive - The SoftLayer_Hardware_Component_HardDrive data type
// abstracts information related to a hard drive.
type SoftLayer_Hardware_Component_HardDrive struct {
}

func (softlayer_hardware_component_harddrive *SoftLayer_Hardware_Component_HardDrive) String() string {
	return "SoftLayer_Hardware_Component_HardDrive"
}

// SoftLayer_Hardware_Component_HardDrive_Extended is SoftLayer_Hardware_Component_HardDrive with all maskable types.
type SoftLayer_Hardware_Component_HardDrive_Extended struct {
	SoftLayer_Hardware_Component_HardDrive

	// PartitionCount - no documentation
	PartitionCount uint64 `json:"partitionCount"`

	// Partitions - no documentation
	Partitions []*SoftLayer_Hardware_Component_Partition `json:"partitions"`
}

func (softlayer_hardware_component_harddrive *SoftLayer_Hardware_Component_HardDrive_Extended) String() string {
	return "SoftLayer_Hardware_Component_HardDrive"
}

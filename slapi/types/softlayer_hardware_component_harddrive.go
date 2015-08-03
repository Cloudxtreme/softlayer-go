package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_HardDrive - The SoftLayer_Hardware_Component_HardDrive data type
// abstracts information related to a hard drive.
type SoftLayer_Hardware_Component_HardDrive struct {

	// PartitionCount - no documentation
	PartitionCount uint64 `json:"partitionCount"`

	// Partitions - no documentation
	Partitions []*SoftLayer_Hardware_Component_Partition `json:"partitions"`
}

func (softlayer_hardware_component_harddrive *SoftLayer_Hardware_Component_HardDrive) String() string {
	return "SoftLayer_Hardware_Component_HardDrive"
}

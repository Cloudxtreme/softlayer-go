package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_Partition - The SoftLayer_Hardware_Component_Partition data type
// contains general information relating to a single hard drive partition.
type SoftLayer_Hardware_Component_Partition struct {

	// HardwareComponentId - A hardware component partition's associated
	// [[SoftLayer_Hardware_Component|hardware component]] Id.
	HardwareComponentId int `json:"hardwareComponentId"`

	// MinimumSize - no documentation
	MinimumSize float64 `json:"minimumSize"`

	// Name - A hardware component partition's name. On a server with windows this may be 'C' and on Linux
	// this may be '/var'
	Name string `json:"name"`

	// DiskNumber - A hardware component partition's order in the [[SoftLayer_Hardware_Hardware|hardware]].
	DiskNumber int `json:"diskNumber"`

	// Grow - A flag indicating if a partition is the grow partition. The grow partition will grow to fill
	// all remaining space on a disk. There can only be one.
	Grow int `json:"grow"`
}

// SoftLayer_Hardware_Component_Partition_Extended is SoftLayer_Hardware_Component_Partition with all maskable types.
type SoftLayer_Hardware_Component_Partition_Extended struct {
	SoftLayer_Hardware_Component_Partition

	// HardwareComponent - A hardware component partitions's associated
	// [[SoftLayer_Hardware_Component|Hardware Component]]. Likely to be a
	// [[SoftLayer_Hardware_Component_HardDrive|Hard Drive]]
	HardwareComponent *SoftLayer_Hardware_Component `json:"hardwareComponent"`
}

func (softlayer_hardware_component_partition *SoftLayer_Hardware_Component_Partition) String() string {
	return "SoftLayer_Hardware_Component_Partition"
}

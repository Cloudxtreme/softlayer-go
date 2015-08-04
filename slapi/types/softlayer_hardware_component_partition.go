package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Hardware_Component_Partition - The SoftLayer_Hardware_Component_Partition data type
// contains general information relating to a single hard drive partition.
type SoftLayer_Hardware_Component_Partition struct {

	// DiskNumber - A hardware component partition's order in the [[SoftLayer_Hardware_Hardware|hardware]].
	DiskNumber int `json:"diskNumber,omitempty"`

	// Grow - A flag indicating if a partition is the grow partition. The grow partition will grow to fill
	// all remaining space on a disk. There can only be one.
	Grow int `json:"grow,omitempty"`

	// HardwareComponentId - A hardware component partition's associated
	// [[SoftLayer_Hardware_Component|hardware component]] Id.
	HardwareComponentId int `json:"hardwareComponentId,omitempty"`

	// MinimumSize - no documentation
	MinimumSize slapi.Float64 `json:"minimumSize,omitempty"`

	// Name - A hardware component partition's name. On a server with windows this may be 'C' and on Linux
	// this may be '/var'
	Name string `json:"name,omitempty"`
}

func (softlayer_hardware_component_partition *SoftLayer_Hardware_Component_Partition) String() string {
	return "SoftLayer_Hardware_Component_Partition"
}

// SoftLayer_Hardware_Component_Partition_Extended is SoftLayer_Hardware_Component_Partition with all maskable types.
type SoftLayer_Hardware_Component_Partition_Extended struct {
	SoftLayer_Hardware_Component_Partition

	// HardwareComponent - A hardware component partitions's associated
	// [[SoftLayer_Hardware_Component|Hardware Component]]. Likely to be a
	// [[SoftLayer_Hardware_Component_HardDrive|Hard Drive]]
	HardwareComponent *SoftLayer_Hardware_Component `json:"hardwareComponent,omitempty"`
}

func (softlayer_hardware_component_partition *SoftLayer_Hardware_Component_Partition_Extended) String() string {
	return "SoftLayer_Hardware_Component_Partition"
}

package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Hardware_Component_Partition_Template_Partition - The
// SoftLayer_Hardware_Component_Partition_Template_Partition data type contains general information
// relating to a single SoftLayer Template Partition.
type SoftLayer_Hardware_Component_Partition_Template_Partition struct {

	// IsGrow - A flag indication if a partition will be the grow partition. The grow partition will have
	// its size adjusted to fill all available space on a hard drive.
	IsGrow bool `json:"isGrow,omitempty"`

	// PartitionName - no documentation
	PartitionName string `json:"partitionName,omitempty"`

	// PartitionSize - no documentation
	PartitionSize slapi.Float64 `json:"partitionSize,omitempty"`

	// PartitionTemplateId - A partition's associated
	// [[SoftLayer_Hardware_Component_Partition_Template|Partition Template]] Id.
	PartitionTemplateId int `json:"partitionTemplateId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// PartitionTemplate - A partition's [[SoftLayer_Hardware_Component_Partition_Template|Partition
	// Template]].
	PartitionTemplate *SoftLayer_Hardware_Component_Partition_Template `json:"partitionTemplate,omitempty"`

	// FilesystemType - no documentation
	FilesystemType *SoftLayer_Configuration_Storage_Filesystem_Type `json:"filesystemType,omitempty"`
}

func (softlayer_hardware_component_partition_template_partition *SoftLayer_Hardware_Component_Partition_Template_Partition) String() string {
	return "SoftLayer_Hardware_Component_Partition_Template_Partition"
}

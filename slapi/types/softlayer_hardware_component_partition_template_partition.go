package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_Partition_Template_Partition - The
// SoftLayer_Hardware_Component_Partition_Template_Partition data type contains general information
// relating to a single SoftLayer Template Partition.
type SoftLayer_Hardware_Component_Partition_Template_Partition struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// IsGrow - A flag indication if a partition will be the grow partition. The grow partition will have
	// its size adjusted to fill all available space on a hard drive.
	IsGrow bool `json:"isGrow,omitempty"`

	// PartitionName - no documentation
	PartitionName string `json:"partitionName,omitempty"`

	// PartitionSize - no documentation
	PartitionSize float64 `json:"partitionSize,omitempty"`

	// PartitionTemplateId - A partition's associated
	// [[SoftLayer_Hardware_Component_Partition_Template|Partition Template]] Id.
	PartitionTemplateId int `json:"partitionTemplateId,omitempty"`
}

func (softlayer_hardware_component_partition_template_partition *SoftLayer_Hardware_Component_Partition_Template_Partition) String() string {
	return "SoftLayer_Hardware_Component_Partition_Template_Partition"
}

// SoftLayer_Hardware_Component_Partition_Template_Partition_Extended is SoftLayer_Hardware_Component_Partition_Template_Partition with all maskable types.
type SoftLayer_Hardware_Component_Partition_Template_Partition_Extended struct {
	SoftLayer_Hardware_Component_Partition_Template_Partition

	// FilesystemType - no documentation
	FilesystemType *SoftLayer_Configuration_Storage_Filesystem_Type `json:"filesystemType,omitempty"`

	// PartitionTemplate - A partition's [[SoftLayer_Hardware_Component_Partition_Template|Partition
	// Template]].
	PartitionTemplate *SoftLayer_Hardware_Component_Partition_Template `json:"partitionTemplate,omitempty"`
}

func (softlayer_hardware_component_partition_template_partition *SoftLayer_Hardware_Component_Partition_Template_Partition_Extended) String() string {
	return "SoftLayer_Hardware_Component_Partition_Template_Partition"
}

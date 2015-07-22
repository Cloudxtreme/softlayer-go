package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_Partition_Template_Partition - The
// SoftLayer_Hardware_Component_Partition_Template_Partition data type contains general information
// relating to a single SoftLayer Template Partition.
type SoftLayer_Hardware_Component_Partition_Template_Partition struct {

	// FilesystemType - no documentation
	FilesystemType *SoftLayer_Configuration_Storage_Filesystem_Type `json:"filesystemType"`

	// Id - no documentation
	Id int `json:"id"`

	// IsGrow - A flag indication if a partition will be the grow partition. The grow partition will have
	// its size adjusted to fill all available space on a hard drive.
	IsGrow bool `json:"isGrow"`

	// PartitionName - no documentation
	PartitionName string `json:"partitionName"`

	// PartitionSize - no documentation
	PartitionSize float64 `json:"partitionSize"`

	// PartitionTemplate - A partition's [[SoftLayer_Hardware_Component_Partition_Template|Partition
	// Template]].
	PartitionTemplate *SoftLayer_Hardware_Component_Partition_Template `json:"partitionTemplate"`

	// PartitionTemplateId - A partition's associated
	// [[SoftLayer_Hardware_Component_Partition_Template|Partition Template]] Id.
	PartitionTemplateId int `json:"partitionTemplateId"`
}

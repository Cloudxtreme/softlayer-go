package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_Partition_OperatingSystem - The
// SoftLayer_Hardware_Component_Partition_OperatingSystem data type contains general information
// relating to a single SoftLayer Operating System Partition Template.
type SoftLayer_Hardware_Component_Partition_OperatingSystem struct {

	// Description - A partition template operating system's description. Typically the title of the
	// Operating System.
	Description string `json:"description,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Notes - Information about the kinds of partition templates assigned to this operating system.
	Notes string `json:"notes,omitempty"`
}

func (softlayer_hardware_component_partition_operatingsystem *SoftLayer_Hardware_Component_Partition_OperatingSystem) String() string {
	return "SoftLayer_Hardware_Component_Partition_OperatingSystem"
}

// SoftLayer_Hardware_Component_Partition_OperatingSystem_Extended is SoftLayer_Hardware_Component_Partition_OperatingSystem with all maskable types.
type SoftLayer_Hardware_Component_Partition_OperatingSystem_Extended struct {
	SoftLayer_Hardware_Component_Partition_OperatingSystem

	// PartitionTemplateCount - A count of information regarding an operating system's
	// [[SoftLayer_Hardware_Component_Partition_Template|Partition Templates]].
	PartitionTemplateCount uint64 `json:"partitionTemplateCount,omitempty"`

	// PartitionTemplates - Information regarding an operating system's
	// [[SoftLayer_Hardware_Component_Partition_Template|Partition Templates]].
	PartitionTemplates []*SoftLayer_Hardware_Component_Partition_Template `json:"partitionTemplates,omitempty"`
}

func (softlayer_hardware_component_partition_operatingsystem *SoftLayer_Hardware_Component_Partition_OperatingSystem_Extended) String() string {
	return "SoftLayer_Hardware_Component_Partition_OperatingSystem"
}

package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_Partition_OperatingSystem - The
// SoftLayer_Hardware_Component_Partition_OperatingSystem data type contains general information
// relating to a single SoftLayer Operating System Partition Template.
type SoftLayer_Hardware_Component_Partition_OperatingSystem struct {

	// Description - A partition template operating system's description. Typically the title of the
	// Operating System.
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`

	// Notes - Information about the kinds of partition templates assigned to this operating system.
	Notes string `json:"notes"`

	// PartitionTemplateCount - A count of information regarding an operating system's
	// [[SoftLayer_Hardware_Component_Partition_Template|Partition Templates]].
	PartitionTemplateCount uint64 `json:"partitionTemplateCount"`

	// PartitionTemplates - Information regarding an operating system's
	// [[SoftLayer_Hardware_Component_Partition_Template|Partition Templates]].
	PartitionTemplates []*SoftLayer_Hardware_Component_Partition_Template `json:"partitionTemplates"`
}

// GetAllObjects - <nil>
func (softlayer_hardware_component_partition_operatingsystem *SoftLayer_Hardware_Component_Partition_OperatingSystem) GetAllObjects() ([]*SoftLayer_Hardware_Component_Partition_OperatingSystem, error) {
	var returnValue []*SoftLayer_Hardware_Component_Partition_OperatingSystem
	return returnValue, nil
}

// GetByDescription - The '''getByDescription''' method retrieves all possible partition templates
// based on the description (required parameter) entered when calling the method. The description is
// typically the operating system's name. Current recognized values include 'linux', 'windows',
// 'freebsd', and 'Debian'.
func (softlayer_hardware_component_partition_operatingsystem *SoftLayer_Hardware_Component_Partition_OperatingSystem) GetByDescription(description string) (*SoftLayer_Hardware_Component_Partition_OperatingSystem, error) {
	var returnValue *SoftLayer_Hardware_Component_Partition_OperatingSystem
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Hardware_Component_Partition_OperatingSystem object
// whose ID number corresponds to the ID number of the init parameter passed to the
// SoftLayer_Hardware_Component_Partition_OperatingSystem service.s
func (softlayer_hardware_component_partition_operatingsystem *SoftLayer_Hardware_Component_Partition_OperatingSystem) GetObject() (*SoftLayer_Hardware_Component_Partition_OperatingSystem, error) {
	var returnValue *SoftLayer_Hardware_Component_Partition_OperatingSystem
	return returnValue, nil
}

package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Hardware_Component_HardDrive - The SoftLayer_Hardware_Component_HardDrive data type
// abstracts information related to a hard drive.
type SoftLayer_Hardware_Component_HardDrive struct {

	// Name - The name of this component as referenced by the operating system.
	Name string `json:"name,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// HardwareId - The internal identifier of the hardware that a hardware component resides inside.
	HardwareId int `json:"hardwareId,omitempty"`

	// SerialNumber - no documentation
	SerialNumber string `json:"serialNumber,omitempty"`

	// HardwareComponentModelId - The internal identifier of a hardware component's component model.
	HardwareComponentModelId int `json:"hardwareComponentModelId,omitempty"`

	// ServiceProviderId - A hardware's internal identification number at its service provider
	ServiceProviderId int `json:"serviceProviderId,omitempty"`

	// ModifyDate - The date that a hardware component was last modified.
	ModifyDate *time.Time `json:"modifyDate,omitempty"`
}

func (softlayer_hardware_component_harddrive *SoftLayer_Hardware_Component_HardDrive) String() string {
	return "SoftLayer_Hardware_Component_HardDrive"
}

// SoftLayer_Hardware_Component_HardDrive_Extended is SoftLayer_Hardware_Component_HardDrive with all maskable types.
type SoftLayer_Hardware_Component_HardDrive_Extended struct {
	SoftLayer_Hardware_Component_HardDrive

	// ChildrenCount - A count of a components sub components. Devices that are usually integrated or in
	// some way attached to a component.
	ChildrenCount uint64 `json:"childrenCount,omitempty"`

	// Capacity - no documentation
	Capacity float64 `json:"capacity,omitempty"`

	// HardwareComponentType - no documentation
	HardwareComponentType *SoftLayer_Hardware_Component_Type `json:"hardwareComponentType,omitempty"`

	// Parent - A components parent. Devices that are usually integrated or in some way attached to a
	// component.
	Parent *SoftLayer_Hardware_Component `json:"parent,omitempty"`

	// DownlinkHardwareComponents - <nil>
	DownlinkHardwareComponents []*SoftLayer_Hardware_Component `json:"downlinkHardwareComponents,omitempty"`

	// HardwareComponentModel - no documentation
	HardwareComponentModel *SoftLayer_Hardware_Component_Model `json:"hardwareComponentModel,omitempty"`

	// ServiceProvider - <nil>
	ServiceProvider *SoftLayer_Service_Provider `json:"serviceProvider,omitempty"`

	// RaidMode - no documentation
	RaidMode string `json:"raidMode,omitempty"`

	// Owner - no documentation
	Owner *SoftLayer_Account `json:"owner,omitempty"`

	// UplinkHardwareComponentCount - no documentation
	UplinkHardwareComponentCount uint64 `json:"uplinkHardwareComponentCount,omitempty"`

	// Children - A components sub components. Devices that are usually integrated or in some way attached
	// to a component.
	Children []*SoftLayer_Hardware_Component `json:"children,omitempty"`

	// Hardware - The hardware object that this component belongs to.
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`

	// NetworkComponents - The components local ethernet and remote management interfaces
	NetworkComponents []*SoftLayer_Network_Component `json:"networkComponents,omitempty"`

	// PartitionCount - no documentation
	PartitionCount uint64 `json:"partitionCount,omitempty"`

	// UplinkHardwareComponents - <nil>
	UplinkHardwareComponents []*SoftLayer_Hardware_Component `json:"uplinkHardwareComponents,omitempty"`

	// DownlinkHardwareComponentCount - no documentation
	DownlinkHardwareComponentCount uint64 `json:"downlinkHardwareComponentCount,omitempty"`

	// NetworkComponentCount - A count of the components local ethernet and remote management interfaces
	NetworkComponentCount uint64 `json:"networkComponentCount,omitempty"`

	// Partitions - no documentation
	Partitions []*SoftLayer_Hardware_Component_Partition `json:"partitions,omitempty"`
}

func (softlayer_hardware_component_harddrive *SoftLayer_Hardware_Component_HardDrive_Extended) String() string {
	return "SoftLayer_Hardware_Component_HardDrive"
}

package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Hardware_Component_RemoteManagement - This class adds functionality to the base
// SoftLayer_Hardware class for web servers (all server hardware)
type SoftLayer_Hardware_Component_RemoteManagement struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// ModifyDate - The date that a hardware component was last modified.
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Name - The name of this component as referenced by the operating system.
	Name string `json:"name,omitempty"`

	// HardwareId - The internal identifier of the hardware that a hardware component resides inside.
	HardwareId int `json:"hardwareId,omitempty"`

	// ServiceProviderId - A hardware's internal identification number at its service provider
	ServiceProviderId int `json:"serviceProviderId,omitempty"`

	// SerialNumber - no documentation
	SerialNumber string `json:"serialNumber,omitempty"`

	// HardwareComponentModelId - The internal identifier of a hardware component's component model.
	HardwareComponentModelId int `json:"hardwareComponentModelId,omitempty"`
}

func (softlayer_hardware_component_remotemanagement *SoftLayer_Hardware_Component_RemoteManagement) String() string {
	return "SoftLayer_Hardware_Component_RemoteManagement"
}

// SoftLayer_Hardware_Component_RemoteManagement_Extended is SoftLayer_Hardware_Component_RemoteManagement with all maskable types.
type SoftLayer_Hardware_Component_RemoteManagement_Extended struct {
	SoftLayer_Hardware_Component_RemoteManagement

	// UplinkHardwareComponentCount - no documentation
	UplinkHardwareComponentCount uint64 `json:"uplinkHardwareComponentCount,omitempty"`

	// NetworkComponentCount - A count of the components local ethernet and remote management interfaces
	NetworkComponentCount uint64 `json:"networkComponentCount,omitempty"`

	// HardwareComponentType - no documentation
	HardwareComponentType *SoftLayer_Hardware_Component_Type `json:"hardwareComponentType,omitempty"`

	// RaidMode - no documentation
	RaidMode string `json:"raidMode,omitempty"`

	// NetworkComponent - no documentation
	NetworkComponent *SoftLayer_Network_Component `json:"networkComponent,omitempty"`

	// Hardware - The hardware object that this component belongs to.
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`

	// Parent - A components parent. Devices that are usually integrated or in some way attached to a
	// component.
	Parent *SoftLayer_Hardware_Component `json:"parent,omitempty"`

	// DownlinkHardwareComponentCount - no documentation
	DownlinkHardwareComponentCount uint64 `json:"downlinkHardwareComponentCount,omitempty"`

	// NetworkComponents - The components local ethernet and remote management interfaces
	NetworkComponents []*SoftLayer_Network_Component `json:"networkComponents,omitempty"`

	// ChildrenCount - A count of a components sub components. Devices that are usually integrated or in
	// some way attached to a component.
	ChildrenCount uint64 `json:"childrenCount,omitempty"`

	// Owner - no documentation
	Owner *SoftLayer_Account `json:"owner,omitempty"`

	// Children - A components sub components. Devices that are usually integrated or in some way attached
	// to a component.
	Children []*SoftLayer_Hardware_Component `json:"children,omitempty"`

	// ServiceProvider - <nil>
	ServiceProvider *SoftLayer_Service_Provider `json:"serviceProvider,omitempty"`

	// DownlinkHardwareComponents - <nil>
	DownlinkHardwareComponents []*SoftLayer_Hardware_Component `json:"downlinkHardwareComponents,omitempty"`

	// HardwareComponentModel - no documentation
	HardwareComponentModel *SoftLayer_Hardware_Component_Model `json:"hardwareComponentModel,omitempty"`

	// UplinkHardwareComponents - <nil>
	UplinkHardwareComponents []*SoftLayer_Hardware_Component `json:"uplinkHardwareComponents,omitempty"`

	// Capacity - no documentation
	Capacity float64 `json:"capacity,omitempty"`
}

func (softlayer_hardware_component_remotemanagement *SoftLayer_Hardware_Component_RemoteManagement_Extended) String() string {
	return "SoftLayer_Hardware_Component_RemoteManagement"
}

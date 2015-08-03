package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Hardware_Component - The SoftLayer_Hardware_Component data type abstracts information
// related to a hardware component.
type SoftLayer_Hardware_Component struct {

	// HardwareId - The internal identifier of the hardware that a hardware component resides inside.
	HardwareId int `json:"hardwareId"`

	// ModifyDate - The date that a hardware component was last modified.
	ModifyDate *time.Time `json:"modifyDate"`

	// ServiceProviderId - A hardware's internal identification number at its service provider
	ServiceProviderId int `json:"serviceProviderId"`

	// HardwareComponentModelId - The internal identifier of a hardware component's component model.
	HardwareComponentModelId int `json:"hardwareComponentModelId"`

	// Id - no documentation
	Id int `json:"id"`

	// Name - The name of this component as referenced by the operating system.
	Name string `json:"name"`

	// SerialNumber - no documentation
	SerialNumber string `json:"serialNumber"`
}

// SoftLayer_Hardware_Component_Extended is SoftLayer_Hardware_Component with all maskable types.
type SoftLayer_Hardware_Component_Extended struct {
	SoftLayer_Hardware_Component

	// Children - A components sub components. Devices that are usually integrated or in some way attached
	// to a component.
	Children []*SoftLayer_Hardware_Component `json:"children"`

	// NetworkComponentCount - A count of the components local ethernet and remote management interfaces
	NetworkComponentCount uint64 `json:"networkComponentCount"`

	// Capacity - no documentation
	Capacity float64 `json:"capacity"`

	// NetworkComponents - The components local ethernet and remote management interfaces
	NetworkComponents []*SoftLayer_Network_Component `json:"networkComponents"`

	// ServiceProvider - <nil>
	ServiceProvider *SoftLayer_Service_Provider `json:"serviceProvider"`

	// UplinkHardwareComponents - <nil>
	UplinkHardwareComponents []*SoftLayer_Hardware_Component `json:"uplinkHardwareComponents"`

	// DownlinkHardwareComponentCount - no documentation
	DownlinkHardwareComponentCount uint64 `json:"downlinkHardwareComponentCount"`

	// UplinkHardwareComponentCount - no documentation
	UplinkHardwareComponentCount uint64 `json:"uplinkHardwareComponentCount"`

	// DownlinkHardwareComponents - <nil>
	DownlinkHardwareComponents []*SoftLayer_Hardware_Component `json:"downlinkHardwareComponents"`

	// HardwareComponentModel - no documentation
	HardwareComponentModel *SoftLayer_Hardware_Component_Model `json:"hardwareComponentModel"`

	// Owner - no documentation
	Owner *SoftLayer_Account `json:"owner"`

	// RaidMode - no documentation
	RaidMode string `json:"raidMode"`

	// Hardware - The hardware object that this component belongs to.
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// HardwareComponentType - no documentation
	HardwareComponentType *SoftLayer_Hardware_Component_Type `json:"hardwareComponentType"`

	// Parent - A components parent. Devices that are usually integrated or in some way attached to a
	// component.
	Parent *SoftLayer_Hardware_Component `json:"parent"`

	// ChildrenCount - A count of a components sub components. Devices that are usually integrated or in
	// some way attached to a component.
	ChildrenCount uint64 `json:"childrenCount"`
}

func (softlayer_hardware_component *SoftLayer_Hardware_Component) String() string {
	return "SoftLayer_Hardware_Component"
}

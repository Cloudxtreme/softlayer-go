package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Configuration_Storage_Group_Array_Type - no documentation
type SoftLayer_Configuration_Storage_Group_Array_Type struct {

	// DriveMultiplier - <nil>
	DriveMultiplier int `json:"driveMultiplier,omitempty"`

	// HotspareAllow - <nil>
	HotspareAllow bool `json:"hotspareAllow,omitempty"`

	// KeyName - <nil>
	KeyName string `json:"keyName,omitempty"`

	// Name - <nil>
	Name string `json:"name,omitempty"`

	// Description - <nil>
	Description string `json:"description,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// MaximumDrives - <nil>
	MaximumDrives int `json:"maximumDrives,omitempty"`

	// MinimumDrives - <nil>
	MinimumDrives int `json:"minimumDrives,omitempty"`

	// HardwareComponentModels - <nil>
	HardwareComponentModels []*SoftLayer_Hardware_Component_Model `json:"hardwareComponentModels,omitempty"`

	// HardwareComponentModelCount - no documentation
	HardwareComponentModelCount uint64 `json:"hardwareComponentModelCount,omitempty"`
}

func (softlayer_configuration_storage_group_array_type *SoftLayer_Configuration_Storage_Group_Array_Type) String() string {
	return "SoftLayer_Configuration_Storage_Group_Array_Type"
}

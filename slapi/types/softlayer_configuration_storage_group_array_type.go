package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Configuration_Storage_Group_Array_Type - no documentation
type SoftLayer_Configuration_Storage_Group_Array_Type struct {

	// MinimumDrives - <nil>
	MinimumDrives int `json:"minimumDrives"`

	// DriveMultiplier - <nil>
	DriveMultiplier int `json:"driveMultiplier"`

	// Id - <nil>
	Id int `json:"id"`

	// MaximumDrives - <nil>
	MaximumDrives int `json:"maximumDrives"`

	// KeyName - <nil>
	KeyName string `json:"keyName"`

	// Name - <nil>
	Name string `json:"name"`

	// Description - <nil>
	Description string `json:"description"`

	// HotspareAllow - <nil>
	HotspareAllow bool `json:"hotspareAllow"`
}

// SoftLayer_Configuration_Storage_Group_Array_Type_Extended is SoftLayer_Configuration_Storage_Group_Array_Type with all maskable types.
type SoftLayer_Configuration_Storage_Group_Array_Type_Extended struct {
	SoftLayer_Configuration_Storage_Group_Array_Type

	// HardwareComponentModelCount - no documentation
	HardwareComponentModelCount uint64 `json:"hardwareComponentModelCount"`

	// HardwareComponentModels - <nil>
	HardwareComponentModels []*SoftLayer_Hardware_Component_Model `json:"hardwareComponentModels"`
}

func (softlayer_configuration_storage_group_array_type *SoftLayer_Configuration_Storage_Group_Array_Type) String() string {
	return "SoftLayer_Configuration_Storage_Group_Array_Type"
}

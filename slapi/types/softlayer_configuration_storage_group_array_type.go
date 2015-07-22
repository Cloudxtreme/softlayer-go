package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Configuration_Storage_Group_Array_Type - no documentation
type SoftLayer_Configuration_Storage_Group_Array_Type struct {

	// Description - <nil>
	Description string `json:"description"`

	// DriveMultiplier - <nil>
	DriveMultiplier int `json:"driveMultiplier"`

	// HardwareComponentModelCount - no documentation
	HardwareComponentModelCount uint64 `json:"hardwareComponentModelCount"`

	// HardwareComponentModels - <nil>
	HardwareComponentModels []*SoftLayer_Hardware_Component_Model `json:"hardwareComponentModels"`

	// HotspareAllow - <nil>
	HotspareAllow bool `json:"hotspareAllow"`

	// Id - <nil>
	Id int `json:"id"`

	// KeyName - <nil>
	KeyName string `json:"keyName"`

	// MaximumDrives - <nil>
	MaximumDrives int `json:"maximumDrives"`

	// MinimumDrives - <nil>
	MinimumDrives int `json:"minimumDrives"`

	// Name - <nil>
	Name string `json:"name"`
}

// GetAllObjects - <nil>
func (softlayer_configuration_storage_group_array_type *SoftLayer_Configuration_Storage_Group_Array_Type) GetAllObjects(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Configuration_Storage_Group_Array_Type, error) {
	var returnValue []*SoftLayer_Configuration_Storage_Group_Array_Type
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_configuration_storage_group_array_type *SoftLayer_Configuration_Storage_Group_Array_Type) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Configuration_Storage_Group_Array_Type, error) {
	var returnValue *SoftLayer_Configuration_Storage_Group_Array_Type
	return returnValue, nil
}
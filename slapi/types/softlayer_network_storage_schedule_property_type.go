package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_Storage_Schedule_Property_Type - A schedule property type is used to allow for a
// standardized method of defining network storage schedules.
type SoftLayer_Network_Storage_Schedule_Property_Type struct {

	// Description - A type's description, for example 'Date for the schedule to start.'.
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`

	// Keyname - no documentation
	Keyname string `json:"keyname"`

	// Name - A schedule property type's name, for example 'Start Date'.
	Name string `json:"name"`

	// NasType - The type of Storage volume type which a property type may be associated with.
	NasType string `json:"nasType"`
}

// GetAllObjects - Use this method to retrieve all network storage schedule property types.
func (softlayer_network_storage_schedule_property_type *SoftLayer_Network_Storage_Schedule_Property_Type) GetAllObjects(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Network_Storage_Schedule_Property_Type, error) {
	var returnValue []*SoftLayer_Network_Storage_Schedule_Property_Type
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_storage_schedule_property_type *SoftLayer_Network_Storage_Schedule_Property_Type) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Network_Storage_Schedule_Property_Type, error) {
	var returnValue *SoftLayer_Network_Storage_Schedule_Property_Type
	return returnValue, nil
}
package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_Storage_Group_Type - <nil>
type SoftLayer_Network_Storage_Group_Type struct {

	// Id - <nil>
	Id int `json:"id"`

	// KeyName - <nil>
	KeyName string `json:"keyName"`

	// Name - <nil>
	Name string `json:"name"`
}

// GetAllObjects - Use this method to retrieve all storage group types available.
func (softlayer_network_storage_group_type *SoftLayer_Network_Storage_Group_Type) GetAllObjects(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Network_Storage_Group_Type, error) {
	var returnValue []*SoftLayer_Network_Storage_Group_Type
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_storage_group_type *SoftLayer_Network_Storage_Group_Type) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Network_Storage_Group_Type, error) {
	var returnValue *SoftLayer_Network_Storage_Group_Type
	return returnValue, nil
}

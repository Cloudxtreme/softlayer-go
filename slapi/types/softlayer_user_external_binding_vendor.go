package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_User_External_Binding_Vendor - The SoftLayer_User_External_Binding_Vendor data type
// contains information for a single external binding vendor. This information includes a user friendly
// vendor name, a unique version of the vendor name, and a unique internal identifier that can be used
// when creating a new external binding.
type SoftLayer_User_External_Binding_Vendor struct {

	// Id - The unique identifier for an external binding vendor.
	Id int `json:"id"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// Name - The user friendly name of an external binding vendor.
	Name string `json:"name"`
}

// GetAllObjects - getAllObjects() will return a list of the available external binding vendors that
// SoftLayer supports. Use this list to select the appropriate vendor when creating a new external
// binding.
func (softlayer_user_external_binding_vendor *SoftLayer_User_External_Binding_Vendor) GetAllObjects(commonOptions *slapi.CommonOptions) ([]*SoftLayer_User_External_Binding_Vendor, error) {
	var returnValue []*SoftLayer_User_External_Binding_Vendor
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_user_external_binding_vendor *SoftLayer_User_External_Binding_Vendor) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_User_External_Binding_Vendor, error) {
	var returnValue *SoftLayer_User_External_Binding_Vendor
	return returnValue, nil
}

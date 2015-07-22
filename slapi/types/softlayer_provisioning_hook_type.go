package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Provisioning_Hook_Type - <nil>
type SoftLayer_Provisioning_Hook_Type struct {

	// Description - <nil>
	Description string `json:"description"`

	// Id - <nil>
	Id int `json:"id"`

	// KeyName - <nil>
	KeyName string `json:"keyName"`

	// Name - <nil>
	Name string `json:"name"`
}

// GetAllHookTypes - <nil>
func (softlayer_provisioning_hook_type *SoftLayer_Provisioning_Hook_Type) GetAllHookTypes(ctx *slapi.RequestContext) ([]*SoftLayer_Provisioning_Hook_Type, error) {
	var returnValue []*SoftLayer_Provisioning_Hook_Type
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_provisioning_hook_type *SoftLayer_Provisioning_Hook_Type) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Provisioning_Hook_Type, error) {
	var returnValue *SoftLayer_Provisioning_Hook_Type
	return returnValue, nil
}

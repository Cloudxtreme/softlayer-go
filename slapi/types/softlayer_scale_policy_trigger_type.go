package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Scale_Policy_Trigger_Type - <nil>
type SoftLayer_Scale_Policy_Trigger_Type struct {

	// Id - no documentation
	Id int `json:"id"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// Name - no documentation
	Name string `json:"name"`
}

// GetAllObjects - <nil>
func (softlayer_scale_policy_trigger_type *SoftLayer_Scale_Policy_Trigger_Type) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_Scale_Policy_Trigger_Type, error) {
	var returnValue []*SoftLayer_Scale_Policy_Trigger_Type
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_scale_policy_trigger_type *SoftLayer_Scale_Policy_Trigger_Type) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Scale_Policy_Trigger_Type, error) {
	var returnValue *SoftLayer_Scale_Policy_Trigger_Type
	return returnValue, nil
}

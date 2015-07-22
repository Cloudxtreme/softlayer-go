package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Scale_Termination_Policy - <nil>
type SoftLayer_Scale_Termination_Policy struct {

	// Id - no documentation
	Id int `json:"id"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// Name - no documentation
	Name string `json:"name"`
}

// GetAllObjects - <nil>
func (softlayer_scale_termination_policy *SoftLayer_Scale_Termination_Policy) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_Scale_Termination_Policy, error) {
	var returnValue []*SoftLayer_Scale_Termination_Policy
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_scale_termination_policy *SoftLayer_Scale_Termination_Policy) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Scale_Termination_Policy, error) {
	var returnValue *SoftLayer_Scale_Termination_Policy
	return returnValue, nil
}

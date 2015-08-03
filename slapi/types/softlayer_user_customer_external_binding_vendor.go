package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_User_Customer_External_Binding_Vendor - The
// SoftLayer_User_Customer_External_Binding_Vendor data type contains information for a single external
// binding vendor. This information includes a user friendly vendor name, a unique version of the
// vendor name, and a unique internal identifier that can be used when creating a new external binding.
type SoftLayer_User_Customer_External_Binding_Vendor struct {
}

func (softlayer_user_customer_external_binding_vendor *SoftLayer_User_Customer_External_Binding_Vendor) String() string {
	return "SoftLayer_User_Customer_External_Binding_Vendor"
}

// GetObject - <nil>
func (softlayer_user_customer_external_binding_vendor *SoftLayer_User_Customer_External_Binding_Vendor) GetObject(ctx *slapi.RequestContext) (*SoftLayer_User_Customer_External_Binding_Vendor, error) {
	var returnValue *SoftLayer_User_Customer_External_Binding_Vendor
	return returnValue, nil
}

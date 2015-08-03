package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_User_Customer_External_Binding - The SoftLayer_User_Customer_External_Binding data type
// contains general information for a single external binding. This includes the 3rd party vendor, type
// of binding, and a unique identifier and password that is used to authenticate against the 3rd party
// service.
type SoftLayer_User_Customer_External_Binding struct {

	// User - The SoftLayer user that the external authentication binding belongs to.
	User *SoftLayer_User_Customer `json:"user"`
}

func (softlayer_user_customer_external_binding *SoftLayer_User_Customer_External_Binding) String() string {
	return "SoftLayer_User_Customer_External_Binding"
}

// Disable - Disabling an external binding will allow you to keep the external binding on your
// SoftLayer account, but will not require you to authentication with our trusted 2 form factor vendor
// when logging into the SoftLayer customer portal. You may supply one of the following reason when you
// disable an external binding: *Unspecified *TemporarilyUnavailable *Lost *Stolen
func (softlayer_user_customer_external_binding *SoftLayer_User_Customer_External_Binding) Disable(ctx *slapi.RequestContext, reason string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// Enable - Enabling an external binding will activate the binding on your account and require you to
// authenticate with our trusted 3rd party 2 form factor vendor when logging into the SoftLayer
// customer portal. Please note that API access will be disabled for users that have an active external
// binding.
func (softlayer_user_customer_external_binding *SoftLayer_User_Customer_External_Binding) Enable(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_user_customer_external_binding *SoftLayer_User_Customer_External_Binding) GetObject(ctx *slapi.RequestContext) (*SoftLayer_User_Customer_External_Binding, error) {
	var returnValue *SoftLayer_User_Customer_External_Binding
	return returnValue, nil
}

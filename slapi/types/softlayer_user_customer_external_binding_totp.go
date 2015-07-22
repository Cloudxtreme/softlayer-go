package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_User_Customer_External_Binding_Totp - The SoftLayer_User_Customer_External_Binding_Totp
// data type contains information about a single time-based one time password external binding. The
// external binding information is used when a SoftLayer customer logs into the SoftLayer customer
// portal to authenticate them. The information provided by this external binding data type includes: *
// The type of credential * The current state of the credential ** Active ** Inactive SoftLayer users
// with an active external binding will be prohibited from using the API for security reasons.
type SoftLayer_User_Customer_External_Binding_Totp struct {
}

// Activate - <nil>
func (softlayer_user_customer_external_binding_totp *SoftLayer_User_Customer_External_Binding_Totp) Activate(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// Deactivate - <nil>
func (softlayer_user_customer_external_binding_totp *SoftLayer_User_Customer_External_Binding_Totp) Deactivate(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GenerateSecretKey - <nil>
func (softlayer_user_customer_external_binding_totp *SoftLayer_User_Customer_External_Binding_Totp) GenerateSecretKey(commonOptions *slapi.CommonOptions) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_user_customer_external_binding_totp *SoftLayer_User_Customer_External_Binding_Totp) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_User_Customer_External_Binding_Totp, error) {
	var returnValue *SoftLayer_User_Customer_External_Binding_Totp
	return returnValue, nil
}

package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_Customer_External_Binding_Totp - The SoftLayer_User_Customer_External_Binding_Totp
// data type contains information about a single time-based one time password external binding. The
// external binding information is used when a SoftLayer customer logs into the SoftLayer customer
// portal to authenticate them. The information provided by this external binding data type includes: *
// The type of credential * The current state of the credential ** Active ** Inactive SoftLayer users
// with an active external binding will be prohibited from using the API for security reasons.
type SoftLayer_User_Customer_External_Binding_Totp struct {
}

func (softlayer_user_customer_external_binding_totp *SoftLayer_User_Customer_External_Binding_Totp) String() string {
	return "SoftLayer_User_Customer_External_Binding_Totp"
}

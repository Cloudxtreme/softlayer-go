package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

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

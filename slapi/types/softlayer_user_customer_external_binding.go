package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_Customer_External_Binding - The SoftLayer_User_Customer_External_Binding data type
// contains general information for a single external binding. This includes the 3rd party vendor, type
// of binding, and a unique identifier and password that is used to authenticate against the 3rd party
// service.
type SoftLayer_User_Customer_External_Binding struct {
}

// SoftLayer_User_Customer_External_Binding_Extended is SoftLayer_User_Customer_External_Binding with all maskable types.
type SoftLayer_User_Customer_External_Binding_Extended struct {
	SoftLayer_User_Customer_External_Binding

	// User - The SoftLayer user that the external authentication binding belongs to.
	User *SoftLayer_User_Customer `json:"user"`
}

func (softlayer_user_customer_external_binding *SoftLayer_User_Customer_External_Binding) String() string {
	return "SoftLayer_User_Customer_External_Binding"
}

package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

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

func (softlayer_user_external_binding_vendor *SoftLayer_User_External_Binding_Vendor) String() string {
	return "SoftLayer_User_External_Binding_Vendor"
}

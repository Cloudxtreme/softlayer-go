package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_User_Customer_External_Binding_Vendor - Container classed used to hold details
// about an external authentication vendor.
type SoftLayer_Container_User_Customer_External_Binding_Vendor struct {

	// KeyName - The keyname used to identify an external authentication vendor.
	KeyName string `json:"keyName,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`
}

func (softlayer_container_user_customer_external_binding_vendor *SoftLayer_Container_User_Customer_External_Binding_Vendor) String() string {
	return "SoftLayer_Container_User_Customer_External_Binding_Vendor"
}

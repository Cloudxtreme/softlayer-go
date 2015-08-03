package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_User_Customer_External_Binding - This container type is used for
// placing orders for external authentication, such as phone-based authentication.
type SoftLayer_Container_Product_Order_User_Customer_External_Binding struct {

	// ExternalId - The external id that access to external authentication is being purchased for.
	ExternalId string `json:"externalId"`

	// UserId - The SoftLayer [[SoftLayer_User_Customer|user]] identifier that an external binding is being
	// purchased for.
	UserId int `json:"userId"`

	// VendorId - The [[SoftLayer_User_Customer_External_Binding_Vendor|vendor]] identifier for the
	// external binding being purchased.
	VendorId int `json:"vendorId"`
}

func (softlayer_container_product_order_user_customer_external_binding *SoftLayer_Container_Product_Order_User_Customer_External_Binding) String() string {
	return "SoftLayer_Container_Product_Order_User_Customer_External_Binding"
}

package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_User_Customer_External_Binding - The
// SoftLayer_Billing_Item_Network_Application_Delivery_Controller data type describes the billing item
// related to an external authentication binding
type SoftLayer_Billing_Item_User_Customer_External_Binding struct {
}

func (softlayer_billing_item_user_customer_external_binding *SoftLayer_Billing_Item_User_Customer_External_Binding) String() string {
	return "SoftLayer_Billing_Item_User_Customer_External_Binding"
}

// SoftLayer_Billing_Item_User_Customer_External_Binding_Extended is SoftLayer_Billing_Item_User_Customer_External_Binding with all maskable types.
type SoftLayer_Billing_Item_User_Customer_External_Binding_Extended struct {
	SoftLayer_Billing_Item_User_Customer_External_Binding

	// Resource - The external authentication binding that a billing item is associated with.
	Resource *SoftLayer_User_Customer_External_Binding `json:"resource"`
}

func (softlayer_billing_item_user_customer_external_binding *SoftLayer_Billing_Item_User_Customer_External_Binding_Extended) String() string {
	return "SoftLayer_Billing_Item_User_Customer_External_Binding"
}

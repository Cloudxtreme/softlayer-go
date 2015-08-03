package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Network_Message_Delivery - The
// SoftLayer_Billing_Item_Network_Message_Delivery data describes the related billing item.
type SoftLayer_Billing_Item_Network_Message_Delivery struct {
}

func (softlayer_billing_item_network_message_delivery *SoftLayer_Billing_Item_Network_Message_Delivery) String() string {
	return "SoftLayer_Billing_Item_Network_Message_Delivery"
}

// SoftLayer_Billing_Item_Network_Message_Delivery_Extended is SoftLayer_Billing_Item_Network_Message_Delivery with all maskable types.
type SoftLayer_Billing_Item_Network_Message_Delivery_Extended struct {
	SoftLayer_Billing_Item_Network_Message_Delivery

	// Resource - no documentation
	Resource *SoftLayer_Network_Message_Delivery `json:"resource"`
}

func (softlayer_billing_item_network_message_delivery *SoftLayer_Billing_Item_Network_Message_Delivery_Extended) String() string {
	return "SoftLayer_Billing_Item_Network_Message_Delivery"
}

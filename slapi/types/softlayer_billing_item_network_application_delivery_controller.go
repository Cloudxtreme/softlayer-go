package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Network_Application_Delivery_Controller - The
// SoftLayer_Billing_Item_Network_Application_Delivery_Controller data type describes the billing item
// related to a NetScaler VPX
type SoftLayer_Billing_Item_Network_Application_Delivery_Controller struct {
}

// SoftLayer_Billing_Item_Network_Application_Delivery_Controller_Extended is SoftLayer_Billing_Item_Network_Application_Delivery_Controller with all maskable types.
type SoftLayer_Billing_Item_Network_Application_Delivery_Controller_Extended struct {
	SoftLayer_Billing_Item_Network_Application_Delivery_Controller

	// BandwidthAllotmentDetail - no documentation
	BandwidthAllotmentDetail *SoftLayer_Network_Bandwidth_Version1_Allotment_Detail `json:"bandwidthAllotmentDetail"`

	// Resource - The network application controller that a billing item is associated with.
	Resource *SoftLayer_Network_Application_Delivery_Controller `json:"resource"`
}

func (softlayer_billing_item_network_application_delivery_controller *SoftLayer_Billing_Item_Network_Application_Delivery_Controller) String() string {
	return "SoftLayer_Billing_Item_Network_Application_Delivery_Controller"
}

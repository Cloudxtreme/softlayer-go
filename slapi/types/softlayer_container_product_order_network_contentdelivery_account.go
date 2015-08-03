package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Network_ContentDelivery_Account - This is the datatype that needs
// to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything
// required to place a CDN order with SoftLayer.
type SoftLayer_Container_Product_Order_Network_ContentDelivery_Account struct {

	// CdnAccountName - no documentation
	CdnAccountName string `json:"cdnAccountName"`
}

func (softlayer_container_product_order_network_contentdelivery_account *SoftLayer_Container_Product_Order_Network_ContentDelivery_Account) String() string {
	return "SoftLayer_Container_Product_Order_Network_ContentDelivery_Account"
}

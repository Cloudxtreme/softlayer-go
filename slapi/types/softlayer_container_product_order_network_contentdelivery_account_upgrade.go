package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Network_ContentDelivery_Account_Upgrade - This is the datatype
// that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has
// everything required to place a CDN order with SoftLayer.
type SoftLayer_Container_Product_Order_Network_ContentDelivery_Account_Upgrade struct {

	// CdnAccountId - ID of an existing CDN account. You can use this to upgrade an existing CDN account.
	CdnAccountId string `json:"cdnAccountId"`
}

func (softlayer_container_product_order_network_contentdelivery_account_upgrade *SoftLayer_Container_Product_Order_Network_ContentDelivery_Account_Upgrade) String() string {
	return "SoftLayer_Container_Product_Order_Network_ContentDelivery_Account_Upgrade"
}

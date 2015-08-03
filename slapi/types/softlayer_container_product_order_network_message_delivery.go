package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Network_Message_Delivery - This is the datatype that needs to be
// populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to
// place a network message delivery order with SoftLayer.
type SoftLayer_Container_Product_Order_Network_Message_Delivery struct {

	// AccountPassword - no documentation
	AccountPassword string `json:"accountPassword"`

	// AccountUsername - no documentation
	AccountUsername string `json:"accountUsername"`

	// EmailAddress - no documentation
	EmailAddress string `json:"emailAddress"`
}

func (softlayer_container_product_order_network_message_delivery *SoftLayer_Container_Product_Order_Network_Message_Delivery) String() string {
	return "SoftLayer_Container_Product_Order_Network_Message_Delivery"
}

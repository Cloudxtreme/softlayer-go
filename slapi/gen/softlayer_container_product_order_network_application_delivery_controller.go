package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Network_Application_Delivery_Controller - This is the datatype
// that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has
// everything required to place an application delivery controller order with SoftLayer.
type SoftLayer_Container_Product_Order_Network_Application_Delivery_Controller struct {

	// ApplicationDeliveryControllerId - An optional [[SoftLayer_Network_Application_Delivery_Controller]]
	// identifier that is used for upgrading an existing application delivery controller.
	ApplicationDeliveryControllerId int `json:"applicationDeliveryControllerId"`
}

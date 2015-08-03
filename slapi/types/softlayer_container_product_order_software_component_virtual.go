package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Software_Component_Virtual - This is the datatype that needs to be
// populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to
// place a virtual license order with SoftLayer.
type SoftLayer_Container_Product_Order_Software_Component_Virtual struct {

	// EndPointIpAddressIds - array of ip address ids for which a license should be allocated for.
	EndPointIpAddressIds []int `json:"endPointIpAddressIds"`
}

func (softlayer_container_product_order_software_component_virtual *SoftLayer_Container_Product_Order_Software_Component_Virtual) String() string {
	return "SoftLayer_Container_Product_Order_Software_Component_Virtual"
}

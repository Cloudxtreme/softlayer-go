package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Property - This is used for storing various items about the order.
// Currently used for storing additional raid information when ordering servers. This is optional
type SoftLayer_Container_Product_Order_Property struct {

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// Value - no documentation
	Value string `json:"value,omitempty"`
}

func (softlayer_container_product_order_property *SoftLayer_Container_Product_Order_Property) String() string {
	return "SoftLayer_Container_Product_Order_Property"
}

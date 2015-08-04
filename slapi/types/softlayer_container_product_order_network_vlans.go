package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Network_Vlans - This class contains the collections of public and
// private VLANs that are available during the ordering process.
type SoftLayer_Container_Product_Order_Network_Vlans struct {

	// PublicVlans - The collection of public vlans available during ordering.
	PublicVlans []*SoftLayer_Container_Product_Order `json:"publicVlans,omitempty"`

	// PrivateVlans - The collection of private vlans available during ordering.
	PrivateVlans []*SoftLayer_Container_Product_Order `json:"privateVlans,omitempty"`
}

func (softlayer_container_product_order_network_vlans *SoftLayer_Container_Product_Order_Network_Vlans) String() string {
	return "SoftLayer_Container_Product_Order_Network_Vlans"
}

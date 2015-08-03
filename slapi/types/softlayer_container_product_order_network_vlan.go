package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Network_Vlan - This is the datatype that needs to be populated and
// sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place a
// network vlan order with SoftLayer.
type SoftLayer_Container_Product_Order_Network_Vlan struct {

	// Id - no documentation
	Id int `json:"id"`

	// Name - no documentation
	Name string `json:"name"`

	// Subnets - The collection of subnets associated with this vlan.
	Subnets []*SoftLayer_Container_Product_Order `json:"subnets"`

	// Description - The description which includes the primary router's hostname plus the vlan number.
	Description string `json:"description"`

	// HostnameRouter - no documentation
	HostnameRouter string `json:"hostnameRouter"`

	// RouterId - The ID of the [[SoftLayer_Hardware_Router]] object on which the new should be created.
	RouterId int `json:"routerId"`

	// VlanNumber - no documentation
	VlanNumber int `json:"vlanNumber"`

	// HostnameDatacenter - no documentation
	HostnameDatacenter string `json:"hostnameDatacenter"`

	// Router - The router object on which the new should be created.
	Router *SoftLayer_Hardware `json:"router"`
}

func (softlayer_container_product_order_network_vlan *SoftLayer_Container_Product_Order_Network_Vlan) String() string {
	return "SoftLayer_Container_Product_Order_Network_Vlan"
}

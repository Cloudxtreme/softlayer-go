package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Network_Vlan - This is the datatype that needs to be populated and
// sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place a
// network vlan order with SoftLayer.
type SoftLayer_Container_Product_Order_Network_Vlan struct {

	// Description - The description which includes the primary router's hostname plus the vlan number.
	Description string `json:"description"`

	// Name - no documentation
	Name string `json:"name"`

	// Router - The router object on which the new should be created.
	Router *SoftLayer_Hardware `json:"router"`

	// RouterId - The ID of the [[SoftLayer_Hardware_Router]] object on which the new should be created.
	RouterId int `json:"routerId"`

	// HostnameDatacenter - no documentation
	HostnameDatacenter string `json:"hostnameDatacenter"`

	// HostnameRouter - no documentation
	HostnameRouter string `json:"hostnameRouter"`

	// Id - no documentation
	Id int `json:"id"`

	// Subnets - The collection of subnets associated with this vlan.
	Subnets []*SoftLayer_Container_Product_Order `json:"subnets"`

	// VlanNumber - no documentation
	VlanNumber int `json:"vlanNumber"`
}

func (softlayer_container_product_order_network_vlan *SoftLayer_Container_Product_Order_Network_Vlan) String() string {
	return "SoftLayer_Container_Product_Order_Network_Vlan"
}

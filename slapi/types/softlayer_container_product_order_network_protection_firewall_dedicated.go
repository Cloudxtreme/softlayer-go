package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Network_Protection_Firewall_Dedicated - This is the datatype that
// needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything
// required to place a hardware (dedicated) firewall order with SoftLayer.
type SoftLayer_Container_Product_Order_Network_Protection_Firewall_Dedicated struct {

	// Vlan - no documentation
	Vlan *SoftLayer_Network_Vlan `json:"vlan"`

	// VlanId - no documentation
	VlanId int `json:"vlanId"`
}

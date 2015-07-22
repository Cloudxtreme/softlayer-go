package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Network_Subnet - This is the datatype that needs to be populated
// and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place a
// subnet order with SoftLayer.
type SoftLayer_Container_Product_Order_Network_Subnet struct {

	// Description - The description which includes the network identifier, Classless Inter-Domain Routing
	// prefix and the available slot count.
	Description string `json:"description"`

	// EndPointIpAddressId - no documentation
	EndPointIpAddressId int `json:"endPointIpAddressId"`

	// EndPointVlanId - no documentation
	EndPointVlanId int `json:"endPointVlanId"`

	// Id - no documentation
	Id int `json:"id"`
}

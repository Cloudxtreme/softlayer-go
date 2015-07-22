package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Network_Vlans - This class contains the collections of public and
// private VLANs that are available during the ordering process.
type SoftLayer_Container_Product_Order_Network_Vlans struct {

	// PrivateVlans - The collection of private vlans available during ordering.
	PrivateVlans []*SoftLayer_Container_Product_Order `json:"privateVlans"`

	// PublicVlans - The collection of public vlans available during ordering.
	PublicVlans []*SoftLayer_Container_Product_Order `json:"publicVlans"`
}

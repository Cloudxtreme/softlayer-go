package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Utility_Network_Subnet_Mask_Generic_Detail - The
// SoftLayer_Container_Utility_Network_Subnet_Mask_Generic_Detail data type contains information
// relating to a subnet mask and details associated with that object.
type SoftLayer_Container_Utility_Network_Subnet_Mask_Generic_Detail struct {

	// Cidr - no documentation
	Cidr string `json:"cidr"`

	// Description - no documentation
	Description string `json:"description"`

	// Mask - no documentation
	Mask string `json:"mask"`
}

func (softlayer_container_utility_network_subnet_mask_generic_detail *SoftLayer_Container_Utility_Network_Subnet_Mask_Generic_Detail) String() string {
	return "SoftLayer_Container_Utility_Network_Subnet_Mask_Generic_Detail"
}

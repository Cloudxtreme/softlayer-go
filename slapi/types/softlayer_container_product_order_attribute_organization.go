package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Attribute_Organization - This is the datatype that needs to be
// populated and sent to SoftLayer_Product_Order::placeOrder. The
// SoftLayer_Container_Product_Order_Attribute_Organization datatype contains the organization
// information.
type SoftLayer_Container_Product_Order_Attribute_Organization struct {

	// Address - no documentation
	Address *SoftLayer_Container_Product_Order_Attribute_Address `json:"address,omitempty"`

	// FaxNumber - The fax number associated with an organization. This is an optional value.
	FaxNumber string `json:"faxNumber,omitempty"`

	// OrganizationName - no documentation
	OrganizationName string `json:"organizationName,omitempty"`

	// PhoneNumber - no documentation
	PhoneNumber string `json:"phoneNumber,omitempty"`
}

func (softlayer_container_product_order_attribute_organization *SoftLayer_Container_Product_Order_Attribute_Organization) String() string {
	return "SoftLayer_Container_Product_Order_Attribute_Organization"
}

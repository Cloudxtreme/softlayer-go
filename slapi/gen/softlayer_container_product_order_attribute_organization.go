package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Attribute_Organization - This is the datatype that needs to be
// populated and sent to SoftLayer_Product_Order::placeOrder. The
// SoftLayer_Container_Product_Order_Attribute_Organization datatype contains the organization
// information.
type SoftLayer_Container_Product_Order_Attribute_Organization struct {

	// Address - no documentation
	Address *SoftLayer_Container_Product_Order_Attribute_Address `json:"address"`

	// FaxNumber - The fax number associated with an organization. This is an optional value.
	FaxNumber string `json:"faxNumber"`

	// OrganizationName - no documentation
	OrganizationName string `json:"organizationName"`

	// PhoneNumber - no documentation
	PhoneNumber string `json:"phoneNumber"`
}

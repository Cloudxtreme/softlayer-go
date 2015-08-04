package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Attribute_Contact - This is the datatype that needs to be
// populated and sent to SoftLayer_Product_Order::placeOrder. The
// SoftLayer_Container_Product_Order_Attribute_Contact datatype contains the contact information.
type SoftLayer_Container_Product_Order_Attribute_Contact struct {

	// EmailAddress - no documentation
	EmailAddress string `json:"emailAddress,omitempty"`

	// FaxNumber - The fax number associated with a contact. This is an optional value.
	FaxNumber string `json:"faxNumber,omitempty"`

	// FirstName - no documentation
	FirstName string `json:"firstName,omitempty"`

	// LastName - no documentation
	LastName string `json:"lastName,omitempty"`

	// OrganizationName - no documentation
	OrganizationName string `json:"organizationName,omitempty"`

	// PhoneNumber - no documentation
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// Title - no documentation
	Title string `json:"title,omitempty"`

	// Address - no documentation
	Address *SoftLayer_Container_Product_Order_Attribute_Address `json:"address,omitempty"`
}

func (softlayer_container_product_order_attribute_contact *SoftLayer_Container_Product_Order_Attribute_Contact) String() string {
	return "SoftLayer_Container_Product_Order_Attribute_Contact"
}

package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Attribute_Address - This is the datatype that needs to be
// populated and sent to SoftLayer_Product_Order::placeOrder. The
// SoftLayer_Container_Product_Order_Attribute_Address datatype contains the address information.
type SoftLayer_Container_Product_Order_Attribute_Address struct {

	// PostalCode - no documentation
	PostalCode string `json:"postalCode,omitempty"`

	// State - no documentation
	State string `json:"state,omitempty"`

	// AddressLine1 - no documentation
	AddressLine1 string `json:"addressLine1,omitempty"`

	// AddressLine2 - The second line in the address. Information such as suite number goes here.
	AddressLine2 string `json:"addressLine2,omitempty"`

	// City - no documentation
	City string `json:"city,omitempty"`

	// CountryCode - no documentation
	CountryCode string `json:"countryCode,omitempty"`

	// NonUsState - no documentation
	NonUsState string `json:"nonUsState,omitempty"`
}

func (softlayer_container_product_order_attribute_address *SoftLayer_Container_Product_Order_Attribute_Address) String() string {
	return "SoftLayer_Container_Product_Order_Attribute_Address"
}

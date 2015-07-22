package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Attribute_Address - This is the datatype that needs to be
// populated and sent to SoftLayer_Product_Order::placeOrder. The
// SoftLayer_Container_Product_Order_Attribute_Address datatype contains the address information.
type SoftLayer_Container_Product_Order_Attribute_Address struct {

	// AddressLine1 - no documentation
	AddressLine1 string `json:"addressLine1"`

	// AddressLine2 - The second line in the address. Information such as suite number goes here.
	AddressLine2 string `json:"addressLine2"`

	// City - no documentation
	City string `json:"city"`

	// CountryCode - no documentation
	CountryCode string `json:"countryCode"`

	// NonUsState - no documentation
	NonUsState string `json:"nonUsState"`

	// PostalCode - no documentation
	PostalCode string `json:"postalCode"`

	// State - no documentation
	State string `json:"state"`
}

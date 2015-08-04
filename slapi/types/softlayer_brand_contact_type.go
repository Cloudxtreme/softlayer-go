package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Brand_Contact_Type - SoftLayer_Brand_Contact_Type contains the contact type information
// for the brand contacts such as Corporate or Support contact type
type SoftLayer_Brand_Contact_Type struct {

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// KeyName - no documentation
	KeyName string `json:"keyName,omitempty"`
}

func (softlayer_brand_contact_type *SoftLayer_Brand_Contact_Type) String() string {
	return "SoftLayer_Brand_Contact_Type"
}

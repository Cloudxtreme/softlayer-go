package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Brand_Contact - SoftLayer_Brand_Contact contains the contact information for the brand
// such as Corporate or Support contact information
type SoftLayer_Brand_Contact struct {

	// Address1 - no documentation
	Address1 string `json:"address1,omitempty"`

	// Email - no documentation
	Email string `json:"email,omitempty"`

	// PostalCode - no documentation
	PostalCode string `json:"postalCode,omitempty"`

	// State - no documentation
	State string `json:"state,omitempty"`

	// AlternatePhone - no documentation
	AlternatePhone string `json:"alternatePhone,omitempty"`

	// City - no documentation
	City string `json:"city,omitempty"`

	// FaxPhone - no documentation
	FaxPhone string `json:"faxPhone,omitempty"`

	// BrandContactTypeId - no documentation
	BrandContactTypeId int `json:"brandContactTypeId,omitempty"`

	// Address2 - no documentation
	Address2 string `json:"address2,omitempty"`

	// FirstName - no documentation
	FirstName string `json:"firstName,omitempty"`

	// LastName - no documentation
	LastName string `json:"lastName,omitempty"`

	// OfficePhone - no documentation
	OfficePhone string `json:"officePhone,omitempty"`

	// Country - no documentation
	Country string `json:"country,omitempty"`

	// Brand - <nil>
	Brand *SoftLayer_Brand `json:"brand,omitempty"`

	// BrandContactType - <nil>
	BrandContactType *SoftLayer_Brand_Contact_Type `json:"brandContactType,omitempty"`
}

func (softlayer_brand_contact *SoftLayer_Brand_Contact) String() string {
	return "SoftLayer_Brand_Contact"
}

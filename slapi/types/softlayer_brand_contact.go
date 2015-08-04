package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Brand_Contact - SoftLayer_Brand_Contact contains the contact information for the brand
// such as Corporate or Support contact information
type SoftLayer_Brand_Contact struct {

	// Address1 - no documentation
	Address1 string `json:"address1,omitempty"`

	// AlternatePhone - no documentation
	AlternatePhone string `json:"alternatePhone,omitempty"`

	// Email - no documentation
	Email string `json:"email,omitempty"`

	// FirstName - no documentation
	FirstName string `json:"firstName,omitempty"`

	// LastName - no documentation
	LastName string `json:"lastName,omitempty"`

	// BrandContactTypeId - no documentation
	BrandContactTypeId int `json:"brandContactTypeId,omitempty"`

	// State - no documentation
	State string `json:"state,omitempty"`

	// Country - no documentation
	Country string `json:"country,omitempty"`

	// FaxPhone - no documentation
	FaxPhone string `json:"faxPhone,omitempty"`

	// OfficePhone - no documentation
	OfficePhone string `json:"officePhone,omitempty"`

	// PostalCode - no documentation
	PostalCode string `json:"postalCode,omitempty"`

	// Address2 - no documentation
	Address2 string `json:"address2,omitempty"`

	// City - no documentation
	City string `json:"city,omitempty"`
}

func (softlayer_brand_contact *SoftLayer_Brand_Contact) String() string {
	return "SoftLayer_Brand_Contact"
}

// SoftLayer_Brand_Contact_Extended is SoftLayer_Brand_Contact with all maskable types.
type SoftLayer_Brand_Contact_Extended struct {
	SoftLayer_Brand_Contact

	// BrandContactType - <nil>
	BrandContactType *SoftLayer_Brand_Contact_Type `json:"brandContactType,omitempty"`

	// Brand - <nil>
	Brand *SoftLayer_Brand `json:"brand,omitempty"`
}

func (softlayer_brand_contact *SoftLayer_Brand_Contact_Extended) String() string {
	return "SoftLayer_Brand_Contact"
}

package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Brand_Contact - SoftLayer_Brand_Contact contains the contact information for the brand
// such as Corporate or Support contact information
type SoftLayer_Brand_Contact struct {

	// Country - no documentation
	Country string `json:"country"`

	// FaxPhone - no documentation
	FaxPhone string `json:"faxPhone"`

	// State - no documentation
	State string `json:"state"`

	// LastName - no documentation
	LastName string `json:"lastName"`

	// PostalCode - no documentation
	PostalCode string `json:"postalCode"`

	// Address2 - no documentation
	Address2 string `json:"address2"`

	// AlternatePhone - no documentation
	AlternatePhone string `json:"alternatePhone"`

	// BrandContactTypeId - no documentation
	BrandContactTypeId int `json:"brandContactTypeId"`

	// Email - no documentation
	Email string `json:"email"`

	// FirstName - no documentation
	FirstName string `json:"firstName"`

	// Address1 - no documentation
	Address1 string `json:"address1"`

	// City - no documentation
	City string `json:"city"`

	// OfficePhone - no documentation
	OfficePhone string `json:"officePhone"`
}

func (softlayer_brand_contact *SoftLayer_Brand_Contact) String() string {
	return "SoftLayer_Brand_Contact"
}

// SoftLayer_Brand_Contact_Extended is SoftLayer_Brand_Contact with all maskable types.
type SoftLayer_Brand_Contact_Extended struct {
	SoftLayer_Brand_Contact

	// Brand - <nil>
	Brand *SoftLayer_Brand `json:"brand"`

	// BrandContactType - <nil>
	BrandContactType *SoftLayer_Brand_Contact_Type `json:"brandContactType"`
}

func (softlayer_brand_contact *SoftLayer_Brand_Contact_Extended) String() string {
	return "SoftLayer_Brand_Contact"
}

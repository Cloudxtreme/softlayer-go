package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Brand_Contact - SoftLayer_Brand_Contact contains the contact information for the brand
// such as Corporate or Support contact information
type SoftLayer_Brand_Contact struct {

	// Address1 - no documentation
	Address1 string `json:"address1"`

	// Address2 - no documentation
	Address2 string `json:"address2"`

	// AlternatePhone - no documentation
	AlternatePhone string `json:"alternatePhone"`

	// Brand - <nil>
	Brand *SoftLayer_Brand `json:"brand"`

	// BrandContactType - <nil>
	BrandContactType *SoftLayer_Brand_Contact_Type `json:"brandContactType"`

	// BrandContactTypeId - no documentation
	BrandContactTypeId int `json:"brandContactTypeId"`

	// City - no documentation
	City string `json:"city"`

	// Country - no documentation
	Country string `json:"country"`

	// Email - no documentation
	Email string `json:"email"`

	// FaxPhone - no documentation
	FaxPhone string `json:"faxPhone"`

	// FirstName - no documentation
	FirstName string `json:"firstName"`

	// LastName - no documentation
	LastName string `json:"lastName"`

	// OfficePhone - no documentation
	OfficePhone string `json:"officePhone"`

	// PostalCode - no documentation
	PostalCode string `json:"postalCode"`

	// State - no documentation
	State string `json:"state"`
}

func (softlayer_brand_contact *SoftLayer_Brand_Contact) String() string {
	return "SoftLayer_Brand_Contact"
}

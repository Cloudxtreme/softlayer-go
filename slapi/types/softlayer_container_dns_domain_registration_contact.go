package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Dns_Domain_Registration_Contact - Contact information container for domain
// registration
type SoftLayer_Container_Dns_Domain_Registration_Contact struct {

	// Address1 - no documentation
	Address1 string `json:"address1"`

	// Address2 - no documentation
	Address2 string `json:"address2"`

	// Address3 - no documentation
	Address3 string `json:"address3"`

	// City - no documentation
	City string `json:"city"`

	// Country - no documentation
	Country string `json:"country"`

	// Email - no documentation
	Email string `json:"email"`

	// Fax - no documentation
	Fax string `json:"fax"`

	// FirstName - no documentation
	FirstName string `json:"firstName"`

	// LastName - no documentation
	LastName string `json:"lastName"`

	// OrganizationName - no documentation
	OrganizationName string `json:"organizationName"`

	// Phone - no documentation
	Phone string `json:"phone"`

	// PostalCode - no documentation
	PostalCode string `json:"postalCode"`

	// State - no documentation
	State string `json:"state"`

	// Type - The type of contact. The following are the valid types of contacts: * admin * owner * billing
	// * tech
	Type string `json:"type"`
}

func (softlayer_container_dns_domain_registration_contact *SoftLayer_Container_Dns_Domain_Registration_Contact) String() string {
	return "SoftLayer_Container_Dns_Domain_Registration_Contact"
}

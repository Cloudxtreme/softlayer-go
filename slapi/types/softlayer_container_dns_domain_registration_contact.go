package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Dns_Domain_Registration_Contact - Contact information container for domain
// registration
type SoftLayer_Container_Dns_Domain_Registration_Contact struct {

	// Phone - no documentation
	Phone string `json:"phone"`

	// State - no documentation
	State string `json:"state"`

	// Email - no documentation
	Email string `json:"email"`

	// City - no documentation
	City string `json:"city"`

	// Fax - no documentation
	Fax string `json:"fax"`

	// LastName - no documentation
	LastName string `json:"lastName"`

	// OrganizationName - no documentation
	OrganizationName string `json:"organizationName"`

	// PostalCode - no documentation
	PostalCode string `json:"postalCode"`

	// Type - The type of contact. The following are the valid types of contacts: * admin * owner * billing
	// * tech
	Type string `json:"type"`

	// Address3 - no documentation
	Address3 string `json:"address3"`

	// Address2 - no documentation
	Address2 string `json:"address2"`

	// FirstName - no documentation
	FirstName string `json:"firstName"`

	// Address1 - no documentation
	Address1 string `json:"address1"`

	// Country - no documentation
	Country string `json:"country"`
}

func (softlayer_container_dns_domain_registration_contact *SoftLayer_Container_Dns_Domain_Registration_Contact) String() string {
	return "SoftLayer_Container_Dns_Domain_Registration_Contact"
}

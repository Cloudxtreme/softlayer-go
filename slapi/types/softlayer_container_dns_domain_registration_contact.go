package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Dns_Domain_Registration_Contact - Contact information container for domain
// registration
type SoftLayer_Container_Dns_Domain_Registration_Contact struct {

	// Address1 - no documentation
	Address1 string `json:"address1,omitempty"`

	// City - no documentation
	City string `json:"city,omitempty"`

	// Country - no documentation
	Country string `json:"country,omitempty"`

	// FirstName - no documentation
	FirstName string `json:"firstName,omitempty"`

	// OrganizationName - no documentation
	OrganizationName string `json:"organizationName,omitempty"`

	// Phone - no documentation
	Phone string `json:"phone,omitempty"`

	// PostalCode - no documentation
	PostalCode string `json:"postalCode,omitempty"`

	// Type - The type of contact. The following are the valid types of contacts: * admin * owner * billing
	// * tech
	Type string `json:"type,omitempty"`

	// Address2 - no documentation
	Address2 string `json:"address2,omitempty"`

	// Address3 - no documentation
	Address3 string `json:"address3,omitempty"`

	// LastName - no documentation
	LastName string `json:"lastName,omitempty"`

	// Email - no documentation
	Email string `json:"email,omitempty"`

	// Fax - no documentation
	Fax string `json:"fax,omitempty"`

	// State - no documentation
	State string `json:"state,omitempty"`
}

func (softlayer_container_dns_domain_registration_contact *SoftLayer_Container_Dns_Domain_Registration_Contact) String() string {
	return "SoftLayer_Container_Dns_Domain_Registration_Contact"
}

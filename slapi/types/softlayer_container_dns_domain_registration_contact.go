package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Dns_Domain_Registration_Contact - Contact information container for domain
// registration
type SoftLayer_Container_Dns_Domain_Registration_Contact struct {

	// Address1 - no documentation
	Address1 string `json:"address1,omitempty"`

	// Fax - no documentation
	Fax string `json:"fax,omitempty"`

	// State - no documentation
	State string `json:"state,omitempty"`

	// Type - The type of contact. The following are the valid types of contacts: * admin * owner * billing
	// * tech
	Type string `json:"type,omitempty"`

	// Country - no documentation
	Country string `json:"country,omitempty"`

	// PostalCode - no documentation
	PostalCode string `json:"postalCode,omitempty"`

	// Phone - no documentation
	Phone string `json:"phone,omitempty"`

	// Address3 - no documentation
	Address3 string `json:"address3,omitempty"`

	// OrganizationName - no documentation
	OrganizationName string `json:"organizationName,omitempty"`

	// Email - no documentation
	Email string `json:"email,omitempty"`

	// FirstName - no documentation
	FirstName string `json:"firstName,omitempty"`

	// LastName - no documentation
	LastName string `json:"lastName,omitempty"`

	// Address2 - no documentation
	Address2 string `json:"address2,omitempty"`

	// City - no documentation
	City string `json:"city,omitempty"`
}

func (softlayer_container_dns_domain_registration_contact *SoftLayer_Container_Dns_Domain_Registration_Contact) String() string {
	return "SoftLayer_Container_Dns_Domain_Registration_Contact"
}

package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Dns_Domain_Registration_Lookup_Items - no documentation
type SoftLayer_Container_Dns_Domain_Registration_Lookup_Items struct {

	// DomainName - no documentation
	DomainName string `json:"domainName,omitempty"`

	// Status - The status of the domain name if available and can be registered.
	Status string `json:"status,omitempty"`
}

func (softlayer_container_dns_domain_registration_lookup_items *SoftLayer_Container_Dns_Domain_Registration_Lookup_Items) String() string {
	return "SoftLayer_Container_Dns_Domain_Registration_Lookup_Items"
}

package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Dns_Domain_Registration_Nameserver_List - Nameservers list container for domain
// registration
type SoftLayer_Container_Dns_Domain_Registration_Nameserver_List struct {

	// Ipv4Address - no documentation
	Ipv4Address string `json:"ipv4Address"`

	// Ipv6Address - no documentation
	Ipv6Address string `json:"ipv6Address"`

	// Name - no documentation
	Name string `json:"name"`

	// SortOrder - no documentation
	SortOrder int `json:"sortOrder"`
}

func (softlayer_container_dns_domain_registration_nameserver_list *SoftLayer_Container_Dns_Domain_Registration_Nameserver_List) String() string {
	return "SoftLayer_Container_Dns_Domain_Registration_Nameserver_List"
}

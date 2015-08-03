package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Dns_Domain_Registration_Nameserver - no documentation
type SoftLayer_Container_Dns_Domain_Registration_Nameserver struct {

	// Nameservers - The list of fully qualified names of the nameserver.
	Nameservers []*SoftLayer_Container_Dns_Domain_Registration_Nameserver_List `json:"nameservers"`
}

func (softlayer_container_dns_domain_registration_nameserver *SoftLayer_Container_Dns_Domain_Registration_Nameserver) String() string {
	return "SoftLayer_Container_Dns_Domain_Registration_Nameserver"
}

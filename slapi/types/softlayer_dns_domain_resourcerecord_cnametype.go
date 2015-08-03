package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Dns_Domain_ResourceRecord_CnameType - SoftLayer_Dns_Domain_ResourceRecord_CnameType is a
// SoftLayer_Dns_Domain_ResourceRecord object whose ''type'' property is set to "cname" and defines a
// DNS record on a SoftLayer hosted domain. A record directs a host name to another host. For instance,
// if the record for "alias.example.org" points to the host "host.example.org" then the ''host''
// property equals "alias" and the ''data'' property equals "host.example.org.". DNS entries defined by
// should not be used as the data field for an MX record.
type SoftLayer_Dns_Domain_ResourceRecord_CnameType struct {
}

func (softlayer_dns_domain_resourcerecord_cnametype *SoftLayer_Dns_Domain_ResourceRecord_CnameType) String() string {
	return "SoftLayer_Dns_Domain_ResourceRecord_CnameType"
}

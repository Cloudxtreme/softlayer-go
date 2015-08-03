package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Dns_Domain_ResourceRecord_AType - SoftLayer_Dns_Domain_ResourceRecord_AType is a
// SoftLayer_Dns_Domain_ResourceRecord object whose ''type'' property is set to "a" and defines a DNS A
// record on a SoftLayer hosted domain. An A record directs a host name to an IP address. For instance
// if the A record for "host.example.org" points to the IP address 10.0.0.1 then the ''host'' property
// for the A record equals "host" and the ''data'' property equals "10.0.0.1".
type SoftLayer_Dns_Domain_ResourceRecord_AType struct {
}

func (softlayer_dns_domain_resourcerecord_atype *SoftLayer_Dns_Domain_ResourceRecord_AType) String() string {
	return "SoftLayer_Dns_Domain_ResourceRecord_AType"
}

package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Dns_Domain_ResourceRecord_SpfType - SoftLayer_Dns_Domain_ResourceRecord_SpfType is a
// SoftLayer_Dns_Domain_ResourceRecord object whose ''type'' property is set to "spf" and defines a DNS
// SPF record on a SoftLayer hosted domain. An SPF record provides sender policy framework data for a
// host. For instance, if defining the SPF record "v=spf1 mx:mail.example.org ~all" for
// "host.example.org". then the ''host'' property equals "host" and the ''data'' property equals
// "v=spf1 mx:mail.example.org ~all". SPF records are commonly used in email verification methods such
// as Sender Policy Framework.
type SoftLayer_Dns_Domain_ResourceRecord_SpfType struct {
}

func (softlayer_dns_domain_resourcerecord_spftype *SoftLayer_Dns_Domain_ResourceRecord_SpfType) String() string {
	return "SoftLayer_Dns_Domain_ResourceRecord_SpfType"
}

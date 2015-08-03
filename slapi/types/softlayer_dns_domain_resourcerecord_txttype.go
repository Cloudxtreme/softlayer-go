package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Dns_Domain_ResourceRecord_TxtType - SoftLayer_Dns_Domain_ResourceRecord_TxtType is a
// SoftLayer_Dns_Domain_ResourceRecord object whose ''type'' property is set to "txt" and defines a DNS
// TXT record on a SoftLayer hosted domain. A TXT record provides a text description for a host. For
// instance, if defining the TXT record "My test host" for "host.example.org". then the ''host''
// property equals "host" and the ''data'' property equals "My test host". TXT records are commonly
// used in email verification methods such as Sender Policy Framework.
type SoftLayer_Dns_Domain_ResourceRecord_TxtType struct {
}

func (softlayer_dns_domain_resourcerecord_txttype *SoftLayer_Dns_Domain_ResourceRecord_TxtType) String() string {
	return "SoftLayer_Dns_Domain_ResourceRecord_TxtType"
}

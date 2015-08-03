package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Dns_Domain_ResourceRecord_MxType - SoftLayer_Dns_Domain_ResourceRecord_MxType is a
// SoftLayer_Dns_Domain_ResourceRecord object whose ''type'' property is set to "mx" and used to
// describe MX resource records. MX records control which hosts are responsible as mail exchangers for
// a domain. For instance, in the domain example.org, an MX record whose host is "@" and data is "mail"
// says that the host "mail.example.org" is responsible for handling mail for example.org. That means
// mail sent to users @example.org are delivered to mail.example.org. Domains can have more than one MX
// record if it uses more than one server to send mail through. Multiple MX records are denoted by
// their priority, defined by the mxPriority property. MX records must be defined for hosts with
// accompanying A or resource records. They may not point mail towards a host defined by a record.
type SoftLayer_Dns_Domain_ResourceRecord_MxType struct {
}

func (softlayer_dns_domain_resourcerecord_mxtype *SoftLayer_Dns_Domain_ResourceRecord_MxType) String() string {
	return "SoftLayer_Dns_Domain_ResourceRecord_MxType"
}

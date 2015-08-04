package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Dns_Domain_ResourceRecord_PtrType - SoftLayer_Dns_Domain_ResourceRecord_PtrType is a
// SoftLayer_Dns_Domain_ResourceRecord object whose ''type'' property is set to "ptr" and defines a
// reverse DNS PTR record on the SoftLayer name servers. The format for a reverse DNS PTR record varies
// based on whether it is for an IPv4 or IPv6 address. For an IPv4 address the ''host'' property for
// every PTR record is the last octet of the IP address that the PTR record belongs to, while the
// ''data'' property is the canonical name of the host that the reverse lookup resolves to. Every PTR
// record belongs to a domain on the SoftLayer name servers named by the first three octets of an IP
// address in reverse order followed by ".in-addr.arpa". For instance, if the reverse DNS record for
// 10.0.0.1 is "host.example.org" then it's corresponding SoftLayer_Dns_Domain_ResourceRecord_PtrType
// host is "1", while it's data property equals "host.example.org". The full name of the reverse record
// for host.example.org including the domain name is "1.0.0.10.in-addr.arpa". For an IPv6 address the
// ''host'' property for every PTR record is the last four octets of the IP address that the PTR record
// belongs to. The last four octets need to be in reversed order and each digit separated by a period.
// The ''data'' property is the canonical name of the host that the reverse lookup resolves to. Every
// PTR record belongs to a domain on the SoftLayer name servers named by the first four octets of an IP
// address in reverse order, split up by digit with a period, and followed by ".ip6.arpa". For
// instance, if the reverse DNS record for fe80:0000:0000:0000:0000:0000:0a00:0001 is
// "host.example.org" then it's corresponding SoftLayer_Dns_Domain_ResourceRecord_PtrType host is
// "1.0.0.0.0.0.a.0.0.0.0.0.0.0.0.0", while it's data property equals "host.example.org". The full name
// of the reverse record for host.example.org including the domain name is
// "1.0.0.0.0.0.a.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.8.e.f.ip6.arpa". PTR record host names
// may not be changed by [[SoftLayer_Dns_Domain_ResourceRecord::editObject]] or
// [[SoftLayer_Dns_Domain_ResourceRecord::editObjects]].
type SoftLayer_Dns_Domain_ResourceRecord_PtrType struct {

	// Expire - The amount of time in seconds that a secondary name server (or servers) will hold a zone
	// before it is no longer considered authoritative.
	Expire int `json:"expire,omitempty"`

	// DomainId - An identifier belonging to the domain that a resource record is associated with.
	DomainId int `json:"domainId,omitempty"`

	// Refresh - The amount of time in seconds that a secondary name server should wait to check for a new
	// copy of a DNS zone from the domain's primary name server. If a zone file has changed then the
	// secondary DNS server will update it's copy of the zone to match the primary DNS server's zone.
	Refresh int `json:"refresh,omitempty"`

	// ResponsiblePerson - The email address of the person responsible for a domain, with the "@" replaced
	// with a For instance, if root@example.org is responsible for example.org, then example.org's SOA
	// responsibility is "root.example.org.".
	ResponsiblePerson string `json:"responsiblePerson,omitempty"`

	// Retry - The amount of time in seconds that a domain's primary name server (or servers) should wait
	// if an attempt to refresh by a secondary name server failed before attempting to refresh a domain's
	// zone with that secondary name server again.
	Retry int `json:"retry,omitempty"`

	// Ttl - The Time To Live value of a resource record, measured in seconds. TTL is used by a name server
	// to determine how long to cache a resource record. An SOA record's TTL value defines the domain's
	// overall
	Ttl int `json:"ttl,omitempty"`

	// Host - The host defined by a resource record. A value of "@" denotes a wildcard.
	Host string `json:"host,omitempty"`

	// Minimum - The amount of time in seconds that a domain's resource records are valid. This is also
	// known as a minimum and can be overridden by an individual resource record's
	Minimum int `json:"minimum,omitempty"`

	// Type - A domain resource record's type. A value of "a" denotes an A (address) record, "aaaa" denotes
	// an (IPv6 address) record, "cname" denotes a (canonical name) record, "mx" denotes an MX (mail
	// exchanger) record, "ns" denotes an NS (nameserver) record, "ptr" denotes a PTR (pointer/reverse)
	// record, "soa" denotes the SOA (start of authority) record, "spf" denotes a SPF (sender policy
	// framework) record, and "txt" denotes a TXT (text) record. A domain record's type also denotes which
	// class in the SoftLayer API is a best match for extending a resource record.
	Type string `json:"type,omitempty"`

	// IsGatewayAddress - Whether the address associated with a PTR record is the gateway address of a
	// subnet.
	IsGatewayAddress bool `json:"isGatewayAddress,omitempty"`

	// Data - The value of a domain's resource record. This can be an IP address or a hostname. Fully
	// qualified host and domain name data must end with the "." character.
	Data string `json:"data,omitempty"`

	// MxPriority - Useful in cases where a domain has more than one mail exchanger, the priority property
	// is the priority of the MTA that delivers mail for a domain. A lower number denotes a higher
	// priority, and mail will attempt to deliver through that MTA before moving to lower priority mail
	// servers. Priority is defaulted to 10 upon resource record creation.
	MxPriority int `json:"mxPriority,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`
}

func (softlayer_dns_domain_resourcerecord_ptrtype *SoftLayer_Dns_Domain_ResourceRecord_PtrType) String() string {
	return "SoftLayer_Dns_Domain_ResourceRecord_PtrType"
}

// SoftLayer_Dns_Domain_ResourceRecord_PtrType_Extended is SoftLayer_Dns_Domain_ResourceRecord_PtrType with all maskable types.
type SoftLayer_Dns_Domain_ResourceRecord_PtrType_Extended struct {
	SoftLayer_Dns_Domain_ResourceRecord_PtrType

	// Domain - no documentation
	Domain *SoftLayer_Dns_Domain `json:"domain,omitempty"`
}

func (softlayer_dns_domain_resourcerecord_ptrtype *SoftLayer_Dns_Domain_ResourceRecord_PtrType_Extended) String() string {
	return "SoftLayer_Dns_Domain_ResourceRecord_PtrType"
}

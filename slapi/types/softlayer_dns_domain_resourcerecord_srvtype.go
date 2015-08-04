package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Dns_Domain_ResourceRecord_SrvType - SoftLayer_Dns_Domain_ResourceRecord_SrvType is a
// SoftLayer_Dns_Domain_ResourceRecord object whose ''type'' property is set to "srv" and defines a DNS
// SRV record on a SoftLayer hosted domain.
type SoftLayer_Dns_Domain_ResourceRecord_SrvType struct {

	// DomainId - An identifier belonging to the domain that a resource record is associated with.
	DomainId int `json:"domainId,omitempty"`

	// MxPriority - Useful in cases where a domain has more than one mail exchanger, the priority property
	// is the priority of the MTA that delivers mail for a domain. A lower number denotes a higher
	// priority, and mail will attempt to deliver through that MTA before moving to lower priority mail
	// servers. Priority is defaulted to 10 upon resource record creation.
	MxPriority int `json:"mxPriority,omitempty"`

	// Protocol - The protocol of the desired service; this is usually either TCP or
	Protocol string `json:"protocol,omitempty"`

	// Service - no documentation
	Service string `json:"service,omitempty"`

	// Minimum - The amount of time in seconds that a domain's resource records are valid. This is also
	// known as a minimum and can be overridden by an individual resource record's
	Minimum int `json:"minimum,omitempty"`

	// Host - The host defined by a resource record. A value of "@" denotes a wildcard.
	Host string `json:"host,omitempty"`

	// ResponsiblePerson - The email address of the person responsible for a domain, with the "@" replaced
	// with a For instance, if root@example.org is responsible for example.org, then example.org's SOA
	// responsibility is "root.example.org.".
	ResponsiblePerson string `json:"responsiblePerson,omitempty"`

	// Port - The TCP or UDP port on which the service is to be found.
	Port int `json:"port,omitempty"`

	// Data - The value of a domain's resource record. This can be an IP address or a hostname. Fully
	// qualified host and domain name data must end with the "." character.
	Data string `json:"data,omitempty"`

	// Expire - The amount of time in seconds that a secondary name server (or servers) will hold a zone
	// before it is no longer considered authoritative.
	Expire int `json:"expire,omitempty"`

	// Retry - The amount of time in seconds that a domain's primary name server (or servers) should wait
	// if an attempt to refresh by a secondary name server failed before attempting to refresh a domain's
	// zone with that secondary name server again.
	Retry int `json:"retry,omitempty"`

	// Type - A domain resource record's type. A value of "a" denotes an A (address) record, "aaaa" denotes
	// an (IPv6 address) record, "cname" denotes a (canonical name) record, "mx" denotes an MX (mail
	// exchanger) record, "ns" denotes an NS (nameserver) record, "ptr" denotes a PTR (pointer/reverse)
	// record, "soa" denotes the SOA (start of authority) record, "spf" denotes a SPF (sender policy
	// framework) record, and "txt" denotes a TXT (text) record. A domain record's type also denotes which
	// class in the SoftLayer API is a best match for extending a resource record.
	Type string `json:"type,omitempty"`

	// Priority - The priority of the target host, lower value means more preferred.
	Priority int `json:"priority,omitempty"`

	// Weight - A relative weight for records with the same priority.
	Weight int `json:"weight,omitempty"`

	// Refresh - The amount of time in seconds that a secondary name server should wait to check for a new
	// copy of a DNS zone from the domain's primary name server. If a zone file has changed then the
	// secondary DNS server will update it's copy of the zone to match the primary DNS server's zone.
	Refresh int `json:"refresh,omitempty"`

	// Ttl - The Time To Live value of a resource record, measured in seconds. TTL is used by a name server
	// to determine how long to cache a resource record. An SOA record's TTL value defines the domain's
	// overall
	Ttl int `json:"ttl,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Domain - no documentation
	Domain *SoftLayer_Dns_Domain `json:"domain,omitempty"`
}

func (softlayer_dns_domain_resourcerecord_srvtype *SoftLayer_Dns_Domain_ResourceRecord_SrvType) String() string {
	return "SoftLayer_Dns_Domain_ResourceRecord_SrvType"
}

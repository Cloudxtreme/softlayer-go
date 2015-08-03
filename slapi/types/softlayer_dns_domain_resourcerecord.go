package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Dns_Domain_ResourceRecord - The SoftLayer_Dns_Domain_ResourceRecord data type represents a
// single resource record entry in a SoftLayer hosted domain. Each resource record contains a ''host''
// and ''data'' property, defining a resource's name and it's target data. Domains contain multiple
// types of resource records. The ''type'' property separates out resource records by type. ''Type''
// can take one of the following values: * '''"a"''' for
// [[SoftLayer_Dns_Domain_ResourceRecord_AType|address]] records * '''"aaaa"''' for
// [[SoftLayer_Dns_Domain_ResourceRecord_AaaaType|address]] records * '''"cname"''' for
// [[SoftLayer_Dns_Domain_ResourceRecord_CnameType|canonical name]] records * '''"mx"''' for
// [[SoftLayer_Dns_Domain_ResourceRecord_MxType|mail exchanger]] records * '''"ns"''' for
// [[SoftLayer_Dns_Domain_ResourceRecord_NsType|name server]] records * '''"ptr"''' for
// [[SoftLayer_Dns_Domain_ResourceRecord_PtrType|pointer]] records in reverse domains * '''"soa"''' for
// a domain's [[SoftLayer_Dns_Domain_ResourceRecord_SoaType|start of authority]] record * '''"spf"'''
// for [[SoftLayer_Dns_Domain_ResourceRecord_SpfType|sender policy framework]] records * '''"srv"'''
// for [[SoftLayer_Dns_Domain_ResourceRecord_SrvType|service]] records * '''"txt"''' for
// [[SoftLayer_Dns_Domain_ResourceRecord_TxtType|text]] records As
// ''SoftLayer_Dns_Domain_ResourceRecord'' objects are created and loaded, the API verifies the
// ''type'' property and casts the object as the appropriate type.
type SoftLayer_Dns_Domain_ResourceRecord struct {

	// DomainId - An identifier belonging to the domain that a resource record is associated with.
	DomainId int `json:"domainId"`

	// Expire - The amount of time in seconds that a secondary name server (or servers) will hold a zone
	// before it is no longer considered authoritative.
	Expire int `json:"expire"`

	// Host - The host defined by a resource record. A value of "@" denotes a wildcard.
	Host string `json:"host"`

	// ResponsiblePerson - The email address of the person responsible for a domain, with the "@" replaced
	// with a For instance, if root@example.org is responsible for example.org, then example.org's SOA
	// responsibility is "root.example.org.".
	ResponsiblePerson string `json:"responsiblePerson"`

	// Type - A domain resource record's type. A value of "a" denotes an A (address) record, "aaaa" denotes
	// an (IPv6 address) record, "cname" denotes a (canonical name) record, "mx" denotes an MX (mail
	// exchanger) record, "ns" denotes an NS (nameserver) record, "ptr" denotes a PTR (pointer/reverse)
	// record, "soa" denotes the SOA (start of authority) record, "spf" denotes a SPF (sender policy
	// framework) record, and "txt" denotes a TXT (text) record. A domain record's type also denotes which
	// class in the SoftLayer API is a best match for extending a resource record.
	Type string `json:"type"`

	// Refresh - The amount of time in seconds that a secondary name server should wait to check for a new
	// copy of a DNS zone from the domain's primary name server. If a zone file has changed then the
	// secondary DNS server will update it's copy of the zone to match the primary DNS server's zone.
	Refresh int `json:"refresh"`

	// Retry - The amount of time in seconds that a domain's primary name server (or servers) should wait
	// if an attempt to refresh by a secondary name server failed before attempting to refresh a domain's
	// zone with that secondary name server again.
	Retry int `json:"retry"`

	// Ttl - The Time To Live value of a resource record, measured in seconds. TTL is used by a name server
	// to determine how long to cache a resource record. An SOA record's TTL value defines the domain's
	// overall
	Ttl int `json:"ttl"`

	// Data - The value of a domain's resource record. This can be an IP address or a hostname. Fully
	// qualified host and domain name data must end with the "." character.
	Data string `json:"data"`

	// Id - no documentation
	Id int `json:"id"`

	// Minimum - The amount of time in seconds that a domain's resource records are valid. This is also
	// known as a minimum and can be overridden by an individual resource record's
	Minimum int `json:"minimum"`

	// MxPriority - Useful in cases where a domain has more than one mail exchanger, the priority property
	// is the priority of the MTA that delivers mail for a domain. A lower number denotes a higher
	// priority, and mail will attempt to deliver through that MTA before moving to lower priority mail
	// servers. Priority is defaulted to 10 upon resource record creation.
	MxPriority int `json:"mxPriority"`
}

func (softlayer_dns_domain_resourcerecord *SoftLayer_Dns_Domain_ResourceRecord) String() string {
	return "SoftLayer_Dns_Domain_ResourceRecord"
}

// SoftLayer_Dns_Domain_ResourceRecord_Extended is SoftLayer_Dns_Domain_ResourceRecord with all maskable types.
type SoftLayer_Dns_Domain_ResourceRecord_Extended struct {
	SoftLayer_Dns_Domain_ResourceRecord

	// Domain - no documentation
	Domain *SoftLayer_Dns_Domain `json:"domain"`
}

func (softlayer_dns_domain_resourcerecord *SoftLayer_Dns_Domain_ResourceRecord_Extended) String() string {
	return "SoftLayer_Dns_Domain_ResourceRecord"
}

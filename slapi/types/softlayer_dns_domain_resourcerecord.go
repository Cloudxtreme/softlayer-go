package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

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

	// Data - The value of a domain's resource record. This can be an IP address or a hostname. Fully
	// qualified host and domain name data must end with the "." character.
	Data string `json:"data"`

	// Domain - no documentation
	Domain *SoftLayer_Dns_Domain `json:"domain"`

	// DomainId - An identifier belonging to the domain that a resource record is associated with.
	DomainId int `json:"domainId"`

	// Expire - The amount of time in seconds that a secondary name server (or servers) will hold a zone
	// before it is no longer considered authoritative.
	Expire int `json:"expire"`

	// Host - The host defined by a resource record. A value of "@" denotes a wildcard.
	Host string `json:"host"`

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

	// Refresh - The amount of time in seconds that a secondary name server should wait to check for a new
	// copy of a DNS zone from the domain's primary name server. If a zone file has changed then the
	// secondary DNS server will update it's copy of the zone to match the primary DNS server's zone.
	Refresh int `json:"refresh"`

	// ResponsiblePerson - The email address of the person responsible for a domain, with the "@" replaced
	// with a For instance, if root@example.org is responsible for example.org, then example.org's SOA
	// responsibility is "root.example.org.".
	ResponsiblePerson string `json:"responsiblePerson"`

	// Retry - The amount of time in seconds that a domain's primary name server (or servers) should wait
	// if an attempt to refresh by a secondary name server failed before attempting to refresh a domain's
	// zone with that secondary name server again.
	Retry int `json:"retry"`

	// Ttl - The Time To Live value of a resource record, measured in seconds. TTL is used by a name server
	// to determine how long to cache a resource record. An SOA record's TTL value defines the domain's
	// overall
	Ttl int `json:"ttl"`

	// Type - A domain resource record's type. A value of "a" denotes an A (address) record, "aaaa" denotes
	// an (IPv6 address) record, "cname" denotes a (canonical name) record, "mx" denotes an MX (mail
	// exchanger) record, "ns" denotes an NS (nameserver) record, "ptr" denotes a PTR (pointer/reverse)
	// record, "soa" denotes the SOA (start of authority) record, "spf" denotes a SPF (sender policy
	// framework) record, and "txt" denotes a TXT (text) record. A domain record's type also denotes which
	// class in the SoftLayer API is a best match for extending a resource record.
	Type string `json:"type"`
}

// CreateObject - createObject creates a new domain resource record. The ''host'' property of the
// templateObject parameter is scrubbed to remove all non-alpha numeric characters except for and The
// ''data'' property of the templateObject parameter is scrubbed to remove all non-alphanumeric
// characters for "." and Creating a resource record updates the serial number of the domain the
// resource record is associated with. ''createObject'' returns Boolean ''true'' on successful create
// or ''false'' if it was unable to create a resource record.
func (softlayer_dns_domain_resourcerecord *SoftLayer_Dns_Domain_ResourceRecord) CreateObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Dns_Domain_ResourceRecord) (*SoftLayer_Dns_Domain_ResourceRecord, error) {
	var returnValue *SoftLayer_Dns_Domain_ResourceRecord
	return returnValue, nil
}

// CreateObjects - Create multiple resource records on a domain. This follows the same logic as
// ''createObject'. The serial number of the domain associated with this resource record is updated
// upon creation. ''createObjects'' returns Boolean ''true'' on successful creation or ''false'' if it
// was unable to create a resource record.
func (softlayer_dns_domain_resourcerecord *SoftLayer_Dns_Domain_ResourceRecord) CreateObjects(commonOptions *slapi.CommonOptions, templateObjects []SoftLayer_Dns_Domain_ResourceRecord) ([]*SoftLayer_Dns_Domain_ResourceRecord, error) {
	var returnValue []*SoftLayer_Dns_Domain_ResourceRecord
	return returnValue, nil
}

// DeleteObject - Delete a domain's resource record. '''This cannot be undone.''' Be wary of running
// this method. If you remove a resource record in error you will need to re-create it by creating a
// new SoftLayer_Dns_Domain_ResourceRecord object. The serial number of the domain associated with this
// resource record is updated upon deletion. You may not delete NS, or PTR resource records.
// ''deleteObject'' returns Boolean ''true'' on successful deletion or ''false'' if it was unable to
// remove a resource record.
func (softlayer_dns_domain_resourcerecord *SoftLayer_Dns_Domain_ResourceRecord) DeleteObject(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteObjects - Remove multiple resource records from a domain. This follows the same logic as
// ''deleteObject'' and '''cannot be undone'''. The serial number of the domain associated with this
// resource record is updated upon deletion. You may not delete SOA records, PTR records, or NS
// resource records that point to ns1.softlayer.com or ns2.softlayer.com. ''deleteObjects'' returns
// Boolean ''true'' on successful deletion or ''false'' if it was unable to remove a resource record.
func (softlayer_dns_domain_resourcerecord *SoftLayer_Dns_Domain_ResourceRecord) DeleteObjects(commonOptions *slapi.CommonOptions, templateObjects []SoftLayer_Dns_Domain_ResourceRecord) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - editObject edits an existing domain resource record. The ''host'' property of the
// templateObject parameter is scrubbed to remove all non-alpha numeric characters except for and The
// ''data'' property of the templateObject parameter is scrubbed to remove all non-alphanumeric
// characters for "." and Editing a resource record updates the serial number of the domain the
// resource record is associated with. ''editObject'' returns Boolean ''true'' on a successful edit or
// ''false'' if it was unable to edit the resource record.
func (softlayer_dns_domain_resourcerecord *SoftLayer_Dns_Domain_ResourceRecord) EditObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Dns_Domain_ResourceRecord) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObjects - Edit multiple resource records on a domain. This follows the same logic as
// ''createObject'. The serial number of the domain associated with this resource record is updated
// upon creation. ''createObjects'' returns Boolean ''true'' on successful creation or ''false'' if it
// was unable to create a resource record.
func (softlayer_dns_domain_resourcerecord *SoftLayer_Dns_Domain_ResourceRecord) EditObjects(commonOptions *slapi.CommonOptions, templateObjects []SoftLayer_Dns_Domain_ResourceRecord) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Dns_Domain_ResourceRecord object whose ID number
// corresponds to the ID number of the init parameter passed to the SoftLayer_Dns_Domain_ResourceRecord
// service. You can only retrieve resource records belonging to domains that are assigned to your
// SoftLayer account.
func (softlayer_dns_domain_resourcerecord *SoftLayer_Dns_Domain_ResourceRecord) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Dns_Domain_ResourceRecord, error) {
	var returnValue *SoftLayer_Dns_Domain_ResourceRecord
	return returnValue, nil
}

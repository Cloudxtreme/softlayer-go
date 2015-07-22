package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Dns_Domain - The SoftLayer_Dns_Domain data type represents a single DNS domain record
// hosted on the SoftLayer nameservers. Domains contain general information about the domain name such
// as name and serial. Individual records such as A, and MX records are stored in the domain's
// associated [[SoftLayer_Dns_Domain_ResourceRecord (type)|SoftLayer_Dns_Domain_ResourceRecord]]
// records.
type SoftLayer_Dns_Domain struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// Id - no documentation
	Id int `json:"id"`

	// ManagedResourceFlag - A flag indicating that the dns domain record is a managed resource.
	ManagedResourceFlag bool `json:"managedResourceFlag"`

	// Name - A domain's name including top-level domain, for example "example.com".
	Name string `json:"name"`

	// ResourceRecordCount - A count of the individual records contained within a domain record. These
	// include but are not limited to A, MX, SPF and TXT records.
	ResourceRecordCount uint64 `json:"resourceRecordCount"`

	// ResourceRecords - The individual records contained within a domain record. These include but are not
	// limited to A, MX, SPF and TXT records.
	ResourceRecords []*SoftLayer_Dns_Domain_ResourceRecord `json:"resourceRecords"`

	// Secondary - The secondary DNS record that defines this domain as being managed through zone
	// transfers.
	Secondary *SoftLayer_Dns_Secondary `json:"secondary"`

	// Serial - A unique number denoting the latest revision of a domain. Whenever a domain is changed its
	// corresponding serial number is also changed. Serial numbers typically follow the format yyyymmdd##
	// where yyyy is the current year, mm is the current month, dd is the current day of the month, and ##
	// is the number of the revision for that day. A domain's serial number is automatically updated when
	// edited via the
	Serial int `json:"serial"`

	// UpdateDate - no documentation
	UpdateDate *time.Time `json:"updateDate"`
}

// CreateARecord - Create an A record on a SoftLayer domain. This is a shortcut method, meant to take
// the work out of creating a SoftLayer_Dns_Domain_ResourceRecord if you already have a domain record
// available. createARecord returns the newly created SoftLayer_Dns_Domain_ResourceRecord_AType.
func (softlayer_dns_domain *SoftLayer_Dns_Domain) CreateARecord(host string, data string, ttl int) (*SoftLayer_Dns_Domain_ResourceRecord_AType, error) {
	var returnValue *SoftLayer_Dns_Domain_ResourceRecord_AType
	return returnValue, nil
}

// CreateAaaaRecord - Create an record on a SoftLayer domain. This is a shortcut method, meant to take
// the work out of creating a SoftLayer_Dns_Domain_ResourceRecord if you already have a domain record
// available. createARecord returns the newly created SoftLayer_Dns_Domain_ResourceRecord_AaaaType.
func (softlayer_dns_domain *SoftLayer_Dns_Domain) CreateAaaaRecord(host string, data string, ttl int) (*SoftLayer_Dns_Domain_ResourceRecord_AaaaType, error) {
	var returnValue *SoftLayer_Dns_Domain_ResourceRecord_AaaaType
	return returnValue, nil
}

// CreateCnameRecord - Create a record on a SoftLayer domain. This is a shortcut method, meant to take
// the work out of creating a SoftLayer_Dns_Domain_ResourceRecord if you already have a domain record
// available. createCnameRecord returns the newly created
// SoftLayer_Dns_Domain_ResourceRecord_CnameType.
func (softlayer_dns_domain *SoftLayer_Dns_Domain) CreateCnameRecord(host string, data string, ttl int) (*SoftLayer_Dns_Domain_ResourceRecord_CnameType, error) {
	var returnValue *SoftLayer_Dns_Domain_ResourceRecord_CnameType
	return returnValue, nil
}

// CreateMxRecord - Create an MX record on a SoftLayer domain. This is a shortcut method, meant to take
// the work out of creating a SoftLayer_Dns_Domain_ResourceRecord if you already have a domain record
// available. MX records are created with a default priority of 10. createMxRecord returns the newly
// created SoftLayer_Dns_Domain_ResourceRecord_MxType.
func (softlayer_dns_domain *SoftLayer_Dns_Domain) CreateMxRecord(host string, data string, ttl int, mxPriority int) (*SoftLayer_Dns_Domain_ResourceRecord_MxType, error) {
	var returnValue *SoftLayer_Dns_Domain_ResourceRecord_MxType
	return returnValue, nil
}

// CreateNsRecord - Create an NS record on a SoftLayer domain. This is a shortcut method, meant to take
// the work out of creating a SoftLayer_Dns_Domain_ResourceRecord if you already have a domain record
// available. createNsRecord returns the newly created SoftLayer_Dns_Domain_ResourceRecord_NsType.
func (softlayer_dns_domain *SoftLayer_Dns_Domain) CreateNsRecord(host string, data string, ttl int) (*SoftLayer_Dns_Domain_ResourceRecord_NsType, error) {
	var returnValue *SoftLayer_Dns_Domain_ResourceRecord_NsType
	return returnValue, nil
}

// CreateObject - Create a new domain on the SoftLayer name servers. The SoftLayer_Dns_Domain object
// passed to this function must have at least one A or resource record. createObject creates a default
// SOA record with the data: * '''host''': "@" * '''data''': "ns1.softlayer.com." * '''responsible
// person''': "root.[your domain name]." * '''expire''': 604800 seconds * '''refresh''': 3600 seconds *
// '''retry''': 300 seconds * '''minimum''': 3600 seconds If your new domain uses the .de top-level
// domain then SOA refresh is set to 10000 seconds, retry is set to 1800 seconds, and minimum to 10000
// seconds. If your domain doesn't contain NS resource records for ns1.softlayer.com or
// ns2.softlayer.com then ''createObject'' will create them for you. ''createObject'' returns a Boolean
// ''true'' on successful object creation or ''false'' if your domain was unable to be created..
func (softlayer_dns_domain *SoftLayer_Dns_Domain) CreateObject(templateObject SoftLayer_Dns_Domain) (*SoftLayer_Dns_Domain, error) {
	var returnValue *SoftLayer_Dns_Domain
	return returnValue, nil
}

// CreateObjects - Create multiple domains on the SoftLayer name servers. Each domain record passed to
// ''createObjects'' follows the logic in the SoftLayer_Dns_Domain ''createObject'' method.
func (softlayer_dns_domain *SoftLayer_Dns_Domain) CreateObjects(templateObjects []SoftLayer_Dns_Domain) ([]*SoftLayer_Dns_Domain, error) {
	var returnValue []*SoftLayer_Dns_Domain
	return returnValue, nil
}

// CreatePtrRecord - setPtrRecordForIpAddress() sets a single reverse DNS record for a single IP
// address and returns the newly created or edited [[SoftLayer_Dns_Domain_ResourceRecord]] record.
// Currently this method only supports IPv4 addresses and performs no operation when given an IPv6
// address.
func (softlayer_dns_domain *SoftLayer_Dns_Domain) CreatePtrRecord(ipAddress string, ptrRecord string, ttl int) (*SoftLayer_Dns_Domain_ResourceRecord, error) {
	var returnValue *SoftLayer_Dns_Domain_ResourceRecord
	return returnValue, nil
}

// CreateSpfRecord - Create an SPF record on a SoftLayer domain. This is a shortcut method, meant to
// take the work out of creating a SoftLayer_Dns_Domain_ResourceRecord if you already have a domain
// record available. createARecord returns the newly created
// SoftLayer_Dns_Domain_ResourceRecord_SpfType.
func (softlayer_dns_domain *SoftLayer_Dns_Domain) CreateSpfRecord(host string, data string, ttl int) (*SoftLayer_Dns_Domain_ResourceRecord_SpfType, error) {
	var returnValue *SoftLayer_Dns_Domain_ResourceRecord_SpfType
	return returnValue, nil
}

// CreateTxtRecord - Create a TXT record on a SoftLayer domain. This is a shortcut method, meant to
// take the work out of creating a SoftLayer_Dns_Domain_ResourceRecord if you already have a domain
// record available. createARecord returns the newly created
// SoftLayer_Dns_Domain_ResourceRecord_TxtType.
func (softlayer_dns_domain *SoftLayer_Dns_Domain) CreateTxtRecord(host string, data string, ttl int) (*SoftLayer_Dns_Domain_ResourceRecord_TxtType, error) {
	var returnValue *SoftLayer_Dns_Domain_ResourceRecord_TxtType
	return returnValue, nil
}

// DeleteObject - deleteObject permanently removes a domain and all of it's associated resource records
// from the softlayer name servers. '''This cannot be undone.''' Be wary of running this method. If you
// remove a domain in error you will need to re-create it by creating a new SoftLayer_Dns_Domain
// object.
func (softlayer_dns_domain *SoftLayer_Dns_Domain) DeleteObject() (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetByDomainName - Search for [[SoftLayer_Dns_Domain]] records by domain name. getByDomainName()
// performs an inclusive search for domain records, returning multiple records based on partial name
// matches. Use this method to locate domain records if you don't have access to their id numbers.
func (softlayer_dns_domain *SoftLayer_Dns_Domain) GetByDomainName(name string) ([]*SoftLayer_Dns_Domain, error) {
	var returnValue []*SoftLayer_Dns_Domain
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Dns_Domain object whose ID number corresponds to the
// ID number of the init parameter passed to the SoftLayer_Dns_Domain service. You can only retrieve
// domains that are assigned to your SoftLayer account.
func (softlayer_dns_domain *SoftLayer_Dns_Domain) GetObject() (*SoftLayer_Dns_Domain, error) {
	var returnValue *SoftLayer_Dns_Domain
	return returnValue, nil
}

// GetZoneFileContents - Return a SoftLayer hosted domain and resource records' data formatted as zone
// file.
func (softlayer_dns_domain *SoftLayer_Dns_Domain) GetZoneFileContents() (string, error) {
	var returnValue string
	return returnValue, nil
}

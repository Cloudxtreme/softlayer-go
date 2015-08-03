package types

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

func (softlayer_dns_domain *SoftLayer_Dns_Domain) String() string {
	return "SoftLayer_Dns_Domain"
}

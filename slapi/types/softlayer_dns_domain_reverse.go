package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Dns_Domain_Reverse - The SoftLayer_Dns_Domain_Reverse data type represents a reverse IP
// address record.
type SoftLayer_Dns_Domain_Reverse struct {

	// Name - A domain's name including top-level domain, for example "example.com".
	Name string `json:"name,omitempty"`

	// Serial - A unique number denoting the latest revision of a domain. Whenever a domain is changed its
	// corresponding serial number is also changed. Serial numbers typically follow the format yyyymmdd##
	// where yyyy is the current year, mm is the current month, dd is the current day of the month, and ##
	// is the number of the revision for that day. A domain's serial number is automatically updated when
	// edited via the
	Serial int `json:"serial,omitempty"`

	// NetworkAddress - no documentation
	NetworkAddress string `json:"networkAddress,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// UpdateDate - no documentation
	UpdateDate *time.Time `json:"updateDate,omitempty"`
}

func (softlayer_dns_domain_reverse *SoftLayer_Dns_Domain_Reverse) String() string {
	return "SoftLayer_Dns_Domain_Reverse"
}

// SoftLayer_Dns_Domain_Reverse_Extended is SoftLayer_Dns_Domain_Reverse with all maskable types.
type SoftLayer_Dns_Domain_Reverse_Extended struct {
	SoftLayer_Dns_Domain_Reverse

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// ManagedResourceFlag - A flag indicating that the dns domain record is a managed resource.
	ManagedResourceFlag bool `json:"managedResourceFlag,omitempty"`

	// ResourceRecords - The individual records contained within a domain record. These include but are not
	// limited to A, MX, SPF and TXT records.
	ResourceRecords []*SoftLayer_Dns_Domain_ResourceRecord `json:"resourceRecords,omitempty"`

	// Secondary - The secondary DNS record that defines this domain as being managed through zone
	// transfers.
	Secondary *SoftLayer_Dns_Secondary `json:"secondary,omitempty"`

	// ResourceRecordCount - A count of the individual records contained within a domain record. These
	// include but are not limited to A, MX, SPF and TXT records.
	ResourceRecordCount uint64 `json:"resourceRecordCount,omitempty"`
}

func (softlayer_dns_domain_reverse *SoftLayer_Dns_Domain_Reverse_Extended) String() string {
	return "SoftLayer_Dns_Domain_Reverse"
}

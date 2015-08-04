package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Dns_Domain_Registration_Information - no documentation
type SoftLayer_Container_Dns_Domain_Registration_Information struct {

	// ExpireDate - no documentation
	ExpireDate *time.Time `json:"expireDate,omitempty"`

	// Nameservers - no documentation
	Nameservers []*SoftLayer_Container_Dns_Domain_Registration_Nameserver `json:"nameservers,omitempty"`

	// RegistryCreateDate - <nil>
	RegistryCreateDate *time.Time `json:"registryCreateDate,omitempty"`

	// RegistryExpireDate - <nil>
	RegistryExpireDate *time.Time `json:"registryExpireDate,omitempty"`

	// RegistryUpdateDate - <nil>
	RegistryUpdateDate *time.Time `json:"registryUpdateDate,omitempty"`

	// Contacts - no documentation
	Contacts []*SoftLayer_Container_Dns_Domain_Registration_Contact `json:"contacts,omitempty"`
}

func (softlayer_container_dns_domain_registration_information *SoftLayer_Container_Dns_Domain_Registration_Information) String() string {
	return "SoftLayer_Container_Dns_Domain_Registration_Information"
}

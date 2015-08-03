package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Dns_Domain_Registration_Information - no documentation
type SoftLayer_Container_Dns_Domain_Registration_Information struct {

	// RegistryExpireDate - <nil>
	RegistryExpireDate *time.Time `json:"registryExpireDate"`

	// RegistryUpdateDate - <nil>
	RegistryUpdateDate *time.Time `json:"registryUpdateDate"`

	// Contacts - no documentation
	Contacts []*SoftLayer_Container_Dns_Domain_Registration_Contact `json:"contacts"`

	// ExpireDate - no documentation
	ExpireDate *time.Time `json:"expireDate"`

	// Nameservers - no documentation
	Nameservers []*SoftLayer_Container_Dns_Domain_Registration_Nameserver `json:"nameservers"`

	// RegistryCreateDate - <nil>
	RegistryCreateDate *time.Time `json:"registryCreateDate"`
}

func (softlayer_container_dns_domain_registration_information *SoftLayer_Container_Dns_Domain_Registration_Information) String() string {
	return "SoftLayer_Container_Dns_Domain_Registration_Information"
}

package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Dns_Domain_Registration - The SoftLayer_Dns_Domain_Registration data type represents a
// domain registration record.
type SoftLayer_Dns_Domain_Registration struct {

	// ExpireDate - no documentation
	ExpireDate *time.Time `json:"expireDate"`

	// Id - no documentation
	Id int `json:"id"`

	// LockedFlag - no documentation
	LockedFlag int `json:"lockedFlag"`

	// Name - no documentation
	Name string `json:"name"`

	// DomainRegistrationStatusId - <nil>
	DomainRegistrationStatusId int `json:"domainRegistrationStatusId"`

	// ServiceProviderId - <nil>
	ServiceProviderId int `json:"serviceProviderId"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// RegistrantVerificationStatusId - <nil>
	RegistrantVerificationStatusId int `json:"registrantVerificationStatusId"`
}

func (softlayer_dns_domain_registration *SoftLayer_Dns_Domain_Registration) String() string {
	return "SoftLayer_Dns_Domain_Registration"
}

// SoftLayer_Dns_Domain_Registration_Extended is SoftLayer_Dns_Domain_Registration with all maskable types.
type SoftLayer_Dns_Domain_Registration_Extended struct {
	SoftLayer_Dns_Domain_Registration

	// Account - The SoftLayer customer account that the domain is registered to.
	Account *SoftLayer_Account `json:"account"`

	// DomainRegistrationStatus - no documentation
	DomainRegistrationStatus *SoftLayer_Dns_Domain_Registration_Status `json:"domainRegistrationStatus"`

	// ServiceProvider - <nil>
	ServiceProvider *SoftLayer_Service_Provider `json:"serviceProvider"`

	// RegistrantVerificationStatus - no documentation
	RegistrantVerificationStatus *SoftLayer_Dns_Domain_Registration_Registrant_Verification_Status `json:"registrantVerificationStatus"`
}

func (softlayer_dns_domain_registration *SoftLayer_Dns_Domain_Registration_Extended) String() string {
	return "SoftLayer_Dns_Domain_Registration"
}

package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Dns_Domain_Registration - The SoftLayer_Dns_Domain_Registration data type represents a
// domain registration record.
type SoftLayer_Dns_Domain_Registration struct {

	// Account - The SoftLayer customer account that the domain is registered to.
	Account *SoftLayer_Account `json:"account"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// DomainRegistrationStatus - no documentation
	DomainRegistrationStatus *SoftLayer_Dns_Domain_Registration_Status `json:"domainRegistrationStatus"`

	// DomainRegistrationStatusId - <nil>
	DomainRegistrationStatusId int `json:"domainRegistrationStatusId"`

	// ExpireDate - no documentation
	ExpireDate *time.Time `json:"expireDate"`

	// Id - no documentation
	Id int `json:"id"`

	// LockedFlag - no documentation
	LockedFlag int `json:"lockedFlag"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// RegistrantVerificationStatus - no documentation
	RegistrantVerificationStatus *SoftLayer_Dns_Domain_Registration_Registrant_Verification_Status `json:"registrantVerificationStatus"`

	// RegistrantVerificationStatusId - <nil>
	RegistrantVerificationStatusId int `json:"registrantVerificationStatusId"`

	// ServiceProvider - <nil>
	ServiceProvider *SoftLayer_Service_Provider `json:"serviceProvider"`

	// ServiceProviderId - <nil>
	ServiceProviderId int `json:"serviceProviderId"`
}

func (softlayer_dns_domain_registration *SoftLayer_Dns_Domain_Registration) String() string {
	return "SoftLayer_Dns_Domain_Registration"
}

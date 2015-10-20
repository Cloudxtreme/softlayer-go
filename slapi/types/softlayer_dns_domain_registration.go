package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Dns_Domain_Registration - The SoftLayer_Dns_Domain_Registration data type represents a
// domain registration record.
type SoftLayer_Dns_Domain_Registration struct {

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// RegistrantVerificationStatusId - <nil>
	RegistrantVerificationStatusId int `json:"registrantVerificationStatusId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// DomainRegistrationStatusId - <nil>
	DomainRegistrationStatusId int `json:"domainRegistrationStatusId,omitempty"`

	// ExpireDate - no documentation
	ExpireDate *time.Time `json:"expireDate,omitempty"`

	// LockedFlag - no documentation
	LockedFlag int `json:"lockedFlag,omitempty"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// ServiceProviderId - <nil>
	ServiceProviderId int `json:"serviceProviderId,omitempty"`

	// Account - The SoftLayer customer account that the domain is registered to.
	Account *SoftLayer_Account `json:"account,omitempty"`

	// DomainRegistrationStatus - no documentation
	DomainRegistrationStatus *SoftLayer_Dns_Domain_Registration_Status `json:"domainRegistrationStatus,omitempty"`

	// RegistrantVerificationStatus - no documentation
	RegistrantVerificationStatus *SoftLayer_Dns_Domain_Registration_Registrant_Verification_Status `json:"registrantVerificationStatus,omitempty"`

	// ServiceProvider - <nil>
	ServiceProvider *SoftLayer_Service_Provider `json:"serviceProvider,omitempty"`
}

func (softlayer_dns_domain_registration *SoftLayer_Dns_Domain_Registration) String() string {
	return "SoftLayer_Dns_Domain_Registration"
}

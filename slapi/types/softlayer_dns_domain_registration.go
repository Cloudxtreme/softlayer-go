package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
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

// AddNameserversToDomain - The addNameserversToDomain method adds nameservers to a domain for a domain
// that already has nameservers assigned to it. This method does not create a nameserver; the
// nameserver must already exist.
func (softlayer_dns_domain_registration *SoftLayer_Dns_Domain_Registration) AddNameserversToDomain(ctx *slapi.RequestContext, nameservers []string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteRegisteredNameserver - The deleteRegisteredNameserver method deletes a nameserver that was
// registered, provided it is not currently serving a domain
func (softlayer_dns_domain_registration *SoftLayer_Dns_Domain_Registration) DeleteRegisteredNameserver(ctx *slapi.RequestContext, nameserver string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetAuthenticationCode - The getAuthenticationCode method retrieves the authentication code for the
// domain.
func (softlayer_dns_domain_registration *SoftLayer_Dns_Domain_Registration) GetAuthenticationCode(ctx *slapi.RequestContext) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetDomainInformation - The getDomainInformation method retrieves all the information for a domain.
func (softlayer_dns_domain_registration *SoftLayer_Dns_Domain_Registration) GetDomainInformation(ctx *slapi.RequestContext) (*SoftLayer_Container_Dns_Domain_Registration_Information, error) {
	var returnValue *SoftLayer_Container_Dns_Domain_Registration_Information
	return returnValue, nil
}

// GetDomainNameservers - The getDomainNameservers method retrieve nameservers information for domain.
func (softlayer_dns_domain_registration *SoftLayer_Dns_Domain_Registration) GetDomainNameservers(ctx *slapi.RequestContext) ([]*SoftLayer_Container_Dns_Domain_Registration_Nameserver, error) {
	var returnValue []*SoftLayer_Container_Dns_Domain_Registration_Nameserver
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Dns_Domain_Registration object whose ID number
// corresponds to the ID number of the init parameter passed to the SoftLayer_Dns_Domain_Registration
// service.
func (softlayer_dns_domain_registration *SoftLayer_Dns_Domain_Registration) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Dns_Domain_Registration, error) {
	var returnValue *SoftLayer_Dns_Domain_Registration
	return returnValue, nil
}

// GetRegisteredNameserver - The getRegisteredNameserver method retrieves registered nameservers.
func (softlayer_dns_domain_registration *SoftLayer_Dns_Domain_Registration) GetRegisteredNameserver(ctx *slapi.RequestContext) (*SoftLayer_Container_Dns_Domain_Registration_Nameserver, error) {
	var returnValue *SoftLayer_Container_Dns_Domain_Registration_Nameserver
	return returnValue, nil
}

// GetRegistrantVerificationStatusDetail - When a domain is registered or transferred, or when the
// registrant contact information is changed, the registrant must reply to an email requesting them to
// confirm that the submitted contact information is correct. This method returns the current state of
// the verification request.
func (softlayer_dns_domain_registration *SoftLayer_Dns_Domain_Registration) GetRegistrantVerificationStatusDetail(ctx *slapi.RequestContext) (*SoftLayer_Container_Dns_Domain_Registration_Registrant_Verification_StatusDetail, error) {
	var returnValue *SoftLayer_Container_Dns_Domain_Registration_Registrant_Verification_StatusDetail
	return returnValue, nil
}

// GetTransferInformation - The getTransferInformation method checks to see if the domain can be
// transferred and also can be used to check the status of the last transfer request.
func (softlayer_dns_domain_registration *SoftLayer_Dns_Domain_Registration) GetTransferInformation(ctx *slapi.RequestContext, domainName string) (*SoftLayer_Container_Dns_Domain_Registration_Transfer_Information, error) {
	var returnValue *SoftLayer_Container_Dns_Domain_Registration_Transfer_Information
	return returnValue, nil
}

// LockDomain - The lockDomain method locks a domain to prevent unauthorized, unwanted or accidental
// changes to the domain name. When set, the following actions are prohibited: * Transferring of the
// domain name * Deletion of the domain name
func (softlayer_dns_domain_registration *SoftLayer_Dns_Domain_Registration) LockDomain(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// LookupDomain - The lookupDomain method checks whether a specified domain name is available for
// registration in TLD's, and suggests other similar domain names, and checks whether they are
// available as well.
func (softlayer_dns_domain_registration *SoftLayer_Dns_Domain_Registration) LookupDomain(ctx *slapi.RequestContext, domainName string) ([]*SoftLayer_Container_Dns_Domain_Registration_Lookup, error) {
	var returnValue []*SoftLayer_Container_Dns_Domain_Registration_Lookup
	return returnValue, nil
}

// ModifyContact - The modifyContact method modifies contact information (admin, billing, owner,
// technical) for a domain.
func (softlayer_dns_domain_registration *SoftLayer_Dns_Domain_Registration) ModifyContact(ctx *slapi.RequestContext, contact SoftLayer_Container_Dns_Domain_Registration_Contact) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// ModifyRegisteredNameserver - The modifyRegisteredNameserver method modifies a nameserver that was
// registered.
func (softlayer_dns_domain_registration *SoftLayer_Dns_Domain_Registration) ModifyRegisteredNameserver(ctx *slapi.RequestContext, oldNameserver string, newNameserver string, ipAddress string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RegisterNameserver - The registerNameserver method creates a nameserver for the domain.
func (softlayer_dns_domain_registration *SoftLayer_Dns_Domain_Registration) RegisterNameserver(ctx *slapi.RequestContext, nameserver string, ipAddress string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveNameserversFromDomain - The removeNameserversFromDomain method removes nameservers from a
// domain for a domain that already has nameservers assigned to it.
func (softlayer_dns_domain_registration *SoftLayer_Dns_Domain_Registration) RemoveNameserversFromDomain(ctx *slapi.RequestContext, nameservers []string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SendAuthenticationCode - The sendAuthenticationCode method sends the authentication code to the
// administrative contact for the domain.
func (softlayer_dns_domain_registration *SoftLayer_Dns_Domain_Registration) SendAuthenticationCode(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SendRegistrantVerificationEmail - When a domain is registered or transferred, or when the registrant
// contact information is changed, the registrant must reply to an email requesting them to confirm
// that the submitted contact information is correct. This method sends the verification email to the
// registrant.
func (softlayer_dns_domain_registration *SoftLayer_Dns_Domain_Registration) SendRegistrantVerificationEmail(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SendTransferApprovalEmail - The sendTransferApprovalEmail method resends a transfer approval email
// message for a transfer that is in 'pending owner approval' state, to the admin contact listed for
// the domain at the time that the transfer request was submitted
func (softlayer_dns_domain_registration *SoftLayer_Dns_Domain_Registration) SendTransferApprovalEmail(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SetAuthenticationCode - The setAuthenticationCode method sets the authentication code for the
// domain. The authentication code is a transfer key and provides an extra level of security,
// safeguarding domain names from unauthorized transfers.
func (softlayer_dns_domain_registration *SoftLayer_Dns_Domain_Registration) SetAuthenticationCode(ctx *slapi.RequestContext, authenticationCode string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// UnlockDomain - no documentation
func (softlayer_dns_domain_registration *SoftLayer_Dns_Domain_Registration) UnlockDomain(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

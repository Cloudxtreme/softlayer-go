package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Dns_Domain_Registration_Registrant_Verification_Status -
// SoftLayer_Dns_Domain_Registration_Registrant_Verification_Status models the state of the registrant.
// Here are the following status codes: *'''Admin Reviewing''': The registrant data has been submitted
// and being reviewed by compliance team. *'''Pending''': The verification process has been inititated,
// and verification email will be sent. *'''Suspended''': The registrant has failed verification and
// the domain has been suspended. *'''Verified''': The registrant has been validated. *'''Verifying''':
// The verification process has been initiated and is waiting for registrant response.
// *'''Unverified''': The verification process has not been inititated.
type SoftLayer_Dns_Domain_Registration_Registrant_Verification_Status struct {

	// Description - The description of the registrant verification status.
	Description string `json:"description"`

	// Id - The unique identifier of the registrant verification status
	Id int `json:"id"`

	// KeyName - The unique keyname of the registrant verification status.
	KeyName string `json:"keyName"`

	// Name - no documentation
	Name string `json:"name"`
}

// GetAllObjects - <nil>
func (softlayer_dns_domain_registration_registrant_verification_status *SoftLayer_Dns_Domain_Registration_Registrant_Verification_Status) GetAllObjects() ([]*SoftLayer_Dns_Domain_Registration_Registrant_Verification_Status, error) {
	var returnValue []*SoftLayer_Dns_Domain_Registration_Registrant_Verification_Status
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_dns_domain_registration_registrant_verification_status *SoftLayer_Dns_Domain_Registration_Registrant_Verification_Status) GetObject() (*SoftLayer_Dns_Domain_Registration_Registrant_Verification_Status, error) {
	var returnValue *SoftLayer_Dns_Domain_Registration_Registrant_Verification_Status
	return returnValue, nil
}

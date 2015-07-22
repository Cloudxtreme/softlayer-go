package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Dns_Domain_Registration_Status - SoftLayer_Dns_Domain_Registration_Status models the state
// of domain name. Here are the following status codes: *'''Active''': This domain name is active.
// *'''Pending Owner Approval''': Pending owner approval for completion of transfer. *'''Pending Admin
// Review''': Pending admin review for transfer. *'''Pending Registry''': Pending registry for
// transfer. *'''Expired''': Domain name has expired.
type SoftLayer_Dns_Domain_Registration_Status struct {

	// Description - The description of the domain registration status names.
	Description string `json:"description"`

	// Id - The unique identifier of the domain registration status
	Id int `json:"id"`

	// KeyName - The unique keyname of the domain registration status.
	KeyName string `json:"keyName"`

	// Name - no documentation
	Name string `json:"name"`
}

// GetAllObjects - <nil>
func (softlayer_dns_domain_registration_status *SoftLayer_Dns_Domain_Registration_Status) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_Dns_Domain_Registration_Status, error) {
	var returnValue []*SoftLayer_Dns_Domain_Registration_Status
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_dns_domain_registration_status *SoftLayer_Dns_Domain_Registration_Status) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Dns_Domain_Registration_Status, error) {
	var returnValue *SoftLayer_Dns_Domain_Registration_Status
	return returnValue, nil
}

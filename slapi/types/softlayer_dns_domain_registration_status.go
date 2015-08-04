package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Dns_Domain_Registration_Status - SoftLayer_Dns_Domain_Registration_Status models the state
// of domain name. Here are the following status codes: *'''Active''': This domain name is active.
// *'''Pending Owner Approval''': Pending owner approval for completion of transfer. *'''Pending Admin
// Review''': Pending admin review for transfer. *'''Pending Registry''': Pending registry for
// transfer. *'''Expired''': Domain name has expired.
type SoftLayer_Dns_Domain_Registration_Status struct {

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// Description - The description of the domain registration status names.
	Description string `json:"description,omitempty"`

	// Id - The unique identifier of the domain registration status
	Id int `json:"id,omitempty"`

	// KeyName - The unique keyname of the domain registration status.
	KeyName string `json:"keyName,omitempty"`
}

func (softlayer_dns_domain_registration_status *SoftLayer_Dns_Domain_Registration_Status) String() string {
	return "SoftLayer_Dns_Domain_Registration_Status"
}

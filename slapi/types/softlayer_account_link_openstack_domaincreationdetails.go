package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Link_OpenStack_DomainCreationDetails - no documentation
type SoftLayer_Account_Link_OpenStack_DomainCreationDetails struct {

	// UserId - Id for the user given the Cloud Admin role for this domain.
	UserId string `json:"userId,omitempty"`

	// UserName - Name for the user given the Cloud Admin role for this domain.
	UserName string `json:"userName,omitempty"`

	// DomainId - no documentation
	DomainId string `json:"domainId,omitempty"`
}

func (softlayer_account_link_openstack_domaincreationdetails *SoftLayer_Account_Link_OpenStack_DomainCreationDetails) String() string {
	return "SoftLayer_Account_Link_OpenStack_DomainCreationDetails"
}

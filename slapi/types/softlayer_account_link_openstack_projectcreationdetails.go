package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Link_OpenStack_ProjectCreationDetails - no documentation
type SoftLayer_Account_Link_OpenStack_ProjectCreationDetails struct {

	// DomainId - no documentation
	DomainId string `json:"domainId,omitempty"`

	// ProjectId - no documentation
	ProjectId string `json:"projectId,omitempty"`

	// ProjectName - no documentation
	ProjectName string `json:"projectName,omitempty"`

	// UserId - Id for the user given the Project Admin role for this project.
	UserId string `json:"userId,omitempty"`

	// UserName - Name for the user given the Project Admin role for this project.
	UserName string `json:"userName,omitempty"`
}

func (softlayer_account_link_openstack_projectcreationdetails *SoftLayer_Account_Link_OpenStack_ProjectCreationDetails) String() string {
	return "SoftLayer_Account_Link_OpenStack_ProjectCreationDetails"
}

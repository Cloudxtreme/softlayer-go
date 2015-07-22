package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Link_OpenStack - <nil>
type SoftLayer_Account_Link_OpenStack struct {

	// DomainId - no documentation
	DomainId string `json:"domainId"`
}

// CreateOSProject - <nil>
func (softlayer_account_link_openstack *SoftLayer_Account_Link_OpenStack) CreateOSProject(request SoftLayer_Account_Link_OpenStack_LinkRequest) error {
	return nil
}

// DeleteOSProject - <nil>
func (softlayer_account_link_openstack *SoftLayer_Account_Link_OpenStack) DeleteOSProject(projectId string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteObject - deleteObject permanently removes an account link and all of it's associated keystone
// data (including users for the associated project). '''This cannot be undone.''' Be wary of running
// this method. If you remove an account link in error you will need to re-create it by creating a new
// SoftLayer_Account_Link_OpenStack object.
func (softlayer_account_link_openstack *SoftLayer_Account_Link_OpenStack) DeleteObject() (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetOSProject - <nil>
func (softlayer_account_link_openstack *SoftLayer_Account_Link_OpenStack) GetOSProject(projectId string) (*SoftLayer_Account_Link_OpenStack_ProjectDetails, error) {
	var returnValue *SoftLayer_Account_Link_OpenStack_ProjectDetails
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_account_link_openstack *SoftLayer_Account_Link_OpenStack) GetObject() (*SoftLayer_Account_Link_OpenStack, error) {
	var returnValue *SoftLayer_Account_Link_OpenStack
	return returnValue, nil
}

// ListOSProjects - <nil>
func (softlayer_account_link_openstack *SoftLayer_Account_Link_OpenStack) ListOSProjects() ([]*SoftLayer_Account_Link_OpenStack_ProjectDetails, error) {
	var returnValue []*SoftLayer_Account_Link_OpenStack_ProjectDetails
	return returnValue, nil
}

package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Account_Link_OpenStack - <nil>
type SoftLayer_Account_Link_OpenStack struct {

	// DomainId - no documentation
	DomainId string `json:"domainId"`
}

func (softlayer_account_link_openstack *SoftLayer_Account_Link_OpenStack) String() string {
	return "SoftLayer_Account_Link_OpenStack"
}

// CreateOSProject - <nil>
func (softlayer_account_link_openstack *SoftLayer_Account_Link_OpenStack) CreateOSProject(ctx *slapi.RequestContext, request SoftLayer_Account_Link_OpenStack_LinkRequest) (*SoftLayer_Account_Link_OpenStack_ProjectCreationDetails, error) {
	var returnValue *SoftLayer_Account_Link_OpenStack_ProjectCreationDetails
	return returnValue, nil
}

// DeleteOSProject - <nil>
func (softlayer_account_link_openstack *SoftLayer_Account_Link_OpenStack) DeleteOSProject(ctx *slapi.RequestContext, projectId string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteObject - deleteObject permanently removes an account link and all of it's associated keystone
// data (including users for the associated project). '''This cannot be undone.''' Be wary of running
// this method. If you remove an account link in error you will need to re-create it by creating a new
// SoftLayer_Account_Link_OpenStack object.
func (softlayer_account_link_openstack *SoftLayer_Account_Link_OpenStack) DeleteObject(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetOSProject - <nil>
func (softlayer_account_link_openstack *SoftLayer_Account_Link_OpenStack) GetOSProject(ctx *slapi.RequestContext, projectId string) (*SoftLayer_Account_Link_OpenStack_ProjectDetails, error) {
	var returnValue *SoftLayer_Account_Link_OpenStack_ProjectDetails
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_account_link_openstack *SoftLayer_Account_Link_OpenStack) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Account_Link_OpenStack, error) {
	var returnValue *SoftLayer_Account_Link_OpenStack
	return returnValue, nil
}

// ListOSProjects - <nil>
func (softlayer_account_link_openstack *SoftLayer_Account_Link_OpenStack) ListOSProjects(ctx *slapi.RequestContext) ([]*SoftLayer_Account_Link_OpenStack_ProjectDetails, error) {
	var returnValue []*SoftLayer_Account_Link_OpenStack_ProjectDetails
	return returnValue, nil
}

package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Resource_Group - <nil>
type SoftLayer_Resource_Group struct {

	// AncestorGroupCount - A count of a resource group's associated group ancestors.
	AncestorGroupCount uint64 `json:"ancestorGroupCount"`

	// AncestorGroups - no documentation
	AncestorGroups []*SoftLayer_Resource_Group `json:"ancestorGroups"`

	// AttributeCount - A count of a resource group's associated attributes.
	AttributeCount uint64 `json:"attributeCount"`

	// Attributes - no documentation
	Attributes []*SoftLayer_Resource_Group_Attribute `json:"attributes"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Description - no documentation
	Description string `json:"description"`

	// HardwareMemberCount - A count of a resource group's associated hardware members.
	HardwareMemberCount uint64 `json:"hardwareMemberCount"`

	// HardwareMembers - no documentation
	HardwareMembers []*SoftLayer_Resource_Group_Member `json:"hardwareMembers"`

	// Id - no documentation
	Id int `json:"id"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// MemberCount - no documentation
	MemberCount uint64 `json:"memberCount"`

	// Members - no documentation
	Members []*SoftLayer_Resource_Group_Member `json:"members"`

	// Name - no documentation
	Name string `json:"name"`

	// RootResourceGroup - no documentation
	RootResourceGroup *SoftLayer_Resource_Group `json:"rootResourceGroup"`

	// RootResourceGroupId - <nil>
	RootResourceGroupId int `json:"rootResourceGroupId"`

	// SubnetMemberCount - A count of a resource group's associated subnet members.
	SubnetMemberCount uint64 `json:"subnetMemberCount"`

	// SubnetMembers - no documentation
	SubnetMembers []*SoftLayer_Resource_Group_Member `json:"subnetMembers"`

	// Template - no documentation
	Template *SoftLayer_Resource_Group_Template `json:"template"`

	// TemplateId - no documentation
	TemplateId int `json:"templateId"`

	// VlanMemberCount - no documentation
	VlanMemberCount uint64 `json:"vlanMemberCount"`

	// VlanMembers - no documentation
	VlanMembers []*SoftLayer_Resource_Group_Member `json:"vlanMembers"`
}

// EditObject - <nil>
func (softlayer_resource_group *SoftLayer_Resource_Group) EditObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Resource_Group) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_resource_group *SoftLayer_Resource_Group) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Resource_Group, error) {
	var returnValue *SoftLayer_Resource_Group
	return returnValue, nil
}

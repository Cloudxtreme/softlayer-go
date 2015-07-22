package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_User_Permission_Group - <nil>
type SoftLayer_User_Permission_Group struct {

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// AccountId - A permission groups associated [[SoftLayer_Account|customer account]] id.
	AccountId int `json:"accountId"`

	// ActionCount - no documentation
	ActionCount uint64 `json:"actionCount"`

	// Actions - <nil>
	Actions []*SoftLayer_User_Permission_Action `json:"actions"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Description - no documentation
	Description string `json:"description"`

	// ExpirationDate - no documentation
	ExpirationDate *time.Time `json:"expirationDate"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyDate - The date the permission group record was last modified.
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// RoleCount - no documentation
	RoleCount uint64 `json:"roleCount"`

	// Roles - <nil>
	Roles []*SoftLayer_User_Permission_Role `json:"roles"`

	// Type - no documentation
	Type *SoftLayer_User_Permission_Group_Type `json:"type"`

	// TypeId - no documentation
	TypeId int `json:"typeId"`
}

// AddAction - <nil>
func (softlayer_user_permission_group *SoftLayer_User_Permission_Group) AddAction(commonOptions *slapi.CommonOptions, action SoftLayer_User_Permission_Action) error {
	return nil
}

// AddBulkActions - <nil>
func (softlayer_user_permission_group *SoftLayer_User_Permission_Group) AddBulkActions(commonOptions *slapi.CommonOptions, actions []SoftLayer_User_Permission_Action) error {
	return nil
}

// AddBulkResourceObjects - <nil>
func (softlayer_user_permission_group *SoftLayer_User_Permission_Group) AddBulkResourceObjects(commonOptions *slapi.CommonOptions, resourceObjects []SoftLayer_Entity, resourceTypeKeyName string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AddResourceObject - <nil>
func (softlayer_user_permission_group *SoftLayer_User_Permission_Group) AddResourceObject(commonOptions *slapi.CommonOptions, resourceObject SoftLayer_Entity, resourceTypeKeyName string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CreateObject - <nil>
func (softlayer_user_permission_group *SoftLayer_User_Permission_Group) CreateObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_User_Permission_Group) (*SoftLayer_User_Permission_Group, error) {
	var returnValue *SoftLayer_User_Permission_Group
	return returnValue, nil
}

// DeleteObject - <nil>
func (softlayer_user_permission_group *SoftLayer_User_Permission_Group) DeleteObject(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - <nil>
func (softlayer_user_permission_group *SoftLayer_User_Permission_Group) EditObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_User_Permission_Group) (*SoftLayer_User_Permission_Group, error) {
	var returnValue *SoftLayer_User_Permission_Group
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_user_permission_group *SoftLayer_User_Permission_Group) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_User_Permission_Group, error) {
	var returnValue *SoftLayer_User_Permission_Group
	return returnValue, nil
}

// LinkRole - <nil>
func (softlayer_user_permission_group *SoftLayer_User_Permission_Group) LinkRole(commonOptions *slapi.CommonOptions, role SoftLayer_User_Permission_Role) error {
	return nil
}

// RemoveAction - <nil>
func (softlayer_user_permission_group *SoftLayer_User_Permission_Group) RemoveAction(commonOptions *slapi.CommonOptions, action SoftLayer_User_Permission_Action) error {
	return nil
}

// RemoveBulkActions - <nil>
func (softlayer_user_permission_group *SoftLayer_User_Permission_Group) RemoveBulkActions(commonOptions *slapi.CommonOptions, actions []SoftLayer_User_Permission_Action) error {
	return nil
}

// RemoveBulkResourceObjects - <nil>
func (softlayer_user_permission_group *SoftLayer_User_Permission_Group) RemoveBulkResourceObjects(commonOptions *slapi.CommonOptions, resourceObjects []SoftLayer_Entity, resourceTypeKeyName string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveResourceObject - <nil>
func (softlayer_user_permission_group *SoftLayer_User_Permission_Group) RemoveResourceObject(commonOptions *slapi.CommonOptions, resourceObject SoftLayer_Entity, resourceTypeKeyName string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// UnlinkRole - <nil>
func (softlayer_user_permission_group *SoftLayer_User_Permission_Group) UnlinkRole(commonOptions *slapi.CommonOptions, role SoftLayer_User_Permission_Role) error {
	return nil
}
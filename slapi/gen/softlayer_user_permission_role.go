package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Permission_Role - <nil>
type SoftLayer_User_Permission_Role struct {

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// AccountId - A permission roles associated [[SoftLayer_Account|customer account]] id.
	AccountId int `json:"accountId"`

	// ActionCount - no documentation
	ActionCount uint64 `json:"actionCount"`

	// Actions - <nil>
	Actions []*SoftLayer_User_Permission_Action `json:"actions"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Description - no documentation
	Description string `json:"description"`

	// GroupCount - no documentation
	GroupCount uint64 `json:"groupCount"`

	// Groups - <nil>
	Groups []*SoftLayer_User_Permission_Group `json:"groups"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyDate - The date the permission role record was last modified.
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// NewUserDefaultFlag - A flag showing if new users should be automatically added to this role.
	NewUserDefaultFlag int `json:"newUserDefaultFlag"`

	// SystemFlag - A flag showing if the permission role was created by our internal system for a single
	// user. If this flag is set only a single user can be assigned to this permission role and it can not
	// be deleted.
	SystemFlag int `json:"systemFlag"`

	// UserCount - no documentation
	UserCount uint64 `json:"userCount"`

	// Users - <nil>
	Users []*SoftLayer_User_Customer `json:"users"`
}

// AddUser - <nil>
func (softlayer_user_permission_role *SoftLayer_User_Permission_Role) AddUser(user SoftLayer_User_Customer) error {
	return nil
}

// CreateObject - <nil>
func (softlayer_user_permission_role *SoftLayer_User_Permission_Role) CreateObject(templateObject SoftLayer_User_Permission_Role) (*SoftLayer_User_Permission_Role, error) {
	var returnValue *SoftLayer_User_Permission_Role
	return returnValue, nil
}

// DeleteObject - <nil>
func (softlayer_user_permission_role *SoftLayer_User_Permission_Role) DeleteObject() (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - <nil>
func (softlayer_user_permission_role *SoftLayer_User_Permission_Role) EditObject(templateObject SoftLayer_User_Permission_Role) (*SoftLayer_User_Permission_Role, error) {
	var returnValue *SoftLayer_User_Permission_Role
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_user_permission_role *SoftLayer_User_Permission_Role) GetObject() (*SoftLayer_User_Permission_Role, error) {
	var returnValue *SoftLayer_User_Permission_Role
	return returnValue, nil
}

// LinkGroup - <nil>
func (softlayer_user_permission_role *SoftLayer_User_Permission_Role) LinkGroup(group SoftLayer_User_Permission_Group) error {
	return nil
}

// RemoveUser - <nil>
func (softlayer_user_permission_role *SoftLayer_User_Permission_Role) RemoveUser(user SoftLayer_User_Customer) error {
	return nil
}

// UnlinkGroup - <nil>
func (softlayer_user_permission_role *SoftLayer_User_Permission_Role) UnlinkGroup(group SoftLayer_User_Permission_Group) error {
	return nil
}

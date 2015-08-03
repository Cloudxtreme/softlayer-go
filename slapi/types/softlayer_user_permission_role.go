package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Permission_Role - <nil>
type SoftLayer_User_Permission_Role struct {

	// ModifyDate - The date the permission role record was last modified.
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// AccountId - A permission roles associated [[SoftLayer_Account|customer account]] id.
	AccountId int `json:"accountId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Description - no documentation
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`

	// NewUserDefaultFlag - A flag showing if new users should be automatically added to this role.
	NewUserDefaultFlag int `json:"newUserDefaultFlag"`

	// SystemFlag - A flag showing if the permission role was created by our internal system for a single
	// user. If this flag is set only a single user can be assigned to this permission role and it can not
	// be deleted.
	SystemFlag int `json:"systemFlag"`
}

func (softlayer_user_permission_role *SoftLayer_User_Permission_Role) String() string {
	return "SoftLayer_User_Permission_Role"
}

// SoftLayer_User_Permission_Role_Extended is SoftLayer_User_Permission_Role with all maskable types.
type SoftLayer_User_Permission_Role_Extended struct {
	SoftLayer_User_Permission_Role

	// GroupCount - no documentation
	GroupCount uint64 `json:"groupCount"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// Groups - <nil>
	Groups []*SoftLayer_User_Permission_Group `json:"groups"`

	// ActionCount - no documentation
	ActionCount uint64 `json:"actionCount"`

	// Actions - <nil>
	Actions []*SoftLayer_User_Permission_Action `json:"actions"`

	// Users - <nil>
	Users []*SoftLayer_User_Customer `json:"users"`

	// UserCount - no documentation
	UserCount uint64 `json:"userCount"`
}

func (softlayer_user_permission_role *SoftLayer_User_Permission_Role_Extended) String() string {
	return "SoftLayer_User_Permission_Role"
}

package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Permission_Role - <nil>
type SoftLayer_User_Permission_Role struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// AccountId - A permission roles associated [[SoftLayer_Account|customer account]] id.
	AccountId int `json:"accountId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// SystemFlag - A flag showing if the permission role was created by our internal system for a single
	// user. If this flag is set only a single user can be assigned to this permission role and it can not
	// be deleted.
	SystemFlag int `json:"systemFlag,omitempty"`

	// ModifyDate - The date the permission role record was last modified.
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// NewUserDefaultFlag - A flag showing if new users should be automatically added to this role.
	NewUserDefaultFlag int `json:"newUserDefaultFlag,omitempty"`
}

func (softlayer_user_permission_role *SoftLayer_User_Permission_Role) String() string {
	return "SoftLayer_User_Permission_Role"
}

// SoftLayer_User_Permission_Role_Extended is SoftLayer_User_Permission_Role with all maskable types.
type SoftLayer_User_Permission_Role_Extended struct {
	SoftLayer_User_Permission_Role

	// Users - <nil>
	Users []*SoftLayer_User_Customer `json:"users,omitempty"`

	// UserCount - no documentation
	UserCount uint64 `json:"userCount,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// Actions - <nil>
	Actions []*SoftLayer_User_Permission_Action `json:"actions,omitempty"`

	// ActionCount - no documentation
	ActionCount uint64 `json:"actionCount,omitempty"`

	// Groups - <nil>
	Groups []*SoftLayer_User_Permission_Group `json:"groups,omitempty"`

	// GroupCount - no documentation
	GroupCount uint64 `json:"groupCount,omitempty"`
}

func (softlayer_user_permission_role *SoftLayer_User_Permission_Role_Extended) String() string {
	return "SoftLayer_User_Permission_Role"
}

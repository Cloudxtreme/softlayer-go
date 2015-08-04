package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Permission_Role - <nil>
type SoftLayer_User_Permission_Role struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// SystemFlag - A flag showing if the permission role was created by our internal system for a single
	// user. If this flag is set only a single user can be assigned to this permission role and it can not
	// be deleted.
	SystemFlag int `json:"systemFlag,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// ModifyDate - The date the permission role record was last modified.
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// NewUserDefaultFlag - A flag showing if new users should be automatically added to this role.
	NewUserDefaultFlag int `json:"newUserDefaultFlag,omitempty"`

	// AccountId - A permission roles associated [[SoftLayer_Account|customer account]] id.
	AccountId int `json:"accountId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// Users - <nil>
	Users []*SoftLayer_User_Customer `json:"users,omitempty"`

	// GroupCount - no documentation
	GroupCount uint64 `json:"groupCount,omitempty"`

	// Actions - <nil>
	Actions []*SoftLayer_User_Permission_Action `json:"actions,omitempty"`

	// ActionCount - no documentation
	ActionCount uint64 `json:"actionCount,omitempty"`

	// UserCount - no documentation
	UserCount uint64 `json:"userCount,omitempty"`

	// Groups - <nil>
	Groups []*SoftLayer_User_Permission_Group `json:"groups,omitempty"`
}

func (softlayer_user_permission_role *SoftLayer_User_Permission_Role) String() string {
	return "SoftLayer_User_Permission_Role"
}

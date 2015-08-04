package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Permission_Group - <nil>
type SoftLayer_User_Permission_Group struct {

	// ExpirationDate - no documentation
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// TypeId - no documentation
	TypeId int `json:"typeId,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// ModifyDate - The date the permission group record was last modified.
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// AccountId - A permission groups associated [[SoftLayer_Account|customer account]] id.
	AccountId int `json:"accountId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`
}

func (softlayer_user_permission_group *SoftLayer_User_Permission_Group) String() string {
	return "SoftLayer_User_Permission_Group"
}

// SoftLayer_User_Permission_Group_Extended is SoftLayer_User_Permission_Group with all maskable types.
type SoftLayer_User_Permission_Group_Extended struct {
	SoftLayer_User_Permission_Group

	// Roles - <nil>
	Roles []*SoftLayer_User_Permission_Role `json:"roles,omitempty"`

	// RoleCount - no documentation
	RoleCount uint64 `json:"roleCount,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// Actions - <nil>
	Actions []*SoftLayer_User_Permission_Action `json:"actions,omitempty"`

	// Type - no documentation
	Type *SoftLayer_User_Permission_Group_Type `json:"type,omitempty"`

	// ActionCount - no documentation
	ActionCount uint64 `json:"actionCount,omitempty"`
}

func (softlayer_user_permission_group *SoftLayer_User_Permission_Group_Extended) String() string {
	return "SoftLayer_User_Permission_Group"
}

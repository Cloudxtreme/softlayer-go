package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Permission_Group - <nil>
type SoftLayer_User_Permission_Group struct {

	// ExpirationDate - no documentation
	ExpirationDate *time.Time `json:"expirationDate"`

	// AccountId - A permission groups associated [[SoftLayer_Account|customer account]] id.
	AccountId int `json:"accountId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// TypeId - no documentation
	TypeId int `json:"typeId"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyDate - The date the permission group record was last modified.
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// Description - no documentation
	Description string `json:"description"`
}

func (softlayer_user_permission_group *SoftLayer_User_Permission_Group) String() string {
	return "SoftLayer_User_Permission_Group"
}

// SoftLayer_User_Permission_Group_Extended is SoftLayer_User_Permission_Group with all maskable types.
type SoftLayer_User_Permission_Group_Extended struct {
	SoftLayer_User_Permission_Group

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// Type - no documentation
	Type *SoftLayer_User_Permission_Group_Type `json:"type"`

	// ActionCount - no documentation
	ActionCount uint64 `json:"actionCount"`

	// RoleCount - no documentation
	RoleCount uint64 `json:"roleCount"`

	// Actions - <nil>
	Actions []*SoftLayer_User_Permission_Action `json:"actions"`

	// Roles - <nil>
	Roles []*SoftLayer_User_Permission_Role `json:"roles"`
}

func (softlayer_user_permission_group *SoftLayer_User_Permission_Group_Extended) String() string {
	return "SoftLayer_User_Permission_Group"
}

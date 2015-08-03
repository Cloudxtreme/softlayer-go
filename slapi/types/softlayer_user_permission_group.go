package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Permission_Group - <nil>
type SoftLayer_User_Permission_Group struct {

	// Description - no documentation
	Description string `json:"description"`

	// TypeId - no documentation
	TypeId int `json:"typeId"`

	// ExpirationDate - no documentation
	ExpirationDate *time.Time `json:"expirationDate"`

	// ModifyDate - The date the permission group record was last modified.
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// AccountId - A permission groups associated [[SoftLayer_Account|customer account]] id.
	AccountId int `json:"accountId"`

	// Id - no documentation
	Id int `json:"id"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`
}

// SoftLayer_User_Permission_Group_Extended is SoftLayer_User_Permission_Group with all maskable types.
type SoftLayer_User_Permission_Group_Extended struct {
	SoftLayer_User_Permission_Group

	// RoleCount - no documentation
	RoleCount uint64 `json:"roleCount"`

	// Actions - <nil>
	Actions []*SoftLayer_User_Permission_Action `json:"actions"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// Roles - <nil>
	Roles []*SoftLayer_User_Permission_Role `json:"roles"`

	// ActionCount - no documentation
	ActionCount uint64 `json:"actionCount"`

	// Type - no documentation
	Type *SoftLayer_User_Permission_Group_Type `json:"type"`
}

func (softlayer_user_permission_group *SoftLayer_User_Permission_Group) String() string {
	return "SoftLayer_User_Permission_Group"
}

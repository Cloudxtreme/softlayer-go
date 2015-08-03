package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Permission_Group_Type - <nil>
type SoftLayer_User_Permission_Group_Type struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Id - <nil>
	Id int `json:"id"`

	// KeyName - <nil>
	KeyName string `json:"keyName"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - <nil>
	Name string `json:"name"`
}

func (softlayer_user_permission_group_type *SoftLayer_User_Permission_Group_Type) String() string {
	return "SoftLayer_User_Permission_Group_Type"
}

// SoftLayer_User_Permission_Group_Type_Extended is SoftLayer_User_Permission_Group_Type with all maskable types.
type SoftLayer_User_Permission_Group_Type_Extended struct {
	SoftLayer_User_Permission_Group_Type

	// GroupCount - no documentation
	GroupCount uint64 `json:"groupCount"`

	// Groups - <nil>
	Groups []*SoftLayer_User_Permission_Group `json:"groups"`
}

func (softlayer_user_permission_group_type *SoftLayer_User_Permission_Group_Type_Extended) String() string {
	return "SoftLayer_User_Permission_Group_Type"
}

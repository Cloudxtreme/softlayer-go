package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Permission_Group_Type - <nil>
type SoftLayer_User_Permission_Group_Type struct {

	// KeyName - <nil>
	KeyName string `json:"keyName,omitempty"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Name - <nil>
	Name string `json:"name,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// GroupCount - no documentation
	GroupCount uint64 `json:"groupCount,omitempty"`

	// Groups - <nil>
	Groups []*SoftLayer_User_Permission_Group `json:"groups,omitempty"`
}

func (softlayer_user_permission_group_type *SoftLayer_User_Permission_Group_Type) String() string {
	return "SoftLayer_User_Permission_Group_Type"
}

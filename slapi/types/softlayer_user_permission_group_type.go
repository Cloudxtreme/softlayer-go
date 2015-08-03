package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_User_Permission_Group_Type - <nil>
type SoftLayer_User_Permission_Group_Type struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// GroupCount - no documentation
	GroupCount uint64 `json:"groupCount"`

	// Groups - <nil>
	Groups []*SoftLayer_User_Permission_Group `json:"groups"`

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

// GetObject - <nil>
func (softlayer_user_permission_group_type *SoftLayer_User_Permission_Group_Type) GetObject(ctx *slapi.RequestContext) (*SoftLayer_User_Permission_Group_Type, error) {
	var returnValue *SoftLayer_User_Permission_Group_Type
	return returnValue, nil
}

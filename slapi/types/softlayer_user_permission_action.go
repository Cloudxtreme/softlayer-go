package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_User_Permission_Action - <nil>
type SoftLayer_User_Permission_Action struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Description - <nil>
	Description string `json:"description"`

	// Id - <nil>
	Id int `json:"id"`

	// KeyName - <nil>
	KeyName string `json:"keyName"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - <nil>
	Name string `json:"name"`
}

// GetAllObjects - <nil>
func (softlayer_user_permission_action *SoftLayer_User_Permission_Action) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_User_Permission_Action, error) {
	var returnValue []*SoftLayer_User_Permission_Action
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_user_permission_action *SoftLayer_User_Permission_Action) GetObject(ctx *slapi.RequestContext) (*SoftLayer_User_Permission_Action, error) {
	var returnValue *SoftLayer_User_Permission_Action
	return returnValue, nil
}

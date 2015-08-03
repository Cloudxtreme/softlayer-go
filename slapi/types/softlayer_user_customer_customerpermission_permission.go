package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_User_Customer_CustomerPermission_Permission - Each SoftLayer portal account is assigned a
// series of permissions that determine what access the user has to functions within the SoftLayer
// customer portal. This status is reflected in the SoftLayer_User_Customer_Status data type.
// Permissions differ from user status in that user status applies globally to the portal while user
// permissions are applied to specific portal functions.
type SoftLayer_User_Customer_CustomerPermission_Permission struct {

	// Key - no documentation
	Key string `json:"key"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// Name - no documentation
	Name string `json:"name"`
}

func (softlayer_user_customer_customerpermission_permission *SoftLayer_User_Customer_CustomerPermission_Permission) String() string {
	return "SoftLayer_User_Customer_CustomerPermission_Permission"
}

// GetAllObjects - no documentation
func (softlayer_user_customer_customerpermission_permission *SoftLayer_User_Customer_CustomerPermission_Permission) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_User_Customer_CustomerPermission_Permission, error) {
	var returnValue []*SoftLayer_User_Customer_CustomerPermission_Permission
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_User_Customer_CustomerPermission_Permission object
// whose ID number corresponds to the ID number of the init parameter passed to the
// SoftLayer_User_Customer_CustomerPermission_Permission service.
func (softlayer_user_customer_customerpermission_permission *SoftLayer_User_Customer_CustomerPermission_Permission) GetObject(ctx *slapi.RequestContext) (*SoftLayer_User_Customer_CustomerPermission_Permission, error) {
	var returnValue *SoftLayer_User_Customer_CustomerPermission_Permission
	return returnValue, nil
}

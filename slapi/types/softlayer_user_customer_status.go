package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_User_Customer_Status - Each SoftLayer portal account is assigned a status code that
// determines how it's treated in the customer portal. This status is reflected in the
// SoftLayer_User_Customer_Status data type. Status differs from user permissions in that user status
// applies globally to the portal while user permissions are applied to specific portal functions.
type SoftLayer_User_Customer_Status struct {

	// Id - no documentation
	Id int `json:"id"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// Name - A user's status. This can be either "Active" for user accounts with portal access, "Inactive"
	// for users disabled by another portal user, "Disabled" for accounts turned off by SoftLayer, or Only"
	// for user accounts with no access to the customer portal but VPN access to the private network.
	Name string `json:"name"`
}

// GetAllObjects - no documentation
func (softlayer_user_customer_status *SoftLayer_User_Customer_Status) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_User_Customer_Status, error) {
	var returnValue []*SoftLayer_User_Customer_Status
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_User_Customer_Status object whose ID number
// corresponds to the ID number of the init parameter passed to the SoftLayer_User_Customer_Status
// service.
func (softlayer_user_customer_status *SoftLayer_User_Customer_Status) GetObject(ctx *slapi.RequestContext) (*SoftLayer_User_Customer_Status, error) {
	var returnValue *SoftLayer_User_Customer_Status
	return returnValue, nil
}

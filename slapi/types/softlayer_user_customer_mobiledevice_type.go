package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_User_Customer_MobileDevice_Type - Describes a supported class of mobile device. In this
// the word class is used in the context of classes of consumer electronic devices, the two most
// prominent examples being mobile phones and tablets.
type SoftLayer_User_Customer_MobileDevice_Type struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Description - no documentation
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`
}

// GetAllObjects - <nil>
func (softlayer_user_customer_mobiledevice_type *SoftLayer_User_Customer_MobileDevice_Type) GetAllObjects(commonOptions *slapi.CommonOptions) ([]*SoftLayer_User_Customer_MobileDevice_Type, error) {
	var returnValue []*SoftLayer_User_Customer_MobileDevice_Type
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_user_customer_mobiledevice_type *SoftLayer_User_Customer_MobileDevice_Type) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_User_Customer_MobileDevice_Type, error) {
	var returnValue *SoftLayer_User_Customer_MobileDevice_Type
	return returnValue, nil
}

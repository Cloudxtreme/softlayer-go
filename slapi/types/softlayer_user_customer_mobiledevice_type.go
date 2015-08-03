package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
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

func (softlayer_user_customer_mobiledevice_type *SoftLayer_User_Customer_MobileDevice_Type) String() string {
	return "SoftLayer_User_Customer_MobileDevice_Type"
}

// GetAllObjects - <nil>
func (softlayer_user_customer_mobiledevice_type *SoftLayer_User_Customer_MobileDevice_Type) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_User_Customer_MobileDevice_Type, error) {
	var returnValue []*SoftLayer_User_Customer_MobileDevice_Type
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_user_customer_mobiledevice_type *SoftLayer_User_Customer_MobileDevice_Type) GetObject(ctx *slapi.RequestContext) (*SoftLayer_User_Customer_MobileDevice_Type, error) {
	var returnValue *SoftLayer_User_Customer_MobileDevice_Type
	return returnValue, nil
}

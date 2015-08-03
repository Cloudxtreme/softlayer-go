package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_User_Customer_MobileDevice_OperatingSystem - This class represents the mobile operating
// system installed on a user's registered mobile device. It assists us when determining the how to get
// a push notification to the user.
type SoftLayer_User_Customer_MobileDevice_OperatingSystem struct {

	// BuildVersion - no documentation
	BuildVersion int `json:"buildVersion"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Description - no documentation
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`

	// MajorVersion - no documentation
	MajorVersion int `json:"majorVersion"`

	// MinorVersion - no documentation
	MinorVersion int `json:"minorVersion"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`
}

func (softlayer_user_customer_mobiledevice_operatingsystem *SoftLayer_User_Customer_MobileDevice_OperatingSystem) String() string {
	return "SoftLayer_User_Customer_MobileDevice_OperatingSystem"
}

// GetAllObjects - <nil>
func (softlayer_user_customer_mobiledevice_operatingsystem *SoftLayer_User_Customer_MobileDevice_OperatingSystem) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_User_Customer_MobileDevice_OperatingSystem, error) {
	var returnValue []*SoftLayer_User_Customer_MobileDevice_OperatingSystem
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_user_customer_mobiledevice_operatingsystem *SoftLayer_User_Customer_MobileDevice_OperatingSystem) GetObject(ctx *slapi.RequestContext) (*SoftLayer_User_Customer_MobileDevice_OperatingSystem, error) {
	var returnValue *SoftLayer_User_Customer_MobileDevice_OperatingSystem
	return returnValue, nil
}

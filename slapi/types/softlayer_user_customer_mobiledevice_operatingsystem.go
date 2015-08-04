package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Customer_MobileDevice_OperatingSystem - This class represents the mobile operating
// system installed on a user's registered mobile device. It assists us when determining the how to get
// a push notification to the user.
type SoftLayer_User_Customer_MobileDevice_OperatingSystem struct {

	// MinorVersion - no documentation
	MinorVersion int `json:"minorVersion,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// BuildVersion - no documentation
	BuildVersion int `json:"buildVersion,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// MajorVersion - no documentation
	MajorVersion int `json:"majorVersion,omitempty"`
}

func (softlayer_user_customer_mobiledevice_operatingsystem *SoftLayer_User_Customer_MobileDevice_OperatingSystem) String() string {
	return "SoftLayer_User_Customer_MobileDevice_OperatingSystem"
}

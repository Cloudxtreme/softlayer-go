package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Customer_MobileDevice_Type - Describes a supported class of mobile device. In this
// the word class is used in the context of classes of consumer electronic devices, the two most
// prominent examples being mobile phones and tablets.
type SoftLayer_User_Customer_MobileDevice_Type struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`
}

func (softlayer_user_customer_mobiledevice_type *SoftLayer_User_Customer_MobileDevice_Type) String() string {
	return "SoftLayer_User_Customer_MobileDevice_Type"
}

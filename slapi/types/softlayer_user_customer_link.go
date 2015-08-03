package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Customer_Link - <nil>
type SoftLayer_User_Customer_Link struct {

	// DestinationUserAlphanumericId - <nil>
	DestinationUserAlphanumericId string `json:"destinationUserAlphanumericId"`

	// DestinationUserId - <nil>
	DestinationUserId int `json:"destinationUserId"`

	// Id - <nil>
	Id int `json:"id"`

	// ServiceProviderId - <nil>
	ServiceProviderId int `json:"serviceProviderId"`

	// UserId - <nil>
	UserId int `json:"userId"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`
}

func (softlayer_user_customer_link *SoftLayer_User_Customer_Link) String() string {
	return "SoftLayer_User_Customer_Link"
}

// SoftLayer_User_Customer_Link_Extended is SoftLayer_User_Customer_Link with all maskable types.
type SoftLayer_User_Customer_Link_Extended struct {
	SoftLayer_User_Customer_Link

	// ServiceProvider - <nil>
	ServiceProvider *SoftLayer_Service_Provider `json:"serviceProvider"`

	// User - <nil>
	User *SoftLayer_User_Customer `json:"user"`
}

func (softlayer_user_customer_link *SoftLayer_User_Customer_Link_Extended) String() string {
	return "SoftLayer_User_Customer_Link"
}

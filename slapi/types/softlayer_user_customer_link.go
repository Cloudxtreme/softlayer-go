package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Customer_Link - <nil>
type SoftLayer_User_Customer_Link struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// DestinationUserAlphanumericId - <nil>
	DestinationUserAlphanumericId string `json:"destinationUserAlphanumericId"`

	// DestinationUserId - <nil>
	DestinationUserId int `json:"destinationUserId"`

	// Id - <nil>
	Id int `json:"id"`

	// ServiceProvider - <nil>
	ServiceProvider *SoftLayer_Service_Provider `json:"serviceProvider"`

	// ServiceProviderId - <nil>
	ServiceProviderId int `json:"serviceProviderId"`

	// User - <nil>
	User *SoftLayer_User_Customer `json:"user"`

	// UserId - <nil>
	UserId int `json:"userId"`
}

func (softlayer_user_customer_link *SoftLayer_User_Customer_Link) String() string {
	return "SoftLayer_User_Customer_Link"
}

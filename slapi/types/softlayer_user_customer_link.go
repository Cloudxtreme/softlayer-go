package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Customer_Link - <nil>
type SoftLayer_User_Customer_Link struct {

	// UserId - <nil>
	UserId int `json:"userId,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// DestinationUserAlphanumericId - <nil>
	DestinationUserAlphanumericId string `json:"destinationUserAlphanumericId,omitempty"`

	// DestinationUserId - <nil>
	DestinationUserId int `json:"destinationUserId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// ServiceProviderId - <nil>
	ServiceProviderId int `json:"serviceProviderId,omitempty"`

	// User - <nil>
	User *SoftLayer_User_Customer `json:"user,omitempty"`

	// ServiceProvider - <nil>
	ServiceProvider *SoftLayer_Service_Provider `json:"serviceProvider,omitempty"`
}

func (softlayer_user_customer_link *SoftLayer_User_Customer_Link) String() string {
	return "SoftLayer_User_Customer_Link"
}

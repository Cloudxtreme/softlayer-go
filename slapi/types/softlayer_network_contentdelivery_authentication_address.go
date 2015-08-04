package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_ContentDelivery_Authentication_Address - The
// SoftLayer_Network_ContentDelivery_Authentication_Address data type models an individual IP address
// that CDN allow or deny access from.
type SoftLayer_Network_ContentDelivery_Authentication_Address struct {

	// CdnAccountId - no documentation
	CdnAccountId int `json:"cdnAccountId,omitempty"`

	// Id - The internal identifier of an authentication IP address
	Id int `json:"id,omitempty"`

	// IpAddress - The IP address that you want to block or allow access to
	IpAddress string `json:"ipAddress,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// UserId - The internal identifier of the user who created an authentication IP record
	UserId int `json:"userId,omitempty"`

	// AccessType - The type of access on an IP address. It can be or
	AccessType string `json:"accessType,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Name - The name of an authentication IP. This helps you to keep track of IP addresses.
	Name string `json:"name,omitempty"`

	// Priority - The priority of an authentication IP address. The smaller number, the higher in priority.
	// Higher priority IP will be matched first.
	Priority int `json:"priority,omitempty"`
}

func (softlayer_network_contentdelivery_authentication_address *SoftLayer_Network_ContentDelivery_Authentication_Address) String() string {
	return "SoftLayer_Network_ContentDelivery_Authentication_Address"
}

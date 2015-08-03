package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_ContentDelivery_Authentication_Token - The
// SoftLayer_Network_ContentDelivery_Authentication_Address data type models an individual IP address
// that CDN allow or deny access from.
type SoftLayer_Network_ContentDelivery_Authentication_Token struct {

	// CdnAccountId - no documentation
	CdnAccountId int `json:"cdnAccountId"`

	// ClientIp - no documentation
	ClientIp string `json:"clientIp"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Name - The customer id. You can use this optional value to tie a user id to an authentication token.
	Name string `json:"name"`

	// Referrer - no documentation
	Referrer string `json:"referrer"`

	// Token - no documentation
	Token string `json:"token"`
}

func (softlayer_network_contentdelivery_authentication_token *SoftLayer_Network_ContentDelivery_Authentication_Token) String() string {
	return "SoftLayer_Network_ContentDelivery_Authentication_Token"
}

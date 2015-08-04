package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_ContentDelivery_Authentication_Token - The
// SoftLayer_Network_ContentDelivery_Authentication_Address data type models an individual IP address
// that CDN allow or deny access from.
type SoftLayer_Network_ContentDelivery_Authentication_Token struct {

	// Name - The customer id. You can use this optional value to tie a user id to an authentication token.
	Name string `json:"name,omitempty"`

	// Referrer - no documentation
	Referrer string `json:"referrer,omitempty"`

	// Token - no documentation
	Token string `json:"token,omitempty"`

	// CdnAccountId - no documentation
	CdnAccountId int `json:"cdnAccountId,omitempty"`

	// ClientIp - no documentation
	ClientIp string `json:"clientIp,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`
}

func (softlayer_network_contentdelivery_authentication_token *SoftLayer_Network_ContentDelivery_Authentication_Token) String() string {
	return "SoftLayer_Network_ContentDelivery_Authentication_Token"
}

package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Customer_Access_Authentication - SoftLayer_User_Customer_Access_Authentication models
// a single attempt to log into the SoftLayer customer portal. A
// SoftLayer_User_Customer_Access_Authentication record is created every time a user attempts to log
// into the portal. Use this service to audit your users' portal activity and diagnose potential
// security breaches of your SoftLayer portal accounts. Unsuccessful login attempts can be caused by an
// incorrect password, failing to answer or not answering a login security question if the user has
// them configured, or attempting to log in from an IP address outside of the user's IP address
// restriction list. SoftLayer employees periodically log into our customer portal as users to diagnose
// portal issues, verify settings and configuration, and to perform maintenance on your account or
// services. SoftLayer employees only log into customer accounts from the following IP ranges: *
// 2607:f0d0:1000::/48 * 2607:f0d0:2000::/48 * 2607:f0d0:3000::/48 * 66.228.118.67/32 *
// 66.228.118.86/32
type SoftLayer_User_Customer_Access_Authentication struct {

	// SuccessFlag - Whether an attempt to log into the SoftLayer customer portal was successful or not.
	SuccessFlag bool `json:"successFlag,omitempty"`

	// UserId - The internal identifier of the user who attempted to log into the SoftLayer customer
	// portal.
	UserId int `json:"userId,omitempty"`

	// Username - The username used when attempting to log into the SoftLayer customer portal
	Username string `json:"username,omitempty"`

	// CreateDate - The date of an attempt to log into the SoftLayer customer portal.
	CreateDate *time.Time `json:"createDate,omitempty"`

	// IpAddress - The IP address of the user who attempted to log into the SoftLayer customer portal.
	IpAddress string `json:"ipAddress,omitempty"`

	// User - The user who has attempted to log into the SoftLayer customer portal.
	User *SoftLayer_User_Customer `json:"user,omitempty"`
}

func (softlayer_user_customer_access_authentication *SoftLayer_User_Customer_Access_Authentication) String() string {
	return "SoftLayer_User_Customer_Access_Authentication"
}

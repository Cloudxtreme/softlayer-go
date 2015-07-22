package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_ContentDelivery_Authentication_Address - The
// SoftLayer_Network_ContentDelivery_Authentication_Address data type models an individual IP address
// that CDN allow or deny access from.
type SoftLayer_Network_ContentDelivery_Authentication_Address struct {

	// AccessType - The type of access on an IP address. It can be or
	AccessType string `json:"accessType"`

	// CdnAccountId - no documentation
	CdnAccountId int `json:"cdnAccountId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Id - The internal identifier of an authentication IP address
	Id int `json:"id"`

	// IpAddress - The IP address that you want to block or allow access to
	IpAddress string `json:"ipAddress"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - The name of an authentication IP. This helps you to keep track of IP addresses.
	Name string `json:"name"`

	// Priority - The priority of an authentication IP address. The smaller number, the higher in priority.
	// Higher priority IP will be matched first.
	Priority int `json:"priority"`

	// UserId - The internal identifier of the user who created an authentication IP record
	UserId int `json:"userId"`
}

// CreateObject - This method creates an authentication IP record. Required parameters are *
// cdnAccountId - A CDN account id that belongs to your SoftLayer Account * ipAddress - An IP address
// or a IP range * accessType- It can be or
func (softlayer_network_contentdelivery_authentication_address *SoftLayer_Network_ContentDelivery_Authentication_Address) CreateObject(ctx *slapi.RequestContext, templateObject SoftLayer_Network_ContentDelivery_Authentication_Address) (*SoftLayer_Network_ContentDelivery_Authentication_Address, error) {
	var returnValue *SoftLayer_Network_ContentDelivery_Authentication_Address
	return returnValue, nil
}

// DeleteObject - no documentation
func (softlayer_network_contentdelivery_authentication_address *SoftLayer_Network_ContentDelivery_Authentication_Address) DeleteObject(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - This method let you edit an authentication IP object by passing a modified object.
func (softlayer_network_contentdelivery_authentication_address *SoftLayer_Network_ContentDelivery_Authentication_Address) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Network_ContentDelivery_Authentication_Address) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Network_ContentDelivery_Authentication_Address object
// whose ID number corresponds to the ID number of the initial parameter passed to the
// SoftLayer_Network_ContentDelivery_Authentication_Address service. You can only retrieve
// authentication IP addresses assigned to one of your CDN account.
func (softlayer_network_contentdelivery_authentication_address *SoftLayer_Network_ContentDelivery_Authentication_Address) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_ContentDelivery_Authentication_Address, error) {
	var returnValue *SoftLayer_Network_ContentDelivery_Authentication_Address
	return returnValue, nil
}

// RearrangeAuthenticationIp - The authentication IP address match occurs from the higher priority IP
// to the lower. This method will be helpful if you want to modify the order (priority) of the
// authentication IP addresses. You can use this method instead of editing individual authentication IP
// addresses. You can retrieve authentication IP address using
// [[SoftLayer_Network_ContentDelivery_Account::getAuthenticationIpAddresses|getAuthenticationIpAddresses]]
// method. Then, rearrange the authentication IP addresses and pass them to this method. When creating
// template objects as parameter, make sure to include the id of each authentication IP addresses. You
// must provide every authentication IP address. New priorities will be assigned to each authentication
// IP addresses in the order of they are passed.
func (softlayer_network_contentdelivery_authentication_address *SoftLayer_Network_ContentDelivery_Authentication_Address) RearrangeAuthenticationIp(ctx *slapi.RequestContext, cdnAccountId int, templateObjects []SoftLayer_Network_ContentDelivery_Authentication_Address) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

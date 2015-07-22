package sl

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

// CreateObject - This method is deprecated! Use the
// [[SoftLayer_Network_ContentDelivery_Authentication_Token::getTimedToken|getTimedToken]] method. This
// method creates a managed authentication token. When passing a parameter, the only required value is
// your CDN account id which can be obtained from the
// [[SoftLayer_Account::getCdnAccounts|getCdnAccounts]] method. There are 3 optional parameters you can
// pass: * name - This helps you keep track of managed tokens. * referrer - If set, the token
// validation will check the client's referrer. Keep in mind, if a client doesn't have the referrer
// information, the token validation will fail. * clientIp - If set, the token validation will check
// the client's IP address.
func (softlayer_network_contentdelivery_authentication_token *SoftLayer_Network_ContentDelivery_Authentication_Token) CreateObject(templateObject SoftLayer_Network_ContentDelivery_Authentication_Token) (*SoftLayer_Network_ContentDelivery_Authentication_Token, error) {
	var returnValue *SoftLayer_Network_ContentDelivery_Authentication_Token
	return returnValue, nil
}

// GetAllManagedTokens - This method is deprecated! This method returns all managed tokens for a CDN
// account.
func (softlayer_network_contentdelivery_authentication_token *SoftLayer_Network_ContentDelivery_Authentication_Token) GetAllManagedTokens(cdnAccountId int) ([]*SoftLayer_Network_ContentDelivery_Authentication_Token, error) {
	var returnValue []*SoftLayer_Network_ContentDelivery_Authentication_Token
	return returnValue, nil
}

// GetObject - This method is deprecated! getObject retrieves the
// SoftLayer_Network_ContentDelivery_Authentication_Token object whose ID number corresponds to the ID
// number of the initial parameter passed to the SoftLayer_Network_ContentDelivery_Authentication_Token
// service. You can only retrieve managed tokens assigned to one of your CDN account.
func (softlayer_network_contentdelivery_authentication_token *SoftLayer_Network_ContentDelivery_Authentication_Token) GetObject() (*SoftLayer_Network_ContentDelivery_Authentication_Token, error) {
	var returnValue *SoftLayer_Network_ContentDelivery_Authentication_Token
	return returnValue, nil
}

// GetTimedToken - This method returns an authentication token that expires after the seconds you
// specify. You can provide number of seconds to manage the token life. This parameter sets the
// expiration time for a token. A valid life time must be an integer between 60 and 604800 (1 week). A
// customer can also provide client ip and (or) referrer information. If used, a client from the same
// IP and referrer can view the protected contents. A valid IP address must be an IPv4 format or an IP
// block. if you want to block access from IP 211.37.0.0/16, you can enter "211.37." instead. IP blocks
// can be specified in the manner of "8bit times n". The referrer is the URL of the previous webpage
// from which a link was followed. A referrer should not include "http://" prefix and it can be maximum
// of 30 characters.
func (softlayer_network_contentdelivery_authentication_token *SoftLayer_Network_ContentDelivery_Authentication_Token) GetTimedToken(cdnAccountId int, tokenLife int, clientIp string, referrer string, mediaType string) (string, error) {
	var returnValue string
	return returnValue, nil
}

// RevokeAllManagedTokens - This method is deprecated! This method revokes all managed tokens belong to
// a CDN account.
func (softlayer_network_contentdelivery_authentication_token *SoftLayer_Network_ContentDelivery_Authentication_Token) RevokeAllManagedTokens(cdnAccountId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RevokeAllTokens - This method revokes all tokens belong to a CDN account. Valid media types are and
func (softlayer_network_contentdelivery_authentication_token *SoftLayer_Network_ContentDelivery_Authentication_Token) RevokeAllTokens(cdnAccountId int, mediaType string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RevokeManagedToken - This method is deprecated! Revokes a managed token. If you revoke a token, the
// token will be removed from SoftLayer's system but it will not remove your content on CDN The content
// that requires token validation will not be available to the visitor who is using a revoked token.
func (softlayer_network_contentdelivery_authentication_token *SoftLayer_Network_ContentDelivery_Authentication_Token) RevokeManagedToken(cdnAccountId int, token string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RevokeManagedTokens - This method is deprecated! Deletes multiple managed tokens
func (softlayer_network_contentdelivery_authentication_token *SoftLayer_Network_ContentDelivery_Authentication_Token) RevokeManagedTokens(templateObjects []SoftLayer_Network_ContentDelivery_Authentication_Token) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

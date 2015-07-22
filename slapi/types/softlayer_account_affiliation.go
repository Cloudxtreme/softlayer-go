package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Account_Affiliation - This service allows for a unique identifier to be associated to an
// existing customer account.
type SoftLayer_Account_Affiliation struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// AffiliateId - An affiliate identifier associated with the customer account.
	AffiliateId string `json:"affiliateId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyDate - The date an account affiliation was last modified.
	ModifyDate *time.Time `json:"modifyDate"`
}

// CreateObject - Create a new affiliation to associate with an existing account.
func (softlayer_account_affiliation *SoftLayer_Account_Affiliation) CreateObject(ctx *slapi.RequestContext, templateObject SoftLayer_Account_Affiliation) (*SoftLayer_Account_Affiliation, error) {
	var returnValue *SoftLayer_Account_Affiliation
	return returnValue, nil
}

// DeleteObject - deleteObject permanently removes an account affiliation
func (softlayer_account_affiliation *SoftLayer_Account_Affiliation) DeleteObject(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - Edit an affiliation that is associated to an existing account.
func (softlayer_account_affiliation *SoftLayer_Account_Affiliation) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Account_Affiliation) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetAccountAffiliationsByAffiliateId - Get account affiliation information associated with affiliate
// id.
func (softlayer_account_affiliation *SoftLayer_Account_Affiliation) GetAccountAffiliationsByAffiliateId(ctx *slapi.RequestContext, affiliateId string) ([]*SoftLayer_Account_Affiliation, error) {
	var returnValue []*SoftLayer_Account_Affiliation
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_account_affiliation *SoftLayer_Account_Affiliation) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Account_Affiliation, error) {
	var returnValue *SoftLayer_Account_Affiliation
	return returnValue, nil
}

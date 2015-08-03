package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Affiliation - This service allows for a unique identifier to be associated to an
// existing customer account.
type SoftLayer_Account_Affiliation struct {

	// ModifyDate - The date an account affiliation was last modified.
	ModifyDate *time.Time `json:"modifyDate"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// AffiliateId - An affiliate identifier associated with the customer account.
	AffiliateId string `json:"affiliateId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`
}

// SoftLayer_Account_Affiliation_Extended is SoftLayer_Account_Affiliation with all maskable types.
type SoftLayer_Account_Affiliation_Extended struct {
	SoftLayer_Account_Affiliation

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`
}

func (softlayer_account_affiliation *SoftLayer_Account_Affiliation) String() string {
	return "SoftLayer_Account_Affiliation"
}

package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Affiliation - This service allows for a unique identifier to be associated to an
// existing customer account.
type SoftLayer_Account_Affiliation struct {

	// AccountId - no documentation
	AccountId int `json:"accountId,omitempty"`

	// AffiliateId - An affiliate identifier associated with the customer account.
	AffiliateId string `json:"affiliateId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// ModifyDate - The date an account affiliation was last modified.
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`
}

func (softlayer_account_affiliation *SoftLayer_Account_Affiliation) String() string {
	return "SoftLayer_Account_Affiliation"
}

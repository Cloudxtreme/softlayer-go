package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Link - <nil>
type SoftLayer_Account_Link struct {

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// AccountId - <nil>
	AccountId int `json:"accountId"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// DestinationAccountAlphanumericId - <nil>
	DestinationAccountAlphanumericId string `json:"destinationAccountAlphanumericId"`

	// DestinationAccountId - <nil>
	DestinationAccountId int `json:"destinationAccountId"`

	// Id - <nil>
	Id int `json:"id"`

	// ServiceProvider - <nil>
	ServiceProvider *SoftLayer_Service_Provider `json:"serviceProvider"`

	// ServiceProviderId - <nil>
	ServiceProviderId int `json:"serviceProviderId"`
}

func (softlayer_account_link *SoftLayer_Account_Link) String() string {
	return "SoftLayer_Account_Link"
}

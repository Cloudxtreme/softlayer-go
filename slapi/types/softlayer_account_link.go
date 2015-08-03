package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Link - <nil>
type SoftLayer_Account_Link struct {

	// DestinationAccountAlphanumericId - <nil>
	DestinationAccountAlphanumericId string `json:"destinationAccountAlphanumericId"`

	// DestinationAccountId - <nil>
	DestinationAccountId int `json:"destinationAccountId"`

	// Id - <nil>
	Id int `json:"id"`

	// ServiceProviderId - <nil>
	ServiceProviderId int `json:"serviceProviderId"`

	// AccountId - <nil>
	AccountId int `json:"accountId"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`
}

// SoftLayer_Account_Link_Extended is SoftLayer_Account_Link with all maskable types.
type SoftLayer_Account_Link_Extended struct {
	SoftLayer_Account_Link

	// ServiceProvider - <nil>
	ServiceProvider *SoftLayer_Service_Provider `json:"serviceProvider"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`
}

func (softlayer_account_link *SoftLayer_Account_Link) String() string {
	return "SoftLayer_Account_Link"
}

package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Link - <nil>
type SoftLayer_Account_Link struct {

	// DestinationAccountId - <nil>
	DestinationAccountId int `json:"destinationAccountId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// ServiceProviderId - <nil>
	ServiceProviderId int `json:"serviceProviderId,omitempty"`

	// AccountId - <nil>
	AccountId int `json:"accountId,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// DestinationAccountAlphanumericId - <nil>
	DestinationAccountAlphanumericId string `json:"destinationAccountAlphanumericId,omitempty"`
}

func (softlayer_account_link *SoftLayer_Account_Link) String() string {
	return "SoftLayer_Account_Link"
}

// SoftLayer_Account_Link_Extended is SoftLayer_Account_Link with all maskable types.
type SoftLayer_Account_Link_Extended struct {
	SoftLayer_Account_Link

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// ServiceProvider - <nil>
	ServiceProvider *SoftLayer_Service_Provider `json:"serviceProvider,omitempty"`
}

func (softlayer_account_link *SoftLayer_Account_Link_Extended) String() string {
	return "SoftLayer_Account_Link"
}

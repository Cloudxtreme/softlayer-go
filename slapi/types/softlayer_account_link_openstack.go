package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Link_OpenStack - <nil>
type SoftLayer_Account_Link_OpenStack struct {

	// DomainId - no documentation
	DomainId string `json:"domainId,omitempty"`

	// DestinationAccountAlphanumericId - <nil>
	DestinationAccountAlphanumericId string `json:"destinationAccountAlphanumericId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// ServiceProviderId - <nil>
	ServiceProviderId int `json:"serviceProviderId,omitempty"`

	// AccountId - <nil>
	AccountId int `json:"accountId,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// DestinationAccountId - <nil>
	DestinationAccountId int `json:"destinationAccountId,omitempty"`
}

func (softlayer_account_link_openstack *SoftLayer_Account_Link_OpenStack) String() string {
	return "SoftLayer_Account_Link_OpenStack"
}

// SoftLayer_Account_Link_OpenStack_Extended is SoftLayer_Account_Link_OpenStack with all maskable types.
type SoftLayer_Account_Link_OpenStack_Extended struct {
	SoftLayer_Account_Link_OpenStack

	// ServiceProvider - <nil>
	ServiceProvider *SoftLayer_Service_Provider `json:"serviceProvider,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`
}

func (softlayer_account_link_openstack *SoftLayer_Account_Link_OpenStack_Extended) String() string {
	return "SoftLayer_Account_Link_OpenStack"
}

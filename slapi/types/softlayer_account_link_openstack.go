package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Link_OpenStack - <nil>
type SoftLayer_Account_Link_OpenStack struct {

	// AccountId - <nil>
	AccountId int `json:"accountId,omitempty"`

	// DestinationAccountAlphanumericId - <nil>
	DestinationAccountAlphanumericId string `json:"destinationAccountAlphanumericId,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// DestinationAccountId - <nil>
	DestinationAccountId int `json:"destinationAccountId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// DomainId - no documentation
	DomainId string `json:"domainId,omitempty"`

	// ServiceProviderId - <nil>
	ServiceProviderId int `json:"serviceProviderId,omitempty"`
}

func (softlayer_account_link_openstack *SoftLayer_Account_Link_OpenStack) String() string {
	return "SoftLayer_Account_Link_OpenStack"
}

// SoftLayer_Account_Link_OpenStack_Extended is SoftLayer_Account_Link_OpenStack with all maskable types.
type SoftLayer_Account_Link_OpenStack_Extended struct {
	SoftLayer_Account_Link_OpenStack

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// ServiceProvider - <nil>
	ServiceProvider *SoftLayer_Service_Provider `json:"serviceProvider,omitempty"`
}

func (softlayer_account_link_openstack *SoftLayer_Account_Link_OpenStack_Extended) String() string {
	return "SoftLayer_Account_Link_OpenStack"
}
